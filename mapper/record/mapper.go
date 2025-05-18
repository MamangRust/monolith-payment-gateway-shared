package recordmapper

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
