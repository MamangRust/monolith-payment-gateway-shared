package apimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// topupResponseMapper provides methods to map gRPC top-up responses to HTTP-compatible API responses.
type topupResponseMapper struct {
}

// NewTopupResponseMapper returns a pointer to a new topupResponseMapper instance.
// This struct provides methods to map gRPC top-up responses to HTTP-compatible API responses.
func NewTopupResponseMapper() *topupResponseMapper {
	return &topupResponseMapper{}
}

// mapPaginationMeta maps a gRPC PaginationMeta to an HTTP-compatible API response PaginationMeta.
//
// Args:
//   - s: A pointer to a pb.PaginationMeta containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.PaginationMeta containing the mapped data, including current page, page size,
//     total records, and total pages.
func mapPaginationMeta(s *pb.PaginationMeta) *response.PaginationMeta {
	return &response.PaginationMeta{
		CurrentPage:  int(s.CurrentPage),
		PageSize:     int(s.PageSize),
		TotalRecords: int(s.TotalRecords),
		TotalPages:   int(s.TotalPages),
	}
}

// ToApiResponseTopup maps a single gRPC top-up response to an HTTP API response.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopup containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopup containing the mapped data, including status, message, and a single
//     mapped top-up response.
func (t *topupResponseMapper) ToApiResponseTopup(s *pb.ApiResponseTopup) *response.ApiResponseTopup {
	return &response.ApiResponseTopup{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopup(s.Data),
	}
}

// ToApiResponseTopupDeleteAt maps a single gRPC soft-deleted top-up response to an HTTP API response.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupDeleteAt containing the mapped data, including status, message, and a single
//     mapped soft-deleted top-up response.
func (t *topupResponseMapper) ToApiResponseTopupDeleteAt(s *pb.ApiResponseTopupDeleteAt) *response.ApiResponseTopupDeleteAt {
	return &response.ApiResponseTopupDeleteAt{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopupDeleteAt(s.Data),
	}
}

// ToApiResponseTopupAll maps a gRPC response containing all top-up records to an HTTP API response.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupAll containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupAll containing the mapped data, including status and message.
func (t *topupResponseMapper) ToApiResponseTopupAll(s *pb.ApiResponseTopupAll) *response.ApiResponseTopupAll {
	return &response.ApiResponseTopupAll{
		Status:  s.Status,
		Message: s.Message,
	}
}

// ToApiResponseTopupDelete maps a single gRPC top-up delete response to an HTTP API response.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupDelete containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupDelete containing the mapped data, including status and message.
func (t *topupResponseMapper) ToApiResponseTopupDelete(s *pb.ApiResponseTopupDelete) *response.ApiResponseTopupDelete {
	return &response.ApiResponseTopupDelete{
		Status:  s.Status,
		Message: s.Message,
	}
}

// ToApiResponsePaginationTopup maps a paginated gRPC top-up response to an HTTP API response. It constructs an
// ApiResponsePaginationTopup by copying the status and message fields, mapping the top-up data slice to a slice
// of TopupResponse, and including pagination metadata.
//
// Args:
//   - s: A pointer to a pb.ApiResponsePaginationTopup containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationTopup with mapped data and pagination info.
func (t *topupResponseMapper) ToApiResponsePaginationTopup(s *pb.ApiResponsePaginationTopup) *response.ApiResponsePaginationTopup {
	return &response.ApiResponsePaginationTopup{
		Status:     s.Status,
		Message:    s.Message,
		Data:       t.mapResponsesTopup(s.Data),
		Pagination: mapPaginationMeta(s.Pagination),
	}
}

