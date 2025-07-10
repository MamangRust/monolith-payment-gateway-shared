package transferapimapper

type transferResponseMapper struct {
	TransferQueryResponseMapper       TransferQueryResponseMapper
	TransferCommandResponseMapper     TransferCommandResponseMapper
	TransferStatsStatusResponseMapper TransferStatsStatusResponseMapper
	TransferStatsAmountResponseMapper TransferStatsAmountResponseMapper
}

func NewTransferResponseMapper() *transferResponseMapper {
	return &transferResponseMapper{
		TransferQueryResponseMapper:       NewTransferQueryResponseMapper(),
		TransferCommandResponseMapper:     NewTransferCommandResponseMapper(),
		TransferStatsStatusResponseMapper: NewTransferStatsStatusResponseMapper(),
		TransferStatsAmountResponseMapper: NewTransferStatsAmountResponseMapper(),
	}
}
