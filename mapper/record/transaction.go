package recordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

// transactionRecordMapper is a struct that provides methods to map
// database rows to transaction domain model records.
type transactionRecordMapper struct{}

// NewTransactionRecordMapper returns a new instance of transactionRecordMapper.
func NewTransactionRecordMapper() *transactionRecordMapper {
	return &transactionRecordMapper{}
}

// ToTransactionRecord maps a Transaction database row to a TransactionRecord domain model.
//
// Args:
//   - transaction: A pointer to a Transaction representing the database row.
//
// Returns:
//   - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionRecordMapper) ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord {
	var deletedAt *string

	if transaction.DeletedAt.Valid {
		formatedDeletedAt := transaction.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.TransactionRecord{
		ID:              int(transaction.TransactionID),
		TransactionNo:   transaction.TransactionNo.String(),
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      int(transaction.MerchantID),
		TransactionTime: transaction.TransactionTime.Format("2006-01-02 15:04:05"),
		CreatedAt:       transaction.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:       transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:       deletedAt,
	}
}

// ToTransactionsRecord maps a slice of Transaction database rows to a slice of TransactionRecord domain models.
//
// Args:
//   - transactions: A slice of pointers to Transaction representing the database rows.
//
// Returns:
//   - A slice of TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionRecordMapper) ToTransactionsRecord(transactions []*db.Transaction) []*record.TransactionRecord {
	var transactionRecords []*record.TransactionRecord
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, s.ToTransactionRecord(transaction))
	}
	return transactionRecords
}

// ToTransactionByCardNumberRecord maps a GetTransactionsByCardNumberRow database row to a TransactionRecord domain model.
//
// Args:
//   - transaction: A pointer to a GetTransactionsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionRecordMapper) ToTransactionByCardNumberRecord(transaction *db.GetTransactionsByCardNumberRow) *record.TransactionRecord {
	var deletedAt *string

	if transaction.DeletedAt.Valid {
		formatedDeletedAt := transaction.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.TransactionRecord{
		ID:              int(transaction.TransactionID),
		TransactionNo:   transaction.TransactionNo.String(),
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      int(transaction.MerchantID),
		TransactionTime: transaction.TransactionTime.Format("2006-01-02 15:04:05"),
		CreatedAt:       transaction.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:       transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:       deletedAt,
	}
}

// ToTransactionsByCardNumberRecord maps a slice of GetTransactionsByCardNumberRow database rows
// to a slice of TransactionRecord domain models.
//
// Args:
//   - transactions: A slice of pointers to GetTransactionsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionRecordMapper) ToTransactionsByCardNumberRecord(transactions []*db.GetTransactionsByCardNumberRow) []*record.TransactionRecord {
	var transactionRecords []*record.TransactionRecord
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, s.ToTransactionByCardNumberRecord(transaction))
	}
	return transactionRecords
}

// ToTransactionRecordAll maps a GetTransactionsRow database row to a TransactionRecord domain model.
//
// Args:
//   - transaction: A pointer to a GetTransactionsRow representing the database row.
//
// Returns:
//   - A pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionRecordMapper) ToTransactionRecordAll(transaction *db.GetTransactionsRow) *record.TransactionRecord {
	var deletedAt *string

	if transaction.DeletedAt.Valid {
		formatedDeletedAt := transaction.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.TransactionRecord{
		ID:              int(transaction.TransactionID),
		TransactionNo:   transaction.TransactionNo.String(),
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      int(transaction.MerchantID),
		TransactionTime: transaction.TransactionTime.Format("2006-01-02 15:04:05"),
		CreatedAt:       transaction.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:       transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:       deletedAt,
	}
}

// ToTransactionsRecordAll maps a slice of GetTransactionsRow database rows to a slice of TransactionRecord domain models.
//
// Args:
//   - transactions: A slice of pointers to GetTransactionsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionRecordMapper) ToTransactionsRecordAll(transactions []*db.GetTransactionsRow) []*record.TransactionRecord {
	var transactionRecords []*record.TransactionRecord
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, s.ToTransactionRecordAll(transaction))
	}
	return transactionRecords
}

