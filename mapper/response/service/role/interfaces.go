package roleservicemapper

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type RoleBaseResponseMapper interface {
	// Converts a single RoleRecord into RoleResponse.
	ToRoleResponse(role *record.RoleRecord) *response.RoleResponse
}

type RoleQueryResponseMapper interface {
	RoleBaseResponseMapper

	// Converts a slice of RoleRecord into a slice of RoleResponse.
	ToRolesResponse(roles []*record.RoleRecord) []*response.RoleResponse

	// Converts a deleted RoleRecord into RoleResponseDeleteAt.
	ToRoleResponseDeleteAt(role *record.RoleRecord) *response.RoleResponseDeleteAt

	// Converts a slice of deleted RoleRecord into RoleResponseDeleteAt slice.
	ToRolesResponseDeleteAt(roles []*record.RoleRecord) []*response.RoleResponseDeleteAt
}

type RoleCommandResponseMapper interface {
	RoleBaseResponseMapper
}
