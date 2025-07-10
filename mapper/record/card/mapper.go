package cardrecordmapper

// cardRecordMapper provides methods to map Card database rows to CardRecord domain models.
type cardRecordMapper struct {
	CardQueryRecordMapper   CardQueryRecordMapper
	CardCommandRecordMapper CardCommandRecordMapper
}

// NewCardRecordMapper creates a new instance of cardRecordMapper, initializing both
// the CardQueryRecordMapper and CardCommandRecordMapper with their respective
// dependencies. This function returns a pointer to cardRecordMapper, which provides
// methods to map Card database rows to CardRecord domain models for both query and
// command operations.
func NewCardRecordMapper() *cardRecordMapper {
	return &cardRecordMapper{
		CardQueryRecordMapper:   NewCardQueryRecordMapper(),
		CardCommandRecordMapper: NewCardCommandRecordMapper(),
	}
}