// ToApiResponsePaginationTopupDeleteAt maps a paginated gRPC response of soft-deleted top-ups
// to an HTTP API response. It constructs an ApiResponsePaginationTopupDeleteAt by copying
// the status and message fields, mapping the top-up data slice to a slice of TopupDeleteAtResponse,
// and including pagination metadata.
//
// Args:
//   - s: A pointer to a pb.ApiResponsePaginationTopupDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationTopupDeleteAt with mapped data and pagination info.
func (t *topupResponseMapper) ToApiResponsePaginationTopupDeleteAt(s *pb.ApiResponsePaginationTopupDeleteAt) *response.ApiResponsePaginationTopupDeleteAt {
	return &response.ApiResponsePaginationTopupDeleteAt{
		Status:     s.Status,
		Message:    s.Message,
		Data:       t.mapResponsesTopupDeleteAt(s.Data),
		Pagination: mapPaginationMeta(s.Pagination),
	}
}

// ToApiResponseTopupMonthStatusSuccess maps a gRPC response containing a month's worth of successful top-up statistics
// to an HTTP API response. It constructs an ApiResponseTopupMonthStatusSuccess by copying the status and message fields,
// mapping the TopupMonthStatusSuccess data slice to a slice of TopupMonthStatusSuccessResponse, and assigning it to the
// response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupMonthStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupMonthStatusSuccess with mapped data.
func (t *topupResponseMapper) ToApiResponseTopupMonthStatusSuccess(s *pb.ApiResponseTopupMonthStatusSuccess) *response.ApiResponseTopupMonthStatusSuccess {
	return &response.ApiResponseTopupMonthStatusSuccess{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponsesTopupMonthStatusSuccess(s.Data),
	}
}

// ToApiResponseTopupYearStatusSuccess maps a gRPC response containing a year's worth of successful top-up statistics
// to an HTTP API response. It constructs an ApiResponseTopupYearStatusSuccess by copying the status and message fields,
// mapping the TopupYearStatusSuccess data slice to a slice of TopupYearStatusSuccessResponse, and assigning it to the
// response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupYearStatusSuccess containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupYearStatusSuccess with mapped data.
func (t *topupResponseMapper) ToApiResponseTopupYearStatusSuccess(s *pb.ApiResponseTopupYearStatusSuccess) *response.ApiResponseTopupYearStatusSuccess {
	return &response.ApiResponseTopupYearStatusSuccess{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapTopupResponsesYearStatusSuccess(s.Data),
	}
}

// ToApiResponseTopupMonthStatusFailed maps a gRPC response containing a month's worth of failed top-up statistics
// to an HTTP API response. It constructs an ApiResponseTopupMonthStatusFailed by copying the status and message fields,
// mapping the TopupMonthStatusFailed data slice to a slice of TopupMonthStatusFailedResponse, and assigning it to the
// response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupMonthStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupMonthStatusFailed with mapped data.
func (t *topupResponseMapper) ToApiResponseTopupMonthStatusFailed(s *pb.ApiResponseTopupMonthStatusFailed) *response.ApiResponseTopupMonthStatusFailed {
	return &response.ApiResponseTopupMonthStatusFailed{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponsesTopupMonthStatusFailed(s.Data),
	}
}

// ToApiResponseTopupYearStatusFailed maps a gRPC response containing a year's worth of failed top-up statistics
// to an HTTP API response. It constructs an ApiResponseTopupYearStatusFailed by copying the status and message fields,
// mapping the TopupYearStatusFailed data slice to a slice of TopupYearStatusFailedResponse, and assigning it to the
// response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupYearStatusFailed containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupYearStatusFailed with mapped data.
func (t *topupResponseMapper) ToApiResponseTopupYearStatusFailed(s *pb.ApiResponseTopupYearStatusFailed) *response.ApiResponseTopupYearStatusFailed {
	return &response.ApiResponseTopupYearStatusFailed{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapTopupResponsesYearStatusFailed(s.Data),
	}
}

