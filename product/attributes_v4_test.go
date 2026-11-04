package product

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

func TestGetProductAttributesV4_DecodesResponseVariants(t *testing.T) {
	tests := []struct {
		name      string
		body      string
		wantID    int64
		wantSKU   string
		wantURL   string
		wantImage GetProductAttributesResponseImage
	}{
		{
			name:    "numeric SKU with URL image",
			body:    `{"result":[{"id":1,"sku":123456,"images":["https://example.invalid/product.jpg"]}]}`,
			wantID:  1,
			wantSKU: "123456",
			wantURL: "https://example.invalid/product.jpg",
		},
		{
			name:      "numeric SKU with object image",
			body:      `{"result":[{"sku":123456,"images":[{"default":true,"file_name":"product.jpg","index":2}]}]}`,
			wantSKU:   "123456",
			wantImage: GetProductAttributesResponseImage{Default: true, FileName: "product.jpg", Index: 2},
		},
		{
			name:    "string SKU with URL image",
			body:    `{"result":[{"sku":"123456","images":["https://example.invalid/product.jpg"]}]}`,
			wantSKU: "123456",
			wantURL: "https://example.invalid/product.jpg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost || r.URL.Path != "/v4/product/info/attributes" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
					http.NotFound(w, r)
					return
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, tt.body)
			}))
			defer server.Close()

			svc := &Service{Client: transport.New("client-id", "api-key", &transport.Options{BaseURL: server.URL})}
			resp, err := svc.GetProductAttributesV4(context.Background(), &V4GetProductAttributesV4Request{})
			if err != nil {
				t.Fatalf("GetProductAttributesV4() error: %v", err)
			}
			if len(resp.Result) != 1 || len(resp.Result[0].Images) != 1 {
				t.Fatalf("unexpected result: %+v", resp.Result)
			}
			result := resp.Result[0]
			if result.ID != tt.wantID {
				t.Errorf("ID = %d, want %d", result.ID, tt.wantID)
			}
			if result.SKU != tt.wantSKU {
				t.Errorf("SKU = %q, want %q", result.SKU, tt.wantSKU)
			}
			if result.Images[0].URL != tt.wantURL {
				t.Errorf("image URL = %q, want %q", result.Images[0].URL, tt.wantURL)
			}
			if result.Images[0].Default != tt.wantImage.Default || result.Images[0].FileName != tt.wantImage.FileName || result.Images[0].Index != tt.wantImage.Index {
				t.Errorf("image = %+v, want %+v", result.Images[0], tt.wantImage)
			}
		})
	}
}
