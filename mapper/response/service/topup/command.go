package topupservicemapper

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type topupCommandResponseMapper struct {
}

func NewTopupCommandResponseMapper() TopupCommandResponseMapper {
	return &topupCommandResponseMapper{}
}

// ToTopupResponse maps a single TopupRecord domain model to a TopupResponse API-compatible response type.
// Args:
//   - topup: A pointer to a TopupRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TopupResponse containing the mapped data, including
//     ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *topupCommandResponseMapper) ToTopupResponse(topup *record.TopupRecord) *response.TopupResponse {
	return &response.TopupResponse{
		ID:          topup.ID,
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: topup.TopupAmount,
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}
