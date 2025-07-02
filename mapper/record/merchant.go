package recordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

// merchantRecordMapper provides methods to map Merchant database rows to MerchantRecord domain models.
type merchantRecordMapper struct {
}

// NewMerchantRecordMapper returns a new instance of merchantRecordMapper which provides methods to map Merchant database rows to MerchantRecord domain models.
func NewMerchantRecordMapper() *merchantRecordMapper {
	return &merchantRecordMapper{}
}

// ToMerchantRecord maps a Merchant to a MerchantRecord domain model.
//
// Args:
//   - merchant: A pointer to a Merchant representing the database row.
//
// Returns:
//   - A pointer to a MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantRecord(merchant *db.Merchant) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

// ToMerchantsRecord maps a slice of Merchant database rows to a slice of MerchantRecord domain models.
//
// Args:
//   - merchants: A slice of Merchant pointers representing the database rows.
//
// Returns:
//   - A slice of MerchantRecord pointers containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantsRecord(merchants []*db.Merchant) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantRecord(merchant))
	}
	return records
}

// ToMerchantTransactionRecord maps a FindAllTransactionsRow to a MerchantTransactionsRecord domain model.
//
// Args:
//   - merchant: A pointer to a FindAllTransactionsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantTransactionRecord(merchant *db.FindAllTransactionsRow) *record.MerchantTransactionsRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantTransactionsRecord{
		TransactionID:   merchant.TransactionID,
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:       merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:       deletedAt,
	}
}

// ToMerchantsTransactionRecord maps multiple FindAllTransactionsRow to multiple
// MerchantTransactionsRecord domain models.
//
// Args:
//   - merchants: A slice of pointers to FindAllTransactionsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantsTransactionRecord(merchants []*db.FindAllTransactionsRow) []*record.MerchantTransactionsRecord {
	var records []*record.MerchantTransactionsRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTransactionRecord(merchant))
	}
	return records
}

// ToMerchantGetAllRecord maps a GetMerchantsRow to a MerchantRecord domain model.
//
// Args:
//   - merchant: A pointer to a GetMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantGetAllRecord(merchant *db.GetMerchantsRow) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

// ToMerchantsGetAllRecord maps a slice of GetMerchantsRow to a slice of
// MerchantRecord domain models.
//
// Args:
//   - merchants: A slice of pointers to GetMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantsGetAllRecord(merchants []*db.GetMerchantsRow) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantGetAllRecord(merchant))
	}
	return records
}

// ToMerchantMonthlyPaymentMethod maps a GetMonthlyPaymentMethodsMerchantRow to a
// MerchantMonthlyPaymentMethod domain model.
//
// Args:
//   - ms: A pointer to a GetMonthlyPaymentMethodsMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethod(ms *db.GetMonthlyPaymentMethodsMerchantRow) *record.MerchantMonthlyPaymentMethod {
	return &record.MerchantMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyPaymentMethods maps a slice of GetMonthlyPaymentMethodsMerchantRow
// to a slice of MerchantMonthlyPaymentMethod domain models.
//
// Args:
//   - ms: A slice of pointers to GetMonthlyPaymentMethodsMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethods(ms []*db.GetMonthlyPaymentMethodsMerchantRow) []*record.MerchantMonthlyPaymentMethod {
	var records []*record.MerchantMonthlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyPaymentMethod(merchant))
	}
	return records
}

// ToMerchantYearlyPaymentMethod maps a GetYearlyPaymentMethodMerchantRow to a
// MerchantYearlyPaymentMethod domain model.
//
// Args:
//   - ms: A pointer to a GetYearlyPaymentMethodMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethod(ms *db.GetYearlyPaymentMethodMerchantRow) *record.MerchantYearlyPaymentMethod {
	return &record.MerchantYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantYearlyPaymentMethods maps a slice of GetYearlyPaymentMethodMerchantRow to a slice
// of MerchantYearlyPaymentMethod domain models.
//
// Args:
//   - ms: A slice of pointers to GetYearlyPaymentMethodMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethods(ms []*db.GetYearlyPaymentMethodMerchantRow) []*record.MerchantYearlyPaymentMethod {
	var records []*record.MerchantYearlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyPaymentMethod(merchant))
	}
	return records
}

// ToMerchantMonthlyAmount maps a GetMonthlyAmountMerchantRow to a MerchantMonthlyAmount domain model.
//
// Args:
//   - ms: A pointer to a GetMonthlyAmountMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyAmount(ms *db.GetMonthlyAmountMerchantRow) *record.MerchantMonthlyAmount {
	return &record.MerchantMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyAmounts maps a slice of GetMonthlyAmountMerchantRow to a slice of MerchantMonthlyAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetMonthlyAmountMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyAmounts(ms []*db.GetMonthlyAmountMerchantRow) []*record.MerchantMonthlyAmount {
	var records []*record.MerchantMonthlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyAmount(merchant))
	}
	return records
}

