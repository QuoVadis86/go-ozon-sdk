package category

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"os"
	"testing"
)

var ctx = context.Background()

func skipNoCreds(t *testing.T) *transport.Client {
	t.Helper()
	if os.Getenv("OZON_CLIENT_ID") == "" || os.Getenv("OZON_API_KEY") == "" {
		t.Skip("set OZON_CLIENT_ID and OZON_API_KEY to run tests")
	}
	return transport.New(os.Getenv("OZON_CLIENT_ID"), os.Getenv("OZON_API_KEY"), nil)
}

func TestGetAttributes(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetAttributes(ctx, &V1GetAttributesRequest{})
	if err != nil {
		t.Fatalf("GetAttributes() error: %v", err)
	}
	_ = resp
}

func TestGetTree(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetTree(ctx, &V1GetTreeRequest{})
	if err != nil {
		t.Fatalf("GetTree() error: %v", err)
	}
	_ = resp
}

func TestSearchAttributeValues(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SearchAttributeValues(ctx, &V1SearchAttributeValuesRequest{})
	if err != nil {
		t.Fatalf("SearchAttributeValues() error: %v", err)
	}
	_ = resp
}

func TestGetAttributeValues(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.GetAttributeValues(ctx, &V1GetAttributeValuesRequest{})
	if err != nil {
		t.Fatalf("GetAttributeValues() error: %v", err)
	}
	_ = resp
}
