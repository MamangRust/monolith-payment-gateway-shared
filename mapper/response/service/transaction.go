package responseservice

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// transactionResponseMapper provides methods to map TransactionRecord domain models to TransactionResponse API-compatible response types.
type transactionResponseMapper struct {
}

// NewTransactionResponseMapper returns a new instance of transactionResponseMapper,
// which provides methods to map TransactionRecord domain models to API-compatible
// TransactionResponse types.
func NewTransactionResponseMapper() *transactionResponseMapper {
	return &transactionResponseMapper{}
}

// ToTransactionResponse converts a single transaction record into a TransactionResponse.
//
// Args:
//   - transaction: A pointer to a TransactionRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransactionResponse containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, and UpdatedAt.
func (s *transactionResponseMapper) ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse {
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

// ToTransactionsResponse converts multiple transaction records into a slice of TransactionResponse.
//
// Args:
//   - transactions: A slice of pointers to TransactionRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponse containing the mapped data, including ID, TransactionNo, CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, and UpdatedAt.
func (s *transactionResponseMapper) ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse {
	responses := make([]*response.TransactionResponse, 0, len(transactions))
	for _, transaction := range transactions {
		responses = append(responses, s.ToTransactionResponse(transaction))
	}
	return responses
}

// ToTransactionResponseDeleteAt maps a soft-deleted transaction record into a TransactionResponseDeleteAt.
//
// Args:
//   - transaction: A pointer to a TransactionRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransactionResponseDeleteAt containing the mapped data,
//     including ID, TransactionNo, CardNumber, Amount, PaymentMethod, MerchantID,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionResponseMapper) ToTransactionResponseDeleteAt(transaction *record.TransactionRecord) *response.TransactionResponseDeleteAt {
	return &response.TransactionResponseDeleteAt{
		ID:              transaction.ID,
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          transaction.Amount,
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      transaction.MerchantID,
		TransactionTime: transaction.TransactionTime,
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
		DeletedAt:       transaction.DeletedAt,
	}
}

// ToTransactionsResponseDeleteAt converts multiple soft-deleted transaction records into a slice of TransactionResponseDeleteAt.
//
// Args:
//   - transactions: A slice of pointers to TransactionRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponseDeleteAt containing the mapped data, including
//     ID, TransactionNo, CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt,
//     UpdatedAt, and DeletedAt.
func (s *transactionResponseMapper) ToTransactionsResponseDeleteAt(transactions []*record.TransactionRecord) []*response.TransactionResponseDeleteAt {
	responses := make([]*response.TransactionResponseDeleteAt, 0, len(transactions))

	for _, transaction := range transactions {
		responses = append(responses, s.ToTransactionResponseDeleteAt(transaction))
	}
	return responses
}

// ToTransactionResponseMonthStatusSuccess maps a record of monthly transaction success status
// into a TransactionResponseMonthStatusSuccess.
//
// Args:
//   - s: A pointer to TransactionRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionResponseMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionResponseMonthStatusSuccess(s *record.TransactionRecordMonthStatusSuccess) *response.TransactionResponseMonthStatusSuccess {
	return &response.TransactionResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTransactionResponsesMonthStatusSuccess converts multiple records of monthly transaction success status into a slice of
// TransactionResponseMonthStatusSuccess.
//
// Args:
//   - Transactions: A slice of pointers to TransactionRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponseMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionResponsesMonthStatusSuccess(Transactions []*record.TransactionRecordMonthStatusSuccess) []*response.TransactionResponseMonthStatusSuccess {
	var TransactionRecords []*response.TransactionResponseMonthStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionResponseMonthStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// ToTransactionResponseYearStatusSuccess maps a record of yearly transaction success status
// into a TransactionResponseYearStatusSuccess.
//
// Args:
//   - s: A pointer to TransactionRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionResponseYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionResponseYearStatusSuccess(s *record.TransactionRecordYearStatusSuccess) *response.TransactionResponseYearStatusSuccess {
	return &response.TransactionResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTransactionResponsesYearStatusSuccess converts multiple records of yearly transaction success status into a slice of
// TransactionResponseYearStatusSuccess.
//
// Args:
//   - Transactions: A slice of pointers to TransactionRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponseYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionResponsesYearStatusSuccess(Transactions []*record.TransactionRecordYearStatusSuccess) []*response.TransactionResponseYearStatusSuccess {
	var TransactionRecords []*response.TransactionResponseYearStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionResponseYearStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// ToTransactionResponseMonthStatusFailed maps a record of monthly transaction failed status into a TransactionResponseMonthStatusFailed.
//
// Args:
//   - s: A pointer to TransactionRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionResponseMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionResponseMonthStatusFailed(s *record.TransactionRecordMonthStatusFailed) *response.TransactionResponseMonthStatusFailed {
	return &response.TransactionResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTransactionResponsesMonthStatusFailed converts multiple records of monthly transaction failed status into a slice of
// TransactionResponseMonthStatusFailed.
//
// Args:
//   - Transactions: A slice of pointers to TransactionRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponseMonthStatusFailed containing the mapped data, including Year, Month, TotalFailed, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionResponsesMonthStatusFailed(Transactions []*record.TransactionRecordMonthStatusFailed) []*response.TransactionResponseMonthStatusFailed {
	var TransactionRecords []*response.TransactionResponseMonthStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionResponseMonthStatusFailed(Transaction))
	}

	return TransactionRecords
}

