package role_errors

import "errors"

// ErrRoleNotFound indicates that the requested Role was not found in the database.
var ErrRoleNotFound = errors.New("role not found")

// ErrFindAllRoles is returned when retrieving all Roles from the database fails.
var ErrFindAllRoles = errors.New("failed to find all Roles")

// ErrFindActiveRoles is returned when retrieving all active Roles fails.
var ErrFindActiveRoles = errors.New("failed to find active Roles")

// ErrFindTrashedRoles is returned when retrieving trashed (soft-deleted) Roles fails.
var ErrFindTrashedRoles = errors.New("failed to find trashed Roles")

// ErrRoleConflict indicates a conflict where a Role already exists.
var ErrRoleConflict = errors.New("failed Role already exists")

// ErrCreateRole is returned when creating a new Role fails.
var ErrCreateRole = errors.New("failed to create Role")

// ErrUpdateRole is returned when updating an existing Role fails.
var ErrUpdateRole = errors.New("failed to update Role")

// ErrTrashedRole is returned when moving a Role to trash (soft-delete) fails.
var ErrTrashedRole = errors.New("failed to move Role to trash")

// ErrRestoreRole is returned when restoring a trashed Role fails.
var ErrRestoreRole = errors.New("failed to restore Role from trash")

// ErrDeleteRolePermanent is returned when permanently deleting a Role fails.
var ErrDeleteRolePermanent = errors.New("failed to permanently delete Role")

// ErrRestoreAllRoles is returned when restoring all trashed Roles fails.
var ErrRestoreAllRoles = errors.New("failed to restore all Roles")

// ErrDeleteAllRoles is returned when permanently deleting all Roles fails.
var ErrDeleteAllRoles = errors.New("failed to permanently delete all Roles")
