package recordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

// saldoRecordMapper provides methods to map Saldo database rows to SaldoRecord domain models.
type saldoRecordMapper struct{}

// NewSaldoRecordMapper returns a new instance of saldoRecordMapper which provides methods to map Saldo database rows to SaldoRecord domain models.
func NewSaldoRecordMapper() *saldoRecordMapper {
	return &saldoRecordMapper{}
}

// ToSaldoRecord maps a Saldo database row to a SaldoRecord domain model.
//
// Args:
//   - saldo: A pointer to a Saldo representing the database row.
//
// Returns:
//   - A pointer to a SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoRecordMapper) ToSaldoRecord(saldo *db.Saldo) *record.SaldoRecord {
	var deletedAt *string

	if saldo.DeletedAt.Valid {
		formatedDeletedAt := saldo.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.SaldoRecord{
		ID:             int(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int(saldo.TotalBalance),
		WithdrawAmount: int(saldo.WithdrawAmount.Int32),
		WithdrawTime:   saldo.WithdrawTime.Time.Format("2006-01-02"),
		CreatedAt:      saldo.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:      saldo.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:      deletedAt,
	}
}

// ToSaldosRecord maps a slice of Saldo database rows to a slice of SaldoRecord domain models.
//
// Args:
//   - saldos: A slice of pointers to Saldo representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoRecordMapper) ToSaldosRecord(saldos []*db.Saldo) []*record.SaldoRecord {
	var saldoRecords []*record.SaldoRecord
	for _, saldo := range saldos {
		saldoRecords = append(saldoRecords, s.ToSaldoRecord(saldo))
	}
	return saldoRecords
}

// ToSaldoRecordAll maps a GetSaldosRow database row to a SaldoRecord domain model.
//
// Args:
//   - saldo: A pointer to a GetSaldosRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoRecordMapper) ToSaldoRecordAll(saldo *db.GetSaldosRow) *record.SaldoRecord {
	var deletedAt *string

	if saldo.DeletedAt.Valid {
		formatedDeletedAt := saldo.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.SaldoRecord{
		ID:             int(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int(saldo.TotalBalance),
		WithdrawAmount: int(saldo.WithdrawAmount.Int32),
		WithdrawTime:   saldo.WithdrawTime.Time.Format("2006-01-02"),
		CreatedAt:      saldo.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:      saldo.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:      deletedAt,
	}
}

// ToSaldosRecordAll maps a slice of GetSaldosRow database rows to a slice of SaldoRecord domain models.
//
// Args:
//   - saldos: A slice of pointers to GetSaldosRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoRecordMapper) ToSaldosRecordAll(saldos []*db.GetSaldosRow) []*record.SaldoRecord {
	var saldoRecords []*record.SaldoRecord
	for _, saldo := range saldos {
		saldoRecords = append(saldoRecords, s.ToSaldoRecordAll(saldo))
	}
	return saldoRecords
}

// ToSaldoRecordActive maps a GetActiveSaldosRow database row to a SaldoRecord domain model.
//
// Args:
//   - saldo: A pointer to a GetActiveSaldosRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoRecordMapper) ToSaldoRecordActive(saldo *db.GetActiveSaldosRow) *record.SaldoRecord {
	var deletedAt *string

	if saldo.DeletedAt.Valid {
		formatedDeletedAt := saldo.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.SaldoRecord{
		ID:             int(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int(saldo.TotalBalance),
		WithdrawAmount: int(saldo.WithdrawAmount.Int32),
		WithdrawTime:   saldo.WithdrawTime.Time.Format("2006-01-02"),
		CreatedAt:      saldo.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:      saldo.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:      deletedAt,
	}
}

// ToSaldosRecordActive maps a slice of GetActiveSaldosRow database rows to a slice of SaldoRecord domain models.
//
// Args:
//   - saldos: A slice of pointers to GetActiveSaldosRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoRecord containing the mapped data, including
//     ID, CardNumber, TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoRecordMapper) ToSaldosRecordActive(saldos []*db.GetActiveSaldosRow) []*record.SaldoRecord {
	var saldoRecords []*record.SaldoRecord
	for _, saldo := range saldos {
		saldoRecords = append(saldoRecords, s.ToSaldoRecordActive(saldo))
	}
	return saldoRecords
}

