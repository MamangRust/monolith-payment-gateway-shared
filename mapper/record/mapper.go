package recordmapper

// RecordMapper provides a set of mapping functions between database rows and
// domain models for records related to the user, role, user role, refresh
// token, reset token, saldo, topup, transfer, withdraw, card, transaction,
// merchant, and merchant document.
type RecordMapper struct {
	UserRecordMapper             UserRecordMapping
	RoleRecordMapper             RoleRecordMapping
	UserRoleRecordMapper         UserRoleRecordMapping
	RefreshTokenRecordMapper     RefreshTokenRecordMapping
	ResetTokenRecordMapper       ResetTokenRecordMapping
	SaldoRecordMapper            SaldoRecordMapping
	TopupRecordMapper            TopupRecordMapping
	TransferRecordMapper         TransferRecordMapping
	WithdrawRecordMapper         WithdrawRecordMapping
	CardRecordMapper             CardRecordMapping
	TransactionRecordMapper      TransactionRecordMapping
	MerchantRecordMapper         MerchantRecordMapping
	MerchantDocumentRecordMapper MerchantDocumentMapping
}

// NewRecordMapper creates a new RecordMapper instance with all the sub-mappers
// properly initialized.
func NewRecordMapper() *RecordMapper {
	return &RecordMapper{
		UserRecordMapper:             NewUserRecordMapper(),
		RoleRecordMapper:             NewRoleRecordMapper(),
		ResetTokenRecordMapper:       NewResetTokenRecordMapper(),
		UserRoleRecordMapper:         NewUserRoleRecordMapper(),
		RefreshTokenRecordMapper:     NewRefreshTokenRecordMapper(),
		SaldoRecordMapper:            NewSaldoRecordMapper(),
		TopupRecordMapper:            NewTopupRecordMapper(),
		TransferRecordMapper:         NewTransferRecordMapper(),
		WithdrawRecordMapper:         NewWithdrawRecordMapper(),
		CardRecordMapper:             NewCardRecordMapper(),
		TransactionRecordMapper:      NewTransactionRecordMapper(),
		MerchantRecordMapper:         NewMerchantRecordMapper(),
		MerchantDocumentRecordMapper: NewMerchantDocumentRecordMapper(),
	}
}