// ToTransactionRecordActive maps a GetActiveTransactionsRow database row to a TransactionRecord domain model.
// It is intended for use with database rows that contain active transaction records.
// It returns a pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
// CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionRecordMapper) ToTransactionRecordActive(transaction *db.GetActiveTransactionsRow) *record.TransactionRecord {
	var deletedAt *string

	if transaction.DeletedAt.Valid {
		formatedDeletedAt := transaction.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.TransactionRecord{
		ID:              int(transaction.TransactionID),
		TransactionNo:   transaction.TransactionNo.String(),
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      int(transaction.MerchantID),
		TransactionTime: transaction.TransactionTime.Format("2006-01-02 15:04:05"),
		CreatedAt:       transaction.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:       transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:       deletedAt,
	}
}

// ToTransactionsRecordActive maps a slice of GetActiveTransactionsRow database rows
// to a slice of TransactionRecord domain models.
//
// Args:
//   - transactions: A slice of pointers to GetActiveTransactionsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionRecordMapper) ToTransactionsRecordActive(transactions []*db.GetActiveTransactionsRow) []*record.TransactionRecord {
	var transactionRecords []*record.TransactionRecord
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, s.ToTransactionRecordActive(transaction))
	}
	return transactionRecords
}

// ToTransactionRecordTrashed maps a GetTrashedTransactionsRow database row to a TransactionRecord domain model.
// It is intended for use with database rows that contain trashed transaction records.
// It returns a pointer to a TransactionRecord containing the mapped data, including ID, TransactionNo,
// CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionRecordMapper) ToTransactionRecordTrashed(transaction *db.GetTrashedTransactionsRow) *record.TransactionRecord {
	var deletedAt *string

	if transaction.DeletedAt.Valid {
		formatedDeletedAt := transaction.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.TransactionRecord{
		ID:              int(transaction.TransactionID),
		TransactionNo:   transaction.TransactionNo.String(),
		CardNumber:      transaction.CardNumber,
		Amount:          int(transaction.Amount),
		PaymentMethod:   transaction.PaymentMethod,
		MerchantID:      int(transaction.MerchantID),
		TransactionTime: transaction.TransactionTime.Format("2006-01-02 15:04:05"),
		CreatedAt:       transaction.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt:       transaction.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
		DeletedAt:       deletedAt,
	}
}

// ToTransactionsRecordTrashed maps a slice of GetTrashedTransactionsRow database rows to a slice of
// TransactionRecord domain models.
//
// Args:
//   - transactions: A slice of pointers to GetTrashedTransactionsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecord containing the mapped data, including ID, TransactionNo,
//     CardNumber, Amount, PaymentMethod, MerchantID, TransactionTime, CreatedAt, UpdatedAt, and DeletedAt.
func (s *transactionRecordMapper) ToTransactionsRecordTrashed(transactions []*db.GetTrashedTransactionsRow) []*record.TransactionRecord {
	var transactionRecords []*record.TransactionRecord
	for _, transaction := range transactions {
		transactionRecords = append(transactionRecords, s.ToTransactionRecordTrashed(transaction))
	}
	return transactionRecords
}

// ToTransactionRecordMonthStatusSuccess maps a single GetMonthTransactionStatusSuccessRow database row
// to a TransactionRecordMonthStatusSuccess domain model.
//
// Args:
//   - s: A pointer to GetMonthTransactionStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordMonthStatusSuccess(s *db.GetMonthTransactionStatusSuccessRow) *record.TransactionRecordMonthStatusSuccess {
	return &record.TransactionRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransactionRecordsMonthStatusSuccess maps a slice of GetMonthTransactionStatusSuccessRow database rows to a slice of
// TransactionRecordMonthStatusSuccess domain models.
//
// Args:
//   - Transactions: A slice of pointers to GetMonthTransactionStatusSuccessRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordsMonthStatusSuccess(Transactions []*db.GetMonthTransactionStatusSuccessRow) []*record.TransactionRecordMonthStatusSuccess {
	var TransactionRecords []*record.TransactionRecordMonthStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordMonthStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordYearStatusSuccess maps a GetYearlyTransactionStatusSuccessRow database row
// to a TransactionRecordYearStatusSuccess domain model.
//
// Args:
//   - s: A pointer to GetYearlyTransactionStatusSuccessRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordYearStatusSuccess(s *db.GetYearlyTransactionStatusSuccessRow) *record.TransactionRecordYearStatusSuccess {
	return &record.TransactionRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransactionRecordsYearStatusSuccess maps a slice of GetYearlyTransactionStatusSuccessRow database rows
