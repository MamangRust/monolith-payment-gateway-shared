package protomapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type merchantProtoMapper struct{}

// NewMerchantProtoMapper initializes and returns a new instance of merchantProtoMapper.
func NewMerchantProtoMapper() *merchantProtoMapper {
	return &merchantProtoMapper{}
}

// ToProtoResponsePaginationMerchant converts a list of MerchantResponse and pagination metadata
// into a protobuf ApiResponsePaginationMerchant. It includes the status and message for the API response.
// Parameters:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//   - merchants: A list of MerchantResponse objects to be included in the response.
//
// Returns:
//
//	A pointer to ApiResponsePaginationMerchant containing the status, message, merchant data, and pagination data.
func (m *merchantProtoMapper) ToProtoResponsePaginationMerchant(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantResponse) *pb.ApiResponsePaginationMerchant {
	return &pb.ApiResponsePaginationMerchant{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantResponses(merchants),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationMerchantDeleteAt maps a pagination meta, status, message and a list of *response.MerchantResponseDeleteAt
// to a *pb.ApiResponsePaginationMerchantDeleteAt proto response.
//
// It is used to generate the response for the MerchantService.ListDeletedMerchant rpc method.
func (m *merchantProtoMapper) ToProtoResponsePaginationMerchantDeleteAt(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantResponseDeleteAt) *pb.ApiResponsePaginationMerchantDeleteAt {
	return &pb.ApiResponsePaginationMerchantDeleteAt{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantResponsesDeleteAt(merchants),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationMerchantTransaction maps a pagination meta, status, message and a list of *response.MerchantTransactionResponse
// to a *pb.ApiResponsePaginationMerchantTransaction proto response.
//
// It is used to generate the response for the MerchantService.ListMerchantTransaction rpc method.
func (m *merchantProtoMapper) ToProtoResponsePaginationMerchantTransaction(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantTransactionResponse) *pb.ApiResponsePaginationMerchantTransaction {

	return &pb.ApiResponsePaginationMerchantTransaction{
		Status:     status,
		Message:    message,
		Data:       m.mapMerchantTransactionResponses(merchants),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponseMonthlyPaymentMethods maps a status, message and a list of *response.MerchantResponseMonthlyPaymentMethod
// to a *pb.ApiResponseMerchantMonthlyPaymentMethod proto response.
//
// It is used to generate the response for the MerchantService.GetMonthlyPaymentMethod rpc method.
func (m *merchantProtoMapper) ToProtoResponseMonthlyPaymentMethods(status string, message string, ms []*response.MerchantResponseMonthlyPaymentMethod) *pb.ApiResponseMerchantMonthlyPaymentMethod {
	return &pb.ApiResponseMerchantMonthlyPaymentMethod{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesMonthlyPaymentMethod(ms),
	}
}

// ToProtoResponseYearlyPaymentMethods maps a status, message and a list of *response.MerchantResponseYearlyPaymentMethod
// to a *pb.ApiResponseMerchantYearlyPaymentMethod proto response.
//
// It is used to generate the response for the MerchantService.GetYearlyPaymentMethod rpc method.
func (m *merchantProtoMapper) ToProtoResponseYearlyPaymentMethods(status string, message string, ms []*response.MerchantResponseYearlyPaymentMethod) *pb.ApiResponseMerchantYearlyPaymentMethod {
	return &pb.ApiResponseMerchantYearlyPaymentMethod{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesYearlyPaymentMethod(ms),
	}
}

// ToProtoResponseMonthlyAmounts maps a status, message and a list of *response.MerchantResponseMonthlyAmount
// to a *pb.ApiResponseMerchantMonthlyAmount proto response.
//
// It is used to generate the response for the MerchantService.GetMonthlyAmount rpc method.
func (m *merchantProtoMapper) ToProtoResponseMonthlyAmounts(status string, message string, ms []*response.MerchantResponseMonthlyAmount) *pb.ApiResponseMerchantMonthlyAmount {
	return &pb.ApiResponseMerchantMonthlyAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesMonthlyAmount(ms),
	}
}

// ToProtoResponseYearlyAmounts maps a status, message and a list of *response.MerchantResponseYearlyAmount
// to a *pb.ApiResponseMerchantYearlyAmount proto response.
//
// It is used to generate the response for the MerchantService.GetYearlyAmount rpc method.
func (m *merchantProtoMapper) ToProtoResponseYearlyAmounts(status string, message string, ms []*response.MerchantResponseYearlyAmount) *pb.ApiResponseMerchantYearlyAmount {
	return &pb.ApiResponseMerchantYearlyAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesYearlyAmount(ms),
	}
}

// ToProtoResponseMonthlyTotalAmounts maps a status, message and a list of *response.MerchantResponseMonthlyTotalAmount
// to a *pb.ApiResponseMerchantMonthlyTotalAmount proto response.
//
// It is used to generate the response for the MerchantService.GetMonthlyTotalAmount rpc method.
func (m *merchantProtoMapper) ToProtoResponseMonthlyTotalAmounts(status string, message string, ms []*response.MerchantResponseMonthlyTotalAmount) *pb.ApiResponseMerchantMonthlyTotalAmount {
	return &pb.ApiResponseMerchantMonthlyTotalAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesMonthlyTotalAmount(ms),
	}
}

// ToProtoResponseYearlyTotalAmounts maps a status, message, and a list of *response.MerchantResponseYearlyTotalAmount
// to a *pb.ApiResponseMerchantYearlyTotalAmount proto response.
//
// It is used to generate the response for the MerchantService.GetYearlyTotalAmount rpc method.
func (m *merchantProtoMapper) ToProtoResponseYearlyTotalAmounts(status string, message string, ms []*response.MerchantResponseYearlyTotalAmount) *pb.ApiResponseMerchantYearlyTotalAmount {
	return &pb.ApiResponseMerchantYearlyTotalAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesYearlyTotalAmount(ms),
	}
}

// ToProtoResponseMerchant maps a status, message and a *response.MerchantResponse
// to a *pb.ApiResponseMerchant proto response.
//
// It is used to generate the response for the MerchantService.GetMerchant rpc method.
func (m *merchantProtoMapper) ToProtoResponseMerchant(status string, message string, res *response.MerchantResponse) *pb.ApiResponseMerchant {
	return &pb.ApiResponseMerchant{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantResponse(res),
	}

}

// ToProtoResponseMerchants maps a status, message and a list of *response.MerchantResponse
// to a *pb.ApiResponsesMerchant proto response.
//
// It is used to generate the response for the MerchantService.ListMerchants rpc method.
func (m *merchantProtoMapper) ToProtoResponseMerchants(status string, message string, res []*response.MerchantResponse) *pb.ApiResponsesMerchant {
	return &pb.ApiResponsesMerchant{
		Status:  status,
		Message: message,
		Data:    m.mapMerchantResponses(res),
	}

}

// ToProtoResponseMerchantAll maps a status and message to its protobuf representation
// specifically for bulk merchant operations.
// It includes the status and message of the API response.
// Parameters:
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//
// Returns:
//
//	A pointer to ApiResponseMerchantAll containing the status and message.
func (m *merchantProtoMapper) ToProtoResponseMerchantAll(status string, message string) *pb.ApiResponseMerchantAll {
	return &pb.ApiResponseMerchantAll{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseMerchantDelete converts a status and message to its protobuf representation
// specifically for merchant deletion operations.
// It includes the status and message of the API response.
// Parameters:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//
// Returns:
//
//	A pointer to ApiResponseMerchantDelete containing the status and message.
func (m *merchantProtoMapper) ToProtoResponseMerchantDelete(status string, message string) *pb.ApiResponseMerchantDelete {
	return &pb.ApiResponseMerchantDelete{
		Status:  status,
		Message: message,
	}
}

// mapMerchantResponse maps a *response.MerchantResponse to a *pb.MerchantResponse proto message.
//
// It is used to generate the response for the MerchantService.GetMerchant rpc method.
func (m *merchantProtoMapper) mapMerchantResponse(merchant *response.MerchantResponse) *pb.MerchantResponse {
	return &pb.MerchantResponse{
		Id:        int32(merchant.ID),
		Name:      merchant.Name,
		Status:    merchant.Status,
		ApiKey:    merchant.ApiKey,
		UserId:    int32(merchant.UserID),
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
	}
}

// mapMerchantResponses maps a list of *response.MerchantResponse to a list of
// *pb.MerchantResponse proto responses.
//
// It iterates over each MerchantResponse in the input slice, converting
// them to their protobuf equivalent using the mapMerchantResponse function.
// This function is used to generate the response data for merchant-related RPC methods.
func (s *merchantProtoMapper) mapMerchantResponses(roles []*response.MerchantResponse) []*pb.MerchantResponse {
	var responseRoles []*pb.MerchantResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapMerchantResponse(role))
	}

	return responseRoles
}

// mapMerchantResponseDeleteAt maps a *response.MerchantResponseDeleteAt to a *pb.MerchantResponseDeleteAt proto message.
//
// It is used to generate the response for the MerchantService.ListDeletedMerchant rpc method.
func (m *merchantProtoMapper) mapMerchantResponseDeleteAt(merchant *response.MerchantResponseDeleteAt) *pb.MerchantResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if merchant.DeletedAt != nil {
		deletedAt = wrapperspb.String(*merchant.DeletedAt)
	}

	return &pb.MerchantResponseDeleteAt{
		Id:        int32(merchant.ID),
		Name:      merchant.Name,
		Status:    merchant.Status,
		UserId:    int32(merchant.UserID),
		ApiKey:    merchant.ApiKey,
		CreatedAt: merchant.CreatedAt,
		UpdatedAt: merchant.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

// mapMerchantResponsesDeleteAt maps a list of *response.MerchantResponseDeleteAt to a list of
// *pb.MerchantResponseDeleteAt proto responses.
//
// It iterates over each MerchantResponseDeleteAt in the input slice, converting
// them to their protobuf equivalent using the mapMerchantResponseDeleteAt function.
// This function is used to generate the response data for the MerchantService.ListDeletedMerchant rpc method.
func (s *merchantProtoMapper) mapMerchantResponsesDeleteAt(roles []*response.MerchantResponseDeleteAt) []*pb.MerchantResponseDeleteAt {
	var responseRoles []*pb.MerchantResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapMerchantResponseDeleteAt(role))
	}

	return responseRoles
}

// mapMerchantTransactionResponse maps a *response.MerchantTransactionResponse to a *pb.MerchantTransactionResponse proto message.
//
// It is used to generate the response for the MerchantService.ListMerchantTransaction rpc method.
func (m *merchantProtoMapper) mapMerchantTransactionResponse(merchant *response.MerchantTransactionResponse) *pb.MerchantTransactionResponse {
	return &pb.MerchantTransactionResponse{
		Id:              int32(merchant.ID),
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantId:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt,
		UpdatedAt:       merchant.UpdatedAt,
	}
}

// mapMerchantTransactionResponses maps a list of *response.MerchantTransactionResponse to a list of
// *pb.MerchantTransactionResponse proto responses.
//
// It iterates over each MerchantTransactionResponse in the input slice, converting
// them to their protobuf equivalent using the mapMerchantTransactionResponse function.
// This function is used to generate the response data for merchant transaction-related RPC methods.
func (s *merchantProtoMapper) mapMerchantTransactionResponses(roles []*response.MerchantTransactionResponse) []*pb.MerchantTransactionResponse {
	var responseRoles []*pb.MerchantTransactionResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapMerchantTransactionResponse(role))
	}

	return responseRoles
}

