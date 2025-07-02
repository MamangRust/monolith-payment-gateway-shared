package topup_errors

import "errors"

// ErrFindAllTopupsFailed indicates failure in retrieving all top-up records from the database.
var ErrFindAllTopupsFailed = errors.New("failed to find all topups")

// ErrFindTopupsByActiveFailed indicates failure in retrieving only the active (non-deleted) top-ups.
var ErrFindTopupsByActiveFailed = errors.New("failed to find active topups")

// ErrFindTopupsByTrashedFailed indicates failure in retrieving trashed (soft-deleted) top-ups.
var ErrFindTopupsByTrashedFailed = errors.New("failed to find trashed topups")

// ErrFindTopupsByCardNumberFailed indicates failure in finding top-ups by a specific card number.
var ErrFindTopupsByCardNumberFailed = errors.New("failed to find topups by card number")

// ErrFindTopupByIdFailed indicates failure in finding a top-up using its unique ID.
var ErrFindTopupByIdFailed = errors.New("failed to find topup by ID")

// ErrGetMonthTopupStatusSuccessFailed indicates failure in getting the monthly count of successful top-ups.
var ErrGetMonthTopupStatusSuccessFailed = errors.New("failed to get monthly topup status success")

// ErrGetYearlyTopupStatusSuccessFailed indicates failure in getting the yearly count of successful top-ups.
var ErrGetYearlyTopupStatusSuccessFailed = errors.New("failed to get yearly topup status success")

// ErrGetMonthTopupStatusSuccessByCardFailed indicates failure in getting monthly successful top-up status by card number.
var ErrGetMonthTopupStatusSuccessByCardFailed = errors.New("failed to get monthly topup status success by card number")

// ErrGetYearlyTopupStatusSuccessByCardFailed indicates failure in getting yearly successful top-up status by card number.
var ErrGetYearlyTopupStatusSuccessByCardFailed = errors.New("failed to get yearly topup status success by card number")

// ErrGetMonthTopupStatusFailedFailed indicates failure in getting the monthly count of failed top-ups.
var ErrGetMonthTopupStatusFailedFailed = errors.New("failed to get monthly topup status failed")

// ErrGetYearlyTopupStatusFailedFailed indicates failure in getting the yearly count of failed top-ups.
var ErrGetYearlyTopupStatusFailedFailed = errors.New("failed to get yearly topup status failed")

// ErrGetMonthTopupStatusFailedByCardFailed indicates failure in getting monthly failed top-up status by card number.
var ErrGetMonthTopupStatusFailedByCardFailed = errors.New("failed to get monthly topup status failed by card number")

// ErrGetYearlyTopupStatusFailedByCardFailed indicates failure in getting yearly failed top-up status by card number.
var ErrGetYearlyTopupStatusFailedByCardFailed = errors.New("failed to get yearly topup status failed by card number")

// ErrGetMonthlyTopupMethodsFailed indicates failure in retrieving monthly top-up payment methods statistics.
var ErrGetMonthlyTopupMethodsFailed = errors.New("failed to get monthly topup methods")

// ErrGetYearlyTopupMethodsFailed indicates failure in retrieving yearly top-up payment methods statistics.
var ErrGetYearlyTopupMethodsFailed = errors.New("failed to get yearly topup methods")

// ErrGetMonthlyTopupAmountsFailed indicates failure in retrieving the monthly top-up amounts.
var ErrGetMonthlyTopupAmountsFailed = errors.New("failed to get monthly topup amounts")

// ErrGetYearlyTopupAmountsFailed indicates failure in retrieving the yearly top-up amounts.
var ErrGetYearlyTopupAmountsFailed = errors.New("failed to get yearly topup amounts")

// ErrGetMonthlyTopupMethodsByCardFailed indicates failure in retrieving monthly payment method stats by card number.
var ErrGetMonthlyTopupMethodsByCardFailed = errors.New("failed to get monthly topup methods by card number")

// ErrGetYearlyTopupMethodsByCardFailed indicates failure in retrieving yearly payment method stats by card number.
var ErrGetYearlyTopupMethodsByCardFailed = errors.New("failed to get yearly topup methods by card number")

// ErrGetMonthlyTopupAmountsByCardFailed indicates failure in retrieving monthly top-up amount stats by card number.
var ErrGetMonthlyTopupAmountsByCardFailed = errors.New("failed to get monthly topup amounts by card number")

// ErrGetYearlyTopupAmountsByCardFailed indicates failure in retrieving yearly top-up amount stats by card number.
var ErrGetYearlyTopupAmountsByCardFailed = errors.New("failed to get yearly topup amounts by card number")

// ErrCreateTopupFailed indicates failure in creating a new top-up record.
var ErrCreateTopupFailed = errors.New("failed to create topup")

// ErrUpdateTopupFailed indicates failure in updating an existing top-up record.
var ErrUpdateTopupFailed = errors.New("failed to update topup")

// ErrUpdateTopupAmountFailed indicates failure in updating only the top-up amount.
var ErrUpdateTopupAmountFailed = errors.New("failed to update topup amount")

// ErrUpdateTopupStatusFailed indicates failure in updating the top-up status (e.g., success/failed).
var ErrUpdateTopupStatusFailed = errors.New("failed to update topup status")

// ErrTrashedTopupFailed indicates failure in soft-deleting (trashing) a top-up.
var ErrTrashedTopupFailed = errors.New("failed to soft-delete (trash) topup")

// ErrRestoreTopupFailed indicates failure in restoring a previously trashed top-up.
var ErrRestoreTopupFailed = errors.New("failed to restore topup")

// ErrDeleteTopupPermanentFailed indicates failure in permanently deleting a top-up.
var ErrDeleteTopupPermanentFailed = errors.New("failed to permanently delete topup")

// ErrRestoreAllTopupFailed indicates failure in restoring all trashed top-ups.
var ErrRestoreAllTopupFailed = errors.New("failed to restore all topups")

// ErrDeleteAllTopupPermanentFailed indicates failure in permanently deleting all trashed top-ups.
var ErrDeleteAllTopupPermanentFailed = errors.New("failed to permanently delete all topups")
