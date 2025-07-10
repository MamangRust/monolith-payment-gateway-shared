package withdrawstatsrecordmapper

// withdrawStatisticRecordMapper is an interface that provides methods to map
// database rows related to withdraw statistics (amount & status) into domain models.
type withdrawStatisticRecordMapper struct {
	WithdrawStatisticAmountRecordMapper WithdrawStatisticAmountRecordMapper
	WithdrawStatisticStatusRecordMapper WithdrawStatisticStatusRecordMapper
}

// NewWithdrawStatisticRecordMapper creates and returns a new instance of
// withdrawStatisticRecordMapper which provides methods to map database rows
// related to withdraw statistics (amount & status) into domain models.
func NewWithdrawStatisticRecordMapper() *withdrawStatisticRecordMapper {
	return &withdrawStatisticRecordMapper{
		WithdrawStatisticAmountRecordMapper: NewWithdrawStatisticAmountRecordMapper(),
		WithdrawStatisticStatusRecordMapper: NewWithdrawStatisticStatusRecordMapper(),
	}
}
