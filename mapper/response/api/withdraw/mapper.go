package withdrawapimapper

type withdrawResponseMapper struct {
	WithdrawQueryResponseMapper       WithdrawQueryResponseMapper
	WithdrawCommandResponseMapper     WithdrawCommandResponseMapper
	WithdrawStatsAmountResponseMapper WithdrawStatsAmountResponseMapper
	WithdrawStatsStatusResponseMapper WithdrawStatsStatusResponseMapper
}

func NewWithdrawResponseMapper() *withdrawResponseMapper {
	return &withdrawResponseMapper{
		WithdrawQueryResponseMapper:       NewWithdrawQueryResponseMapper(),
		WithdrawCommandResponseMapper:     NewWithdrawCommandResponseMapper(),
		WithdrawStatsAmountResponseMapper: NewWithdrawStatsAmountResponseMapper(),
		WithdrawStatsStatusResponseMapper: NewWithdrawStatsStatusResponseMapper(),
	}
}