// ToApiResponseTopupMonthMethod maps a gRPC response containing a month's worth of top-up statistics by payment method
// to an HTTP API response. It constructs an ApiResponseTopupMonthMethod by copying the status and message fields,
// mapping the TopupMonthMethod data slice to a slice of TopupMonthMethodResponse, and assigning it to the response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupMonthMethod containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupMonthMethod with mapped data.
func (t *topupResponseMapper) ToApiResponseTopupMonthMethod(s *pb.ApiResponseTopupMonthMethod) *response.ApiResponseTopupMonthMethod {
	return &response.ApiResponseTopupMonthMethod{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopupMonthlyMethods(s.Data),
	}
}

// ToApiResponseTopupYearMethod maps a gRPC response containing a year's worth of top-up statistics by payment method
// to an HTTP API response. It constructs an ApiResponseTopupYearMethod by copying the status and message fields,
// mapping the TopupYearMethod data slice to a slice of TopupYearMethodResponse, and assigning it to the response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupYearMethod containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupYearMethod with mapped data.
func (t *topupResponseMapper) ToApiResponseTopupYearMethod(s *pb.ApiResponseTopupYearMethod) *response.ApiResponseTopupYearMethod {
	return &response.ApiResponseTopupYearMethod{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopupYearlyMethods(s.Data),
	}
}

// ToApiResponseTopupMonthAmount maps a gRPC response containing a month's worth of top-up amounts
// to an HTTP API response. It constructs an ApiResponseTopupMonthAmount by copying the status and message fields,
// mapping the TopupMonthAmount data slice to a slice of TopupMonthAmountResponse, and assigning it to the response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupMonthAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupMonthAmount with mapped data.
func (t *topupResponseMapper) ToApiResponseTopupMonthAmount(s *pb.ApiResponseTopupMonthAmount) *response.ApiResponseTopupMonthAmount {
	return &response.ApiResponseTopupMonthAmount{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopupMonthlyAmounts(s.Data),
	}
}

// ToApiResponseTopupYearAmount maps a gRPC response containing a year's worth of top-up amounts
// to an HTTP API response. It constructs an ApiResponseTopupYearAmount by copying the status and message fields,
// mapping the TopupYearAmount data slice to a slice of TopupYearAmountResponse, and assigning it to the response's Data field.
//
// Args:
//   - s: A pointer to a pb.ApiResponseTopupYearAmount containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTopupYearAmount with mapped data.
func (t *topupResponseMapper) ToApiResponseTopupYearAmount(s *pb.ApiResponseTopupYearAmount) *response.ApiResponseTopupYearAmount {
	return &response.ApiResponseTopupYearAmount{
		Status:  s.Status,
		Message: s.Message,
		Data:    t.mapResponseTopupYearlyAmounts(s.Data),
	}
}

