package merchantprotomapper

type merchantProtoMapper struct {
	MerchantQueryProtoMapper            MerchantQueryProtoMapper
	MerchantCommandProtoMapper          MerchantCommandProtoMapper
	MerchantStatsAmountProtoMapper      MerchantStatsAmountProtoMapper
	MerchantStatsMethodProtoMapper      MerchantStatsMethodProtoMapper
	MerchantStatsTotalAmountProtoMapper MerchantStatsTotalAmountProtoMapper
}

func NewMerchantProtoMapper() *merchantProtoMapper {
	return &merchantProtoMapper{
		MerchantQueryProtoMapper:            NewMerchantQueryProtoMapper(),
		MerchantCommandProtoMapper:          NewMerchantCommandProtoMapper(),
		MerchantStatsAmountProtoMapper:      NewMerchantStatsAmountProtoMapper(),
		MerchantStatsMethodProtoMapper:      NewMerchantStatsMethodProtoMapper(),
		MerchantStatsTotalAmountProtoMapper: NewMerchantStatsTotalAmount(),
	}
}
