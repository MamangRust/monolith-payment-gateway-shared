package recordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

// transferRecordMapper maps a Transfer entity to a TransferRecord struct.
type transferRecordMapper struct {
}

// NewTransferRecordMapper returns a new instance of transferRecordMapper which provides methods to map Transfer database rows to TransferRecord domain models.
func NewTransferRecordMapper() *transferRecordMapper {
	return &transferRecordMapper{}
}

// ToTransferRecord maps a Transfer database row to a TransferRecord domain model.
//
// Args:
//   - transfer: A pointer to a Transfer representing the database row.
//
// Returns:
//   - A pointer to a TransferRecord containing the mapped data, including
//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
//     CreatedAt, UpdatedAt, and DeletedAt.
func (t *transferRecordMapper) ToTransferRecord(transfer *db.Transfer) *record.TransferRecord {
	var deletedAt *string

	return &record.TransferRecord{
		ID:             int(transfer.TransferID),
		TransferNo:     transfer.TransferNo.String(),
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime.String(),
		CreatedAt:      transfer.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      transfer.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

// ToTransfersRecord maps a slice of Transfer database rows to a slice of TransferRecord domain models.
//
// Args:
//   - transfers: A slice of pointers to Transfer structs representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecord structs containing the mapped data.
func (s *transferRecordMapper) ToTransfersRecord(transfers []*db.Transfer) []*record.TransferRecord {
	var transferRecords []*record.TransferRecord

	for _, transfer := range transfers {
		transferRecords = append(transferRecords, s.ToTransferRecord(transfer))
	}

	return transferRecords
}

// ToTransferRecordAll maps a GetTransfersRow database row to a TransferRecord domain model.
//
// Args:
//   - transfer: A pointer to a GetTransfersRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecord containing the mapped data, including
//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
//     CreatedAt, UpdatedAt, and DeletedAt.
func (t *transferRecordMapper) ToTransferRecordAll(transfer *db.GetTransfersRow) *record.TransferRecord {
	var deletedAt *string

	return &record.TransferRecord{
		ID:             int(transfer.TransferID),
		TransferNo:     transfer.TransferNo.String(),
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime.String(),
		CreatedAt:      transfer.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      transfer.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

func (s *transferRecordMapper) ToTransfersRecordAll(transfers []*db.GetTransfersRow) []*record.TransferRecord {
	var transferRecords []*record.TransferRecord

	for _, transfer := range transfers {
		transferRecords = append(transferRecords, s.ToTransferRecordAll(transfer))
	}

	return transferRecords
}

// ToTransferRecordActive maps a GetActiveTransfersRow database row to a TransferRecord domain model.
//
// Args:
//   - transfer: A pointer to a GetActiveTransfersRow representing the active transfer database row.
//
// Returns:
//   - A pointer to a TransferRecord containing the mapped data, including
//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
//     CreatedAt, UpdatedAt, and DeletedAt.
func (t *transferRecordMapper) ToTransferRecordActive(transfer *db.GetActiveTransfersRow) *record.TransferRecord {
	var deletedAt *string

	return &record.TransferRecord{
		ID:             int(transfer.TransferID),
		TransferNo:     transfer.TransferNo.String(),
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime.String(),
		CreatedAt:      transfer.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      transfer.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

// ToTransfersRecordActive maps a slice of GetActiveTransfersRow database rows to a slice of TransferRecord domain models.
//
// Args:
//   - transfers: A slice of pointers to GetActiveTransfersRow structs representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecord structs containing the mapped data.
func (s *transferRecordMapper) ToTransfersRecordActive(transfers []*db.GetActiveTransfersRow) []*record.TransferRecord {
	var transferRecords []*record.TransferRecord

	for _, transfer := range transfers {
		transferRecords = append(transferRecords, s.ToTransferRecordActive(transfer))
	}

	return transferRecords
}

// ToTransferRecordTrashed maps a GetTrashedTransfersRow database row to a TransferRecord domain model.
//
// Args:
//   - transfer: A pointer to a GetTrashedTransfersRow representing the active transfer database row.
//
// Returns:
//   - A pointer to a TransferRecord containing the mapped data, including
//     ID, TransferNo, TransferFrom, TransferTo, TransferAmount, TransferTime,
//     CreatedAt, UpdatedAt, and DeletedAt.
func (t *transferRecordMapper) ToTransferRecordTrashed(transfer *db.GetTrashedTransfersRow) *record.TransferRecord {
	var deletedAt *string

	if transfer.DeletedAt.Valid {
		formatedDeletedAt := transfer.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.TransferRecord{
		ID:             int(transfer.TransferID),
		TransferNo:     transfer.TransferNo.String(),
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime.String(),
		CreatedAt:      transfer.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      transfer.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

// ToTransfersRecordTrashed maps a slice of GetTrashedTransfersRow database rows to a slice of TransferRecord domain models.
//
// Args:
//   - transfers: A slice of pointers to GetTrashedTransfersRow structs representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecord structs containing the mapped data.
func (s *transferRecordMapper) ToTransfersRecordTrashed(transfers []*db.GetTrashedTransfersRow) []*record.TransferRecord {
	var transferRecords []*record.TransferRecord

	for _, transfer := range transfers {
		transferRecords = append(transferRecords, s.ToTransferRecordTrashed(transfer))
	}

	return transferRecords
}

// ToTransferRecordMonthStatusSuccess maps a GetMonthTransferStatusSuccessRow database row
// to a TransferRecordMonthStatusSuccess domain model.
//
// Args:
//   - s: A pointer to a GetMonthTransferStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordMonthStatusSuccess containing the mapped data, including
//     Year, Month, TotalSuccess, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordMonthStatusSuccess(s *db.GetMonthTransferStatusSuccessRow) *record.TransferRecordMonthStatusSuccess {
	return &record.TransferRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransferRecordsMonthStatusSuccess maps a slice of GetMonthTransferStatusSuccessRow database rows
// to a slice of TransferRecordMonthStatusSuccess domain models.
//
// Args:
//   - Transfers: A slice of pointers to GetMonthTransferStatusSuccessRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordsMonthStatusSuccess(Transfers []*db.GetMonthTransferStatusSuccessRow) []*record.TransferRecordMonthStatusSuccess {
	var TransferRecords []*record.TransferRecordMonthStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordMonthStatusSuccess(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordYearStatusSuccess maps a GetYearlyTransferStatusSuccessRow database row
// to a TransferRecordYearStatusSuccess domain model.
//
// Args:
//   - s: A pointer to a GetYearlyTransferStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordYearStatusSuccess(s *db.GetYearlyTransferStatusSuccessRow) *record.TransferRecordYearStatusSuccess {
	return &record.TransferRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransferRecordsYearStatusSuccess maps a slice of GetYearlyTransferStatusSuccessRow database rows
// to a slice of TransferRecordYearStatusSuccess domain models.
//
// Args:
//   - Transfers: A slice of pointers to GetYearlyTransferStatusSuccessRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordsYearStatusSuccess(Transfers []*db.GetYearlyTransferStatusSuccessRow) []*record.TransferRecordYearStatusSuccess {
	var TransferRecords []*record.TransferRecordYearStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordYearStatusSuccess(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordMonthStatusFailed maps a GetMonthTransferStatusFailedRow database row
// to a TransferRecordMonthStatusFailed domain model.
//
// Args:
//   - s: A pointer to a GetMonthTransferStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordMonthStatusFailed(s *db.GetMonthTransferStatusFailedRow) *record.TransferRecordMonthStatusFailed {
	return &record.TransferRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferRecordsMonthStatusFailed maps a slice of GetMonthTransferStatusFailedRow database rows
// to a slice of TransferRecordMonthStatusFailed domain models.
//
// Args:
//   - Transfers: A slice of pointers to GetMonthTransferStatusFailedRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordsMonthStatusFailed(Transfers []*db.GetMonthTransferStatusFailedRow) []*record.TransferRecordMonthStatusFailed {
	var TransferRecords []*record.TransferRecordMonthStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordMonthStatusFailed(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordYearStatusFailed maps a GetYearlyTransferStatusFailedRow database row
// to a TransferRecordYearStatusFailed domain model.
//
// Args:
//   - s: A pointer to a GetYearlyTransferStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordYearStatusFailed(s *db.GetYearlyTransferStatusFailedRow) *record.TransferRecordYearStatusFailed {
	return &record.TransferRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferRecordsYearStatusFailed maps a slice of GetYearlyTransferStatusFailedRow database rows
// to a slice of TransferRecordYearStatusFailed domain models.
//
// Args:
//   - Transfers: A slice of pointers to GetYearlyTransferStatusFailedRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordsYearStatusFailed(Transfers []*db.GetYearlyTransferStatusFailedRow) []*record.TransferRecordYearStatusFailed {
	var TransferRecords []*record.TransferRecordYearStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordYearStatusFailed(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordMonthStatusSuccessCardNumber maps a GetMonthTransferStatusSuccessCardNumberRow database row
// to a TransferRecordMonthStatusSuccess domain model.
//
// Args:
//   - s: A pointer to GetMonthTransferStatusSuccessCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordMonthStatusSuccessCardNumber(s *db.GetMonthTransferStatusSuccessCardNumberRow) *record.TransferRecordMonthStatusSuccess {
	return &record.TransferRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransferRecordsMonthStatusSuccessCardNumber maps a slice of GetMonthTransferStatusSuccessCardNumberRow database rows
// to a slice of TransferRecordMonthStatusSuccess domain models.
//
// Args:
//   - Transfers: A slice of pointers to GetMonthTransferStatusSuccessCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordsMonthStatusSuccessCardNumber(Transfers []*db.GetMonthTransferStatusSuccessCardNumberRow) []*record.TransferRecordMonthStatusSuccess {
	var TransferRecords []*record.TransferRecordMonthStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordMonthStatusSuccessCardNumber(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordYearStatusSuccessCardNumber maps a GetYearlyTransferStatusSuccessCardNumberRow database row
// to a TransferRecordYearStatusSuccess domain model.
//
// Args:
//   - s: A pointer to a GetYearlyTransferStatusSuccessCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordYearStatusSuccessCardNumber(s *db.GetYearlyTransferStatusSuccessCardNumberRow) *record.TransferRecordYearStatusSuccess {
	return &record.TransferRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransferRecordsYearStatusSuccessCardNumber maps a slice of GetYearlyTransferStatusSuccessCardNumberRow database rows
// to a slice of TransferRecordYearStatusSuccess domain models.
//
// Args:
//   - Transfers: A slice of pointers to GetYearlyTransferStatusSuccessCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordsYearStatusSuccessCardNumber(Transfers []*db.GetYearlyTransferStatusSuccessCardNumberRow) []*record.TransferRecordYearStatusSuccess {
	var TransferRecords []*record.TransferRecordYearStatusSuccess

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordYearStatusSuccessCardNumber(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordMonthStatusFailedCardNumber maps a GetMonthTransferStatusFailedCardNumberRow database row
// to a TransferRecordMonthStatusFailed domain model.
//
// Args:
//   - s: A pointer to a GetMonthTransferStatusFailedCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordMonthStatusFailedCardNumber(s *db.GetMonthTransferStatusFailedCardNumberRow) *record.TransferRecordMonthStatusFailed {
	return &record.TransferRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferRecordsMonthStatusFailedCardNumber maps a slice of GetMonthTransferStatusFailedCardNumberRow database rows
// to a slice of TransferRecordMonthStatusFailed domain models.
//
// Args:
//   - Transfers: A slice of pointers to GetMonthTransferStatusFailedCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordsMonthStatusFailedCardNumber(Transfers []*db.GetMonthTransferStatusFailedCardNumberRow) []*record.TransferRecordMonthStatusFailed {
	var TransferRecords []*record.TransferRecordMonthStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordMonthStatusFailedCardNumber(Transfer))
	}

	return TransferRecords
}

// ToTransferRecordYearStatusFailedCardNumber maps a GetYearlyTransferStatusFailedCardNumberRow database row
// to a TransferRecordYearStatusFailed domain model.
//
// Args:
//   - s: A pointer to a GetYearlyTransferStatusFailedCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordYearStatusFailedCardNumber(s *db.GetYearlyTransferStatusFailedCardNumberRow) *record.TransferRecordYearStatusFailed {
	return &record.TransferRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransferRecordsYearStatusFailedCardNumber maps a slice of GetYearlyTransferStatusFailedCardNumberRow database rows
// to a slice of TransferRecordYearStatusFailed domain models.
//
// Args:
//   - Transfers: A slice of pointers to GetYearlyTransferStatusFailedCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transferRecordMapper) ToTransferRecordsYearStatusFailedCardNumber(Transfers []*db.GetYearlyTransferStatusFailedCardNumberRow) []*record.TransferRecordYearStatusFailed {
	var TransferRecords []*record.TransferRecordYearStatusFailed

	for _, Transfer := range Transfers {
		TransferRecords = append(TransferRecords, t.ToTransferRecordYearStatusFailedCardNumber(Transfer))
	}

	return TransferRecords
}

// ToTransferMonthAmount maps a GetMonthlyTransferAmountsRow database row
// to a TransferMonthAmount domain model.
//
// Args:
//   - ss: A pointer to a GetMonthlyTransferAmountsRow representing the database row.
//
// Returns:
//   - A pointer to a TransferMonthAmount containing the mapped data, including Month
//     and TotalAmount.
func (s *transferRecordMapper) ToTransferMonthAmount(ss *db.GetMonthlyTransferAmountsRow) *record.TransferMonthAmount {
	return &record.TransferMonthAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferMonthAmounts maps a slice of GetMonthlyTransferAmountsRow database rows
// to a slice of TransferMonthAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetMonthlyTransferAmountsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferMonthAmount containing the mapped data, including Month
//     and TotalAmount.
func (s *transferRecordMapper) ToTransferMonthAmounts(ss []*db.GetMonthlyTransferAmountsRow) []*record.TransferMonthAmount {
	var transferRecords []*record.TransferMonthAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferMonthAmount(transfer))
	}

	return transferRecords
}

// ToTransferYearAmount maps a GetYearlyTransferAmountsRow database row to a TransferYearAmount domain model.
//
// Args:
//   - ss: A pointer to a GetYearlyTransferAmountsRow representing the database row.
//
// Returns:
//   - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferRecordMapper) ToTransferYearAmount(ss *db.GetYearlyTransferAmountsRow) *record.TransferYearAmount {
	return &record.TransferYearAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferYearAmounts maps a slice of GetYearlyTransferAmountsRow database rows
// to a slice of TransferYearAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetYearlyTransferAmountsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferRecordMapper) ToTransferYearAmounts(ss []*db.GetYearlyTransferAmountsRow) []*record.TransferYearAmount {
	var transferRecords []*record.TransferYearAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferYearAmount(transfer))
	}

	return transferRecords
}

// ToTransferMonthAmountSender maps a GetMonthlyTransferAmountsBySenderCardNumberRow database row
// to a TransferMonthAmount domain model.
//
// Args:
//   - ss: A pointer to a GetMonthlyTransferAmountsBySenderCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferMonthAmount containing the mapped data, including Month
//     and TotalAmount.
func (s *transferRecordMapper) ToTransferMonthAmountSender(ss *db.GetMonthlyTransferAmountsBySenderCardNumberRow) *record.TransferMonthAmount {
	return &record.TransferMonthAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferMonthAmountsSender maps a slice of GetMonthlyTransferAmountsBySenderCardNumberRow
// database rows to a slice of TransferMonthAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetMonthlyTransferAmountsBySenderCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transferRecordMapper) ToTransferMonthAmountsSender(ss []*db.GetMonthlyTransferAmountsBySenderCardNumberRow) []*record.TransferMonthAmount {
	var transferRecords []*record.TransferMonthAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferMonthAmountSender(transfer))
	}

	return transferRecords
}

// ToTransferYearAmountSender maps a GetYearlyTransferAmountsBySenderCardNumberRow database row
// to a TransferYearAmount domain model.
//
// Args:
//   - ss: A pointer to a GetYearlyTransferAmountsBySenderCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferRecordMapper) ToTransferYearAmountSender(ss *db.GetYearlyTransferAmountsBySenderCardNumberRow) *record.TransferYearAmount {
	return &record.TransferYearAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferYearAmountsSender maps a slice of GetYearlyTransferAmountsBySenderCardNumberRow database rows
// to a slice of TransferYearAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetYearlyTransferAmountsBySenderCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferRecordMapper) ToTransferYearAmountsSender(ss []*db.GetYearlyTransferAmountsBySenderCardNumberRow) []*record.TransferYearAmount {
	var transferRecords []*record.TransferYearAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferYearAmountSender(transfer))
	}

	return transferRecords
}

// ToTransferMonthAmountReceiver maps a GetMonthlyTransferAmountsByReceiverCardNumberRow database row
// to a TransferMonthAmount domain model.
//
// Args:
//   - ss: A pointer to a GetMonthlyTransferAmountsByReceiverCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transferRecordMapper) ToTransferMonthAmountReceiver(ss *db.GetMonthlyTransferAmountsByReceiverCardNumberRow) *record.TransferMonthAmount {
	return &record.TransferMonthAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferMonthAmountsReceiver maps a slice of GetMonthlyTransferAmountsByReceiverCardNumberRow database rows
// to a slice of TransferMonthAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetMonthlyTransferAmountsByReceiverCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transferRecordMapper) ToTransferMonthAmountsReceiver(ss []*db.GetMonthlyTransferAmountsByReceiverCardNumberRow) []*record.TransferMonthAmount {
	var transferRecords []*record.TransferMonthAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferMonthAmountReceiver(transfer))
	}

	return transferRecords
}

// ToTransferYearAmountReceiver maps a GetYearlyTransferAmountsByReceiverCardNumberRow database row
// to a TransferYearAmount domain model.
//
// Args:
//   - ss: A pointer to a GetYearlyTransferAmountsByReceiverCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferRecordMapper) ToTransferYearAmountReceiver(ss *db.GetYearlyTransferAmountsByReceiverCardNumberRow) *record.TransferYearAmount {
	return &record.TransferYearAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalTransferAmount),
	}
}

// ToTransferYearAmountsReceiver maps a slice of GetYearlyTransferAmountsByReceiverCardNumberRow database rows
// to a slice of TransferYearAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetYearlyTransferAmountsByReceiverCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransferYearAmount containing the mapped data, including Year and TotalAmount.
func (s *transferRecordMapper) ToTransferYearAmountsReceiver(ss []*db.GetYearlyTransferAmountsByReceiverCardNumberRow) []*record.TransferYearAmount {
	var transferRecords []*record.TransferYearAmount

	for _, transfer := range ss {
		transferRecords = append(transferRecords, s.ToTransferYearAmountReceiver(transfer))
	}

	return transferRecords
}
