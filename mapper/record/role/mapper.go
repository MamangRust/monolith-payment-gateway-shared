package rolerecord

// roleRecordMapper provides methods to map Role database rows to RoleRecord domain models.
type roleRecordMapper struct {
	RoleQueryRecordMapping   RoleQueryRecordMapping
	RoleCommandRecordMapping RoleCommandRecordMapping
}

// NewRoleRecordMapper returns a new instance of roleRecordMapper which provides methods to map Role database rows to RoleRecord domain models.
func NewRoleRecordMapper() *roleRecordMapper {
	return &roleRecordMapper{
		RoleQueryRecordMapping:   NewRoleQueryRecordMapping(),
		RoleCommandRecordMapping: NewRoleCommandRecordMapping(),
	}
}
