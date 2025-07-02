package transaction_errors

import "errors"

// ErrFindAllTransactionsFailed indicates a failure when attempting to retrieve all transactions.
var ErrFindAllTransactionsFailed = errors.New("failed to find all transactions")

// ErrFindActiveTransactionsFailed indicates a failure when retrieving active (non-deleted) transactions.
var ErrFindActiveTransactionsFailed = errors.New("failed to find active transactions")

// ErrFindTrashedTransactionsFailed indicates a failure when retrieving soft-deleted (trashed) transactions.
var ErrFindTrashedTransactionsFailed = errors.New("failed to find trashed transactions")

// ErrFindTransactionsByCardNumberFailed indicates a failure when retrieving transactions by card number.
var ErrFindTransactionsByCardNumberFailed = errors.New("failed to find transactions by card number")

// ErrFindTransactionByIdFailed indicates a failure when retrieving a transaction by its ID.
var ErrFindTransactionByIdFailed = errors.New("failed to find transaction by ID")

// ErrGetMonthTransactionStatusSuccessFailed indicates a failure when retrieving monthly successful transaction status.
var ErrGetMonthTransactionStatusSuccessFailed = errors.New("failed to get monthly transaction status success")

// ErrGetYearlyTransactionStatusSuccessFailed indicates a failure when retrieving yearly successful transaction status.
var ErrGetYearlyTransactionStatusSuccessFailed = errors.New("failed to get yearly transaction status success")

// ErrGetMonthTransactionStatusSuccessByCardFailed indicates a failure when retrieving monthly successful transactions by card number.
var ErrGetMonthTransactionStatusSuccessByCardFailed = errors.New("failed to get monthly transaction status success by card number")

// ErrGetYearlyTransactionStatusSuccessByCardFailed indicates a failure when retrieving yearly successful transactions by card number.
var ErrGetYearlyTransactionStatusSuccessByCardFailed = errors.New("failed to get yearly transaction status success by card number")

// ErrGetMonthTransactionStatusFailedFailed indicates a failure when retrieving monthly failed transaction status.
var ErrGetMonthTransactionStatusFailedFailed = errors.New("failed to get monthly transaction status failed")

// ErrGetYearlyTransactionStatusFailedFailed indicates a failure when retrieving yearly failed transaction status.
var ErrGetYearlyTransactionStatusFailedFailed = errors.New("failed to get yearly transaction status failed")

// ErrGetMonthTransactionStatusFailedByCardFailed indicates a failure when retrieving monthly failed transactions by card number.
var ErrGetMonthTransactionStatusFailedByCardFailed = errors.New("failed to get monthly transaction status failed by card number")

// ErrGetYearlyTransactionStatusFailedByCardFailed indicates a failure when retrieving yearly failed transactions by card number.
var ErrGetYearlyTransactionStatusFailedByCardFailed = errors.New("failed to get yearly transaction status failed by card number")

// ErrGetMonthlyPaymentMethodsFailed indicates a failure when retrieving monthly payment method statistics.
var ErrGetMonthlyPaymentMethodsFailed = errors.New("failed to get monthly payment methods")

// ErrGetYearlyPaymentMethodsFailed indicates a failure when retrieving yearly payment method statistics.
var ErrGetYearlyPaymentMethodsFailed = errors.New("failed to get yearly payment methods")

// ErrGetMonthlyAmountsFailed indicates a failure when retrieving total monthly transaction amounts.
var ErrGetMonthlyAmountsFailed = errors.New("failed to get monthly amounts")

// ErrGetYearlyAmountsFailed indicates a failure when retrieving total yearly transaction amounts.
var ErrGetYearlyAmountsFailed = errors.New("failed to get yearly amounts")

// ErrGetMonthlyPaymentMethodsByCardFailed indicates a failure when retrieving monthly payment methods by card number.
var ErrGetMonthlyPaymentMethodsByCardFailed = errors.New("failed to get monthly payment methods by card number")

// ErrGetYearlyPaymentMethodsByCardFailed indicates a failure when retrieving yearly payment methods by card number.
var ErrGetYearlyPaymentMethodsByCardFailed = errors.New("failed to get yearly payment methods by card number")

// ErrGetMonthlyAmountsByCardFailed indicates a failure when retrieving monthly amounts by card number.
var ErrGetMonthlyAmountsByCardFailed = errors.New("failed to get monthly amounts by card number")

// ErrGetYearlyAmountsByCardFailed indicates a failure when retrieving yearly amounts by card number.
var ErrGetYearlyAmountsByCardFailed = errors.New("failed to get yearly amounts by card number")

// ErrFindTransactionByMerchantIdFailed indicates a failure when retrieving transactions by merchant ID.
var ErrFindTransactionByMerchantIdFailed = errors.New("failed to find transaction by merchant ID")

// ErrCreateTransactionFailed indicates a failure when creating a new transaction.
var ErrCreateTransactionFailed = errors.New("failed to create transaction")

// ErrUpdateTransactionFailed indicates a failure when updating a transaction.
var ErrUpdateTransactionFailed = errors.New("failed to update transaction")

// ErrUpdateTransactionStatusFailed indicates a failure when updating the status of a transaction.
var ErrUpdateTransactionStatusFailed = errors.New("failed to update transaction status")

// ErrTrashedTransactionFailed indicates a failure when soft-deleting (trashing) a transaction.
var ErrTrashedTransactionFailed = errors.New("failed to soft-delete (trash) transaction")

// ErrRestoreTransactionFailed indicates a failure when restoring a trashed transaction.
var ErrRestoreTransactionFailed = errors.New("failed to restore transaction")

// ErrDeleteTransactionPermanentFailed indicates a failure when permanently deleting a transaction.
var ErrDeleteTransactionPermanentFailed = errors.New("failed to permanently delete transaction")

// ErrRestoreAllTransactionsFailed indicates a failure when restoring all trashed transactions.
var ErrRestoreAllTransactionsFailed = errors.New("failed to restore all transactions")

// ErrDeleteAllTransactionsPermanentFailed indicates a failure when permanently deleting all transactions.
var ErrDeleteAllTransactionsPermanentFailed = errors.New("failed to permanently delete all transactions")