// to a slice of TransactionRecordYearStatusSuccess domain models.
//
// Args:
//   - Transactions: A slice of pointers to GetYearlyTransactionStatusSuccessRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordsYearStatusSuccess(Transactions []*db.GetYearlyTransactionStatusSuccessRow) []*record.TransactionRecordYearStatusSuccess {
	var TransactionRecords []*record.TransactionRecordYearStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordYearStatusSuccess(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordMonthStatusFailed maps a GetMonthTransactionStatusFailedRow database row
// to a TransactionRecordMonthStatusFailed domain model.
//
// Args:
//   - s: A pointer to GetMonthTransactionStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordMonthStatusFailed(s *db.GetMonthTransactionStatusFailedRow) *record.TransactionRecordMonthStatusFailed {
	return &record.TransactionRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionRecordsMonthStatusFailed maps a slice of GetMonthTransactionStatusFailedRow database rows
// to a slice of TransactionRecordMonthStatusFailed domain models.
//
// Args:
//   - Transactions: A slice of pointers to GetMonthTransactionStatusFailedRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordsMonthStatusFailed(Transactions []*db.GetMonthTransactionStatusFailedRow) []*record.TransactionRecordMonthStatusFailed {
	var TransactionRecords []*record.TransactionRecordMonthStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordMonthStatusFailed(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordYearStatusFailed maps a GetYearlyTransactionStatusFailedRow database row
// to a TransactionRecordYearStatusFailed domain model.
//
// Args:
//   - s: A pointer to GetYearlyTransactionStatusFailedRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordYearStatusFailed(s *db.GetYearlyTransactionStatusFailedRow) *record.TransactionRecordYearStatusFailed {
	return &record.TransactionRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionRecordsYearStatusFailed maps a slice of GetYearlyTransactionStatusFailedRow database rows
// to a slice of TransactionRecordYearStatusFailed domain models.
//
// Args:
//   - Transactions: A slice of pointers to GetYearlyTransactionStatusFailedRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordsYearStatusFailed(Transactions []*db.GetYearlyTransactionStatusFailedRow) []*record.TransactionRecordYearStatusFailed {
	var TransactionRecords []*record.TransactionRecordYearStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordYearStatusFailed(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordMonthStatusSuccessCardNumber maps a GetMonthTransactionStatusSuccessCardNumberRow database row
// to a TransactionRecordMonthStatusSuccess domain model.
//
// Args:
//   - s: A pointer to GetMonthTransactionStatusSuccessCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordMonthStatusSuccessCardNumber(s *db.GetMonthTransactionStatusSuccessCardNumberRow) *record.TransactionRecordMonthStatusSuccess {
	return &record.TransactionRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransactionRecordsMonthStatusSuccessCardNumber maps a slice of GetMonthTransactionStatusSuccessCardNumberRow database rows
// to a slice of TransactionRecordMonthStatusSuccess domain models.
//
// Args:
//   - Transactions: A slice of pointers to GetMonthTransactionStatusSuccessCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordMonthStatusSuccess containing the mapped data, including Year,
//     Month, TotalSuccess, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordsMonthStatusSuccessCardNumber(Transactions []*db.GetMonthTransactionStatusSuccessCardNumberRow) []*record.TransactionRecordMonthStatusSuccess {
	var TransactionRecords []*record.TransactionRecordMonthStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordMonthStatusSuccessCardNumber(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordYearStatusSuccessCardNumber maps a GetYearlyTransactionStatusSuccessCardNumberRow database row
// to a TransactionRecordYearStatusSuccess domain model.
//
// Args:
//   - s: A pointer to GetYearlyTransactionStatusSuccessCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordYearStatusSuccessCardNumber(s *db.GetYearlyTransactionStatusSuccessCardNumberRow) *record.TransactionRecordYearStatusSuccess {
	return &record.TransactionRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTransactionRecordsYearStatusSuccessCardNumber maps a slice of GetYearlyTransactionStatusSuccessCardNumberRow database rows
// to a slice of TransactionRecordYearStatusSuccess domain models.
//
// Args:
//   - Transactions: A slice of pointers to GetYearlyTransactionStatusSuccessCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordYearStatusSuccess containing the mapped data, including Year,
//     TotalSuccess, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordsYearStatusSuccessCardNumber(Transactions []*db.GetYearlyTransactionStatusSuccessCardNumberRow) []*record.TransactionRecordYearStatusSuccess {
	var TransactionRecords []*record.TransactionRecordYearStatusSuccess

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordYearStatusSuccessCardNumber(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordMonthStatusFailedCardNumber maps a GetMonthTransactionStatusFailedCardNumberRow
// database row to a TransactionRecordMonthStatusFailed domain model.
//
// Args:
//   - s: A pointer to GetMonthTransactionStatusFailedCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordMonthStatusFailedCardNumber(s *db.GetMonthTransactionStatusFailedCardNumberRow) *record.TransactionRecordMonthStatusFailed {
	return &record.TransactionRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionRecordsMonthStatusFailedCardNumber maps a slice of GetMonthTransactionStatusFailedCardNumberRow
// database rows to a slice of TransactionRecordMonthStatusFailed domain models.
//
// Args:
//   - Transactions: A slice of pointers to GetMonthTransactionStatusFailedCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordMonthStatusFailed containing the mapped data, including Year,
//     Month, TotalFailed, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordsMonthStatusFailedCardNumber(Transactions []*db.GetMonthTransactionStatusFailedCardNumberRow) []*record.TransactionRecordMonthStatusFailed {
	var TransactionRecords []*record.TransactionRecordMonthStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordMonthStatusFailedCardNumber(Transaction))
	}

	return TransactionRecords
}

// ToTransactionRecordYearStatusFailedCardNumber maps a GetYearlyTransactionStatusFailedCardNumberRow
// database row to a TransactionRecordYearStatusFailed domain model.
//
// Args:
//   - s: A pointer to GetYearlyTransactionStatusFailedCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordYearStatusFailedCardNumber(s *db.GetYearlyTransactionStatusFailedCardNumberRow) *record.TransactionRecordYearStatusFailed {
	return &record.TransactionRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTransactionRecordsYearStatusFailedCardNumber maps a slice of GetYearlyTransactionStatusFailedCardNumberRow
// database rows to a slice of TransactionRecordYearStatusFailed domain models.
//
// Args:
//   - Transactions: A slice of pointers to GetYearlyTransactionStatusFailedCardNumberRow representing the
//     database rows.
//
// Returns:
//   - A slice of pointers to TransactionRecordYearStatusFailed containing the mapped data, including Year,
//     TotalFailed, and TotalAmount.
func (t *transactionRecordMapper) ToTransactionRecordsYearStatusFailedCardNumber(Transactions []*db.GetYearlyTransactionStatusFailedCardNumberRow) []*record.TransactionRecordYearStatusFailed {
	var TransactionRecords []*record.TransactionRecordYearStatusFailed

	for _, Transaction := range Transactions {
		TransactionRecords = append(TransactionRecords, t.ToTransactionRecordYearStatusFailedCardNumber(Transaction))
	}

	return TransactionRecords
}

// ToTransactionMonthlyMethod maps a GetMonthlyPaymentMethodsRow database row to a TransactionMonthMethod
// domain model.
//
// Args:
//   - s: A pointer to GetMonthlyPaymentMethodsRow representing the database row.
//
// Returns:
//   - A pointer to TransactionMonthMethod containing the mapped data, including Month,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionRecordMapper) ToTransactionMonthlyMethod(ss *db.GetMonthlyPaymentMethodsRow) *record.TransactionMonthMethod {
	return &record.TransactionMonthMethod{
		Month:             ss.Month,
		PaymentMethod:     ss.PaymentMethod,
		TotalTransactions: int(ss.TotalTransactions),
		TotalAmount:       int(ss.TotalAmount),
	}
}

// ToTransactionMonthlyMethods maps a slice of GetMonthlyPaymentMethodsRow database rows
// to a slice of TransactionMonthMethod domain models.
//
// Args:
//   - ss: A slice of pointers to GetMonthlyPaymentMethodsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionMonthMethod containing the mapped data, including Month,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionRecordMapper) ToTransactionMonthlyMethods(ss []*db.GetMonthlyPaymentMethodsRow) []*record.TransactionMonthMethod {
	var transactionRecords []*record.TransactionMonthMethod
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionMonthlyMethod(transaction))
	}
	return transactionRecords
}

