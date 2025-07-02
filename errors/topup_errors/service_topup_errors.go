package topup_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrTopupNotFoundRes indicates that a requested top-up was not found.
var ErrTopupNotFoundRes = response.NewErrorResponse("Topup not found", http.StatusNotFound)

// ErrFailedFindAllTopups indicates failure in retrieving all top-up records.
var ErrFailedFindAllTopups = response.NewErrorResponse("Failed to fetch Topups", http.StatusInternalServerError)

// ErrFailedFindAllTopupsByCardNumber indicates failure in retrieving top-ups by card number.
var ErrFailedFindAllTopupsByCardNumber = response.NewErrorResponse("Failed to fetch Topups by card number", http.StatusInternalServerError)

// ErrFailedFindTopupById indicates failure in finding a top-up by its ID.
var ErrFailedFindTopupById = response.NewErrorResponse("Failed to find Topup by ID", http.StatusInternalServerError)

// ErrFailedFindActiveTopups indicates failure in retrieving active (non-trashed) top-up records.
var ErrFailedFindActiveTopups = response.NewErrorResponse("Failed to fetch active Topups", http.StatusInternalServerError)

// ErrFailedFindTrashedTopups indicates failure in retrieving trashed (soft-deleted) top-up records.
var ErrFailedFindTrashedTopups = response.NewErrorResponse("Failed to fetch trashed Topups", http.StatusInternalServerError)

// ErrFailedFindMonthTopupStatusSuccess indicates failure in retrieving monthly successful top-up status.
var ErrFailedFindMonthTopupStatusSuccess = response.NewErrorResponse("Failed to get monthly topup success status", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusSuccess indicates failure in retrieving yearly successful top-up status.
var ErrFailedFindYearlyTopupStatusSuccess = response.NewErrorResponse("Failed to get yearly topup success status", http.StatusInternalServerError)

// ErrFailedFindMonthTopupStatusFailed indicates failure in retrieving monthly failed top-up status.
var ErrFailedFindMonthTopupStatusFailed = response.NewErrorResponse("Failed to get monthly topup failed status", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusFailed indicates failure in retrieving yearly failed top-up status.
var ErrFailedFindYearlyTopupStatusFailed = response.NewErrorResponse("Failed to get yearly topup failed status", http.StatusInternalServerError)

// ErrFailedFindMonthTopupStatusSuccessByCard indicates failure in retrieving monthly successful top-up status by card.
var ErrFailedFindMonthTopupStatusSuccessByCard = response.NewErrorResponse("Failed to get monthly topup success status by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusSuccessByCard indicates failure in retrieving yearly successful top-up status by card.
var ErrFailedFindYearlyTopupStatusSuccessByCard = response.NewErrorResponse("Failed to get yearly topup success status by card", http.StatusInternalServerError)

// ErrFailedFindMonthTopupStatusFailedByCard indicates failure in retrieving monthly failed top-up status by card.
var ErrFailedFindMonthTopupStatusFailedByCard = response.NewErrorResponse("Failed to get monthly topup failed status by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupStatusFailedByCard indicates failure in retrieving yearly failed top-up status by card.
var ErrFailedFindYearlyTopupStatusFailedByCard = response.NewErrorResponse("Failed to get yearly topup failed status by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTopupMethods indicates failure in retrieving monthly top-up methods.
var ErrFailedFindMonthlyTopupMethods = response.NewErrorResponse("Failed to get monthly topup methods", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupMethods indicates failure in retrieving yearly top-up methods.
var ErrFailedFindYearlyTopupMethods = response.NewErrorResponse("Failed to get yearly topup methods", http.StatusInternalServerError)

// ErrFailedFindMonthlyTopupMethodsByCard indicates failure in retrieving monthly top-up methods by card.
var ErrFailedFindMonthlyTopupMethodsByCard = response.NewErrorResponse("Failed to get monthly topup methods by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupMethodsByCard indicates failure in retrieving yearly top-up methods by card.
var ErrFailedFindYearlyTopupMethodsByCard = response.NewErrorResponse("Failed to get yearly topup methods by card", http.StatusInternalServerError)

// ErrFailedFindMonthlyTopupAmounts indicates failure in retrieving monthly top-up amounts.
var ErrFailedFindMonthlyTopupAmounts = response.NewErrorResponse("Failed to get monthly topup amounts", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupAmounts indicates failure in retrieving yearly top-up amounts.
var ErrFailedFindYearlyTopupAmounts = response.NewErrorResponse("Failed to get yearly topup amounts", http.StatusInternalServerError)

// ErrFailedFindMonthlyTopupAmountsByCard indicates failure in retrieving monthly top-up amounts by card.
var ErrFailedFindMonthlyTopupAmountsByCard = response.NewErrorResponse("Failed to get monthly topup amounts by card", http.StatusInternalServerError)

// ErrFailedFindYearlyTopupAmountsByCard indicates failure in retrieving yearly top-up amounts by card.
var ErrFailedFindYearlyTopupAmountsByCard = response.NewErrorResponse("Failed to get yearly topup amounts by card", http.StatusInternalServerError)

// ErrFailedCreateTopup indicates failure in creating a new top-up record.
var ErrFailedCreateTopup = response.NewErrorResponse("Failed to create Topup", http.StatusInternalServerError)

// ErrFailedUpdateTopup indicates failure in updating an existing top-up record.
var ErrFailedUpdateTopup = response.NewErrorResponse("Failed to update Topup", http.StatusInternalServerError)

// ErrFailedTrashTopup indicates failure in soft-deleting (trashing) a top-up.
var ErrFailedTrashTopup = response.NewErrorResponse("Failed to trash Topup", http.StatusInternalServerError)

// ErrFailedRestoreTopup indicates failure in restoring a previously trashed top-up.
var ErrFailedRestoreTopup = response.NewErrorResponse("Failed to restore Topup", http.StatusInternalServerError)

// ErrFailedDeleteTopup indicates failure in permanently deleting a top-up.
var ErrFailedDeleteTopup = response.NewErrorResponse("Failed to delete Topup permanently", http.StatusInternalServerError)

// ErrFailedRestoreAllTopups indicates failure in restoring all trashed top-up records.
var ErrFailedRestoreAllTopups = response.NewErrorResponse("Failed to restore all Topups", http.StatusInternalServerError)

// ErrFailedDeleteAllTopups indicates failure in permanently deleting all trashed top-up records.
var ErrFailedDeleteAllTopups = response.NewErrorResponse("Failed to delete all Topups permanently", http.StatusInternalServerError)
