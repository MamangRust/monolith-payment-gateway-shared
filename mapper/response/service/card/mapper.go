package cardservicemapper

// cardResponseMapper aggregates multiple response mappers including query, command, statistic amount, and statistic balance mappers.
type cardResponseMapper struct {
	CardQueryResponseMapper            CardQueryResponseMapper
	CardCommandResponseMapper          CardCommandResponseMapper
	CardStatisticAmountResponseMapper  CardStatisticAmountResponseMapper
	CardStatisticBalanceResponseMapper CardStatisticBalanceResponseMapper
}


// NewCardResponseMapper creates and returns a new instance of cardResponseMapper, which aggregates multiple
// response mappers including query, command, statistic amount, and statistic balance mappers. These mappers
// provide methods for converting domain models to response models for various card-related operations.
func NewCardResponseMapper() *cardResponseMapper {
	return &cardResponseMapper{
		CardQueryResponseMapper:            NewCardQueryResponseMapper(),
		CardCommandResponseMapper:          NewCardCommandResponseMapper(),
		CardStatisticAmountResponseMapper:  NewCardStatsAmountResponseMapper(),
		CardStatisticBalanceResponseMapper: NewCardStatsBalanceResponseMapper(),
	}
}
