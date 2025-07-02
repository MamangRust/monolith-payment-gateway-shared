package refreshtoken_errors

import "errors"

// ErrTokenNotFound indicates that the refresh token could not be found.
var ErrTokenNotFound = errors.New("refresh token not found")

// ErrFindByToken is returned when a lookup for the refresh token by token value fails.
var ErrFindByToken = errors.New("failed to find refresh token by token")

// ErrFindByUserID is returned when a lookup for the refresh token by user ID fails.
var ErrFindByUserID = errors.New("failed to find refresh token by user ID")

// ErrCreateRefreshToken indicates that an error occurred while creating a new refresh token.
var ErrCreateRefreshToken = errors.New("failed to create refresh token")

// ErrUpdateRefreshToken is returned when the refresh token update process fails.
var ErrUpdateRefreshToken = errors.New("failed to update refresh token")

// ErrDeleteRefreshToken is returned when deleting a refresh token fails.
var ErrDeleteRefreshToken = errors.New("failed to delete refresh token")

// ErrDeleteByUserID indicates a failure when attempting to delete a refresh token using the user ID.
var ErrDeleteByUserID = errors.New("failed to delete refresh token by user ID")

// ErrParseDate is returned when parsing the expiration date of a token fails.
var ErrParseDate = errors.New("failed to parse expiration date")
