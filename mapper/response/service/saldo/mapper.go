package saldoservicemapper

// saldoResponseMapper provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types for query, command, and statistic operations.
type saldoResponseMapper struct {
	SaldoQueryResponseMapper                 SaldoQueryResponseMapper
	SaldoCommandResponseMapper               SaldoCommandResponseMapper
	SaldoStatisticBalanceResponseMapper      SaldoStatisticBalanceResponseMapper
	SaldoStatisticTotalBalanceResponseMapper SaldoStatisticTotalBalanceResponseMapper
}

// NewSaldoResponseMapper returns a new instance of SaldoResponseMapper which provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types for query, command, and statistic operations.
func NewSaldoResponseMapper() *saldoResponseMapper {
	return &saldoResponseMapper{
		SaldoQueryResponseMapper:                 NewSaldoQueryResponseMapper(),
		SaldoCommandResponseMapper:               NewSaldoCommandResponseMapper(),
		SaldoStatisticBalanceResponseMapper:      NewSaldoStatsBalanceResponseMapper(),
		SaldoStatisticTotalBalanceResponseMapper: NewSaldoTotalBalanceResponseMapper(),
	}
}
