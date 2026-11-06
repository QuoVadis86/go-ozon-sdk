package transport

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientPost_SendsConfiguredLanguage(t *testing.T) {
	tests := []struct {
		name       string
		language   Language
		wantHeader string
	}{
		{name: "default omits language"},
		{name: "simplified Chinese", language: LanguageZHHans, wantHeader: "zh-Hans"},
		{name: "English", language: LanguageEN, wantHeader: "EN"},
		{name: "Russian", language: LanguageRU, wantHeader: "RU"},
		{name: "Turkish", language: LanguageTR, wantHeader: "TR"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost || r.URL.Path != "/v3/product/info/list" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				if got := r.Header.Get("x-o3-language"); got != tt.wantHeader {
					t.Errorf("x-o3-language = %q, want %q", got, tt.wantHeader)
				}
				_, present := r.Header[http.CanonicalHeaderKey("x-o3-language")]
				if present != (tt.wantHeader != "") {
					t.Errorf("x-o3-language presence = %t, want %t", present, tt.wantHeader != "")
				}
				if got := r.Header.Get(HeaderClientID); got != "client-id" {
					t.Errorf("Client-Id = %q, want %q", got, "client-id")
				}
				if got := r.Header.Get(HeaderAPIKey); got != "api-key" {
					t.Errorf("Api-Key = %q, want %q", got, "api-key")
				}
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			client := New("client-id", "api-key", &Options{BaseURL: server.URL, Language: tt.language})
			if err := client.Post(context.Background(), "/v3/product/info/list", nil, nil); err != nil {
				t.Fatalf("Post() error: %v", err)
			}
		})
	}
}

func TestClientPost_RejectsUnsupportedLanguage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Error("request sent with unsupported language")
	}))
	defer server.Close()

	client := New("client-id", "api-key", &Options{BaseURL: server.URL, Language: Language(255)})
	if err := client.Post(context.Background(), "/v3/product/info/list", nil, nil); err == nil {
		t.Fatal("Post() succeeded with unsupported language")
	}
}

func TestClientPost_OverridesLanguageWithoutChangingClientDefault(t *testing.T) {
	tests := []struct {
		name            string
		defaultLanguage Language
		wantDefault     string
	}{
		{name: "client without language"},
		{name: "client with English default", defaultLanguage: LanguageEN, wantDefault: "EN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := make(chan string, 4)
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				headers <- r.Header.Get(HeaderLanguage)
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			client := New("client-id", "api-key", &Options{BaseURL: server.URL, Language: tt.defaultLanguage})
			base := context.Background()
			requests := []struct {
				ctx  context.Context
				want string
			}{
				{ctx: base, want: tt.wantDefault},
				{ctx: WithLanguage(base, LanguageZHHans), want: "zh-Hans"},
				{ctx: base, want: tt.wantDefault},
				{ctx: WithLanguage(base, LanguageDefault)},
			}

			for _, request := range requests {
				if err := client.Post(request.ctx, "/v3/product/info/list", nil, nil); err != nil {
					t.Fatalf("Post() error: %v", err)
				}
				if got := <-headers; got != request.want {
					t.Errorf("x-o3-language = %q, want %q", got, request.want)
				}
			}
		})
	}
}

func TestClientPost_UsesSeparateLanguagesForConcurrentRequests(t *testing.T) {
	type receivedHeader struct {
		path     string
		language string
	}
	headers := make(chan receivedHeader, 2)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers <- receivedHeader{path: r.URL.Path, language: r.Header.Get(HeaderLanguage)}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New("client-id", "api-key", &Options{BaseURL: server.URL})
	start := make(chan struct{})
	errors := make(chan error, 2)
	for _, request := range []struct {
		path     string
		language Language
	}{
		{path: "/zh", language: LanguageZHHans},
		{path: "/en", language: LanguageEN},
	} {
		go func(path string, language Language) {
			<-start
			errors <- client.Post(WithLanguage(context.Background(), language), path, nil, nil)
		}(request.path, request.language)
	}
	close(start)

	for i := 0; i < 2; i++ {
		if err := <-errors; err != nil {
			t.Fatalf("Post() error: %v", err)
		}
	}
	want := map[string]string{"/zh": "zh-Hans", "/en": "EN"}
	for i := 0; i < 2; i++ {
		got := <-headers
		if got.language != want[got.path] {
			t.Errorf("%s x-o3-language = %q, want %q", got.path, got.language, want[got.path])
		}
	}
}
