package merchantdocumentservicemapper

// merchantDocumentResponseMapper provides methods to map MerchantDocumentRecord domain models to MerchantDocumentResponse API-compatible responses.
type merchantDocumentResponseMapper struct {
	MerchantDocumentQueryResponseMapper   MerchantDocumentQueryResponseMapper
	MerchantDocumentCommandResponseMapper MerchantDocumentCommandResponseMapper
}

// NewMerchantDocumentResponseMapper creates a new instance of merchantDocumentResponseMapper.
//
// It initializes both the query and command mappers to facilitate a full range of mapping operations.
func NewMerchantDocumentResponseMapper() *merchantDocumentResponseMapper {
	return &merchantDocumentResponseMapper{
		MerchantDocumentQueryResponseMapper:   NewMerchantDocumentQueryResponseMapper(),
		MerchantDocumentCommandResponseMapper: NewMerchantDocumentCommandResponseMapper(),
	}
}
