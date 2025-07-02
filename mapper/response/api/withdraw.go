package apimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// withdrawResponseMapper provides methods to map gRPC withdraw responses to HTTP API responses
type withdrawResponseMapper struct {
}

// NewWithdrawResponseMapper creates and returns a new instance of withdrawResponseMapper.
// This instance provides methods to map gRPC withdraw responses to HTTP API responses.
func NewWithdrawResponseMapper() *withdrawResponseMapper {
	return &withdrawResponseMapper{}
}

// ToApiResponseWithdraw maps a single withdraw gRPC response to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponseWithdraw containing the mapped data.
func (m *withdrawResponseMapper) ToApiResponseWithdraw(pbResponse *pb.ApiResponseWithdraw) *response.ApiResponseWithdraw {
	return &response.ApiResponseWithdraw{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseWithdrawal(pbResponse.Data),
	}
}

// ToApiResponsesWithdraw maps a list of withdraw gRPC responses to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponsesWithdraw containing the mapped data.
func (m *withdrawResponseMapper) ToApiResponsesWithdraw(pbResponse *pb.ApiResponsesWithdraw) *response.ApiResponsesWithdraw {
	return &response.ApiResponsesWithdraw{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesWithdrawal(pbResponse.Data),
	}
}

// ToApiResponseWithdrawDelete maps a gRPC response indicating a withdraw has been deleted to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponseWithdrawDelete containing the mapped data.
func (m *withdrawResponseMapper) ToApiResponseWithdrawDelete(pbResponse *pb.ApiResponseWithdrawDelete) *response.ApiResponseWithdrawDelete {
	return &response.ApiResponseWithdrawDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseWithdrawAll maps a gRPC response containing all withdraw records to an API response.
//
// Args:
//   - pbResponse: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to an ApiResponseWithdrawAll containing the mapped data.
func (m *withdrawResponseMapper) ToApiResponseWithdrawAll(pbResponse *pb.ApiResponseWithdrawAll) *response.ApiResponseWithdrawAll {
	return &response.ApiResponseWithdrawAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponsePaginationWithdraw maps a pagination meta, status, message, and a list of WithdrawResponse
// to a response.ApiResponsePaginationWithdraw proto message.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationWithdraw containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationWithdraw containing the mapped data.
func (m *withdrawResponseMapper) ToApiResponsePaginationWithdraw(pbResponse *pb.ApiResponsePaginationWithdraw) *response.ApiResponsePaginationWithdraw {
	return &response.ApiResponsePaginationWithdraw{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.mapResponsesWithdrawal(pbResponse.Data),
		Pagination: *mapPaginationMeta(pbResponse.Pagination),
	}
}

// ToApiResponsePaginationWithdrawDeleteAt maps a pagination meta, status, message, and a list of WithdrawResponseDeleteAt
// to a response.ApiResponsePaginationWithdrawDeleteAt proto message.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationWithdrawDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationWithdrawDeleteAt containing the mapped data.
func (m *withdrawResponseMapper) ToApiResponsePaginationWithdrawDeleteAt(pbResponse *pb.ApiResponsePaginationWithdrawDeleteAt) *response.ApiResponsePaginationWithdrawDeleteAt {
	return &response.ApiResponsePaginationWithdrawDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.mapResponsesWithdrawalDeleteAt(pbResponse.Data),
		Pagination: *mapPaginationMeta(pbResponse.Pagination),
	}
}

// ToApiResponseWithdrawMonthStatusSuccess converts a gRPC response containing monthly successful withdraw statistics
// into an API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawMonthStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawMonthStatusSuccess containing the mapped data including status, message,
//     and detailed information about the monthly successful withdraw statistics.
func (m *withdrawResponseMapper) ToApiResponseWithdrawMonthStatusSuccess(pbResponse *pb.ApiResponseWithdrawMonthStatusSuccess) *response.ApiResponseWithdrawMonthStatusSuccess {
	return &response.ApiResponseWithdrawMonthStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesWithdrawMonthStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseWithdrawYearStatusSuccess maps a gRPC response containing yearly successful withdraw statistics
// to an HTTP API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawYearStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawYearStatusSuccess containing the mapped data, including status, message,
//     and detailed information about the yearly successful withdraw statistics.
func (m *withdrawResponseMapper) ToApiResponseWithdrawYearStatusSuccess(pbResponse *pb.ApiResponseWithdrawYearStatusSuccess) *response.ApiResponseWithdrawYearStatusSuccess {
	return &response.ApiResponseWithdrawYearStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapWithdrawResponsesYearStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseWithdrawMonthStatusFailed maps a gRPC response containing monthly failed withdraw statistics
// into an API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawMonthStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawMonthStatusFailed containing the mapped data including status, message,
//     and detailed information about the monthly failed withdraw statistics.
func (m *withdrawResponseMapper) ToApiResponseWithdrawMonthStatusFailed(pbResponse *pb.ApiResponseWithdrawMonthStatusFailed) *response.ApiResponseWithdrawMonthStatusFailed {
	return &response.ApiResponseWithdrawMonthStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesWithdrawMonthStatusFailed(pbResponse.Data),
	}
}

// ToApiResponseWithdrawYearStatusFailed maps a gRPC response containing yearly failed withdraw statistics
// into an API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawYearStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawYearStatusFailed containing the mapped data including status, message,
//     and detailed information about the yearly failed withdraw statistics.
func (m *withdrawResponseMapper) ToApiResponseWithdrawYearStatusFailed(pbResponse *pb.ApiResponseWithdrawYearStatusFailed) *response.ApiResponseWithdrawYearStatusFailed {
	return &response.ApiResponseWithdrawYearStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapWithdrawResponsesYearStatusFailed(pbResponse.Data),
	}
}

// ToApiResponseWithdrawMonthAmount converts a gRPC response containing monthly withdraw amounts
// into an API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawMonthAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawMonthAmount containing the mapped data, including status, message,
//     and detailed information about the monthly withdraw amounts.
func (m *withdrawResponseMapper) ToApiResponseWithdrawMonthAmount(pbResponse *pb.ApiResponseWithdrawMonthAmount) *response.ApiResponseWithdrawMonthAmount {
	return &response.ApiResponseWithdrawMonthAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseWithdrawMonthlyAmounts(pbResponse.Data),
	}
}

