package refreshtoken_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrRefreshTokenNotFound indicates that the refresh token was not found.
var ErrRefreshTokenNotFound = response.NewErrorResponse("Refresh token not found", http.StatusNotFound)

// ErrFailedExpire occurs when expiring a refresh token fails.
var ErrFailedExpire = response.NewErrorResponse("Failed to find refresh token by token", http.StatusInternalServerError)

// ErrFailedFindByToken indicates failure when searching for a refresh token by its token value.
var ErrFailedFindByToken = response.NewErrorResponse("Failed to find refresh token by token", http.StatusInternalServerError)

// ErrFailedFindByUserID indicates failure when searching for a refresh token by user ID.
var ErrFailedFindByUserID = response.NewErrorResponse("Failed to find refresh token by user ID", http.StatusInternalServerError)

// ErrFailedInValidToken is returned when an access token is invalid.
var ErrFailedInValidToken = response.NewErrorResponse("Failed to invalid access token", http.StatusInternalServerError)

// ErrFailedInValidUserId is returned when a user ID is invalid.
var ErrFailedInValidUserId = response.NewErrorResponse("Failed to invalid user id", http.StatusInternalServerError)

// ErrFailedCreateAccess is returned when the creation of an access token fails.
var ErrFailedCreateAccess = response.NewErrorResponse("Failed to create access token", http.StatusInternalServerError)

// ErrFailedCreateRefresh is returned when the creation of a refresh token fails.
var ErrFailedCreateRefresh = response.NewErrorResponse("Failed to create refresh token", http.StatusInternalServerError)

// ErrFailedCreateRefreshToken is returned when refresh token creation fails.
var ErrFailedCreateRefreshToken = response.NewErrorResponse("Failed to create refresh token", http.StatusInternalServerError)

// ErrFailedUpdateRefreshToken is returned when refresh token update fails.
var ErrFailedUpdateRefreshToken = response.NewErrorResponse("Failed to update refresh token", http.StatusInternalServerError)

// ErrFailedDeleteRefreshToken is returned when refresh token deletion fails.
var ErrFailedDeleteRefreshToken = response.NewErrorResponse("Failed to delete refresh token", http.StatusInternalServerError)

// ErrFailedDeleteByUserID is returned when deletion of a refresh token by user ID fails.
var ErrFailedDeleteByUserID = response.NewErrorResponse("Failed to delete refresh token by user ID", http.StatusInternalServerError)

// ErrFailedParseExpirationDate indicates failure when parsing the expiration date of a token.
var ErrFailedParseExpirationDate = response.NewErrorResponse("Failed to parse expiration date", http.StatusBadRequest)
