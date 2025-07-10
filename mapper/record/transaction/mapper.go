package transactionrecordmapper

// transactionRecordMapper is a struct that provides methods to map database rows to transaction domain model records.
type transactionRecordMapper struct {
	TransactionQueryRecordMapper   TransactionQueryRecordMapper
	TransactionCommandRecordMapper TransactionCommandRecordMapper
}

// NewTransactionRecordMapper returns a new instance of transactionRecordMapper,
// initializing it with query and command record mappers for mapping
// Transaction database rows to TransactionRecord domain models.
func NewTransactionRecordMapper() *transactionRecordMapper {
	return &transactionRecordMapper{
		TransactionQueryRecordMapper:   NewTransactionQueryRecordMapper(),
		TransactionCommandRecordMapper: NewTransactionCommandRecordMapper(),
	}
}
