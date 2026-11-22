package prices

import (
	"encoding/json"
	"testing"
)

func TestGetProductInfoPricesMarketingActionValue(t *testing.T) {
	for _, tc := range []struct {
		name  string
		value string
		want  float64
	}{
		{name: "fractional", value: "46.68", want: 46.68},
		{name: "integer", value: "46", want: 46},
	} {
		t.Run(tc.name, func(t *testing.T) {
			body := []byte(`{"items":[{"price":{"price":100.5},"marketing_actions":{"actions":[{"value":` + tc.value + `}]}}]}`)
			var response V5GetProductInfoPricesV5Response
			if err := json.Unmarshal(body, &response); err != nil {
				t.Fatal(err)
			}
			if len(response.Items) != 1 || len(response.Items[0].MarketingActions.Actions) != 1 {
				t.Fatalf("unexpected response shape: %+v", response)
			}
			if got := response.Items[0].MarketingActions.Actions[0].Value; got != tc.want {
				t.Errorf("marketing action value = %v, want %v", got, tc.want)
			}
			if got := response.Items[0].Price.Price; got != 100.5 {
				t.Errorf("price = %v, want 100.5", got)
			}
		})
	}
}
