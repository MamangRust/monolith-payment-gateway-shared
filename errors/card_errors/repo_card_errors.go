package card_errors

import "errors"

// ErrFindAllCardsFailed is returned when fetching all card records fails.
var ErrFindAllCardsFailed = errors.New("failed to find all cards")

// ErrFindActiveCardsFailed is returned when fetching active card records fails.
var ErrFindActiveCardsFailed = errors.New("failed to find active cards")

// ErrFindTrashedCardsFailed is returned when fetching trashed card records fails.
var ErrFindTrashedCardsFailed = errors.New("failed to find trashed cards")

// ErrFindCardByIdFailed is returned when fetching a card by its ID fails.
var ErrFindCardByIdFailed = errors.New("failed to find card by ID")

// ErrFindCardByUserIdFailed is returned when fetching a card by user ID fails.
var ErrFindCardByUserIdFailed = errors.New("failed to find card by user ID")

// ErrFindCardByCardNumberFailed is returned when fetching a card by card number fails.
var ErrFindCardByCardNumberFailed = errors.New("failed to find card by card number")

// ErrGetTotalBalancesFailed is returned when fetching total balances fails.
var ErrGetTotalBalancesFailed = errors.New("failed to get total balances")

// ErrGetTotalTopAmountFailed is returned when fetching the total top-up amount fails.
var ErrGetTotalTopAmountFailed = errors.New("failed to get total topup amount")

// ErrGetTotalWithdrawAmountFailed is returned when fetching the total withdrawal amount fails.
var ErrGetTotalWithdrawAmountFailed = errors.New("failed to get total withdraw amount")

// ErrGetTotalTransactionAmountFailed is returned when fetching the total transaction amount fails.
var ErrGetTotalTransactionAmountFailed = errors.New("failed to get total transaction amount")

// ErrGetTotalTransferAmountFailed is returned when fetching the total transfer amount fails.
var ErrGetTotalTransferAmountFailed = errors.New("failed to get total transfer amount")

// ErrGetTotalBalanceByCardFailed is returned when fetching the total balance by card fails.
var ErrGetTotalBalanceByCardFailed = errors.New("failed to get total balance by card")

// ErrGetTotalTopupAmountByCardFailed is returned when fetching the total top-up amount by card fails.
var ErrGetTotalTopupAmountByCardFailed = errors.New("failed to get total topup amount by card")

// ErrGetTotalWithdrawAmountByCardFailed is returned when fetching the total withdrawal amount by card fails.
var ErrGetTotalWithdrawAmountByCardFailed = errors.New("failed to get total withdraw amount by card")

// ErrGetTotalTransactionAmountByCardFailed is returned when fetching the total transaction amount by card fails.
var ErrGetTotalTransactionAmountByCardFailed = errors.New("failed to get total transaction amount by card")

// ErrGetTotalTransferAmountBySenderFailed is returned when fetching the total transfer amount by sender fails.
var ErrGetTotalTransferAmountBySenderFailed = errors.New("failed to get total transfer amount by sender")

// ErrGetTotalTransferAmountByReceiverFailed is returned when fetching the total transfer amount by receiver fails.
var ErrGetTotalTransferAmountByReceiverFailed = errors.New("failed to get total transfer amount by receiver")

// ErrGetMonthlyBalanceFailed is returned when fetching monthly balances fails.
var ErrGetMonthlyBalanceFailed = errors.New("failed to get monthly balance")

// ErrGetYearlyBalanceFailed is returned when fetching yearly balances fails.
var ErrGetYearlyBalanceFailed = errors.New("failed to get yearly balance")

// ErrGetMonthlyTopupAmountFailed is returned when fetching monthly top-up amounts fails.
var ErrGetMonthlyTopupAmountFailed = errors.New("failed to get monthly topup amount")

// ErrGetYearlyTopupAmountFailed is returned when fetching yearly top-up amounts fails.
var ErrGetYearlyTopupAmountFailed = errors.New("failed to get yearly topup amount")

// ErrGetMonthlyWithdrawAmountFailed is returned when fetching monthly withdrawal amounts fails.
var ErrGetMonthlyWithdrawAmountFailed = errors.New("failed to get monthly withdraw amount")

// ErrGetYearlyWithdrawAmountFailed is returned when fetching yearly withdrawal amounts fails.
var ErrGetYearlyWithdrawAmountFailed = errors.New("failed to get yearly withdraw amount")

// ErrGetMonthlyTransactionAmountFailed is returned when fetching monthly transaction amounts fails.
var ErrGetMonthlyTransactionAmountFailed = errors.New("failed to get monthly transaction amount")

// ErrGetYearlyTransactionAmountFailed is returned when fetching yearly transaction amounts fails.
var ErrGetYearlyTransactionAmountFailed = errors.New("failed to get yearly transaction amount")

// ErrGetMonthlyTransferAmountSenderFailed is returned when fetching monthly transfer amounts by sender fails.
var ErrGetMonthlyTransferAmountSenderFailed = errors.New("failed to get monthly transfer amount by sender")

