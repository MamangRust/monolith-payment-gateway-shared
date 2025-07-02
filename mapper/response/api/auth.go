package apimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// authResponseMapper provides methods to map the auth response from the domain
// to the api response.
type authResponseMapper struct {
}

// NewAuthResponseMapper creates a new instance of authResponseMapper.
//
// This function is used to create a new instance of authResponseMapper which
// is used to map the auth response from the domain to the api response.
func NewAuthResponseMapper() *authResponseMapper {
	return &authResponseMapper{}
}

// ToResponseVerifyCode maps the ApiResponseVerifyCode from the domain to the
// ApiResponseVerifyCode of the api.
func (s *authResponseMapper) ToResponseVerifyCode(res *pb.ApiResponseVerifyCode) *response.ApiResponseVerifyCode {
	return &response.ApiResponseVerifyCode{
		Status:  res.Status,
		Message: res.Message,
	}
}

// ToResponseForgotPassword maps the ApiResponseForgotPassword from the domain to the
// ApiResponseForgotPassword of the api.
func (s *authResponseMapper) ToResponseForgotPassword(res *pb.ApiResponseForgotPassword) *response.ApiResponseForgotPassword {
	return &response.ApiResponseForgotPassword{
		Status:  res.Status,
		Message: res.Message,
	}
}

// ToResponseResetPassword maps the ApiResponseResetPassword from the domain to the
// ApiResponseResetPassword of the api.
func (s *authResponseMapper) ToResponseResetPassword(res *pb.ApiResponseResetPassword) *response.ApiResponseResetPassword {
	return &response.ApiResponseResetPassword{
		Status:  res.Status,
		Message: res.Message,
	}
}

// ToResponseLogin maps the ApiResponseLogin from the domain to the ApiResponseLogin of the api.
func (s *authResponseMapper) ToResponseLogin(res *pb.ApiResponseLogin) *response.ApiResponseLogin {
	return &response.ApiResponseLogin{
		Status:  res.Status,
		Message: res.Message,
		Data: &response.TokenResponse{
			AccessToken:  res.Data.AccessToken,
			RefreshToken: res.Data.RefreshToken,
		},
	}
}

// ToResponseRegister maps the ApiResponseRegister from the domain to the
// ApiResponseRegister of the api.
func (s *authResponseMapper) ToResponseRegister(res *pb.ApiResponseRegister) *response.ApiResponseRegister {
	return &response.ApiResponseRegister{
		Status:  res.Status,
		Message: res.Message,
		Data: &response.UserResponse{
			ID:        int(res.Data.Id),
			FirstName: res.Data.Firstname,
			LastName:  res.Data.Lastname,
			Email:     res.Data.Email,
			CreatedAt: res.Data.CreatedAt,
			UpdatedAt: res.Data.UpdatedAt,
		},
	}
}

// ToResponseRefreshToken maps the ApiResponseRefreshToken from the domain to the
// ApiResponseRefreshToken of the api.
//
// This function takes a pointer to a pb.ApiResponseRefreshToken and returns a
// pointer to a response.ApiResponseRefreshToken, which includes status, message,
// and token data such as access and refresh tokens.
func (s *authResponseMapper) ToResponseRefreshToken(res *pb.ApiResponseRefreshToken) *response.ApiResponseRefreshToken {
	return &response.ApiResponseRefreshToken{
		Status:  res.Status,
		Message: res.Message,
		Data: &response.TokenResponse{
			AccessToken:  res.Data.AccessToken,
			RefreshToken: res.Data.RefreshToken,
		},
	}
}

// ToResponseGetMe maps the ApiResponseGetMe from the domain to the ApiResponseGetMe of the api.
//
// Args:
//   - res: A pointer to a pb.ApiResponseGetMe representing the ApiResponseGetMe from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseGetMe containing the mapped data, including status, message,
//     and user data such as ID, FirstName, LastName, Email, CreatedAt, and UpdatedAt.
func (s *authResponseMapper) ToResponseGetMe(res *pb.ApiResponseGetMe) *response.ApiResponseGetMe {
	return &response.ApiResponseGetMe{
		Status:  res.Status,
		Message: res.Message,
		Data: &response.UserResponse{
			ID:        int(res.Data.Id),
			FirstName: res.Data.Firstname,
			LastName:  res.Data.Lastname,
			Email:     res.Data.Email,
			CreatedAt: res.Data.CreatedAt,
			UpdatedAt: res.Data.UpdatedAt,
		},
	}
}
