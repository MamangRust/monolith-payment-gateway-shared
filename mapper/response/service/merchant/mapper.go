package merchantservicemapper

// merchantResponseMapper is a struct that provides methods to map MerchantRecord domain models to their corresponding API-compatible response types.
type merchantResponseMapper struct {
	MerchantQueryResponseMapper       MerchantQueryResponseMapper
	MerchantCommandResponseMapper     MerchantCommandResponseMapper
	MerchantTotalAmountResponseMapper MerchantTotalAmountResponseMapper
	MerchantAmountResponseMapper      MerchantAmountResponseMapper
}

// NewMerchantResponseMapper returns a new instance of merchantResponseMapper, which provides methods to map MerchantRecord domain models to their corresponding API-compatible response types.
func NewMerchantResponseMapper() *merchantResponseMapper {
	return &merchantResponseMapper{
		MerchantQueryResponseMapper:       NewMerchantQueryResponseMapper(),
		MerchantCommandResponseMapper:     NewMerchantCommandResponseMapper(),
		MerchantTotalAmountResponseMapper: NewMerchantTotalAmountResponseMapper(),
		MerchantAmountResponseMapper:      NewMerchantAmountResponseMapper(),
	}
}
