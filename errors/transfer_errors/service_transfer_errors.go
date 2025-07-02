package transfer_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrFailedFindAllTransfers indicates a failure when retrieving all transfer records.
var ErrFailedFindAllTransfers = response.NewErrorResponse("Failed to fetch all transfers", http.StatusInternalServerError)

// ErrTransferNotFound indicates that a specific transfer record was not found.
var ErrTransferNotFound = response.NewErrorResponse("Transfer not found", http.StatusNotFound)

// ErrFailedFindMonthTransferStatusSuccess indicates a failure when retrieving monthly successful transfer statistics.
var ErrFailedFindMonthTransferStatusSuccess = response.NewErrorResponse("Failed to fetch monthly successful transfers", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusSuccess indicates a failure when retrieving yearly successful transfer statistics.
var ErrFailedFindYearTransferStatusSuccess = response.NewErrorResponse("Failed to fetch yearly successful transfers", http.StatusInternalServerError)

// ErrFailedFindMonthTransferStatusFailed indicates a failure when retrieving monthly failed transfer statistics.
var ErrFailedFindMonthTransferStatusFailed = response.NewErrorResponse("Failed to fetch monthly failed transfers", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusFailed indicates a failure when retrieving yearly failed transfer statistics.
var ErrFailedFindYearTransferStatusFailed = response.NewErrorResponse("Failed to fetch yearly failed transfers", http.StatusInternalServerError)

// ErrFailedFindMonthTransferStatusSuccessByCard indicates a failure when retrieving monthly successful transfers by card number.
var ErrFailedFindMonthTransferStatusSuccessByCard = response.NewErrorResponse("Failed to fetch monthly successful transfers by card", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusSuccessByCard indicates a failure when retrieving yearly successful transfers by card number.
var ErrFailedFindYearTransferStatusSuccessByCard = response.NewErrorResponse("Failed to fetch yearly successful transfers by card", http.StatusInternalServerError)

// ErrFailedFindMonthTransferStatusFailedByCard indicates a failure when retrieving monthly failed transfers by card number.
var ErrFailedFindMonthTransferStatusFailedByCard = response.NewErrorResponse("Failed to fetch monthly failed transfers by card", http.StatusInternalServerError)

// ErrFailedFindYearTransferStatusFailedByCard indicates a failure when retrieving yearly failed transfers by card number.
var ErrFailedFindYearTransferStatusFailedByCard = response.NewErrorResponse("Failed to fetch yearly failed transfers by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmounts indicates a failure when retrieving the total monthly transfer amounts.
var ErrFailedFindMonthlyTransferAmounts = response.NewErrorResponse("Failed to fetch monthly transfer amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmounts indicates a failure when retrieving the total yearly transfer amounts.
var ErrFailedFindYearlyTransferAmounts = response.NewErrorResponse("Failed to fetch yearly transfer amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountsBySenderCard indicates a failure when retrieving monthly transfer amounts by sender card number.
var ErrFailedFindMonthlyTransferAmountsBySenderCard = response.NewErrorResponse("Failed to fetch monthly transfer amounts by sender card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTransferAmountsByReceiverCard indicates a failure when retrieving monthly transfer amounts by receiver card number.
var ErrFailedFindMonthlyTransferAmountsByReceiverCard = response.NewErrorResponse("Failed to fetch monthly transfer amounts by receiver card", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountsBySenderCard indicates a failure when retrieving yearly transfer amounts by sender card number.
var ErrFailedFindYearlyTransferAmountsBySenderCard = response.NewErrorResponse("Failed to fetch yearly transfer amounts by sender card", http.StatusInternalServerError)

// ErrFailedFindYearlyTransferAmountsByReceiverCard indicates a failure when retrieving yearly transfer amounts by receiver card number.
var ErrFailedFindYearlyTransferAmountsByReceiverCard = response.NewErrorResponse("Failed to fetch yearly transfer amounts by receiver card", http.StatusInternalServerError)

// ErrFailedFindActiveTransfers indicates a failure when retrieving active transfer records.
var ErrFailedFindActiveTransfers = response.NewErrorResponse("Failed to fetch active transfers", http.StatusInternalServerError)

// ErrFailedFindTrashedTransfers indicates a failure when retrieving trashed (soft-deleted) transfer records.
var ErrFailedFindTrashedTransfers = response.NewErrorResponse("Failed to fetch trashed transfers", http.StatusInternalServerError)

// ErrFailedFindTransfersBySender indicates a failure when retrieving transfers filtered by sender card.
var ErrFailedFindTransfersBySender = response.NewErrorResponse("Failed to fetch transfers by sender", http.StatusInternalServerError)

// ErrFailedFindTransfersByReceiver indicates a failure when retrieving transfers filtered by receiver card.
var ErrFailedFindTransfersByReceiver = response.NewErrorResponse("Failed to fetch transfers by receiver", http.StatusInternalServerError)

// ErrFailedCreateTransfer indicates a failure when attempting to create a new transfer record.
var ErrFailedCreateTransfer = response.NewErrorResponse("Failed to create transfer", http.StatusInternalServerError)

// ErrFailedUpdateTransfer indicates a failure when attempting to update an existing transfer record.
var ErrFailedUpdateTransfer = response.NewErrorResponse("Failed to update transfer", http.StatusInternalServerError)

// ErrFailedTrashedTransfer indicates a failure when attempting to soft-delete (trash) a transfer.
var ErrFailedTrashedTransfer = response.NewErrorResponse("Failed to trash transfer", http.StatusInternalServerError)

// ErrFailedRestoreTransfer indicates a failure when attempting to restore a previously trashed transfer.
var ErrFailedRestoreTransfer = response.NewErrorResponse("Failed to restore transfer", http.StatusInternalServerError)

// ErrFailedDeleteTransferPermanent indicates a failure when attempting to permanently delete a transfer.
var ErrFailedDeleteTransferPermanent = response.NewErrorResponse("Failed to permanently delete transfer", http.StatusInternalServerError)

// ErrFailedRestoreAllTransfers indicates a failure when attempting to restore all trashed transfers.
var ErrFailedRestoreAllTransfers = response.NewErrorResponse("Failed to restore all transfers", http.StatusInternalServerError)

// ErrFailedDeleteAllTransfersPermanent indicates a failure when attempting to permanently delete all transfer records.
var ErrFailedDeleteAllTransfersPermanent = response.NewErrorResponse("Failed to permanently delete all transfers", http.StatusInternalServerError)
