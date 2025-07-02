package responseservice

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// withdrawResponseMapper provides methods to map WithdrawRecord domain models to WithdrawResponse API-compatible response types.
type withdrawResponseMapper struct {
}

// NewWithdrawResponseMapper creates and returns a new instance of withdrawResponseMapper.
// This instance provides methods to map WithdrawRecord domain models to
// API-compatible WithdrawResponse objects.
func NewWithdrawResponseMapper() *withdrawResponseMapper {
	return &withdrawResponseMapper{}
}

// ToWithdrawResponse converts a single WithdrawRecord into a WithdrawResponse.
//
// Args:
//   - withdraw: The WithdrawRecord that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponse containing the mapped data, including ID, WithdrawNo,
//     CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.
func (s *withdrawResponseMapper) ToWithdrawResponse(withdraw *record.WithdrawRecord) *response.WithdrawResponse {
	return &response.WithdrawResponse{
		ID:             withdraw.ID,
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: withdraw.WithdrawAmount,
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
	}
}

// ToWithdrawsResponse converts multiple WithdrawRecords into a slice of WithdrawResponse.
//
// Args:
//   - withdraws: The WithdrawRecords that needs to be converted.
//
// Returns:
//   - A slice of WithdrawResponse containing the mapped data, including ID, WithdrawNo,
//     CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, and UpdatedAt.
func (s *withdrawResponseMapper) ToWithdrawsResponse(withdraws []*record.WithdrawRecord) []*response.WithdrawResponse {
	var withdrawResponses []*response.WithdrawResponse
	for _, withdraw := range withdraws {
		withdrawResponses = append(withdrawResponses, s.ToWithdrawResponse(withdraw))
	}
	return withdrawResponses
}

// ToWithdrawResponseDeleteAt converts a single WithdrawRecord into a WithdrawResponseDeleteAt.
//
// Args:
//   - withdraw: The WithdrawRecord that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseDeleteAt containing the mapped data, including ID, WithdrawNo,
//     CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawResponseMapper) ToWithdrawResponseDeleteAt(withdraw *record.WithdrawRecord) *response.WithdrawResponseDeleteAt {
	return &response.WithdrawResponseDeleteAt{
		ID:             withdraw.ID,
		WithdrawNo:     withdraw.WithdrawNo,
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: withdraw.WithdrawAmount,
		WithdrawTime:   withdraw.WithdrawTime,
		CreatedAt:      withdraw.CreatedAt,
		UpdatedAt:      withdraw.UpdatedAt,
		DeletedAt:      withdraw.DeletedAt,
	}
}

func (s *withdrawResponseMapper) ToWithdrawsResponseDeleteAt(withdraws []*record.WithdrawRecord) []*response.WithdrawResponseDeleteAt {
	var withdrawResponses []*response.WithdrawResponseDeleteAt

	for _, withdraw := range withdraws {
		withdrawResponses = append(withdrawResponses, s.ToWithdrawResponseDeleteAt(withdraw))
	}
	return withdrawResponses
}

// ToWithdrawResponseMonthStatusSuccess converts a single WithdrawRecordMonthStatusSuccess into a WithdrawResponseMonthStatusSuccess.
//
// Args:
//   - s: The WithdrawRecordMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseMonthStatusSuccess containing the mapped data, including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawResponseMapper) ToWithdrawResponseMonthStatusSuccess(s *record.WithdrawRecordMonthStatusSuccess) *response.WithdrawResponseMonthStatusSuccess {
	return &response.WithdrawResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToWithdrawResponsesMonthStatusSuccess converts multiple WithdrawRecordMonthStatusSuccess into a slice of WithdrawResponseMonthStatusSuccess.
//
// Args:
//   - Withdraws: The WithdrawRecordMonthStatusSuccess that needs to be converted.
//
// Returns:
//   - A slice of WithdrawResponseMonthStatusSuccess containing the mapped data, including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawResponseMapper) ToWithdrawResponsesMonthStatusSuccess(Withdraws []*record.WithdrawRecordMonthStatusSuccess) []*response.WithdrawResponseMonthStatusSuccess {
	var WithdrawRecords []*response.WithdrawResponseMonthStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawResponseMonthStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawResponseYearStatusSuccess converts a single WithdrawRecordYearStatusSuccess into a WithdrawResponseYearStatusSuccess.
