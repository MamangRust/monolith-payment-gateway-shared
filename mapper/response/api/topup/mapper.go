package topupapimapper

type topupResponseMapper struct {
	TopupQueryResponseMapper       TopupQueryResponseMapper
	TopupCommandResponseMapper     TopupCommandResponseMapper
	TopupStatsStatusResponseMapper TopupStatsStatusResponseMapper
	TopupStatsAmountResponseMapper TopupStatsAmountResponseMapper
	TopupStatsMethodResponseMapper TopupStatsMethodResponseMapper
}

func NewTopupResponseMapper() *topupResponseMapper {
	return &topupResponseMapper{
		TopupQueryResponseMapper:       NewTopupQueryResponseMapper(),
		TopupCommandResponseMapper:     NewTopupCommandResponseMapper(),
		TopupStatsStatusResponseMapper: NewTopupStatsStatusResponseMapper(),
		TopupStatsAmountResponseMapper: NewTopupStatsAmountResponseMapper(),
		TopupStatsMethodResponseMapper: NewTopupStatsMethodResponseMapper(),
	}
}