// mapResponseTopup maps a single gRPC top-up response to an HTTP API response.
// It constructs a TopupResponse by copying the fields from the gRPC response.
//
// Args:
//   - topup: A pointer to a pb.TopupResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TopupResponse with mapped data.
func (t *topupResponseMapper) mapResponseTopup(topup *pb.TopupResponse) *response.TopupResponse {
	return &response.TopupResponse{
		ID:          int(topup.Id),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}

// mapResponsesTopup maps a slice of gRPC top-up responses to a slice of HTTP API responses.
// It iterates over the slice of gRPC responses and maps each one to an HTTP API response
// using mapResponseTopup, returning the slice of mapped responses.
//
// Args:
//   - topups: A slice of pointers to pb.TopupResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponse with the mapped data.
func (t *topupResponseMapper) mapResponsesTopup(topups []*pb.TopupResponse) []*response.TopupResponse {
	var responses []*response.TopupResponse

	for _, topup := range topups {
		responses = append(responses, t.mapResponseTopup(topup))
	}

	return responses
}

// mapResponseTopupDeleteAt maps a gRPC soft-deleted top-up response to an HTTP API response format.
// It constructs a TopupResponseDeleteAt by copying the fields from the gRPC response, converting
// the ID and TopupAmount fields to integers, and handling the potentially nil DeletedAt field.
//
// Args:
//   - topup: A pointer to a pb.TopupResponseDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TopupResponseDeleteAt with the mapped data, including a non-nil DeletedAt field if present.
func (t *topupResponseMapper) mapResponseTopupDeleteAt(topup *pb.TopupResponseDeleteAt) *response.TopupResponseDeleteAt {
	var deletedAt string
	if topup.DeletedAt != nil {
		deletedAt = topup.DeletedAt.Value
	}

	return &response.TopupResponseDeleteAt{
		ID:          int(topup.Id),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
		DeletedAt:   &deletedAt,
	}
}

// mapResponsesTopupDeleteAt maps a slice of gRPC soft-deleted top-up responses to a slice
// of HTTP API responses. It iterates over the gRPC response slice, mapping each individual
// response using mapResponseTopupDeleteAt, and returns a slice of mapped TopupResponseDeleteAt.
//
// Args:
//   - topups: A slice of pointers to pb.TopupResponseDeleteAt containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponseDeleteAt with the mapped data.
func (t *topupResponseMapper) mapResponsesTopupDeleteAt(topups []*pb.TopupResponseDeleteAt) []*response.TopupResponseDeleteAt {
	var responses []*response.TopupResponseDeleteAt

	for _, topup := range topups {
		responses = append(responses, t.mapResponseTopupDeleteAt(topup))
	}

	return responses
}

// mapResponseTopupMonthStatusSuccess maps a gRPC response containing a month's worth of successful top-up statistics
// to an HTTP API response format. It constructs a TopupResponseMonthStatusSuccess by copying the year, month,
// total success count, and total amount fields from the gRPC response to the API response.
func (t *topupResponseMapper) mapResponseTopupMonthStatusSuccess(s *pb.TopupMonthStatusSuccessResponse) *response.TopupResponseMonthStatusSuccess {
	return &response.TopupResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapResponsesTopupMonthStatusSuccess maps a slice of gRPC responses containing a month's worth of successful
// top-up statistics to a slice of HTTP API responses. It iterates over the gRPC response slice, mapping
// each individual response using mapResponseTopupMonthStatusSuccess, and returns a slice of mapped
// TopupResponseMonthStatusSuccess.
//
// Args:
//   - topups: A slice of pointers to pb.TopupMonthStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponseMonthStatusSuccess with the mapped data.
func (t *topupResponseMapper) mapResponsesTopupMonthStatusSuccess(topups []*pb.TopupMonthStatusSuccessResponse) []*response.TopupResponseMonthStatusSuccess {
	var topupRecords []*response.TopupResponseMonthStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapResponseTopupMonthStatusSuccess(topup))
	}

	return topupRecords
}

// mapTopupResponseYearStatusSuccess maps a gRPC response containing a year's worth of successful top-up statistics
// to an HTTP API response format. It constructs a TopupResponseYearStatusSuccess by copying the year, total success count,
// and total amount fields from the gRPC response to the API response.
func (t *topupResponseMapper) mapTopupResponseYearStatusSuccess(s *pb.TopupYearStatusSuccessResponse) *response.TopupResponseYearStatusSuccess {
	return &response.TopupResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// mapTopupResponsesYearStatusSuccess maps a slice of gRPC responses containing a year's worth of successful
// top-up statistics to a slice of HTTP API responses. It iterates over the gRPC response slice, mapping
// each individual response using mapTopupResponseYearStatusSuccess, and returns a slice of mapped
// TopupResponseYearStatusSuccess.
//
// Args:
//   - topups: A slice of pointers to pb.TopupYearStatusSuccessResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponseYearStatusSuccess with the mapped data.
func (t *topupResponseMapper) mapTopupResponsesYearStatusSuccess(topups []*pb.TopupYearStatusSuccessResponse) []*response.TopupResponseYearStatusSuccess {
	var topupRecords []*response.TopupResponseYearStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapTopupResponseYearStatusSuccess(topup))
	}

	return topupRecords
}

