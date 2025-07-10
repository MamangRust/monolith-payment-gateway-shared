package merchantdocumentapimapper

type merchantDocumentResponseMapper struct {
	MerchantDocumentQueryResponseMapper   MerchantDocumentQueryResponseMapper
	MerchantDocumentCommandResponseMapper MerchantDocumentCommandResponseMapper
}

func NewMerchantDocumentResponseMapper() *merchantDocumentResponseMapper {
	return &merchantDocumentResponseMapper{
		MerchantDocumentQueryResponseMapper:   NewMerchantDocumentQueryResponseMapper(),
		MerchantDocumentCommandResponseMapper: NewMerchantDocumentCommandResponseMapper(),
	}
}
