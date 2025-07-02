package responseservice

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// transferResponseMapper provides methods to map TransferRecord domain models to TransferResponse API-compatible response types.
type transferResponseMapper struct {
}

// NewTransferResponseMapper returns a new instance of transferResponseMapper,
// which provides methods to map TransferRecord domain models to API-compatible
// TransferResponse types.
func NewTransferResponseMapper() *transferResponseMapper {
	return &transferResponseMapper{}
}

// ToTransferResponse converts a single transfer record into a TransferResponse.
//
// Args:
//   - transfer: The transfer record to be converted.
//
// Returns:
//   - A pointer to a response.TransferResponse containing the mapped data.
func (s *transferResponseMapper) ToTransferResponse(transfer *record.TransferRecord) *response.TransferResponse {
	return &response.TransferResponse{
		ID:             transfer.ID,
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: transfer.TransferAmount,
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}
}

// ToTransfersResponse converts multiple transfer records into a slice of TransferResponse.
//
// Args:
//   - transfers: A slice of transfer records to be converted.
//
// Returns:
//   - A slice of pointers to response.TransferResponse containing the mapped data.
func (s *transferResponseMapper) ToTransfersResponse(transfers []*record.TransferRecord) []*response.TransferResponse {
	var responses []*response.TransferResponse

	for _, response := range transfers {
		responses = append(responses, s.ToTransferResponse(response))
	}

	return responses
}

// ToTransferResponseDeleteAt converts a single transfer record into a TransferResponseDeleteAt.
//
// Args:
//   - transfer: The transfer record to be converted.
//
// Returns:
//   - A pointer to a response.TransferResponseDeleteAt containing the mapped data.
func (s *transferResponseMapper) ToTransferResponseDeleteAt(transfer *record.TransferRecord) *response.TransferResponseDeleteAt {
	return &response.TransferResponseDeleteAt{
		ID:             transfer.ID,
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: transfer.TransferAmount,
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
		DeletedAt:      transfer.DeletedAt,
	}
}

// ToTransfersResponseDeleteAt converts multiple soft-deleted transfer records into a slice of TransferResponseDeleteAt.
//
// Args:
//   - transfers: A slice of pointers to TransferRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferResponseDeleteAt containing the mapped data, including
//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime, CreatedAt,
//     UpdatedAt, and DeletedAt.
func (s *transferResponseMapper) ToTransfersResponseDeleteAt(transfers []*record.TransferRecord) []*response.TransferResponseDeleteAt {
	var responses []*response.TransferResponseDeleteAt

	for _, response := range transfers {
		responses = append(responses, s.ToTransferResponseDeleteAt(response))
	}

	return responses
}

// ToTransferResponseMonthStatusSuccess converts a monthly success status transfer record
// into a TransferResponseMonthStatusSuccess.
//
// Args:
//   - s: A pointer to a TransferRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferResponseMonthStatusSuccess containing the mapped data, including
//     Year, Month, TotalSuccess, and TotalAmount.
func (t *transferResponseMapper) ToTransferResponseMonthStatusSuccess(s *record.TransferRecordMonthStatusSuccess) *response.TransferResponseMonthStatusSuccess {
	return &response.TransferResponseMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTransferResponsesMonthStatusSuccess converts multiple monthly success status transfer records
// into a slice of TransferResponseMonthStatusSuccess.
//
// Args:
//   - Transfers: A slice of pointers to TransferRecordMonthStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferResponseMonthStatusSuccess containing the mapped data, including
//     Year, Month, TotalSuccess, and TotalAmount.
func (t *transferResponseMapper) ToTransferResponsesMonthStatusSuccess(Transfers []*record.TransferRecordMonthStatusSuccess) []*response.TransferResponseMonthStatusSuccess {
	var TransferRecords []*response.TransferResponseMonthStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferResponseMonthStatusSuccess(Transfer))
	}

	return TransferRecords
}

// ToTransferResponseYearStatusSuccess converts a yearly success status transfer record
// into a TransferResponseYearStatusSuccess.
//
// Args:
//   - s: A pointer to a TransferRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferResponseYearStatusSuccess containing the mapped data, including
//     Year, TotalSuccess, and TotalAmount.
func (t *transferResponseMapper) ToTransferResponseYearStatusSuccess(s *record.TransferRecordYearStatusSuccess) *response.TransferResponseYearStatusSuccess {
	return &response.TransferResponseYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  s.TotalAmount,
	}
}

// ToTransferResponsesYearStatusSuccess converts multiple yearly success status transfer records
// into a slice of TransferResponseYearStatusSuccess.
//
// Args:
//   - Transfers: A slice of pointers to TransferRecordYearStatusSuccess containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferResponseYearStatusSuccess containing the mapped data, including
//     Year, TotalSuccess, and TotalAmount.
func (t *transferResponseMapper) ToTransferResponsesYearStatusSuccess(Transfers []*record.TransferRecordYearStatusSuccess) []*response.TransferResponseYearStatusSuccess {
	var TransferRecords []*response.TransferResponseYearStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferResponseYearStatusSuccess(Transfer))
	}

	return TransferRecords
}

