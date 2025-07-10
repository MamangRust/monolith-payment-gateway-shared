package transferresponsemapper

// transferResponseMapper provides methods to map TransferRecord, TransferResponseDeleteAt, TransferMonthAmount,  TransferYearAmount, TransferRecordMonthStatusSuccess, TransferRecordMonthStatusFailed, TransferRecordYearStatusSuccess, and TransferRecordYearStatusFailed domain models to API-compatible response types.
type transferResponseMapper struct {
	TransferQueryResponseMapper       TransferQueryResponseMapper
	TransferCommandResponseMapper     TransferCommandResponseMapper
	TransferStatsAmountResponseMapper TransferAmountResponseMapper
	TransferStatsStatusResponseMapper TransferStatsStatusResponseMapper
}

// NewTransferBaseResponseMapper returns a new instance of transferResponseMapper,
// which provides methods to map TransferRecord, TransferResponseDeleteAt,
// TransferMonthAmount, TransferYearAmount, TransferRecordMonthStatusSuccess,
// TransferRecordMonthStatusFailed, TransferRecordYearStatusSuccess, and
// TransferRecordYearStatusFailed domain models to API-compatible response
// types.
func NewTransferBaseResponseMapper() *transferResponseMapper {
	return &transferResponseMapper{
		TransferQueryResponseMapper:       NewTransferQueryResponseMapper(),
		TransferCommandResponseMapper:     NewTransferCommandResponseMapper(),
		TransferStatsAmountResponseMapper: NewTransferStatsAmountResponseMapper(),
		TransferStatsStatusResponseMapper: NewTransferStatsStatusResponseMapper(),
	}
}
