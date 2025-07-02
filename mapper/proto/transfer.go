package protomapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// transferProtoMapper is responsible for mapping Transfer responses to their corresponding protobuf representations.
type transferProtoMapper struct{}

// NewTransferProtoMapper returns a new instance of transferProtoMapper.
func NewTransferProtoMapper() *transferProtoMapper {
	return &transferProtoMapper{}
}

// ToProtoResponseTransferMonthStatusSuccess maps a status, message, and a list of TransferResponseMonthStatusSuccess
// to an ApiResponseTransferMonthStatusSuccess proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponseMonthStatusSuccess to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferMonthStatusSuccess containing the mapped data.
func (m *transferProtoMapper) ToProtoResponseTransferMonthStatusSuccess(status string, message string, pbResponse []*response.TransferResponseMonthStatusSuccess) *pb.ApiResponseTransferMonthStatusSuccess {
	return &pb.ApiResponseTransferMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransferMonthStatusSuccess(pbResponse),
	}
}

// ToProtoResponseTransferYearStatusSuccess maps a status, message, and a list of TransferResponseYearStatusSuccess
// to an ApiResponseTransferYearStatusSuccess proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponseYearStatusSuccess to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferYearStatusSuccess containing the mapped data.
func (m *transferProtoMapper) ToProtoResponseTransferYearStatusSuccess(status string, message string, pbResponse []*response.TransferResponseYearStatusSuccess) *pb.ApiResponseTransferYearStatusSuccess {
	return &pb.ApiResponseTransferYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    m.mapTransferResponsesYearStatusSuccess(pbResponse),
	}
}

// ToProtoResponseTransferMonthStatusFailed maps a status, message, and a list of TransferResponseMonthStatusFailed
// to an ApiResponseTransferMonthStatusFailed proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponseMonthStatusFailed to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferMonthStatusFailed containing the mapped data.
func (m *transferProtoMapper) ToProtoResponseTransferMonthStatusFailed(status string, message string, pbResponse []*response.TransferResponseMonthStatusFailed) *pb.ApiResponseTransferMonthStatusFailed {
	return &pb.ApiResponseTransferMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransferMonthStatusFailed(pbResponse),
	}
}

// ToProtoResponseTransferYearStatusFailed maps a status, message, and a list of TransferResponseYearStatusFailed
// to an ApiResponseTransferYearStatusFailed proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponseYearStatusFailed to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferYearStatusFailed containing the mapped data.
func (m *transferProtoMapper) ToProtoResponseTransferYearStatusFailed(status string, message string, pbResponse []*response.TransferResponseYearStatusFailed) *pb.ApiResponseTransferYearStatusFailed {
	return &pb.ApiResponseTransferYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    m.mapTransferResponsesYearStatusFailed(pbResponse),
	}
}

// ToProtoResponseTransferMonthAmount maps a status, message, and a list of TransferMonthAmountResponse
// to an ApiResponseTransferMonthAmount proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferMonthAmountResponse to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferMonthAmount containing the mapped data.
func (m *transferProtoMapper) ToProtoResponseTransferMonthAmount(status string, message string, pbResponse []*response.TransferMonthAmountResponse) *pb.ApiResponseTransferMonthAmount {
	return &pb.ApiResponseTransferMonthAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransferMonthAmounts(pbResponse),
	}
}

// ToProtoResponseTransferYearAmount maps a status, message, and a list of TransferYearAmountResponse
// to an ApiResponseTransferYearAmount proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferYearAmountResponse to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransferYearAmount containing the mapped data.
func (m *transferProtoMapper) ToProtoResponseTransferYearAmount(status string, message string, pbResponse []*response.TransferYearAmountResponse) *pb.ApiResponseTransferYearAmount {
	return &pb.ApiResponseTransferYearAmount{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransferYearAmounts(pbResponse),
	}
}

// ToProtoResponseTransfer maps a status, message, and a single TransferResponse
// to an ApiResponseTransfer proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: The TransferResponse to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransfer containing the mapped data.
func (m *transferProtoMapper) ToProtoResponseTransfer(status string, message string, pbResponse *response.TransferResponse) *pb.ApiResponseTransfer {
	return &pb.ApiResponseTransfer{
		Status:  status,
		Message: message,
		Data:    m.mapResponseTransfer(pbResponse),
	}
}

