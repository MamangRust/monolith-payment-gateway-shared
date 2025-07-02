package recordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

// cardRecordMapper provides methods to map Card database rows to CardRecord domain models.
type cardRecordMapper struct {
}

// NewCardRecordMapper creates a new CardRecordMapper, which is responsible for mapping
// between data transfer objects and the internal database schema.
func NewCardRecordMapper() *cardRecordMapper {
	return &cardRecordMapper{}
}

// ToCardEmailRecord maps a GetUserEmailByCardNumberRow database row to a CardEmailRecord domain model.
//
// Args:
//   - card: A pointer to a GetUserEmailByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardEmailRecord containing the mapped data, including ID, UserID,
//     Email, CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
func (s *cardRecordMapper) ToCardEmailRecord(card *db.GetUserEmailByCardNumberRow) *record.CardEmailRecord {

	return &record.CardEmailRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		Email:        card.Email,
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
	}
}

// ToCardRecord maps a Card database row to a CardRecord domain model.
//
// Args:
//   - card: A pointer to a Card representing the database row.
//
// Returns:
//   - A pointer to a CardRecord containing the mapped data, including ID, UserID, CardNumber,
//     CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardRecordMapper) ToCardRecord(card *db.Card) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formatedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.CardRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

// ToCardGetAll maps a GetCardsRow database row to a CardRecord domain model.
//
// Args:
//   - card: A pointer to a GetCardsRow representing the database row.
//
// Returns:
//   - A pointer to a CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardRecordMapper) ToCardGetAll(card *db.GetCardsRow) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formatedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.CardRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

// ToCardsRecord maps a slice of GetCardsRow database rows to a slice of CardRecord domain models.
//
// Args:
//   - cards: A slice of pointers to GetCardsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardRecordMapper) ToCardsRecord(cards []*db.GetCardsRow) []*record.CardRecord {
	var records []*record.CardRecord
	for _, card := range cards {
		records = append(records, s.ToCardGetAll(card))
	}
	return records
}

// ToCardRecordActive maps a GetActiveCardsWithCountRow database row to a CardRecord domain model.
//
// Args:
//   - card: A pointer to a GetActiveCardsWithCountRow representing the database row.
//
// Returns:
//   - A pointer to a CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardRecordMapper) ToCardRecordActive(card *db.GetActiveCardsWithCountRow) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formattedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formattedDeletedAt
	}

	return &record.CardRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

// ToCardRecordsActive maps a slice of GetActiveCardsWithCountRow database rows to a slice of CardRecord domain models.
//
// Args:
//   - cards: A slice of pointers to GetActiveCardsWithCountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardRecordMapper) ToCardRecordsActive(cards []*db.GetActiveCardsWithCountRow) []*record.CardRecord {
	var records []*record.CardRecord
	for _, card := range cards {
		records = append(records, s.ToCardRecordActive(card))
	}
	return records
}

