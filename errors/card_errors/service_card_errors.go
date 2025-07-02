package card_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrCardNotFoundRes is an error response when a requested card was not found.
// Its HTTP status code is 404.
var ErrCardNotFoundRes = response.NewErrorResponse("Card not found", http.StatusNotFound)

// ErrFailedFindAllCards is an error response when retrieving all card records
// fails. Its HTTP status code is 500.
var ErrFailedFindAllCards = response.NewErrorResponse("Failed to fetch Cards", http.StatusInternalServerError)

// ErrFailedFindActiveCards is an error response when retrieving active card
// records fails. Its HTTP status code is 500.
var ErrFailedFindActiveCards = response.NewErrorResponse("Failed to fetch active Cards", http.StatusInternalServerError)

// ErrFailedFindTrashedCards is an error response when retrieving trashed card
// records fails. Its HTTP status code is 500.
var ErrFailedFindTrashedCards = response.NewErrorResponse("Failed to fetch trashed Cards", http.StatusInternalServerError)

// ErrFailedFindById is an error response when finding a card by its ID fails.
// Its HTTP status code is 500.
var ErrFailedFindById = response.NewErrorResponse("Failed to find Card by ID", http.StatusInternalServerError)

// ErrFailedFindByUserID is an error response when finding a card by its user ID
// fails. Its HTTP status code is 500.
var ErrFailedFindByUserID = response.NewErrorResponse("Failed to find Card by User ID", http.StatusInternalServerError)

// ErrFailedFindByCardNumber is an error response when finding a card by its card
// number fails. Its HTTP status code is 500.
var ErrFailedFindByCardNumber = response.NewErrorResponse("Failed to find Card by Card Number", http.StatusInternalServerError)

// ErrFailedFindTotalBalances is an error response when retrieving total balances
// fails. Its HTTP status code is 500.
var ErrFailedFindTotalBalances = response.NewErrorResponse("Failed to Find total balances", http.StatusInternalServerError)

// ErrFailedFindTotalTopAmount is an error response when retrieving total topup
// amount fails. Its HTTP status code is 500.
var ErrFailedFindTotalTopAmount = response.NewErrorResponse("Failed to Find total topup amount", http.StatusInternalServerError)

// ErrFailedFindTotalWithdrawAmount is an error response when retrieving total
// withdraw amount fails. Its HTTP status code is 500.
var ErrFailedFindTotalWithdrawAmount = response.NewErrorResponse("Failed to Find total withdraw amount", http.StatusInternalServerError)

// ErrFailedFindTotalTransactionAmount is an error response when retrieving total
// transaction amount fails. Its HTTP status code is 500.
var ErrFailedFindTotalTransactionAmount = response.NewErrorResponse("Failed to Find total transaction amount", http.StatusInternalServerError)

// ErrFailedFindTotalTransferAmount is an error response when retrieving total
// transfer amount fails. Its HTTP status code is 500.
var ErrFailedFindTotalTransferAmount = response.NewErrorResponse("Failed to Find total transfer amount", http.StatusInternalServerError)

// ErrFailedFindTotalBalanceByCard is an error response when retrieving total
// balance by card fails. Its HTTP status code is 500.
var ErrFailedFindTotalBalanceByCard = response.NewErrorResponse("Failed to Find total balance by card", http.StatusInternalServerError)

// ErrFailedFindTotalTopupAmountByCard is an error response when retrieving total
// topup amount by card fails. Its HTTP status code is 500.
var ErrFailedFindTotalTopupAmountByCard = response.NewErrorResponse("Failed to Find total topup amount by card", http.StatusInternalServerError)

// ErrFailedFindTotalWithdrawAmountByCard is an error response when retrieving
// total withdraw amount by card fails. Its HTTP status code is 500.
var ErrFailedFindTotalWithdrawAmountByCard = response.NewErrorResponse("Failed to Find total withdraw amount by card", http.StatusInternalServerError)