// mapResponseTopupMonthStatusFailed maps a gRPC response containing a month's worth of failed top-up statistics
// to an HTTP API response format. It constructs a TopupResponseMonthStatusFailed by copying the year, month,
// total failed count, and total amount fields from the gRPC response to the API response.
func (t *topupResponseMapper) mapResponseTopupMonthStatusFailed(s *pb.TopupMonthStatusFailedResponse) *response.TopupResponseMonthStatusFailed {
	return &response.TopupResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponsesTopupMonthStatusFailed maps a slice of gRPC responses containing a month's worth of failed
// top-up statistics to a slice of HTTP API responses. It iterates over the gRPC response slice, mapping
// each individual response using mapResponseTopupMonthStatusFailed, and returns a slice of mapped
// TopupResponseMonthStatusFailed.
//
// Args:
//   - topups: A slice of pointers to pb.TopupMonthStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponseMonthStatusFailed with the mapped data.
func (t *topupResponseMapper) mapResponsesTopupMonthStatusFailed(topups []*pb.TopupMonthStatusFailedResponse) []*response.TopupResponseMonthStatusFailed {
	var topupRecords []*response.TopupResponseMonthStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapResponseTopupMonthStatusFailed(topup))
	}

	return topupRecords
}

// mapTopupResponseYearStatusFailed maps a gRPC response containing a year's worth
// of failed top-up statistics to an HTTP API response format. It constructs a
// TopupResponseYearStatusFailed by copying the year, total failed count, and
// total amount fields from the gRPC response to the API response.
func (t *topupResponseMapper) mapTopupResponseYearStatusFailed(s *pb.TopupYearStatusFailedResponse) *response.TopupResponseYearStatusFailed {
	return &response.TopupResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapTopupResponsesYearStatusFailed maps a slice of gRPC responses containing a year's worth
// of failed top-up statistics to a slice of HTTP API responses. It iterates over the gRPC
// response slice, mapping each individual response using mapTopupResponseYearStatusFailed,
// and returns a slice of mapped TopupResponseYearStatusFailed.
//
// Args:
//   - topups: A slice of pointers to pb.TopupYearStatusFailedResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupResponseYearStatusFailed with the mapped data.
func (t *topupResponseMapper) mapTopupResponsesYearStatusFailed(topups []*pb.TopupYearStatusFailedResponse) []*response.TopupResponseYearStatusFailed {
	var topupRecords []*response.TopupResponseYearStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.mapTopupResponseYearStatusFailed(topup))
	}

	return topupRecords
}

// mapResponseTopupMonthlyMethod maps a gRPC response containing a month's worth of top-up statistics
// by payment method to an HTTP API response format. It constructs a TopupMonthMethodResponse by
// copying the month, top-up method, total top-ups, and total amount fields from the gRPC response
// to the API response.
func (t *topupResponseMapper) mapResponseTopupMonthlyMethod(s *pb.TopupMonthMethodResponse) *response.TopupMonthMethodResponse {
	return &response.TopupMonthMethodResponse{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTopupMonthlyMethods maps a slice of gRPC responses containing monthly top-up statistics by payment method
// to a slice of HTTP API response formats. It iterates over the gRPC response slice, mapping each individual response
// using mapResponseTopupMonthlyMethod, and returns a slice of mapped TopupMonthMethodResponse.
//
// Args:
//   - s: A slice of pointers to pb.TopupMonthMethodResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupMonthMethodResponse with the mapped data.
func (t *topupResponseMapper) mapResponseTopupMonthlyMethods(s []*pb.TopupMonthMethodResponse) []*response.TopupMonthMethodResponse {
	var topupProtos []*response.TopupMonthMethodResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupMonthlyMethod(topup))
	}
	return topupProtos
}

