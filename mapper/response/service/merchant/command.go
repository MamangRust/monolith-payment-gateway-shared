package merchantservicemapper

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// MerchantCommandResponseMapper provides methods to map MerchantRecord domain models to MerchantResponse API-compatible response types for command operations.
type merchantCommandResponseMapper struct{}

// NewMerchantCommandResponseMapper returns a new instance of merchantCommandResponseMapper, which provides methods to map
// MerchantRecord domain models to MerchantResponse API-compatible response types for command operations.
func NewMerchantCommandResponseMapper() MerchantCommandResponseMapper {
	return &merchantCommandResponseMapper{}
}

// ToMerchantResponse maps a single MerchantRecord to a MerchantResponse API-compatible response.
// Args:
//   - merchant: A pointer to a MerchantRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponse containing the mapped data, including
//     ID, Name, UserID, Status, ApiKey, CreatedAt, and UpdatedAt.
func (s *merchantCommandResponseMapper) ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse {
	return &response.MerchantResponse{
		ID:        merchant.ID,
		Name:      merchant.Name,
		UserID:    merchant.UserID,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}
}
