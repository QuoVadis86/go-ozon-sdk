package product

import (
	"context"
	"github.com/QuoVadis86/go-ozon-sdk/transport"
	"os"
	"testing"
)

func TestIntegration_GetUploadQuota(t *testing.T) {
	clientID := os.Getenv("OZON_CLIENT_ID")
	apiKey := os.Getenv("OZON_API_KEY")
	if clientID == "" || apiKey == "" {
		t.Skip("set OZON_CLIENT_ID and OZON_API_KEY to run integration tests")
	}
	cl := transport.New(clientID, apiKey, nil)
	svc := &Service{Client: cl}
	ctx := context.Background()
	resp, err := svc.GetUploadQuota(ctx)
	if err != nil {
		t.Fatalf("GetUploadQuota() error: %v", err)
	}
	if resp == nil {
		t.Fatal("GetUploadQuota() returned nil")
	}
}
