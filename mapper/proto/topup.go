package protomapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// topupProtoMapper is responsible for mapping Topup responses to their corresponding protobuf representations.
type topupProtoMapper struct {
}

// NewTopupProtoMapper returns a new instance of topupProtoMapper.
func NewTopupProtoMapper() *topupProtoMapper {
	return &topupProtoMapper{}
}

// ToProtoResponsePaginationTopup maps paginated top-up records to a protobuf response.
//
// Args:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponsePaginationTopup containing the status, message, top-up data, and pagination data.
func (t *topupProtoMapper) ToProtoResponsePaginationTopup(pagination *pb.PaginationMeta, status string, message string, s []*response.TopupResponse) *pb.ApiResponsePaginationTopup {
	return &pb.ApiResponsePaginationTopup{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesTopup(s),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationTopupDeleteAt maps paginated soft-deleted top-up records to a protobuf response.
//
// Args:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponseDeleteAt that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponsePaginationTopupDeleteAt containing the status, message, top-up data, and pagination data.
func (t *topupProtoMapper) ToProtoResponsePaginationTopupDeleteAt(pagination *pb.PaginationMeta, status string, message string, s []*response.TopupResponseDeleteAt) *pb.ApiResponsePaginationTopupDeleteAt {
	return &pb.ApiResponsePaginationTopupDeleteAt{
		Status:     status,
		Message:    message,
		Data:       t.mapResponsesTopupDeleteAt(s),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponseTopupMonthStatusSuccess maps monthly successful top-ups to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupMonthStatusSuccess containing the status, message, and top-up data.
func (t *topupProtoMapper) ToProtoResponseTopupMonthStatusSuccess(status string, message string, s []*response.TopupResponseMonthStatusSuccess) *pb.ApiResponseTopupMonthStatusSuccess {
	return &pb.ApiResponseTopupMonthStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupMonthStatusSuccess(s),
	}
}

// ToProtoResponseTopupYearStatusSuccess maps yearly successful top-ups to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupYearStatusSuccess containing the status, message, and top-up data.
func (t *topupProtoMapper) ToProtoResponseTopupYearStatusSuccess(status string, message string, s []*response.TopupResponseYearStatusSuccess) *pb.ApiResponseTopupYearStatusSuccess {
	return &pb.ApiResponseTopupYearStatusSuccess{
		Status:  status,
		Message: message,
		Data:    t.mapTopupResponsesYearStatusSuccess(s),
	}
}

// ToProtoResponseTopupMonthStatusFailed maps monthly failed top-ups to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupMonthStatusFailed containing the status, message, and top-up data.
func (t *topupProtoMapper) ToProtoResponseTopupMonthStatusFailed(status string, message string, s []*response.TopupResponseMonthStatusFailed) *pb.ApiResponseTopupMonthStatusFailed {
	return &pb.ApiResponseTopupMonthStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapResponsesTopupMonthStatusFailed(s),
	}
}

// ToProtoResponseTopupYearStatusFailed maps yearly failed top-ups to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupResponseYearStatusFailed that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupYearStatusFailed containing the status, message, and top-up data.
func (t *topupProtoMapper) ToProtoResponseTopupYearStatusFailed(status string, message string, s []*response.TopupResponseYearStatusFailed) *pb.ApiResponseTopupYearStatusFailed {
	return &pb.ApiResponseTopupYearStatusFailed{
		Status:  status,
		Message: message,
		Data:    t.mapTopupResponsesYearStatusFailed(s),
	}
}

// ToProtoResponseTopupMonthMethod maps monthly top-up methods to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupMonthMethodResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupMonthMethod containing the status, message, and top-up data.
func (t *topupProtoMapper) ToProtoResponseTopupMonthMethod(status string, message string, s []*response.TopupMonthMethodResponse) *pb.ApiResponseTopupMonthMethod {
	return &pb.ApiResponseTopupMonthMethod{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupMonthlyMethods(s),
	}
}

// ToProtoResponseTopupYearMethod maps yearly top-up methods to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupYearlyMethodResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopupYearMethod containing the status, message, and top-up data.
func (t *topupProtoMapper) ToProtoResponseTopupYearMethod(status string, message string, s []*response.TopupYearlyMethodResponse) *pb.ApiResponseTopupYearMethod {
	return &pb.ApiResponseTopupYearMethod{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupYearlyMethods(s),
	}
}

// ToProtoResponseTopupMonthAmount maps monthly top-up amounts to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupMonthAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to ApiResponseTopupMonthAmount containing the status, message, and mapped top-up amount data.
func (t *topupProtoMapper) ToProtoResponseTopupMonthAmount(status string, message string, s []*response.TopupMonthAmountResponse) *pb.ApiResponseTopupMonthAmount {
	return &pb.ApiResponseTopupMonthAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupMonthlyAmounts(s),
	}
}

// ToProtoResponseTopupYearAmount maps yearly top-up amounts to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A slice of TopupYearlyAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to ApiResponseTopupYearAmount containing the status, message, and mapped top-up amount data.
func (t *topupProtoMapper) ToProtoResponseTopupYearAmount(status string, message string, s []*response.TopupYearlyAmountResponse) *pb.ApiResponseTopupYearAmount {
	return &pb.ApiResponseTopupYearAmount{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupYearlyAmounts(s),
	}
}

// ToProtoResponseTopup maps a single top-up record to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A pointer to TopupResponse that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponseTopup containing the status, message, and top-up data.
func (t *topupProtoMapper) ToProtoResponseTopup(status string, message string, s *response.TopupResponse) *pb.ApiResponseTopup {
	return &pb.ApiResponseTopup{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopup(s),
	}
}

// ToProtoResponseTopupDeletAt maps a soft-deleted top-up record to a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//   - s: A pointer to TopupResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to ApiResponseTopupDeleteAt containing the status, message, and mapped top-up data.
func (t *topupProtoMapper) ToProtoResponseTopupDeletAt(status string, message string, s *response.TopupResponseDeleteAt) *pb.ApiResponseTopupDeleteAt {
	return &pb.ApiResponseTopupDeleteAt{
		Status:  status,
		Message: message,
		Data:    t.mapResponseTopupDeleteAt(s),
	}
}

// ToProtoResponseTopupDelete returns a response indicating a top-up record has been deleted.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//
// Returns:
//
//	A pointer to ApiResponseTopupDelete containing the status and message.
func (t topupProtoMapper) ToProtoResponseTopupDelete(status string, message string) *pb.ApiResponseTopupDelete {
	return &pb.ApiResponseTopupDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseTopupAll returns all top-up records in a protobuf response.
//
// Args:
//   - status: The status of the API response.
//   - message: A descriptive message of the API response.
//
// Returns:
//   - A pointer to ApiResponseTopupAll containing the status and message.
func (t topupProtoMapper) ToProtoResponseTopupAll(status string, message string) *pb.ApiResponseTopupAll {
	return &pb.ApiResponseTopupAll{
		Status:  status,
		Message: message,
	}
}

// mapResponseTopup maps a single TopupResponse to its corresponding protobuf response.
//
// Args:
//   - topup: A pointer to TopupResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupResponse containing the mapped top-up data.
func (t *topupProtoMapper) mapResponseTopup(topup *response.TopupResponse) *pb.TopupResponse {
	return &pb.TopupResponse{
		Id:          int32(topup.ID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}

// mapResponsesTopup maps a slice of TopupResponse to a slice of TopupResponse.
//
// Args:
//   - topups: A slice of TopupResponse that needs to be converted.
//
// Returns:
//   - A slice of TopupResponse containing the mapped top-up data.
func (t *topupProtoMapper) mapResponsesTopup(topups []*response.TopupResponse) []*pb.TopupResponse {
	var responses []*pb.TopupResponse

	for _, response := range topups {
		responses = append(responses, t.mapResponseTopup(response))
	}

	return responses
}

// mapResponseTopupDeleteAt maps a single TopupResponseDeleteAt to its corresponding protobuf response.
//
// Args:
//   - topup: A pointer to TopupResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A pointer to TopupResponseDeleteAt containing the mapped top-up data.
func (t *topupProtoMapper) mapResponseTopupDeleteAt(topup *response.TopupResponseDeleteAt) *pb.TopupResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if topup.DeletedAt != nil {
		deletedAt = wrapperspb.String(*topup.DeletedAt)
	}

	return &pb.TopupResponseDeleteAt{
		Id:          int32(topup.ID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int32(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
		DeletedAt:   deletedAt,
	}
}

// mapResponsesTopupDeleteAt maps a slice of TopupResponseDeleteAt to a slice of TopupResponseDeleteAt.
//
// Args:
//   - topups: A slice of TopupResponseDeleteAt that needs to be converted.
//
// Returns:
//   - A slice of TopupResponseDeleteAt containing the mapped top-up data.
func (t *topupProtoMapper) mapResponsesTopupDeleteAt(topups []*response.TopupResponseDeleteAt) []*pb.TopupResponseDeleteAt {
	var responses []*pb.TopupResponseDeleteAt

	for _, response := range topups {
		responses = append(responses, t.mapResponseTopupDeleteAt(response))
	}

	return responses
}

// mapResponseTopupMonthStatusSuccess maps a TopupResponseMonthStatusSuccess to a TopupMonthStatusSuccessResponse.
//
// Args:
//   - s: A pointer to TopupResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to TopupMonthStatusSuccessResponse containing the mapped top-up data.
func (t *topupProtoMapper) mapResponseTopupMonthStatusSuccess(s *response.TopupResponseMonthStatusSuccess) *pb.TopupMonthStatusSuccessResponse {
	return &pb.TopupMonthStatusSuccessResponse{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapResponsesTopupMonthStatusSuccess maps a slice of TopupResponseMonthStatusSuccess to a slice of
// TopupMonthStatusSuccessResponse protobuf messages.
//
// Args:
//   - topups: A slice of TopupResponseMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A slice of TopupMonthStatusSuccessResponse containing the mapped monthly top-up success data.
func (t *topupProtoMapper) mapResponsesTopupMonthStatusSuccess(topups []*response.TopupResponseMonthStatusSuccess) []*pb.TopupMonthStatusSuccessResponse {
	var topupRecords []*pb.TopupMonthStatusSuccessResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapResponseTopupMonthStatusSuccess(topup))
	}

	return topupRecords
}

// mapTopupResponseYearStatusSuccess maps a TopupResponseYearStatusSuccess to a TopupYearStatusSuccessResponse.
//
// Args:
//   - s: A pointer to TopupResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to TopupYearStatusSuccessResponse containing the mapped yearly top-up success data.
func (t *topupProtoMapper) mapTopupResponseYearStatusSuccess(s *response.TopupResponseYearStatusSuccess) *pb.TopupYearStatusSuccessResponse {
	return &pb.TopupYearStatusSuccessResponse{
		Year:         s.Year,
		TotalSuccess: int32(s.TotalSuccess),
		TotalAmount:  int32(s.TotalAmount),
	}
}

// mapTopupResponsesYearStatusSuccess maps a slice of TopupResponseYearStatusSuccess to a slice of
// TopupYearStatusSuccessResponse protobuf messages.
//
// Args:
//   - topups: A slice of TopupResponseYearStatusSuccess that needs to be converted.
//
// Returns:
//   - A slice of TopupYearStatusSuccessResponse containing the mapped yearly top-up success data.
func (t *topupProtoMapper) mapTopupResponsesYearStatusSuccess(topups []*response.TopupResponseYearStatusSuccess) []*pb.TopupYearStatusSuccessResponse {
	var topupRecords []*pb.TopupYearStatusSuccessResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapTopupResponseYearStatusSuccess(topup))
	}

	return topupRecords
}

// mapResponseTopupMonthStatusFailed maps a TopupResponseMonthStatusFailed to a TopupMonthStatusFailedResponse protobuf message.
//
// Args:
//   - s: A pointer to TopupResponseMonthStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to TopupMonthStatusFailedResponse containing the mapped monthly top-up failure data.
func (t *topupProtoMapper) mapResponseTopupMonthStatusFailed(s *response.TopupResponseMonthStatusFailed) *pb.TopupMonthStatusFailedResponse {
	return &pb.TopupMonthStatusFailedResponse{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponsesTopupMonthStatusFailed converts a slice of TopupResponseMonthStatusFailed
// into a slice of TopupMonthStatusFailedResponse protobuf messages.
//
// Args:
//   - topups: A slice of TopupResponseMonthStatusFailed to be converted.
//
// Returns:
//   - A slice of TopupMonthStatusFailedResponse containing the mapped data
//     for each monthly top-up failure entry.
func (t *topupProtoMapper) mapResponsesTopupMonthStatusFailed(topups []*response.TopupResponseMonthStatusFailed) []*pb.TopupMonthStatusFailedResponse {
	var topupRecords []*pb.TopupMonthStatusFailedResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapResponseTopupMonthStatusFailed(topup))
	}

	return topupRecords
}

// mapTopupResponseYearStatusFailed maps a TopupResponseYearStatusFailed to a TopupYearStatusFailedResponse.
//
// Args:
//   - s: A pointer to TopupResponseYearStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to TopupYearStatusFailedResponse containing the mapped yearly top-up failure data.
func (t *topupProtoMapper) mapTopupResponseYearStatusFailed(s *response.TopupResponseYearStatusFailed) *pb.TopupYearStatusFailedResponse {
	return &pb.TopupYearStatusFailedResponse{
		Year:        s.Year,
		TotalFailed: int32(s.TotalFailed),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapTopupResponsesYearStatusFailed maps a slice of TopupResponseYearStatusFailed
// into a slice of TopupYearStatusFailedResponse protobuf messages.
//
// Args:
//   - topups: A slice of TopupResponseYearStatusFailed to be converted.
//
// Returns:
//   - A slice of TopupYearStatusFailedResponse containing the mapped yearly top-up failure data
//     for each yearly top-up failure entry.
func (t *topupProtoMapper) mapTopupResponsesYearStatusFailed(topups []*response.TopupResponseYearStatusFailed) []*pb.TopupYearStatusFailedResponse {
	var topupRecords []*pb.TopupYearStatusFailedResponse

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapTopupResponseYearStatusFailed(topup))
	}

	return topupRecords
}

// mapResponseTopupMonthlyMethod maps a TopupMonthMethodResponse to a TopupMonthMethodResponse protobuf message.
//
// Args:
//   - s: A pointer to TopupMonthMethodResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupMonthMethodResponse containing the mapped monthly top-up method data.
func (t *topupProtoMapper) mapResponseTopupMonthlyMethod(s *response.TopupMonthMethodResponse) *pb.TopupMonthMethodResponse {
	return &pb.TopupMonthMethodResponse{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int32(s.TotalTopups),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTopupMonthlyMethods maps a slice of TopupMonthMethodResponse to a slice of
// TopupMonthMethodResponse protobuf messages.
//
// Args:
//   - s: A slice of TopupMonthMethodResponse that needs to be converted.
//
// Returns:
//   - A slice of TopupMonthMethodResponse containing the mapped monthly top-up method data.
func (t *topupProtoMapper) mapResponseTopupMonthlyMethods(s []*response.TopupMonthMethodResponse) []*pb.TopupMonthMethodResponse {
	var topupProtos []*pb.TopupMonthMethodResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupMonthlyMethod(topup))
	}
	return topupProtos
}

// mapResponseTopupYearlyMethod maps a TopupYearlyMethodResponse to a TopupYearlyMethodResponse protobuf message.
//
// Args:
//   - s: A pointer to TopupYearlyMethodResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupYearlyMethodResponse containing the mapped yearly top-up method data.
func (t *topupProtoMapper) mapResponseTopupYearlyMethod(s *response.TopupYearlyMethodResponse) *pb.TopupYearlyMethodResponse {
	return &pb.TopupYearlyMethodResponse{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int32(s.TotalTopups),
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTopupYearlyMethods maps a slice of TopupYearlyMethodResponse to a slice of
// TopupYearlyMethodResponse protobuf messages.
//
// Args:
//   - s: A slice of TopupYearlyMethodResponse that needs to be converted.
//
// Returns:
//   - A slice of TopupYearlyMethodResponse containing the mapped yearly top-up method data.
func (t *topupProtoMapper) mapResponseTopupYearlyMethods(s []*response.TopupYearlyMethodResponse) []*pb.TopupYearlyMethodResponse {
	var topupProtos []*pb.TopupYearlyMethodResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupYearlyMethod(topup))
	}
	return topupProtos
}

// mapResponseTopupMonthlyAmount maps a TopupMonthAmountResponse to a TopupMonthAmountResponse protobuf message.
//
// Args:
//   - s: A pointer to TopupMonthAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupMonthAmountResponse containing the mapped monthly top-up amount data.
func (t *topupProtoMapper) mapResponseTopupMonthlyAmount(s *response.TopupMonthAmountResponse) *pb.TopupMonthAmountResponse {
	return &pb.TopupMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTopupMonthlyAmounts maps a slice of TopupMonthAmountResponse to a slice of
// TopupMonthAmountResponse protobuf messages.
//
// Args:
//   - s: A slice of TopupMonthAmountResponse that needs to be converted.
//
// Returns:
//   - A slice of TopupMonthAmountResponse containing the mapped monthly top-up amount data.
func (t *topupProtoMapper) mapResponseTopupMonthlyAmounts(s []*response.TopupMonthAmountResponse) []*pb.TopupMonthAmountResponse {
	var topupProtos []*pb.TopupMonthAmountResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupMonthlyAmount(topup))
	}
	return topupProtos
}

// mapResponseTopupYearlyAmount maps a TopupYearlyAmountResponse to a TopupYearlyAmountResponse protobuf message.
//
// Args:
//   - s: A pointer to TopupYearlyAmountResponse that needs to be converted.
//
// Returns:
//   - A pointer to TopupYearlyAmountResponse containing the mapped yearly top-up amount data.
func (t *topupProtoMapper) mapResponseTopupYearlyAmount(s *response.TopupYearlyAmountResponse) *pb.TopupYearlyAmountResponse {
	return &pb.TopupYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int32(s.TotalAmount),
	}
}

// mapResponseTopupYearlyAmounts maps a slice of TopupYearlyAmountResponse to a slice of
// TopupYearlyAmountResponse protobuf messages.
//
// Args:
//   - s: A slice of TopupYearlyAmountResponse that needs to be converted.
//
// Returns:
//   - A slice of TopupYearlyAmountResponse containing the mapped yearly top-up amount data.
func (t *topupProtoMapper) mapResponseTopupYearlyAmounts(s []*response.TopupYearlyAmountResponse) []*pb.TopupYearlyAmountResponse {
	var topupProtos []*pb.TopupYearlyAmountResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupYearlyAmount(topup))
	}
	return topupProtos
}
