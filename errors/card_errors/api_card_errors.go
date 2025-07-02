package card_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiInvalidCardID returns a 400 Bad Request error when the card ID is invalid.
var ErrApiInvalidCardID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card ID", http.StatusBadRequest)
}

// ErrApiInvalidUserID returns a 400 Bad Request error when the user ID is invalid.
var ErrApiInvalidUserID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid user ID", http.StatusBadRequest)
}

// ErrApiInvalidCardNumber returns a 400 Bad Request error when the card number is invalid.
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card number", http.StatusBadRequest)
}

// ErrApiInvalidMonth returns a 400 Bad Request error when the month value is invalid.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}

// ErrApiInvalidYear returns a 400 Bad Request error when the year value is invalid.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}

// ErrApiFailedFindAllCards returns a 500 Internal Server Error when fetching all cards fails.
var ErrApiFailedFindAllCards = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all cards", http.StatusInternalServerError)
}

// ErrApiFailedFindByIdCard returns a 500 Internal Server Error when fetching a card by ID fails.
var ErrApiFailedFindByIdCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card by ID", http.StatusInternalServerError)
}

// ErrApiFailedDashboardCard returns a 500 Internal Server Error when fetching the card dashboard fails.
var ErrApiFailedDashboardCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card dashboard", http.StatusInternalServerError)
}

// ErrApiFailedDashboardCardByCardNumber returns a 500 Internal Server Error when fetching the dashboard by card number fails.
var ErrApiFailedDashboardCardByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch dashboard by card number", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyBalance returns an API error response when fetching the monthly balance fails.
var ErrApiFailedFindMonthlyBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly balance", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyBalance returns an API error response when fetching the yearly balance fails.
var ErrApiFailedFindYearlyBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly balance", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupAmount returns an API error response when fetching the monthly top-up amount fails.
var ErrApiFailedFindMonthlyTopupAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amount", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupAmount returns an API error response when fetching the yearly top-up amount fails.
var ErrApiFailedFindYearlyTopupAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amount", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawAmount returns an API error response when fetching the monthly withdrawal amount fails.
var ErrApiFailedFindMonthlyWithdrawAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amount", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawAmount returns an API error response when fetching the yearly withdrawal amount fails.
var ErrApiFailedFindYearlyWithdrawAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amount", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransactionAmount returns an API error response when fetching the monthly transaction amount fails.
var ErrApiFailedFindMonthlyTransactionAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transaction amount", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransactionAmount returns an API error response when fetching the yearly transaction amount fails.
var ErrApiFailedFindYearlyTransactionAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transaction amount", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferSenderAmount returns an API error response when fetching the monthly transfer sender amount fails.
var ErrApiFailedFindMonthlyTransferSenderAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer sender amount", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferSenderAmount returns an API error response when fetching the yearly transfer sender amount fails.
var ErrApiFailedFindYearlyTransferSenderAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer sender amount", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferReceiverAmount returns an API error response when fetching the monthly transfer receiver amount fails.
var ErrApiFailedFindMonthlyTransferReceiverAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer receiver amount", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferReceiverAmount returns an API error response when fetching the yearly transfer receiver amount fails.
var ErrApiFailedFindYearlyTransferReceiverAmount = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer receiver amount", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyBalanceByCard returns an API error response when fetching the monthly balance by card fails.
var ErrApiFailedFindMonthlyBalanceByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly balance by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyBalanceByCard returns an API error response when fetching the yearly balance by card fails.
var ErrApiFailedFindYearlyBalanceByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly balance by card", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTopupAmountByCard returns an API error response when fetching the monthly top-up amount by card fails.
var ErrApiFailedFindMonthlyTopupAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly topup amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTopupAmountByCard returns an API error response when fetching the yearly top-up amount by card fails.
var ErrApiFailedFindYearlyTopupAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly topup amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyWithdrawAmountByCard returns an API error response when fetching the monthly withdrawal amount by card fails.
var ErrApiFailedFindMonthlyWithdrawAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly withdraw amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyWithdrawAmountByCard returns an API error response when fetching the yearly withdrawal amount by card fails.
var ErrApiFailedFindYearlyWithdrawAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly withdraw amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransactionAmountByCard returns an API error response when fetching the monthly transaction amount by card fails.
var ErrApiFailedFindMonthlyTransactionAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transaction amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransactionAmountByCard returns an API error response when fetching the yearly transaction amount by card fails.
var ErrApiFailedFindYearlyTransactionAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transaction amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferSenderAmountByCard represents error when failing to fetch monthly transfer amount by sender card.
var ErrApiFailedFindMonthlyTransferSenderAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer sender amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferSenderAmountByCard represents error when failing to fetch yearly transfer amount by sender card.
var ErrApiFailedFindYearlyTransferSenderAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer sender amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindMonthlyTransferReceiverAmountByCard represents error when failing to fetch monthly transfer amount by receiver card.
var ErrApiFailedFindMonthlyTransferReceiverAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly transfer receiver amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindYearlyTransferReceiverAmountByCard represents error when failing to fetch yearly transfer amount by receiver card.
var ErrApiFailedFindYearlyTransferReceiverAmountByCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly transfer receiver amount by card", http.StatusInternalServerError)
}

// ErrApiFailedFindByUserIdCard represents error when failing to fetch cards by user ID.
var ErrApiFailedFindByUserIdCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card by user ID", http.StatusInternalServerError)
}

// ErrApiFailedFindByActiveCard represents error when failing to fetch active cards.
var ErrApiFailedFindByActiveCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch active cards", http.StatusInternalServerError)
}

// ErrApiFailedFindByTrashedCard represents error when failing to fetch trashed cards.
var ErrApiFailedFindByTrashedCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch trashed cards", http.StatusInternalServerError)
}

// ErrApiFailedFindByCardNumber represents error when failing to fetch card by card number.
var ErrApiFailedFindByCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch card by card number", http.StatusInternalServerError)
}

// ErrApiFailedCreateCard represents error when failing to create a new card.
var ErrApiFailedCreateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create card", http.StatusInternalServerError)
}

// ErrApiFailedUpdateCard represents error when failing to update a card.
var ErrApiFailedUpdateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update card", http.StatusInternalServerError)
}

// ErrApiValidateCreateCard represents validation error for create card request.
var ErrApiValidateCreateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Card request", http.StatusBadRequest)
}

// ErrApiValidateUpdateCard represents validation error for update card request.
var ErrApiValidateUpdateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Card request", http.StatusBadRequest)
}

// ErrApiBindCreateCard represents binding error for create card request.
var ErrApiBindCreateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Card request", http.StatusBadRequest)
}

// ErrApiBindUpdateCard represents binding error for update card request.
var ErrApiBindUpdateCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Card request", http.StatusBadRequest)
}

// ErrApiFailedTrashCard represents error when failing to move a card to trash.
var ErrApiFailedTrashCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash card", http.StatusInternalServerError)
}

// ErrApiFailedRestoreCard represents error when failing to restore a trashed card.
var ErrApiFailedRestoreCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore card", http.StatusInternalServerError)
}

// ErrApiFailedDeleteCardPermanent represents error when failing to permanently delete a card.
var ErrApiFailedDeleteCardPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete card", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllCard represents error when failing to restore all trashed cards.
var ErrApiFailedRestoreAllCard = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all cards", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllCardPermanent represents error when failing to permanently delete all trashed cards.
var ErrApiFailedDeleteAllCardPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all cards", http.StatusInternalServerError)
}
