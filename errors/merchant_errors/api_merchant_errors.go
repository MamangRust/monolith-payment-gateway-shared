package merchant_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiInvalidCardID is returned when the card ID provided is invalid.
var ErrApiInvalidCardID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card ID", http.StatusBadRequest)
}

// ErrApiInvalidMerchantID is returned when the merchant ID provided is invalid.
var ErrApiInvalidMerchantID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid merchant ID", http.StatusBadRequest)
}

// ErrApiInvalidMonth is returned when the month parameter is invalid or malformed.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}

// ErrApiInvalidYear is returned when the year parameter is invalid or malformed.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}

// ErrApiInvalidUserID is returned when the user ID provided is invalid.
var ErrApiInvalidUserID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid user ID", http.StatusBadRequest)
}

// ErrApiInvalidApiKey is returned when the API key provided is invalid or unauthorized.
var ErrApiInvalidApiKey = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid API key", http.StatusUnauthorized)
}

// ErrApiFailedFindAllMerchants is returned when failing to retrieve all merchants.
var ErrApiFailedFindAllMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchants", http.StatusInternalServerError)
}

// ErrApiFailedFindAllMerchantsActive is returned when failing to retrieve all active merchants.
var ErrApiFailedFindAllMerchantsActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchants active", http.StatusInternalServerError)
}

// ErrApiFailedFindAllMerchantsTrashed is returned when failing to retrieve all trashed merchants.
var ErrApiFailedFindAllMerchantsTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchants trashed", http.StatusInternalServerError)
}

// ErrApiFailedFindByUserId is returned when failing to retrieve a merchant by user ID.
var ErrApiFailedFindByUserId = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant by user ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdMerchant is returned when failing to retrieve a merchant by ID.
var ErrApiFailedFindByIdMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant by ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByApiKeyMerchant is returned when failing to retrieve a merchant by API key.
var ErrApiFailedFindByApiKeyMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant by api key", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTransactions is returned when failing to retrieve all merchant transactions.
var ErrApiFailedFindAllTransactions = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all transactions", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTransactionByMerchant is returned when failing to retrieve transactions by merchant.
var ErrApiFailedFindAllTransactionByMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindAllTransactionByApikey is returned when failing to retrieve transactions by API key.
var ErrApiFailedFindAllTransactionByApikey = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch transactions by API key", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyPaymentMethodsMerchant is returned when failing to retrieve monthly payment methods for a merchant.
var ErrApiFailedFindMonthlyPaymentMethodsMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyPaymentMethodMerchant is returned when failing to retrieve yearly payment methods for a merchant.
var ErrApiFailedFindYearlyPaymentMethodMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyAmountMerchant is returned when failing to retrieve monthly transaction amount for a merchant.
var ErrApiFailedFindMonthlyAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amount by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyAmountMerchant is returned when failing to retrieve yearly transaction amount for a merchant.
var ErrApiFailedFindYearlyAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amount by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTotalAmountMerchant is returned when failing to retrieve total monthly transaction amount for a merchant.
var ErrApiFailedFindMonthlyTotalAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly total amount by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTotalAmountMerchant is returned when failing to retrieve total yearly transaction amount for a merchant.
var ErrApiFailedFindYearlyTotalAmountMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly total amount by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyPaymentMethodByMerchants is returned when failing to fetch monthly payment methods for all merchants.
var ErrApiFailedFindMonthlyPaymentMethodByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyPaymentMethodByMerchants is returned when failing to fetch yearly payment methods for all merchants.
var ErrApiFailedFindYearlyPaymentMethodByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyAmountByMerchants is returned when failing to fetch monthly amounts for all merchants.
var ErrApiFailedFindMonthlyAmountByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyAmountByMerchants is returned when failing to fetch yearly amounts for all merchants.
var ErrApiFailedFindYearlyAmountByMerchants = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts by merchant", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyPaymentMethodByApikeys is returned when failing to fetch monthly payment methods by API key.
var ErrApiFailedFindMonthlyPaymentMethodByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly payment methods by API key", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyPaymentMethodByApikeys is returned when failing to fetch yearly payment methods by API key.
var ErrApiFailedFindYearlyPaymentMethodByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly payment methods by API key", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyAmountByApikeys is returned when failing to fetch monthly amounts by API key.
var ErrApiFailedFindMonthlyAmountByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly amounts by API key", http.StatusInternalServerError)
}

var ErrApiFailedFindYearlyAmountByApikeys = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly amounts by API key", http.StatusInternalServerError)
}

var ErrApiFailedCreateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create merchant", http.StatusInternalServerError)
}

var ErrApiFailedUpdateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant", http.StatusInternalServerError)
}

var ErrApiFailedUpdateMerchantStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant status", http.StatusInternalServerError)
}

var ErrApiValidateCreateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create merchant request", http.StatusBadRequest)
}

var ErrApiValidateUpdateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant request", http.StatusBadRequest)
}

var ErrApiValidateUpdateMerchantStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant status request", http.StatusBadRequest)
}

var ErrApiBindCreateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create merchant request", http.StatusBadRequest)
}

var ErrApiBindUpdateMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant request", http.StatusBadRequest)
}

var ErrApiBindUpdateMerchantStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant status request", http.StatusBadRequest)
}

var ErrApiFailedTrashMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash merchant", http.StatusInternalServerError)
}

var ErrApiFailedRestoreMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore merchant", http.StatusInternalServerError)
}

var ErrApiFailedDeleteMerchantPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete merchant", http.StatusInternalServerError)
}

var ErrApiFailedRestoreAllMerchant = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all merchants", http.StatusInternalServerError)
}

var ErrApiFailedDeleteAllMerchantPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all merchants", http.StatusInternalServerError)
}
