package apimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// merchantResponse provides methods to map gRPC merchant responses to HTTP API responses.
type merchantResponse struct{}

// NewMerchantResponseMapper returns a new instance of merchantResponse, a type that
// defines methods for mapping gRPC merchant-related responses to HTTP API responses.
func NewMerchantResponseMapper() *merchantResponse {
	return &merchantResponse{}
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
func (m *merchantResponse) ToApiResponseMerchant(merchants *pb.ApiResponseMerchant) *response.ApiResponseMerchant {
	return &response.ApiResponseMerchant{
		Status:  merchants.Status,
		Message: merchants.Message,
		Data:    m.mapMerchantResponse(merchants.Data),
	}
}

// ToApiResponseMerchants maps a gRPC merchant response slice to an HTTP API response. It constructs
// an ApiResponsesMerchant by copying the status and message fields and mapping the merchant
// data to a slice of MerchantResponse.
func (m *merchantResponse) ToApiResponseMerchants(merchants *pb.ApiResponsesMerchant) *response.ApiResponsesMerchant {
	return &response.ApiResponsesMerchant{
		Status:  merchants.Status,
		Message: merchants.Message,
		Data:    m.mapMerchantResponses(merchants.Data),
	}
}

// ToApiResponsesMerchant maps a paginated list of merchants to a paginated HTTP API response.
// It constructs an ApiResponsePaginationMerchant by copying the status and message fields
// and mapping the merchant data slice to a slice of MerchantResponse, and including pagination
// metadata.
//
// Args:
//
//	merchants: A pointer to a pb.ApiResponsePaginationMerchant containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationMerchant with mapped data and pagination info.
func (m *merchantResponse) ToApiResponsesMerchant(merchants *pb.ApiResponsePaginationMerchant) *response.ApiResponsePaginationMerchant {
	return &response.ApiResponsePaginationMerchant{
		Status:     merchants.Status,
		Message:    merchants.Message,
		Data:       m.mapMerchantResponses(merchants.Data),
		Pagination: mapPaginationMeta(merchants.Pagination),
	}
}

// ToApiResponsesMerchantDeleteAt maps a paginated list of soft-deleted merchants to an HTTP API response.
// It constructs an ApiResponsePaginationMerchantDeleteAt by copying the status and message fields,
// mapping the merchant data slice to a slice of MerchantResponseDeleteAt, and including pagination metadata.
//
// Args:
//
//	merchants: A pointer to a pb.ApiResponsePaginationMerchantDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationMerchantDeleteAt with mapped data and pagination info.
func (m *merchantResponse) ToApiResponsesMerchantDeleteAt(merchants *pb.ApiResponsePaginationMerchantDeleteAt) *response.ApiResponsePaginationMerchantDeleteAt {
	return &response.ApiResponsePaginationMerchantDeleteAt{
		Status:     merchants.Status,
		Message:    merchants.Message,
		Data:       m.mapMerchantResponsesDeleteAt(merchants.Data),
		Pagination: mapPaginationMeta(merchants.Pagination),
	}
}

// ToApiResponseMerchantsTransactionResponse maps a paginated gRPC response of merchant transactions to an HTTP API response.
// It constructs an ApiResponsePaginationMerchantTransaction by copying the status and message fields,
// mapping the transaction data slice to a slice of MerchantTransactionResponse, and including pagination metadata.
//
// Args:
//
//	merchants: A pointer to a pb.ApiResponsePaginationMerchantTransaction containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationMerchantTransaction with mapped data and pagination info.
func (m *merchantResponse) ToApiResponseMerchantsTransactionResponse(merchants *pb.ApiResponsePaginationMerchantTransaction) *response.ApiResponsePaginationMerchantTransaction {

	return &response.ApiResponsePaginationMerchantTransaction{
		Status:     merchants.Status,
		Message:    merchants.Message,
		Data:       m.mapMerchantTransactionResponses(merchants.Data),
		Pagination: mapPaginationMeta(merchants.Pagination),
	}
}

// ToApiResponseMonthlyPaymentMethods converts monthly payment method statistics from a gRPC
// response to an HTTP API response. It constructs an ApiResponseMerchantMonthlyPaymentMethod
// by copying the status and message fields and mapping the data to a slice of
// MerchantResponseMonthlyPaymentMethod.
//
// Args:
//
//	ms: A pointer to a pb.ApiResponseMerchantMonthlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantMonthlyPaymentMethod with mapped data.
func (m *merchantResponse) ToApiResponseMonthlyPaymentMethods(ms *pb.ApiResponseMerchantMonthlyPaymentMethod) *response.ApiResponseMerchantMonthlyPaymentMethod {
	return &response.ApiResponseMerchantMonthlyPaymentMethod{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesMonthlyPaymentMethod(ms.Data),
	}
}

// ToApiResponseYearlyPaymentMethods converts yearly payment method statistics from
// a gRPC response to an HTTP API response. It constructs an ApiResponseMerchantYearlyPaymentMethod
// by copying the status and message fields and mapping the data to a slice of
// MerchantResponseYearlyPaymentMethod.
//
// Args:
//
//	ms: A pointer to a pb.ApiResponseMerchantYearlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantYearlyPaymentMethod with mapped data.
func (m *merchantResponse) ToApiResponseYearlyPaymentMethods(ms *pb.ApiResponseMerchantYearlyPaymentMethod) *response.ApiResponseMerchantYearlyPaymentMethod {
	return &response.ApiResponseMerchantYearlyPaymentMethod{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesYearlyPaymentMethod(ms.Data),
	}
}

// ToApiResponseMonthlyAmounts converts monthly financial amounts data from a gRPC response
// into an HTTP API response format. It constructs an ApiResponseMerchantMonthlyAmount by
// copying the status and message fields and mapping the data to a slice of
// MerchantResponseMonthlyAmount.
//
// Args:
//
//	ms: A pointer to a pb.ApiResponseMerchantMonthlyAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantMonthlyAmount with mapped data.
func (m *merchantResponse) ToApiResponseMonthlyAmounts(ms *pb.ApiResponseMerchantMonthlyAmount) *response.ApiResponseMerchantMonthlyAmount {
	return &response.ApiResponseMerchantMonthlyAmount{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesMonthlyAmount(ms.Data),
	}
}

// ToApiResponseYearlyAmounts converts yearly financial amounts data from a gRPC response
// into an HTTP API response format. It constructs an ApiResponseMerchantYearlyAmount by
// copying the status and message fields and mapping the data to a slice of
// MerchantResponseYearlyAmount.
//
// Args:
//
//	ms: A pointer to a pb.ApiResponseMerchantYearlyAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantYearlyAmount with mapped data.
func (m *merchantResponse) ToApiResponseYearlyAmounts(ms *pb.ApiResponseMerchantYearlyAmount) *response.ApiResponseMerchantYearlyAmount {
	return &response.ApiResponseMerchantYearlyAmount{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesYearlyAmount(ms.Data),
	}
}

// ToApiResponseMonthlyTotalAmounts converts monthly total amount data from a gRPC response
// into an HTTP API response format. It constructs an ApiResponseMerchantMonthlyTotalAmount by
// copying the status and message fields and mapping the data to a slice of
// MerchantResponseMonthlyTotalAmount.
//
// Args:
//
//	ms: A pointer to a pb.ApiResponseMerchantMonthlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantMonthlyTotalAmount with mapped data.
func (m *merchantResponse) ToApiResponseMonthlyTotalAmounts(ms *pb.ApiResponseMerchantMonthlyTotalAmount) *response.ApiResponseMerchantMonthlyTotalAmount {
	return &response.ApiResponseMerchantMonthlyTotalAmount{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesMonthlyTotalAmount(ms.Data),
	}
}

// ToApiResponseYearlyTotalAmounts converts yearly total amount data from a gRPC response
// into an HTTP API response format. It constructs an ApiResponseMerchantYearlyTotalAmount by
// copying the status and message fields and mapping the data to a slice of
// MerchantResponseYearlyTotalAmount.
//
// Args:
//
//	ms: A pointer to a pb.ApiResponseMerchantYearlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantYearlyTotalAmount with mapped data.
func (m *merchantResponse) ToApiResponseYearlyTotalAmounts(ms *pb.ApiResponseMerchantYearlyTotalAmount) *response.ApiResponseMerchantYearlyTotalAmount {
	return &response.ApiResponseMerchantYearlyTotalAmount{
		Status:  ms.Status,
		Message: ms.Message,
		Data:    m.mapResponsesYearlyTotalAmount(ms.Data),
	}
}

// ToApiResponseMerchantDeleteAt maps a gRPC response for a deleted merchant to an HTTP API response. It
// constructs an ApiResponseMerchantDelete by copying the status and message fields.
//
// Args:
//
//	card: A pointer to a pb.ApiResponseMerchantDelete containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantDelete with mapped data.
func (s *merchantResponse) ToApiResponseMerchantDeleteAt(card *pb.ApiResponseMerchantDelete) *response.ApiResponseMerchantDelete {
	return &response.ApiResponseMerchantDelete{
		Status:  card.Status,
		Message: card.Message,
	}
}

// ToApiResponseMerchantAll maps a gRPC response containing all merchants to an HTTP API response.
// It constructs an ApiResponseMerchantAll by copying the status and message fields from the
// gRPC response.
//
// Args:
//
//	card: A pointer to a pb.ApiResponseMerchantAll containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantAll with the status and message set.
func (s *merchantResponse) ToApiResponseMerchantAll(card *pb.ApiResponseMerchantAll) *response.ApiResponseMerchantAll {
	return &response.ApiResponseMerchantAll{
		Status:  card.Status,
		Message: card.Message,
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
func (m *merchantResponse) mapMerchantResponse(merchant *pb.MerchantResponse) *response.MerchantResponse {
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

// mapMerchantResponses maps a slice of gRPC MerchantResponse to a slice of HTTP API responses. It
// constructs a slice of ApiResponseMerchant by copying the status, message, ID, name,
// status, API key, user ID, created at, and updated at fields from the gRPC
// response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponse with the mapped data.
func (m *merchantResponse) mapMerchantResponses(r []*pb.MerchantResponse) []*response.MerchantResponse {
	var responseMerchants []*response.MerchantResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapMerchantResponse(merchant))
	}

	return responseMerchants
}

// mapMerchantResponseDeleteAt maps a gRPC MerchantResponseDeleteAt to an HTTP API response. It
// constructs an ApiResponseMerchantDeleteAt by copying the status, message, ID, name,
// status, API key, user ID, created at, updated at, and deleted at fields from the gRPC
// response.
//
// Args:
//
//	merchant: A pointer to a pb.MerchantResponseDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseDeleteAt with the mapped data.
func (m *merchantResponse) mapMerchantResponseDeleteAt(merchant *pb.MerchantResponseDeleteAt) *response.MerchantResponseDeleteAt {
	var deletedAt string
	if merchant.DeletedAt != nil {
		deletedAt = merchant.DeletedAt.Value
	}

	return &response.MerchantResponseDeleteAt{
		ID:        int(merchant.Id),
		Name:      merchant.Name,
		Status:    merchant.Status,
		UserID:    int(merchant.UserId),
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
		DeletedAt: &deletedAt,
	}
}

// mapMerchantResponsesDeleteAt maps a slice of gRPC MerchantResponseDeleteAt to a slice of HTTP API responses.
// It constructs a slice of ApiResponseMerchantDeleteAt by iterating over each MerchantResponseDeleteAt
// in the input slice and mapping them using mapMerchantResponseDeleteAt. It handles the optional DeletedAt field.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseDeleteAt with the mapped data.
func (m *merchantResponse) mapMerchantResponsesDeleteAt(r []*pb.MerchantResponseDeleteAt) []*response.MerchantResponseDeleteAt {
	var responseMerchants []*response.MerchantResponseDeleteAt
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapMerchantResponseDeleteAt(merchant))
	}

	return responseMerchants
}

