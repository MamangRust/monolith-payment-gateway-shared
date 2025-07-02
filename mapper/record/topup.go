package recordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

type topupRecordMapper struct {
}

// NewTopupRecordMapper creates a new instance of topupRecordMapper.
// It returns a pointer to the newly created topupRecordMapper.

func NewTopupRecordMapper() *topupRecordMapper {
	return &topupRecordMapper{}
}

// ToTopupRecord converts a db.Topup to a record.TopupRecord.
// It takes a pointer to a db.Topup as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.Topup to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.Topup is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupRecordMapper) ToTopupRecord(topup *db.Topup) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

// ToTopupRecords converts a slice of db.Topup to a slice of record.TopupRecord.
// It iterates over the provided slice of db.Topup, converting each element
// using the ToTopupRecord method and appending the result to a new slice.
// The function returns a slice of pointers to record.TopupRecord.
func (t *topupRecordMapper) ToTopupRecords(topups []*db.Topup) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecord(topup))
	}

	return topupRecords
}

// ToTopupByCardNumberRecord converts a db.GetTopupsByCardNumberRow to a record.TopupRecord.
// It takes a pointer to a db.GetTopupsByCardNumberRow as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.GetTopupsByCardNumberRow to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.GetTopupsByCardNumberRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupRecordMapper) ToTopupByCardNumberRecord(topup *db.GetTopupsByCardNumberRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

// ToTopupByCardNumberRecords converts a slice of db.GetTopupsByCardNumberRow to a slice of record.TopupRecord.
// It takes a slice of pointers to db.GetTopupsByCardNumberRow as a parameter and returns a slice of pointers to record.TopupRecord.
// The function iterates over the provided slice of db.GetTopupsByCardNumberRow, converting each element
// using the ToTopupByCardNumberRecord method and appending the result to a new slice.
// The function returns a slice of pointers to record.TopupRecord.
func (t *topupRecordMapper) ToTopupByCardNumberRecords(topups []*db.GetTopupsByCardNumberRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupByCardNumberRecord(topup))
	}

	return topupRecords
}

// ToTopupRecordAll converts a db.GetTopupsRow to a record.TopupRecord.
// It takes a pointer to a db.GetTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.GetTopupsRow to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.GetTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupRecordMapper) ToTopupRecordAll(topup *db.GetTopupsRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

// ToTopupRecordsAll converts a slice of db.GetTopupsRow to a slice of record.TopupRecord.
// It iterates over the provided slice of db.GetTopupsRow, converting each element
// using the ToTopupRecordAll method and appending the result to a new slice.
// The function returns a slice of pointers to record.TopupRecord.
func (t *topupRecordMapper) ToTopupRecordsAll(topups []*db.GetTopupsRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordAll(topup))
	}

	return topupRecords
}

// ToTopupRecordActive converts a db.GetActiveTopupsRow to a record.TopupRecord.
// It takes a pointer to a db.GetActiveTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.GetActiveTopupsRow to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.GetActiveTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupRecordMapper) ToTopupRecordActive(topup *db.GetActiveTopupsRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

// ToTopupRecordsActive maps a slice of GetActiveTopupsRow database rows to a slice of TopupRecord domain models.
// It iterates over the provided slice of GetActiveTopupsRow, converting each element
// using the ToTopupRecordActive method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecord.
func (t *topupRecordMapper) ToTopupRecordsActive(topups []*db.GetActiveTopupsRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordActive(topup))
	}

	return topupRecords
}

// ToTopupRecordTrashed converts a db.GetTrashedTopupsRow to a record.TopupRecord.
// It takes a pointer to a db.GetTrashedTopupsRow as a parameter and returns a pointer to a record.TopupRecord.
// The function maps the fields of the db.GetTrashedTopupsRow to the corresponding fields of the record.TopupRecord.
// If the DeletedAt field of the db.GetTrashedTopupsRow is not valid, the function returns nil for the DeletedAt field of the record.TopupRecord.
func (t *topupRecordMapper) ToTopupRecordTrashed(topup *db.GetTrashedTopupsRow) *record.TopupRecord {
	var deleted_at *string

	if topup.DeletedAt.Valid {
		formatedDeletedAt := topup.DeletedAt.Time.Format("2006-01-02")
		deleted_at = &formatedDeletedAt
	}

	return &record.TopupRecord{
		ID:          int(topup.TopupID),
		CardNumber:  topup.CardNumber,
		TopupNo:     topup.TopupNo.String(),
		TopupAmount: int(topup.TopupAmount),
		TopupMethod: topup.TopupMethod,
		TopupTime:   topup.TopupTime.Format("2006-01-02 15:04:05.000"),
		CreatedAt:   topup.CreatedAt.Time.Format("2006-01-02 15:04:05.000"),
		UpdatedAt:   topup.UpdatedAt.Time.Format("2006-01-02 15:04:05.000"),
		DeletedAt:   deleted_at,
	}
}

