package role_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrRoleNotFoundRes is returned when the requested role is not found.
var ErrRoleNotFoundRes = response.NewErrorResponse("Role not found", http.StatusNotFound)

// ErrFailedFindAll is returned when there is a failure in fetching all roles.
var ErrFailedFindAll = response.NewErrorResponse("Failed to fetch Roles", http.StatusInternalServerError)

// ErrFailedFindActive is returned when there is a failure in fetching active roles.
var ErrFailedFindActive = response.NewErrorResponse("Failed to fetch active Roles", http.StatusInternalServerError)

// ErrFailedFindTrashed is returned when there is a failure in fetching trashed roles.
var ErrFailedFindTrashed = response.NewErrorResponse("Failed to fetch trashed Roles", http.StatusInternalServerError)

// ErrFailedCreateRole is returned when there is a failure in creating a role.
var ErrFailedCreateRole = response.NewErrorResponse("Failed to create Role", http.StatusInternalServerError)

// ErrFailedUpdateRole is returned when there is a failure in updating a role.
var ErrFailedUpdateRole = response.NewErrorResponse("Failed to update Role", http.StatusInternalServerError)

// ErrFailedTrashedRole is returned when there is a failure in moving a role to trash.
var ErrFailedTrashedRole = response.NewErrorResponse("Failed to move Role to trash", http.StatusInternalServerError)

// ErrFailedRestoreRole is returned when there is a failure in restoring a trashed role.
var ErrFailedRestoreRole = response.NewErrorResponse("Failed to restore Role", http.StatusInternalServerError)

// ErrFailedDeletePermanent is returned when there is a failure in permanently deleting a role.
var ErrFailedDeletePermanent = response.NewErrorResponse("Failed to delete Role permanently", http.StatusInternalServerError)

// ErrFailedRestoreAll is returned when there is a failure in restoring all trashed roles.
var ErrFailedRestoreAll = response.NewErrorResponse("Failed to restore all Roles", http.StatusInternalServerError)

// ErrFailedDeleteAll is returned when there is a failure in permanently deleting all roles.
var ErrFailedDeleteAll = response.NewErrorResponse("Failed to delete all Roles permanently", http.StatusInternalServerError)
