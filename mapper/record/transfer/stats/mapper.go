package transferstatsrecordmapper

type transferStatisticRecordMapper struct {
	TransferStatisticAmountRecordMapper TransferStatisticAmountRecordMapper
	TransferStatisticStatusRecordMapper TransferStatisticStatusRecordMapper
}

// NewTransferStatisticRecordMapper returns a new instance of transferStatisticRecordMapper
// which provides methods to map database rows to TransferStatisticAmount and
// TransferStatisticStatus domain models.
func NewTransferStatisticRecordMapper() *transferStatisticRecordMapper {
	return &transferStatisticRecordMapper{
		TransferStatisticAmountRecordMapper: NewTransferStatisticAmountRecordMapper(),
		TransferStatisticStatusRecordMapper: NewTransferStatisticStatusRecordMapper(),
	}
}