// ToTopupRecordsTrashed maps a slice of GetTrashedTopupsRow database rows to a slice of TopupRecord domain models.
// It iterates over the provided slice of GetTrashedTopupsRow, converting each element
// using the ToTopupRecordTrashed method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecord.
func (t *topupRecordMapper) ToTopupRecordsTrashed(topups []*db.GetTrashedTopupsRow) []*record.TopupRecord {
	var topupRecords []*record.TopupRecord

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordTrashed(topup))
	}

	return topupRecords
}

// ToTopupRecordMonthStatusSuccess maps a GetMonthTopupStatusSuccessRow database row to a TopupRecordMonthStatusSuccess domain model.
// It takes a pointer to a GetMonthTopupStatusSuccessRow as a parameter and returns a pointer to a TopupRecordMonthStatusSuccess.
// The function maps the fields of the GetMonthTopupStatusSuccessRow to the corresponding fields of the TopupRecordMonthStatusSuccess.
func (t *topupRecordMapper) ToTopupRecordMonthStatusSuccess(s *db.GetMonthTopupStatusSuccessRow) *record.TopupRecordMonthStatusSuccess {
	return &record.TopupRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTopupRecordsMonthStatusSuccess maps a slice of GetMonthTopupStatusSuccessRow database rows to a slice of TopupRecordMonthStatusSuccess domain models.
// It iterates over the provided slice of GetMonthTopupStatusSuccessRow, converting each element
// using the ToTopupRecordMonthStatusSuccess method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordMonthStatusSuccess.
func (t *topupRecordMapper) ToTopupRecordsMonthStatusSuccess(topups []*db.GetMonthTopupStatusSuccessRow) []*record.TopupRecordMonthStatusSuccess {
	var topupRecords []*record.TopupRecordMonthStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordMonthStatusSuccess(topup))
	}

	return topupRecords
}

// ToTopupRecordYearStatusSuccess maps a GetYearlyTopupStatusSuccessRow database row to a TopupRecordYearStatusSuccess domain model.
// It takes a pointer to a GetYearlyTopupStatusSuccessRow as a parameter and returns a pointer to a TopupRecordYearStatusSuccess.
// The function maps the fields of the GetYearlyTopupStatusSuccessRow to the corresponding fields of the TopupRecordYearStatusSuccess.
func (t *topupRecordMapper) ToTopupRecordYearStatusSuccess(s *db.GetYearlyTopupStatusSuccessRow) *record.TopupRecordYearStatusSuccess {
	return &record.TopupRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTopupRecordsYearStatusSuccess maps a slice of GetYearlyTopupStatusSuccessRow database rows to a slice of TopupRecordYearStatusSuccess domain models.
// It iterates over the provided slice of GetYearlyTopupStatusSuccessRow, converting each element
// using the ToTopupRecordYearStatusSuccess method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordYearStatusSuccess.
func (t *topupRecordMapper) ToTopupRecordsYearStatusSuccess(topups []*db.GetYearlyTopupStatusSuccessRow) []*record.TopupRecordYearStatusSuccess {
	var topupRecords []*record.TopupRecordYearStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordYearStatusSuccess(topup))
	}

	return topupRecords
}

// ToTopupRecordMonthStatusFailed maps a GetMonthTopupStatusFailedRow database row
// to a TopupRecordMonthStatusFailed domain model. It takes a pointer to a
// GetMonthTopupStatusFailedRow as a parameter and returns a pointer to a
// TopupRecordMonthStatusFailed. The function maps the fields of the
// GetMonthTopupStatusFailedRow to the corresponding fields of the
// TopupRecordMonthStatusFailed.
func (t *topupRecordMapper) ToTopupRecordMonthStatusFailed(s *db.GetMonthTopupStatusFailedRow) *record.TopupRecordMonthStatusFailed {
	return &record.TopupRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupRecordsMonthStatusFailed maps a slice of GetMonthTopupStatusFailedRow database rows to a slice of TopupRecordMonthStatusFailed domain models.
// It iterates over the provided slice of GetMonthTopupStatusFailedRow, converting each element
// using the ToTopupRecordMonthStatusFailed method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordMonthStatusFailed.
func (t *topupRecordMapper) ToTopupRecordsMonthStatusFailed(topups []*db.GetMonthTopupStatusFailedRow) []*record.TopupRecordMonthStatusFailed {
	var topupRecords []*record.TopupRecordMonthStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordMonthStatusFailed(topup))
	}

	return topupRecords
}

