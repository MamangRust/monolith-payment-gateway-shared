package apimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// transactionResponseMapper provides methods to map gRPC transaction responses to HTTP-compatible API responses.
type transactionResponseMapper struct {
}

// NewTransactionResponseMapper creates a new transactionResponseMapper.
func NewTransactionResponseMapper() *transactionResponseMapper {
	return &transactionResponseMapper{}
}

// ToApiResponseTransactionMonthStatusSuccess maps a single gRPC transaction month status response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionMonthStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionMonthStatusSuccess containing the mapped data, including status, message,
//     and a single mapped transaction month status response.
func (m *transactionResponseMapper) ToApiResponseTransactionMonthStatusSuccess(pbResponse *pb.ApiResponseTransactionMonthStatusSuccess) *response.ApiResponseTransactionMonthStatusSuccess {

	return &response.ApiResponseTransactionMonthStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.ToResponsesTransactionMonthStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseTransactionYearStatusSuccess maps a single gRPC transaction year status response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionYearStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionYearStatusSuccess containing the mapped data, including status, message,
//     and a single mapped transaction year status response.
func (m *transactionResponseMapper) ToApiResponseTransactionYearStatusSuccess(pbResponse *pb.ApiResponseTransactionYearStatusSuccess) *response.ApiResponseTransactionYearStatusSuccess {

	return &response.ApiResponseTransactionYearStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapTransactionResponsesYearStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseTransactionMonthStatusFailed maps a single gRPC transaction month status failed response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionMonthStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionMonthStatusFailed containing the mapped data, including status, message,
//     and a single mapped transaction month status failed response.
func (m *transactionResponseMapper) ToApiResponseTransactionMonthStatusFailed(pbResponse *pb.ApiResponseTransactionMonthStatusFailed) *response.ApiResponseTransactionMonthStatusFailed {

	return &response.ApiResponseTransactionMonthStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesTransactionMonthStatusFailed(pbResponse.Data),
	}
}

// ToApiResponseTransactionYearStatusFailed maps a single gRPC transaction year status failed response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionYearStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionYearStatusFailed containing the mapped data, including status, message,
//     and a single mapped transaction year status failed response.
func (m *transactionResponseMapper) ToApiResponseTransactionYearStatusFailed(pbResponse *pb.ApiResponseTransactionYearStatusFailed) *response.ApiResponseTransactionYearStatusFailed {

	return &response.ApiResponseTransactionYearStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapTransactionResponsesYearStatusFailed(pbResponse.Data),
	}
}

// ToApiResponseTransactionMonthMethod maps a single gRPC transaction month method response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionMonthMethod containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionMonthMethod containing the mapped data, including status, message,
//     and a single mapped transaction month method response.
func (m *transactionResponseMapper) ToApiResponseTransactionMonthMethod(pbResponse *pb.ApiResponseTransactionMonthMethod) *response.ApiResponseTransactionMonthMethod {

	return &response.ApiResponseTransactionMonthMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransactionMonthMethods(pbResponse.Data),
	}
}

// ToApiResponseTransactionYearMethod maps a single gRPC transaction year method response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionYearMethod containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionYearMethod containing the mapped data, including status, message,
//     and a single mapped transaction year method response.
func (m *transactionResponseMapper) ToApiResponseTransactionYearMethod(pbResponse *pb.ApiResponseTransactionYearMethod) *response.ApiResponseTransactionYearMethod {

	return &response.ApiResponseTransactionYearMethod{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransactionYearMethods(pbResponse.Data),
	}
}

// ToApiResponseTransactionMonthAmount maps a gRPC transaction month amount response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionMonthAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionMonthAmount containing the mapped data, including status, message,
//     and a single mapped transaction month amount response.
func (m *transactionResponseMapper) ToApiResponseTransactionMonthAmount(pbResponse *pb.ApiResponseTransactionMonthAmount) *response.ApiResponseTransactionMonthAmount {

	return &response.ApiResponseTransactionMonthAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransactionMonthAmounts(pbResponse.Data),
	}
}

// ToApiResponseTransactionYearAmount maps a single gRPC transaction year amount response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionYearAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionYearAmount containing the mapped data, including status, message,
//     and a single mapped transaction year amount response.
func (m *transactionResponseMapper) ToApiResponseTransactionYearAmount(pbResponse *pb.ApiResponseTransactionYearAmount) *response.ApiResponseTransactionYearAmount {

	return &response.ApiResponseTransactionYearAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransactionYearlyAmounts(pbResponse.Data),
	}
}

