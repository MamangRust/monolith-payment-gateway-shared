package apimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// roleResponseMapper maps between gRPC role-related responses and HTTP-compatible API response formats
type roleResponseMapper struct {
}

// NewRoleResponseMapper creates a new instance of roleResponseMapper, which is used to map between
// gRPC role-related responses and HTTP-compatible API response formats used in the presentation layer.
func NewRoleResponseMapper() *roleResponseMapper {
	return &roleResponseMapper{}
}

// ToApiResponseRoleAll maps a gRPC response containing all roles to an HTTP API response format.
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseRoleAll containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseRoleAll containing the mapped data, including status and message.
func (s *roleResponseMapper) ToApiResponseRoleAll(pbResponse *pb.ApiResponseRoleAll) *response.ApiResponseRoleAll {
	return &response.ApiResponseRoleAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseRoleDelete maps a gRPC delete role response to an HTTP API response format.
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseRoleDelete containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseRoleDelete containing the mapped status and message.
func (s *roleResponseMapper) ToApiResponseRoleDelete(pbResponse *pb.ApiResponseRoleDelete) *response.ApiResponseRoleDelete {
	return &response.ApiResponseRoleDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseRole maps a single gRPC role response into an HTTP API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseRole containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseRole containing the mapped status, message, and data.
func (s *roleResponseMapper) ToApiResponseRole(pbResponse *pb.ApiResponseRole) *response.ApiResponseRole {
	return &response.ApiResponseRole{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseRole(pbResponse.Data),
	}
}

// ToApiResponsesRole maps a gRPC response containing multiple roles
// into a list HTTP API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsesRole containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsesRole containing the mapped data, including status, message,
//     and data.
func (s *roleResponseMapper) ToApiResponsesRole(pbResponse *pb.ApiResponsesRole) *response.ApiResponsesRole {
	return &response.ApiResponsesRole{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesRole(pbResponse.Data),
	}
}

// ToApiResponsePaginationRole maps a paginated gRPC response of roles
// into a paginated HTTP API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationRole containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationRole containing the mapped data, including status,
//     message, data, and pagination metadata.
func (s *roleResponseMapper) ToApiResponsePaginationRole(pbResponse *pb.ApiResponsePaginationRole) *response.ApiResponsePaginationRole {
	return &response.ApiResponsePaginationRole{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesRole(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

// ToApiResponsePaginationRoleDeleteAt maps a paginated gRPC response of soft-deleted roles
// into a paginated HTTP API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponsePaginationRoleDeleteAt containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponsePaginationRoleDeleteAt containing the mapped data, including status,
//     message, data, and pagination metadata.
func (s *roleResponseMapper) ToApiResponsePaginationRoleDeleteAt(pbResponse *pb.ApiResponsePaginationRoleDeleteAt) *response.ApiResponsePaginationRoleDeleteAt {
	return &response.ApiResponsePaginationRoleDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesRoleDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

// mapResponseRole maps a gRPC RoleResponse to an HTTP-compatible RoleResponse.
//
// Args:
//   - role: A pointer to a pb.RoleResponse containing the gRPC role response data.
//
// Returns:
//   - A pointer to a response.RoleResponse containing the mapped data, with fields ID, Name,
//     CreatedAt, and UpdatedAt extracted from the gRPC response.
func (s *roleResponseMapper) mapResponseRole(role *pb.RoleResponse) *response.RoleResponse {
	return &response.RoleResponse{
		ID:        int(role.Id),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}

// mapResponsesRole maps a slice of gRPC RoleResponse to a slice of HTTP-compatible RoleResponse.
//
// Args:
//   - roles: A slice of pointers to pb.RoleResponse containing the gRPC role response data.
//
// Returns:
//   - A slice of pointers to response.RoleResponse containing the mapped data, with fields ID, Name,
//     CreatedAt, and UpdatedAt extracted from the gRPC response.
func (s *roleResponseMapper) mapResponsesRole(roles []*pb.RoleResponse) []*response.RoleResponse {
	var responseRoles []*response.RoleResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRole(role))
	}

	return responseRoles
}

// mapResponseRoleDeleteAt maps a gRPC RoleResponseDeleteAt to an HTTP-compatible RoleResponseDeleteAt.
//
// Args:
//   - role: A pointer to a pb.RoleResponseDeleteAt containing the gRPC role response data including deletion info.
//
// Returns:
//   - A pointer to a response.RoleResponseDeleteAt containing the mapped data, with fields ID, Name,
//     CreatedAt, UpdatedAt, and DeletedAt extracted from the gRPC response. If DeletedAt is present,
//     it is mapped to a string; otherwise, it is set to nil.
func (s *roleResponseMapper) mapResponseRoleDeleteAt(role *pb.RoleResponseDeleteAt) *response.RoleResponseDeleteAt {
	var deletedAt string
	if role.DeletedAt != nil {
		deletedAt = role.DeletedAt.Value
	}

	return &response.RoleResponseDeleteAt{
		ID:        int(role.Id),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
		DeletedAt: &deletedAt,
	}
}

// mapResponsesRoleDeleteAt maps a slice of gRPC RoleResponseDeleteAt to a slice of HTTP-compatible
// RoleResponseDeleteAt.
//
// Args:
//   - roles: A slice of pointers to pb.RoleResponseDeleteAt containing the gRPC role response data
//     including deletion info.
//
// Returns:
//   - A slice of pointers to response.RoleResponseDeleteAt containing the mapped data, with fields ID,
//     Name, CreatedAt, UpdatedAt, and DeletedAt extracted from the gRPC response. If DeletedAt is
//     present, it is mapped to a string; otherwise, it is set to nil.
func (s *roleResponseMapper) mapResponsesRoleDeleteAt(roles []*pb.RoleResponseDeleteAt) []*response.RoleResponseDeleteAt {
	var responseRoles []*response.RoleResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapResponseRoleDeleteAt(role))
	}

	return responseRoles
}
