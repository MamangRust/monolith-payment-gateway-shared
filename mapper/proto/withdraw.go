package protomapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type withdrawProtoMapper struct {
}

// NewWithdrawProtoMapper creates a new WithdrawProtoMapper.
func NewWithdrawProtoMapper() *withdrawProtoMapper {
	return &withdrawProtoMapper{}
}

// ToProtoResponseWithdraw maps a status, message, and a WithdrawResponse
// to an ApiResponseWithdraw proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - withdraw: The WithdrawResponse that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdraw containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponseWithdraw(status string, message string, withdraw *response.WithdrawResponse) *pb.ApiResponseWithdraw {
	return &pb.ApiResponseWithdraw{
		Status:  status,
		Message: message,
		Data:    m.mapResponseWithdrawal(withdraw),
	}
}

// ToProtoResponsesWithdraw maps a status, message, and a list of WithdrawResponse
// to an ApiResponsesWithdraw proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - withdraws: The list of WithdrawResponse that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponsesWithdraw containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponsesWithdraw(status string, message string, pbResponse []*response.WithdrawResponse) *pb.ApiResponsesWithdraw {
	return &pb.ApiResponsesWithdraw{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesWithdrawal(pbResponse),
	}
}

// ToProtoResponseWithdrawDelete maps a status and message to an ApiResponseWithdrawDelete proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawDelete containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponseWithdrawDelete(status string, message string) *pb.ApiResponseWithdrawDelete {
	return &pb.ApiResponseWithdrawDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseWithdrawAll maps a status and message to an ApiResponseWithdrawAll proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawAll containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponseWithdrawAll(status string, message string) *pb.ApiResponseWithdrawAll {
	return &pb.ApiResponseWithdrawAll{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponsePaginationWithdraw maps a status, message, a pagination meta, and a list of WithdrawResponse
// to an ApiResponsePaginationWithdraw proto message.
//
// Args:
//   - pagination: The pagination meta that needs to be converted.
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - withdraws: The list of WithdrawResponse that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponsePaginationWithdraw containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponsePaginationWithdraw(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponse) *pb.ApiResponsePaginationWithdraw {
	return &pb.ApiResponsePaginationWithdraw{
		Status:     status,
		Message:    message,
		Data:       m.mapResponsesWithdrawal(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationWithdrawDeleteAt maps a pagination meta, status, message, and a list of WithdrawResponseDeleteAt
// to an ApiResponsePaginationWithdrawDeleteAt proto message.
//
// Args:
//   - pagination: The pagination meta that needs to be converted.
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - withdraws: The list of WithdrawResponseDeleteAt that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponsePaginationWithdrawDeleteAt containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponsePaginationWithdrawDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponseDeleteAt) *pb.ApiResponsePaginationWithdrawDeleteAt {
	return &pb.ApiResponsePaginationWithdrawDeleteAt{
		Status:     status,
		Message:    message,
		Data:       m.mapResponsesWithdrawalDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponseWithdrawMonthStatusSuccess maps a status, message, and a list of WithdrawResponseMonthStatusSuccess
// to an ApiResponseWithdrawMonthStatusSuccess proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawMonthStatusSuccess containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponseWithdrawMonthStatusSuccess(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusSuccess) *pb.ApiResponseWithdrawMonthStatusSuccess {
	return &pb.ApiResponseWithdrawMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesWithdrawMonthStatusSuccess(pbResponse),
	}
}

// ToProtoResponseWithdrawYearStatusSuccess maps a status, message, and a list of WithdrawResponseYearStatusSuccess
// to an ApiResponseWithdrawYearStatusSuccess proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawYearStatusSuccess containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponseWithdrawYearStatusSuccess(status string, message string, pbResponse []*response.WithdrawResponseYearStatusSuccess) *pb.ApiResponseWithdrawYearStatusSuccess {
	return &pb.ApiResponseWithdrawYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapWithdrawResponsesYearStatusSuccess(pbResponse),
	}
}

// ToProtoResponseWithdrawMonthStatusFailed maps a status, message, and a list of WithdrawResponseMonthStatusFailed
// to an ApiResponseWithdrawMonthStatusFailed proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawMonthStatusFailed containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponseWithdrawMonthStatusFailed(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusFailed) *pb.ApiResponseWithdrawMonthStatusFailed {
	return &pb.ApiResponseWithdrawMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesWithdrawMonthStatusFailed(pbResponse),
	}
}