// ToTransferResponseMonthStatusFailed converts a monthly failed status transfer record
// into a TransferResponseMonthStatusFailed.
//
// Args:
//   - s: A pointer to a TransferRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferResponseMonthStatusFailed containing the mapped data, including
//     Year, Month, TotalFailed, and TotalAmount.
func (t *transferResponseMapper) ToTransferResponseMonthStatusFailed(s *record.TransferRecordMonthStatusFailed) *response.TransferResponseMonthStatusFailed {
	return &response.TransferResponseMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTransferResponsesMonthStatusFailed converts multiple monthly failed status transfer records
// into a slice of TransferResponseMonthStatusFailed.
//
// Args:
//   - Transfers: A slice of pointers to TransferRecordMonthStatusFailed containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferResponseMonthStatusFailed containing the mapped data, including
//     Year, Month, TotalFailed, and TotalAmount.
func (t *transferResponseMapper) ToTransferResponsesMonthStatusFailed(Transfers []*record.TransferRecordMonthStatusFailed) []*response.TransferResponseMonthStatusFailed {
	var TransferRecords []*response.TransferResponseMonthStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferResponseMonthStatusFailed(Transfer))
	}

	return TransferRecords
}

// ToTransferResponseYearStatusFailed converts a yearly failed status transfer record
// into a TransferResponseYearStatusFailed.
//
// Args:
//   - s: A pointer to a TransferRecordYearStatusFailed containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *transferResponseMapper) ToTransferResponseYearStatusFailed(s *record.TransferRecordYearStatusFailed) *response.TransferResponseYearStatusFailed {
	return &response.TransferResponseYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: s.TotalAmount,
	}
}

// ToTransferResponsesYearStatusFailed converts multiple yearly failed status transfer records
// into a slice of TransferResponseYearStatusFailed.
//
// Args:
//   - Transfers: A slice of pointers to TransferRecordYearStatusFailed containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferResponseYearStatusFailed containing the mapped data, including
//     Year, TotalFailed, and TotalAmount.
func (t *transferResponseMapper) ToTransferResponsesYearStatusFailed(Transfers []*record.TransferRecordYearStatusFailed) []*response.TransferResponseYearStatusFailed {
	var TransferRecords []*response.TransferResponseYearStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferResponseYearStatusFailed(Transfer))
	}

	return TransferRecords
}

// ToTransferResponseMonthAmount maps a TransferMonthAmount record to a TransferMonthAmountResponse.
//
// Args:
//   - s: A pointer to a TransferMonthAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferMonthAmountResponse containing the mapped data.
func (t *transferResponseMapper) ToTransferResponseMonthAmount(s *record.TransferMonthAmount) *response.TransferMonthAmountResponse {
	return &response.TransferMonthAmountResponse{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferResponsesMonthAmount converts a slice of TransferMonthAmount records
// into a slice of TransferMonthAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransferMonthAmount containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferMonthAmountResponse containing the mapped data.
func (t *transferResponseMapper) ToTransferResponsesMonthAmount(s []*record.TransferMonthAmount) []*response.TransferMonthAmountResponse {
	var transferResponses []*response.TransferMonthAmountResponse
	for _, transfer := range s {
		transferResponses = append(transferResponses, t.ToTransferResponseMonthAmount(transfer))
	}
	return transferResponses
}

// ToTransferResponseYearAmount converts a yearly transfer amount record into a TransferYearAmountResponse.
//
// Args:
//   - s: A pointer to a TransferYearAmount containing the data to be mapped.
//
// Returns:
//   - A pointer to a TransferYearAmountResponse containing the mapped data, including Year and TotalAmount.
func (t *transferResponseMapper) ToTransferResponseYearAmount(s *record.TransferYearAmount) *response.TransferYearAmountResponse {
	return &response.TransferYearAmountResponse{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferResponsesYearAmount converts multiple yearly transfer amount records into a slice of TransferYearAmountResponse API-compatible response types.
//
// Args:
//   - s: A slice of pointers to TransferYearAmount containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to TransferYearAmountResponse containing the mapped data.
func (t *transferResponseMapper) ToTransferResponsesYearAmount(s []*record.TransferYearAmount) []*response.TransferYearAmountResponse {
	var transferResponses []*response.TransferYearAmountResponse
	for _, transfer := range s {
		transferResponses = append(transferResponses, t.ToTransferResponseYearAmount(transfer))
	}
	return transferResponses
}