// ErrFailedFindTotalTransactionAmountByCard is an error response when
// retrieving total transaction amount by card fails. Its HTTP status code is
// 500.
var ErrFailedFindTotalTransactionAmountByCard = response.NewErrorResponse("Failed to Find total transaction amount by card", http.StatusInternalServerError)

// ErrFailedFindTotalTransferAmountBySender is an error response when
// retrieving total transfer amount by sender fails. Its HTTP status code is
// 500.
var ErrFailedFindTotalTransferAmountBySender = response.NewErrorResponse("Failed to Find total transfer amount by sender", http.StatusInternalServerError)

// ErrFailedFindTotalTransferAmountByReceiver is an error response when
// retrieving total transfer amount by receiver fails. Its HTTP status code is
// 500.
var ErrFailedFindTotalTransferAmountByReceiver = response.NewErrorResponse("Failed to Find total transfer amount by receiver", http.StatusInternalServerError)

// ErrFailedDashboardCard is an error response when retrieving card dashboard
// fails. Its HTTP status code is 500.
var ErrFailedDashboardCard = response.NewErrorResponse("Failed to get Card dashboard", http.StatusInternalServerError)

// ErrFailedDashboardCardNumber is an error response when retrieving card
// dashboard by card number fails. Its HTTP status code is 500.
var ErrFailedDashboardCardNumber = response.NewErrorResponse("Failed to get Card dashboard by card number", http.StatusInternalServerError)

// ErrFailedFindMonthlyBalance indicates a failure in retrieving the monthly balance.
var ErrFailedFindMonthlyBalance = response.NewErrorResponse("Failed to get monthly balance", http.StatusInternalServerError)

// ErrFailedFindYearlyBalance indicates a failure in retrieving the yearly balance.
var ErrFailedFindYearlyBalance = response.NewErrorResponse("Failed to get yearly balance", http.StatusInternalServerError)