// ToCardRecordTrashed maps a GetTrashedCardsWithCountRow database row to a CardRecord domain model.
//
// Args:
//   - card: A pointer to a GetTrashedCardsWithCountRow representing the database row.
//
// Returns:
//   - A pointer to a CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardRecordMapper) ToCardRecordTrashed(card *db.GetTrashedCardsWithCountRow) *record.CardRecord {
	var deletedAt *string

	if card.DeletedAt.Valid {
		formattedDeletedAt := card.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formattedDeletedAt
	}

	return &record.CardRecord{
		ID:           int(card.CardID),
		UserID:       int(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate.Format("2006-01-02"),
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    card.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

// ToCardRecordsTrashed maps a slice of GetTrashedCardsWithCountRow database rows to a slice of CardRecord domain models.
//
// Args:
//   - cards: A slice of pointers to GetTrashedCardsWithCountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardRecord containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, UpdatedAt, and DeletedAt.
func (s *cardRecordMapper) ToCardRecordsTrashed(cards []*db.GetTrashedCardsWithCountRow) []*record.CardRecord {
	var records []*record.CardRecord
	for _, card := range cards {
		records = append(records, s.ToCardRecordTrashed(card))
	}
	return records
}

// ToMonthlyBalance maps a GetMonthlyBalancesRow database row to a CardMonthBalance domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyBalancesRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthBalance containing the mapped data, including Month and TotalBalance.
func (s *cardRecordMapper) ToMonthlyBalance(card *db.GetMonthlyBalancesRow) *record.CardMonthBalance {
	return &record.CardMonthBalance{
		Month:        card.Month,
		TotalBalance: int64(card.TotalBalance),
	}
}

// ToMonthlyBalances maps a slice of GetMonthlyBalancesRow database rows to a slice of CardMonthBalance domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyBalancesRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthBalance containing the mapped data, including Month and TotalBalance.
func (s *cardRecordMapper) ToMonthlyBalances(cards []*db.GetMonthlyBalancesRow) []*record.CardMonthBalance {
	var records []*record.CardMonthBalance
	for _, card := range cards {
		records = append(records, s.ToMonthlyBalance(card))
	}
	return records
}

// ToYearlyBalance maps a GetYearlyBalancesRow database row to a CardYearlyBalance domain model.
//
// Args:
//   - card: A pointer to a GetYearlyBalancesRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearlyBalance containing the mapped data, including Year and TotalBalance.
func (s *cardRecordMapper) ToYearlyBalance(card *db.GetYearlyBalancesRow) *record.CardYearlyBalance {
	return &record.CardYearlyBalance{
		Year:         card.Year,
		TotalBalance: card.TotalBalance,
	}
}

// ToYearlyBalances maps a slice of GetYearlyBalancesRow database rows to a slice of CardYearlyBalance domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyBalancesRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearlyBalance containing the mapped data, including Year and TotalBalance.
func (s *cardRecordMapper) ToYearlyBalances(cards []*db.GetYearlyBalancesRow) []*record.CardYearlyBalance {
	var records []*record.CardYearlyBalance
	for _, card := range cards {
		records = append(records, s.ToYearlyBalance(card))
	}
	return records
}

// ToMonthlyTopupAmount maps a GetMonthlyTopupAmountRow database row to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyTopupAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTopupAmount(card *db.GetMonthlyTopupAmountRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTopupAmount),
	}
}

// ToMonthlyTopupAmounts maps a slice of GetMonthlyTopupAmountRow database rows to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyTopupAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTopupAmounts(cards []*db.GetMonthlyTopupAmountRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTopupAmount(card))
	}
	return records
}

// ToYearlyTopupAmount maps a GetYearlyTopupAmountRow database row to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyTopupAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTopupAmount(card *db.GetYearlyTopupAmountRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTopupAmount,
	}
}

// ToYearlyTopupAmounts maps a slice of GetYearlyTopupAmountRow database rows to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyTopupAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTopupAmounts(cards []*db.GetYearlyTopupAmountRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTopupAmount(card))
	}
	return records
}

// ToMonthlyWithdrawAmount maps a GetMonthlyWithdrawAmountRow database row to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyWithdrawAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyWithdrawAmount(card *db.GetMonthlyWithdrawAmountRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalWithdrawAmount),
	}
}

// ToMonthlyWithdrawAmounts maps a slice of GetMonthlyWithdrawAmountRow database rows to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyWithdrawAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyWithdrawAmounts(cards []*db.GetMonthlyWithdrawAmountRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyWithdrawAmount(card))
	}
	return records
}

// ToYearlyWithdrawAmount maps a GetYearlyWithdrawAmountRow database row to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyWithdrawAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyWithdrawAmount(card *db.GetYearlyWithdrawAmountRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalWithdrawAmount,
	}
}

// ToYearlyWithdrawAmounts maps a slice of GetYearlyWithdrawAmountRow database rows to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyWithdrawAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyWithdrawAmounts(cards []*db.GetYearlyWithdrawAmountRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyWithdrawAmount(card))
	}
	return records
}

// ToMonthlyTransactionAmount maps a GetMonthlyTransactionAmountRow database row to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyTransactionAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransactionAmount(card *db.GetMonthlyTransactionAmountRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTransactionAmount),
	}
}

// ToMonthlyTransactionAmounts maps a slice of GetMonthlyTransactionAmountRow database rows to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyTransactionAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransactionAmounts(cards []*db.GetMonthlyTransactionAmountRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransactionAmount(card))
	}
	return records
}

