package transaction_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiInvalidTransactionApiKey returns an API error response
// when the transaction API key is missing or invalid.
var ErrApiInvalidTransactionApiKey = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction api-key", http.StatusBadRequest)
}

// ErrApiInvalidTransactionID returns an API error response
// when the transaction ID is invalid or improperly formatted.
var ErrApiInvalidTransactionID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction ID", http.StatusBadRequest)
}

// ErrApiInvalidTransactionCardNumber returns an API error response
// when the provided transaction card number is invalid.
var ErrApiInvalidTransactionCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction Card Number", http.StatusBadRequest)
}

// ErrApiInvalidTransactionMerchantID returns an API error response
// when the provided merchant ID for the transaction is invalid.
var ErrApiInvalidTransactionMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid transaction merchant ID", http.StatusBadRequest)
}

// ErrApiInvalidMonth returns an API error response
// when the month value is missing, invalid, or out of range.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}

// ErrApiInvalidYear returns an API error response
// when the year value is missing, invalid, or out of range.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}

// ErrApiFailedFindAllTransactions returns an API error response
// when fetching all transactions from the database fails.
var ErrApiFailedFindAllTransactions = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTransactionsActive returns an API error response
// when fetching active (non-deleted) transactions fails.
var ErrApiFailedFindAllTransactionsActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions active", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTransactionsTrashed returns an API error response
// when fetching trashed (soft-deleted) transactions fails.
var ErrApiFailedFindAllTransactionsTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions active", http.StatusInternalServerError)
}

// ErrApiFailedFindByCardNumber returns an API error response
// when fetching transactions filtered by card number fails.
var ErrApiFailedFindByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindById returns an API error response
// when fetching a transaction by its ID fails.
var ErrApiFailedFindById = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transaction by ID", http.StatusInternalServerError)
}

// ErrApiFailedMonthlySuccess returns an API error response
// when retrieving monthly successful transactions fails.
var ErrApiFailedMonthlySuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transactions", http.StatusInternalServerError)
}

// ErrApiFailedYearlySuccess returns an API error response
// when retrieving yearly successful transactions fails.
var ErrApiFailedYearlySuccess = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transactions", http.StatusInternalServerError)
}

// ErrApiFailedMonthlyFailed returns an API error response
// when retrieving monthly failed transactions fails.
var ErrApiFailedMonthlyFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transactions", http.StatusInternalServerError)
}

// ErrApiFailedYearlyFailed returns an API error response
// when retrieving yearly failed transactions fails.
var ErrApiFailedYearlyFailed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transactions", http.StatusInternalServerError)
}

// ErrApiFailedMonthlySuccessByCard returns an API error response
// when retrieving monthly successful transactions by card number fails.
var ErrApiFailedMonthlySuccessByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly successful transactions by card number", http.StatusInternalServerError)
}

// ErrApiFailedYearlySuccessByCard returns an API error response
// when retrieving yearly successful transactions by card number fails.
var ErrApiFailedYearlySuccessByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly successful transactions by card number", http.StatusInternalServerError)
}

// ErrApiFailedMonthlyFailedByCard returns an API error response
// when retrieving monthly failed transactions by card number fails.
var ErrApiFailedMonthlyFailedByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly failed transactions by card number", http.StatusInternalServerError)
}

// ErrApiFailedYearlyFailedByCard returns an API error response
// when retrieving yearly failed transactions by card number fails.
var ErrApiFailedYearlyFailedByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly failed transactions by card number", http.StatusInternalServerError)
}

// ErrApiFailedMonthlyMethods returns an API error response
// when retrieving monthly payment methods fails.
var ErrApiFailedMonthlyMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods", http.StatusInternalServerError)
}

// ErrApiFailedYearlyMethods returns an API error response
// when retrieving yearly payment methods fails.
var ErrApiFailedYearlyMethods = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods", http.StatusInternalServerError)
}

// ErrApiFailedMonthlyAmounts returns an API error response
// when retrieving total monthly transaction amounts fails.
var ErrApiFailedMonthlyAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts", http.StatusInternalServerError)
}

// ErrApiFailedYearlyAmounts returns an API error response
// when retrieving total yearly transaction amounts fails.
var ErrApiFailedYearlyAmounts = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts", http.StatusInternalServerError)
}

// ErrApiFailedMonthlyMethodsByCard returns an API error response
// when retrieving monthly payment methods by card number fails.
var ErrApiFailedMonthlyMethodsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by card number", http.StatusInternalServerError)
}

// ErrApiFailedYearlyMethodsByCard returns an API error response
// when retrieving yearly payment methods by card number fails.
var ErrApiFailedYearlyMethodsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by card number", http.StatusInternalServerError)
}

// ErrApiFailedMonthlyAmountsByCard returns an API error response
// when retrieving total monthly transaction amounts by card number fails.
var ErrApiFailedMonthlyAmountsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts by card number", http.StatusInternalServerError)
}

// ErrApiFailedYearlyAmountsByCard returns an API error response
// when retrieving total yearly transaction amounts by card number fails.
var ErrApiFailedYearlyAmountsByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindByMerchantID returns an API error response
// when retrieving transactions by merchant ID fails.
var ErrApiFailedFindByMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by merchant ID", http.StatusInternalServerError)
}

// ErrApiFailedFindActive returns an API error response
// when retrieving active (non-deleted) transactions fails.
var ErrApiFailedFindActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active transactions", http.StatusInternalServerError)
}

// ErrApiFailedFindTrashed returns an API error response
// when retrieving trashed (soft-deleted) transactions fails.
var ErrApiFailedFindTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed transactions", http.StatusInternalServerError)
}

// ErrApiFailedCreateTransaction returns an API error response
// when creating a new transaction fails.
var ErrApiFailedCreateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create transaction", http.StatusInternalServerError)
}

// ErrApiFailedUpdateTransaction returns an API error response
// when updating an existing transaction fails.
var ErrApiFailedUpdateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update transaction", http.StatusInternalServerError)
}

// ErrApiValidateCreateTransaction returns an API error response
// when validation of the create transaction request fails.
var ErrApiValidateCreateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create transaction request", http.StatusBadRequest)
}

// ErrApiValidateUpdateTransaction returns an API error response
// when validation of the update transaction request fails.
var ErrApiValidateUpdateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update transaction request", http.StatusBadRequest)
}

// ErrApiBindCreateTransaction returns an API error response
// when binding the create transaction request body to struct fails.
var ErrApiBindCreateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create transaction request", http.StatusBadRequest)
}

// ErrApiBindUpdateTransaction returns an API error response
// when binding the update transaction request body to struct fails.
var ErrApiBindUpdateTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update transaction request", http.StatusBadRequest)
}

// ErrApiFailedRestoreTransaction returns an API error response
// when restoring a soft-deleted transaction fails.
var ErrApiFailedRestoreTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore transaction", http.StatusInternalServerError)
}

// ErrApiFailedTrashTransaction returns an API error response
// when moving a transaction to trash (soft-delete) fails.
var ErrApiFailedTrashTransaction = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to move transaction to trash", http.StatusInternalServerError)
}

// ErrApiFailedDeletePermanent returns an API error response
// when permanently deleting a transaction fails.
var ErrApiFailedDeletePermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete transaction", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllTransactions returns an API error response
// when restoring all trashed transactions fails.
var ErrApiFailedRestoreAllTransactions = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all transactions", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllPermanent returns an API error response
// when permanently deleting all trashed transactions fails.
var ErrApiFailedDeleteAllPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all trashed transactions", http.StatusInternalServerError)
}
