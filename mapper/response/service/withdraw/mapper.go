package withdrawresponseservice

// withdrawResponseMapper consolidates methods to map WithdrawRecord domain models into API-compatible WithdrawResponse, WithdrawResponseMonthStatusSuccess, WithdrwResponseYearStatusSuccess, WithdrawMonthlyStatusFailed, WithdrawYearStatusFailed, WithdrawMonthlyAmount, and WithdrawYearAmount objects.
type withdrawResponseMapper struct {
	WithdrawQueryResponseMapper       WithdrawQueryResponseMapper
	WithdrawCommandResponseMapper     WithdrawCommandResponseMapper
	WithdrawStatsStatusResponseMapper WithdrawStatsStatusResponseMapper
	WithdrawStatsAmountResponseMapper WithdrawStatsAmountResponseMapper
}

// NewWithdrawResponseMapper creates and returns a new instance of withdrawResponseMapper.
// This instance consolidates various response mapping methods, including those for
// query responses, command responses, status statistics, and amount statistics,
// enabling conversion of WithdrawRecord domain models into API-compatible
// WithdrawResponse objects.
func NewWithdrawResponseMapper() *withdrawResponseMapper {
	return &withdrawResponseMapper{
		WithdrawQueryResponseMapper:       NewWithdrawQueryResponseMapper(),
		WithdrawCommandResponseMapper:     NewWithdrawCommandResponseMapper(),
		WithdrawStatsStatusResponseMapper: NewWithdrawStatsStatusResponseMapper(),
		WithdrawStatsAmountResponseMapper: NewWithdrawStatsAmountResponseMapper(),
	}
}
