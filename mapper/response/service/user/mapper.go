package userservicemapper

// userResponseMapper is a struct that implements the UserQueryResponseMapper and UserCommandResponseMapper interfaces.
type userResponseMapper struct {
	UserQueryResponseMapper   UserQueryResponseMapper
	UserCommandResponseMapper UserCommandResponseMapper
}

// NewUserResponseMapper creates and returns a new instance of userResponseMapper.
// This mapper combines both query and command response mappers to provide a
// comprehensive tool for converting UserRecord domain models into various
// API-compatible UserResponse and UserResponseDeleteAt types.
func NewUserResponseMapper() *userResponseMapper {
	return &userResponseMapper{
		UserQueryResponseMapper:   NewUserQueryResponseMapper(),
		UserCommandResponseMapper: NewUserCommandResponseMapper(),
	}
}
