package apimapper

type ResponseApiMapper struct {
	AuthResponseMapper        AuthResponseMapper
	CardResponseMapper        CardResponseMapper
	RoleResponseMapper        RoleResponseMapper
	SaldoResponseMapper       SaldoResponseMapper
	TransactionResponseMapper TransactionResponseMapper
	TransferResponseMapper    TransferResponseMapper
	TopupResponseMapper       TopupResponseMapper
	WithdrawResponseMapper    WithdrawResponseMapper
	UserResponseMapper        UserResponseMapper
	MerchantResponseMapper    MerchantResponseMapper
	MerchantDocumentProMapper MerchantDocumentResponseMapper
}

// NewResponseApiMapper creates a new instance of ResponseApiMapper with all the
// other response mappers initialized.
func NewResponseApiMapper() *ResponseApiMapper {
	return &ResponseApiMapper{
		AuthResponseMapper:        NewAuthResponseMapper(),
		CardResponseMapper:        NewCardResponseMapper(),
		SaldoResponseMapper:       NewSaldoResponseMapper(),
		TransactionResponseMapper: NewTransactionResponseMapper(),
		TransferResponseMapper:    NewTransferResponseMapper(),
		TopupResponseMapper:       NewTopupResponseMapper(),
		WithdrawResponseMapper:    NewWithdrawResponseMapper(),
		UserResponseMapper:        NewUserResponseMapper(),
		RoleResponseMapper:        NewRoleResponseMapper(),
		MerchantResponseMapper:    NewMerchantResponseMapper(),
		MerchantDocumentProMapper: NewMerchantDocumentResponseMapper(),
	}
}