// mapResponseMonthlyPaymentMethod maps a *response.MerchantResponseMonthlyPaymentMethod to a *pb.MerchantResponseMonthlyPaymentMethod proto message.
//
// It is used to generate the response for the MerchantService.GetMonthlyPaymentMethod rpc method.
func (m *merchantProtoMapper) mapResponseMonthlyPaymentMethod(ms *response.MerchantResponseMonthlyPaymentMethod) *pb.MerchantResponseMonthlyPaymentMethod {
	return &pb.MerchantResponseMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int64(ms.TotalAmount),
	}
}

// mapResponsesMonthlyPaymentMethod maps a list of *response.MerchantResponseMonthlyPaymentMethod to a list of
// *pb.MerchantResponseMonthlyPaymentMethod proto responses.
//
// It iterates over each MerchantResponseMonthlyPaymentMethod in the input slice, converting
// them to their protobuf equivalent using the mapResponseMonthlyPaymentMethod function. This
// function is used to generate the response data for monthly payment method RPC methods.
func (s *merchantProtoMapper) mapResponsesMonthlyPaymentMethod(roles []*response.MerchantResponseMonthlyPaymentMethod) []*pb.MerchantResponseMonthlyPaymentMethod {
	var responseRoles []*pb.MerchantResponseMonthlyPaymentMethod

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseMonthlyPaymentMethod(role))
	}

	return responseRoles
}