// ToTopupRecordYearStatusFailed maps a GetYearlyTopupStatusFailedRow database row to a TopupRecordYearStatusFailed domain model.
// It takes a pointer to a GetYearlyTopupStatusFailedRow as a parameter and returns a pointer to a TopupRecordYearStatusFailed.
// The function maps the fields of the GetYearlyTopupStatusFailedRow to the corresponding fields of the TopupRecordYearStatusFailed.
func (t *topupRecordMapper) ToTopupRecordYearStatusFailed(s *db.GetYearlyTopupStatusFailedRow) *record.TopupRecordYearStatusFailed {
	return &record.TopupRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupRecordsYearStatusFailed maps a slice of GetYearlyTopupStatusFailedRow database rows to a slice of TopupRecordYearStatusFailed domain models.
// It iterates over the provided slice of GetYearlyTopupStatusFailedRow, converting each element
// using the ToTopupRecordYearStatusFailed method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordYearStatusFailed.
func (t *topupRecordMapper) ToTopupRecordsYearStatusFailed(topups []*db.GetYearlyTopupStatusFailedRow) []*record.TopupRecordYearStatusFailed {
	var topupRecords []*record.TopupRecordYearStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordYearStatusFailed(topup))
	}

	return topupRecords
}

// ToTopupRecordMonthStatusSuccessByCardNumber maps a GetMonthTopupStatusSuccessCardNumberRow database row to a TopupRecordMonthStatusSuccess domain model.
// It takes a pointer to a GetMonthTopupStatusSuccessCardNumberRow as a parameter and returns a pointer to a TopupRecordMonthStatusSuccess.
// The function maps the fields of the GetMonthTopupStatusSuccessCardNumberRow to the corresponding fields of the TopupRecordMonthStatusSuccess.
func (t *topupRecordMapper) ToTopupRecordMonthStatusSuccessByCardNumber(s *db.GetMonthTopupStatusSuccessCardNumberRow) *record.TopupRecordMonthStatusSuccess {
	return &record.TopupRecordMonthStatusSuccess{
		Year:         s.Year,
		Month:        s.Month,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTopupRecordsMonthStatusSuccessByCardNumber maps a slice of GetMonthTopupStatusSuccessCardNumberRow database rows to a slice of TopupRecordMonthStatusSuccess domain models.
// It iterates over the provided slice of GetMonthTopupStatusSuccessCardNumberRow, converting each element
// using the ToTopupRecordMonthStatusSuccessByCardNumber method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordMonthStatusSuccess.
func (t *topupRecordMapper) ToTopupRecordsMonthStatusSuccessByCardNumber(topups []*db.GetMonthTopupStatusSuccessCardNumberRow) []*record.TopupRecordMonthStatusSuccess {
	var topupRecords []*record.TopupRecordMonthStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordMonthStatusSuccessByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupRecordYearStatusSuccessByCardNumber maps a GetYearlyTopupStatusSuccessCardNumberRow database row
// to a TopupRecordYearStatusSuccess domain model. It takes a pointer to a
// GetYearlyTopupStatusSuccessCardNumberRow as a parameter and returns a pointer
// to a TopupRecordYearStatusSuccess. The function maps the fields of
// GetYearlyTopupStatusSuccessCardNumberRow to the corresponding fields of
// TopupRecordYearStatusSuccess.
func (t *topupRecordMapper) ToTopupRecordYearStatusSuccessByCardNumber(s *db.GetYearlyTopupStatusSuccessCardNumberRow) *record.TopupRecordYearStatusSuccess {
	return &record.TopupRecordYearStatusSuccess{
		Year:         s.Year,
		TotalSuccess: int(s.TotalSuccess),
		TotalAmount:  int(s.TotalAmount),
	}
}

// ToTopupRecordsYearStatusSuccessByCardNumber maps a slice of GetYearlyTopupStatusSuccessCardNumberRow database rows to a slice of TopupRecordYearStatusSuccess domain models.
// It iterates over the provided slice of GetYearlyTopupStatusSuccessCardNumberRow, converting each element
// using the ToTopupRecordYearStatusSuccessByCardNumber method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordYearStatusSuccess.
func (t *topupRecordMapper) ToTopupRecordsYearStatusSuccessByCardNumber(topups []*db.GetYearlyTopupStatusSuccessCardNumberRow) []*record.TopupRecordYearStatusSuccess {
	var topupRecords []*record.TopupRecordYearStatusSuccess

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordYearStatusSuccessByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupRecordMonthStatusFailedByCardNumber maps a GetMonthTopupStatusFailedCardNumberRow database row
// to a TopupRecordMonthStatusFailed domain model. It takes a pointer to a
// GetMonthTopupStatusFailedCardNumberRow as a parameter and returns a pointer to a
// TopupRecordMonthStatusFailed. The function maps the fields of the
// GetMonthTopupStatusFailedCardNumberRow to the corresponding fields of the
// TopupRecordMonthStatusFailed.
func (t *topupRecordMapper) ToTopupRecordMonthStatusFailedByCardNumber(s *db.GetMonthTopupStatusFailedCardNumberRow) *record.TopupRecordMonthStatusFailed {
	return &record.TopupRecordMonthStatusFailed{
		Year:        s.Year,
		Month:       s.Month,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupRecordsMonthStatusFailedByCardNumber maps a slice of GetMonthTopupStatusFailedCardNumberRow database rows to a slice of TopupRecordMonthStatusFailed domain models.
// It iterates over the provided slice of GetMonthTopupStatusFailedCardNumberRow, converting each element
// using the ToTopupRecordMonthStatusFailedByCardNumber method and appending the result to a new slice.
// The function returns a slice of pointers to TopupRecordMonthStatusFailed.
func (t *topupRecordMapper) ToTopupRecordsMonthStatusFailedByCardNumber(topups []*db.GetMonthTopupStatusFailedCardNumberRow) []*record.TopupRecordMonthStatusFailed {
	var topupRecords []*record.TopupRecordMonthStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordMonthStatusFailedByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupRecordYearStatusFailedByCardNumber maps a GetYearlyTopupStatusFailedCardNumberRow database row
// to a TopupRecordYearStatusFailed domain model. It takes a pointer to a
// GetYearlyTopupStatusFailedCardNumberRow as a parameter and returns a pointer
// to a TopupRecordYearStatusFailed. The function maps the fields of
// GetYearlyTopupStatusFailedCardNumberRow to the corresponding fields of
// TopupRecordYearStatusFailed.
func (t *topupRecordMapper) ToTopupRecordYearStatusFailedByCardNumber(s *db.GetYearlyTopupStatusFailedCardNumberRow) *record.TopupRecordYearStatusFailed {
	return &record.TopupRecordYearStatusFailed{
		Year:        s.Year,
		TotalFailed: int(s.TotalFailed),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupRecordsYearStatusFailedByCardNumber maps a slice of GetYearlyTopupStatusFailedCardNumberRow database rows
// to a slice of TopupRecordYearStatusFailed domain models. It iterates over the provided slice, converting each element
// using the ToTopupRecordYearStatusFailedByCardNumber method, and appends the result to a new slice.
// The function returns a slice of pointers to TopupRecordYearStatusFailed.
func (t *topupRecordMapper) ToTopupRecordsYearStatusFailedByCardNumber(topups []*db.GetYearlyTopupStatusFailedCardNumberRow) []*record.TopupRecordYearStatusFailed {
	var topupRecords []*record.TopupRecordYearStatusFailed

	for _, topup := range topups {
		topupRecords = append(topupRecords, t.ToTopupRecordYearStatusFailedByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupMonthlyMethod maps a GetMonthlyTopupMethodsRow database row to a TopupMonthMethod domain model.
// It takes a pointer to a GetMonthlyTopupMethodsRow as a parameter and returns a pointer to a TopupMonthMethod.
// The function maps the fields of GetMonthlyTopupMethodsRow to the corresponding fields of TopupMonthMethod.
func (t *topupRecordMapper) ToTopupMonthlyMethod(s *db.GetMonthlyTopupMethodsRow) *record.TopupMonthMethod {
	return &record.TopupMonthMethod{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyMethods maps a slice of GetMonthlyTopupMethodsRow database rows to a slice of TopupMonthMethod domain models.
// It iterates over the provided slice, converting each element using the ToTopupMonthlyMethod method,
// and appends the result to a new slice. The function returns a slice of pointers to TopupMonthMethod.
func (t *topupRecordMapper) ToTopupMonthlyMethods(s []*db.GetMonthlyTopupMethodsRow) []*record.TopupMonthMethod {
	var topupRecords []*record.TopupMonthMethod

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupMonthlyMethod(topup))
	}

	return topupRecords
}

// ToTopupYearlyMethod maps a GetYearlyTopupMethodsRow database row to a TopupYearlyMethod domain model.
// It takes a pointer to a GetYearlyTopupMethodsRow as a parameter and returns a pointer to a TopupYearlyMethod.
// The function maps the fields of GetYearlyTopupMethodsRow to the corresponding fields of TopupYearlyMethod.
func (t *topupRecordMapper) ToTopupYearlyMethod(s *db.GetYearlyTopupMethodsRow) *record.TopupYearlyMethod {
	return &record.TopupYearlyMethod{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyMethods maps a slice of GetYearlyTopupMethodsRow database rows to a slice of TopupYearlyMethod domain models.
// It iterates over the provided slice, converting each element using the ToTopupYearlyMethod method,
// and appends the result to a new slice. The function returns a slice of pointers to TopupYearlyMethod.
func (t *topupRecordMapper) ToTopupYearlyMethods(s []*db.GetYearlyTopupMethodsRow) []*record.TopupYearlyMethod {
	var topupRecords []*record.TopupYearlyMethod

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupYearlyMethod(topup))
	}

	return topupRecords
}

// ToTopupMonthlyAmount maps a GetMonthlyTopupAmountsRow database row to a TopupMonthAmount domain model.
// It takes a pointer to a GetMonthlyTopupAmountsRow as a parameter and returns a pointer to a TopupMonthAmount.
// The function maps the fields of GetMonthlyTopupAmountsRow to the corresponding fields of TopupMonthAmount.
func (t *topupRecordMapper) ToTopupMonthlyAmount(s *db.GetMonthlyTopupAmountsRow) *record.TopupMonthAmount {
	return &record.TopupMonthAmount{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyAmounts maps a slice of GetMonthlyTopupAmountsRow database rows to a slice of TopupMonthAmount domain models.
// It iterates over the provided slice, converting each element using the ToTopupMonthlyAmount method,
// and appends the result to a new slice. The function returns a slice of pointers to TopupMonthAmount.
func (t *topupRecordMapper) ToTopupMonthlyAmounts(s []*db.GetMonthlyTopupAmountsRow) []*record.TopupMonthAmount {
	var topupRecords []*record.TopupMonthAmount

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupMonthlyAmount(topup))
	}

	return topupRecords
}

// ToTopupYearlyAmount maps a GetYearlyTopupAmountsRow database row to a TopupYearlyAmount domain model.
// It takes a pointer to a GetYearlyTopupAmountsRow as a parameter and returns a pointer to a TopupYearlyAmount.
// The function maps the fields of GetYearlyTopupAmountsRow to the corresponding fields of TopupYearlyAmount.
func (t *topupRecordMapper) ToTopupYearlyAmount(s *db.GetYearlyTopupAmountsRow) *record.TopupYearlyAmount {
	return &record.TopupYearlyAmount{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyAmounts maps a slice of GetYearlyTopupAmountsRow database rows to a slice of TopupYearlyAmount domain models.
// It iterates over the provided slice, converting each element using the ToTopupYearlyAmount method,
// and appends the result to a new slice. The function returns a slice of pointers to TopupYearlyAmount.
func (t *topupRecordMapper) ToTopupYearlyAmounts(s []*db.GetYearlyTopupAmountsRow) []*record.TopupYearlyAmount {
	var topupRecords []*record.TopupYearlyAmount

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupYearlyAmount(topup))
	}

	return topupRecords
}

//

// ToTopupMonthlyMethodByCardNumber maps a GetMonthlyTopupMethodsByCardNumberRow database row
// to a TopupMonthMethod domain model. It takes a pointer to a
// GetMonthlyTopupMethodsByCardNumberRow as a parameter and returns a pointer to a
// TopupMonthMethod. The function maps the fields of GetMonthlyTopupMethodsByCardNumberRow
// to the corresponding fields of TopupMonthMethod.
func (t *topupRecordMapper) ToTopupMonthlyMethodByCardNumber(s *db.GetMonthlyTopupMethodsByCardNumberRow) *record.TopupMonthMethod {
	return &record.TopupMonthMethod{
		Month:       s.Month,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyMethodsByCardNumber maps a slice of GetMonthlyTopupMethodsByCardNumberRow database rows
// to a slice of TopupMonthMethod domain models. It iterates over the provided slice, converting
// each element using the ToTopupMonthlyMethodByCardNumber method, and appends the result to a new
// slice. The function returns a slice of pointers to TopupMonthMethod.
func (t *topupRecordMapper) ToTopupMonthlyMethodsByCardNumber(s []*db.GetMonthlyTopupMethodsByCardNumberRow) []*record.TopupMonthMethod {
	var topupRecords []*record.TopupMonthMethod

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupMonthlyMethodByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupYearlyMethodByCardNumber maps a GetYearlyTopupMethodsByCardNumberRow database row to a TopupYearlyMethod domain model.
// It takes a pointer to a GetYearlyTopupMethodsByCardNumberRow as a parameter and returns a pointer to a TopupYearlyMethod.
// The function maps the fields of GetYearlyTopupMethodsByCardNumberRow to the corresponding fields of TopupYearlyMethod.
func (t *topupRecordMapper) ToTopupYearlyMethodByCardNumber(s *db.GetYearlyTopupMethodsByCardNumberRow) *record.TopupYearlyMethod {
	return &record.TopupYearlyMethod{
		Year:        s.Year,
		TopupMethod: s.TopupMethod,
		TotalTopups: int(s.TotalTopups),
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyMethodsByCardNumber maps a slice of GetYearlyTopupMethodsByCardNumberRow database rows
// to a slice of TopupYearlyMethod domain models. It iterates over the provided slice, converting
// each element using the ToTopupYearlyMethodByCardNumber method, and appends the result to a new
// slice. The function returns a slice of pointers to TopupYearlyMethod.
func (t *topupRecordMapper) ToTopupYearlyMethodsByCardNumber(s []*db.GetYearlyTopupMethodsByCardNumberRow) []*record.TopupYearlyMethod {
	var topupRecords []*record.TopupYearlyMethod

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupYearlyMethodByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupMonthlyAmountByCardNumber maps a GetMonthlyTopupAmountsByCardNumberRow to a TopupMonthAmount domain model.
//
// Args:
//   - s: A pointer to a GetMonthlyTopupAmountsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TopupMonthAmount containing the mapped data, including Month and TotalAmount.
func (t *topupRecordMapper) ToTopupMonthlyAmountByCardNumber(s *db.GetMonthlyTopupAmountsByCardNumberRow) *record.TopupMonthAmount {
	return &record.TopupMonthAmount{
		Month:       s.Month,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupMonthlyAmountsByCardNumber maps a slice of GetMonthlyTopupAmountsByCardNumberRow database rows to a slice of TopupMonthAmount domain models.
// It iterates over the provided slice, converting each element using the ToTopupMonthlyAmountByCardNumber method,
// and appends the result to a new slice. The function returns a slice of pointers to TopupMonthAmount.
func (t *topupRecordMapper) ToTopupMonthlyAmountsByCardNumber(s []*db.GetMonthlyTopupAmountsByCardNumberRow) []*record.TopupMonthAmount {
	var topupRecords []*record.TopupMonthAmount

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupMonthlyAmountByCardNumber(topup))
	}

	return topupRecords
}

// ToTopupYearlyAmountByCardNumber maps a GetYearlyTopupAmountsByCardNumberRow to a TopupYearlyAmount domain model.
//
// Args:
//   - s: A pointer to a GetYearlyTopupAmountsByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a TopupYearlyAmount containing the mapped data, including Year and TotalAmount.
func (t *topupRecordMapper) ToTopupYearlyAmountByCardNumber(s *db.GetYearlyTopupAmountsByCardNumberRow) *record.TopupYearlyAmount {
	return &record.TopupYearlyAmount{
		Year:        s.Year,
		TotalAmount: int(s.TotalAmount),
	}
}

// ToTopupYearlyAmountsByCardNumber maps a slice of GetYearlyTopupAmountsByCardNumberRow to a
// slice of TopupYearlyAmount domain models.
//
// Args:
//   - s: A slice of pointers to GetYearlyTopupAmountsByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to TopupYearlyAmount containing the mapped data, including Year and TotalAmount.
func (t *topupRecordMapper) ToTopupYearlyAmountsByCardNumber(s []*db.GetYearlyTopupAmountsByCardNumberRow) []*record.TopupYearlyAmount {
	var topupRecords []*record.TopupYearlyAmount

	for _, topup := range s {
		topupRecords = append(topupRecords, t.ToTopupYearlyAmountByCardNumber(topup))
	}

	return topupRecords
}
