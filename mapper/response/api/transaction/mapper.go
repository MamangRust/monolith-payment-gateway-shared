package transactionapimapper

type trnasactionResponseMapper struct {
	TransactionQueryResponseMapper       TransactionQueryResponseMapper
	TransactionCommandResponseMapper     TransactionCommandResponseMapper
	TransactionStatsStatusResponseMapper TransactionStatsStatusResponseMapper
	TransactionStatsMethodResponseMapper TransactionStatsMethodResponseMapper
	TransactionStatsAmountResponseMapper TransactionStatsAmountResponseMapper
}

func NewTransactionResponseMapper() *trnasactionResponseMapper {
	return &trnasactionResponseMapper{
		TransactionQueryResponseMapper:       NewTransactionQueryResponseMapper(),
		TransactionCommandResponseMapper:     NewTransactionCommandResponseMapper(),
		TransactionStatsStatusResponseMapper: NewTransactionStatsStatusResponseMapper(),
		TransactionStatsMethodResponseMapper: NewTransactionStatsMethodResponseMapper(),
		TransactionStatsAmountResponseMapper: NewTransactionStatsAmountResponseMapper(),
	}
}
