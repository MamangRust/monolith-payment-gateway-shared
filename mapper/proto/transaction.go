package protomapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// transactionProtoMapper is responsible for mapping Transaction responses to their corresponding protobuf representations.
type transactionProtoMapper struct{}

// NewTransactionProtoMapper returns a new instance of transactionProtoMapper.
func NewTransactionProtoMapper() *transactionProtoMapper {
	return &transactionProtoMapper{}
}

// ToProtoResponseTransactionMonthStatusSuccess converts a status, message, and a list of TransactionResponseMonthStatusSuccess
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionMonthStatusSuccess containing the status, message, and transaction data.
func (m *transactionProtoMapper) ToProtoResponseTransactionMonthStatusSuccess(status string, message string, pbResponse []*response.TransactionResponseMonthStatusSuccess) *pb.ApiResponseTransactionMonthStatusSuccess {
	return &pb.ApiResponseTransactionMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransactionMonthStatusSuccess(pbResponse),
	}
}

// ToProtoResponseTransactionYearStatusSuccess converts a status, message, and a list of TransactionResponseYearStatusSuccess
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionYearStatusSuccess containing the status, message, and transaction data.
func (m *transactionProtoMapper) ToProtoResponseTransactionYearStatusSuccess(status string, message string, pbResponse []*response.TransactionResponseYearStatusSuccess) *pb.ApiResponseTransactionYearStatusSuccess {
	return &pb.ApiResponseTransactionYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapTransactionResponsesYearStatusSuccess(pbResponse),
	}
}

// ToProtoResponseTransactionMonthStatusFailed converts a status, message, and a list of TransactionResponseMonthStatusFailed
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionMonthStatusFailed containing the status, message, and transaction data.
func (m *transactionProtoMapper) ToProtoResponseTransactionMonthStatusFailed(status string, message string, pbResponse []*response.TransactionResponseMonthStatusFailed) *pb.ApiResponseTransactionMonthStatusFailed {
	return &pb.ApiResponseTransactionMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransactionMonthStatusFailed(pbResponse),
	}
}

// ToProtoResponseTransactionYearStatusFailed converts a status, message, and a list of TransactionResponseYearStatusFailed
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionResponseYearStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionYearStatusFailed containing the status, message, and transaction data.
func (m *transactionProtoMapper) ToProtoResponseTransactionYearStatusFailed(status string, message string, pbResponse []*response.TransactionResponseYearStatusFailed) *pb.ApiResponseTransactionYearStatusFailed {
	return &pb.ApiResponseTransactionYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapTransactionResponsesYearStatusFailed(pbResponse),
	}
}

// ToProtoResponseTransactionMonthMethod converts a status, message, and a list of TransactionMonthMethodResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction method data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionMonthMethodResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionMonthMethod containing the status, message, and transaction method data.
func (m *transactionProtoMapper) ToProtoResponseTransactionMonthMethod(status string, message string, pbResponse []*response.TransactionMonthMethodResponse) *pb.ApiResponseTransactionMonthMethod {
	return &pb.ApiResponseTransactionMonthMethod{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransactionMonthMethods(pbResponse),
	}
}

// ToProtoResponseTransactionYearMethod converts a status, message, and a list of TransactionYearMethodResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction method data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionYearMethodResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionYearMethod containing the status, message, and transaction method data.
func (m *transactionProtoMapper) ToProtoResponseTransactionYearMethod(status string, message string, pbResponse []*response.TransactionYearMethodResponse) *pb.ApiResponseTransactionYearMethod {

	return &pb.ApiResponseTransactionYearMethod{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransactionYearMethods(pbResponse),
	}
}

// ToProtoResponseTransactionMonthAmount converts a status, message, and a list of TransactionMonthAmountResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction month amount data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionMonthAmountResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionMonthAmount containing the status, message, and transaction month amount data.
func (m *transactionProtoMapper) ToProtoResponseTransactionMonthAmount(status string, message string, pbResponse []*response.TransactionMonthAmountResponse) *pb.ApiResponseTransactionMonthAmount {
	return &pb.ApiResponseTransactionMonthAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransactionMonthAmounts(pbResponse),
	}
}