// ToSaldoRecordTrashed maps a GetTrashedSaldosRow database row to a SaldoRecord domain model.
//
// Args:
//   - saldo: A pointer to a GetTrashedSaldosRow representing the trashed saldo database row.
//
// Returns:
//   - A pointer to a SaldoRecord containing the mapped data, including ID, CardNumber,
//     TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoRecordMapper) ToSaldoRecordTrashed(saldo *db.GetTrashedSaldosRow) *record.SaldoRecord {
	var deletedAt *string

	if saldo.DeletedAt.Valid {
		formatedDeletedAt := saldo.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.SaldoRecord{
		ID:             int(saldo.SaldoID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int(saldo.TotalBalance),
		WithdrawAmount: int(saldo.WithdrawAmount.Int32),
		WithdrawTime:   saldo.WithdrawTime.Time.Format("2006-01-02"),
		CreatedAt:      saldo.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:      saldo.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:      deletedAt,
	}
}

// ToSaldosRecordTrashed maps a slice of GetTrashedSaldosRow database rows to a slice of SaldoRecord domain models.
//
// Args:
//   - saldos: A slice of pointers to GetTrashedSaldosRow representing the trashed saldo database rows.
//
// Returns:
//   - A slice of pointers to SaldoRecord containing the mapped data, including ID, CardNumber,
//     TotalBalance, WithdrawAmount, WithdrawTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *saldoRecordMapper) ToSaldosRecordTrashed(saldos []*db.GetTrashedSaldosRow) []*record.SaldoRecord {
	var saldoRecords []*record.SaldoRecord
	for _, saldo := range saldos {
		saldoRecords = append(saldoRecords, s.ToSaldoRecordTrashed(saldo))
	}
	return saldoRecords
}

// ToSaldoMonthTotalBalance maps a GetMonthlyTotalSaldoBalanceRow database row to a SaldoMonthTotalBalance domain model.
//
// Args:
//   - ss: A pointer to a GetMonthlyTotalSaldoBalanceRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoMonthTotalBalance containing the mapped data, including Month, Year, and TotalBalance.
func (s *saldoRecordMapper) ToSaldoMonthTotalBalance(ss *db.GetMonthlyTotalSaldoBalanceRow) *record.SaldoMonthTotalBalance {
	totalBalance := 0
	if ss.TotalBalance != 0 {
		totalBalance = int(ss.TotalBalance)
	}

	return &record.SaldoMonthTotalBalance{
		Month:        ss.Month,
		Year:         ss.Year,
		TotalBalance: totalBalance,
	}
}

// ToSaldoMonthTotalBalances maps a slice of GetMonthlyTotalSaldoBalanceRow database rows to a slice of SaldoMonthTotalBalance domain models.
//
// Args:
//   - ss: A slice of pointers to GetMonthlyTotalSaldoBalanceRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoMonthTotalBalance containing the mapped data, including Month, Year, and TotalBalance.
func (s *saldoRecordMapper) ToSaldoMonthTotalBalances(ss []*db.GetMonthlyTotalSaldoBalanceRow) []*record.SaldoMonthTotalBalance {
	var saldoRecords []*record.SaldoMonthTotalBalance
	for _, saldo := range ss {
		saldoRecords = append(saldoRecords, s.ToSaldoMonthTotalBalance(saldo))
	}
	return saldoRecords
}

