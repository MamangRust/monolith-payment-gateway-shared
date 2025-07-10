package transactionstatbycard

// transactionStatisticByCardRecordMapper is a struct that implements the TransactionStatisticByCardRecordMapper interface.
type transactionStatisticByCardRecordMapper struct {
	TransactionStatisticStatusByCardMapper TransactionStatisticByCardStatusMapper
	TransactionStatisticAmountByCardMapper TransactionStatisticByCardAmountMapper
	TransactionStatisticMethodByCardMapper TransactionStatisticByCardMethodMapper
}

// NewTransactionStatisticsByCardRecordMapper returns a new instance of
// transactionStatisticByCardRecordMapper, which provides methods to map
// database rows related to transaction statistics by card number to domain
// models.
func NewTransactionStatisticsByCardRecordMapper() *transactionStatisticByCardRecordMapper {
	return &transactionStatisticByCardRecordMapper{
		TransactionStatisticStatusByCardMapper: NewTransactionStatisticMethodByCardMapper(),
		TransactionStatisticAmountByCardMapper: NewTransactionStatisticAmountByCardMapper(),
		TransactionStatisticMethodByCardMapper: NewTransactionStatisticMethodByCardMapper(),
	}
}
