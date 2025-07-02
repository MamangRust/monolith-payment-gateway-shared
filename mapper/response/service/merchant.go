package responseservice

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// merchantResponseMapper provides methods to map MerchantRecord domain models to MerchantResponse API-compatible response types.
type merchantResponseMapper struct{}

// NewMerchantResponseMapper returns a new instance of merchantResponse, a type that
// defines methods for mapping gRPC merchant-related responses to HTTP API responses.

func NewMerchantResponseMapper() *merchantResponseMapper {
	return &merchantResponseMapper{}
}

// ToMerchantResponse maps a single MerchantRecord to a MerchantResponse API-compatible response.
// Args:
//   - merchant: A pointer to a MerchantRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponse containing the mapped data, including
//     ID, Name, UserID, Status, ApiKey, CreatedAt, and UpdatedAt.
func (s *merchantResponseMapper) ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse {
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

// ToMerchantsResponse maps multiple MerchantRecords to a slice of MerchantResponse API-compatible responses.
// Args:
//   - merchants: A slice of pointers to MerchantRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponse containing the mapped data, including
//     ID, Name, UserID, Status, ApiKey, CreatedAt, and UpdatedAt.
func (s *merchantResponseMapper) ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse {
	var response []*response.MerchantResponse
	for _, merchant := range merchants {
		response = append(response, s.ToMerchantResponse(merchant))
	}
	return response
}

// ToMerchantResponseDeleteAt maps a MerchantRecord to a MerchantResponseDeleteAt,
// which includes additional deletion data.
// This function is useful for handling soft-deleted merchants where the deletion
// timestamp is relevant.
//
// Args:
//   - merchant: A pointer to a MerchantRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseDeleteAt containing the mapped data,
//     including ID, Name, UserID, Status, ApiKey, CreatedAt, UpdatedAt, and DeletedAt.
func (s *merchantResponseMapper) ToMerchantResponseDeleteAt(merchant *record.MerchantRecord) *response.MerchantResponseDeleteAt {
	return &response.MerchantResponseDeleteAt{
		ID:        merchant.ID,
		Name:      merchant.Name,
		UserID:    merchant.UserID,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
		DeletedAt: merchant.DeletedAt,
	}
}

// ToMerchantsResponseDeleteAt maps multiple soft-deleted MerchantRecords to their corresponding responses.
// Args:
//   - merchants: A slice of pointers to MerchantRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseDeleteAt containing the mapped data, including
//     ID, Name, UserID, Status, ApiKey, CreatedAt, UpdatedAt, and DeletedAt.
func (s *merchantResponseMapper) ToMerchantsResponseDeleteAt(merchants []*record.MerchantRecord) []*response.MerchantResponseDeleteAt {
	var response []*response.MerchantResponseDeleteAt
	for _, merchant := range merchants {
		response = append(response, s.ToMerchantResponseDeleteAt(merchant))
	}
	return response
}

// ToMerchantTransactionResponse maps a single MerchantTransactionsRecord to a MerchantTransactionResponse API-compatible response.
//
// Args:
//   - merchant: A pointer to a MerchantTransactionsRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantTransactionResponse containing the mapped data, including
//     ID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName, TransactionTime, CreatedAt, and UpdatedAt.
func (m *merchantResponseMapper) ToMerchantTransactionResponse(merchant *record.MerchantTransactionsRecord) *response.MerchantTransactionResponse {

	return &response.MerchantTransactionResponse{
		ID:              int(merchant.TransactionID),
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime.Format("2006-01-02"),
		CreatedAt:       merchant.CreatedAt,
		UpdatedAt:       merchant.UpdatedAt,
	}
}

// ToMerchantsTransactionResponse maps multiple MerchantTransactionsRecords to a slice of MerchantTransactionResponse API-compatible responses.
// Args:
//   - merchants: A slice of pointers to MerchantTransactionsRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantTransactionResponse containing the mapped data, including
//     ID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName, TransactionTime, CreatedAt, and UpdatedAt.
func (m *merchantResponseMapper) ToMerchantsTransactionResponse(merchants []*record.MerchantTransactionsRecord) []*response.MerchantTransactionResponse {
	var records []*response.MerchantTransactionResponse
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTransactionResponse(merchant))
	}
	return records
}

func (s *merchantResponseMapper) ToMerchantMonthlyPaymentMethod(ms *record.MerchantMonthlyPaymentMethod) *response.MerchantResponseMonthlyPaymentMethod {
	return &response.MerchantResponseMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   ms.TotalAmount,
	}
}

// ToMerchantMonthlyPaymentMethods maps multiple monthly payment method records into a slice of MerchantResponseMonthlyPaymentMethod.
//
// Args:
//   - ms: A slice of pointers to MerchantMonthlyPaymentMethod records containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (s *merchantResponseMapper) ToMerchantMonthlyPaymentMethods(ms []*record.MerchantMonthlyPaymentMethod) []*response.MerchantResponseMonthlyPaymentMethod {
	var response []*response.MerchantResponseMonthlyPaymentMethod
	for _, merchant := range ms {
		response = append(response, s.ToMerchantMonthlyPaymentMethod(merchant))
	}
	return response
}

// ToMerchantYearlyPaymentMethod maps a single yearly payment method record to a
// MerchantResponseYearlyPaymentMethod API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantYearlyPaymentMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (s *merchantResponseMapper) ToMerchantYearlyPaymentMethod(ms *record.MerchantYearlyPaymentMethod) *response.MerchantResponseYearlyPaymentMethod {
	return &response.MerchantResponseYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   ms.TotalAmount,
	}
}