// ToProtoResponseWithdrawYearStatusFailed maps a status, message, and a list of WithdrawResponseYearStatusFailed
// to an ApiResponseWithdrawYearStatusFailed proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawResponseYearStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawYearStatusFailed containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponseWithdrawYearStatusFailed(status string, message string, pbResponse []*response.WithdrawResponseYearStatusFailed) *pb.ApiResponseWithdrawYearStatusFailed {
	return &pb.ApiResponseWithdrawYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapWithdrawResponsesYearStatusFailed(pbResponse),
	}
}

// ToProtoResponseWithdrawMonthAmount maps a status, message, and a list of WithdrawMonthlyAmountResponse
// to an ApiResponseWithdrawMonthAmount proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawMonthlyAmountResponse that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawMonthAmount containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponseWithdrawMonthAmount(status string, message string, pbResponse []*response.WithdrawMonthlyAmountResponse) *pb.ApiResponseWithdrawMonthAmount {
	return &pb.ApiResponseWithdrawMonthAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseWithdrawMonthlyAmounts(pbResponse),
	}
}

// ToProtoResponseWithdrawYearAmount maps a status, message, and a list of WithdrawYearlyAmountResponse
// to an ApiResponseWithdrawYearAmount proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The message accompanying the response.
//   - pbResponse: The list of WithdrawYearlyAmountResponse that needs to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseWithdrawYearAmount containing the mapped data.
func (m *withdrawProtoMapper) ToProtoResponseWithdrawYearAmount(status string, message string, pbResponse []*response.WithdrawYearlyAmountResponse) *pb.ApiResponseWithdrawYearAmount {
	return &pb.ApiResponseWithdrawYearAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseWithdrawYearlyAmounts(pbResponse),
	}
}

// mapResponseWithdrawal maps a single WithdrawResponse to its protobuf representation.
//
// Args:
//   - withdraw: The WithdrawResponse that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponse protobuf message containing the mapped data.
func (w *withdrawProtoMapper) mapResponseWithdrawal(withdraw *response.WithdrawResponse) *pb.WithdrawResponse {
	return &pb.WithdrawResponse{
		WithdrawId:     int32(withdraw.ID),
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int32(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
	}
}

// mapResponsesWithdrawal maps a slice of WithdrawResponse to a slice of protobuf WithdrawResponse.
//
// It takes a slice of WithdrawResponse as input and returns a slice of corresponding protobuf WithdrawResponse.
// The mapping includes fields like WithdrawId, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.
func (w *withdrawProtoMapper) mapResponsesWithdrawal(withdraws []*response.WithdrawResponse) []*pb.WithdrawResponse {
	var responseWithdraws []*pb.WithdrawResponse

	for _, withdraw := range withdraws {
		responseWithdraws = append(responseWithdraws, w.mapResponseWithdrawal(withdraw))
	}

	return responseWithdraws
}

// mapResponseWithdrawalDeleteAt maps a single WithdrawResponseDeleteAt to its protobuf representation.
//
// Args:
//   - withdraw: The WithdrawResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseDeleteAt protobuf message containing the mapped data.
func (w *withdrawProtoMapper) mapResponseWithdrawalDeleteAt(withdraw *response.WithdrawResponseDeleteAt) *pb.WithdrawResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if withdraw.DeletedAt != nil {
		deletedAt = wrapperspb.String(*withdraw.DeletedAt)
	}

	return &pb.WithdrawResponseDeleteAt{
		WithdrawId:     int32(withdraw.ID),
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int32(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
		DeletedAt:      deletedAt,
	}
}

// mapResponsesWithdrawalDeleteAt maps a slice of WithdrawResponseDeleteAt to a slice of protobuf WithdrawResponseDeleteAt.
//
// It takes a slice of WithdrawResponseDeleteAt as input and returns a slice of corresponding protobuf WithdrawResponseDeleteAt.
// The mapping includes fields like WithdrawId, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (w *withdrawProtoMapper) mapResponsesWithdrawalDeleteAt(withdraws []*response.WithdrawResponseDeleteAt) []*pb.WithdrawResponseDeleteAt {
	var responseWithdraws []*pb.WithdrawResponseDeleteAt

	for _, withdraw := range withdraws {
		responseWithdraws = append(responseWithdraws, w.mapResponseWithdrawalDeleteAt(withdraw))
	}

	return responseWithdraws
}