// ToYearlyTransactionAmount maps a GetYearlyTransactionAmountRow database row to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyTransactionAmountRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransactionAmount(card *db.GetYearlyTransactionAmountRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTransactionAmount,
	}
}

// ToYearlyTransactionAmounts maps a slice of GetYearlyTransactionAmountRow database rows to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyTransactionAmountRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransactionAmounts(cards []*db.GetYearlyTransactionAmountRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransactionAmount(card))
	}
	return records
}

// ToMonthlyTransferSenderAmount maps a GetMonthlyTransferAmountSenderRow database row to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyTransferAmountSenderRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransferSenderAmount(card *db.GetMonthlyTransferAmountSenderRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalSentAmount),
	}
}

// ToMonthlyTransferSenderAmounts maps a slice of GetMonthlyTransferAmountSenderRow database rows to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyTransferAmountSenderRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransferSenderAmounts(cards []*db.GetMonthlyTransferAmountSenderRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferSenderAmount(card))
	}
	return records
}

// ToYearlyTransferSenderAmount maps a GetYearlyTransferAmountSenderRow database row to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyTransferAmountSenderRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransferSenderAmount(card *db.GetYearlyTransferAmountSenderRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalSentAmount,
	}
}

// ToYearlyTransferSenderAmounts maps a slice of GetYearlyTransferAmountSenderRow database rows to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyTransferAmountSenderRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransferSenderAmounts(cards []*db.GetYearlyTransferAmountSenderRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferSenderAmount(card))
	}
	return records
}

// ToMonthlyTransferReceiverAmount maps a GetMonthlyTransferAmountReceiverRow database row to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyTransferAmountReceiverRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransferReceiverAmount(card *db.GetMonthlyTransferAmountReceiverRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalReceivedAmount),
	}
}

// ToMonthlyTransferReceiverAmounts maps a slice of GetMonthlyTransferAmountReceiverRow database rows to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyTransferAmountReceiverRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransferReceiverAmounts(cards []*db.GetMonthlyTransferAmountReceiverRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferReceiverAmount(card))
	}
	return records
}

// ToYearlyTransferReceiverAmount maps a GetYearlyTransferAmountReceiverRow database row to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyTransferAmountReceiverRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransferReceiverAmount(card *db.GetYearlyTransferAmountReceiverRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalReceivedAmount,
	}
}

// ToYearlyTransferReceiverAmounts maps a slice of GetYearlyTransferAmountReceiverRow database rows to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyTransferAmountReceiverRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransferReceiverAmounts(cards []*db.GetYearlyTransferAmountReceiverRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferReceiverAmount(card))
	}
	return records
}

// ToMonthlyBalanceCardNumber maps a GetMonthlyBalancesByCardNumberRow database row to a CardMonthBalance domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyBalancesByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthBalance containing the mapped data, including Month and TotalBalance.
func (s *cardRecordMapper) ToMonthlyBalanceCardNumber(card *db.GetMonthlyBalancesByCardNumberRow) *record.CardMonthBalance {
	return &record.CardMonthBalance{
		Month:        card.Month,
		TotalBalance: int64(card.TotalBalance),
	}
}

// ToMonthlyBalancesCardNumber maps a slice of GetMonthlyBalancesByCardNumberRow database rows to a slice of CardMonthBalance domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyBalancesByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthBalance containing the mapped data, including Month and TotalBalance.
func (s *cardRecordMapper) ToMonthlyBalancesCardNumber(cards []*db.GetMonthlyBalancesByCardNumberRow) []*record.CardMonthBalance {
	var records []*record.CardMonthBalance
	for _, card := range cards {
		records = append(records, s.ToMonthlyBalanceCardNumber(card))
	}
	return records
}

// ToYearlyBalanceCardNumber maps a GetYearlyBalancesByCardNumberRow database row to a CardYearlyBalance domain model.
//
// Args:
//   - card: A pointer to a GetYearlyBalancesByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearlyBalance containing the mapped data, including Year and TotalBalance.
func (s *cardRecordMapper) ToYearlyBalanceCardNumber(card *db.GetYearlyBalancesByCardNumberRow) *record.CardYearlyBalance {
	return &record.CardYearlyBalance{
		Year:         card.Year,
		TotalBalance: card.TotalBalance,
	}
}