// ToMerchantYearlyPaymentMethods maps multiple yearly payment method records into a slice of MerchantResponseYearlyPaymentMethod.
//
// Args:
//   - ms: A slice of pointers to MerchantYearlyPaymentMethod records containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (s *merchantResponseMapper) ToMerchantYearlyPaymentMethods(ms []*record.MerchantYearlyPaymentMethod) []*response.MerchantResponseYearlyPaymentMethod {
	var response []*response.MerchantResponseYearlyPaymentMethod
	for _, merchant := range ms {
		response = append(response, s.ToMerchantYearlyPaymentMethod(merchant))
	}
	return response
}

// ToMerchantMonthlyAmount maps a single monthly amount record to a
// MerchantResponseMonthlyAmount API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantMonthlyAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseMonthlyAmount containing the mapped data, including
//     Month and TotalAmount.
func (s *merchantResponseMapper) ToMerchantMonthlyAmount(ms *record.MerchantMonthlyAmount) *response.MerchantResponseMonthlyAmount {
	return &response.MerchantResponseMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: ms.TotalAmount,
	}
}

// ToMerchantMonthlyAmounts maps multiple monthly amount records into a slice of MerchantResponseMonthlyAmount.
//
// Args:
//   - ms: A slice of pointers to MerchantMonthlyAmount records containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseMonthlyAmount containing the mapped data, including
//     Month and TotalAmount.
func (s *merchantResponseMapper) ToMerchantMonthlyAmounts(ms []*record.MerchantMonthlyAmount) []*response.MerchantResponseMonthlyAmount {
	var response []*response.MerchantResponseMonthlyAmount
	for _, merchant := range ms {
		response = append(response, s.ToMerchantMonthlyAmount(merchant))
	}
	return response
}

// ToMerchantYearlyAmount maps a single yearly amount record to a
// MerchantResponseYearlyAmount API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantYearlyAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (s *merchantResponseMapper) ToMerchantYearlyAmount(ms *record.MerchantYearlyAmount) *response.MerchantResponseYearlyAmount {
	return &response.MerchantResponseYearlyAmount{
		Year:        ms.Year,
		TotalAmount: ms.TotalAmount,
	}
}

// ToMerchantYearlyAmounts maps multiple yearly amount records into a slice of
// MerchantResponseYearlyAmount.
//
// Args:
//   - ms: A slice of pointers to MerchantYearlyAmount records containing the data
//     to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseYearlyAmount containing the mapped
//     data, including Year and TotalAmount.
func (s *merchantResponseMapper) ToMerchantYearlyAmounts(ms []*record.MerchantYearlyAmount) []*response.MerchantResponseYearlyAmount {
	var response []*response.MerchantResponseYearlyAmount
	for _, merchant := range ms {
		response = append(response, s.ToMerchantYearlyAmount(merchant))
	}
	return response
}

// ToMerchantMonthlyTotalAmount maps a single MerchantMonthlyTotalAmount record to a
// MerchantResponseMonthlyTotalAmount API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantMonthlyTotalAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (s *merchantResponseMapper) ToMerchantMonthlyTotalAmount(ms *record.MerchantMonthlyTotalAmount) *response.MerchantResponseMonthlyTotalAmount {
	return &response.MerchantResponseMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: ms.TotalAmount,
	}
}

// ToMerchantMonthlyTotalAmounts maps multiple monthly total amount records into a slice of
// MerchantResponseMonthlyTotalAmount.
//
// Args:
//   - ms: A slice of pointers to MerchantMonthlyTotalAmount records containing the data
//     to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseMonthlyTotalAmount containing the mapped
//     data, including Month, Year, and TotalAmount.
func (s *merchantResponseMapper) ToMerchantMonthlyTotalAmounts(ms []*record.MerchantMonthlyTotalAmount) []*response.MerchantResponseMonthlyTotalAmount {
	var response []*response.MerchantResponseMonthlyTotalAmount
	for _, merchant := range ms {
		response = append(response, s.ToMerchantMonthlyTotalAmount(merchant))
	}
	return response
}

// ToMerchantYearlyTotalAmount maps a single yearly total amount record to a
// MerchantResponseYearlyTotalAmount API-compatible response.
//
// Args:
//   - ms: A pointer to a MerchantYearlyTotalAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantResponseYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (s *merchantResponseMapper) ToMerchantYearlyTotalAmount(ms *record.MerchantYearlyTotalAmount) *response.MerchantResponseYearlyTotalAmount {
	return &response.MerchantResponseYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: ms.TotalAmount,
	}
}

// ToMerchantYearlyTotalAmounts maps multiple yearly total amount records into a slice of
// MerchantResponseYearlyTotalAmount API-compatible responses.
//
// Args:
//   - ms: A slice of pointers to MerchantYearlyTotalAmount records containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantResponseYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.

func (s *merchantResponseMapper) ToMerchantYearlyTotalAmounts(ms []*record.MerchantYearlyTotalAmount) []*response.MerchantResponseYearlyTotalAmount {
	var response []*response.MerchantResponseYearlyTotalAmount
	for _, merchant := range ms {
		response = append(response, s.ToMerchantYearlyTotalAmount(merchant))
	}
	return response
}
