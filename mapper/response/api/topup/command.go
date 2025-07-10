package topupapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/topup"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type topupCommandResponseMapper struct{}

func NewTopupCommandResponseMapper() TopupCommandResponseMapper {
	return &topupCommandResponseMapper{}
}

// ToApiResponseTopup maps a single gRPC top-up response to an HTTP API response.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopup containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopup containing the mapped data, including status, message, and a single
//     mapped top-up response.
func (t *topupCommandResponseMapper) ToApiResponseTopup(s *pb.ApiResponseTopup) *response.ApiResponseTopup {
	return &response.ApiResponseTopup{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopup(s.Data),
	}
}

// mapResponseTopup maps a single gRPC top-up response to an HTTP API response.
// It constructs a TopupResponse by copying the fields from the gRPC response.
//
// Args:
//   - topup: A pointer to a pb.TopupResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TopupResponse with mapped data.
func (t *topupCommandResponseMapper) mapResponseTopup(topup *pb.TopupResponse) *response.TopupResponse {
	return &response.TopupResponse{
		ID:          int(topup.Id),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}
