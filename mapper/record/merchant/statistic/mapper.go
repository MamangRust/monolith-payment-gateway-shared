package merchantstats

// merchantStatisticRecordMapper maps database rows to domain models for statistic operations.
type merchantStatisticRecordMapper struct {
	MerchantStatisticTotalAmountRecordMapper MerchantStatisticTotalAmountRecordMapper
	MerchantStatisticMethodRecordMapper      MerchantStatisticMethodRecordMapper
	MerchantStatisticAmountRecordMapper      MerchantStatisticAmountRecordMapper
}

// NewMerchantStatisticRecordMapper creates a new instance of
// MerchantStatisticRecordMapper, providing methods to map database rows
// to MerchantMonthlyPaymentMethod, MerchantMonthlyAmount, and
// MerchantMonthlyTotalAmount domain models.
func NewMerchantStatisticRecordMapper() *merchantStatisticRecordMapper {
	return &merchantStatisticRecordMapper{
		MerchantStatisticTotalAmountRecordMapper: NewMerchantStatisticTotalAmountRecordMapper(),
		MerchantStatisticMethodRecordMapper:      NewMerchantStatisticMethodRecordMapper(),
		MerchantStatisticAmountRecordMapper:      NewMerchantStatisticAmountRecordMapper(),
	}
}
