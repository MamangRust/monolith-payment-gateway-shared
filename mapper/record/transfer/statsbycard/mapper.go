package transferstatsbycardrecordmapper

// transferStatisticByCardRecordMapper is a struct that holds the mappers for transfer statistics related to card number.
type transferStatisticByCardRecordMapper struct {
	TransferStatisticStatusByCardRecordMapper         TransferStatisticStatusByCardRecordMapper
	TransferStatisticSenderAmountByCardRecordMapper   TransferStatisticSenderAmountByCardRecordMapper
	TransferStatisticReceiverAmountByCardRecordMapper TransferStatisticReceiverAmountByCardRecordMapper
}

// NewTransferStatisticByCardRecordMapper creates and returns a new instance of transferStatisticByCardRecordMapper.
// This mapper provides methods to map database rows related to transfer statistics, including status,
// sender amounts, and receiver amounts, filtered by card number to their corresponding domain models.
func NewTransferStatisticByCardRecordMapper() *transferStatisticByCardRecordMapper {
	return &transferStatisticByCardRecordMapper{
		TransferStatisticStatusByCardRecordMapper:         NewTransferStatisticStatusByCardRecordMapper(),
		TransferStatisticSenderAmountByCardRecordMapper:   NewTransferStatisticSenderAmountByCardRecordMapper(),
		TransferStatisticReceiverAmountByCardRecordMapper: NewTransferStatisticReceiverAmountByCardRecordMapper(),
	}
}
