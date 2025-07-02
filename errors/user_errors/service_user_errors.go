package user_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrUserNotFoundRes is returned when a user is not found.
var ErrUserNotFoundRes = response.NewErrorResponse("User not found", http.StatusNotFound)

// ErrUserEmailAlready is returned when a user email already exists.
var ErrUserEmailAlready = response.NewErrorResponse("User email already exists", http.StatusBadRequest)

// ErrUserPassword is returned when there is an invalid password.
var ErrUserPassword = response.NewErrorResponse("Failed invalid password", http.StatusBadRequest)

// ErrFailedPasswordNoMatch is returned when passwords do not match.
var ErrFailedPasswordNoMatch = response.NewErrorResponse("Failed password not match", http.StatusBadRequest)

// ErrFailedFindAll is returned when fetching users fails.
var ErrFailedFindAll = response.NewErrorResponse("Failed to fetch users", http.StatusInternalServerError)

// ErrFailedFindActive is returned when fetching active users fails.
var ErrFailedFindActive = response.NewErrorResponse("Failed to fetch active users", http.StatusInternalServerError)

// ErrFailedFindTrashed is returned when fetching trashed users fails.
var ErrFailedFindTrashed = response.NewErrorResponse("Failed to fetch trashed users", http.StatusInternalServerError)

// ErrInternalServerError is a generic internal server error.
var ErrInternalServerError = response.NewErrorResponse("Internal server error", http.StatusInternalServerError)

// ErrFailedSendEmail is returned when sending an email fails.
var ErrFailedSendEmail = response.NewErrorResponse("Failed to send email", http.StatusInternalServerError)

// ErrFailedCreateUser is returned when creating a user fails.
var ErrFailedCreateUser = response.NewErrorResponse("Failed to create user", http.StatusInternalServerError)

// ErrFailedUpdateUser is returned when updating a user fails.
var ErrFailedUpdateUser = response.NewErrorResponse("Failed to update user", http.StatusInternalServerError)

// ErrFailedTrashedUser is returned when moving a user to trash fails.
var ErrFailedTrashedUser = response.NewErrorResponse("Failed to move user to trash", http.StatusInternalServerError)

// ErrFailedRestoreUser is returned when restoring a user fails.
var ErrFailedRestoreUser = response.NewErrorResponse("Failed to restore user", http.StatusInternalServerError)

// ErrFailedDeletePermanent is returned when permanently deleting a user fails.
var ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete user permanently", http.StatusInternalServerError)

// ErrFailedRestoreAll is returned when restoring all users fails.
var ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all users", http.StatusInternalServerError)

// ErrFailedDeleteAll is returned when permanently deleting all users fails.
var ErrFailedDeleteAll = response.NewErrorResponse("Failed to delete all users permanently", http.StatusInternalServerError)
