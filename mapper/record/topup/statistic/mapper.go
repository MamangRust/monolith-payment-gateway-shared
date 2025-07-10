package topupstatsrecord

// topupStatisticRecordMapper provides methods to map Topup database rows to TopupRecord domain models for statistic operations.
type topupStatisticRecordMapper struct {
	TopupStatisticStatusMapper TopupStatisticStatusMapper
	TopupStatisticMethodMapper TopupStatisticMethodMapper
	TopupStatisticAmountMapper TopupStatisticAmountMapper
}

// NewTopupStatisticRecordMapper returns a new instance of topupStatisticRecordMapper which provides methods to map Topup database rows to TopupRecord domain models for statistic operations.
func NewTopupStatisticRecordMapper() *topupStatisticRecordMapper {
	return &topupStatisticRecordMapper{
		TopupStatisticStatusMapper: NewTopupStatisticStatusRecordMapper(),
		TopupStatisticMethodMapper: NewTopupStatisticMethodRecordMapper(),
		TopupStatisticAmountMapper: NewTopupStatisticAmountRecordMapper(),
	}
}
