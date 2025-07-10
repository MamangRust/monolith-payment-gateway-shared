package merchantdocumentrecordmapper

// merchantDocumentRecordMapper provides methods to map MerchantDocument database rows to MerchantDocumentRecord domain models.
type merchantDocumentRecordMapper struct {
	MerchantDocumentQueryMapper   MerchantDocumentQueryMapper
	MerchantDocumentCommandMapper MerchantDocumentCommandMapper
}

// NewMerchantDocumentRecordMapper returns a new instance of merchantDocumentRecordMapper,
// which provides methods to map MerchantDocument database rows to MerchantDocumentRecord domain models.
// It initializes both the query and command mappers to facilitate a full range of mapping operations.
func NewMerchantDocumentRecordMapper() *merchantDocumentRecordMapper {
	return &merchantDocumentRecordMapper{
		MerchantDocumentQueryMapper:   NewMerchantDocumentQueryMapper(),
		MerchantDocumentCommandMapper: NewMerchantDocumentCommandMapper(),
	}
}