//
// Args:
//   - s: The WithdrawRecordYearStatusSuccess that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseYearStatusSuccess containing the mapped data, including Year, TotalSuccess, and TotalAmount.
func (t *withdrawResponseMapper) ToWithdrawResponseYearStatusSuccess(s *record.WithdrawRecordYearStatusSuccess) *response.WithdrawResponseYearStatusSuccess {
	return &response.WithdrawResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToWithdrawResponsesYearStatusSuccess converts multiple yearly success-status WithdrawRecords
// into a slice of WithdrawResponseYearStatusSuccess.
//
// Args:
//   - Withdraws: A slice of pointers to WithdrawRecordYearStatusSuccess that need to be converted.
//
// Returns:
//   - A slice of pointers to WithdrawResponseYearStatusSuccess, each containing the mapped data
//     including Year, TotalSuccess, and TotalAmount.
func (t *withdrawResponseMapper) ToWithdrawResponsesYearStatusSuccess(Withdraws []*record.WithdrawRecordYearStatusSuccess) []*response.WithdrawResponseYearStatusSuccess {
	var WithdrawRecords []*response.WithdrawResponseYearStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawResponseYearStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawResponseMonthStatusFailed converts a single WithdrawRecordMonthStatusFailed into a WithdrawResponseMonthStatusFailed.
//
// Args:
//   - s: The WithdrawRecordMonthStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseMonthStatusFailed containing the mapped data, including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawResponseMapper) ToWithdrawResponseMonthStatusFailed(s *record.WithdrawRecordMonthStatusFailed) *response.WithdrawResponseMonthStatusFailed {
	return &response.WithdrawResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToWithdrawResponsesMonthStatusFailed converts multiple WithdrawRecordMonthStatusFailed into a slice of WithdrawResponseMonthStatusFailed.
//
// Args:
//   - Withdraws: A slice of pointers to WithdrawRecordMonthStatusFailed that need to be converted.
//
// Returns:
//   - A slice of pointers to WithdrawResponseMonthStatusFailed, each containing the mapped data
//     including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawResponseMapper) ToWithdrawResponsesMonthStatusFailed(Withdraws []*record.WithdrawRecordMonthStatusFailed) []*response.WithdrawResponseMonthStatusFailed {
	var WithdrawRecords []*response.WithdrawResponseMonthStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawResponseMonthStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawResponseYearStatusFailed converts a single WithdrawRecordYearStatusFailed into
// a WithdrawResponseYearStatusFailed.
//
// Args:
//   - s: The WithdrawRecordYearStatusFailed that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *withdrawResponseMapper) ToWithdrawResponseYearStatusFailed(s *record.WithdrawRecordYearStatusFailed) *response.WithdrawResponseYearStatusFailed {
	return &response.WithdrawResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToWithdrawResponsesYearStatusFailed converts multiple yearly failed-status WithdrawRecords
// into a slice of WithdrawResponseYearStatusFailed.
//
// Args:
//   - Withdraws: A slice of pointers to WithdrawRecordYearStatusFailed that need to be converted.
//
// Returns:
//   - A slice of pointers to WithdrawResponseYearStatusFailed, each containing the mapped data
//     including Year, TotalFailed, and TotalAmount.
func (t *withdrawResponseMapper) ToWithdrawResponsesYearStatusFailed(Withdraws []*record.WithdrawRecordYearStatusFailed) []*response.WithdrawResponseYearStatusFailed {
	var WithdrawRecords []*response.WithdrawResponseYearStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawResponseYearStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawAmountMonthlyResponse converts a single WithdrawMonthlyAmount record
// into a WithdrawMonthlyAmountResponse.
//
// Args:
//   - s: The WithdrawMonthlyAmount that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawMonthlyAmountResponse containing the mapped data,
//     including Month and TotalAmount.
func (w *withdrawResponseMapper) ToWithdrawAmountMonthlyResponse(s *record.WithdrawMonthlyAmount) *response.WithdrawMonthlyAmountResponse {
	return &response.WithdrawMonthlyAmountResponse{
		Month:       s.Month,
		TotalAmount: s.TotalAmount,
	}
}

// ToWithdrawsAmountMonthlyResponses converts a slice of WithdrawMonthlyAmount
// records into a slice of WithdrawMonthlyAmountResponse objects.
//
// Args:
//   - s: A slice of pointers to WithdrawMonthlyAmount records that need to be converted.
//
// Returns:
//   - A slice of pointers to WithdrawMonthlyAmountResponse, each containing the mapped data,
//     including Month and TotalAmount.
func (w *withdrawResponseMapper) ToWithdrawsAmountMonthlyResponses(s []*record.WithdrawMonthlyAmount) []*response.WithdrawMonthlyAmountResponse {
	var withdrawResponses []*response.WithdrawMonthlyAmountResponse
	for _, withdraw := range s {
		withdrawResponses = append(withdrawResponses, w.ToWithdrawAmountMonthlyResponse(withdraw))
	}
	return withdrawResponses
}

// ToWithdrawAmountYearlyResponse converts a WithdrawYearlyAmount record
// into a WithdrawYearlyAmountResponse.
//
// Args:
//   - s: The WithdrawYearlyAmount that needs to be converted.
//
// Returns:
//   - A pointer to a WithdrawYearlyAmountResponse containing the mapped data, including
//     Year and TotalAmount.
func (w *withdrawResponseMapper) ToWithdrawAmountYearlyResponse(s *record.WithdrawYearlyAmount) *response.WithdrawYearlyAmountResponse {
	return &response.WithdrawYearlyAmountResponse{
		Year:        s.Year,
		TotalAmount: s.TotalAmount,
	}
}

// ToWithdrawsAmountYearlyResponses converts a slice of WithdrawYearlyAmount records
// into a slice of WithdrawYearlyAmountResponse.
//
// Args:
//   - s: A slice of pointers to WithdrawYearlyAmount records that need to be converted.
//
// Returns:
//   - A slice of pointers to WithdrawYearlyAmountResponse, each containing the mapped data,
//     including Year and TotalAmount.
func (w *withdrawResponseMapper) ToWithdrawsAmountYearlyResponses(s []*record.WithdrawYearlyAmount) []*response.WithdrawYearlyAmountResponse {
	var withdrawResponses []*response.WithdrawYearlyAmountResponse
	for _, withdraw := range s {
		withdrawResponses = append(withdrawResponses, w.ToWithdrawAmountYearlyResponse(withdraw))
	}
	return withdrawResponses
}