// ToProtoResponseTransactionYearAmount converts a status, message, and a list of TransactionYearlyAmountResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction year amount data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionYearlyAmountResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactionYearAmount containing the status, message, and transaction year amount data.
func (m *transactionProtoMapper) ToProtoResponseTransactionYearAmount(status string, message string, pbResponse []*response.TransactionYearlyAmountResponse) *pb.ApiResponseTransactionYearAmount {
	return &pb.ApiResponseTransactionYearAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransactionYearlyAmounts(pbResponse),
	}
}

// ToProtoResponseTransaction converts a status, message, and a TransactionResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The TransactionResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransaction containing the status, message, and transaction data.
func (m *transactionProtoMapper) ToProtoResponseTransaction(status string, message string, pbResponse *response.TransactionResponse) *pb.ApiResponseTransaction {
	return &pb.ApiResponseTransaction{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransaction(pbResponse),
	}
}

// ToProtoResponseTransactions converts a status, message, and a list of TransactionResponse
// to its protobuf representation.
//
// It includes the status, message, and mapped transaction data.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - pbResponse: The list of TransactionResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTransactions containing the status, message, and transaction data.
func (m *transactionProtoMapper) ToProtoResponseTransactions(status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponseTransactions {
	return &pb.ApiResponseTransactions{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransaction(pbResponse),
	}
}

// ToProtoResponseTransactionDelete generates a protobuf API response for a transaction delete operation.
//
// It includes the status and message for the API response.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//
// Returns:
//
//	A pointer to ApiResponseTransactionDelete containing the status and message.
func (m *transactionProtoMapper) ToProtoResponseTransactionDelete(status string, message string) *pb.ApiResponseTransactionDelete {
	return &pb.ApiResponseTransactionDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseTransactionAll generates a protobuf API response for all transactions.
//
// It includes the status and message for the API response.
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//
// Returns:
//
//	A pointer to ApiResponseTransactionAll containing the status and message.
func (m *transactionProtoMapper) ToProtoResponseTransactionAll(status string, message string) *pb.ApiResponseTransactionAll {
	return &pb.ApiResponseTransactionAll{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponsePaginationTransaction maps a pagination meta, status, message and a list of *response.TransactionResponse
// to a *pb.ApiResponsePaginationTransaction proto response.
//
// It is used to generate the response for the TransactionService.ListTransaction rpc method.
func (m *transactionProtoMapper) ToProtoResponsePaginationTransaction(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponsePaginationTransaction {

	return &pb.ApiResponsePaginationTransaction{
		Status:     status,
		Message:    message,
		Data:       m.mapResponsesTransaction(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationTransactionDeleteAt maps a pagination meta, status, message and a list of *response.TransactionResponseDeleteAt
// to a *pb.ApiResponsePaginationTransactionDeleteAt proto response.
//
// It is used to generate the response for the TransactionService.ListTransactionDeleteAt rpc method.
func (m *transactionProtoMapper) ToProtoResponsePaginationTransactionDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponseDeleteAt) *pb.ApiResponsePaginationTransactionDeleteAt {

	return &pb.ApiResponsePaginationTransactionDeleteAt{
		Status:     status,
		Message:    message,
		Data:       m.mapResponsesTransactionDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

// mapResponseTransaction maps a single response.TransactionResponse to its protobuf representation.
//
// It takes a response.TransactionResponse as input and returns a pointer to the converted TransactionResponse proto object.
func (m *transactionProtoMapper) mapResponseTransaction(transaction *response.TransactionResponse) *pb.TransactionResponse {
	return &pb.TransactionResponse{
		Id:              int32(transaction.ID),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int32(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantId:      int32(transaction.MerchantID),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
	}

}

// mapResponsesTransaction maps a slice of response.TransactionResponse to a slice of protobuf TransactionResponse.
//
// It takes a slice of response.TransactionResponse as input and returns a slice of corresponding protobuf TransactionResponse.
// The mapping includes fields like ID, TransactionNo, CardNumber, Amount, PaymentMethod, TransactionTime, MerchantId, CreatedAt, and UpdatedAt.
func (m *transactionProtoMapper) mapResponsesTransaction(transactions []*response.TransactionResponse) []*pb.TransactionResponse {
	var result []*pb.TransactionResponse
	for _, transaction := range transactions {
		result = append(result, m.mapResponseTransaction(transaction))
	}
	return result
}

// mapResponseTransactionDeleteAt maps a single response.TransactionResponseDeleteAt to its protobuf representation.
//
// It takes a response.TransactionResponseDeleteAt as input and returns a pointer to the converted TransactionResponseDeleteAt proto object.
// The mapping includes fields like ID, TransactionNo, CardNumber, Amount, PaymentMethod, TransactionTime, MerchantId, CreatedAt, UpdatedAt, and DeletedAt.
func (m *transactionProtoMapper) mapResponseTransactionDeleteAt(transaction *response.TransactionResponseDeleteAt) *pb.TransactionResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if transaction.DeletedAt != nil {
		deletedAt = wrapperspb.String(*transaction.DeletedAt)
	}

	return &pb.TransactionResponseDeleteAt{
		Id:              int32(transaction.ID),
		TransactionNo:   transaction.TransactionNo,
		CardNumber:      transaction.CardNumber,
		Amount:          int32(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		TransactionTime: transaction.TransactionTime,
		MerchantId:      int32(transaction.MerchantID),
		CreatedAt:       transaction.CreatedAt,
		UpdatedAt:       transaction.UpdatedAt,
		DeletedAt:       deletedAt,
	}

}

// mapResponsesTransactionDeleteAt maps a slice of response.TransactionResponseDeleteAt to a slice of protobuf TransactionResponseDeleteAt.
//
// Args:
//   - transactions: A slice of TransactionResponseDeleteAt to be converted.
//
// Returns:
//
//	A slice of pointers to TransactionResponseDeleteAt containing the mapped data.
func (m *transactionProtoMapper) mapResponsesTransactionDeleteAt(transactions []*response.TransactionResponseDeleteAt) []*pb.TransactionResponseDeleteAt {
	var result []*pb.TransactionResponseDeleteAt

	for _, transaction := range transactions {
		result = append(result, m.mapResponseTransactionDeleteAt(transaction))
	}
	return result
}

// mapResponseTransactionMonthStatusSuccess maps a response.TransactionResponseMonthStatusSuccess to its protobuf representation.
//
// It takes a response.TransactionResponseMonthStatusSuccess as input and returns a pointer to the converted TransactionMonthStatusSuccessResponse proto object.
// The mapping includes fields like Year, Month, TotalSuccess, and TotalAmount.
func (t *transactionProtoMapper) mapResponseTransactionMonthStatusSuccess(s *response.TransactionResponseMonthStatusSuccess) *pb.TransactionMonthStatusSuccessResponse {
	return &pb.TransactionMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapResponsesTransactionMonthStatusSuccess converts a slice of TransactionResponseMonthStatusSuccess
// to a slice of protobuf TransactionMonthStatusSuccessResponse.
//
// It iterates through each TransactionResponseMonthStatusSuccess, mapping it to its protobuf equivalent
// using the mapResponseTransactionMonthStatusSuccess function.
//
// Args:
//   - Transactions: A slice of TransactionResponseMonthStatusSuccess to be converted.
//
// Returns:
//
//	A slice of pointers to TransactionMonthStatusSuccessResponse containing the mapped data.
func (t *transactionProtoMapper) mapResponsesTransactionMonthStatusSuccess(Transactions []*response.TransactionResponseMonthStatusSuccess) []*pb.TransactionMonthStatusSuccessResponse {
	var TransactionRecords []*pb.TransactionMonthStatusSuccessResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.mapResponseTransactionMonthStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// mapTransactionResponseYearStatusSuccess maps a response.TransactionResponseYearStatusSuccess to its protobuf representation.
//
// It takes a response.TransactionResponseYearStatusSuccess as input and returns a pointer to the converted TransactionYearStatusSuccessResponse proto object.
// The mapping includes fields like Year, TotalSuccess, and TotalAmount.
func (t *transactionProtoMapper) mapTransactionResponseYearStatusSuccess(s *response.TransactionResponseYearStatusSuccess) *pb.TransactionYearStatusSuccessResponse {
	return &pb.TransactionYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapTransactionResponsesYearStatusSuccess maps a slice of TransactionResponseYearStatusSuccess to a slice of protobuf TransactionYearStatusSuccessResponse.
//
// It takes a slice of TransactionResponseYearStatusSuccess as input and returns a slice of corresponding protobuf TransactionYearStatusSuccessResponse.
// The mapping includes fields like Year, TotalSuccess, and TotalAmount.
func (t *transactionProtoMapper) mapTransactionResponsesYearStatusSuccess(Transactions []*response.TransactionResponseYearStatusSuccess) []*pb.TransactionYearStatusSuccessResponse {
	var TransactionRecords []*pb.TransactionYearStatusSuccessResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.mapTransactionResponseYearStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// mapResponseTransactionMonthStatusFailed maps a response.TransactionResponseMonthStatusFailed to its protobuf representation.
//
// It takes a response.TransactionResponseMonthStatusFailed as input and returns a pointer to the converted TransactionMonthStatusFailedResponse proto object.
// The mapping includes fields like Year, Month, TotalFailed, and TotalAmount.
func (t *transactionProtoMapper) mapResponseTransactionMonthStatusFailed(s *response.TransactionResponseMonthStatusFailed) *pb.TransactionMonthStatusFailedResponse {
	return &pb.TransactionMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponsesTransactionMonthStatusFailed converts a slice of TransactionResponseMonthStatusFailed
// to a slice of protobuf TransactionMonthStatusFailedResponse.
//
// It iterates through each TransactionResponseMonthStatusFailed, mapping it to its protobuf equivalent
// using the mapResponseTransactionMonthStatusFailed function.
//
// Args:
//   - Transactions: A slice of TransactionResponseMonthStatusFailed to be converted.
//
// Returns:
//
//	A slice of pointers to TransactionMonthStatusFailedResponse containing the mapped data.
func (t *transactionProtoMapper) mapResponsesTransactionMonthStatusFailed(Transactions []*response.TransactionResponseMonthStatusFailed) []*pb.TransactionMonthStatusFailedResponse {
	var TransactionRecords []*pb.TransactionMonthStatusFailedResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.mapResponseTransactionMonthStatusFailed(Transaction))
	}

	return TransactionRecords
}

// mapTransactionResponseYearStatusFailed maps a response.TransactionResponseYearStatusFailed to its protobuf representation.
//
// It takes a response.TransactionResponseYearStatusFailed as input and returns a pointer to the converted
// TransactionYearStatusFailedResponse proto object. The mapping includes fields like Year, TotalFailed, and TotalAmount.
//
// Args:
//   - s: A pointer to a TransactionResponseYearStatusFailed to be converted.
//
// Returns:
//   - A pointer to a TransactionYearStatusFailedResponse containing the mapped data.
func (t *transactionProtoMapper) mapTransactionResponseYearStatusFailed(s *response.TransactionResponseYearStatusFailed) *pb.TransactionYearStatusFailedResponse {
	return &pb.TransactionYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapTransactionResponsesYearStatusFailed maps a slice of TransactionResponseYearStatusFailed
// to a slice of protobuf TransactionYearStatusFailedResponse.
//
// It iterates through each TransactionResponseYearStatusFailed, mapping it to its protobuf equivalent
// using the mapTransactionResponseYearStatusFailed function.
//
// Args:
//   - Transactions: A slice of TransactionResponseYearStatusFailed to be converted.
//
// Returns:
//
//	A slice of pointers to TransactionYearStatusFailedResponse containing the mapped data.
func (t *transactionProtoMapper) mapTransactionResponsesYearStatusFailed(Transactions []*response.TransactionResponseYearStatusFailed) []*pb.TransactionYearStatusFailedResponse {
	var TransactionRecords []*pb.TransactionYearStatusFailedResponse

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.mapTransactionResponseYearStatusFailed(Transaction))
	}

	return TransactionRecords
}

// mapResponseTransactionMonthMethod maps a *response.TransactionMonthMethodResponse to a *pb.TransactionMonthMethodResponse proto message.
//
// It takes a *response.TransactionMonthMethodResponse as input and returns a pointer to the converted TransactionMonthMethodResponse proto object.
// The mapping includes fields like Month, PaymentMethod, TotalTransactions, and TotalAmount.
func (m *transactionProtoMapper) mapResponseTransactionMonthMethod(s *response.TransactionMonthMethodResponse) *pb.TransactionMonthMethodResponse {
	return &pb.TransactionMonthMethodResponse{
		Month:             s.Month,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int32(s.TotalTransactions),
		TotalAmount:       int32(s.TotalAmount),
	}
}

// mapResponseTransactionMonthMethods maps a slice of TransactionMonthMethodResponse to a slice of protobuf TransactionMonthMethodResponse.
//
// It takes a slice of TransactionMonthMethodResponse as input and returns a slice of corresponding protobuf TransactionMonthMethodResponse.
// The mapping includes fields like Month, PaymentMethod, TotalTransactions, and TotalAmount.
func (m *transactionProtoMapper) mapResponseTransactionMonthMethods(s []*response.TransactionMonthMethodResponse) []*pb.TransactionMonthMethodResponse {
	var protoResponses []*pb.TransactionMonthMethodResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.mapResponseTransactionMonthMethod(transaction))
	}
	return protoResponses
}

// mapResponseTransactionYearMethod maps a *response.TransactionYearMethodResponse to a *pb.TransactionYearMethodResponse proto message.
//
// It takes a *response.TransactionYearMethodResponse as input and returns a pointer to the converted TransactionYearMethodResponse proto object.
// The mapping includes fields like Year, PaymentMethod, TotalTransactions, and TotalAmount.
func (m *transactionProtoMapper) mapResponseTransactionYearMethod(s *response.TransactionYearMethodResponse) *pb.TransactionYearMethodResponse {
	return &pb.TransactionYearMethodResponse{
		Year:              s.Year,
		PaymentMethod:     s.PaymentMethod,
		TotalTransactions: int32(s.TotalTransactions),
		TotalAmount:       int32(s.TotalAmount),
	}
}

// mapResponseTransactionYearMethods maps a slice of TransactionYearMethodResponse to a slice of protobuf TransactionYearMethodResponse.
//
// It takes a slice of TransactionYearMethodResponse as input and returns a slice of corresponding protobuf TransactionYearMethodResponse.
// The mapping includes fields like Year, PaymentMethod, TotalTransactions, and TotalAmount.
func (m *transactionProtoMapper) mapResponseTransactionYearMethods(s []*response.TransactionYearMethodResponse) []*pb.TransactionYearMethodResponse {
	var protoResponses []*pb.TransactionYearMethodResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.mapResponseTransactionYearMethod(transaction))
	}
	return protoResponses
}

// mapResponseTransactionMonthAmount maps a *response.TransactionMonthAmountResponse to a *pb.TransactionMonthAmountResponse proto message.
//
// It takes a *response.TransactionMonthAmountResponse as input and returns a pointer to the converted TransactionMonthAmountResponse proto object.
// The mapping includes fields like Month and TotalAmount.
func (m *transactionProtoMapper) mapResponseTransactionMonthAmount(s *response.TransactionMonthAmountResponse) *pb.TransactionMonthAmountResponse {
	return &pb.TransactionMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTransactionMonthAmounts maps a slice of TransactionMonthAmountResponse
// to a slice of protobuf TransactionMonthAmountResponse.
//
// It iterates through each TransactionMonthAmountResponse, converting it to its
// protobuf equivalent using the mapResponseTransactionMonthAmount function.
//
// Args:
//   - s: A slice of TransactionMonthAmountResponse to be converted.
//
// Returns:
//   - A slice of pointers to TransactionMonthAmountResponse containing the mapped data.
func (m *transactionProtoMapper) mapResponseTransactionMonthAmounts(s []*response.TransactionMonthAmountResponse) []*pb.TransactionMonthAmountResponse {
	var protoResponses []*pb.TransactionMonthAmountResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.mapResponseTransactionMonthAmount(transaction))
	}
	return protoResponses
}

// mapResponseTransactionYearlyAmount maps a *response.TransactionYearlyAmountResponse to a *pb.TransactionYearlyAmountResponse proto message.
//
// It takes a *response.TransactionYearlyAmountResponse as input and returns a pointer to the converted TransactionYearlyAmountResponse proto object.
// The mapping includes fields like Year and TotalAmount.
func (m *transactionProtoMapper) mapResponseTransactionYearlyAmount(s *response.TransactionYearlyAmountResponse) *pb.TransactionYearlyAmountResponse {
	return &pb.TransactionYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTransactionYearlyAmounts maps a slice of TransactionYearlyAmountResponse
// to a slice of protobuf TransactionYearlyAmountResponse.
//
// It iterates through each TransactionYearlyAmountResponse, converting it to its
// protobuf equivalent using the mapResponseTransactionYearlyAmount function.
//
// Args:
//   - s: A slice of TransactionYearlyAmountResponse to be converted.
//
// Returns:
//   - A slice of pointers to TransactionYearlyAmountResponse containing the mapped data.
func (m *transactionProtoMapper) mapResponseTransactionYearlyAmounts(s []*response.TransactionYearlyAmountResponse) []*pb.TransactionYearlyAmountResponse {
	var protoResponses []*pb.TransactionYearlyAmountResponse
	for _, transaction := range s {
		protoResponses = append(protoResponses, m.mapResponseTransactionYearlyAmount(transaction))
	}
	return protoResponses
}
