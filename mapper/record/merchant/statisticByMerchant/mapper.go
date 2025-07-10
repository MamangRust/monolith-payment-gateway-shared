package merchantstatbyid

// merchantStatisticByMerchantRecordMapper provides methods to map Merchant database rows to MerchantRecord domain models for statistic operations.
type merchantStatisticByMerchantRecordMapper struct {
	MerchantStatisticMethodByMerchantMapper MerchantStatisticMethodByMerchantMapper

	MerchantStatisticAmountByMerchantMapper MerchantStatisticAmountByMerchantMapper

	MerchantStatisticTotalAmountByMerchantMapper MerchantStatisticTotalAmountByMerchantMapper
}

// NewMerchantStatisticByMerchantRecordMapper returns a new instance of
// MerchantStatisticByMerchantMapper, which provides methods to map Merchant
// database rows to MerchantRecord domain models for statistic operations.
func NewMerchantStatisticByMerchantRecordMapper() *merchantStatisticByMerchantRecordMapper {
	return &merchantStatisticByMerchantRecordMapper{
		MerchantStatisticMethodByMerchantMapper:      NewMerchantStatisticMethodByMerchantRecordMapper(),
		MerchantStatisticAmountByMerchantMapper:      NewMerchantStatisticAmountByMerchantRecordMapper(),
		MerchantStatisticTotalAmountByMerchantMapper: NewMerchantStatisticTotalAmountByMerchantRecordMapper(),
	}
}