// mapMerchantTransactionResponse maps a gRPC MerchantTransactionResponse to an HTTP API response.
// It constructs a MerchantTransactionResponse by copying fields such as ID, CardNumber,
// Amount, PaymentMethod, MerchantID, MerchantName, TransactionTime, CreatedAt, and UpdatedAt
// from the gRPC response.
//
// Args:
//
//	merchant: A pointer to a pb.MerchantTransactionResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantTransactionResponse with the mapped data.
func (m *merchantResponse) mapMerchantTransactionResponse(merchant *pb.MerchantTransactionResponse) *response.MerchantTransactionResponse {

	return &response.MerchantTransactionResponse{
		ID:              int(merchant.Id),
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantId,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt,
		UpdatedAt:       merchant.UpdatedAt,
	}
}

// mapMerchantTransactionResponses maps a slice of gRPC MerchantTransactionResponses to a slice of HTTP API responses.
// It constructs a slice of MerchantTransactionResponse by copying fields such as ID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
// TransactionTime, CreatedAt, and UpdatedAt from the gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantTransactionResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantTransactionResponse with the mapped data.
func (m *merchantResponse) mapMerchantTransactionResponses(r []*pb.MerchantTransactionResponse) []*response.MerchantTransactionResponse {
	var responseMerchants []*response.MerchantTransactionResponse
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapMerchantTransactionResponse(merchant))
	}

	return responseMerchants
}

