package withdraw_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiWithdrawInvalidID is an error response for invalid withdraw ID.
var ErrApiWithdrawInvalidID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Withdraw ID", http.StatusBadRequest)
}

// ErrApiWithdrawInvalidUserID is an error response for invalid withdraw merchant ID.
var ErrApiWithdrawInvalidUserID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Withdraw Merchant ID", http.StatusBadRequest)
}

// ErrApiInvalidCardNumber is an error response for invalid card number.
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card number", http.StatusBadRequest)
}

// ErrApiInvalidMonth is an error response for invalid month.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month", http.StatusBadRequest)
}

// ErrApiInvalidYear is an error response for invalid year.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "year", "Invalid year", http.StatusBadRequest)
}

// ErrApiWithdrawNotFound is an error response for withdraw not found.
var ErrApiWithdrawNotFound = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Withdraw not found", http.StatusNotFound)
}

// ErrApiFailedFindAllWithdraw is an error response for failed to fetch all withdraws.
var ErrApiFailedFindAllWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindAllWithdrawByCardNumber is an error response for failed to fetch withdraws by card number.
var ErrApiFailedFindAllWithdrawByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch withdraws by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdWithdraw is an error response for failed to fetch withdraw by ID.
var ErrApiFailedFindByIdWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch withdraw by ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByCardNumber is an error response for failed to fetch withdraws using card number.
var ErrApiFailedFindByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch withdraws using card number", http.StatusInternalServerError)
}

// ErrApiFailedFindByActiveWithdraw is an error response for failed to fetch active withdraws.
var ErrApiFailedFindByActiveWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindByTrashedWithdraw is an error response for failed to fetch trashed withdraws.
var ErrApiFailedFindByTrashedWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawStatusSuccess is an error response for failed to fetch monthly successful withdraws.
var ErrApiFailedFindMonthlyWithdrawStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawStatusSuccess is an error response for failed to fetch yearly successful withdraws.
var ErrApiFailedFindYearlyWithdrawStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawStatusFailed is an error response for failed to fetch monthly failed withdraws.
var ErrApiFailedFindMonthlyWithdrawStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawStatusFailed is an error response for failed to fetch yearly failed withdraws.
var ErrApiFailedFindYearlyWithdrawStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed withdraws", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawStatusSuccessCardNumber is an error response for failed to fetch monthly successful withdraws by card number.
var ErrApiFailedFindMonthlyWithdrawStatusSuccessCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful withdraws by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawStatusSuccessCardNumber is an error response for failed to fetch yearly successful withdraws by card number.
var ErrApiFailedFindYearlyWithdrawStatusSuccessCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful withdraws by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawStatusFailedCardNumber is an error response for failed to fetch monthly failed withdraws by card number.
var ErrApiFailedFindMonthlyWithdrawStatusFailedCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed withdraws by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawStatusFailedCardNumber is an error response for failed to fetch yearly failed withdraws by card number.
var ErrApiFailedFindYearlyWithdrawStatusFailedCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed withdraws by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdraws is an error response for failed to fetch monthly withdraw amounts.
var ErrApiFailedFindMonthlyWithdraws = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdraws is an error response for failed to fetch yearly withdraw amounts.
var ErrApiFailedFindYearlyWithdraws = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawsByCardNumber is an error response for failed to fetch monthly withdraw amounts by card number.
var ErrApiFailedFindMonthlyWithdrawsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amounts by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawsByCardNumber is an error response for failed to fetch yearly withdraw amounts by card number.
var ErrApiFailedFindYearlyWithdrawsByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amounts by card number", http.StatusInternalServerError)
}

// ErrApiFailedCreateWithdraw is an error response for failed to create withdraw.
var ErrApiFailedCreateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create withdraw", http.StatusInternalServerError)
}

// ErrApiFailedUpdateWithdraw is an error response for failed to update withdraw.
var ErrApiFailedUpdateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update withdraw", http.StatusInternalServerError)
}

// ErrApiBindCreateWithdraw is an error response for bind failed: invalid create withdraw request.
var ErrApiBindCreateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create withdraw request", http.StatusBadRequest)
}

// ErrApiBindUpdateWithdraw is an error response for bind failed: invalid update withdraw request.
var ErrApiBindUpdateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update withdraw request", http.StatusBadRequest)
}

// ErrApiValidateCreateWithdraw is an error response for validation failed: invalid create withdraw request.
var ErrApiValidateCreateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create withdraw request", http.StatusBadRequest)
}

// ErrApiValidateUpdateWithdraw is an error response for validation failed: invalid update withdraw request.
var ErrApiValidateUpdateWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update withdraw request", http.StatusBadRequest)
}

// ErrApiFailedTrashedWithdraw is an error response for failed to move withdraw to trash.
var ErrApiFailedTrashedWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to move withdraw to trash", http.StatusInternalServerError)
}

// ErrApiFailedRestoreWithdraw is an error response for failed to restore withdraw.
var ErrApiFailedRestoreWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore withdraw", http.StatusInternalServerError)
}

// ErrApiFailedDeleteWithdrawPermanent is an error response for failed to permanently delete withdraw.
var ErrApiFailedDeleteWithdrawPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete withdraw", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllWithdraw is an error response for failed to restore all withdraws.
var ErrApiFailedRestoreAllWithdraw = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all withdraws", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllWithdrawPermanent is an error response for failed to permanently delete all withdraws.
var ErrApiFailedDeleteAllWithdrawPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all withdraws", http.StatusInternalServerError)
}
