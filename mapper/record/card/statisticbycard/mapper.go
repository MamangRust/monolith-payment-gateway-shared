package statisticbycard

// cardStatisticByCardRecordMapper provides methods to map database rows to CardStatisticByCardRecord domain models.
type cardStatisticByCardRecordMapper struct {
	CardStatisticBalanceByCardRecordMapper     CardStatisticBalanceByCardRecordMapper
	CardStatisticTopupByCardRecordMapper       CardStatisticTopupByCardRecordMapper
	CardStatisticWithdrawByCardRecordMapper    CardStatisticWithdrawByCardRecordMapper
	CardStatisticTransferByCardRecordMapper    CardStatisticTransferByCardRecordMapper
	CardStatisticTransactionByCardRecordMapper CardStatisticTransactionByCardRecordMapper
}

// NewCardStatisticByCardRecordMapper creates a new instance of cardStatisticByCardRecordMapper,
// initializing all the individual mappers for balance, topup, withdraw, transfer, and transaction
// statistics by card. This function returns a pointer to cardStatisticByCardRecordMapper with all
// the necessary dependencies set up for mapping database rows to their respective domain models.
func NewCardStatisticByCardRecordMapper() *cardStatisticByCardRecordMapper {
	return &cardStatisticByCardRecordMapper{
		CardStatisticBalanceByCardRecordMapper:     NewCardStatisticBalanceByCardRecordMapper(),
		CardStatisticTopupByCardRecordMapper:       NewCardStatisticTopupRecordMapper(),
		CardStatisticWithdrawByCardRecordMapper:    NewCardStatisticWithdrawRecordMapper(),
		CardStatisticTransferByCardRecordMapper:    NewCardStatisticTransferRecordMapper(),
		CardStatisticTransactionByCardRecordMapper: NewCardStatisticTransactionRecordMapper(),
	}
}
