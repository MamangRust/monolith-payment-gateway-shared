package transactionservicemapper

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transactionCommandResponseMapper struct {
}

func NewTransactionCommandResponseMapper() TransactionCommandResponseMapper {
	return &transactionCommandResponseMapper{}
}

// ToTransactionResponse converts a single transaction record into a TransactionResponse.
//
// Args:
//   - transaction: A pointer to a TransactionRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransactionResponse containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, and UpdatedAt.
func (s *transactionCommandResponseMapper) ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse {
	return &response.TransactionResponse{
		ID:              transaction.ID,
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          transaction.Amount,
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      transaction.MerchantID,
		TransactionTime: transaction.TransactionTime,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}
}