// ToMerchantYearlyAmount maps a GetYearlyAmountMerchantRow to a MerchantYearlyAmount domain model.
//
// Args:
//   - ms: A pointer to a GetYearlyAmountMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyAmount containing the mapped data, including Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyAmount(ms *db.GetYearlyAmountMerchantRow) *record.MerchantYearlyAmount {
	return &record.MerchantYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyAmounts maps a slice of GetYearlyAmountMerchantRow to a slice
// of MerchantYearlyAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetYearlyAmountMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyAmounts(ms []*db.GetYearlyAmountMerchantRow) []*record.MerchantYearlyAmount {
	var records []*record.MerchantYearlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyAmount(merchant))
	}
	return records
}

// ToMerchantMonthlyTotalAmount maps a GetMonthlyTotalAmountMerchantRow to a
// MerchantMonthlyTotalAmount domain model.
//
// Args:
//   - ms: A pointer to a GetMonthlyTotalAmountMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmount(ms *db.GetMonthlyTotalAmountMerchantRow) *record.MerchantMonthlyTotalAmount {
	return &record.MerchantMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyTotalAmounts maps a slice of GetMonthlyTotalAmountMerchantRow to a slice
// of MerchantMonthlyTotalAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetMonthlyTotalAmountMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmounts(ms []*db.GetMonthlyTotalAmountMerchantRow) []*record.MerchantMonthlyTotalAmount {
	var records []*record.MerchantMonthlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyTotalAmount(merchant))
	}
	return records
}

// ToMerchantYearlyTotalAmount maps a GetYearlyTotalAmountMerchantRow to a
// MerchantYearlyTotalAmount domain model.
//
// Args:
//   - ms: A pointer to a GetYearlyTotalAmountMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmount(ms *db.GetYearlyTotalAmountMerchantRow) *record.MerchantYearlyTotalAmount {
	return &record.MerchantYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyTotalAmounts maps a slice of GetYearlyTotalAmountMerchantRow to a slice
// of MerchantYearlyTotalAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetYearlyTotalAmountMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmounts(ms []*db.GetYearlyTotalAmountMerchantRow) []*record.MerchantYearlyTotalAmount {
	var records []*record.MerchantYearlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyTotalAmount(merchant))
	}
	return records
}

// ToMerchantTransactionByMerchantRecord maps a FindAllTransactionsByMerchantRow to a MerchantTransactionsRecord domain model.
//
// Args:
//   - merchant: A pointer to a FindAllTransactionsByMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantTransactionByMerchantRecord(merchant *db.FindAllTransactionsByMerchantRow) *record.MerchantTransactionsRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantTransactionsRecord{
		TransactionID:   merchant.TransactionID,
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:       merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:       deletedAt,
	}
}

// ToMerchantsTransactionByMerchantRecord maps a slice of FindAllTransactionsByMerchantRow to a slice
// of MerchantTransactionsRecord domain models.
//
// Args:
//   - merchants: A slice of pointers to FindAllTransactionsByMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantsTransactionByMerchantRecord(merchants []*db.FindAllTransactionsByMerchantRow) []*record.MerchantTransactionsRecord {
	var records []*record.MerchantTransactionsRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTransactionByMerchantRecord(merchant))
	}
	return records
}

// ToMerchantMonthlyPaymentMethodByMerchant maps a GetMonthlyPaymentMethodByMerchantsRow to a
// MerchantMonthlyPaymentMethod domain model.
//
// Args:
//   - ms: A pointer to a GetMonthlyPaymentMethodByMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethodByMerchant(ms *db.GetMonthlyPaymentMethodByMerchantsRow) *record.MerchantMonthlyPaymentMethod {
	return &record.MerchantMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyPaymentMethodsByMerchant maps a slice of GetMonthlyPaymentMethodByMerchantsRow
// to a slice of MerchantMonthlyPaymentMethod domain models.
//
// Args:
//   - ms: A slice of pointers to GetMonthlyPaymentMethodByMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethodsByMerchant(ms []*db.GetMonthlyPaymentMethodByMerchantsRow) []*record.MerchantMonthlyPaymentMethod {
	var records []*record.MerchantMonthlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyPaymentMethodByMerchant(merchant))
	}
	return records
}

