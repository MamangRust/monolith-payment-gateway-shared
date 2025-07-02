package saldo_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

// ErrApiInvalidSaldoID indicates an invalid saldo ID in the request.
var ErrApiInvalidSaldoID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid saldo ID", http.StatusBadRequest)
}

// ErrApiInvalidMonth indicates an invalid month value in the request.
var ErrApiInvalidMonth = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid month value", http.StatusBadRequest)
}

// ErrApiInvalidYear indicates an invalid year value in the request.
var ErrApiInvalidYear = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid year value", http.StatusBadRequest)
}

// ErrApiInvalidCardNumber indicates an invalid card number value in the request.
var ErrApiInvalidCardNumber = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid card-number value", http.StatusBadRequest)
}

// ErrApiFailedFindAllSaldo is returned when fetching all saldo records fails.
var ErrApiFailedFindAllSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all saldos", http.StatusInternalServerError)
}

// ErrApiFailedFindAllSaldoTrashed is returned when fetching trashed saldo records fails.
var ErrApiFailedFindAllSaldoTrashed = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all saldos trashed", http.StatusInternalServerError)
}

// ErrApiFailedFindAllSaldoActive is returned when fetching active saldo records fails.
var ErrApiFailedFindAllSaldoActive = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all saldos active", http.StatusInternalServerError)
}

var ErrApiFailedFindByIdSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch saldo by ID", http.StatusInternalServerError)
}

var ErrApiFailedFindMonthlyTotalSaldoBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly total saldo balance", http.StatusInternalServerError)
}

var ErrApiFailedFindYearTotalSaldoBalance = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly total saldo balance", http.StatusInternalServerError)
}

var ErrApiFailedFindMonthlySaldoBalances = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch monthly saldo balances", http.StatusInternalServerError)
}

var ErrApiFailedFindYearlySaldoBalances = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch yearly saldo balances", http.StatusInternalServerError)
}

var ErrApiFailedFindByCardNumberSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch saldo by card number", http.StatusInternalServerError)
}

var ErrApiFailedCreateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create saldo", http.StatusInternalServerError)
}

var ErrApiFailedUpdateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update saldo", http.StatusInternalServerError)
}

// ErrApiValidateCreateSaldo is returned when validation fails for saldo creation.
var ErrApiValidateCreateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create Saldo request", http.StatusBadRequest)
}

// ErrApiValidateUpdateSaldo is returned when validation fails for saldo update.
var ErrApiValidateUpdateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update Saldo request", http.StatusBadRequest)
}

// ErrApiBindCreateSaldo is returned when binding fails for saldo creation.
var ErrApiBindCreateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create Saldo request", http.StatusBadRequest)
}

// ErrApiBindUpdateSaldo is returned when binding fails for saldo update.
var ErrApiBindUpdateSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update Saldo request", http.StatusBadRequest)
}

// ErrApiFailedTrashSaldo is returned when soft-deleting saldo fails.
var ErrApiFailedTrashSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash saldo", http.StatusInternalServerError)
}

// ErrApiFailedRestoreSaldo is returned when restoring saldo fails.
var ErrApiFailedRestoreSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore saldo", http.StatusInternalServerError)
}

// ErrApiFailedDeleteSaldoPermanent is returned when permanently deleting saldo fails.
var ErrApiFailedDeleteSaldoPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete saldo", http.StatusInternalServerError)
}

// ErrApiFailedRestoreAllSaldo is returned when restoring all saldos fails.
var ErrApiFailedRestoreAllSaldo = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all saldos", http.StatusInternalServerError)
}

// ErrApiFailedDeleteAllSaldoPermanent is returned when permanently deleting all saldos fails.
var ErrApiFailedDeleteAllSaldoPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all saldos", http.StatusInternalServerError)
}
