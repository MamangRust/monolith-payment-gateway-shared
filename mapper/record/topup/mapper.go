package topuprecordmapper

// topupRecordMapper provides methods to map Topup database rows to TopupRecord domain models for both query and command operations.
type topupRecordMapper struct {
	TopupQueryRecordMapping   TopupQueryRecordMapping
	TopupCommandRecordMapping TopupCommandRecordMapping
}

// NewTopupRecordMapper returns a new instance of topupRecordMapper,
// for both query and command operations. It initializes the TopupQueryRecordMapping
// and TopupCommandRecordMapping fields using their respective constructor functions.
func NewTopupRecordMapper() *topupRecordMapper {
	return &topupRecordMapper{
		TopupQueryRecordMapping:   NewTopupQueryRecordMapper(),
		TopupCommandRecordMapping: NewTopupCommandRecordMapper(),
	}
}
