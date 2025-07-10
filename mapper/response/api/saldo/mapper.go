package saldoapimapper

type saldoResponseMapper struct {
	SaldoQueryResponseMapper        SaldoQueryResponseMapper
	SaldoCommandResponseMapper      SaldoCommandResponseMapper
	SaldoStatsBalanceResponseMapper SaldoStatsBalanceResponseMapper
	SaldoStatsTotalResponseMapper   SaldoStatsTotalResponseMapper
}

func NewSaldoResponseMapper() *saldoResponseMapper {
	return &saldoResponseMapper{
		SaldoQueryResponseMapper:        NewSaldoQueryResponseMapper(),
		SaldoCommandResponseMapper:      NewSaldoCommandResponseMapper(),
		SaldoStatsBalanceResponseMapper: NewSaldoStatsBalanceResponseMapper(),
		SaldoStatsTotalResponseMapper:   NewSaldoStatsTotalResponseMapper(),
	}
}
