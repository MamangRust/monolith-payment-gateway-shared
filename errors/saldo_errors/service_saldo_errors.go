package saldo_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrFailedFindAllSaldos returns a 500 error when fetching all saldo records fails.
var ErrFailedFindAllSaldos = response.NewErrorResponse("Failed to fetch saldos", http.StatusInternalServerError)

// ErrFailedSaldoNotFound returns a 404 error when a requested saldo record is not found.
var ErrFailedSaldoNotFound = response.NewErrorResponse("Saldo not found", http.StatusNotFound)

// ErrFailedInsuffientBalance returns a 400 error when a transaction is attempted with insufficient balance.
var ErrFailedInsuffientBalance = response.NewErrorResponse("Insufficient balance", http.StatusBadRequest)

// ErrFailedFindMonthlyTotalSaldoBalance returns a 500 error when fetching monthly total saldo balance fails.
var ErrFailedFindMonthlyTotalSaldoBalance = response.NewErrorResponse("Failed to fetch monthly total saldo balance", http.StatusInternalServerError)

// ErrFailedFindYearTotalSaldoBalance returns a 500 error when fetching yearly total saldo balance fails.
var ErrFailedFindYearTotalSaldoBalance = response.NewErrorResponse("Failed to fetch yearly total saldo balance", http.StatusInternalServerError)

// ErrFailedFindMonthlySaldoBalances returns a 500 error when fetching monthly saldo balances fails.
var ErrFailedFindMonthlySaldoBalances = response.NewErrorResponse("Failed to fetch monthly saldo balances", http.StatusInternalServerError)

// ErrFailedFindYearlySaldoBalances returns a 500 error when fetching yearly saldo balances fails.
var ErrFailedFindYearlySaldoBalances = response.NewErrorResponse("Failed to fetch yearly saldo balances", http.StatusInternalServerError)

// ErrFailedFindSaldoByCardNumber returns a 500 error when fetching saldo by card number fails.
var ErrFailedFindSaldoByCardNumber = response.NewErrorResponse("Failed to find saldo by card number", http.StatusInternalServerError)

// ErrFailedFindActiveSaldos returns a 500 error when fetching active saldo records fails.
var ErrFailedFindActiveSaldos = response.NewErrorResponse("Failed to fetch active saldos", http.StatusInternalServerError)

// ErrFailedFindTrashedSaldos returns a 500 error when fetching trashed saldo records fails.
var ErrFailedFindTrashedSaldos = response.NewErrorResponse("Failed to fetch trashed saldos", http.StatusInternalServerError)

// ErrFailedCreateSaldo returns a 500 error when creating a saldo record fails.
var ErrFailedCreateSaldo = response.NewErrorResponse("Failed to create saldo", http.StatusInternalServerError)

// ErrFailedUpdateSaldo returns a 500 error when updating a saldo record fails.
var ErrFailedUpdateSaldo = response.NewErrorResponse("Failed to update saldo", http.StatusInternalServerError)

// ErrFailedTrashSaldo returns a 500 error when moving a saldo record to trash fails.
var ErrFailedTrashSaldo = response.NewErrorResponse("Failed to trash saldo", http.StatusInternalServerError)

// ErrFailedRestoreSaldo returns a 500 error when restoring a trashed saldo record fails.
var ErrFailedRestoreSaldo = response.NewErrorResponse("Failed to restore saldo", http.StatusInternalServerError)

// ErrFailedDeleteSaldoPermanent returns a 500 error when permanently deleting a saldo record fails.
var ErrFailedDeleteSaldoPermanent = response.NewErrorResponse("Failed to permanently delete saldo", http.StatusInternalServerError)

// ErrFailedRestoreAllSaldo returns a 500 error when restoring all trashed saldo records fails.
var ErrFailedRestoreAllSaldo = response.NewErrorResponse("Failed to restore all saldos", http.StatusInternalServerError)

// ErrFailedDeleteAllSaldoPermanent returns a 500 error when permanently deleting all trashed saldo records fails.
var ErrFailedDeleteAllSaldoPermanent = response.NewErrorResponse("Failed to permanently delete all saldos", http.StatusInternalServerError)