// mapResponseMonthlyPaymentMethod maps a single gRPC MerchantResponseMonthlyPaymentMethod to an HTTP API response.
// It constructs a MerchantResponseMonthlyPaymentMethod by copying the month, payment method, and total amount
// from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pb.MerchantResponseMonthlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseMonthlyPaymentMethod with the mapped data.
func (m *merchantResponse) mapResponseMonthlyPaymentMethod(ms *pb.MerchantResponseMonthlyPaymentMethod) *response.MerchantResponseMonthlyPaymentMethod {
	return &response.MerchantResponseMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// mapResponsesMonthlyPaymentMethod maps a slice of gRPC MerchantResponseMonthlyPaymentMethods to a slice of HTTP API responses.
// It constructs a slice of MerchantResponseMonthlyPaymentMethod by copying the month, payment method, and total amount
// from the gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseMonthlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseMonthlyPaymentMethod with the mapped data.
func (m *merchantResponse) mapResponsesMonthlyPaymentMethod(r []*pb.MerchantResponseMonthlyPaymentMethod) []*response.MerchantResponseMonthlyPaymentMethod {
	var responseMerchants []*response.MerchantResponseMonthlyPaymentMethod
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseMonthlyPaymentMethod(merchant))
	}

	return responseMerchants
}

// mapResponseYearlyPaymentMethod maps a single gRPC MerchantResponseYearlyPaymentMethod to an HTTP API response.
// It constructs a MerchantResponseYearlyPaymentMethod by copying the year, payment method, and total amount
// from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pb.MerchantResponseYearlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseYearlyPaymentMethod with the mapped data.
func (m *merchantResponse) mapResponseYearlyPaymentMethod(ms *pb.MerchantResponseYearlyPaymentMethod) *response.MerchantResponseYearlyPaymentMethod {
	return &response.MerchantResponseYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// mapResponsesYearlyPaymentMethod maps a slice of gRPC MerchantResponseYearlyPaymentMethod
// to a slice of HTTP API responses. It constructs a slice of MerchantResponseYearlyPaymentMethod
// by copying the year, payment method, and total amount from each gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseYearlyPaymentMethod containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseYearlyPaymentMethod with the mapped data.
func (m *merchantResponse) mapResponsesYearlyPaymentMethod(r []*pb.MerchantResponseYearlyPaymentMethod) []*response.MerchantResponseYearlyPaymentMethod {
	var responseMerchants []*response.MerchantResponseYearlyPaymentMethod
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseYearlyPaymentMethod(merchant))
	}

	return responseMerchants
}

