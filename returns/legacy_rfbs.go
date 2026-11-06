package returns

import "context"

// V1Empty is the response type used by the retired rFBS return actions.
type V1Empty any

type V2ReturnsRfbsVerifyRequest struct {
	ReturnID                int64  `json:"return_id"`
	ReturnMethodDescription string `json:"return_method_description"`
}

type V2ReturnsRfbsReturnMoneyRequest struct {
	ReturnID         int64 `json:"return_id"`
	ReturnForBackWay int64 `json:"return_for_back_way"`
}

type V2ReturnsRfbsCompensateRequest struct {
	ReturnID           int64  `json:"return_id"`
	CompensationAmount string `json:"compensation_amount"`
}

type V2ReturnsRfbsRejectRequest struct {
	ReturnID          int64  `json:"return_id"`
	Comment           string `json:"comment"`
	RejectionReasonID int64  `json:"rejection_reason_id"`
}

type V2ReturnsRfbsReceiveReturnRequest struct {
	ReturnID int64 `json:"return_id"`
}

func (s *Service) postLegacyRFBS(ctx context.Context, path string, req any) (*V1Empty, error) {
	var resp V1Empty
	if err := s.Client.Post(ctx, path, req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReturnsRfbsVerifyV2 approves an rFBS return request.
//
// Deprecated: this endpoint is absent from the current Ozon Swagger and may no longer work. Use the documented rFBS returns API.
func (s *Service) ReturnsRfbsVerifyV2(ctx context.Context, req *V2ReturnsRfbsVerifyRequest) (*V1Empty, error) {
	return s.postLegacyRFBS(ctx, "/v2/returns/rfbs/verify", req)
}

// ReturnsRfbsReturnMoneyV2 refunds an rFBS return request.
//
// Deprecated: this endpoint is absent from the current Ozon Swagger and may no longer work. Use the documented rFBS returns API.
func (s *Service) ReturnsRfbsReturnMoneyV2(ctx context.Context, req *V2ReturnsRfbsReturnMoneyRequest) (*V1Empty, error) {
	return s.postLegacyRFBS(ctx, "/v2/returns/rfbs/return-money", req)
}

// ReturnsRfbsCompensateV2 partially refunds an rFBS return request.
//
// Deprecated: this endpoint is absent from the current Ozon Swagger and may no longer work. Use the documented rFBS returns API.
func (s *Service) ReturnsRfbsCompensateV2(ctx context.Context, req *V2ReturnsRfbsCompensateRequest) (*V1Empty, error) {
	return s.postLegacyRFBS(ctx, "/v2/returns/rfbs/compensate", req)
}

// ReturnsRfbsRejectV2 rejects an rFBS return request.
//
// Deprecated: this endpoint is absent from the current Ozon Swagger and may no longer work. Use the documented rFBS returns API.
func (s *Service) ReturnsRfbsRejectV2(ctx context.Context, req *V2ReturnsRfbsRejectRequest) (*V1Empty, error) {
	return s.postLegacyRFBS(ctx, "/v2/returns/rfbs/reject", req)
}

// ReturnsRfbsReceiveReturnV2 confirms receipt of an rFBS return.
//
// Deprecated: this endpoint is absent from the current Ozon Swagger and may no longer work. Use the documented rFBS returns API.
func (s *Service) ReturnsRfbsReceiveReturnV2(ctx context.Context, req *V2ReturnsRfbsReceiveReturnRequest) (*V1Empty, error) {
	return s.postLegacyRFBS(ctx, "/v2/returns/rfbs/receive-return", req)
}
