package roleservicemapper

// roleResponseMapper provides methods to map RoleRecord domain models to RoleResponse API-compatible response types.
type roleResponseMapper struct {
	RoleQueryResponseMapper   RoleQueryResponseMapper
	RoleCommandResponseMapper RoleCommandResponseMapper
}

// NewRoleResponseMapper returns a new instance of roleResponseMapper, which provides methods to map RoleRecord domain models to RoleResponse API-compatible response types.
func NewRoleResponseMapper() *roleResponseMapper {
	return &roleResponseMapper{
		RoleQueryResponseMapper:   NewRoleQueryResponseMapper(),
		RoleCommandResponseMapper: NewRoleCommandResponseMapper(),
	}
}
