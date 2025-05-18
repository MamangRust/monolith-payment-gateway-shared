package protomapper

type ProtoMapper struct {
	AuthProtoMapper             AuthProtoMapper
	RoleProtoMapper             RoleProtoMapper
	CardProtoMapper             CardProtoMapper
	MerchantProtoMapper         MerchantProtoMapper
	MerchantDocumentProtoMapper MerchantDocumentProtoMapper
	SaldoProtoMapper            SaldoProtoMapper
	TopupProtoMapper            TopupProtoMapper
	TransactionProtoMapper      TransactionProtoMapper
	TransferProtoMapper         TransferProtoMapper
	UserProtoMapper             UserProtoMapper
	WithdrawalProtoMapper       WithdrawalProtoMapper
}

func NewProtoMapper() *ProtoMapper {
	return &ProtoMapper{
		AuthProtoMapper:             NewAuthProtoMapper(),
		RoleProtoMapper:             NewRoleProtoMapper(),
		CardProtoMapper:             NewCardProtoMapper(),
		MerchantProtoMapper:         NewMerchantProtoMapper(),
		MerchantDocumentProtoMapper: NewMerchantDocumentProtoMapper(),
		SaldoProtoMapper:            NewSaldoProtoMapper(),
		TopupProtoMapper:            NewTopupProtoMapper(),
		TransactionProtoMapper:      NewTransactionProtoMapper(),
		TransferProtoMapper:         NewTransferProtoMapper(),
		UserProtoMapper:             NewUserProtoMapper(),
		WithdrawalProtoMapper:       NewWithdrawProtoMapper(),
	}
}
