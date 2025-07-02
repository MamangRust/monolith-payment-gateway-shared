package transfer_errors

import "errors"

// ErrFindAllTransfersFailed indicates a failure when retrieving all transfer records.
var ErrFindAllTransfersFailed = errors.New("failed to find all transfers")

// ErrFindActiveTransfersFailed indicates a failure when retrieving active (non-trashed) transfer records.
var ErrFindActiveTransfersFailed = errors.New("failed to find active transfers")

// ErrFindTrashedTransfersFailed indicates a failure when retrieving trashed (soft-deleted) transfer records.
var ErrFindTrashedTransfersFailed = errors.New("failed to find trashed transfers")

// ErrFindTransferByIdFailed indicates a failure when retrieving a transfer record by its ID.
var ErrFindTransferByIdFailed = errors.New("failed to find transfer by ID")

// ErrGetMonthTransferStatusSuccessFailed indicates a failure when retrieving monthly successful transfer statistics.
var ErrGetMonthTransferStatusSuccessFailed = errors.New("failed to get monthly transfer status success")

// ErrGetYearlyTransferStatusSuccessFailed indicates a failure when retrieving yearly successful transfer statistics.
var ErrGetYearlyTransferStatusSuccessFailed = errors.New("failed to get yearly transfer status success")

// ErrGetMonthTransferStatusSuccessByCardFailed indicates a failure when retrieving monthly successful transfers filtered by card number.
var ErrGetMonthTransferStatusSuccessByCardFailed = errors.New("failed to get monthly transfer status success by card number")

// ErrGetYearlyTransferStatusSuccessByCardFailed indicates a failure when retrieving yearly successful transfers filtered by card number.
var ErrGetYearlyTransferStatusSuccessByCardFailed = errors.New("failed to get yearly transfer status success by card number")

// ErrGetMonthTransferStatusFailedFailed indicates a failure when retrieving monthly failed transfer statistics.
var ErrGetMonthTransferStatusFailedFailed = errors.New("failed to get monthly transfer status failed")

// ErrGetYearlyTransferStatusFailedFailed indicates a failure when retrieving yearly failed transfer statistics.
var ErrGetYearlyTransferStatusFailedFailed = errors.New("failed to get yearly transfer status failed")

// ErrGetMonthTransferStatusFailedByCardFailed indicates a failure when retrieving monthly failed transfers filtered by card number.
var ErrGetMonthTransferStatusFailedByCardFailed = errors.New("failed to get monthly transfer status failed by card number")

// ErrGetYearlyTransferStatusFailedByCardFailed indicates a failure when retrieving yearly failed transfers filtered by card number.
var ErrGetYearlyTransferStatusFailedByCardFailed = errors.New("failed to get yearly transfer status failed by card number")

// ErrGetMonthlyTransferAmountsFailed indicates a failure when retrieving the total amount of monthly transfers.
var ErrGetMonthlyTransferAmountsFailed = errors.New("failed to get monthly transfer amounts")

// ErrGetYearlyTransferAmountsFailed indicates a failure when retrieving the total amount of yearly transfers.
var ErrGetYearlyTransferAmountsFailed = errors.New("failed to get yearly transfer amounts")

// ErrGetMonthlyTransferAmountsBySenderCardFailed indicates a failure when retrieving monthly transfer amounts filtered by sender card number.
var ErrGetMonthlyTransferAmountsBySenderCardFailed = errors.New("failed to get monthly transfer amounts by sender card number")

// ErrGetYearlyTransferAmountsBySenderCardFailed indicates a failure when retrieving yearly transfer amounts filtered by sender card number.
var ErrGetYearlyTransferAmountsBySenderCardFailed = errors.New("failed to get yearly transfer amounts by sender card number")

// ErrGetMonthlyTransferAmountsByReceiverCardFailed indicates a failure when retrieving monthly transfer amounts filtered by receiver card number.
var ErrGetMonthlyTransferAmountsByReceiverCardFailed = errors.New("failed to get monthly transfer amounts by receiver card number")

// ErrGetYearlyTransferAmountsByReceiverCardFailed indicates a failure when retrieving yearly transfer amounts filtered by receiver card number.
var ErrGetYearlyTransferAmountsByReceiverCardFailed = errors.New("failed to get yearly transfer amounts by receiver card number")

// ErrFindTransferByTransferFromFailed indicates a failure when retrieving transfers by the sender (transfer from).
var ErrFindTransferByTransferFromFailed = errors.New("failed to find transfer by transfer from")

// ErrFindTransferByTransferToFailed indicates a failure when retrieving transfers by the receiver (transfer to).
var ErrFindTransferByTransferToFailed = errors.New("failed to find transfer by transfer to")

// ErrCreateTransferFailed indicates a failure when creating a new transfer.
var ErrCreateTransferFailed = errors.New("failed to create transfer")

// ErrUpdateTransferFailed indicates a failure when updating an existing transfer.
var ErrUpdateTransferFailed = errors.New("failed to update transfer")

// ErrUpdateTransferAmountFailed indicates a failure when updating the amount of a transfer.
var ErrUpdateTransferAmountFailed = errors.New("failed to update transfer amount")

// ErrUpdateTransferStatusFailed indicates a failure when updating the status of a transfer.
var ErrUpdateTransferStatusFailed = errors.New("failed to update transfer status")

// ErrTrashedTransferFailed indicates a failure when soft-deleting (trashing) a transfer.
var ErrTrashedTransferFailed = errors.New("failed to soft-delete (trash) transfer")

// ErrRestoreTransferFailed indicates a failure when restoring a previously trashed transfer.
var ErrRestoreTransferFailed = errors.New("failed to restore transfer")

// ErrDeleteTransferPermanentFailed indicates a failure when permanently deleting a transfer.
var ErrDeleteTransferPermanentFailed = errors.New("failed to permanently delete transfer")

// ErrRestoreAllTransfersFailed indicates a failure when restoring all trashed transfers.
var ErrRestoreAllTransfersFailed = errors.New("failed to restore all transfers")

// ErrDeleteAllTransfersPermanentFailed indicates a failure when permanently deleting all transfers.
var ErrDeleteAllTransfersPermanentFailed = errors.New("failed to permanently delete all transfers")
