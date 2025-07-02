package protomapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

// UserProtoMapper provides methods for mapping domain models to pb.UserResponse
type userProtoMapper struct {
}

// NewUserProtoMapper returns a new instance of UserProtoMapper
func NewUserProtoMapper() *userProtoMapper {
	return &userProtoMapper{}
}

// ToProtoResponseUser maps a UserResponse to a pb.UserResponse
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//   - pbResponse: A pointer to the UserResponse to be mapped.
//
// Returns:
//   - A pointer to a pb.UserResponse containing the mapped data.
//   - If pbResponse is nil, returns nil.
func (u *userProtoMapper) ToProtoResponseUser(status string, message string, pbResponse *response.UserResponse) *pb.ApiResponseUser {
	return &pb.ApiResponseUser{
		Status:  status,
		Message: message,
		Data:    u.mapResponseUser(pbResponse),
	}
}

// ToProtoResponseUserDeleteAt maps a UserResponseDeleteAt to a pb.ApiResponseUserDeleteAt proto message.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//   - pbResponse: A pointer to the UserResponseDeleteAt to be mapped.
//
// Returns:
//   - A pointer to a pb.ApiResponseUserDeleteAt containing the mapped data.
//   - If pbResponse is nil, returns nil.
func (u *userProtoMapper) ToProtoResponseUserDeleteAt(status string, message string, pbResponse *response.UserResponseDeleteAt) *pb.ApiResponseUserDeleteAt {
	return &pb.ApiResponseUserDeleteAt{
		Status:  status,
		Message: message,
		Data:    u.mapResponseUserDelete(pbResponse),
	}
}

// ToProtoResponsesUser maps a list of UserResponse to a pb.ApiResponsesUser proto message.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//   - pbResponse: A slice of pointers to UserResponse to be mapped.
//
// Returns:
//   - A pointer to a pb.ApiResponsesUser containing the mapped data.
//   - If pbResponse is nil, returns nil.
func (u *userProtoMapper) ToProtoResponsesUser(status string, message string, pbResponse []*response.UserResponse) *pb.ApiResponsesUser {
	return &pb.ApiResponsesUser{
		Status:  status,
		Message: message,
		Data:    u.mapResponsesUser(pbResponse),
	}
}

// ToProtoResponseUserDelete maps a UserResponseDelete to a pb.ApiResponseUserDelete proto message.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//
// Returns:
//   - A pointer to a pb.ApiResponseUserDelete containing the mapped data.
func (u *userProtoMapper) ToProtoResponseUserDelete(status string, message string) *pb.ApiResponseUserDelete {
	return &pb.ApiResponseUserDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseUserAll returns a protobuf message for all users without pagination.
//
// Args:
//   - status: A string representing the status of the response.
//   - message: A string representing the message associated with the response.
//
// Returns:
//   - A pointer to a pb.ApiResponseUserAll containing the mapped data.
func (u *userProtoMapper) ToProtoResponseUserAll(status string, message string) *pb.ApiResponseUserAll {
	return &pb.ApiResponseUserAll{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponsePaginationUserDeleteAt maps a list of UserResponseDeleteAt and pagination metadata
// to a protobuf ApiResponsePaginationUserDeleteAt. It includes the status and message for the API response.
//
// Args:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//   - users: A list of UserResponseDeleteAt objects to be included in the response.
//
// Returns:
//
//	A pointer to ApiResponsePaginationUserDeleteAt containing the status, message, user data, and pagination data.
func (u *userProtoMapper) ToProtoResponsePaginationUserDeleteAt(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponseDeleteAt) *pb.ApiResponsePaginationUserDeleteAt {
	return &pb.ApiResponsePaginationUserDeleteAt{
		Status:     status,
		Message:    message,
		Data:       u.mapResponsesUserDeleteAt(users),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationUser maps a list of UserResponse and pagination metadata
// to a protobuf ApiResponsePaginationUser. It includes the status and message for the API response.
//
// Args:
//   - pagination: The pagination metadata for the response.
//   - status: The status of the API response.
//   - message: A descriptive message for the API response.
//   - users: A list of UserResponse objects to be included in the response.
//
// Returns:
//
//	A pointer to ApiResponsePaginationUser containing the status, message, user data, and pagination data.
func (u *userProtoMapper) ToProtoResponsePaginationUser(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponse) *pb.ApiResponsePaginationUser {
	return &pb.ApiResponsePaginationUser{
		Status:     status,
		Message:    message,
		Data:       u.mapResponsesUser(users),
		Pagination: mapPaginationMeta(pagination),
	}
}

// mapResponseUser maps a UserResponse to a protobuf UserResponse.
//
// Args:
//   - user: A pointer to a UserResponse to be mapped.
//
// Returns:
//   - A pointer to a protobuf UserResponse containing the mapped data.
func (u *userProtoMapper) mapResponseUser(user *response.UserResponse) *pb.UserResponse {
	return &pb.UserResponse{
		Id:        int32(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

// mapResponsesUser maps a slice of UserResponse to a slice of protobuf UserResponse.
//
// Args:
//   - users: A slice of pointers to UserResponse to be mapped.
//
// Returns:
//   - A slice of pointers to protobuf UserResponse containing the mapped data for each user.
func (u *userProtoMapper) mapResponsesUser(users []*response.UserResponse) []*pb.UserResponse {
	var mappedUsers []*pb.UserResponse

	for _, user := range users {
		mappedUsers = append(mappedUsers, u.mapResponseUser(user))
	}

	return mappedUsers
}

// mapResponseUserDelete maps a UserResponseDeleteAt to a protobuf UserResponseDeleteAt.
//
// Args:
//   - user: A pointer to a UserResponseDeleteAt containing the user's details.
//
// Returns:
//   - A pointer to a protobuf UserResponseDeleteAt containing the mapped data, including
//     ID, Firstname, Lastname, Email, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.
func (u *userProtoMapper) mapResponseUserDelete(user *response.UserResponseDeleteAt) *pb.UserResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if user.DeletedAt != nil {
		deletedAt = wrapperspb.String(*user.DeletedAt)
	}

	return &pb.UserResponseDeleteAt{
		Id:        int32(user.ID),
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: deletedAt,
	}
}

// mapResponsesUserDeleteAt maps a slice of UserResponseDeleteAt to a slice of protobuf UserResponseDeleteAt.
//
// Args:
//   - users: A slice of pointers to UserResponseDeleteAt to be mapped.
//
// Returns:
//   - A slice of pointers to protobuf UserResponseDeleteAt containing the mapped data for each user, including
//     ID, Firstname, Lastname, Email, CreatedAt, UpdatedAt, and potentially DeletedAt, if available.
func (u *userProtoMapper) mapResponsesUserDeleteAt(users []*response.UserResponseDeleteAt) []*pb.UserResponseDeleteAt {
	var mappedUsers []*pb.UserResponseDeleteAt

	for _, user := range users {
		mappedUsers = append(mappedUsers, u.mapResponseUserDelete(user))
	}

	return mappedUsers
}

// mapPaginationMeta maps a sharedPagination.PaginationMeta to a protobuf PaginationMeta.
//
// Args:
//   - s: A pointer to a sharedPagination.PaginationMeta to be mapped.
//
// Returns:
//   - A pointer to a protobuf PaginationMeta containing the mapped data.
func mapPaginationMeta(s *pb.PaginationMeta) *pb.PaginationMeta {
	return &pb.PaginationMeta{
		CurrentPage:  int32(s.CurrentPage),
		PageSize:     int32(s.PageSize),
		TotalPages:   int32(s.TotalPages),
		TotalRecords: int32(s.TotalRecords),
	}
}
