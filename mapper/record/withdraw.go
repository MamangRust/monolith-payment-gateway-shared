package recordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

// withdrawRecordMapper provides methods to map Withdraw database rows to WithdrawRecord domain models.
type withdrawRecordMapper struct{}

// NewWithdrawRecordMapper creates a new instance of withdrawRecordMapper.
// It returns a pointer to a withdrawRecordMapper, which provides methods
// for mapping Withdraw database rows to WithdrawRecord domain models.
func NewWithdrawRecordMapper() *withdrawRecordMapper {
	return &withdrawRecordMapper{}
}

// ToWithdrawRecord maps a Withdraw database row to a WithdrawRecord domain model.
//
// Args:
//   - withdraw: A pointer to a Withdraw representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawRecordMapper) ToWithdrawRecord(withdraw *db.Withdraw) *record.WithdrawRecord {
	var deletedAt *string

	if withdraw.DeletedAt.Valid {
		formatedDeletedAt := withdraw.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.WithdrawRecord{
		ID:             int(withdraw.WithdrawID),
		WithdrawNo:     withdraw.WithdrawNo.String(),
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime.String(),
		CreatedAt:      withdraw.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      withdraw.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

// ToWithdrawsRecord maps a slice of Withdraw database rows to a slice of WithdrawRecord
// domain models.
//
// Args:
//   - withdraws: A slice of pointers to Withdraw representing the database rows.
//
// Returns:
//   - A slice of pointers to WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawRecordMapper) ToWithdrawsRecord(withdraws []*db.Withdraw) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecord(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawByCardNumberRecord maps a database row representing a withdraw record
// associated with a given card number to a WithdrawRecord domain model.
//
// Args:
//   - withdraw: A pointer to a GetWithdrawsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawRecordMapper) ToWithdrawByCardNumberRecord(withdraw *db.GetWithdrawsByCardNumberRow) *record.WithdrawRecord {
	var deletedAt *string

	if withdraw.DeletedAt.Valid {
		formatedDeletedAt := withdraw.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.WithdrawRecord{
		ID:             int(withdraw.WithdrawID),
		WithdrawNo:     withdraw.WithdrawNo.String(),
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime.String(),
		CreatedAt:      withdraw.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      withdraw.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

// ToWithdrawsByCardNumberRecord maps a slice of GetWithdrawsByCardNumberRow database rows to a slice of WithdrawRecord domain models.
//
// Args:
//   - withdraws: A slice of pointers to GetWithdrawsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawRecordMapper) ToWithdrawsByCardNumberRecord(withdraws []*db.GetWithdrawsByCardNumberRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawByCardNumberRecord(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawRecordAll maps a GetWithdrawsRow database row to a WithdrawRecord domain model.
//
// Args:
//   - withdraw: A pointer to a GetWithdrawsRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and optionally DeletedAt if it is valid.
func (s *withdrawRecordMapper) ToWithdrawRecordAll(withdraw *db.GetWithdrawsRow) *record.WithdrawRecord {
	var deletedAt *string

	if withdraw.DeletedAt.Valid {
		formatedDeletedAt := withdraw.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.WithdrawRecord{
		ID:             int(withdraw.WithdrawID),
		WithdrawNo:     withdraw.WithdrawNo.String(),
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime.String(),
		CreatedAt:      withdraw.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      withdraw.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

// ToWithdrawsRecordALl maps a slice of GetWithdrawsRow database rows to a slice of WithdrawRecord domain models.
//
// Args:
//   - withdraws: A slice of pointers to GetWithdrawsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and optionally DeletedAt if it is valid.
func (s *withdrawRecordMapper) ToWithdrawsRecordALl(withdraws []*db.GetWithdrawsRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecordAll(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawRecordActive maps a single GetActiveWithdrawsRow database row to a WithdrawRecord
// domain model.
//
// Args:
//   - withdraw: A pointer to a GetActiveWithdrawsRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and
//     optionally DeletedAt if it is valid.
func (s *withdrawRecordMapper) ToWithdrawRecordActive(withdraw *db.GetActiveWithdrawsRow) *record.WithdrawRecord {
	var deletedAt *string

	if withdraw.DeletedAt.Valid {
		formatedDeletedAt := withdraw.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.WithdrawRecord{
		ID:             int(withdraw.WithdrawID),
		WithdrawNo:     withdraw.WithdrawNo.String(),
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime.String(),
		CreatedAt:      withdraw.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      withdraw.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

// ToWithdrawsRecordActive maps a slice of GetActiveWithdrawsRow database rows to a slice of WithdrawRecord domain models.
//
// Args:
//   - withdraws: A slice of pointers to GetActiveWithdrawsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to WithdrawRecord containing the mapped data, including
//     ID, WithdrawNo, CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and
//     optionally DeletedAt if it is valid.
func (s *withdrawRecordMapper) ToWithdrawsRecordActive(withdraws []*db.GetActiveWithdrawsRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecordActive(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawRecordTrashed maps a GetTrashedWithdrawsRow database row to a WithdrawRecord domain model.
// It is intended for use with database rows that contain trashed withdraw records.
// It returns a pointer to a WithdrawRecord containing the mapped data, including ID, WithdrawNo,
// CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawRecordMapper) ToWithdrawRecordTrashed(withdraw *db.GetTrashedWithdrawsRow) *record.WithdrawRecord {
	var deletedAt *string

	if withdraw.DeletedAt.Valid {
		formatedDeletedAt := withdraw.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.WithdrawRecord{
		ID:             int(withdraw.WithdrawID),
		WithdrawNo:     withdraw.WithdrawNo.String(),
		CardNumber:     withdraw.CardNumber,
		WithdrawAmount: int(withdraw.WithdrawAmount),
		WithdrawTime:   withdraw.WithdrawTime.String(),
		CreatedAt:      withdraw.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:      withdraw.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:      deletedAt,
	}
}

// ToWithdrawsRecordTrashed maps a slice of GetTrashedWithdrawsRow database rows to a slice of WithdrawRecord
// domain models. It is intended for use with database rows that contain trashed withdraw records.
//
// Args:
//   - withdraws: A slice of pointers to GetTrashedWithdrawsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to WithdrawRecord containing the mapped data, including ID, WithdrawNo,
//     CardNumber, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *withdrawRecordMapper) ToWithdrawsRecordTrashed(withdraws []*db.GetTrashedWithdrawsRow) []*record.WithdrawRecord {
	var withdrawRecords []*record.WithdrawRecord

	for _, withdraw := range withdraws {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawRecordTrashed(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawRecordMonthStatusSuccess maps a database row representing monthly
// successful withdraw statistics to a WithdrawRecordMonthStatusSuccess domain model.
//
// Args:
//   - s: A pointer to a GetMonthWithdrawStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordMonthStatusSuccess containing the mapped data,
//     including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordMonthStatusSuccess(s *db.GetMonthWithdrawStatusSuccessRow) *record.WithdrawRecordMonthStatusSuccess {
	return &record.WithdrawRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToWithdrawRecordsMonthStatusSuccess converts a slice of GetMonthWithdrawStatusSuccessRow
// database rows to a slice of WithdrawRecordMonthStatusSuccess domain models.
//
// Args:
//   - Withdraws: A slice of pointers to GetMonthWithdrawStatusSuccessRow representing
//     the database rows containing monthly successful withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawRecordMonthStatusSuccess, each containing
//     the mapped data including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordsMonthStatusSuccess(Withdraws []*db.GetMonthWithdrawStatusSuccessRow) []*record.WithdrawRecordMonthStatusSuccess {
	var WithdrawRecords []*record.WithdrawRecordMonthStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordMonthStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordYearStatusSuccess maps a database row representing yearly
// successful withdraw statistics to a WithdrawRecordYearStatusSuccess domain model.
//
// Args:
//   - s: A pointer to a GetYearlyWithdrawStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordYearStatusSuccess containing the mapped data,
//     including Year, TotalSuccess, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordYearStatusSuccess(s *db.GetYearlyWithdrawStatusSuccessRow) *record.WithdrawRecordYearStatusSuccess {
	return &record.WithdrawRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToWithdrawRecordsYearStatusSuccess maps a slice of GetYearlyWithdrawStatusSuccessRow
// database rows to a slice of WithdrawRecordYearStatusSuccess domain models.
//
// Args:
//   - Withdraws: A slice of pointers to GetYearlyWithdrawStatusSuccessRow representing
//     the database rows containing yearly successful withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawRecordYearStatusSuccess, each containing
//     the mapped data including Year, TotalSuccess, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordsYearStatusSuccess(Withdraws []*db.GetYearlyWithdrawStatusSuccessRow) []*record.WithdrawRecordYearStatusSuccess {
	var WithdrawRecords []*record.WithdrawRecordYearStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordYearStatusSuccess(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordMonthStatusFailed maps a database row representing monthly
// failed withdraw statistics to a WithdrawRecordMonthStatusFailed domain model.
//
// Args:
//   - s: A pointer to a GetMonthWithdrawStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordMonthStatusFailed containing the mapped data,
//     including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordMonthStatusFailed(s *db.GetMonthWithdrawStatusFailedRow) *record.WithdrawRecordMonthStatusFailed {
	return &record.WithdrawRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToWithdrawRecordsMonthStatusFailed maps a slice of GetMonthWithdrawStatusFailedRow
// database rows to a slice of WithdrawRecordMonthStatusFailed domain models.
//
// Args:
//   - Withdraws: A slice of pointers to GetMonthWithdrawStatusFailedRow representing
//     the database rows containing monthly failed withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawRecordMonthStatusFailed, each containing
//     the mapped data including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordsMonthStatusFailed(Withdraws []*db.GetMonthWithdrawStatusFailedRow) []*record.WithdrawRecordMonthStatusFailed {
	var WithdrawRecords []*record.WithdrawRecordMonthStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordMonthStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordYearStatusFailed maps a database row representing yearly
// failed withdraw statistics to a WithdrawRecordYearStatusFailed domain model.
//
// Args:
//   - s: A pointer to a GetYearlyWithdrawStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordYearStatusFailed containing the mapped data,
//     including Year, TotalFailed, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordYearStatusFailed(s *db.GetYearlyWithdrawStatusFailedRow) *record.WithdrawRecordYearStatusFailed {
	return &record.WithdrawRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToWithdrawRecordsYearStatusFailed maps a slice of GetYearlyWithdrawStatusFailedRow
// database rows to a slice of WithdrawRecordYearStatusFailed domain models.
//
// Args:
//   - Withdraws: A slice of pointers to GetYearlyWithdrawStatusFailedRow representing
//     the database rows containing yearly failed withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawRecordYearStatusFailed, each containing
//     the mapped data including Year, TotalFailed, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordsYearStatusFailed(Withdraws []*db.GetYearlyWithdrawStatusFailedRow) []*record.WithdrawRecordYearStatusFailed {
	var WithdrawRecords []*record.WithdrawRecordYearStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordYearStatusFailed(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordMonthStatusSuccessCardNumber maps a database row representing
// monthly successful withdraw statistics, filtered by card number, to a
// WithdrawRecordMonthStatusSuccess domain model.
//
// Args:
//   - s: A pointer to a GetMonthWithdrawStatusSuccessCardNumberRow representing
//     the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordMonthStatusSuccess containing the mapped data,
//     including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordMonthStatusSuccessCardNumber(s *db.GetMonthWithdrawStatusSuccessCardNumberRow) *record.WithdrawRecordMonthStatusSuccess {
	return &record.WithdrawRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToWithdrawRecordsMonthStatusSuccessCardNumber maps a slice of
// GetMonthWithdrawStatusSuccessCardNumberRow database rows to a slice of
// WithdrawRecordMonthStatusSuccess domain models, filtered by card number.
//
// Args:
//   - Withdraws: A slice of pointers to GetMonthWithdrawStatusSuccessCardNumberRow
//     representing the database rows with monthly successful withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawRecordMonthStatusSuccess, each containing
//     the mapped data including Year, Month, TotalSuccess, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordsMonthStatusSuccessCardNumber(Withdraws []*db.GetMonthWithdrawStatusSuccessCardNumberRow) []*record.WithdrawRecordMonthStatusSuccess {
	var WithdrawRecords []*record.WithdrawRecordMonthStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordMonthStatusSuccessCardNumber(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordYearStatusSuccessCardNumber maps a database row representing yearly
// successful withdraw statistics, filtered by card number, to a
// WithdrawRecordYearStatusSuccess domain model.
//
// Args:
//   - s: A pointer to a GetYearlyWithdrawStatusSuccessCardNumberRow representing
//     the database row.
//
// Returns:
//   - A pointer to a WithdrawRecordYearStatusSuccess containing the mapped data,
//     including Year, TotalSuccess, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordYearStatusSuccessCardNumber(s *db.GetYearlyWithdrawStatusSuccessCardNumberRow) *record.WithdrawRecordYearStatusSuccess {
	return &record.WithdrawRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToWithdrawRecordsYearStatusSuccessCardNumber maps a slice of
// GetYearlyWithdrawStatusSuccessCardNumberRow database rows to a slice of
// WithdrawRecordYearStatusSuccess domain models, filtered by card number.
//
// Args:
//   - Withdraws: A slice of pointers to GetYearlyWithdrawStatusSuccessCardNumberRow
//     representing the database rows with yearly successful withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawRecordYearStatusSuccess, each containing
//     the mapped data including Year, TotalSuccess, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordsYearStatusSuccessCardNumber(Withdraws []*db.GetYearlyWithdrawStatusSuccessCardNumberRow) []*record.WithdrawRecordYearStatusSuccess {
	var WithdrawRecords []*record.WithdrawRecordYearStatusSuccess

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordYearStatusSuccessCardNumber(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordMonthStatusFailedCardNumber maps a GetMonthWithdrawStatusFailedCardNumberRow database row
// to a WithdrawRecordMonthStatusFailed domain model, filtered by card number.
//
// Args:
//   - s: A pointer to a GetMonthWithdrawStatusFailedCardNumberRow representing the
//     database row.
//
// Returns:
//   - A pointer to a WithdrawRecordMonthStatusFailed containing the mapped data,
//     including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordMonthStatusFailedCardNumber(s *db.GetMonthWithdrawStatusFailedCardNumberRow) *record.WithdrawRecordMonthStatusFailed {
	return &record.WithdrawRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToWithdrawRecordsMonthStatusFailedCardNumber maps a slice of
// GetMonthWithdrawStatusFailedCardNumberRow database rows to a slice of
// WithdrawRecordMonthStatusFailed domain models, filtered by card number.
//
// Args:
//   - Withdraws: A slice of pointers to GetMonthWithdrawStatusFailedCardNumberRow
//     representing the database rows with monthly failed withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawRecordMonthStatusFailed, each containing
//     the mapped data including Year, Month, TotalFailed, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordsMonthStatusFailedCardNumber(Withdraws []*db.GetMonthWithdrawStatusFailedCardNumberRow) []*record.WithdrawRecordMonthStatusFailed {
	var WithdrawRecords []*record.WithdrawRecordMonthStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordMonthStatusFailedCardNumber(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawRecordYearStatusFailedCardNumber maps a single
// GetYearlyWithdrawStatusFailedCardNumberRow database row to a
// WithdrawRecordYearStatusFailed domain model, filtered by card number.
//
// Args:
//   - s: A pointer to a GetYearlyWithdrawStatusFailedCardNumberRow
//     representing the database row with yearly failed withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A pointer to a WithdrawRecordYearStatusFailed containing the mapped
//     data, including Year, TotalFailed, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordYearStatusFailedCardNumber(s *db.GetYearlyWithdrawStatusFailedCardNumberRow) *record.WithdrawRecordYearStatusFailed {
	return &record.WithdrawRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToWithdrawRecordsYearStatusFailedCardNumber maps a slice of
// GetYearlyWithdrawStatusFailedCardNumberRow database rows to a slice of
// WithdrawRecordYearStatusFailed domain models, filtered by card number.
//
// Args:
//   - Withdraws: A slice of pointers to GetYearlyWithdrawStatusFailedCardNumberRow
//     representing the database rows with yearly failed withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawRecordYearStatusFailed, each containing
//     the mapped data including Year, TotalFailed, and TotalAmount.
func (t *withdrawRecordMapper) ToWithdrawRecordsYearStatusFailedCardNumber(Withdraws []*db.GetYearlyWithdrawStatusFailedCardNumberRow) []*record.WithdrawRecordYearStatusFailed {
	var WithdrawRecords []*record.WithdrawRecordYearStatusFailed

	for _, Withdraw := range Withdraws {
		WithdrawRecords = append(WithdrawRecords, t.ToWithdrawRecordYearStatusFailedCardNumber(Withdraw))
	}

	return WithdrawRecords
}

// ToWithdrawAmountMonthly maps a single database row to a WithdrawMonthlyAmount
// domain model. It is intended for use with database rows that contain monthly
// withdraw statistics.
//
// Args:
//   - ss: A pointer to a GetMonthlyWithdrawsRow representing the database row
//     with monthly withdraw statistics.
//
// Returns:
//   - A pointer to a WithdrawMonthlyAmount containing the mapped data,
//     including Month and TotalAmount.
func (r *withdrawRecordMapper) ToWithdrawAmountMonthly(ss *db.GetMonthlyWithdrawsRow) *record.WithdrawMonthlyAmount {
	return &record.WithdrawMonthlyAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalWithdrawAmount),
	}
}

// ToWithdrawsAmountMonthly maps a slice of GetMonthlyWithdrawsRow database rows
// to a slice of WithdrawMonthlyAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetMonthlyWithdrawsRow representing the database
//     rows with monthly withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawMonthlyAmount, each containing the mapped
//     data including Month and TotalAmount.
func (s *withdrawRecordMapper) ToWithdrawsAmountMonthly(ss []*db.GetMonthlyWithdrawsRow) []*record.WithdrawMonthlyAmount {
	var withdrawRecords []*record.WithdrawMonthlyAmount

	for _, withdraw := range ss {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawAmountMonthly(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawAmountYearly maps a single database row to a WithdrawYearlyAmount
// domain model. It is intended for use with database rows that contain yearly
// withdraw statistics.
//
// Args:
//   - ss: A pointer to a GetYearlyWithdrawsRow representing the database row
//     with yearly withdraw statistics.
//
// Returns:
//   - A pointer to a WithdrawYearlyAmount containing the mapped data,
//     including Year and TotalAmount.
func (r *withdrawRecordMapper) ToWithdrawAmountYearly(ss *db.GetYearlyWithdrawsRow) *record.WithdrawYearlyAmount {
	return &record.WithdrawYearlyAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalWithdrawAmount),
	}
}

// ToWithdrawsAmountYearly maps a slice of GetYearlyWithdrawsRow database rows to a slice
// of WithdrawYearlyAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetYearlyWithdrawsRow representing the database rows
//     with yearly withdraw statistics.
//
// Returns:
//   - A slice of pointers to WithdrawYearlyAmount, each containing the mapped
//     data including Year and TotalAmount.
func (s *withdrawRecordMapper) ToWithdrawsAmountYearly(ss []*db.GetYearlyWithdrawsRow) []*record.WithdrawYearlyAmount {
	var withdrawRecords []*record.WithdrawYearlyAmount

	for _, withdraw := range ss {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawAmountYearly(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawAmountMonthlyByCardNumber maps a database row representing monthly
// withdraw statistics, filtered by card number, to a WithdrawMonthlyAmount
// domain model.
//
// Args:
//   - ss: A pointer to a GetMonthlyWithdrawsByCardNumberRow representing the
//     database row with monthly withdraw statistics filtered by card number.
//
// Returns:
//   - A pointer to a WithdrawMonthlyAmount containing the mapped data,
//     including Month and TotalAmount.
func (r *withdrawRecordMapper) ToWithdrawAmountMonthlyByCardNumber(ss *db.GetMonthlyWithdrawsByCardNumberRow) *record.WithdrawMonthlyAmount {
	return &record.WithdrawMonthlyAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalWithdrawAmount),
	}
}

// ToWithdrawsAmountMonthlyByCardNumber maps a slice of
// GetMonthlyWithdrawsByCardNumberRow database rows to a slice of
// WithdrawMonthlyAmount domain models, filtered by card number.
//
// Args:
//   - ss: A slice of pointers to GetMonthlyWithdrawsByCardNumberRow
//     representing the database rows with monthly withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawMonthlyAmount, each containing the
//     mapped data including Month and TotalAmount.
func (s *withdrawRecordMapper) ToWithdrawsAmountMonthlyByCardNumber(ss []*db.GetMonthlyWithdrawsByCardNumberRow) []*record.WithdrawMonthlyAmount {
	var withdrawRecords []*record.WithdrawMonthlyAmount

	for _, withdraw := range ss {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawAmountMonthlyByCardNumber(withdraw))
	}

	return withdrawRecords
}

// ToWithdrawAmountYearlyByCardNumber maps a database row representing yearly
// withdraw statistics, filtered by card number, to a WithdrawYearlyAmount
// domain model.
//
// Args:
//   - ss: A pointer to a GetYearlyWithdrawsByCardNumberRow representing the
//     database row with yearly withdraw statistics filtered by card number.
//
// Returns:
//   - A pointer to a WithdrawYearlyAmount containing the mapped data,
//     including Year and TotalAmount.
func (r *withdrawRecordMapper) ToWithdrawAmountYearlyByCardNumber(ss *db.GetYearlyWithdrawsByCardNumberRow) *record.WithdrawYearlyAmount {
	return &record.WithdrawYearlyAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalWithdrawAmount),
	}
}

// ToWithdrawsAmountYearlyByCardNumber maps a slice of
// GetYearlyWithdrawsByCardNumberRow database rows to a slice of
// WithdrawYearlyAmount domain models, filtered by card number.
//
// Args:
//   - ss: A slice of pointers to GetYearlyWithdrawsByCardNumberRow
//     representing the database rows with yearly withdraw statistics
//     filtered by card number.
//
// Returns:
//   - A slice of pointers to WithdrawYearlyAmount, each containing the
//     mapped data including Year and TotalAmount.
func (s *withdrawRecordMapper) ToWithdrawsAmountYearlyByCardNumber(ss []*db.GetYearlyWithdrawsByCardNumberRow) []*record.WithdrawYearlyAmount {
	var withdrawRecords []*record.WithdrawYearlyAmount

	for _, withdraw := range ss {
		withdrawRecords = append(withdrawRecords, s.ToWithdrawAmountYearlyByCardNumber(withdraw))
	}

	return withdrawRecords
}
