package userapimapper

type userResponseMapper struct {
	UserQueryResponseMapper   UserQueryResponseMapper
	UserCommandResponseMapper UserCommandResponseMapper
}

func NewUserResponseMapper() *userResponseMapper {
	return &userResponseMapper{
		UserQueryResponseMapper:   NewUserQueryResponseMapper(),
		UserCommandResponseMapper: NewUserCommandResponseMapper(),
	}
}