// ToApiResponseWithdrawYearAmount maps a gRPC response containing yearly withdraw amounts
// into an API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseWithdrawYearAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseWithdrawYearAmount containing the mapped data, including status, message,
//     and detailed information about the yearly withdraw amounts.
func (m *withdrawResponseMapper) ToApiResponseWithdrawYearAmount(pbResponse *pb.ApiResponseWithdrawYearAmount) *response.ApiResponseWithdrawYearAmount {
	return &response.ApiResponseWithdrawYearAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseWithdrawYearlyAmounts(pbResponse.Data),
	}
}

// mapResponseWithdrawal maps a single withdraw gRPC response to an API response.
//
// Args:
//   - withdraw: The gRPC response that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponse containing the mapped data.
func (w *withdrawResponseMapper) mapResponseWithdrawal(withdraw *pb.WithdrawResponse) *response.WithdrawResponse {
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

// mapResponsesWithdrawal maps a slice of WithdrawResponse to a slice of WithdrawResponse.
//
// It takes a slice of WithdrawResponse as input and returns a slice of corresponding WithdrawResponse.
// The mapping includes fields like WithdrawId, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.
func (w *withdrawResponseMapper) mapResponsesWithdrawal(withdraws []*pb.WithdrawResponse) []*response.WithdrawResponse {
	var responseWithdraws []*response.WithdrawResponse

	for _, withdraw := range withdraws {
		responseWithdraws = append(responseWithdraws, w.mapResponseWithdrawal(withdraw))
	}

	return responseWithdraws
}

// mapResponseWithdrawalDeleteAt maps a single WithdrawResponseDeleteAt to an API response.
//
// Args:
//   - withdraw: The WithdrawResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseDeleteAt containing the mapped data, including id, withdraw_no, card_number,
//     withdraw_amount, withdraw_time, created_at, updated_at, and deleted_at.
func (w *withdrawResponseMapper) mapResponseWithdrawalDeleteAt(withdraw *pb.WithdrawResponseDeleteAt) *response.WithdrawResponseDeleteAt {
	var deletedAt string
	if withdraw.DeletedAt != nil {
		deletedAt = withdraw.DeletedAt.Value
	}

	return &response.WithdrawResponseDeleteAt{
		ID:             int(withdraw.WithdrawId),
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
		DeletedAt:      &deletedAt,
	}
}

// mapResponsesWithdrawalDeleteAt maps a slice of WithdrawResponseDeleteAt gRPC messages
// to a slice of API response WithdrawResponseDeleteAt objects.
//
// Args:
//   - withdraws: A slice of WithdrawResponseDeleteAt gRPC messages that need to be converted.
//
// Returns:
//   - A slice of WithdrawResponseDeleteAt API responses containing the mapped data.
func (w *withdrawResponseMapper) mapResponsesWithdrawalDeleteAt(withdraws []*pb.WithdrawResponseDeleteAt) []*response.WithdrawResponseDeleteAt {
	var responseWithdraws []*response.WithdrawResponseDeleteAt

	for _, withdraw := range withdraws {
		responseWithdraws = append(responseWithdraws, w.mapResponseWithdrawalDeleteAt(withdraw))
	}

	return responseWithdraws
}

func (t *withdrawResponseMapper) mapResponseWithdrawMonthStatusSuccess(s *pb.WithdrawMonthStatusSuccessResponse) *response.WithdrawResponseMonthStatusSuccess {
	return &response.WithdrawResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapResponsesWithdrawMonthStatusSuccess maps a slice of WithdrawMonthStatusSuccessResponse gRPC messages to a slice
// of API response WithdrawResponseMonthStatusSuccess objects.
//
// Args:
//   - Withdraws: A slice of WithdrawMonthStatusSuccessResponse gRPC messages that need to be converted.
//
// Returns:
//   - A slice of WithdrawResponseMonthStatusSuccess API responses containing the mapped data.
func (t *withdrawResponseMapper) mapResponsesWithdrawMonthStatusSuccess(Withdraws []*pb.WithdrawMonthStatusSuccessResponse) []*response.WithdrawResponseMonthStatusSuccess {
	var WithdrawRecords []*response.WithdrawResponseMonthStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapResponseWithdrawMonthStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// mapWithdrawResponseYearStatusSuccess maps a single WithdrawYearStatusSuccessResponse gRPC message
// to an API response WithdrawResponseYearStatusSuccess object.
//
// Args:
//   - s: The WithdrawYearStatusSuccessResponse gRPC message that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseYearStatusSuccess containing the mapped data, including
//     fields like Year, TotalSuccess, and TotalAmount.
func (t *withdrawResponseMapper) mapWithdrawResponseYearStatusSuccess(s *pb.WithdrawYearStatusSuccessResponse) *response.WithdrawResponseYearStatusSuccess {
	return &response.WithdrawResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapWithdrawResponsesYearStatusSuccess maps a slice of gRPC WithdrawYearStatusSuccessResponse messages
// to a slice of API WithdrawResponseYearStatusSuccess objects.
//
// Args:
//   - Withdraws: A slice of pointers to pb.WithdrawYearStatusSuccessResponse messages that need to be converted.
//
// Returns:
//   - A slice of pointers to response.WithdrawResponseYearStatusSuccess objects containing the mapped data,
//     including fields like Year, TotalSuccess, and TotalAmount.
func (t *withdrawResponseMapper) mapWithdrawResponsesYearStatusSuccess(Withdraws []*pb.WithdrawYearStatusSuccessResponse) []*response.WithdrawResponseYearStatusSuccess {
	var WithdrawRecords []*response.WithdrawResponseYearStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapWithdrawResponseYearStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// mapResponseWithdrawMonthStatusFailed maps a single WithdrawMonthStatusFailedResponse gRPC message
// to an API response WithdrawResponseMonthStatusFailed object.
//
// Args:
//   - s: The WithdrawMonthStatusFailedResponse gRPC message that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseMonthStatusFailed containing the mapped data, including
//     fields like Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawResponseMapper) mapResponseWithdrawMonthStatusFailed(s *pb.WithdrawMonthStatusFailedResponse) *response.WithdrawResponseMonthStatusFailed {
	return &response.WithdrawResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponsesWithdrawMonthStatusFailed maps a slice of WithdrawMonthStatusFailedResponse gRPC messages
// to a slice of API response WithdrawResponseMonthStatusFailed objects.
//
// Args:
//   - Withdraws: A slice of pointers to pb.WithdrawMonthStatusFailedResponse messages that need to be converted.
//
// Returns:
//   - A slice of pointers to response.WithdrawResponseMonthStatusFailed objects containing the mapped data,
//     including fields like Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawResponseMapper) mapResponsesWithdrawMonthStatusFailed(Withdraws []*pb.WithdrawMonthStatusFailedResponse) []*response.WithdrawResponseMonthStatusFailed {
	var WithdrawRecords []*response.WithdrawResponseMonthStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapResponseWithdrawMonthStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// mapWithdrawResponseYearStatusFailed maps a WithdrawYearStatusFailedResponse gRPC message
// to an API response WithdrawResponseYearStatusFailed object.
//
// Args:
//   - s: The WithdrawYearStatusFailedResponse gRPC message that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseYearStatusFailed containing the mapped data, including
//     fields like Year, TotalFailed, and TotalAmount.
func (t *withdrawResponseMapper) mapWithdrawResponseYearStatusFailed(s *pb.WithdrawYearStatusFailedResponse) *response.WithdrawResponseYearStatusFailed {
	return &response.WithdrawResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapWithdrawResponsesYearStatusFailed maps a slice of gRPC WithdrawYearStatusFailedResponse messages
// to a slice of API WithdrawResponseYearStatusFailed objects.
//
// Args:
//   - Withdraws: A slice of pointers to pb.WithdrawYearStatusFailedResponse messages that need to be converted.
//
// Returns:
//   - A slice of pointers to response.WithdrawResponseYearStatusFailed objects containing the mapped data,
//     including fields like Year, TotalFailed, and TotalAmount.s
func (t *withdrawResponseMapper) mapWithdrawResponsesYearStatusFailed(Withdraws []*pb.WithdrawYearStatusFailedResponse) []*response.WithdrawResponseYearStatusFailed {
	var WithdrawRecords []*response.WithdrawResponseYearStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.mapWithdrawResponseYearStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// mapResponseWithdrawMonthlyAmount maps a single WithdrawMonthlyAmountResponse gRPC message
// to an API response WithdrawMonthlyAmountResponse object.
//
// Args:
//   - s: The WithdrawMonthlyAmountResponse gRPC message that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawMonthlyAmountResponse containing the mapped data, including
//     fields like Month and TotalAmount.
func (m *withdrawResponseMapper) mapResponseWithdrawMonthlyAmount(s *pb.WithdrawMonthlyAmountResponse) *response.WithdrawMonthlyAmountResponse {
	return &response.WithdrawMonthlyAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseWithdrawMonthlyAmounts maps a slice of gRPC WithdrawMonthlyAmountResponse messages
// to a slice of API response WithdrawMonthlyAmountResponse objects.
//
// Args:
//   - s: A slice of pointers to pb.WithdrawMonthlyAmountResponse messages that need to be converted.
//
// Returns:
//   - A slice of pointers to response.WithdrawMonthlyAmountResponse objects containing the mapped data,
//     including fields like Month and TotalAmount.
func (m *withdrawResponseMapper) mapResponseWithdrawMonthlyAmounts(s []*pb.WithdrawMonthlyAmountResponse) []*response.WithdrawMonthlyAmountResponse {
	var protoResponses []*response.WithdrawMonthlyAmountResponse
	for _, withdraw := range s {
		protoResponses = append(protoResponses, m.mapResponseWithdrawMonthlyAmount(withdraw))
	}
	return protoResponses
}

// mapResponseWithdrawYearlyAmount maps a single WithdrawYearlyAmountResponse gRPC message
// to an API response WithdrawYearlyAmountResponse object.
//
// Args:
//   - s: The WithdrawYearlyAmountResponse gRPC message that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawYearlyAmountResponse containing the mapped data, including
//     fields like Year and TotalAmount.
func (m *withdrawResponseMapper) mapResponseWithdrawYearlyAmount(s *pb.WithdrawYearlyAmountResponse) *response.WithdrawYearlyAmountResponse {
	return &response.WithdrawYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseWithdrawYearlyAmounts maps a slice of gRPC WithdrawYearlyAmountResponse messages
// to a slice of API response WithdrawYearlyAmountResponse objects.
//
// Args:
//   - s: A slice of pointers to pb.WithdrawYearlyAmountResponse messages that need to be converted.
//
// Returns:
//   - A slice of pointers to response.WithdrawYearlyAmountResponse objects containing the mapped data,
//     including fields like Year and TotalAmount.
func (m *withdrawResponseMapper) mapResponseWithdrawYearlyAmounts(s []*pb.WithdrawYearlyAmountResponse) []*response.WithdrawYearlyAmountResponse {
	var protoResponses []*response.WithdrawYearlyAmountResponse
	for _, withdraw := range s {
		protoResponses = append(protoResponses, m.mapResponseWithdrawYearlyAmount(withdraw))
	}
	return protoResponses
}