// ToTransactionResponseYearStatusFailed maps a yearly transaction failed status record
// into a TransactionResponseYearStatusFailed.
//
// Args:
//   - s: A pointer to TransactionRecordYearStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionResponseYearStatusFailed(s *record.TransactionRecordYearStatusFailed) *response.TransactionResponseYearStatusFailed {
	return &response.TransactionResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTransactionResponsesYearStatusFailed converts multiple yearly transaction failed status records
// into a slice of TransactionResponseYearStatusFailed.
//
// Args:
//   - Transactions: A slice of pointers to TransactionRecordYearStatusFailed containing the records to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionResponsesYearStatusFailed(Transactions []*record.TransactionRecordYearStatusFailed) []*response.TransactionResponseYearStatusFailed {
	var TransactionRecords []*response.TransactionResponseYearStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionResponseYearStatusFailed(Transaction))
	}

	return TransactionRecords
}

// ToTransactionMonthlyMethodResponse maps a record of monthly transaction methods into a TransactionMonthMethodResponse.
//
// Args:
//   - s: A pointer to TransactionMonthMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionMonthMethodResponse containing the mapped data, including Month,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionMonthlyMethodResponse(s *record.TransactionMonthMethod) *response.TransactionMonthMethodResponse {
	return &response.TransactionMonthMethodResponse{
		Month:             s.Month,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int(s.TotalTransactions),
		TotalAmount:       int(s.TotalAmount),
	}
}

// ToTransactionMonthlyMethodResponses maps a slice of TransactionMonthMethod domain models
// into a slice of TransactionMonthMethodResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransactionMonthMethod containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionMonthMethodResponse containing the mapped data, including
//     Month, PaymentMethod, TotalTransactions, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionMonthlyMethodResponses(s []*record.TransactionMonthMethod) []*response.TransactionMonthMethodResponse {
	var transactionResponses []*response.TransactionMonthMethodResponse
	for _, transaction := range s {
		transactionResponses = append(transactionResponses, t.ToTransactionMonthlyMethodResponse(transaction))
	}
	return transactionResponses
}

// ToTransactionYearlyMethodResponse maps a single yearly transaction method record into a TransactionYearMethodResponse.
//
// Args:
//   - s: A pointer to TransactionYearMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionYearMethodResponse containing the mapped data, including Year,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionYearlyMethodResponse(s *record.TransactionYearMethod) *response.TransactionYearMethodResponse {
	return &response.TransactionYearMethodResponse{
		Year:              s.Year,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int(s.TotalTransactions),
		TotalAmount:       int(s.TotalAmount),
	}
}

// ToTransactionYearlyMethodResponses maps a slice of TransactionYearMethod domain models
// into a slice of TransactionYearMethodResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransactionYearMethod containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionYearMethodResponse containing the mapped data, including
//     Year, PaymentMethod, TotalTransactions, and TotalAmount.
func (t *transactionResponseMapper) ToTransactionYearlyMethodResponses(s []*record.TransactionYearMethod) []*response.TransactionYearMethodResponse {
	var transactionResponses []*response.TransactionYearMethodResponse
	for _, transaction := range s {
		transactionResponses = append(transactionResponses, t.ToTransactionYearlyMethodResponse(transaction))
	}
	return transactionResponses
}

// ToTransactionMonthlyAmountResponse converts a monthly amount record into a TransactionMonthAmountResponse.
//
// Args:
//   - s: A pointer to TransactionMonthAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionMonthAmountResponse containing the mapped data, including Month and TotalAmount.
func (t *transactionResponseMapper) ToTransactionMonthlyAmountResponse(s *record.TransactionMonthAmount) *response.TransactionMonthAmountResponse {
	return &response.TransactionMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionMonthlyAmountResponses maps a slice of TransactionMonthAmount domain models
// into a slice of TransactionMonthAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransactionMonthAmount containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionMonthAmountResponse containing the mapped data, including
//     Month and TotalAmount.
func (t *transactionResponseMapper) ToTransactionMonthlyAmountResponses(s []*record.TransactionMonthAmount) []*response.TransactionMonthAmountResponse {
	var transactionResponses []*response.TransactionMonthAmountResponse
	for _, transaction := range s {
		transactionResponses = append(transactionResponses, t.ToTransactionMonthlyAmountResponse(transaction))
	}
	return transactionResponses
}

// ToTransactionYearlyAmountResponse maps a yearly amount record into a TransactionYearlyAmountResponse.
//
// Args:
//   - s: A pointer to TransactionYearlyAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to TransactionYearlyAmountResponse containing the mapped data, including Year and TotalAmount.
func (t *transactionResponseMapper) ToTransactionYearlyAmountResponse(s *record.TransactionYearlyAmount) *response.TransactionYearlyAmountResponse {
	return &response.TransactionYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionYearlyAmountResponses maps a slice of TransactionYearlyAmount domain models
// into a slice of TransactionYearlyAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransactionYearlyAmount containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransactionYearlyAmountResponse containing the mapped data, including
//     Year and TotalAmount.
func (t *transactionResponseMapper) ToTransactionYearlyAmountResponses(s []*record.TransactionYearlyAmount) []*response.TransactionYearlyAmountResponse {
	var transactionResponses []*response.TransactionYearlyAmountResponse
	for _, transaction := range s {
		transactionResponses = append(transactionResponses, t.ToTransactionYearlyAmountResponse(transaction))
	}
	return transactionResponses
}