// ToTransactionYearlyMethod maps a GetYearlyPaymentMethodsRow database row to a TransactionYearMethod
// domain model.
//
// Args:
//   - s: A pointer to GetYearlyPaymentMethodsRow representing the database row.
//
// Returns:
//   - A pointer to TransactionYearMethod containing the mapped data, including Year,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionRecordMapper) ToTransactionYearlyMethod(ss *db.GetYearlyPaymentMethodsRow) *record.TransactionYearMethod {
	return &record.TransactionYearMethod{
		Year:              ss.Year,
		PaymentMethod:     ss.PaymentMethod,
		TotalTransactions: int(ss.TotalTransactions),
		TotalAmount:       int(ss.TotalAmount),
	}
}

// ToTransactionYearlyMethods maps a slice of GetYearlyPaymentMethodsRow database rows
// to a slice of TransactionYearMethod domain models.
//
// Args:
//   - ss: A slice of pointers to GetYearlyPaymentMethodsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionYearMethod containing the mapped data, including Year,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionRecordMapper) ToTransactionYearlyMethods(ss []*db.GetYearlyPaymentMethodsRow) []*record.TransactionYearMethod {
	var transactionRecords []*record.TransactionYearMethod
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionYearlyMethod(transaction))
	}
	return transactionRecords
}

