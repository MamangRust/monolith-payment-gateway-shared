package userapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/user"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type userCommandResponseMapper struct {
}

func NewUserCommandResponseMapper() UserCommandResponseMapper {
	return &userCommandResponseMapper{}
}

// ToApiResponseUser converts a single user response into an API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseUser containing the user data.
//
// Returns:
//   - A pointer to a response.ApiResponseUser containing the mapped user data, including
//     Status, Message, and Data.
func (u *userCommandResponseMapper) ToApiResponseUser(pbResponse *pb.ApiResponseUser) *response.ApiResponseUser {
	return &response.ApiResponseUser{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    u.ToResponseUser(pbResponse.Data),
	}
}

// ToResponseUser maps a protobuf UserResponse to a domain UserResponse.
//
// Args:
//   - user: A pointer to a pb.UserResponse containing the user data.
//
// Returns:
//   - A pointer to a response.UserResponse containing the mapped user data, including
//     ID, FirstName, LastName, Email, CreatedAt, and UpdatedAt.
func (u *userCommandResponseMapper) ToResponseUser(user *pb.UserResponse) *response.UserResponse {
	return &response.UserResponse{
		ID:        int(user.Id),
		FirstName: user.Firstname,
		LastName:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
