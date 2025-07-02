package responseservice

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// userResponseMapper provides methods to map UserRecord domain models to UserResponse API-compatible response types.
type userResponseMapper struct {
}

// NewUserResponseMapper creates and returns a new instance of userResponseMapper.
// This struct provides methods for converting UserRecord domain models into
// API-compatible UserResponse types.
func NewUserResponseMapper() *userResponseMapper {
	return &userResponseMapper{}
}

// ToUserResponse maps a UserRecord domain model to a UserResponse API-compatible format.
//
// Args:
//   - user: A pointer to a UserRecord containing the user data to be mapped.
//
// Returns:
//   - A pointer to a UserResponse with the mapped user data, including ID, FirstName, LastName,
//     Email, IsVerified, CreatedAt, and UpdatedAt.
func (s *userResponseMapper) ToUserResponse(user *record.UserRecord) *response.UserResponse {
	return &response.UserResponse{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}

// ToUsersResponse converts a slice of UserRecord domain models into a slice of UserResponse API-compatible formats.
//
// Args:
//   - users: A slice of pointers to UserRecord containing the user data to be mapped.
//
// Returns:
//   - A slice of pointers to UserResponse with the mapped user data for each user.
func (s *userResponseMapper) ToUsersResponse(users []*record.UserRecord) []*response.UserResponse {
	var responses []*response.UserResponse

	for _, user := range users {
		responses = append(responses, s.ToUserResponse(user))
	}

	return responses
}

// ToUserResponseDeleteAt converts a UserRecord domain model into a UserResponseDeleteAt API-compatible format.
//
// Args:
//   - user: A pointer to a UserRecord containing the user data to be mapped.
//
// Returns:
//   - A pointer to a UserResponseDeleteAt with the mapped user data, including ID, FirstName, LastName,
//     Email, IsVerified, CreatedAt, UpdatedAt, and DeletedAt.
func (s *userResponseMapper) ToUserResponseDeleteAt(user *record.UserRecord) *response.UserResponseDeleteAt {
	return &response.UserResponseDeleteAt{
		ID:         user.ID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		IsVerified: user.IsVerified,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		DeletedAt:  user.DeletedAt,
	}
}

// ToUsersResponseDeleteAt converts a slice of UserRecord domain models into a slice of UserResponseDeleteAt API-compatible formats.
//
// Args:
//   - users: A slice of pointers to UserRecord containing the user data to be mapped.
//
// Returns:
//   - A slice of pointers to UserResponseDeleteAt with the mapped user data for each user.
func (s *userResponseMapper) ToUsersResponseDeleteAt(users []*record.UserRecord) []*response.UserResponseDeleteAt {
	var responses []*response.UserResponseDeleteAt

	for _, user := range users {
		responses = append(responses, s.ToUserResponseDeleteAt(user))
	}

	return responses
}
