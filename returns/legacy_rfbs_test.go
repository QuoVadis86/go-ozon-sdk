package returns

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/QuoVadis86/go-ozon-sdk/transport"
)

func TestLegacyRFBSMethodsKeepTheirPaths(t *testing.T) {
	tests := []struct {
		name string
		path string
		call func(*Service) (*V1Empty, error)
	}{
		{"verify", "/v2/returns/rfbs/verify", func(s *Service) (*V1Empty, error) {
			return s.ReturnsRfbsVerifyV2(context.Background(), &V2ReturnsRfbsVerifyRequest{ReturnID: 42})
		}},
		{"return-money", "/v2/returns/rfbs/return-money", func(s *Service) (*V1Empty, error) {
			return s.ReturnsRfbsReturnMoneyV2(context.Background(), &V2ReturnsRfbsReturnMoneyRequest{ReturnID: 42})
		}},
		{"compensate", "/v2/returns/rfbs/compensate", func(s *Service) (*V1Empty, error) {
			return s.ReturnsRfbsCompensateV2(context.Background(), &V2ReturnsRfbsCompensateRequest{ReturnID: 42})
		}},
		{"reject", "/v2/returns/rfbs/reject", func(s *Service) (*V1Empty, error) {
			return s.ReturnsRfbsRejectV2(context.Background(), &V2ReturnsRfbsRejectRequest{ReturnID: 42})
		}},
		{"receive-return", "/v2/returns/rfbs/receive-return", func(s *Service) (*V1Empty, error) {
			return s.ReturnsRfbsReceiveReturnV2(context.Background(), &V2ReturnsRfbsReceiveReturnRequest{ReturnID: 42})
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPost || r.URL.Path != tt.path {
					t.Errorf("request = %s %s, want POST %s", r.Method, r.URL.Path, tt.path)
				}
				var body struct {
					ReturnID int64 `json:"return_id"`
				}
				if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.ReturnID != 42 {
					t.Errorf("request body = %+v, err = %v", body, err)
				}
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte("{}"))
			}))
			defer server.Close()
			svc := &Service{Client: transport.New("id", "key", &transport.Options{BaseURL: server.URL})}
			if response, err := tt.call(svc); err != nil || response == nil {
				t.Fatalf("response = %v, err = %v", response, err)
			}
		})
	}
}