// mapResponseYearlyPaymentMethod maps a *response.MerchantResponseYearlyPaymentMethod to a
// *pb.MerchantResponseYearlyPaymentMethod proto message.
//
// It is used to generate the response for the MerchantService.GetYearlyPaymentMethod rpc method.
func (m *merchantProtoMapper) mapResponseYearlyPaymentMethod(ms *response.MerchantResponseYearlyPaymentMethod) *pb.MerchantResponseYearlyPaymentMethod {
	return &pb.MerchantResponseYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int64(ms.TotalAmount),
	}
}

// mapResponsesYearlyPaymentMethod maps a list of *response.MerchantResponseYearlyPaymentMethod to a list of
// *pb.MerchantResponseYearlyPaymentMethod proto responses.
//
// It iterates over each MerchantResponseYearlyPaymentMethod in the input slice, converting
// them to their protobuf equivalent using the mapResponseYearlyPaymentMethod function. This
// function is used to generate the response data for yearly payment method RPC methods.
func (s *merchantProtoMapper) mapResponsesYearlyPaymentMethod(roles []*response.MerchantResponseYearlyPaymentMethod) []*pb.MerchantResponseYearlyPaymentMethod {
	var responseRoles []*pb.MerchantResponseYearlyPaymentMethod

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseYearlyPaymentMethod(role))
	}

	return responseRoles
}