// mapResponseTopupYearlyMethod maps a gRPC response containing a year's worth of top-up statistics
// by payment method to an HTTP API response format. It constructs a TopupYearlyMethodResponse by
// copying the year, top-up method, total top-ups, and total amount fields from the gRPC response
// to the API response.
//
// Args:
//   - s: A pointer to a pb.TopupYearlyMethodResponse containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.TopupYearlyMethodResponse with the mapped data.
func (t *topupResponseMapper) mapResponseTopupYearlyMethod(s *pb.TopupYearlyMethodResponse) *response.TopupYearlyMethodResponse {
	return &response.TopupYearlyMethodResponse{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTopupYearlyMethods maps a slice of gRPC responses containing yearly top-up statistics
// by payment method to a slice of HTTP API response formats. It iterates over the gRPC response slice,
// mapping each individual response using mapResponseTopupYearlyMethod, and returns a slice of mapped
// TopupYearlyMethodResponse.
//
// Args:
//   - s: A slice of pointers to pb.TopupYearlyMethodResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupYearlyMethodResponse with the mapped data.
func (t *topupResponseMapper) mapResponseTopupYearlyMethods(s []*pb.TopupYearlyMethodResponse) []*response.TopupYearlyMethodResponse {
	var topupProtos []*response.TopupYearlyMethodResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupYearlyMethod(topup))
	}
	return topupProtos
}

// mapResponseTopupMonthlyAmount maps a gRPC response containing a month's worth of top-up
// amount statistics to an HTTP API response format. It constructs a TopupMonthAmountResponse
// by copying the month and total amount fields from the gRPC response to the API response.
func (t *topupResponseMapper) mapResponseTopupMonthlyAmount(s *pb.TopupMonthAmountResponse) *response.TopupMonthAmountResponse {
	return &response.TopupMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTopupMonthlyAmounts maps a slice of gRPC responses containing monthly top-up amount statistics
// to a slice of HTTP API response formats. It iterates over the gRPC response slice, mapping each individual response
// using mapResponseTopupMonthlyAmount, and returns a slice of mapped TopupMonthAmountResponse.
//
// Args:
//   - s: A slice of pointers to pb.TopupMonthAmountResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupMonthAmountResponse with the mapped data.
func (t *topupResponseMapper) mapResponseTopupMonthlyAmounts(s []*pb.TopupMonthAmountResponse) []*response.TopupMonthAmountResponse {
	var topupProtos []*response.TopupMonthAmountResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupMonthlyAmount(topup))
	}
	return topupProtos
}

// mapResponseTopupYearlyAmount maps a gRPC response containing a year's worth of top-up amount statistics
// to an HTTP API response format. It constructs a TopupYearlyAmountResponse by copying the year and
// total amount fields from the gRPC response to the API response.
func (t *topupResponseMapper) mapResponseTopupYearlyAmount(s *pb.TopupYearlyAmountResponse) *response.TopupYearlyAmountResponse {
	return &response.TopupYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// mapResponseTopupYearlyAmounts maps a slice of gRPC responses containing yearly top-up amount statistics
// to a slice of HTTP API response formats. It iterates over the gRPC response slice, mapping each individual response
// using mapResponseTopupYearlyAmount, and returns a slice of mapped TopupYearlyAmountResponse.
//
// Args:
//   - s: A slice of pointers to pb.TopupYearlyAmountResponse containing the gRPC response data.
//
// Returns:
//   - A slice of pointers to response.TopupYearlyAmountResponse with the mapped data.
func (t *topupResponseMapper) mapResponseTopupYearlyAmounts(s []*pb.TopupYearlyAmountResponse) []*response.TopupYearlyAmountResponse {
	var topupProtos []*response.TopupYearlyAmountResponse
	for _, topup := range s {
		topupProtos = append(topupProtos, t.mapResponseTopupYearlyAmount(topup))
	}
	return topupProtos
}