// ErrFailedFindMonthlyTopupAmount indicates a failure in retrieving the monthly top-up amount.
var ErrFailedFindMonthlyTopupAmount = response.NewErrorResponse("Failed to get monthly topup amount", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupAmount indicates a failure in retrieving the yearly top-up amount.
var ErrFailedFindYearlyTopupAmount = response.NewErrorResponse("Failed to get yearly topup amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyWithdrawAmount indicates a failure in retrieving the monthly withdraw amount.
var ErrFailedFindMonthlyWithdrawAmount = response.NewErrorResponse("Failed to get monthly withdraw amount", http.StatusInternalServerError)

// ErrFailedFindYearlyWithdrawAmount indicates a failure in retrieving the yearly withdraw amount.
var ErrFailedFindYearlyWithdrawAmount = response.NewErrorResponse("Failed to get yearly withdraw amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransactionAmount indicates a failure in retrieving the monthly transaction amount.
var ErrFailedFindMonthlyTransactionAmount = response.NewErrorResponse("Failed to get monthly transaction amount", http.StatusInternalServerError)

// ErrFailedFindYearlyTransactionAmount indicates a failure in retrieving the yearly transaction amount.
var ErrFailedFindYearlyTransactionAmount = response.NewErrorResponse("Failed to get yearly transaction amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountSender indicates a failure in retrieving the monthly transfer amount by sender.
var ErrFailedFindMonthlyTransferAmountSender = response.NewErrorResponse("Failed to get monthly transfer amount by sender", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountSender indicates a failure in retrieving the yearly transfer amount by sender.
var ErrFailedFindYearlyTransferAmountSender = response.NewErrorResponse("Failed to get yearly transfer amount by sender", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountReceiver indicates a failure in retrieving the monthly transfer amount by receiver.
var ErrFailedFindMonthlyTransferAmountReceiver = response.NewErrorResponse("Failed to get monthly transfer amount by receiver", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountReceiver indicates a failure in retrieving the yearly transfer amount by receiver.
var ErrFailedFindYearlyTransferAmountReceiver = response.NewErrorResponse("Failed to get yearly transfer amount by receiver", http.StatusInternalServerError)

// ErrFailedFindMonthlyBalanceByCard returns an error response when retrieving
// monthly balance by card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyBalanceByCard = response.NewErrorResponse("Failed to get monthly balance by card", http.StatusInternalServerError)

// ErrFailedFindYearlyBalanceByCard returns an error response when retrieving
// yearly balance by card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyBalanceByCard = response.NewErrorResponse("Failed to get yearly balance by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTopupAmountByCard returns an error response when retrieving
// monthly top-up amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyTopupAmountByCard = response.NewErrorResponse("Failed to get monthly topup amount by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupAmountByCard returns an error response when retrieving
// yearly top-up amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyTopupAmountByCard = response.NewErrorResponse("Failed to get yearly topup amount by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyWithdrawAmountByCard returns an error response when retrieving
// monthly withdraw amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyWithdrawAmountByCard = response.NewErrorResponse("Failed to get monthly withdraw amount by card", http.StatusInternalServerError)

// ErrFailedFindYearlyWithdrawAmountByCard returns an error response when retrieving
// yearly withdraw amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyWithdrawAmountByCard = response.NewErrorResponse("Failed to get yearly withdraw amount by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransactionAmountByCard returns an error response when retrieving
// monthly transaction amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyTransactionAmountByCard = response.NewErrorResponse("Failed to get monthly transaction amount by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTransactionAmountByCard returns an error response when retrieving
// yearly transaction amount by card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyTransactionAmountByCard = response.NewErrorResponse("Failed to get yearly transaction amount by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountBySender returns an error response when retrieving
// monthly transfer amount by sender card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyTransferAmountBySender = response.NewErrorResponse("Failed to get monthly transfer amount by sender", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountBySender returns an error response when retrieving
// yearly transfer amount by sender card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyTransferAmountBySender = response.NewErrorResponse("Failed to get yearly transfer amount by sender", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountByReceiver returns an error response when retrieving
// monthly transfer amount by receiver card number fails. Its HTTP status code is 500.
var ErrFailedFindMonthlyTransferAmountByReceiver = response.NewErrorResponse("Failed to get monthly transfer amount by receiver", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountByReceiver returns an error response when retrieving
// yearly transfer amount by receiver card number fails. Its HTTP status code is 500.
var ErrFailedFindYearlyTransferAmountByReceiver = response.NewErrorResponse("Failed to get yearly transfer amount by receiver", http.StatusInternalServerError)

// ErrFailedCreateCard is returned when creating a new Card record fails.
// Its HTTP status code is 500.
var ErrFailedCreateCard = response.NewErrorResponse("Failed to create Card", http.StatusInternalServerError)

// ErrFailedUpdateCard is returned when updating an existing Card record fails.
// Its HTTP status code is 500.
var ErrFailedUpdateCard = response.NewErrorResponse("Failed to update Card", http.StatusInternalServerError)

// ErrFailedTrashCard is returned when moving a Card to trash fails.
// Its HTTP status code is 500.
var ErrFailedTrashCard = response.NewErrorResponse("Failed to trash Card", http.StatusInternalServerError)

// ErrFailedRestoreCard is returned when restoring a trashed Card fails.
// Its HTTP status code is 500.
var ErrFailedRestoreCard = response.NewErrorResponse("Failed to restore Card", http.StatusInternalServerError)

// ErrFailedDeleteCard is returned when permanently deleting a Card fails.
// Its HTTP status code is 500.
var ErrFailedDeleteCard = response.NewErrorResponse("Failed to delete Card permanently", http.StatusInternalServerError)

// ErrFailedRestoreAllCards is returned when restoring all trashed Cards fails.
// Its HTTP status code is 500.
var ErrFailedRestoreAllCards = response.NewErrorResponse("Failed to restore all Cards", http.StatusInternalServerError)

// ErrFailedDeleteAllCards is returned when permanently deleting all Cards fails.
// Its HTTP status code is 500.
var ErrFailedDeleteAllCards = response.NewErrorResponse("Failed to delete all Cards permanently", http.StatusInternalServerError)
