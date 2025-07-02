package topup_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiInvalidTopupID returns a 400 Bad Request error when the provided top-up ID is invalid.
var ErrApiInvalidTopupID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid topup ID", http.StatusBadRequest)
}

// ErrApiInvalidCardNumber returns a 400 Bad Request error when the provided card number is invalid.
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card number", http.StatusBadRequest)
}

// ErrApiInvalidMonth returns a 400 Bad Request error when the provided month value is invalid or missing.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}

// ErrApiInvalidYear returns a 400 Bad Request error when the provided year value is invalid or missing.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}

// ErrApiFailedFindAllTopups returns a 500 Internal Server Error when fetching all top-up records fails.
var ErrApiFailedFindAllTopups = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all topups", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTopupsTrashed returns a 500 error when retrieving all trashed (soft-deleted) top-ups fails.
var ErrApiFailedFindAllTopupsTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all topups trashed", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTopupsActive returns a 500 error when retrieving all active (non-deleted) top-ups fails.
var ErrApiFailedFindAllTopupsActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all topups active", http.StatusInternalServerError)
}

// ErrApiFailedFindAllByCardNumberTopup returns a 500 error when retrieving top-ups by card number fails.
var ErrApiFailedFindAllByCardNumberTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch topups by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdTopup returns a 500 error when retrieving a specific top-up by ID fails.
var ErrApiFailedFindByIdTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch topup by ID", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupStatusSuccess returns a 500 error when fetching monthly successful top-up stats fails.
var ErrApiFailedFindMonthlyTopupStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful topups", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupStatusSuccess returns a 500 error when fetching yearly successful top-up stats fails.
var ErrApiFailedFindYearlyTopupStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful topups", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupStatusFailed returns a 500 error when fetching monthly failed top-up stats fails.
var ErrApiFailedFindMonthlyTopupStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed topups", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupStatusFailed returns a 500 error when fetching yearly failed top-up stats fails.
var ErrApiFailedFindYearlyTopupStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed topups", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupStatusSuccessByCardNumber returns a 500 error when fetching monthly successful top-ups by card number fails.
var ErrApiFailedFindMonthlyTopupStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful topups by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupStatusSuccessByCardNumber returns a 500 error when fetching yearly successful top-ups by card number fails.
var ErrApiFailedFindYearlyTopupStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful topups by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupStatusFailedByCardNumber returns a 500 error when fetching monthly failed topups by card number fails.
var ErrApiFailedFindMonthlyTopupStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed topups by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupStatusFailedByCardNumber returns a 500 error when fetching yearly failed topups by card number fails.
var ErrApiFailedFindYearlyTopupStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed topups by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupMethods returns a 500 error when fetching monthly topup methods fails.
var ErrApiFailedFindMonthlyTopupMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup methods", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupMethods returns a 500 error when fetching yearly topup methods fails.
var ErrApiFailedFindYearlyTopupMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup methods", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupMethodsByCardNumber returns a 500 error when fetching monthly topup methods by card number fails.
var ErrApiFailedFindMonthlyTopupMethodsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup methods by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupMethodsByCardNumber returns a 500 error when fetching yearly topup methods by card number fails.
var ErrApiFailedFindYearlyTopupMethodsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup methods by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupAmounts returns a 500 error when fetching monthly topup amounts fails.
var ErrApiFailedFindMonthlyTopupAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupAmounts returns a 500 error when fetching yearly topup amounts fails.
var ErrApiFailedFindYearlyTopupAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupAmountsByCardNumber returns a 500 error when fetching monthly topup amounts by card number fails.
var ErrApiFailedFindMonthlyTopupAmountsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amounts by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupAmountsByCardNumber returns a 500 error when fetching yearly topup amounts by card number fails.
var ErrApiFailedFindYearlyTopupAmountsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amounts by card number", http.StatusInternalServerError)
}

// ErrApiFailedCreateTopup returns a 500 error when creating a new topup fails.
var ErrApiFailedCreateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create topup", http.StatusInternalServerError)
}

// ErrApiFailedUpdateTopup returns a 500 error when updating a topup fails.
var ErrApiFailedUpdateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update topup", http.StatusInternalServerError)
}

// ErrApiValidateCreateTopup returns a 400 error when the create topup request fails validation.
var ErrApiValidateCreateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create topup request", http.StatusBadRequest)
}

// ErrApiValidateUpdateTopup returns a 400 error when the update topup request fails validation.
var ErrApiValidateUpdateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update topup request", http.StatusBadRequest)
}

// ErrApiBindCreateTopup returns a 400 error when binding the create topup request fails.
var ErrApiBindCreateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create topup request", http.StatusBadRequest)
}

// ErrApiBindUpdateTopup returns a 400 error when binding the update topup request fails.
var ErrApiBindUpdateTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update topup request", http.StatusBadRequest)
}

// ErrApiFailedTrashTopup returns a 500 error when moving a topup to trash fails.
var ErrApiFailedTrashTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash topup", http.StatusInternalServerError)
}

// ErrApiFailedRestoreTopup returns a 500 error when restoring a trashed topup fails.
var ErrApiFailedRestoreTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore topup", http.StatusInternalServerError)
}

// ErrApiFailedDeletePermanentTopup returns a 500 error when permanently deleting a topup fails.
var ErrApiFailedDeletePermanentTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete topup", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllTopup returns a 500 error when restoring all trashed topups fails.
var ErrApiFailedRestoreAllTopup = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all topups", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllTopupPermanent returns a 500 error when permanently deleting all topups fails.
var ErrApiFailedDeleteAllTopupPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all topups", http.StatusInternalServerError)
}
