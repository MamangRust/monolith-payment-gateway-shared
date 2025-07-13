package cardapimapper

type cardResponseMapper struct {
	CardQueryResponseMapper        CardQueryResponseMapper
	CardCommandResponseMapper      CardCommandResponseMapper
	CardStatsAmountResponseMapper  CardStatsAmountResponseMapper
	CardStatsBalanceResponseMapper CardStatsBalanceResponseMapper
	CardDashboardResponseMapper    CardDashboardResponseMapper
}

func NewCardResponseMapper() *cardResponseMapper {
	return &cardResponseMapper{
		CardQueryResponseMapper:        NewCardQueryResponseMapper(),
		CardCommandResponseMapper:      NewCardCommandResponseMapper(),
		CardStatsAmountResponseMapper:  NewCardStatsAmountResponseMapper(),
		CardStatsBalanceResponseMapper: NewCardStatsBalanceResponseMapper(),
		CardDashboardResponseMapper:    NewCardDashboardResponseMapper(),
	}
}