//

// ToTransactionMonthlyAmount maps a GetMonthlyAmountsRow database row to a TransactionMonthAmount
// domain model.
//
// Args:
//   - ss: A pointer to GetMonthlyAmountsRow representing the database row.
//
// Returns:
//   - A pointer to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transactionRecordMapper) ToTransactionMonthlyAmount(ss *db.GetMonthlyAmountsRow) *record.TransactionMonthAmount {
	return &record.TransactionMonthAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalAmount),
	}
}

// ToTransactionMonthlyAmounts maps a slice of GetMonthlyAmountsRow database rows
// to a slice of TransactionMonthAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetMonthlyAmountsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transactionRecordMapper) ToTransactionMonthlyAmounts(ss []*db.GetMonthlyAmountsRow) []*record.TransactionMonthAmount {
	var transactionRecords []*record.TransactionMonthAmount
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionMonthlyAmount(transaction))
	}
	return transactionRecords
}

// ToTransactionYearlyAmount maps a GetYearlyAmountsRow database row to a TransactionYearlyAmount
// domain model.
//
// Args:
//   - ss: A pointer to GetYearlyAmountsRow representing the database row.
//
// Returns:
//   - A pointer to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
func (s *transactionRecordMapper) ToTransactionYearlyAmount(ss *db.GetYearlyAmountsRow) *record.TransactionYearlyAmount {
	return &record.TransactionYearlyAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalAmount),
	}
}

// ToTransactionYearlyAmounts maps a slice of GetYearlyAmountsRow database rows
// to a slice of TransactionYearlyAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetYearlyAmountsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
func (s *transactionRecordMapper) ToTransactionYearlyAmounts(ss []*db.GetYearlyAmountsRow) []*record.TransactionYearlyAmount {
	var transactionRecords []*record.TransactionYearlyAmount
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionYearlyAmount(transaction))
	}
	return transactionRecords
}

/////

// ToTransactionMonthlyMethodByCardNumber maps a GetMonthlyPaymentMethodsByCardNumberRow database row to a TransactionMonthMethod
// domain model.
//
// Args:
//   - ss: A pointer to GetMonthlyPaymentMethodsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionMonthMethod containing the mapped data, including Month,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionRecordMapper) ToTransactionMonthlyMethodByCardNumber(ss *db.GetMonthlyPaymentMethodsByCardNumberRow) *record.TransactionMonthMethod {
	return &record.TransactionMonthMethod{
		Month:             ss.Month,
		PaymentMethod:     ss.PaymentMethod,
		TotalTransactions: int(ss.TotalTransactions),
		TotalAmount:       int(ss.TotalAmount),
	}
}

