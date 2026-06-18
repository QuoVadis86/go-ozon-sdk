package seller

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

func TestSellerInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerInfo(ctx)
	if err != nil {
		t.Fatalf("SellerInfo() error: %v", err)
	}
	_ = resp
}

func TestRolesByToken(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.RolesByToken(ctx)
	if err != nil {
		t.Fatalf("RolesByToken() error: %v", err)
	}
	_ = resp
}

func TestSellerOzonLogisticsInfo(t *testing.T) {
	cl := skipNoCreds(t)
	svc := &Service{Client: cl}
	resp, err := svc.SellerOzonLogisticsInfo(ctx)
	if err != nil {
		t.Fatalf("SellerOzonLogisticsInfo() error: %v", err)
	}
	_ = resp
}