// ToMerchantYearlyPaymentMethodByMerchant maps a GetYearlyPaymentMethodByMerchantsRow to a
// MerchantYearlyPaymentMethod domain model.
//
// Args:
//   - ms: A pointer to a GetYearlyPaymentMethodByMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethodByMerchant(ms *db.GetYearlyPaymentMethodByMerchantsRow) *record.MerchantYearlyPaymentMethod {
	return &record.MerchantYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantYearlyPaymentMethodsByMerchant maps a slice of GetYearlyPaymentMethodByMerchantsRow
// to a slice of MerchantYearlyPaymentMethod domain models.
//
// Args:
//   - ms: A slice of pointers to GetYearlyPaymentMethodByMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.

func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethodsByMerchant(ms []*db.GetYearlyPaymentMethodByMerchantsRow) []*record.MerchantYearlyPaymentMethod {
	var records []*record.MerchantYearlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyPaymentMethodByMerchant(merchant))
	}
	return records
}

// ToMerchantMonthlyAmountByMerchant maps a GetMonthlyAmountByMerchantsRow to a
// MerchantMonthlyAmount domain model.
//
// Args:
//   - ms: A pointer to a GetMonthlyAmountByMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyAmountByMerchant(ms *db.GetMonthlyAmountByMerchantsRow) *record.MerchantMonthlyAmount {
	return &record.MerchantMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyAmountsByMerchant maps a slice of GetMonthlyAmountByMerchantsRow
// to a slice of MerchantMonthlyAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetMonthlyAmountByMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including
//     Month and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyAmountsByMerchant(ms []*db.GetMonthlyAmountByMerchantsRow) []*record.MerchantMonthlyAmount {
	var records []*record.MerchantMonthlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyAmountByMerchant(merchant))
	}
	return records
}

// ToMerchantYearlyAmountByMerchant maps a GetYearlyAmountByMerchantsRow to a
// MerchantYearlyAmount domain model.
//
// Args:
//   - ms: A pointer to a GetYearlyAmountByMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyAmountByMerchant(ms *db.GetYearlyAmountByMerchantsRow) *record.MerchantYearlyAmount {
	return &record.MerchantYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyAmountsByMerchant maps a slice of GetYearlyAmountByMerchantsRow
// to a slice of MerchantYearlyAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetYearlyAmountByMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyAmountsByMerchant(ms []*db.GetYearlyAmountByMerchantsRow) []*record.MerchantYearlyAmount {
	var records []*record.MerchantYearlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyAmountByMerchant(merchant))
	}
	return records
}

// ToMerchantMonthlyTotalAmountByMerchant maps a GetMonthlyTotalAmountByMerchantRow to a
// MerchantMonthlyTotalAmount domain model.
//
// Args:
//   - ms: A pointer to a GetMonthlyTotalAmountByMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmountByMerchant(ms *db.GetMonthlyTotalAmountByMerchantRow) *record.MerchantMonthlyTotalAmount {
	return &record.MerchantMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyTotalAmountsByMerchant maps a slice of GetMonthlyTotalAmountByMerchantRow to a
// slice of MerchantMonthlyTotalAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetMonthlyTotalAmountByMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmountsByMerchant(ms []*db.GetMonthlyTotalAmountByMerchantRow) []*record.MerchantMonthlyTotalAmount {
	var records []*record.MerchantMonthlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyTotalAmountByMerchant(merchant))
	}
	return records
}

// ToMerchantYearlyTotalAmountByMerchant maps a GetYearlyTotalAmountByMerchantRow to a
// MerchantYearlyTotalAmount domain model.
//
// Args:
//   - ms: A pointer to a GetYearlyTotalAmountByMerchantRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmountByMerchant(ms *db.GetYearlyTotalAmountByMerchantRow) *record.MerchantYearlyTotalAmount {
	return &record.MerchantYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyTotalAmountsByMerchant maps a slice of GetYearlyTotalAmountByMerchantRow to a
// slice of MerchantYearlyTotalAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetYearlyTotalAmountByMerchantRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmountsByMerchant(ms []*db.GetYearlyTotalAmountByMerchantRow) []*record.MerchantYearlyTotalAmount {
	var records []*record.MerchantYearlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyTotalAmountByMerchant(merchant))
	}
	return records
}

// ToMerchantTransactionByApikeyRecord maps a FindAllTransactionsByApikeyRow to a MerchantTransactionsRecord domain model.
//
// Args:
//   - merchant: A pointer to a FindAllTransactionsByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantTransactionByApikeyRecord(merchant *db.FindAllTransactionsByApikeyRow) *record.MerchantTransactionsRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantTransactionsRecord{
		TransactionID:   merchant.TransactionID,
		CardNumber:      merchant.CardNumber,
		Amount:          merchant.Amount,
		PaymentMethod:   merchant.PaymentMethod,
		MerchantID:      merchant.MerchantID,
		MerchantName:    merchant.MerchantName,
		TransactionTime: merchant.TransactionTime,
		CreatedAt:       merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:       merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:       deletedAt,
	}
}

