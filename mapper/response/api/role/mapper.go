package roleapimapper

type roleResponseMapper struct {
	RoleQueryResponseMapper   RoleQueryResponseMapper
	RoleCommandResponseMapper RoleCommandResponseMapper
}

func NewRoleResponseMapper() *roleResponseMapper {
	return &roleResponseMapper{
		RoleQueryResponseMapper:   NewRoleQueryResponseMapper(),
		RoleCommandResponseMapper: NewRoleCommandResponseMapper(),
	}
}