// mapResponseWithdrawMonthStatusSuccess maps a single WithdrawResponseMonthStatusSuccess to its protobuf representation.
//
// Args:
//   - s: The WithdrawResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawMonthStatusSuccessResponse protobuf message containing the mapped data.
func (t *withdrawProtoMapper) mapResponseWithdrawMonthStatusSuccess(s *response.WithdrawResponseMonthStatusSuccess) *pb.WithdrawMonthStatusSuccessResponse {
	return &pb.WithdrawMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapResponsesWithdrawMonthStatusSuccess maps a slice of WithdrawResponseMonthStatusSuccess to a slice of protobuf WithdrawMonthStatusSuccessResponse.
//
// It takes a slice of WithdrawResponseMonthStatusSuccess as input and returns a slice of corresponding protobuf WithdrawMonthStatusSuccessResponse.
// The mapping includes fields like Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawProtoMapper) mapResponsesWithdrawMonthStatusSuccess(Withdraws []*response.WithdrawResponseMonthStatusSuccess) []*pb.WithdrawMonthStatusSuccessResponse {
	var WithdrawRecords []*pb.WithdrawMonthStatusSuccessResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapResponseWithdrawMonthStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// mapWithdrawResponseYearStatusSuccess maps a single WithdrawResponseYearStatusSuccess to a WithdrawYearStatusSuccessResponse proto message.
//
// Args:
//   - s: The WithdrawResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawYearStatusSuccessResponse containing the mapped data, with fields Year, TotalSuccess, and TotalAmount.
func (t *withdrawProtoMapper) mapWithdrawResponseYearStatusSuccess(s *response.WithdrawResponseYearStatusSuccess) *pb.WithdrawYearStatusSuccessResponse {
	return &pb.WithdrawYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapWithdrawResponsesYearStatusSuccess maps a slice of WithdrawResponseYearStatusSuccess to a slice of protobuf WithdrawYearStatusSuccessResponse.
//
// It takes a slice of WithdrawResponseYearStatusSuccess as input and returns a slice of corresponding protobuf WithdrawYearStatusSuccessResponse.
// The mapping includes fields like Year, TotalSuccess, and TotalAmount.
func (t *withdrawProtoMapper) mapWithdrawResponsesYearStatusSuccess(Withdraws []*response.WithdrawResponseYearStatusSuccess) []*pb.WithdrawYearStatusSuccessResponse {
	var WithdrawRecords []*pb.WithdrawYearStatusSuccessResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapWithdrawResponseYearStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// mapResponseWithdrawMonthStatusFailed maps a single WithdrawResponseMonthStatusFailed to its protobuf representation.
//
// Args:
//   - s: The WithdrawResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawMonthStatusFailedResponse containing the mapped data, with fields Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawProtoMapper) mapResponseWithdrawMonthStatusFailed(s *response.WithdrawResponseMonthStatusFailed) *pb.WithdrawMonthStatusFailedResponse {
	return &pb.WithdrawMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponsesWithdrawMonthStatusFailed maps a slice of WithdrawResponseMonthStatusFailed to a slice of protobuf WithdrawMonthStatusFailedResponse.
//
// It takes a slice of WithdrawResponseMonthStatusFailed as input and returns a slice of corresponding protobuf WithdrawMonthStatusFailedResponse.
// The mapping includes fields like Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawProtoMapper) mapResponsesWithdrawMonthStatusFailed(Withdraws []*response.WithdrawResponseMonthStatusFailed) []*pb.WithdrawMonthStatusFailedResponse {
	var WithdrawRecords []*pb.WithdrawMonthStatusFailedResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapResponseWithdrawMonthStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// mapWithdrawResponseYearStatusFailed maps a WithdrawResponseYearStatusFailed to a WithdrawYearStatusFailedResponse proto message.
//
// Args:
//   - s: The WithdrawResponseYearStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawYearStatusFailedResponse containing the mapped data, with fields Year, TotalFailed, and TotalAmount.
func (t *withdrawProtoMapper) mapWithdrawResponseYearStatusFailed(s *response.WithdrawResponseYearStatusFailed) *pb.WithdrawYearStatusFailedResponse {
	return &pb.WithdrawYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapWithdrawResponsesYearStatusFailed maps a slice of WithdrawResponseYearStatusFailed to a slice of protobuf WithdrawYearStatusFailedResponse.
//
// It takes a slice of WithdrawResponseYearStatusFailed as input and returns a slice of corresponding protobuf WithdrawYearStatusFailedResponse.
// The mapping includes fields like Year, TotalFailed, and TotalAmount.
func (t *withdrawProtoMapper) mapWithdrawResponsesYearStatusFailed(Withdraws []*response.WithdrawResponseYearStatusFailed) []*pb.WithdrawYearStatusFailedResponse {
	var WithdrawRecords []*pb.WithdrawYearStatusFailedResponse

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapWithdrawResponseYearStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// mapResponseWithdrawMonthlyAmount maps a WithdrawMonthlyAmountResponse to its protobuf representation.
//
// Args:
//   - s: The WithdrawMonthlyAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawMonthlyAmountResponse containing the mapped data, with fields Month and TotalAmount.
func (m *withdrawProtoMapper) mapResponseWithdrawMonthlyAmount(s *response.WithdrawMonthlyAmountResponse) *pb.WithdrawMonthlyAmountResponse {
	return &pb.WithdrawMonthlyAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseWithdrawMonthlyAmounts maps a slice of WithdrawMonthlyAmountResponse to a slice of protobuf WithdrawMonthlyAmountResponse.
//
// It takes a slice of WithdrawMonthlyAmountResponse as input and returns a slice of corresponding protobuf WithdrawMonthlyAmountResponse.
// The mapping includes fields like Month and TotalAmount.
func (m *withdrawProtoMapper) mapResponseWithdrawMonthlyAmounts(s []*response.WithdrawMonthlyAmountResponse) []*pb.WithdrawMonthlyAmountResponse {
	var protoResponses []*pb.WithdrawMonthlyAmountResponse
	for _, withdraw := range s {
		protoResponses = append(protoResponses, m.mapResponseWithdrawMonthlyAmount(withdraw))
	}
	return protoResponses
}

// mapResponseWithdrawYearlyAmount maps a WithdrawYearlyAmountResponse to its protobuf representation.
//
// Args:
//   - s: The WithdrawYearlyAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawYearlyAmountResponse containing the mapped data, with fields Year and TotalAmount.
func (m *withdrawProtoMapper) mapResponseWithdrawYearlyAmount(s *response.WithdrawYearlyAmountResponse) *pb.WithdrawYearlyAmountResponse {
	return &pb.WithdrawYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseWithdrawYearlyAmounts maps a slice of WithdrawYearlyAmountResponse to a slice of protobuf WithdrawYearlyAmountResponse.
//
// It takes a slice of WithdrawYearlyAmountResponse as input and returns a slice of corresponding protobuf WithdrawYearlyAmountResponse.
// The mapping includes fields like Year and TotalAmount.
func (m *withdrawProtoMapper) mapResponseWithdrawYearlyAmounts(s []*response.WithdrawYearlyAmountResponse) []*pb.WithdrawYearlyAmountResponse {
	var protoResponses []*pb.WithdrawYearlyAmountResponse
	for _, withdraw := range s {
		protoResponses = append(protoResponses, m.mapResponseWithdrawYearlyAmount(withdraw))
	}
	return protoResponses
}