// ErrGetYearlyTransferAmountSenderFailed is returned when fetching yearly transfer amounts by sender fails.
var ErrGetYearlyTransferAmountSenderFailed = errors.New("failed to get yearly transfer amount by sender")

// ErrGetMonthlyTransferAmountReceiverFailed is returned when fetching monthly transfer amounts by receiver fails.
var ErrGetMonthlyTransferAmountReceiverFailed = errors.New("failed to get monthly transfer amount by receiver")

// ErrGetYearlyTransferAmountReceiverFailed is returned when fetching yearly transfer amounts by receiver fails.
var ErrGetYearlyTransferAmountReceiverFailed = errors.New("failed to get yearly transfer amount by receiver")

// ErrGetMonthlyBalanceByCardFailed is returned when fetching monthly balances by card fails.
var ErrGetMonthlyBalanceByCardFailed = errors.New("failed to get monthly balance by card")

// ErrGetYearlyBalanceByCardFailed is returned when fetching yearly balances by card fails.
var ErrGetYearlyBalanceByCardFailed = errors.New("failed to get yearly balance by card")

// ErrGetMonthlyTopupAmountByCardFailed is returned when fetching monthly top-up amounts by card fails.
var ErrGetMonthlyTopupAmountByCardFailed = errors.New("failed to get monthly topup amount by card")

// ErrGetYearlyTopupAmountByCardFailed is returned when fetching yearly top-up amounts by card fails.
var ErrGetYearlyTopupAmountByCardFailed = errors.New("failed to get yearly topup amount by card")

// ErrGetMonthlyWithdrawAmountByCardFailed is returned when fetching monthly withdrawal amounts by card fails.
var ErrGetMonthlyWithdrawAmountByCardFailed = errors.New("failed to get monthly withdraw amount by card")

// ErrGetYearlyWithdrawAmountByCardFailed is returned when fetching yearly withdrawal amounts by card fails.
var ErrGetYearlyWithdrawAmountByCardFailed = errors.New("failed to get yearly withdraw amount by card")

// ErrGetMonthlyTransactionAmountByCardFailed is returned when fetching monthly transaction amounts by card fails.
var ErrGetMonthlyTransactionAmountByCardFailed = errors.New("failed to get monthly transaction amount by card")

// ErrGetYearlyTransactionAmountByCardFailed is returned when fetching yearly transaction amounts by card fails.
var ErrGetYearlyTransactionAmountByCardFailed = errors.New("failed to get yearly transaction amount by card")

// ErrGetMonthlyTransferAmountBySenderFailed is returned when fetching monthly transfer amount by sender fails.
var ErrGetMonthlyTransferAmountBySenderFailed = errors.New("failed to get monthly transfer amount by sender")

// ErrGetYearlyTransferAmountBySenderFailed is returned when fetching yearly transfer amount by sender fails.
var ErrGetYearlyTransferAmountBySenderFailed = errors.New("failed to get yearly transfer amount by sender")

// ErrGetMonthlyTransferAmountByReceiverFailed is returned when fetching monthly transfer amount by receiver fails.
var ErrGetMonthlyTransferAmountByReceiverFailed = errors.New("failed to get monthly transfer amount by receiver")

// ErrGetYearlyTransferAmountByReceiverFailed is returned when fetching yearly transfer amount by receiver fails.
var ErrGetYearlyTransferAmountByReceiverFailed = errors.New("failed to get yearly transfer amount by receiver")

// ErrCreateCardFailed is returned when creating a new card fails.
var ErrCreateCardFailed = errors.New("failed to create card")

// ErrUpdateCardFailed is returned when updating a card fails.
var ErrUpdateCardFailed = errors.New("failed to update card")

// ErrTrashCardFailed is returned when trashing a card fails.
var ErrTrashCardFailed = errors.New("failed to trash card")

// ErrRestoreCardFailed is returned when restoring a trashed card fails.
var ErrRestoreCardFailed = errors.New("failed to restore card")

// ErrDeleteCardPermanentFailed is returned when permanently deleting a card fails.
var ErrDeleteCardPermanentFailed = errors.New("failed to delete card permanently")

// ErrRestoreAllCardsFailed is returned when restoring all trashed cards fails.
var ErrRestoreAllCardsFailed = errors.New("failed to restore all cards")

// ErrDeleteAllCardsPermanentFailed is returned when permanently deleting all cards fails.
var ErrDeleteAllCardsPermanentFailed = errors.New("failed to delete all cards permanently")

// ErrInvalidCardRequest is returned when the card request data is invalid.
var ErrInvalidCardRequest = errors.New("invalid card request data")

// ErrInvalidCardId is returned when the card ID is invalid.
var ErrInvalidCardId = errors.New("invalid card ID")

// ErrInvalidUserId is returned when the user ID is invalid.
var ErrInvalidUserId = errors.New("invalid user ID")

// ErrInvalidCardNumber is returned when the card number is invalid.
var ErrInvalidCardNumber = errors.New("invalid card number")

// ErrCardAlreadyExists is returned when a card already exists.
var ErrCardAlreadyExists = errors.New("card already exists")

// ErrCardNotFound is returned when a card is not found.
var ErrCardNotFound = errors.New("card not found")
