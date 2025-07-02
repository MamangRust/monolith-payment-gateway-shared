package transfer_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiTransferInvalidID is an error response for invalid transfer ID.
var ErrApiTransferInvalidID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Transfer ID", http.StatusBadRequest)
}

// ErrApiTransferInvalidMerchantID is an error response for invalid transfer merchant ID.
var ErrApiTransferInvalidMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid Transfer Merchant ID", http.StatusBadRequest)
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
	return response.NewApiErrorResponse(c, "error", "Invalid year", http.StatusBadRequest)
}

// ErrApiFailedFindAllTransfers is an error response for failed to fetch all transfers.
var ErrApiFailedFindAllTransfers = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdTransfer is an error response for failed to fetch transfer by ID.
var ErrApiFailedFindByIdTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transfer by ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByTransferFrom is an error response for failed to fetch transfers by transfer_from.
var ErrApiFailedFindByTransferFrom = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transfers by transfer_from", http.StatusInternalServerError)
}

// ErrApiFailedFindByTransferTo is an error response for failed to fetch transfers by transfer_to.
var ErrApiFailedFindByTransferTo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transfers by transfer_to", http.StatusInternalServerError)
}

// ErrApiFailedFindByActiveTransfer is an error response for failed to fetch active transfers.
var ErrApiFailedFindByActiveTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindByTrashedTransfer returns an API error response when fetching trashed (soft-deleted) transfers fails.
var ErrApiFailedFindByTrashedTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferStatusSuccess returns an API error response when fetching monthly successful transfer data fails.
var ErrApiFailedFindMonthlyTransferStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferStatusSuccess returns an API error response when fetching yearly successful transfer data fails.
var ErrApiFailedFindYearlyTransferStatusSuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferStatusFailed returns an API error response when fetching monthly failed transfer data fails.
var ErrApiFailedFindMonthlyTransferStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferStatusFailed returns an API error response when fetching yearly failed transfer data fails.
var ErrApiFailedFindYearlyTransferStatusFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transfers", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferStatusSuccessByCardNumber returns an API error response when fetching monthly successful transfers by card number fails.
var ErrApiFailedFindMonthlyTransferStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transfers by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferStatusSuccessByCardNumber returns an API error response when fetching yearly successful transfers by card number fails.
var ErrApiFailedFindYearlyTransferStatusSuccessByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transfers by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferStatusFailedByCardNumber returns an API error response when fetching monthly failed transfers by card number fails.
var ErrApiFailedFindMonthlyTransferStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transfers by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferStatusFailedByCardNumber returns an API error response
// when fetching yearly failed transfers by card number fails.
var ErrApiFailedFindYearlyTransferStatusFailedByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transfers by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferAmounts returns an API error response
// when fetching monthly total transfer amounts fails.
var ErrApiFailedFindMonthlyTransferAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferAmounts returns an API error response
// when fetching yearly total transfer amounts fails.
var ErrApiFailedFindYearlyTransferAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer amounts", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferAmountsBySenderCardNumber returns an API error response
// when fetching monthly transfer amounts by sender card number fails.
var ErrApiFailedFindMonthlyTransferAmountsBySenderCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer amounts by sender card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferAmountsByReceiverCardNumber returns an API error response
// when fetching monthly transfer amounts by receiver card number fails.
var ErrApiFailedFindMonthlyTransferAmountsByReceiverCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer amounts by receiver card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferAmountsBySenderCardNumber returns an API error response
// when fetching yearly transfer amounts by sender card number fails.
var ErrApiFailedFindYearlyTransferAmountsBySenderCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer amounts by sender card number", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferAmountsByReceiverCardNumber returns an API error response
// when fetching yearly transfer amounts by receiver card number fails.
var ErrApiFailedFindYearlyTransferAmountsByReceiverCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer amounts by receiver card number", http.StatusInternalServerError)
}

// ErrApiFailedCreateTransfer returns an API error response
// when creating a new transfer fails.
var ErrApiFailedCreateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create transfer", http.StatusInternalServerError)
}

// ErrApiFailedUpdateTransfer returns an API error response
// when updating a transfer fails.
var ErrApiFailedUpdateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update transfer", http.StatusInternalServerError)
}

// ErrApiBindCreateTransfer returns an API error response
// when binding the request for creating a transfer fails due to invalid input.
var ErrApiBindCreateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create transfer request", http.StatusBadRequest)
}

// ErrApiBindUpdateTransfer returns an API error response
// when binding the request for updating a transfer fails due to invalid input.
var ErrApiBindUpdateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update transfer request", http.StatusBadRequest)
}

// ErrApiValidateCreateTransfer returns an API error response
// when validation fails on the create transfer request payload.
var ErrApiValidateCreateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create transfer request", http.StatusBadRequest)
}

// ErrApiValidateUpdateTransfer returns an API error response
// when validation fails on the update transfer request payload.
var ErrApiValidateUpdateTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update transfer request", http.StatusBadRequest)
}

// ErrApiFailedTrashedTransfer returns an API error response
// when attempting to trash (soft-delete) a transfer fails.
var ErrApiFailedTrashedTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash transfer", http.StatusInternalServerError)
}

// ErrApiFailedRestoreTransfer returns an API error response
// when attempting to restore a trashed transfer fails.
var ErrApiFailedRestoreTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore transfer", http.StatusInternalServerError)
}

// ErrApiFailedDeleteTransferPermanent returns an API error response
// when attempting to permanently delete a transfer fails.
var ErrApiFailedDeleteTransferPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete transfer", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllTransfer returns an API error response
// when attempting to restore all trashed transfers fails.
var ErrApiFailedRestoreAllTransfer = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all transfers", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllTransferPermanent returns an API error response
// when attempting to permanently delete all transfers fails.
var ErrApiFailedDeleteAllTransferPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all transfers", http.StatusInternalServerError)
}