// ToYearlyBalancesCardNumber maps a slice of GetYearlyBalancesByCardNumberRow database rows to a slice of CardYearlyBalance domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyBalancesByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearlyBalance containing the mapped data, including Year and TotalBalance.
func (s *cardRecordMapper) ToYearlyBalancesCardNumber(cards []*db.GetYearlyBalancesByCardNumberRow) []*record.CardYearlyBalance {
	var records []*record.CardYearlyBalance
	for _, card := range cards {
		records = append(records, s.ToYearlyBalanceCardNumber(card))
	}
	return records
}

// ToMonthlyTopupAmountByCardNumber maps a GetMonthlyTopupAmountByCardNumberRow database row to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyTopupAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTopupAmountByCardNumber(card *db.GetMonthlyTopupAmountByCardNumberRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTopupAmount),
	}
}

// ToMonthlyTopupAmountsByCardNumber maps a slice of GetMonthlyTopupAmountByCardNumberRow database rows to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyTopupAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTopupAmountsByCardNumber(cards []*db.GetMonthlyTopupAmountByCardNumberRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTopupAmountByCardNumber(card))
	}
	return records
}

// ToYearlyTopupAmountByCardNumber maps a GetYearlyTopupAmountByCardNumberRow database row to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyTopupAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTopupAmountByCardNumber(card *db.GetYearlyTopupAmountByCardNumberRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTopupAmount,
	}
}

// ToYearlyTopupAmountsByCardNumber maps a slice of GetYearlyTopupAmountByCardNumberRow database rows
// to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyTopupAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTopupAmountsByCardNumber(cards []*db.GetYearlyTopupAmountByCardNumberRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTopupAmountByCardNumber(card))
	}
	return records
}

// ToMonthlyWithdrawAmountByCardNumber maps a GetMonthlyWithdrawAmountByCardNumberRow database row to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyWithdrawAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyWithdrawAmountByCardNumber(card *db.GetMonthlyWithdrawAmountByCardNumberRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalWithdrawAmount),
	}
}

// ToMonthlyWithdrawAmountsByCardNumber maps a slice of GetMonthlyWithdrawAmountByCardNumberRow database rows
// to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyWithdrawAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyWithdrawAmountsByCardNumber(cards []*db.GetMonthlyWithdrawAmountByCardNumberRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyWithdrawAmountByCardNumber(card))
	}
	return records
}

// ToYearlyWithdrawAmountByCardNumber maps a GetYearlyWithdrawAmountByCardNumberRow database row to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyWithdrawAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyWithdrawAmountByCardNumber(card *db.GetYearlyWithdrawAmountByCardNumberRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalWithdrawAmount,
	}
}

// ToYearlyWithdrawAmountsByCardNumber maps a slice of GetYearlyWithdrawAmountByCardNumberRow database rows
// to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyWithdrawAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyWithdrawAmountsByCardNumber(cards []*db.GetYearlyWithdrawAmountByCardNumberRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyWithdrawAmountByCardNumber(card))
	}
	return records
}

// ToMonthlyTransactionAmountByCardNumber maps a GetMonthlyTransactionAmountByCardNumberRow database row
// to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyTransactionAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransactionAmountByCardNumber(card *db.GetMonthlyTransactionAmountByCardNumberRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalTransactionAmount),
	}
}

// ToMonthlyTransactionAmountsByCardNumber maps a slice of GetMonthlyTransactionAmountByCardNumberRow database rows
// to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyTransactionAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransactionAmountsByCardNumber(cards []*db.GetMonthlyTransactionAmountByCardNumberRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransactionAmountByCardNumber(card))
	}
	return records
}

// ToYearlyTransactionAmountByCardNumber maps a GetYearlyTransactionAmountByCardNumberRow database row
// to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyTransactionAmountByCardNumberRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransactionAmountByCardNumber(card *db.GetYearlyTransactionAmountByCardNumberRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalTransactionAmount,
	}
}

