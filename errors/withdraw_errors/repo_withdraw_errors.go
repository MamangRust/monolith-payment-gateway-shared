package withdraw_errors

import "errors"

// ErrFindAllWithdrawsFailed is used when the system fails to find all withdraws
var ErrFindAllWithdrawsFailed = errors.New("failed to find all withdraws")

// ErrFindActiveWithdrawsFailed is used when the system fails to find active withdraws
var ErrFindActiveWithdrawsFailed = errors.New("failed to find active withdraws")

// ErrFindTrashedWithdrawsFailed is used when the system fails to find trashed withdraws
var ErrFindTrashedWithdrawsFailed = errors.New("failed to find trashed withdraws")

// ErrFindWithdrawsByCardNumberFailed is used when the system fails to find withdraws by card number
var ErrFindWithdrawsByCardNumberFailed = errors.New("failed to find withdraws by card number")

// ErrFindWithdrawByIdFailed is used when the system fails to find a withdraw by ID
var ErrFindWithdrawByIdFailed = errors.New("failed to find withdraw by ID")

// ErrGetMonthWithdrawStatusSuccessFailed is used when the system fails to get the monthly withdraw status success
var ErrGetMonthWithdrawStatusSuccessFailed = errors.New("failed to get monthly withdraw status success")

// ErrGetYearlyWithdrawStatusSuccessFailed is used when the system fails to get the yearly withdraw status success
var ErrGetYearlyWithdrawStatusSuccessFailed = errors.New("failed to get yearly withdraw status success")

// ErrGetMonthWithdrawStatusSuccessByCardFailed is used when the system fails to get the monthly withdraw status success by card number
var ErrGetMonthWithdrawStatusSuccessByCardFailed = errors.New("failed to get monthly withdraw status success by card number")

// ErrGetYearlyWithdrawStatusSuccessByCardFailed is used when the system fails to get the yearly withdraw status success by card number
var ErrGetYearlyWithdrawStatusSuccessByCardFailed = errors.New("failed to get yearly withdraw status success by card number")

// ErrGetMonthWithdrawStatusFailedFailed is used when the system fails to get the monthly withdraw status failed
var ErrGetMonthWithdrawStatusFailedFailed = errors.New("failed to get monthly withdraw status failed")

// ErrGetYearlyWithdrawStatusFailedFailed is used when the system fails to get the yearly withdraw status failed
var ErrGetYearlyWithdrawStatusFailedFailed = errors.New("failed to get yearly withdraw status failed")

// ErrGetMonthWithdrawStatusFailedByCardFailed is used when the system fails to get the monthly withdraw status failed by card number
var ErrGetMonthWithdrawStatusFailedByCardFailed = errors.New("failed to get monthly withdraw status failed by card number")

// ErrGetYearlyWithdrawStatusFailedByCardFailed is used when the system fails to get the yearly withdraw status failed by card number
var ErrGetYearlyWithdrawStatusFailedByCardFailed = errors.New("failed to get yearly withdraw status failed by card number")

// ErrGetMonthlyWithdrawsFailed is used when the system fails to get monthly withdraw amounts
var ErrGetMonthlyWithdrawsFailed = errors.New("failed to get monthly withdraw amounts")

// ErrGetYearlyWithdrawsFailed is used when the system fails to get yearly withdraw amounts
var ErrGetYearlyWithdrawsFailed = errors.New("failed to get yearly withdraw amounts")

// ErrGetMonthlyWithdrawsByCardFailed indicates a failure when retrieving monthly withdraw amounts by card number.
var ErrGetMonthlyWithdrawsByCardFailed = errors.New("failed to get monthly withdraw amounts by card number")

// ErrGetYearlyWithdrawsByCardFailed indicates a failure when retrieving yearly withdraw amounts by card number.
var ErrGetYearlyWithdrawsByCardFailed = errors.New("failed to get yearly withdraw amounts by card number")

// ErrCreateWithdrawFailed indicates a failure when creating a new withdraw record.
var ErrCreateWithdrawFailed = errors.New("failed to create withdraw")

// ErrUpdateWithdrawFailed indicates a failure when updating a withdraw record.
var ErrUpdateWithdrawFailed = errors.New("failed to update withdraw")

// ErrUpdateWithdrawStatusFailed indicates a failure when updating the status of a withdraw record.
var ErrUpdateWithdrawStatusFailed = errors.New("failed to update withdraw status")

// ErrTrashedWithdrawFailed indicates a failure when soft-deleting (trashing) a withdraw record.
var ErrTrashedWithdrawFailed = errors.New("failed to soft-delete (trash) withdraw")

// ErrRestoreWithdrawFailed indicates a failure when restoring a previously trashed withdraw record.
var ErrRestoreWithdrawFailed = errors.New("failed to restore withdraw")

// ErrDeleteWithdrawPermanentFailed indicates a failure when permanently deleting a withdraw record.
var ErrDeleteWithdrawPermanentFailed = errors.New("failed to permanently delete withdraw")

// ErrRestoreAllWithdrawsFailed indicates a failure when restoring all trashed withdraw records.
var ErrRestoreAllWithdrawsFailed = errors.New("failed to restore all withdraws")

// ErrDeleteAllWithdrawsPermanentFailed indicates a failure when permanently deleting all withdraw records.
var ErrDeleteAllWithdrawsPermanentFailed = errors.New("failed to permanently delete all withdraws")
