package merchantapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchant"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantCommandResponse struct {
}

func NewMerchantCommandResponseMapper() MerchantCommandResponseMapper {
	return &merchantCommandResponse{}
}

// ToApiResponseMerchant maps a gRPC merchant response to an HTTP API response. It constructs
// an ApiResponseMerchant by copying the status and message fields and mapping the merchant
// data to a MerchantResponse.
//
// Args:
//
//	merchants: A pointer to a pb.ApiResponseMerchant containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchant with mapped data.
func (m *merchantCommandResponse) ToApiResponseMerchant(merchants *pb.ApiResponseMerchant) *response.ApiResponseMerchant {
	return &response.ApiResponseMerchant{
		Status:  merchants.Status,
		Message: merchants.Message,
		Data:    m.mapMerchantResponse(merchants.Data),
	}
}

// mapMerchantResponse maps a gRPC MerchantResponse to an HTTP API response. It
// constructs an ApiResponseMerchant by copying the status, message, ID, name,
// status, API key, user ID, created at, and updated at fields from the gRPC
// response.
//
// Args:
//
//	merchant: A pointer to a pb.MerchantResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponse with the mapped data.
func (m *merchantCommandResponse) mapMerchantResponse(merchant *pb.MerchantResponse) *response.MerchantResponse {
	return &response.MerchantResponse{
		ID:        int(merchant.Id),
		Name:      merchant.Name,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserId),
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}
}
