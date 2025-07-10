package userrecord

// userRecordMapper provides methods to map User database rows to UserRecord domain models.
type userRecordMapper struct {
	UserQueryRecordMapper   UserQueryRecordMapper
	UserCommandRecordMapper UserCommandRecordMapper
}

// NewUserRecordMapper creates a new instance of userRecordMapper with the provided
// UserQueryRecordMapper and UserCommandRecordMapper. This mapper is responsible for
// mapping User database rows to UserRecord domain models, and can be used for both
// query and command operations.
func NewUserRecordMapper() *userRecordMapper {
	return &userRecordMapper{
		UserQueryRecordMapper:   NewUserQueryRecordMapper(),
		UserCommandRecordMapper: NewUserCommandRecordMapper(),
	}
}
