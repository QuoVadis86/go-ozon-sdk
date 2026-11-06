package ozon

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/QuoVadis86/go-ozon-sdk/product"
)

func TestClient_UsesLanguageOnlyForSelectedRequest(t *testing.T) {
	headers := make(chan string, 3)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v3/product/info/list" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		headers <- r.Header.Get("x-o3-language")
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"items":[]}`))
	}))
	defer server.Close()

	client := NewClient("client-id", "api-key", &ClientOptions{BaseURL: server.URL})
	ctx := context.Background()
	req := &product.V3GetProductInfoListRequest{}
	for _, call := range []struct {
		ctx  context.Context
		want string
	}{
		{ctx: ctx},
		{ctx: WithLanguage(ctx, LanguageZHHans), want: "zh-Hans"},
		{ctx: ctx},
	} {
		if _, err := client.Product.GetProductInfoList(call.ctx, req); err != nil {
			t.Fatalf("GetProductInfoList() error: %v", err)
		}
		if got := <-headers; got != call.want {
			t.Errorf("x-o3-language = %q, want %q", got, call.want)
		}
	}
}
