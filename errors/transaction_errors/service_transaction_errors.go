package transaction_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrFailedFindAllTransactions indicates a failure when retrieving all transaction records.
var ErrFailedFindAllTransactions = response.NewErrorResponse("Failed to fetch all transactions", http.StatusInternalServerError)

// ErrFailedFindAllByCardNumber indicates a failure when retrieving transactions filtered by card number.
var ErrFailedFindAllByCardNumber = response.NewErrorResponse("Failed to fetch transactions by card number", http.StatusInternalServerError)

// ErrTransactionNotFound indicates that the requested transaction could not be found.
var ErrTransactionNotFound = response.NewErrorResponse("Transaction not found", http.StatusNotFound)

// ErrFailedFindMonthTransactionSuccess indicates a failure when retrieving monthly successful transactions.
var ErrFailedFindMonthTransactionSuccess = response.NewErrorResponse("Failed to fetch monthly successful transactions", http.StatusInternalServerError)

// ErrFailedFindYearTransactionSuccess indicates a failure when retrieving yearly successful transactions.
var ErrFailedFindYearTransactionSuccess = response.NewErrorResponse("Failed to fetch yearly successful transactions", http.StatusInternalServerError)

// ErrFailedFindMonthTransactionFailed indicates a failure when retrieving monthly failed transactions.
var ErrFailedFindMonthTransactionFailed = response.NewErrorResponse("Failed to fetch monthly failed transactions", http.StatusInternalServerError)

// ErrFailedFindYearTransactionFailed indicates a failure when retrieving yearly failed transactions.
var ErrFailedFindYearTransactionFailed = response.NewErrorResponse("Failed to fetch yearly failed transactions", http.StatusInternalServerError)

// ErrFailedFindMonthTransactionSuccessByCard indicates a failure when retrieving monthly successful transactions filtered by card number.
var ErrFailedFindMonthTransactionSuccessByCard = response.NewErrorResponse("Failed to fetch monthly successful transactions by card", http.StatusInternalServerError)

// ErrFailedFindYearTransactionSuccessByCard indicates a failure when retrieving yearly successful transactions filtered by card number.
var ErrFailedFindYearTransactionSuccessByCard = response.NewErrorResponse("Failed to fetch yearly successful transactions by card", http.StatusInternalServerError)

// ErrFailedFindMonthTransactionFailedByCard indicates a failure when retrieving monthly failed transactions filtered by card number.
var ErrFailedFindMonthTransactionFailedByCard = response.NewErrorResponse("Failed to fetch monthly failed transactions by card", http.StatusInternalServerError)

// ErrFailedFindYearTransactionFailedByCard indicates a failure when retrieving yearly failed transactions filtered by card number.
var ErrFailedFindYearTransactionFailedByCard = response.NewErrorResponse("Failed to fetch yearly failed transactions by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyPaymentMethods indicates a failure when retrieving monthly statistics of payment methods used.
var ErrFailedFindMonthlyPaymentMethods = response.NewErrorResponse("Failed to fetch monthly payment methods", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethods indicates a failure when retrieving yearly statistics of payment methods used.
var ErrFailedFindYearlyPaymentMethods = response.NewErrorResponse("Failed to fetch yearly payment methods", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmounts indicates a failure when retrieving the total monthly transaction amounts.
var ErrFailedFindMonthlyAmounts = response.NewErrorResponse("Failed to fetch monthly amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyAmounts indicates a failure when retrieving the total yearly transaction amounts.
var ErrFailedFindYearlyAmounts = response.NewErrorResponse("Failed to fetch yearly amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyPaymentMethodsByCard indicates a failure when retrieving monthly payment methods filtered by card.
var ErrFailedFindMonthlyPaymentMethodsByCard = response.NewErrorResponse("Failed to fetch monthly payment methods by card", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodsByCard indicates a failure when retrieving yearly payment methods filtered by card.
var ErrFailedFindYearlyPaymentMethodsByCard = response.NewErrorResponse("Failed to fetch yearly payment methods by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmountsByCard indicates a failure when retrieving monthly transaction amounts filtered by card.
var ErrFailedFindMonthlyAmountsByCard = response.NewErrorResponse("Failed to fetch monthly amounts by card", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountsByCard indicates a failure when retrieving yearly transaction amounts filtered by card.
var ErrFailedFindYearlyAmountsByCard = response.NewErrorResponse("Failed to fetch yearly amounts by card", http.StatusInternalServerError)

// ErrFailedFindByActiveTransactions indicates a failure when retrieving active (non-deleted) transactions.
var ErrFailedFindByActiveTransactions = response.NewErrorResponse("Failed to fetch active transactions", http.StatusInternalServerError)

// ErrFailedFindByTrashedTransactions indicates a failure when retrieving trashed (soft-deleted) transactions.
var ErrFailedFindByTrashedTransactions = response.NewErrorResponse("Failed to fetch trashed transactions", http.StatusInternalServerError)

// ErrFailedFindByMerchantID indicates a failure when retrieving transactions filtered by merchant ID.
var ErrFailedFindByMerchantID = response.NewErrorResponse("Failed to fetch transactions by merchant ID", http.StatusInternalServerError)

// ErrFailedCreateTransaction indicates a failure when creating a new transaction record.
var ErrFailedCreateTransaction = response.NewErrorResponse("Failed to create transaction", http.StatusInternalServerError)

// ErrFailedUpdateTransaction indicates a failure when updating an existing transaction record.
var ErrFailedUpdateTransaction = response.NewErrorResponse("Failed to update transaction", http.StatusInternalServerError)

// ErrFailedTrashedTransaction indicates a failure when soft-deleting (trashing) a transaction.
var ErrFailedTrashedTransaction = response.NewErrorResponse("Failed to trash transaction", http.StatusInternalServerError)

// ErrFailedRestoreTransaction indicates a failure when restoring a previously trashed transaction.
var ErrFailedRestoreTransaction = response.NewErrorResponse("Failed to restore transaction", http.StatusInternalServerError)

// ErrFailedDeleteTransactionPermanent indicates a failure when permanently deleting a transaction.
var ErrFailedDeleteTransactionPermanent = response.NewErrorResponse("Failed to permanently delete transaction", http.StatusInternalServerError)

// ErrFailedRestoreAllTransactions indicates a failure when restoring all trashed transactions.
var ErrFailedRestoreAllTransactions = response.NewErrorResponse("Failed to restore all transactions", http.StatusInternalServerError)

// ErrFailedDeleteAllTransactionsPermanent indicates a failure when permanently deleting all transactions.
var ErrFailedDeleteAllTransactionsPermanent = response.NewErrorResponse("Failed to permanently delete all transactions", http.StatusInternalServerError)
