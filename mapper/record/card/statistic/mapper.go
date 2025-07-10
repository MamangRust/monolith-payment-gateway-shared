package cardstatisticrecordmapper

// cardStatisticRecordMapper provides methods to map database rows to CardStatisticRecord domain models.
type cardStatisticRecordMapper struct {
	CardStatisticBalanceRecordMapper          CardStatisticBalanceRecordMapper
	CardStatisticTopupRecordMapper            CardStatisticTopupRecordMapper
	CardStatisticTransactionRecordMapper      CardStatisticTransactionRecordMapper
	CardStatisticTopupTransactionRecordMapper CardStatisticTopupRecordMapper
	CardStatTisticTransferRecordMapper        CardStatisticTransferRecordMapper
}

// NewCardStatisticRecordMapper returns a new instance of cardStatisticRecordMapper, which provides methods to map database rows to CardStatisticRecord domain models.
func NewCardStatisticRecordMapper() *cardStatisticRecordMapper {
	return &cardStatisticRecordMapper{
		CardStatisticBalanceRecordMapper:          NewCardStatisticBalanceRecordMapper(),
		CardStatisticTopupRecordMapper:            NewCardStatisticTopupRecordMapper(),
		CardStatisticTransactionRecordMapper:      NewCardStatisticTransactionRecordMapper(),
		CardStatisticTopupTransactionRecordMapper: NewCardStatisticTopupRecordMapper(),
		CardStatTisticTransferRecordMapper:        NewCardStatisticTransferRecordMapper(),
	}
}