// ToTransactionMonthlyMethodsByCardNumber maps a slice of GetMonthlyPaymentMethodsByCardNumberRow database rows
// to a slice of TransactionMonthMethod domain models.
//
// Args:
//   - ss: A slice of pointers to GetMonthlyPaymentMethodsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionMonthMethod containing the mapped data, including Month,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionRecordMapper) ToTransactionMonthlyMethodsByCardNumber(ss []*db.GetMonthlyPaymentMethodsByCardNumberRow) []*record.TransactionMonthMethod {
	var transactionRecords []*record.TransactionMonthMethod
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionMonthlyMethodByCardNumber(transaction))
	}
	return transactionRecords
}

// ToTransactionYearlyMethodByCardNumber maps a GetYearlyPaymentMethodsByCardNumberRow database row to a TransactionYearMethod
// domain model.
//
// Args:
//   - ss: A pointer to GetYearlyPaymentMethodsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionYearMethod containing the mapped data, including Year,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionRecordMapper) ToTransactionYearlyMethodByCardNumber(ss *db.GetYearlyPaymentMethodsByCardNumberRow) *record.TransactionYearMethod {
	return &record.TransactionYearMethod{
		Year:              ss.Year,
		PaymentMethod:     ss.PaymentMethod,
		TotalTransactions: int(ss.TotalTransactions),
		TotalAmount:       int(ss.TotalAmount),
	}
}

// ToTransactionYearlyMethodsByCardNumber maps a slice of GetYearlyPaymentMethodsByCardNumberRow database rows
// to a slice of TransactionYearMethod domain models.
//
// Args:
//   - ss: A slice of pointers to GetYearlyPaymentMethodsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionYearMethod containing the mapped data, including Year,
//     PaymentMethod, TotalTransactions, and TotalAmount.
func (s *transactionRecordMapper) ToTransactionYearlyMethodsByCardNumber(ss []*db.GetYearlyPaymentMethodsByCardNumberRow) []*record.TransactionYearMethod {
	var transactionRecords []*record.TransactionYearMethod
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionYearlyMethodByCardNumber(transaction))
	}
	return transactionRecords
}

//

// ToTransactionMonthlyAmountByCardNumber maps a GetMonthlyAmountsByCardNumberRow database row
// to a TransactionMonthAmount domain model.
//
// Args:
//   - ss: A pointer to GetMonthlyAmountsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transactionRecordMapper) ToTransactionMonthlyAmountByCardNumber(ss *db.GetMonthlyAmountsByCardNumberRow) *record.TransactionMonthAmount {
	return &record.TransactionMonthAmount{
		Month:       ss.Month,
		TotalAmount: int(ss.TotalAmount),
	}
}

// ToTransactionMonthlyAmountsByCardNumber maps a slice of GetMonthlyAmountsByCardNumberRow database rows
// to a slice of TransactionMonthAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetMonthlyAmountsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *transactionRecordMapper) ToTransactionMonthlyAmountsByCardNumber(ss []*db.GetMonthlyAmountsByCardNumberRow) []*record.TransactionMonthAmount {
	var transactionRecords []*record.TransactionMonthAmount
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionMonthlyAmountByCardNumber(transaction))
	}
	return transactionRecords

}

// ToTransactionYearlyAmountByCardNumber maps a GetYearlyAmountsByCardNumberRow database row
// to a TransactionYearlyAmount domain model.
//
// Args:
//   - ss: A pointer to GetYearlyAmountsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
func (s *transactionRecordMapper) ToTransactionYearlyAmountByCardNumber(ss *db.GetYearlyAmountsByCardNumberRow) *record.TransactionYearlyAmount {
	return &record.TransactionYearlyAmount{
		Year:        ss.Year,
		TotalAmount: int(ss.TotalAmount),
	}
}

// ToTransactionYearlyAmountsByCardNumber maps a slice of GetYearlyAmountsByCardNumberRow database rows
// to a slice of TransactionYearlyAmount domain models.
//
// Args:
//   - ss: A slice of pointers to GetYearlyAmountsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TransactionYearlyAmount containing the mapped data, including Year and TotalAmount.
func (s *transactionRecordMapper) ToTransactionYearlyAmountsByCardNumber(ss []*db.GetYearlyAmountsByCardNumberRow) []*record.TransactionYearlyAmount {
	var transactionRecords []*record.TransactionYearlyAmount
	for _, transaction := range ss {
		transactionRecords = append(transactionRecords, s.ToTransactionYearlyAmountByCardNumber(transaction))
	}
	return transactionRecords
}
