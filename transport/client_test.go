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
