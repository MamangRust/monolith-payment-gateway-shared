package transactionservicemapper

// transactionResponseMapper provides methods to map TransactionRecord domain models to API-compatible
// TransactionResponse types.
type transactionResponseMapper struct {
	TransactionQueryResponseMapper       TransactionQueryResponseMapper
	TransactionCommandResponseMapper     TransactionCommandResponseMapper
	TransactionStatsAmountResponseMapper TransactionStatsAmountResponseMapper
	TransactionStatsMethodResponseMapper TransactionStatsMethodResponseMapper
	TransactionStatsStatusResponseMapper TransactionStatsStatusResponseMapper
}

// NewTransactionResponseMapper constructs and returns a new instance of transactionResponseMapper.
// This function initializes the mapper with various sub-mappers responsible for converting
// different aspects of transaction records into their corresponding API-compatible response types.
// The sub-mappers include query, command, stats for amount, method, and status.
func NewTransactionResponseMapper() *transactionResponseMapper {
	return &transactionResponseMapper{
		TransactionQueryResponseMapper:       NewTransactionQueryResponseMapper(),
		TransactionCommandResponseMapper:     NewTransactionCommandResponseMapper(),
		TransactionStatsAmountResponseMapper: NewTransactionStatsAmountResponseMapper(),
		TransactionStatsMethodResponseMapper: NewTransactionStatsMethodResponseMapper(),
		TransactionStatsStatusResponseMapper: NewTransactionStatsStatusResponseMapper(),
	}
}
