package apimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transferResponseMapper struct{}

// NewTransferResponseMapper creates a new TransferResponseMapper.
func NewTransferResponseMapper() *transferResponseMapper {
	return &transferResponseMapper{}
}

// ToApiResponseTransferMonthStatusSuccess converts a gRPC response containing a month's worth
// of successful transfer statistics into an HTTP API response. It constructs an
// ApiResponseTransferMonthStatusSuccess by copying the status and message fields, mapping the
// TransferMonthStatusSuccess data slice to a slice of TransferResponseMonthStatusSuccess, and
// assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferMonthStatusSuccess containing the gRPC
//     response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferMonthStatusSuccess with mapped data.
func (m *transferResponseMapper) ToApiResponseTransferMonthStatusSuccess(pbResponse *pb.ApiResponseTransferMonthStatusSuccess) *response.ApiResponseTransferMonthStatusSuccess {
	return &response.ApiResponseTransferMonthStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesTransferMonthStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseTransferYearStatusSuccess maps a gRPC response containing a year's worth of successful transfer
// statistics to an HTTP API response. It constructs an ApiResponseTransferYearStatusSuccess by copying the
// status and message fields, mapping the TransferYearStatusSuccess data slice to a slice of
// TransferResponseYearStatusSuccess, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferYearStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferYearStatusSuccess with mapped data.
func (m *transferResponseMapper) ToApiResponseTransferYearStatusSuccess(pbResponse *pb.ApiResponseTransferYearStatusSuccess) *response.ApiResponseTransferYearStatusSuccess {
	return &response.ApiResponseTransferYearStatusSuccess{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapTransferResponsesYearStatusSuccess(pbResponse.Data),
	}
}

// ToApiResponseTransferMonthStatusFailed maps a gRPC response containing a month's worth of failed transfer
// statistics into an HTTP API response. It constructs an ApiResponseTransferMonthStatusFailed by copying the
// status and message fields, mapping the TransferMonthStatusFailed data slice to a slice of
// TransferResponseMonthStatusFailed, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferMonthStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferMonthStatusFailed with mapped data.
func (m *transferResponseMapper) ToApiResponseTransferMonthStatusFailed(pbResponse *pb.ApiResponseTransferMonthStatusFailed) *response.ApiResponseTransferMonthStatusFailed {
	return &response.ApiResponseTransferMonthStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesTransferMonthStatusFailed(pbResponse.Data),
	}
}

// ToApiResponseTransferYearStatusFailed maps a gRPC response containing a year's worth of failed transfer
// statistics into an HTTP API response. It constructs an ApiResponseTransferYearStatusFailed by copying the
// status and message fields, mapping the TransferYearStatusFailed data slice to a slice of
// TransferResponseYearStatusFailed, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferYearStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferYearStatusFailed with mapped data.
func (m *transferResponseMapper) ToApiResponseTransferYearStatusFailed(pbResponse *pb.ApiResponseTransferYearStatusFailed) *response.ApiResponseTransferYearStatusFailed {
	return &response.ApiResponseTransferYearStatusFailed{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapTransferResponsesYearStatusFailed(pbResponse.Data),
	}
}

// ToApiResponseTransferMonthAmount maps a gRPC response containing a month's worth of transfer amounts
// into an HTTP API response. It constructs an ApiResponseTransferMonthAmount by copying the status and message fields,
// mapping the TransferMonthAmount data slice to a slice of TransferMonthAmountResponse, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferMonthAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferMonthAmount with mapped data.
func (m *transferResponseMapper) ToApiResponseTransferMonthAmount(pbResponse *pb.ApiResponseTransferMonthAmount) *response.ApiResponseTransferMonthAmount {
	return &response.ApiResponseTransferMonthAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransferMonthAmounts(pbResponse.Data),
	}
}

// ToApiResponseTransferYearAmount maps a gRPC response containing a year's worth of transfer amounts
// into an HTTP API response. It constructs an ApiResponseTransferYearAmount by copying the status and message fields,
// mapping the TransferYearAmount data slice to a slice of TransferYearAmountResponse, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferYearAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferYearAmount with mapped data.
func (m *transferResponseMapper) ToApiResponseTransferYearAmount(pbResponse *pb.ApiResponseTransferYearAmount) *response.ApiResponseTransferYearAmount {
	return &response.ApiResponseTransferYearAmount{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransferYearAmounts(pbResponse.Data),
	}
}

