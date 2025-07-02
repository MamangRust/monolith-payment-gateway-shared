package user_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiUserNotFound is an error response for when the user is not found.
var ErrApiUserNotFound = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "User not found", http.StatusNotFound)
}

// ErrApiUserInvalidId is an error response for an invalid user ID.
var ErrApiUserInvalidId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid User id", http.StatusNotFound)
}

// ErrApiFailedFindAll is an error response for failing to fetch all users.
var ErrApiFailedFindAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch Users", http.StatusInternalServerError)
}

// ErrApiFailedFindActive is an error response for failing to fetch active users.
var ErrApiFailedFindActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch active Users", http.StatusInternalServerError)
}

// ErrApiFailedFindTrashed is an error response for failing to fetch trashed users.
var ErrApiFailedFindTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Users", http.StatusInternalServerError)
}

// ErrApiFailedCreateUser is an error response for failing to create a user.
var ErrApiFailedCreateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to create User", http.StatusInternalServerError)
}

// ErrApiFailedUpdateUser is an error response for failing to update a user.
var ErrApiFailedUpdateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to update User", http.StatusInternalServerError)
}

// ErrApiValidateCreateUser is an error response for validation failure in creating a user.
var ErrApiValidateCreateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create User request", http.StatusBadRequest)
}

// ErrApiValidateUpdateUser is an error response for validation failure in updating a user.
var ErrApiValidateUpdateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update User request", http.StatusBadRequest)
}

// ErrInvalidUserId is an error response for an invalid user ID.
var ErrInvalidUserId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid User id", http.StatusBadRequest)
}

// ErrApiBindCreateUser is an error response for bind failure: invalid create user request.
var ErrApiBindCreateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create User request", http.StatusBadRequest)
}

// ErrApiBindUpdateUser is an error response for bind failure: invalid update user request.
var ErrApiBindUpdateUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update User request", http.StatusBadRequest)
}

// ErrApiFailedTrashedUser is an error response for failed to move user to trash.
var ErrApiFailedTrashedUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to move User to trash", http.StatusInternalServerError)
}

// ErrApiFailedRestoreUser is an error response for failed to restore user.
var ErrApiFailedRestoreUser = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore User", http.StatusInternalServerError)
}

// ErrApiFailedDeletePermanent is an error response for failed to delete user permanently.
var ErrApiFailedDeletePermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete User permanently", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAll is an error response for failed to restore all users.
var ErrApiFailedRestoreAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore all Users", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAll is an error response for failed to permanently delete all users.
var ErrApiFailedDeleteAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete all Users permanently", http.StatusInternalServerError)
}