// mapResponseMonthlyAmount maps a single gRPC MerchantResponseMonthlyAmount to an HTTP API response.
// It constructs a MerchantResponseMonthlyAmount by copying the month and total amount from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pb.MerchantResponseMonthlyAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseMonthlyAmount with the mapped data.
func (m *merchantResponse) mapResponseMonthlyAmount(ms *pb.MerchantResponseMonthlyAmount) *response.MerchantResponseMonthlyAmount {
	return &response.MerchantResponseMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int(ms.TotalAmount),
	}
}

// mapResponsesMonthlyAmount maps a slice of gRPC MerchantResponseMonthlyAmount
// to a slice of HTTP API responses. It constructs a slice of MerchantResponseMonthlyAmount
// by copying the month and total amount from each gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseMonthlyAmount containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseMonthlyAmount with the mapped data.
func (m *merchantResponse) mapResponsesMonthlyAmount(r []*pb.MerchantResponseMonthlyAmount) []*response.MerchantResponseMonthlyAmount {
	var responseMerchants []*response.MerchantResponseMonthlyAmount
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseMonthlyAmount(merchant))
	}

	return responseMerchants
}

// mapResponseYearlyAmount maps a single gRPC MerchantResponseYearlyAmount to an HTTP API response.
// It constructs a MerchantResponseYearlyAmount by copying the year and total amount from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pb.MerchantResponseYearlyAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseYearlyAmount with the mapped data.
func (m *merchantResponse) mapResponseYearlyAmount(ms *pb.MerchantResponseYearlyAmount) *response.MerchantResponseYearlyAmount {
	return &response.MerchantResponseYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// mapResponsesYearlyAmount maps a slice of gRPC MerchantResponseYearlyAmount
// to a slice of HTTP API responses. It constructs a slice of MerchantResponseYearlyAmount
// by copying the year and total amount from each gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseYearlyAmount containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseYearlyAmount with the mapped data.
func (m *merchantResponse) mapResponsesYearlyAmount(r []*pb.MerchantResponseYearlyAmount) []*response.MerchantResponseYearlyAmount {
	var responseMerchants []*response.MerchantResponseYearlyAmount
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseYearlyAmount(merchant))
	}

	return responseMerchants
}