// mapResponseMonthlyAmount maps a *response.MerchantResponseMonthlyAmount to a *pb.MerchantResponseMonthlyAmount proto message.
//
// It is used to generate the response for the MerchantService.GetMonthlyAmount rpc method.
func (m *merchantProtoMapper) mapResponseMonthlyAmount(ms *response.MerchantResponseMonthlyAmount) *pb.MerchantResponseMonthlyAmount {
	return &pb.MerchantResponseMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int64(ms.TotalAmount),
	}
}

// mapResponsesMonthlyAmount maps a list of *response.MerchantResponseMonthlyAmount to a list of
// *pb.MerchantResponseMonthlyAmount proto responses.
//
// It iterates over each MerchantResponseMonthlyAmount in the input slice, converting
// them to their protobuf equivalent using the mapResponseMonthlyAmount function. This
// function is used to generate the response data for monthly amount RPC methods.
func (s *merchantProtoMapper) mapResponsesMonthlyAmount(roles []*response.MerchantResponseMonthlyAmount) []*pb.MerchantResponseMonthlyAmount {
	var responseRoles []*pb.MerchantResponseMonthlyAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseMonthlyAmount(role))
	}

	return responseRoles
}

// mapResponseYearlyAmount maps a *response.MerchantResponseYearlyAmount to a *pb.MerchantResponseYearlyAmount proto message.
//
// It is used to generate the response for the MerchantService.GetYearlyAmount rpc method.
func (m *merchantProtoMapper) mapResponseYearlyAmount(ms *response.MerchantResponseYearlyAmount) *pb.MerchantResponseYearlyAmount {
	return &pb.MerchantResponseYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int64(ms.TotalAmount),
	}
}

