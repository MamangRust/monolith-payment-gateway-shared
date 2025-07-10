package withdrawstatsbycardrecordmapper

// withdrawStatisticRecordMapper is an interface that provides methods to map
// database rows related to withdraw statistics (amount & status) into domain models.
type withdrawStatisticRecordMapper struct {
	WithdrawStatisticAmountByCardRecordMapper WithdrawStatisticAmountByCardRecordMapper
	WithdrawStatisticStatusByCardRecordMapper WithdrawStatisticStatusByCardRecordMapper
}

// NewWithdrawStatisticRecordMapper creates and returns a new instance of
// withdrawStatisticRecordMapper which provides methods to map database rows
// related to withdraw statistics (amount & status) filtered by card number
// into domain models.
func NewWithdrawStatisticRecordMapper() *withdrawStatisticRecordMapper {
	return &withdrawStatisticRecordMapper{
		WithdrawStatisticAmountByCardRecordMapper: NewWithdrawStatisticAmountByCardRecordMapper(),
		WithdrawStatisticStatusByCardRecordMapper: NewWithdrawStatisticStatusByCardRecordMapper(),
	}
}
