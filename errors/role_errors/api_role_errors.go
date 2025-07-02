package role_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiRoleNotFound returns an API error response when the requested Role is not found.
var ErrApiRoleNotFound = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Role not found", http.StatusNotFound)
}

// ErrApiRoleInvalidId returns an API error response for an invalid Role ID.
var ErrApiRoleInvalidId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid Role id", http.StatusNotFound)
}

// ErrApiFailedFindAll returns an API error response when fetching all Roles fails.
var ErrApiFailedFindAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch Roles", http.StatusInternalServerError)
}

// ErrApiFailedFindActive returns an API error response when fetching active Roles fails.
var ErrApiFailedFindActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch active Roles", http.StatusInternalServerError)
}

// ErrApiFailedFindTrashed returns an API error response when fetching trashed Roles fails.
var ErrApiFailedFindTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to fetch trashed Roles", http.StatusInternalServerError)
}

// ErrApiFailedCreateRole returns an API error response when Role creation fails.
var ErrApiFailedCreateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to create Role", http.StatusInternalServerError)
}

// ErrApiFailedUpdateRole returns an API error response when Role update fails.
var ErrApiFailedUpdateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to update Role", http.StatusInternalServerError)
}

// ErrApiValidateCreateRole returns an API error response for invalid Role creation request validation.
var ErrApiValidateCreateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Role request", http.StatusBadRequest)
}

// ErrApiValidateUpdateRole returns an API error response for invalid Role update request validation.
var ErrApiValidateUpdateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Role request", http.StatusBadRequest)
}

// ErrInvalidRoleId returns an API error response for an invalid Role ID format.
var ErrInvalidRoleId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "invalid Role id", http.StatusBadRequest)
}

// ErrApiBindCreateRole returns an API error response for failed Role creation request binding.
var ErrApiBindCreateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Role request", http.StatusBadRequest)
}

// ErrApiBindUpdateRole returns an API error response for failed Role update request binding.
var ErrApiBindUpdateRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Role request", http.StatusBadRequest)
}

// ErrApiFailedTrashedRole returns an API error response when soft-deleting (trashing) a Role fails.
var ErrApiFailedTrashedRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to move Role to trash", http.StatusInternalServerError)
}

// ErrApiFailedRestoreRole returns an API error response when restoring a trashed Role fails.
var ErrApiFailedRestoreRole = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore Role", http.StatusInternalServerError)
}

// ErrApiFailedDeletePermanent returns an API error response when permanently deleting a Role fails.
var ErrApiFailedDeletePermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete Role permanently", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAll returns an API error response when restoring all trashed Roles fails.
var ErrApiFailedRestoreAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to restore all Roles", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAll returns an API error response when permanently deleting all Roles fails.
var ErrApiFailedDeleteAll = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "failed to delete all Roles permanently", http.StatusInternalServerError)
}
