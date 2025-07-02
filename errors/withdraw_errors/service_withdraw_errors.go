package withdraw_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrFailedFindAllWithdraws is used when failed to fetch all withdraws
var ErrFailedFindAllWithdraws = response.NewErrorResponse("Failed to fetch all withdraws", http.StatusInternalServerError)

// ErrWithdrawNotFound is used when withdraw is not found
var ErrWithdrawNotFound = response.NewErrorResponse("Withdraw not found", http.StatusNotFound)

// ErrFailedFindAllWithdrawsByCard is used when failed to fetch all withdraws by card number
var ErrFailedFindAllWithdrawsByCard = response.NewErrorResponse("Failed to fetch all withdraws by card number", http.StatusInternalServerError)

// ErrFailedSendEmail is used when failed to send email
var ErrFailedSendEmail = response.NewErrorResponse("Failed to send email", http.StatusInternalServerError)

// ErrFailedFindMonthWithdrawStatusSuccess is used when failed to fetch monthly successful withdraws
var ErrFailedFindMonthWithdrawStatusSuccess = response.NewErrorResponse("Failed to fetch monthly successful withdraws", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusSuccess is used when failed to fetch yearly successful withdraws
var ErrFailedFindYearWithdrawStatusSuccess = response.NewErrorResponse("Failed to fetch yearly successful withdraws", http.StatusInternalServerError)

// ErrFailedFindMonthWithdrawStatusFailed is used when failed to fetch monthly failed withdraws
var ErrFailedFindMonthWithdrawStatusFailed = response.NewErrorResponse("Failed to fetch monthly failed withdraws", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusFailed is used when failed to fetch yearly failed withdraws
var ErrFailedFindYearWithdrawStatusFailed = response.NewErrorResponse("Failed to fetch yearly failed withdraws", http.StatusInternalServerError)

// ErrFailedFindMonthWithdrawStatusSuccessByCard is used when failed to fetch monthly successful withdraws by card
var ErrFailedFindMonthWithdrawStatusSuccessByCard = response.NewErrorResponse("Failed to fetch monthly successful withdraws by card", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusSuccessByCard is used when failed to fetch yearly successful withdraws by card
var ErrFailedFindYearWithdrawStatusSuccessByCard = response.NewErrorResponse("Failed to fetch yearly successful withdraws by card", http.StatusInternalServerError)

// ErrFailedFindMonthWithdrawStatusFailedByCard is used when failed to fetch monthly failed withdraws by card
var ErrFailedFindMonthWithdrawStatusFailedByCard = response.NewErrorResponse("Failed to fetch monthly failed withdraws by card", http.StatusInternalServerError)

// ErrFailedFindYearWithdrawStatusFailedByCard is used when failed to fetch yearly failed withdraws by card
var ErrFailedFindYearWithdrawStatusFailedByCard = response.NewErrorResponse("Failed to fetch yearly failed withdraws by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyWithdraws is used when failed to fetch monthly withdraw amounts
var ErrFailedFindMonthlyWithdraws = response.NewErrorResponse("Failed to fetch monthly withdraw amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyWithdraws is used when failed to fetch yearly withdraw amounts
var ErrFailedFindYearlyWithdraws = response.NewErrorResponse("Failed to fetch yearly withdraw amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyWithdrawsByCardNumber is used when failed to fetch monthly withdraw amounts by card
var ErrFailedFindMonthlyWithdrawsByCardNumber = response.NewErrorResponse("Failed to fetch monthly withdraw amounts by card", http.StatusInternalServerError)

// ErrFailedFindYearlyWithdrawsByCardNumber is used when failed to fetch yearly withdraw amounts by card
var ErrFailedFindYearlyWithdrawsByCardNumber = response.NewErrorResponse("Failed to fetch yearly withdraw amounts by card", http.StatusInternalServerError)

// ErrFailedFindActiveWithdraws is used when failed to fetch active withdraws
var ErrFailedFindActiveWithdraws = response.NewErrorResponse("Failed to fetch active withdraws", http.StatusInternalServerError)

// ErrFailedFindTrashedWithdraws is used when failed to fetch trashed withdraws
var ErrFailedFindTrashedWithdraws = response.NewErrorResponse("Failed to fetch trashed withdraws", http.StatusInternalServerError)

// ErrFailedCreateWithdraw is used when failed to create withdraw
var ErrFailedCreateWithdraw = response.NewErrorResponse("Failed to create withdraw", http.StatusInternalServerError)

// ErrFailedUpdateWithdraw is used when failed to update withdraw
var ErrFailedUpdateWithdraw = response.NewErrorResponse("Failed to update withdraw", http.StatusInternalServerError)

// ErrFailedTrashedWithdraw is used when failed to trash withdraw
var ErrFailedTrashedWithdraw = response.NewErrorResponse("Failed to trash withdraw", http.StatusInternalServerError)

// ErrFailedRestoreWithdraw is used when failed to restore withdraw
var ErrFailedRestoreWithdraw = response.NewErrorResponse("Failed to restore withdraw", http.StatusInternalServerError)

// ErrFailedDeleteWithdrawPermanent is used when failed to permanently delete withdraw
var ErrFailedDeleteWithdrawPermanent = response.NewErrorResponse("Failed to permanently delete withdraw", http.StatusInternalServerError)

// ErrFailedRestoreAllWithdraw is used when failed to restore all withdraws
var ErrFailedRestoreAllWithdraw = response.NewErrorResponse("Failed to restore all withdraws", http.StatusInternalServerError)

// ErrFailedDeleteAllWithdrawPermanent is used when failed to permanently delete all withdraws
var ErrFailedDeleteAllWithdrawPermanent = response.NewErrorResponse("Failed to permanently delete all withdraws", http.StatusInternalServerError)