// ToApiResponseTransaction maps a single gRPC transaction response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransaction containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransaction containing the mapped data, including status, message,
//     and a single mapped transaction response.
func (m *transactionResponseMapper) ToApiResponseTransaction(pbResponse *pb.ApiResponseTransaction) *response.ApiResponseTransaction {
	return &response.ApiResponseTransaction{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransaction(pbResponse.Data),
	}
}

// ToApiResponseTransactions maps multiple gRPC transaction responses into an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactions containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactions containing the mapped data, including status, message,
//     and a slice of mapped transaction responses.
func (m *transactionResponseMapper) ToApiResponseTransactions(pbResponse *pb.ApiResponseTransactions) *response.ApiResponseTransactions {

	return &response.ApiResponseTransactions{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesTransaction(pbResponse.Data),
	}
}

// ToApiResponseTransactionDelete maps a single gRPC transaction delete response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionDelete containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionDelete containing the mapped data, including status and message.
func (m *transactionResponseMapper) ToApiResponseTransactionDelete(pbResponse *pb.ApiResponseTransactionDelete) *response.ApiResponseTransactionDelete {
	return &response.ApiResponseTransactionDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseTransactionAll maps a single gRPC transaction all response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransactionAll containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransactionAll containing the mapped data, including status and message.
func (m *transactionResponseMapper) ToApiResponseTransactionAll(pbResponse *pb.ApiResponseTransactionAll) *response.ApiResponseTransactionAll {
	return &response.ApiResponseTransactionAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponsePaginationTransaction maps a paginated gRPC response of transactions to an HTTP API response.
// It constructs an ApiResponsePaginationTransaction by copying the status and message fields,
// mapping the transaction data slice to a slice of TransactionResponse, and including pagination metadata.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponsePaginationTransaction containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationTransaction with mapped data and pagination info.
func (m *transactionResponseMapper) ToApiResponsePaginationTransaction(pbResponse *pb.ApiResponsePaginationTransaction) *response.ApiResponsePaginationTransaction {

	return &response.ApiResponsePaginationTransaction{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.mapResponsesTransaction(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

// ToApiResponsePaginationTransactionDeleteAt maps a paginated gRPC response of
// soft-deleted transactions to an HTTP API response. It constructs an
// ApiResponsePaginationTransactionDeleteAt by copying the status and message fields,
// mapping the transaction data slice to a slice of TransactionResponseDeleteAt, and
// including pagination metadata.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationTransactionDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationTransactionDeleteAt with mapped data and pagination info.
func (m *transactionResponseMapper) ToApiResponsePaginationTransactionDeleteAt(pbResponse *pb.ApiResponsePaginationTransactionDeleteAt) *response.ApiResponsePaginationTransactionDeleteAt {

	return &response.ApiResponsePaginationTransactionDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.ToResponsesTransactionDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

// mapResponseTransaction maps a single gRPC transaction response to an HTTP API response.
//
// Args:
//   - transaction: A pointer to a pb.TransactionResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponse containing the mapped data.
func (m *transactionResponseMapper) mapResponseTransaction(transaction *pb.TransactionResponse) *response.TransactionResponse {
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

// mapResponsesTransaction maps a slice of gRPC transaction responses to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponse containing the mapped data.
func (m *transactionResponseMapper) mapResponsesTransaction(transactions []*pb.TransactionResponse) []*response.TransactionResponse {
	var result []*response.TransactionResponse
	for _, transaction := range transactions {
		result = append(result, m.mapResponseTransaction(transaction))
	}
	return result
}

// mapResponseTransactionDeleteAt maps a single gRPC transaction delete response to an HTTP API response.
//
// Args:
//   - transaction: A pointer to a pb.TransactionResponseDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseDeleteAt containing the mapped data.
func (m *transactionResponseMapper) mapResponseTransactionDeleteAt(transaction *pb.TransactionResponseDeleteAt) *response.TransactionResponseDeleteAt {
	var deletedAt string
	if transaction.DeletedAt != nil {
		deletedAt = transaction.DeletedAt.Value
	}

	return &response.TransactionResponseDeleteAt{
		ID:              int(transaction.Id),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantID:      int(transaction.MerchantId),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
		DeletedAt:       &deletedAt,
	}
}

// ToResponsesTransactionDeleteAt maps a slice of gRPC transaction delete responses to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionResponseDeleteAt containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponseDeleteAt containing the mapped data.
func (m *transactionResponseMapper) ToResponsesTransactionDeleteAt(transactions []*pb.TransactionResponseDeleteAt) []*response.TransactionResponseDeleteAt {
	var result []*response.TransactionResponseDeleteAt
	for _, transaction := range transactions {
		result = append(result, m.mapResponseTransactionDeleteAt(transaction))
	}
	return result
}

// mapResponseTransactionMonthStatusSuccess maps a gRPC transaction month status success response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionMonthStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseMonthStatusSuccess containing the mapped data,
//     including the year, month, total successful transactions, and total amount.
func (m *transactionResponseMapper) mapResponseTransactionMonthStatusSuccess(s *pb.TransactionMonthStatusSuccessResponse) *response.TransactionResponseMonthStatusSuccess {
	return &response.TransactionResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToResponsesTransactionMonthStatusSuccess maps a slice of gRPC transaction month status success responses
// to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionMonthStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponseMonthStatusSuccess containing the mapped data.
func (m *transactionResponseMapper) ToResponsesTransactionMonthStatusSuccess(transactions []*pb.TransactionMonthStatusSuccessResponse) []*response.TransactionResponseMonthStatusSuccess {
	var transactionRecords []*response.TransactionResponseMonthStatusSuccess
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, m.mapResponseTransactionMonthStatusSuccess(transaction))
	}
	return transactionRecords
}

// mapTransactionResponseYearStatusSuccess maps a gRPC transaction year status success response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionYearStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseYearStatusSuccess containing the mapped data,
//     including the year, total successful transactions, and total amount.
func (m *transactionResponseMapper) mapTransactionResponseYearStatusSuccess(s *pb.TransactionYearStatusSuccessResponse) *response.TransactionResponseYearStatusSuccess {
	return &response.TransactionResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapTransactionResponsesYearStatusSuccess maps a slice of gRPC transaction year status success responses
// to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionYearStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponseYearStatusSuccess containing the mapped data.
func (m *transactionResponseMapper) mapTransactionResponsesYearStatusSuccess(transactions []*pb.TransactionYearStatusSuccessResponse) []*response.TransactionResponseYearStatusSuccess {
	var transactionRecords []*response.TransactionResponseYearStatusSuccess
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, m.mapTransactionResponseYearStatusSuccess(transaction))
	}
	return transactionRecords
}

// mapResponseTransactionMonthStatusFailed maps a gRPC transaction month status failed response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionMonthStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseMonthStatusFailed containing the mapped data,
//     including the year, month, total failed transactions, and total amount.
func (m *transactionResponseMapper) mapResponseTransactionMonthStatusFailed(s *pb.TransactionMonthStatusFailedResponse) *response.TransactionResponseMonthStatusFailed {
	return &response.TransactionResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponsesTransactionMonthStatusFailed maps a slice of gRPC transaction month status failed responses
// to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionMonthStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponseMonthStatusFailed containing the mapped data.
func (m *transactionResponseMapper) mapResponsesTransactionMonthStatusFailed(transactions []*pb.TransactionMonthStatusFailedResponse) []*response.TransactionResponseMonthStatusFailed {
	var transactionRecords []*response.TransactionResponseMonthStatusFailed
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, m.mapResponseTransactionMonthStatusFailed(transaction))
	}
	return transactionRecords
}

// mapTransactionResponseYearStatusFailed maps a gRPC transaction year status failed response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionYearStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionResponseYearStatusFailed containing the mapped data,
//     including the year, total failed transactions, and total amount.
func (m *transactionResponseMapper) mapTransactionResponseYearStatusFailed(s *pb.TransactionYearStatusFailedResponse) *response.TransactionResponseYearStatusFailed {
	return &response.TransactionResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapTransactionResponsesYearStatusFailed maps a slice of gRPC transaction year status failed responses
// to a slice of HTTP API responses.
//
// Args:
//   - transactions: A slice of pointers to pb.TransactionYearStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionResponseYearStatusFailed containing the mapped data.
func (m *transactionResponseMapper) mapTransactionResponsesYearStatusFailed(transactions []*pb.TransactionYearStatusFailedResponse) []*response.TransactionResponseYearStatusFailed {
	var transactionRecords []*response.TransactionResponseYearStatusFailed
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, m.mapTransactionResponseYearStatusFailed(transaction))
	}
	return transactionRecords
}

// mapResponseTransactionMonthMethod maps a gRPC transaction month method response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionMonthMethodResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionMonthMethodResponse containing the mapped data,
//     including the month, payment method, total transactions, and total amount.
func (m *transactionResponseMapper) mapResponseTransactionMonthMethod(s *pb.TransactionMonthMethodResponse) *response.TransactionMonthMethodResponse {
	return &response.TransactionMonthMethodResponse{
		Month:             s.Month,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int(s.TotalTransactions),
		TotalAmount:       int(s.TotalAmount),
	}
}

// mapResponseTransactionMonthMethods maps a slice of gRPC transaction month method responses
// to a slice of HTTP API responses.
//
// Args:
//   - s: A slice of pointers to pb.TransactionMonthMethodResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionMonthMethodResponse containing the mapped data,
//     including the month, payment method, total transactions, and total amount.
func (m *transactionResponseMapper) mapResponseTransactionMonthMethods(s []*pb.TransactionMonthMethodResponse) []*response.TransactionMonthMethodResponse {
	var responses []*response.TransactionMonthMethodResponse
	for _, transaction := range s {
		responses = append(responses, m.mapResponseTransactionMonthMethod(transaction))
	}
	return responses
}

// mapResponseTransactionYearMethod maps a single gRPC transaction year method response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionYearMethodResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionYearMethodResponse containing the mapped data,
//     including the year, payment method, total transactions, and total amount.
func (m *transactionResponseMapper) mapResponseTransactionYearMethod(s *pb.TransactionYearMethodResponse) *response.TransactionYearMethodResponse {
	return &response.TransactionYearMethodResponse{
		Year:              s.Year,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int(s.TotalTransactions),
		TotalAmount:       int(s.TotalAmount),
	}
}

// mapResponseTransactionYearMethods maps a slice of gRPC transaction year method responses
// to a slice of HTTP API responses.
//
// Args:
//   - s: A slice of pointers to pb.TransactionYearMethodResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionYearMethodResponse containing the mapped data,
//     including the year, payment method, total transactions, and total amount.
func (m *transactionResponseMapper) mapResponseTransactionYearMethods(s []*pb.TransactionYearMethodResponse) []*response.TransactionYearMethodResponse {
	var responses []*response.TransactionYearMethodResponse
	for _, transaction := range s {
		responses = append(responses, m.mapResponseTransactionYearMethod(transaction))
	}
	return responses
}

// mapResponseTransactionMonthAmount maps a gRPC transaction month amount response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionMonthAmountResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionMonthAmountResponse containing the mapped data,
//     including the month and total amount.
func (m *transactionResponseMapper) mapResponseTransactionMonthAmount(s *pb.TransactionMonthAmountResponse) *response.TransactionMonthAmountResponse {
	return &response.TransactionMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTransactionMonthAmounts maps a slice of gRPC transaction month amount responses
// to a slice of HTTP API responses.
//
// Args:
//   - s: A slice of pointers to pb.TransactionMonthAmountResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionMonthAmountResponse containing the mapped data,
//     including the month and total amount.
func (m *transactionResponseMapper) mapResponseTransactionMonthAmounts(s []*pb.TransactionMonthAmountResponse) []*response.TransactionMonthAmountResponse {
	var responses []*response.TransactionMonthAmountResponse
	for _, transaction := range s {
		responses = append(responses, m.mapResponseTransactionMonthAmount(transaction))
	}
	return responses
}

// mapResponseTransactionYearlyAmount maps a gRPC transaction yearly amount response
// to an HTTP API response format.
//
// Args:
//   - s: A pointer to a pb.TransactionYearlyAmountResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TransactionYearlyAmountResponse containing the mapped data,
//     including the year and total amount.
func (m *transactionResponseMapper) mapResponseTransactionYearlyAmount(s *pb.TransactionYearlyAmountResponse) *response.TransactionYearlyAmountResponse {
	return &response.TransactionYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTransactionYearlyAmounts maps a slice of gRPC transaction yearly amount responses
// to a slice of HTTP API responses.
//
// Args:
//   - s: A slice of pointers to pb.TransactionYearlyAmountResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TransactionYearlyAmountResponse containing the mapped data,
//     including the year and total amount.
func (m *transactionResponseMapper) mapResponseTransactionYearlyAmounts(s []*pb.TransactionYearlyAmountResponse) []*response.TransactionYearlyAmountResponse {
	var responses []*response.TransactionYearlyAmountResponse
	for _, transaction := range s {
		responses = append(responses, m.mapResponseTransactionYearlyAmount(transaction))
	}
	return responses
}