// mapResponseMonthlyTotalAmount maps a single gRPC MerchantResponseMonthlyTotalAmount to an HTTP API response.
// It constructs a MerchantResponseMonthlyTotalAmount by copying the month, year, and total amount from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pb.MerchantResponseMonthlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseMonthlyTotalAmount with the mapped data.
func (m *merchantResponse) mapResponseMonthlyTotalAmount(ms *pb.MerchantResponseMonthlyTotalAmount) *response.MerchantResponseMonthlyTotalAmount {
	return &response.MerchantResponseMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// mapResponsesMonthlyTotalAmount maps a slice of gRPC MerchantResponseMonthlyTotalAmount
// to a slice of HTTP API responses. It constructs a slice of MerchantResponseMonthlyTotalAmount
// by iterating over each gRPC response and mapping them using mapResponseMonthlyTotalAmount.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseMonthlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseMonthlyTotalAmount with the mapped data.
func (m *merchantResponse) mapResponsesMonthlyTotalAmount(r []*pb.MerchantResponseMonthlyTotalAmount) []*response.MerchantResponseMonthlyTotalAmount {
	var responseMerchants []*response.MerchantResponseMonthlyTotalAmount
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseMonthlyTotalAmount(merchant))
	}

	return responseMerchants
}

// mapResponseYearlyTotalAmount maps a single gRPC MerchantResponseYearlyTotalAmount to an HTTP API response.
// It constructs a MerchantResponseYearlyTotalAmount by copying the year and total amount from the gRPC response.
//
// Args:
//
//	ms: A pointer to a pb.MerchantResponseYearlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantResponseYearlyTotalAmount with the mapped data.
func (m *merchantResponse) mapResponseYearlyTotalAmount(ms *pb.MerchantResponseYearlyTotalAmount) *response.MerchantResponseYearlyTotalAmount {
	return &response.MerchantResponseYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// mapResponsesYearlyTotalAmount maps a slice of gRPC MerchantResponseYearlyTotalAmount
// to a slice of HTTP API responses. It constructs a slice of MerchantResponseYearlyTotalAmount
// by copying the year and total amount from each gRPC response.
//
// Args:
//
//	r: A slice of pointers to pb.MerchantResponseYearlyTotalAmount containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.MerchantResponseYearlyTotalAmount with the mapped data.
func (m *merchantResponse) mapResponsesYearlyTotalAmount(r []*pb.MerchantResponseYearlyTotalAmount) []*response.MerchantResponseYearlyTotalAmount {
	var responseMerchants []*response.MerchantResponseYearlyTotalAmount
	for _, merchant := range r {
		responseMerchants = append(responseMerchants, m.mapResponseYearlyTotalAmount(merchant))
	}

	return responseMerchants
}
