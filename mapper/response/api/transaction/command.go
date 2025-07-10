package transactionapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transaction"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transactionCommandResponseMapper struct{}

func NewTransactionCommandResponseMapper() *transactionCommandResponseMapper {
	return &transactionCommandResponseMapper{}
}

// ToApiResponseTransaction maps a single gRPC transaction response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransaction containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransaction containing the mapped data, including status, message,
//     and a single mapped transaction response.
func (m *transactionCommandResponseMapper) ToApiResponseTransaction(pbResponse *pb.ApiResponseTransaction) *response.ApiResponseTransaction {
	return &response.ApiResponseTransaction{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransaction(pbResponse.Data),
	}
}

// mapResponseTransaction maps a single gRPC transaction response to an HTTP API response.
//
// Args:
//   - transaction: A pointer to a pb.TransactionResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponse containing the mapped data.
func (m *transactionCommandResponseMapper) mapResponseTransaction(transaction *pb.TransactionResponse) *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:              int(transaction.Id),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantID:      int(transaction.MerchantId),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}
}
