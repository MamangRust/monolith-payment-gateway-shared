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

// NewProtoMapper initializes and returns a new instance of ProtoMapper.
// This function sets up all the necessary proto mappers, including those for
// authentication, roles, cards, merchants, merchant documents, saldo, top-ups,
// transactions, transfers, users, and withdrawals. Each field is initialized
// with its respective proto mapper constructor function.
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