// ToMerchantsTransactionByApikeyRecord maps a slice of FindAllTransactionsByApikeyRow to a slice of
// MerchantTransactionsRecord domain models.
//
// Args:
//   - merchants: A slice of pointers to FindAllTransactionsByApikeyRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantTransactionsRecord containing the mapped data, including
//     TransactionID, CardNumber, Amount, PaymentMethod, MerchantID, MerchantName,
//     TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantsTransactionByApikeyRecord(merchants []*db.FindAllTransactionsByApikeyRow) []*record.MerchantTransactionsRecord {
	var records []*record.MerchantTransactionsRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTransactionByApikeyRecord(merchant))
	}
	return records
}

// ToMerchantMonthlyPaymentMethodByApikey maps a GetMonthlyPaymentMethodByApikeyRow to a
// MerchantMonthlyPaymentMethod domain model.
//
// Args:
//   - ms: A pointer to a GetMonthlyPaymentMethodByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethodByApikey(ms *db.GetMonthlyPaymentMethodByApikeyRow) *record.MerchantMonthlyPaymentMethod {
	return &record.MerchantMonthlyPaymentMethod{
		Month:         ms.Month,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyPaymentMethodsByApikey maps a slice of GetMonthlyPaymentMethodByApikeyRow to a slice
// of MerchantMonthlyPaymentMethod domain models.
//
// Args:
//   - ms: A slice of pointers to GetMonthlyPaymentMethodByApikeyRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyPaymentMethod containing the mapped data, including
//     Month, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyPaymentMethodsByApikey(ms []*db.GetMonthlyPaymentMethodByApikeyRow) []*record.MerchantMonthlyPaymentMethod {
	var records []*record.MerchantMonthlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyPaymentMethodByApikey(merchant))
	}
	return records
}

// ToMerchantYearlyPaymentMethodByApikey maps a GetYearlyPaymentMethodByApikeyRow to a
// MerchantYearlyPaymentMethod domain model.
//
// Args:
//   - ms: A pointer to a GetYearlyPaymentMethodByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethodByApikey(ms *db.GetYearlyPaymentMethodByApikeyRow) *record.MerchantYearlyPaymentMethod {
	return &record.MerchantYearlyPaymentMethod{
		Year:          ms.Year,
		PaymentMethod: ms.PaymentMethod,
		TotalAmount:   int(ms.TotalAmount),
	}
}

// ToMerchantYearlyPaymentMethodsByApikey maps a slice of GetYearlyPaymentMethodByApikeyRow
// to a slice of MerchantYearlyPaymentMethod domain models.
//
// Args:
//   - ms: A slice of pointers to GetYearlyPaymentMethodByApikeyRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyPaymentMethod containing the mapped data, including
//     Year, PaymentMethod, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyPaymentMethodsByApikey(ms []*db.GetYearlyPaymentMethodByApikeyRow) []*record.MerchantYearlyPaymentMethod {
	var records []*record.MerchantYearlyPaymentMethod
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyPaymentMethodByApikey(merchant))
	}
	return records
}

// ToMerchantMonthlyAmountByApikey maps a GetMonthlyAmountByApikeyRow to a
// MerchantMonthlyAmount domain model.
//
// Args:
//   - ms: A pointer to a GetMonthlyAmountByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyAmount containing the mapped data, including
//     Month, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyAmountByApikey(ms *db.GetMonthlyAmountByApikeyRow) *record.MerchantMonthlyAmount {
	return &record.MerchantMonthlyAmount{
		Month:       ms.Month,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyAmountsByApikey maps a slice of GetMonthlyAmountByApikeyRow to a slice of
// MerchantMonthlyAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetMonthlyAmountByApikeyRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyAmount containing the mapped data, including Month and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyAmountsByApikey(ms []*db.GetMonthlyAmountByApikeyRow) []*record.MerchantMonthlyAmount {
	var records []*record.MerchantMonthlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyAmountByApikey(merchant))
	}
	return records
}

// ToMerchantYearlyAmountByApikey maps a GetYearlyAmountByApikeyRow to a
// MerchantYearlyAmount domain model.
//
// Args:
//   - ms: A pointer to a GetYearlyAmountByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyAmountByApikey(ms *db.GetYearlyAmountByApikeyRow) *record.MerchantYearlyAmount {
	return &record.MerchantYearlyAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyAmountsByApikey maps a slice of GetYearlyAmountByApikeyRow to a slice
