package saldo_errors

import "errors"

// ErrFindAllSaldosFailed is returned when fetching all saldo records fails.
var ErrFindAllSaldosFailed = errors.New("failed to find all saldo records")

// ErrFindActiveSaldosFailed is returned when fetching active saldo records fails.
var ErrFindActiveSaldosFailed = errors.New("failed to find active saldo records")

// ErrFindTrashedSaldosFailed is returned when fetching trashed saldo records fails.
var ErrFindTrashedSaldosFailed = errors.New("failed to find trashed saldo records")

// ErrFindSaldoByIdFailed is returned when fetching a saldo by its ID fails.
var ErrFindSaldoByIdFailed = errors.New("failed to find saldo by ID")

// ErrFindSaldoByCardNumberFailed is returned when fetching saldo by card number fails.
var ErrFindSaldoByCardNumberFailed = errors.New("failed to find saldo by card number")

// ErrGetMonthlyTotalSaldoBalanceFailed is returned when fetching monthly total saldo balance fails.
var ErrGetMonthlyTotalSaldoBalanceFailed = errors.New("failed to get monthly total saldo balance")

// ErrGetYearTotalSaldoBalanceFailed is returned when fetching yearly total saldo balance fails.
var ErrGetYearTotalSaldoBalanceFailed = errors.New("failed to get yearly total saldo balance")

// ErrGetMonthlySaldoBalancesFailed is returned when fetching monthly saldo balances fails.
var ErrGetMonthlySaldoBalancesFailed = errors.New("failed to get monthly saldo balances")

// ErrGetYearlySaldoBalancesFailed is returned when fetching yearly saldo balances fails.
var ErrGetYearlySaldoBalancesFailed = errors.New("failed to get yearly saldo balances")

// ErrCreateSaldoFailed is returned when creating a saldo record fails.
var ErrCreateSaldoFailed = errors.New("failed to create saldo record")

// ErrUpdateSaldoFailed is returned when updating a saldo record fails.
var ErrUpdateSaldoFailed = errors.New("failed to update saldo record")

// ErrUpdateSaldoBalanceFailed is returned when updating saldo balance fails.
var ErrUpdateSaldoBalanceFailed = errors.New("failed to update saldo balance")

// ErrUpdateSaldoWithdrawFailed is returned when updating saldo for a withdrawal fails.
var ErrUpdateSaldoWithdrawFailed = errors.New("failed to update saldo withdrawal")

// ErrTrashSaldoFailed is returned when soft-deleting (trashing) a saldo record fails.
var ErrTrashSaldoFailed = errors.New("failed to trash saldo record")

// ErrRestoreSaldoFailed is returned when restoring a trashed saldo record fails.
var ErrRestoreSaldoFailed = errors.New("failed to restore saldo record")

// ErrDeleteSaldoPermanentFailed is returned when permanently deleting a saldo record fails.
var ErrDeleteSaldoPermanentFailed = errors.New("failed to delete saldo permanently")

// ErrRestoreAllSaldosFailed is returned when restoring all trashed saldo records fails.
var ErrRestoreAllSaldosFailed = errors.New("failed to restore all saldo records")

// ErrDeleteAllSaldosPermanentFailed is returned when permanently deleting all saldo records fails.
var ErrDeleteAllSaldosPermanentFailed = errors.New("failed to delete all saldo records permanently")
