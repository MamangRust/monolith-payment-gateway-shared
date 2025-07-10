package roleapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/role"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type roleCommandResponseMapper struct {
}

// NewRoleCommandResponseMapper returns a new instance of roleCommandResponseMapper, which provides methods to map RoleRecord domain models to RoleResponse API-compatible response types for command operations.
func NewRoleCommandResponseMapper() RoleCommandResponseMapper {
	return &roleCommandResponseMapper{}
}

// ToApiResponseRole maps a single gRPC role response into an HTTP API response format.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseRole containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseRole containing the mapped status, message, and data.
func (s *roleCommandResponseMapper) ToApiResponseRole(pbResponse *pb.ApiResponseRole) *response.ApiResponseRole {
	return &response.ApiResponseRole{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseRole(pbResponse.Data),
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
func (s *roleCommandResponseMapper) mapResponseRole(role *pb.RoleResponse) *response.RoleResponse {
	return &response.RoleResponse{
		ID:        int(role.Id),
		Name:      role.Name,
		CreatedAt: role.CreatedAt,
		UpdatedAt: role.UpdatedAt,
	}
}
