package withdrawapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/withdraw"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// WithdrawCommandResponseMapper provides methods to map gRPC withdraw responses to HTTP API responses
type withdrawCommandResponseMapper struct {
}

// NewWithdrawCommandResponseMapper creates and returns a new instance of withdrawCommandResponseMapper.
// This mapper provides methods to convert withdraw gRPC responses into API-compatible responses for command operations.
func NewWithdrawCommandResponseMapper() WithdrawCommandResponseMapper {
	return &withdrawCommandResponseMapper{}
}

// ToApiResponseWithdraw maps a single withdraw gRPC response to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponseWithdraw containing the mapped data.
func (m *withdrawCommandResponseMapper) ToApiResponseWithdraw(pbResponse *pb.ApiResponseWithdraw) *response.ApiResponseWithdraw {
	return &response.ApiResponseWithdraw{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseWithdrawal(pbResponse.Data),
	}
}

// mapResponseWithdrawal maps a single withdraw gRPC response to an API response.
//
// Args:
//   - withdraw: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponse containing the mapped data.
func (w *withdrawCommandResponseMapper) mapResponseWithdrawal(withdraw *pb.WithdrawResponse) *response.WithdrawResponse {
	return &response.WithdrawResponse{
		ID:             int(withdraw.WithdrawId),
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
	}
}
