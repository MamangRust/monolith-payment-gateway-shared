package responseservice

type ResponseServiceMapper struct {
	CardResponseMapper             CardResponseMapper
	RoleResponseMapper             RoleResponseMapper
	RefreshTokenResponseMapper     RefreshTokenResponseMapper
	SaldoResponseMapper            SaldoResponseMapper
	TransactionResponseMapper      TransactionResponseMapper
	TransferResponseMapper         TransferResponseMapper
	TopupResponseMapper            TopupResponseMapper
	WithdrawResponseMapper         WithdrawResponseMapper
	UserResponseMapper             UserResponseMapper
	MerchantResponseMapper         MerchantResponseMapper
	MerchantDocumentResponseMapper MerchantDocumentResponseMapper
}

// NewResponseServiceMapper creates a new ResponseServiceMapper, which is a
// collection of mappers that can be used to map the internal domain models
// into structured API response objects.
func NewResponseServiceMapper() *ResponseServiceMapper {
	return &ResponseServiceMapper{
		CardResponseMapper:             NewCardResponseMapper(),
		SaldoResponseMapper:            NewSaldoResponseMapper(),
		TransactionResponseMapper:      NewTransactionResponseMapper(),
		TransferResponseMapper:         NewTransferResponseMapper(),
		TopupResponseMapper:            NewTopupResponseMapper(),
		WithdrawResponseMapper:         NewWithdrawResponseMapper(),
		UserResponseMapper:             NewUserResponseMapper(),
		RefreshTokenResponseMapper:     NewRefreshTokenResponseMapper(),
		RoleResponseMapper:             NewRoleResponseMapper(),
		MerchantResponseMapper:         NewMerchantResponseMapper(),
		MerchantDocumentResponseMapper: NewMerchantDocumentResponseMapper(),
	}
}