// ToSaldoYearTotalBalance maps a GetYearlyTotalSaldoBalancesRow database row to a SaldoYearTotalBalance domain model.
//
// Args:
//   - ss: A pointer to a GetYearlyTotalSaldoBalancesRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoYearTotalBalance containing the mapped data, including Year, and TotalBalance.
func (s *saldoRecordMapper) ToSaldoYearTotalBalance(ss *db.GetYearlyTotalSaldoBalancesRow) *record.SaldoYearTotalBalance {
	totalBalance := 0
	if ss.TotalBalance != 0 {
		totalBalance = int(ss.TotalBalance)
	}

	return &record.SaldoYearTotalBalance{
		Year:         ss.Year,
		TotalBalance: int(totalBalance),
	}
}

// ToSaldoYearTotalBalances maps a slice of GetYearlyTotalSaldoBalancesRow database rows to a slice of SaldoYearTotalBalance domain models.
//
// Args:
//   - ss: A slice of pointers to GetYearlyTotalSaldoBalancesRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoYearTotalBalance containing the mapped data, including Year, and TotalBalance.
func (s *saldoRecordMapper) ToSaldoYearTotalBalances(ss []*db.GetYearlyTotalSaldoBalancesRow) []*record.SaldoYearTotalBalance {
	var saldoRecords []*record.SaldoYearTotalBalance
	for _, saldo := range ss {
		saldoRecords = append(saldoRecords, s.ToSaldoYearTotalBalance(saldo))
	}
	return saldoRecords
}

// ToSaldoMonthBalance maps a MonthlySaldoBalancesRow database row to a SaldoMonthSaldoBalance domain model.
//
// Args:
//   - ss: A pointer to a MonthlySaldoBalancesRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoMonthSaldoBalance containing the mapped data, including Month, and TotalBalance.
func (s *saldoRecordMapper) ToSaldoMonthBalance(ss *db.GetMonthlySaldoBalancesRow) *record.SaldoMonthSaldoBalance {
	return &record.SaldoMonthSaldoBalance{
		Month:        ss.Month,
		TotalBalance: int(ss.TotalBalance),
	}
}

// ToSaldoMonthBalances maps a slice of MonthlySaldoBalancesRow database rows to a slice of SaldoMonthSaldoBalance domain models.
//
// Args:
//   - ss: A slice of pointers to MonthlySaldoBalancesRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoMonthSaldoBalance containing the mapped data, including Month, and TotalBalance.
func (s *saldoRecordMapper) ToSaldoMonthBalances(ss []*db.GetMonthlySaldoBalancesRow) []*record.SaldoMonthSaldoBalance {
	var saldoRecords []*record.SaldoMonthSaldoBalance
	for _, saldo := range ss {
		saldoRecords = append(saldoRecords, s.ToSaldoMonthBalance(saldo))
	}
	return saldoRecords
}

// ToSaldoYearSaldoBalance maps a GetYearlySaldoBalancesRow database row to a SaldoYearSaldoBalance domain model.
//
// Args:
//   - ss: A pointer to a GetYearlySaldoBalancesRow representing the database row.
//
// Returns:
//   - A pointer to a SaldoYearSaldoBalance containing the mapped data, including Year and TotalBalance.
func (s *saldoRecordMapper) ToSaldoYearSaldoBalance(ss *db.GetYearlySaldoBalancesRow) *record.SaldoYearSaldoBalance {
	return &record.SaldoYearSaldoBalance{
		Year:         ss.Year,
		TotalBalance: int(ss.TotalBalance),
	}
}

// ToSaldoYearSaldoBalances maps a slice of GetYearlySaldoBalancesRow database rows to a slice of SaldoYearSaldoBalance domain models.
//
// Args:
//   - ss: A slice of pointers to GetYearlySaldoBalancesRow representing the database rows.
//
// Returns:
//   - A slice of pointers to SaldoYearSaldoBalance containing the mapped data, including Year and TotalBalance.
func (s *saldoRecordMapper) ToSaldoYearSaldoBalances(ss []*db.GetYearlySaldoBalancesRow) []*record.SaldoYearSaldoBalance {
	var saldoRecords []*record.SaldoYearSaldoBalance
	for _, saldo := range ss {
		saldoRecords = append(saldoRecords, s.ToSaldoYearSaldoBalance(saldo))
	}
	return saldoRecords
}