// mapResponsesYearlyAmount maps a list of *response.MerchantResponseYearlyAmount to a list of
// *pb.MerchantResponseYearlyAmount proto responses.
//
// It iterates over each MerchantResponseYearlyAmount in the input slice, converting
// them to their protobuf equivalent using the mapResponseYearlyAmount function.
// This function is used to generate the response data for yearly amount RPC methods.
func (s *merchantProtoMapper) mapResponsesYearlyAmount(roles []*response.MerchantResponseYearlyAmount) []*pb.MerchantResponseYearlyAmount {
	var responseRoles []*pb.MerchantResponseYearlyAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseYearlyAmount(role))
	}

	return responseRoles
}

// mapResponseMonthlyTotalAmount maps a *response.MerchantResponseMonthlyTotalAmount to a *pb.MerchantResponseMonthlyTotalAmount proto message.
//
// It is used to generate the response for the MerchantService.GetMonthlyTotalAmount rpc method.
func (m *merchantProtoMapper) mapResponseMonthlyTotalAmount(ms *response.MerchantResponseMonthlyTotalAmount) *pb.MerchantResponseMonthlyTotalAmount {
	return &pb.MerchantResponseMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int64(ms.TotalAmount),
	}
}

// mapResponsesMonthlyTotalAmount maps a list of *response.MerchantResponseMonthlyTotalAmount to a list of
// *pb.MerchantResponseMonthlyTotalAmount proto responses.
//
// It iterates over each MerchantResponseMonthlyTotalAmount in the input slice, converting
// them to their protobuf equivalent using the mapResponseMonthlyTotalAmount function.
// This function is used to generate the response data for monthly total amount RPC methods.
func (s *merchantProtoMapper) mapResponsesMonthlyTotalAmount(roles []*response.MerchantResponseMonthlyTotalAmount) []*pb.MerchantResponseMonthlyTotalAmount {
	var responseRoles []*pb.MerchantResponseMonthlyTotalAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseMonthlyTotalAmount(role))
	}

	return responseRoles
}

// mapResponseYearlyTotalAmount maps a *response.MerchantResponseYearlyTotalAmount to a
// *pb.MerchantResponseYearlyTotalAmount proto message.
//
// It is used to generate the response for the MerchantService.GetYearlyTotalAmount rpc method.
func (m *merchantProtoMapper) mapResponseYearlyTotalAmount(ms *response.MerchantResponseYearlyTotalAmount) *pb.MerchantResponseYearlyTotalAmount {
	return &pb.MerchantResponseYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: int64(ms.TotalAmount),
	}
}

// mapResponsesYearlyTotalAmount maps a list of *response.MerchantResponseYearlyTotalAmount to a list of
// *pb.MerchantResponseYearlyTotalAmount proto responses.
//
// It iterates over each MerchantResponseYearlyTotalAmount in the input slice, converting
// them to their protobuf equivalent using the mapResponseYearlyTotalAmount function.
// This function is used to generate the response data for yearly total amount RPC methods.
func (s *merchantProtoMapper) mapResponsesYearlyTotalAmount(roles []*response.MerchantResponseYearlyTotalAmount) []*pb.MerchantResponseYearlyTotalAmount {
	var responseRoles []*pb.MerchantResponseYearlyTotalAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseYearlyTotalAmount(role))
	}

	return responseRoles
}
