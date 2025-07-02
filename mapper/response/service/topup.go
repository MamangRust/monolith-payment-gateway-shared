package responseservice

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// topupResponseMapper provides methods to map TopupRecord domain models to TopupResponse API-compatible response types.
type topupResponseMapper struct {
}

// NewTopupResponseMapper returns a pointer to a new topupResponseMapper instance.
// This struct provides methods to map TopupRecord domain models to TopupResponse API-compatible response types.
func NewTopupResponseMapper() *topupResponseMapper {
	return &topupResponseMapper{}
}

// ToTopupResponse maps a single TopupRecord domain model to a TopupResponse API-compatible response type.
// Args:
//   - topup: A pointer to a TopupRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TopupResponse containing the mapped data, including
//     ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *topupResponseMapper) ToTopupResponse(topup *record.TopupRecord) *response.TopupResponse {
	return &response.TopupResponse{
		ID:          topup.ID,
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: topup.TopupAmount,
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
	}
}

// ToTopupResponses maps a slice of TopupRecord to a slice of TopupResponse API-compatible response types.
//
// Args:
//   - topups: A slice of TopupRecord containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponse containing the mapped data, including
//     ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *topupResponseMapper) ToTopupResponses(topups []*record.TopupRecord) []*response.TopupResponse {
	var responses []*response.TopupResponse

	for _, response := range topups {
		responses = append(responses, s.ToTopupResponse(response))
	}

	return responses
}

// ToTopupResponseDeleteAt maps a single TopupRecord domain model to a TopupResponseDeleteAt API-compatible response type.
// It includes soft delete information by mapping the DeletedAt field.
//
// Args:
//   - topup: A pointer to a TopupRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a TopupResponseDeleteAt containing the mapped data, including
//     ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *topupResponseMapper) ToTopupResponseDeleteAt(topup *record.TopupRecord) *response.TopupResponseDeleteAt {
	return &response.TopupResponseDeleteAt{
		ID:          topup.ID,
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo,
		TopupAmount: topup.TopupAmount,
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime,
		CreatedAt:   topup.CreatedAt,
		UpdatedAt:   topup.UpdatedAt,
		DeletedAt:   topup.DeletedAt,
	}
}

// ToTopupResponsesDeleteAt maps a slice of TopupRecord domain models to a slice of TopupResponseDeleteAt API-compatible response types.
// It includes soft delete information by mapping the DeletedAt field.
//
// Args:
//   - topups: A slice of TopupRecord containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponseDeleteAt containing the mapped data, including
//     ID, CardNumber, TopupNo, TopupAmount, TopupMethod, TopupTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *topupResponseMapper) ToTopupResponsesDeleteAt(topups []*record.TopupRecord) []*response.TopupResponseDeleteAt {
	var responses []*response.TopupResponseDeleteAt

	for _, response := range topups {
		responses = append(responses, s.ToTopupResponseDeleteAt(response))
	}

	return responses
}

// ToTopupResponseMonthStatusSuccess converts a TopupRecordMonthStatusSuccess domain model
// into a TopupResponseMonthStatusSuccess API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupResponseMonthStatusSuccess containing the mapped data, including
//     Year, Month, TotalSuccess, and TotalAmount.
func (t *topupResponseMapper) ToTopupResponseMonthStatusSuccess(s *record.TopupRecordMonthStatusSuccess) *response.TopupResponseMonthStatusSuccess {
	return &response.TopupResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTopupResponsesMonthStatusSuccess maps a slice of TopupRecordMonthStatusSuccess domain models
// to a slice of TopupResponseMonthStatusSuccess API-compatible response types.
//
// Args:
//   - topups: A slice of TopupRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponseMonthStatusSuccess containing the mapped data, including
//     Year, Month, TotalSuccess, and TotalAmount.
func (t *topupResponseMapper) ToTopupResponsesMonthStatusSuccess(topups []*record.TopupRecordMonthStatusSuccess) []*response.TopupResponseMonthStatusSuccess {
	var topupRecords []*response.TopupResponseMonthStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupResponseMonthStatusSuccess(topup))
	}

	return topupRecords
}