// of MerchantYearlyAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetYearlyAmountByApikeyRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyAmountsByApikey(ms []*db.GetYearlyAmountByApikeyRow) []*record.MerchantYearlyAmount {
	var records []*record.MerchantYearlyAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyAmountByApikey(merchant))
	}
	return records
}

// ToMerchantMonthlyTotalAmountByApikey maps a GetMonthlyTotalAmountByApikeyRow to a
// MerchantMonthlyTotalAmount domain model.
//
// Args:
//   - ms: A pointer to a GetMonthlyTotalAmountByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantMonthlyTotalAmount containing the mapped data, including
//     Month, Year, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmountByApikey(ms *db.GetMonthlyTotalAmountByApikeyRow) *record.MerchantMonthlyTotalAmount {
	return &record.MerchantMonthlyTotalAmount{
		Month:       ms.Month,
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantMonthlyTotalAmountsByApikey maps a slice of GetMonthlyTotalAmountByApikeyRow
// to a slice of MerchantMonthlyTotalAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetMonthlyTotalAmountByApikeyRow representing the
//     database rows.
//
// Returns:
//   - A slice of pointers to MerchantMonthlyTotalAmount containing the mapped data,
//     including Month, Year, and TotalAmount.
func (m *merchantRecordMapper) ToMerchantMonthlyTotalAmountsByApikey(ms []*db.GetMonthlyTotalAmountByApikeyRow) []*record.MerchantMonthlyTotalAmount {
	var records []*record.MerchantMonthlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantMonthlyTotalAmountByApikey(merchant))
	}
	return records
}

// ToMerchantYearlyTotalAmountByApikey maps a GetYearlyTotalAmountByApikeyRow to a
// MerchantYearlyTotalAmount domain model.
//
// Args:
//   - ms: A pointer to a GetYearlyTotalAmountByApikeyRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantYearlyTotalAmount containing the mapped data, including
//     Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmountByApikey(ms *db.GetYearlyTotalAmountByApikeyRow) *record.MerchantYearlyTotalAmount {
	return &record.MerchantYearlyTotalAmount{
		Year:        ms.Year,
		TotalAmount: int(ms.TotalAmount),
	}
}

// ToMerchantYearlyTotalAmountsByApikey maps a slice of GetYearlyTotalAmountByApikeyRow
// to a slice of MerchantYearlyTotalAmount domain models.
//
// Args:
//   - ms: A slice of pointers to GetYearlyTotalAmountByApikeyRow representing the
//     database rows.
//
// Returns:
//   - A slice of pointers to MerchantYearlyTotalAmount containing the mapped data,
//     including Year and TotalAmount.
func (m *merchantRecordMapper) ToMerchantYearlyTotalAmountsByApikey(ms []*db.GetYearlyTotalAmountByApikeyRow) []*record.MerchantYearlyTotalAmount {
	var records []*record.MerchantYearlyTotalAmount
	for _, merchant := range ms {
		records = append(records, m.ToMerchantYearlyTotalAmountByApikey(merchant))
	}
	return records
}

//

// ToMerchantActiveRecord maps a GetActiveMerchantsRow to a MerchantRecord domain model.
//
// Args:
//   - merchant: A pointer to a GetActiveMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantActiveRecord(merchant *db.GetActiveMerchantsRow) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

// ToMerchantsActiveRecord maps a slice of GetActiveMerchantsRow to a slice of
// MerchantRecord domain models.
//
// Args:
//   - merchants: A slice of pointers to GetActiveMerchantsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantsActiveRecord(merchants []*db.GetActiveMerchantsRow) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantActiveRecord(merchant))
	}
	return records
}

// ToMerchantTrashedRecord maps a GetTrashedMerchantsRow to a MerchantRecord domain model.
//
// Args:
//   - merchant: A pointer to a GetTrashedMerchantsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantTrashedRecord(merchant *db.GetTrashedMerchantsRow) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}

// ToMerchantsTrashedRecord maps a slice of GetTrashedMerchantsRow to a slice of
// MerchantRecord domain models.
//
// Args:
//   - merchants: A slice of pointers to GetTrashedMerchantsRow representing the trashed database rows.
//
// Returns:
//   - A slice of pointers to MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantRecordMapper) ToMerchantsTrashedRecord(merchants []*db.GetTrashedMerchantsRow) []*record.MerchantRecord {
	var records []*record.MerchantRecord
	for _, merchant := range merchants {
		records = append(records, m.ToMerchantTrashedRecord(merchant))
	}
	return records
}
