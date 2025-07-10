package statisticbycard

// topStatisticByCardRecordMapper provides methods to map Topup database rows to TopupRecord domain models for statistic operations.
type topStatisticByCardRecordMapper struct {
	TopupStatisticStatusMapper TopupStatisticStatusByCardNumberMapper
	TopupStatisticMethodMapper TopupStatisticMethodByCardNumberMapper
	TopupStatisticAmountMapper TopupStatisticAmountByCardNumberMapper
}

// NewTopupStatisticByCardRecordMapper returns a new instance of topStatisticByCardRecordMapper which provides methods to map Topup database rows to TopupRecord domain models for statistic operations.
func NewTopupStatisticByCardRecordMapper() *topStatisticByCardRecordMapper {
	return &topStatisticByCardRecordMapper{
		TopupStatisticStatusMapper: NewTopupStatisticStatusByCardNumberMapper(),
		TopupStatisticMethodMapper: NewTopupStatisticMethodByCardNumberMapper(),
		TopupStatisticAmountMapper: NewTopupStatisticAmountByCardNumberMapper(),
	}
}
