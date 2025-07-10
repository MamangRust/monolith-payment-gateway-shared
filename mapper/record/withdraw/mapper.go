package withdrawrecordmapper

// withdrawRecordMapper is a struct that provides methods to map Withdraw database rows to WithdrawRecord domain models for both query and command operations.
type withdrawRecordMapper struct {
	WithdrawQueryRecordMapping   WithdrawQueryRecordMapping
	WithdrawCommandRecordMapping WithdrawCommandRecordMapping
}

// NewWithdrawRecordMapper creates a new instance of withdrawRecordMapper, which provides methods
// to map Withdraw database rows to WithdrawRecord domain models for both query and command operations.
// It returns a pointer to a withdrawRecordMapper.
func NewWithdrawRecordMapper() *withdrawRecordMapper {
	return &withdrawRecordMapper{
		WithdrawQueryRecordMapping:   NewWithdrawQueryRecordMapper(),
		WithdrawCommandRecordMapping: NewWithdrawCommandRecordMapper(),
	}
}