// ToApiResponseTransfer maps a single gRPC transfer response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransfer containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransfer containing the mapped data, including status, message, and a single mapped transfer response.
func (m *transferResponseMapper) ToApiResponseTransfer(pbResponse *pb.ApiResponseTransfer) *response.ApiResponseTransfer {
	return &response.ApiResponseTransfer{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransfer(pbResponse.Data),
	}
}

// ToApiResponseTransfers maps a gRPC response containing multiple transfer records into an HTTP API response.
// It constructs an ApiResponseTransfers by copying the status and message fields, mapping the Transfer data slice to a slice of
// TransferResponse, and assigning it to the response's Data field.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransfers containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransfers with mapped data.
func (m *transferResponseMapper) ToApiResponseTransfers(pbResponse *pb.ApiResponseTransfers) *response.ApiResponseTransfers {
	return &response.ApiResponseTransfers{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponsesTransfer(pbResponse.Data),
	}
}

// ToApiResponseTransferDelete maps a gRPC transfer delete response to an HTTP API response.
// It constructs an ApiResponseTransferDelete by copying the status and message fields from
// the gRPC response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferDelete containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferDelete with the mapped status and message.
func (m *transferResponseMapper) ToApiResponseTransferDelete(pbResponse *pb.ApiResponseTransferDelete) *response.ApiResponseTransferDelete {
	return &response.ApiResponseTransferDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseTransferAll maps a gRPC response containing all transfer records into an HTTP API response.
// It constructs an ApiResponseTransferAll by copying the status and message fields from the gRPC response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransferAll containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransferAll with the mapped status and message.
func (m *transferResponseMapper) ToApiResponseTransferAll(pbResponse *pb.ApiResponseTransferAll) *response.ApiResponseTransferAll {
	return &response.ApiResponseTransferAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponsePaginationTransfer maps a pagination meta, status, message, and a list of TransferResponse
// to a response.ApiResponsePaginationTransfer proto message.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationTransfer containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationTransfer containing the mapped data.
func (m *transferResponseMapper) ToApiResponsePaginationTransfer(pbResponse *pb.ApiResponsePaginationTransfer) *response.ApiResponsePaginationTransfer {
	return &response.ApiResponsePaginationTransfer{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.mapResponsesTransfer(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

// ToApiResponsePaginationTransferDeleteAt maps a pagination meta, status, message, and a list of TransferResponseDeleteAt
// to a response.ApiResponsePaginationTransferDeleteAt proto message.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationTransferDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationTransferDeleteAt containing the mapped data.
func (m *transferResponseMapper) ToApiResponsePaginationTransferDeleteAt(pbResponse *pb.ApiResponsePaginationTransferDeleteAt) *response.ApiResponsePaginationTransferDeleteAt {
	return &response.ApiResponsePaginationTransferDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       m.mapResponsesTransferDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

// mapResponseTransfer maps a TransferResponse from a protobuf message to a response.TransferResponse.
//
// Args:
//   - transfer: A pointer to a pb.TransferResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponse containing the mapped data.
func (t *transferResponseMapper) mapResponseTransfer(transfer *pb.TransferResponse) *response.TransferResponse {
	return &response.TransferResponse{
		ID:             int(transfer.Id),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}
}

// mapResponsesTransfer maps a slice of TransferResponse from protobuf messages to a slice of response.TransferResponse.
//
// Args:
//   - transfers: A slice of pointers to pb.TransferResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponse containing the mapped data.
func (t *transferResponseMapper) mapResponsesTransfer(transfers []*pb.TransferResponse) []*response.TransferResponse {
	var responses []*response.TransferResponse

	for _, response := range transfers {
		responses = append(responses, t.mapResponseTransfer(response))
	}

	return responses
}

// mapResponseTransferDeleteAt maps a TransferResponseDeleteAt from a protobuf message to a response.TransferResponseDeleteAt.
//
// Args:
//   - transfer: A pointer to a pb.TransferResponseDeleteAt containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseDeleteAt containing the mapped data.
func (t *transferResponseMapper) mapResponseTransferDeleteAt(transfer *pb.TransferResponseDeleteAt) *response.TransferResponseDeleteAt {
	var deletedAt string
	if transfer.DeletedAt != nil {
		deletedAt = transfer.DeletedAt.Value
	}

	return &response.TransferResponseDeleteAt{
		ID:             int(transfer.Id),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
		DeletedAt:      &deletedAt,
	}
}

// mapResponsesTransferDeleteAt maps a slice of TransferResponseDeleteAt from protobuf messages to a slice of response.TransferResponseDeleteAt.
//
// It iterates over the input slice, converting each TransferResponseDeleteAt to its corresponding response.TransferResponseDeleteAt
// using the mapResponseTransferDeleteAt method.
//
// Args:
//   - transfers: A slice of pointers to pb.TransferResponseDeleteAt containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponseDeleteAt containing the mapped data.
func (t *transferResponseMapper) mapResponsesTransferDeleteAt(transfers []*pb.TransferResponseDeleteAt) []*response.TransferResponseDeleteAt {
	var responses []*response.TransferResponseDeleteAt

	for _, response := range transfers {
		responses = append(responses, t.mapResponseTransferDeleteAt(response))
	}

	return responses
}

// mapResponseTransferMonthStatusSuccess maps a TransferMonthStatusSuccessResponse protobuf message to a
// response.TransferResponseMonthStatusSuccess.
//
// Args:
//   - s: A pointer to a pb.TransferMonthStatusSuccessResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transferResponseMapper) mapResponseTransferMonthStatusSuccess(s *pb.TransferMonthStatusSuccessResponse) *response.TransferResponseMonthStatusSuccess {
	return &response.TransferResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapResponsesTransferMonthStatusSuccess maps a slice of TransferMonthStatusSuccessResponse
// protobuf messages to a slice of TransferResponseMonthStatusSuccess.
//
// It iterates over the input slice, converting each TransferMonthStatusSuccessResponse
// to its corresponding response representation using the mapResponseTransferMonthStatusSuccess method.
//
// Args:
//   - Transfers: A slice of pointers to pb.TransferMonthStatusSuccessResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponseMonthStatusSuccess containing the mapped data.
func (t *transferResponseMapper) mapResponsesTransferMonthStatusSuccess(Transfers []*pb.TransferMonthStatusSuccessResponse) []*response.TransferResponseMonthStatusSuccess {
	var TransferRecords []*response.TransferResponseMonthStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapResponseTransferMonthStatusSuccess(Transfer))
	}

	return TransferRecords
}

// mapTransferResponseYearStatusSuccess maps a TransferYearStatusSuccessResponse protobuf message to a
// response.TransferResponseYearStatusSuccess.
//
// Args:
//   - s: A pointer to a pb.TransferYearStatusSuccessResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transferResponseMapper) mapTransferResponseYearStatusSuccess(s *pb.TransferYearStatusSuccessResponse) *response.TransferResponseYearStatusSuccess {
	return &response.TransferResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapTransferResponsesYearStatusSuccess maps a slice of TransferYearStatusSuccessResponse
// protobuf messages to a slice of TransferResponseYearStatusSuccess.
//
// It iterates over the input slice, converting each TransferYearStatusSuccessResponse
// to its corresponding response representation using the mapTransferResponseYearStatusSuccess method.
//
// Args:
//   - Transfers: A slice of pointers to pb.TransferYearStatusSuccessResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponseYearStatusSuccess containing the mapped data.
func (t *transferResponseMapper) mapTransferResponsesYearStatusSuccess(Transfers []*pb.TransferYearStatusSuccessResponse) []*response.TransferResponseYearStatusSuccess {
	var TransferRecords []*response.TransferResponseYearStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapTransferResponseYearStatusSuccess(Transfer))
	}

	return TransferRecords
}

// mapResponseTransferMonthStatusFailed maps a TransferMonthStatusFailedResponse protobuf message to a
// response.TransferResponseMonthStatusFailed.
//
// Args:
//   - s: A pointer to a pb.TransferMonthStatusFailedResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transferResponseMapper) mapResponseTransferMonthStatusFailed(s *pb.TransferMonthStatusFailedResponse) *response.TransferResponseMonthStatusFailed {
	return &response.TransferResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponsesTransferMonthStatusFailed maps a slice of TransferMonthStatusFailedResponse
// protobuf messages to a slice of TransferResponseMonthStatusFailed.
//
// It iterates over the input slice, converting each TransferMonthStatusFailedResponse
// to its corresponding response representation using the mapResponseTransferMonthStatusFailed method.
//
// Args:
//   - Transfers: A slice of pointers to pb.TransferMonthStatusFailedResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponseMonthStatusFailed containing the mapped data.
func (t *transferResponseMapper) mapResponsesTransferMonthStatusFailed(Transfers []*pb.TransferMonthStatusFailedResponse) []*response.TransferResponseMonthStatusFailed {
	var TransferRecords []*response.TransferResponseMonthStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapResponseTransferMonthStatusFailed(Transfer))
	}

	return TransferRecords
}