// ToTopupResponseYearStatusSuccess converts a TopupRecordYearStatusSuccess domain model
// into a TopupResponseYearStatusSuccess API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupResponseYearStatusSuccess containing the mapped data, including
//     Year, TotalSuccess, and TotalAmount.
func (t *topupResponseMapper) ToTopupResponseYearStatusSuccess(s *record.TopupRecordYearStatusSuccess) *response.TopupResponseYearStatusSuccess {
	return &response.TopupResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTopupResponsesYearStatusSuccess maps a slice of TopupRecordYearStatusSuccess domain models
// to a slice of TopupResponseYearStatusSuccess API-compatible response types.
//
// Args:
//   - topups: A slice of TopupRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponseYearStatusSuccess containing the mapped data, including
//     Year, TotalSuccess, and TotalAmount.
func (t *topupResponseMapper) ToTopupResponsesYearStatusSuccess(topups []*record.TopupRecordYearStatusSuccess) []*response.TopupResponseYearStatusSuccess {
	var topupRecords []*response.TopupResponseYearStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupResponseYearStatusSuccess(topup))
	}

	return topupRecords
}

// ToTopupResponseMonthStatusFailed converts a TopupRecordMonthStatusFailed domain model
// into a TopupResponseMonthStatusFailed API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupResponseMonthStatusFailed containing the mapped data, including
//     Year, Month, TotalFailed, and TotalAmount.
func (t *topupResponseMapper) ToTopupResponseMonthStatusFailed(s *record.TopupRecordMonthStatusFailed) *response.TopupResponseMonthStatusFailed {
	return &response.TopupResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTopupResponsesMonthStatusFailed maps a slice of TopupRecordMonthStatusFailed domain models
// to a slice of TopupResponseMonthStatusFailed API-compatible response types.
//
// Args:
//   - topups: A slice of TopupRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponseMonthStatusFailed containing the mapped data, including
//     Year, Month, TotalFailed, and TotalAmount.
func (t *topupResponseMapper) ToTopupResponsesMonthStatusFailed(topups []*record.TopupRecordMonthStatusFailed) []*response.TopupResponseMonthStatusFailed {
	var topupRecords []*response.TopupResponseMonthStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupResponseMonthStatusFailed(topup))
	}

	return topupRecords
}

// ToTopupResponseYearStatusFailed maps a TopupRecordYearStatusFailed domain model
// to a TopupResponseYearStatusFailed API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupRecordYearStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *topupResponseMapper) ToTopupResponseYearStatusFailed(s *record.TopupRecordYearStatusFailed) *response.TopupResponseYearStatusFailed {
	return &response.TopupResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTopupResponsesYearStatusFailed maps a slice of TopupRecordYearStatusFailed domain models
// to a slice of TopupResponseYearStatusFailed API-compatible response types.
//
// Args:
//   - topups: A slice of TopupRecordYearStatusFailed containing the data to be mapped.
//
// Returns:
//   - A slice of TopupResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *topupResponseMapper) ToTopupResponsesYearStatusFailed(topups []*record.TopupRecordYearStatusFailed) []*response.TopupResponseYearStatusFailed {
	var topupRecords []*response.TopupResponseYearStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupResponseYearStatusFailed(topup))
	}

	return topupRecords
}

