package report

import (
	"context"
	"net/http"
	"testing"

	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

func TestFinanceCashFlowStatementList_DecodesPaymentArray(t *testing.T) {
	client, server := transport.NewTestClient(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v1/finance/cash-flow-statement/list" {
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"result":{"details":[{"payments":[{"currency_code":"RUB","payment":42.5}]}],"cash_flows":[],"page_count":1}}`))
	})
	defer server.Close()

	service := &Service{Client: client}
	response, err := service.FinanceCashFlowStatementList(context.Background(), &V3FinanceCashFlowStatementListRequest{WithDetails: true})
	if err != nil {
		t.Fatalf("decode cash-flow response: %v", err)
	}
	if len(response.Result.Details) != 1 {
		t.Fatalf("details length = %d, want 1", len(response.Result.Details))
	}
	payments := response.Result.Details[0].Payments
	if len(payments) != 1 {
		t.Fatalf("payments = %v, want one-element slice", response.Result.Details[0].Payments)
	}
	if got := payments[0]; got.CurrencyCode != "RUB" || got.Payment != 42.5 {
		t.Fatalf("payment = %+v, want RUB 42.5", got)
	}
}
