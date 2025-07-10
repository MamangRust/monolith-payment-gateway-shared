package merchantrecordmapper

// merchantRecordMapper is a struct that provides methods to map Merchant database rows to MerchantRecord domain models.
type merchantRecordMapper struct {
	MerchantQueryRecordMapper       MerchantQueryRecordMapper
	MerchantCommandRecordMapper     MerchantCommandRecordMapper
	MerchantTransactionRecordMapper MerchantTransactionRecordMapper
}

// NewMerchantRecordMapper returns a new instance of merchantRecordMapper,
// initializing it with query, command, and transaction record mappers for
// mapping Merchant database rows to MerchantRecord domain models.
func NewMerchantRecordMapper() *merchantRecordMapper {
	return &merchantRecordMapper{
		MerchantQueryRecordMapper:       NewMerchantQueryRecordMapper(),
		MerchantCommandRecordMapper:     NewMerchantDocumentCommandMapper(),
		MerchantTransactionRecordMapper: NewMerchantTransactionRecordMapper(),
	}
}