// mapTransferResponseYearStatusFailed maps a TransferYearStatusFailedResponse protobuf message to a
// response.TransferResponseYearStatusFailed.
//
// Args:
//   - s: A pointer to a pb.TransferYearStatusFailedResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponseYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transferResponseMapper) mapTransferResponseYearStatusFailed(s *pb.TransferYearStatusFailedResponse) *response.TransferResponseYearStatusFailed {
	return &response.TransferResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapTransferResponsesYearStatusFailed maps a slice of TransferYearStatusFailedResponse
// protobuf messages to a slice of TransferResponseYearStatusFailed.
//
// It iterates over the input slice, converting each TransferYearStatusFailedResponse
// to its corresponding response representation using the mapTransferResponseYearStatusFailed method.
//
// Args:
//   - Transfers: A slice of pointers to pb.TransferYearStatusFailedResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferResponseYearStatusFailed containing the mapped data.
func (t *transferResponseMapper) mapTransferResponsesYearStatusFailed(Transfers []*pb.TransferYearStatusFailedResponse) []*response.TransferResponseYearStatusFailed {
	var TransferRecords []*response.TransferResponseYearStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapTransferResponseYearStatusFailed(Transfer))
	}

	return TransferRecords
}

// mapResponseTransferMonthAmount maps a TransferMonthAmountResponse protobuf message to a
// response.TransferMonthAmountResponse.
//
// Args:
//   - s: A pointer to a pb.TransferMonthAmountResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferMonthAmountResponse containing the mapped data, including Month,
//     and TotalAmount.
func (m *transferResponseMapper) mapResponseTransferMonthAmount(s *pb.TransferMonthAmountResponse) *response.TransferMonthAmountResponse {
	return &response.TransferMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTransferMonthAmounts maps a slice of TransferMonthAmountResponse
// protobuf messages to a slice of TransferMonthAmountResponse.
//
// It iterates over the input slice, converting each TransferMonthAmountResponse
// to its corresponding response representation using the mapResponseTransferMonthAmount method.
//
// Args:
//   - s: A slice of pointers to pb.TransferMonthAmountResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferMonthAmountResponse containing the mapped data.
func (m *transferResponseMapper) mapResponseTransferMonthAmounts(s []*pb.TransferMonthAmountResponse) []*response.TransferMonthAmountResponse {
	var protoResponses []*response.TransferMonthAmountResponse
	for _, transfer := range s {
		protoResponses = append(protoResponses, m.mapResponseTransferMonthAmount(transfer))
	}
	return protoResponses
}

// mapResponseTransferYearAmount maps a TransferYearAmountResponse protobuf message to a
// response.TransferYearAmountResponse.
//
// Args:
//   - s: A pointer to a pb.TransferYearAmountResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferYearAmountResponse containing the mapped data, including Year
//     and TotalAmount.
func (m *transferResponseMapper) mapResponseTransferYearAmount(s *pb.TransferYearAmountResponse) *response.TransferYearAmountResponse {
	return &response.TransferYearAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTransferYearAmounts maps a slice of TransferYearAmountResponse
// protobuf messages to a slice of TransferYearAmountResponse.
//
// It iterates over the input slice, converting each TransferYearAmountResponse
// to its corresponding response representation using the mapResponseTransferYearAmount method.
//
// Args:
//   - s: A slice of pointers to pb.TransferYearAmountResponse containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to response.TransferYearAmountResponse containing the mapped data.
func (m *transferResponseMapper) mapResponseTransferYearAmounts(s []*pb.TransferYearAmountResponse) []*response.TransferYearAmountResponse {
	var protoResponses []*response.TransferYearAmountResponse
	for _, transfer := range s {
		protoResponses = append(protoResponses, m.mapResponseTransferYearAmount(transfer))
	}
	return protoResponses
}