// ToProtoResponseTransfers maps a status, message, and a list of TransferResponse
// to an ApiResponseTransfers proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponse to be converted.
//
// Returns:
//   - A pointer to an ApiResponseTransfers containing the mapped data.
func (m *transferProtoMapper) ToProtoResponseTransfers(status string, message string, pbResponse []*response.TransferResponse) *pb.ApiResponseTransfers {
	return &pb.ApiResponseTransfers{
		Status:  status,
		Message: message,
		Data:    m.mapResponsesTransfer(pbResponse),
	}
}

// ToProtoResponseTransferDelete maps a status, message, and no data
// to an ApiResponseTransferDelete proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//
// Returns:
//   - A pointer to an ApiResponseTransferDelete containing the mapped data.
func (m *transferProtoMapper) ToProtoResponseTransferDelete(status string, message string) *pb.ApiResponseTransferDelete {
	return &pb.ApiResponseTransferDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseTransferAll maps a status, message, and no data
// to an ApiResponseTransferAll proto message.
//
// Args:
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//
// Returns:
//   - A pointer to an ApiResponseTransferAll containing the mapped data.
func (m *transferProtoMapper) ToProtoResponseTransferAll(status string, message string) *pb.ApiResponseTransferAll {
	return &pb.ApiResponseTransferAll{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponsePaginationTransfer maps a pagination meta, status, message, and a list of TransferResponse
// to an ApiResponsePaginationTransfer proto message.
//
// Args:
//   - pagination: The pagination meta that needs to be converted.
//   - status: The status of the response.
//   - message: The accompanying message of the response.
//   - pbResponse: A slice of TransferResponse to be converted.
//
// Returns:
//   - A pointer to an ApiResponsePaginationTransfer containing the mapped data.
func (m *transferProtoMapper) ToProtoResponsePaginationTransfer(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransferResponse) *pb.ApiResponsePaginationTransfer {
	return &pb.ApiResponsePaginationTransfer{
		Status:     status,
		Message:    message,
		Data:       m.mapResponsesTransfer(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationTransferDeleteAt maps paginated soft-deleted transfer responses to a protobuf message.
//
// Args:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//   - pbResponse: A slice of TransferResponseDeleteAt objects to be included in the response.
//
// Returns:
//   - A pointer to an ApiResponsePaginationTransferDeleteAt containing the status, message, transfer data, and pagination data.
func (m *transferProtoMapper) ToProtoResponsePaginationTransferDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransferResponseDeleteAt) *pb.ApiResponsePaginationTransferDeleteAt {
	return &pb.ApiResponsePaginationTransferDeleteAt{
		Status:     status,
		Message:    message,
		Data:       m.mapResponsesTransferDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

// mapResponseTransfer maps a TransferResponse to a TransferResponse proto message.
//
// Args:
//   - transfer: The TransferResponse to be converted.
//
// Returns:
//   - A pointer to a TransferResponse containing the mapped data.
func (t *transferProtoMapper) mapResponseTransfer(transfer *response.TransferResponse) *pb.TransferResponse {
	return &pb.TransferResponse{
		Id:             int32(transfer.ID),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int32(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}
}

// mapResponsesTransfer maps a slice of TransferResponse to a slice of TransferResponse proto messages.
//
// Args:
//   - transfers: A slice of TransferResponse to be converted.
//
// Returns:
//   - A slice of pointers to TransferResponse containing the mapped data.
func (t *transferProtoMapper) mapResponsesTransfer(transfers []*response.TransferResponse) []*pb.TransferResponse {
	var responses []*pb.TransferResponse

	for _, response := range transfers {
		responses = append(responses, t.mapResponseTransfer(response))
	}

	return responses
}

// mapResponseTransferDeleteAt maps a TransferResponseDeleteAt to its protobuf representation.
//
// Args:
//   - transfer: The TransferResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to a TransferResponseDeleteAt protobuf message containing the mapped data.
func (t *transferProtoMapper) mapResponseTransferDeleteAt(transfer *response.TransferResponseDeleteAt) *pb.TransferResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if transfer.DeletedAt != nil {
		deletedAt = wrapperspb.String(*transfer.DeletedAt)
	}

	return &pb.TransferResponseDeleteAt{
		Id:             int32(transfer.ID),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int32(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
		DeletedAt:      deletedAt,
	}
}

// mapResponsesTransferDeleteAt maps a slice of TransferResponseDeleteAt to a slice of
// their protobuf representations.
//
// Args:
//   - transfers: A slice of TransferResponseDeleteAt to be converted.
//
// Returns:
//   - A slice of pointers to TransferResponseDeleteAt protobuf messages containing the mapped data.
func (t *transferProtoMapper) mapResponsesTransferDeleteAt(transfers []*response.TransferResponseDeleteAt) []*pb.TransferResponseDeleteAt {
	var responses []*pb.TransferResponseDeleteAt

	for _, response := range transfers {
		responses = append(responses, t.mapResponseTransferDeleteAt(response))
	}

	return responses
}

// mapResponseTransferMonthStatusSuccess maps a TransferResponseMonthStatusSuccess to its protobuf representation.
//
// Args:
//   - s: The TransferResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a TransferMonthStatusSuccessResponse protobuf message containing the mapped data.
func (t *transferProtoMapper) mapResponseTransferMonthStatusSuccess(s *response.TransferResponseMonthStatusSuccess) *pb.TransferMonthStatusSuccessResponse {
	return &pb.TransferMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapResponsesTransferMonthStatusSuccess maps a slice of TransferResponseMonthStatusSuccess
// to a slice of TransferMonthStatusSuccessResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferResponseMonthStatusSuccess
// to its corresponding protobuf representation using the mapResponseTransferMonthStatusSuccess method.
//
// Args:
//   - Transfers: A slice of TransferResponseMonthStatusSuccess to be converted.
//
// Returns:
//   - A slice of TransferMonthStatusSuccessResponse containing the mapped protobuf data.
func (t *transferProtoMapper) mapResponsesTransferMonthStatusSuccess(Transfers []*response.TransferResponseMonthStatusSuccess) []*pb.TransferMonthStatusSuccessResponse {
	var TransferRecords []*pb.TransferMonthStatusSuccessResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapResponseTransferMonthStatusSuccess(Transfer))
	}

	return TransferRecords
}

// mapTransferResponseYearStatusSuccess maps a TransferResponseYearStatusSuccess to its protobuf representation.
//
// Args:
//   - s: The TransferResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a TransferYearStatusSuccessResponse protobuf message containing the mapped data.
func (t *transferProtoMapper) mapTransferResponseYearStatusSuccess(s *response.TransferResponseYearStatusSuccess) *pb.TransferYearStatusSuccessResponse {
	return &pb.TransferYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapTransferResponsesYearStatusSuccess maps a slice of TransferResponseYearStatusSuccess
// to a slice of TransferYearStatusSuccessResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferResponseYearStatusSuccess
// to its corresponding protobuf representation using the mapTransferResponseYearStatusSuccess method.
//
// Args:
//   - Transfers: A slice of TransferResponseYearStatusSuccess to be converted.
//
// Returns:
//   - A slice of TransferYearStatusSuccessResponse containing the mapped protobuf data.
func (t *transferProtoMapper) mapTransferResponsesYearStatusSuccess(Transfers []*response.TransferResponseYearStatusSuccess) []*pb.TransferYearStatusSuccessResponse {
	var TransferRecords []*pb.TransferYearStatusSuccessResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapTransferResponseYearStatusSuccess(Transfer))
	}

	return TransferRecords
}

// mapResponseTransferMonthStatusFailed maps a TransferResponseMonthStatusFailed to its protobuf representation.
//
// Args:
//   - s: The TransferResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a TransferMonthStatusFailedResponse protobuf message containing the mapped data.
func (t *transferProtoMapper) mapResponseTransferMonthStatusFailed(s *response.TransferResponseMonthStatusFailed) *pb.TransferMonthStatusFailedResponse {
	return &pb.TransferMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponsesTransferMonthStatusFailed maps a slice of TransferResponseMonthStatusFailed
// to a slice of TransferMonthStatusFailedResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferResponseMonthStatusFailed
// to its corresponding protobuf representation using the mapResponseTransferMonthStatusFailed method.
//
// Args:
//   - Transfers: A slice of TransferResponseMonthStatusFailed to be converted.
//
// Returns:
//   - A slice of TransferMonthStatusFailedResponse containing the mapped protobuf data.
func (t *transferProtoMapper) mapResponsesTransferMonthStatusFailed(Transfers []*response.TransferResponseMonthStatusFailed) []*pb.TransferMonthStatusFailedResponse {
	var TransferRecords []*pb.TransferMonthStatusFailedResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapResponseTransferMonthStatusFailed(Transfer))
	}

	return TransferRecords
}

// mapTransferResponseYearStatusFailed maps a TransferResponseYearStatusFailed to its protobuf representation.
//
// Args:
//   - s: The TransferResponseYearStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a TransferYearStatusFailedResponse protobuf message containing the mapped data.
func (t *transferProtoMapper) mapTransferResponseYearStatusFailed(s *response.TransferResponseYearStatusFailed) *pb.TransferYearStatusFailedResponse {
	return &pb.TransferYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapTransferResponsesYearStatusFailed maps a slice of TransferResponseYearStatusFailed
// to a slice of TransferYearStatusFailedResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferResponseYearStatusFailed
// to its corresponding protobuf representation using the mapTransferResponseYearStatusFailed method.
//
// Args:
//   - Transfers: A slice of TransferResponseYearStatusFailed to be converted.
//
// Returns:
//   - A slice of TransferYearStatusFailedResponse containing the mapped protobuf data.
func (t *transferProtoMapper) mapTransferResponsesYearStatusFailed(Transfers []*response.TransferResponseYearStatusFailed) []*pb.TransferYearStatusFailedResponse {
	var TransferRecords []*pb.TransferYearStatusFailedResponse

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.mapTransferResponseYearStatusFailed(Transfer))
	}

	return TransferRecords
}

// mapResponseTransferMonthAmount maps a TransferMonthAmountResponse to its protobuf representation.
//
// Args:
//   - s: The TransferMonthAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to a TransferMonthAmountResponse protobuf message containing the mapped data.
func (m *transferProtoMapper) mapResponseTransferMonthAmount(s *response.TransferMonthAmountResponse) *pb.TransferMonthAmountResponse {
	return &pb.TransferMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTransferMonthAmounts maps a slice of TransferMonthAmountResponse
// to a slice of TransferMonthAmountResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferMonthAmountResponse
// to its corresponding protobuf representation using the mapResponseTransferMonthAmount method.
//
// Args:
//   - s: A slice of TransferMonthAmountResponse to be converted.
//
// Returns:
//   - A slice of TransferMonthAmountResponse containing the mapped protobuf data.
func (m *transferProtoMapper) mapResponseTransferMonthAmounts(s []*response.TransferMonthAmountResponse) []*pb.TransferMonthAmountResponse {
	var protoResponses []*pb.TransferMonthAmountResponse
	for _, transfer := range s {
		protoResponses = append(protoResponses, m.mapResponseTransferMonthAmount(transfer))
	}
	return protoResponses
}

// mapResponseTransferYearAmount maps a TransferYearAmountResponse to its protobuf representation.
//
// Args:
//   - s: The TransferYearAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to a TransferYearAmountResponse protobuf message containing the mapped data.
func (m *transferProtoMapper) mapResponseTransferYearAmount(s *response.TransferYearAmountResponse) *pb.TransferYearAmountResponse {
	return &pb.TransferYearAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTransferYearAmounts maps a slice of TransferYearAmountResponse
// to a slice of TransferYearAmountResponse protobuf messages.
//
// It iterates over the input slice, converting each TransferYearAmountResponse
// to its corresponding protobuf representation using the mapResponseTransferYearAmount method.
//
// Args:
//   - s: A slice of TransferYearAmountResponse to be converted.
//
// Returns:
//   - A slice of TransferYearAmountResponse containing the mapped protobuf data.
func (m *transferProtoMapper) mapResponseTransferYearAmounts(s []*response.TransferYearAmountResponse) []*pb.TransferYearAmountResponse {
	var protoResponses []*pb.TransferYearAmountResponse
	for _, transfer := range s {
		protoResponses = append(protoResponses, m.mapResponseTransferYearAmount(transfer))
	}
	return protoResponses
}