// ToYearlyTransactionAmountsByCardNumber maps a slice of GetYearlyTransactionAmountByCardNumberRow database rows
// to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyTransactionAmountByCardNumberRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransactionAmountsByCardNumber(cards []*db.GetYearlyTransactionAmountByCardNumberRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransactionAmountByCardNumber(card))
	}
	return records
}

// ToMonthlyTransferSenderAmountByCardNumber maps a GetMonthlyTransferAmountBySenderRow database row to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyTransferAmountBySenderRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransferSenderAmountByCardNumber(card *db.GetMonthlyTransferAmountBySenderRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalSentAmount),
	}
}

// ToMonthlyTransferSenderAmountsByCardNumber maps a slice of GetMonthlyTransferAmountBySenderRow database rows
// to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyTransferAmountBySenderRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransferSenderAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountBySenderRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferSenderAmountByCardNumber(card))
	}
	return records
}

// ToYearlyTransferSenderAmountByCardNumber maps a GetYearlyTransferAmountBySenderRow database row
// to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyTransferAmountBySenderRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransferSenderAmountByCardNumber(card *db.GetYearlyTransferAmountBySenderRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalSentAmount,
	}
}

// ToYearlyTransferSenderAmountsByCardNumber maps a slice of GetYearlyTransferAmountBySenderRow database rows
// to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyTransferAmountBySenderRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransferSenderAmountsByCardNumber(cards []*db.GetYearlyTransferAmountBySenderRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferSenderAmountByCardNumber(card))
	}
	return records
}

// ToMonthlyTransferReceiverAmountByCardNumber maps a GetMonthlyTransferAmountByReceiverRow database row
// to a CardMonthAmount domain model.
//
// Args:
//   - card: A pointer to a GetMonthlyTransferAmountByReceiverRow representing the database row.
//
// Returns:
//   - A pointer to a CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransferReceiverAmountByCardNumber(card *db.GetMonthlyTransferAmountByReceiverRow) *record.CardMonthAmount {
	return &record.CardMonthAmount{
		Month:       card.Month,
		TotalAmount: int64(card.TotalReceivedAmount),
	}
}

// ToMonthlyTransferReceiverAmountsByCardNumber maps a slice of GetMonthlyTransferAmountByReceiverRow database rows
// to a slice of CardMonthAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetMonthlyTransferAmountByReceiverRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardMonthAmount containing the mapped data, including Month and TotalAmount.
func (s *cardRecordMapper) ToMonthlyTransferReceiverAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountByReceiverRow) []*record.CardMonthAmount {
	var records []*record.CardMonthAmount
	for _, card := range cards {
		records = append(records, s.ToMonthlyTransferReceiverAmountByCardNumber(card))
	}
	return records
}

// ToYearlyTransferReceiverAmountByCardNumber maps a GetYearlyTransferAmountByReceiverRow database row
// to a CardYearAmount domain model.
//
// Args:
//   - card: A pointer to a GetYearlyTransferAmountByReceiverRow representing the database row.
//
// Returns:
//   - A pointer to a CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransferReceiverAmountByCardNumber(card *db.GetYearlyTransferAmountByReceiverRow) *record.CardYearAmount {
	return &record.CardYearAmount{
		Year:        card.Year,
		TotalAmount: card.TotalReceivedAmount,
	}
}

// ToYearlyTransferReceiverAmountsByCardNumber maps a slice of GetYearlyTransferAmountByReceiverRow database rows
// to a slice of CardYearAmount domain models.
//
// Args:
//   - cards: A slice of pointers to GetYearlyTransferAmountByReceiverRow representing the database rows.
//
// Returns:
//   - A slice of pointers to CardYearAmount containing the mapped data, including Year and TotalAmount.
func (s *cardRecordMapper) ToYearlyTransferReceiverAmountsByCardNumber(cards []*db.GetYearlyTransferAmountByReceiverRow) []*record.CardYearAmount {
	var records []*record.CardYearAmount
	for _, card := range cards {
		records = append(records, s.ToYearlyTransferReceiverAmountByCardNumber(card))
	}
	return records
}