// ToTopupMonthlyMethodResponse converts a TopupMonthMethod domain model
// into a TopupMonthMethodResponse API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupMonthMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupMonthMethodResponse containing the mapped data, including
//     Month, TopupMethod, TotalTopups, and TotalAmount.
func (t *topupResponseMapper) ToTopupMonthlyMethodResponse(s *record.TopupMonthMethod) *response.TopupMonthMethodResponse {
	return &response.TopupMonthMethodResponse{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyMethodResponses maps a slice of TopupMonthMethod domain models
// to a slice of TopupMonthMethodResponse API-compatible response types.
//
// Args:
//   - s: A slice of TopupMonthMethod containing the data to be mapped.
//
// Returns:
//   - A slice of TopupMonthMethodResponse containing the mapped data, including
//     Month, TopupMethod, TotalTopups, and TotalAmount.
func (t *topupResponseMapper) ToTopupMonthlyMethodResponses(s []*record.TopupMonthMethod) []*response.TopupMonthMethodResponse {
	var topupResponses []*response.TopupMonthMethodResponse
	for _, topup := range s {
		topupResponses = append(topupResponses, t.ToTopupMonthlyMethodResponse(topup))
	}
	return topupResponses
}

// ToTopupYearlyMethodResponse converts a TopupYearlyMethod domain model
// into a TopupYearlyMethodResponse API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupYearlyMethod containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupYearlyMethodResponse containing the mapped data, including
//     Year, TopupMethod, TotalTopups, and TotalAmount.
func (t *topupResponseMapper) ToTopupYearlyMethodResponse(s *record.TopupYearlyMethod) *response.TopupYearlyMethodResponse {
	return &response.TopupYearlyMethodResponse{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyMethodResponses maps a slice of TopupYearlyMethod domain models
// to a slice of TopupYearlyMethodResponse API-compatible response types.
//
// Args:
//   - s: A slice of TopupYearlyMethod containing the data to be mapped.
//
// Returns:
//   - A slice of TopupYearlyMethodResponse containing the mapped data, including
//     Year, TopupMethod, TotalTopups, and TotalAmount.
func (t *topupResponseMapper) ToTopupYearlyMethodResponses(s []*record.TopupYearlyMethod) []*response.TopupYearlyMethodResponse {
	var topupResponses []*response.TopupYearlyMethodResponse
	for _, topup := range s {
		topupResponses = append(topupResponses, t.ToTopupYearlyMethodResponse(topup))
	}
	return topupResponses
}

// ToTopupMonthlyAmountResponse converts a TopupMonthAmount domain model
// into a TopupMonthAmountResponse API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupMonthAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupMonthAmountResponse containing the mapped data, including
//     Month, TotalAmount.
func (t *topupResponseMapper) ToTopupMonthlyAmountResponse(s *record.TopupMonthAmount) *response.TopupMonthAmountResponse {
	return &response.TopupMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyAmountResponses converts a slice of TopupMonthAmount domain models
// into a slice of TopupMonthAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of TopupMonthAmount containing the data to be mapped.
//
// Returns:
//   - A slice of TopupMonthAmountResponse containing the mapped data, including
//     Month and TotalAmount.
func (t *topupResponseMapper) ToTopupMonthlyAmountResponses(s []*record.TopupMonthAmount) []*response.TopupMonthAmountResponse {
	var topupResponses []*response.TopupMonthAmountResponse
	for _, topup := range s {
		topupResponses = append(topupResponses, t.ToTopupMonthlyAmountResponse(topup))
	}
	return topupResponses
}

// ToTopupYearlyAmountResponse converts a TopupYearlyAmount domain model
// into a TopupYearlyAmountResponse API-compatible response type.
//
// Args:
//   - s: A pointer to a TopupYearlyAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to TopupYearlyAmountResponse containing the mapped data, including
//     Year, TotalAmount.
func (t *topupResponseMapper) ToTopupYearlyAmountResponse(s *record.TopupYearlyAmount) *response.TopupYearlyAmountResponse {
	return &response.TopupYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyAmountResponses maps a slice of TopupYearlyAmount domain models
// to a slice of TopupYearlyAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of TopupYearlyAmount containing the data to be mapped.
//
// Returns:
//   - A slice of TopupYearlyAmountResponse containing the mapped data, including
//     Year, TotalAmount.
func (t *topupResponseMapper) ToTopupYearlyAmountResponses(s []*record.TopupYearlyAmount) []*response.TopupYearlyAmountResponse {
	var topupResponses []*response.TopupYearlyAmountResponse
	for _, topup := range s {
		topupResponses = append(topupResponses, t.ToTopupYearlyAmountResponse(topup))
	}
	return topupResponses
}
