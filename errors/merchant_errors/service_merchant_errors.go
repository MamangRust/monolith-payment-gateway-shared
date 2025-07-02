package merchant_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrMerchantNotFoundRes is returned when the merchant is not found.
var ErrMerchantNotFoundRes = response.NewErrorResponse("Merchant not found", http.StatusNotFound)

// ErrFailedFindAllMerchants indicates failure in fetching all merchants.
var ErrFailedFindAllMerchants = response.NewErrorResponse("Failed to fetch Merchants", http.StatusInternalServerError)

// ErrFailedFindActiveMerchants indicates failure in fetching active merchants.
var ErrFailedFindActiveMerchants = response.NewErrorResponse("Failed to fetch active Merchants", http.StatusInternalServerError)

// ErrFailedFindTrashedMerchants indicates failure in fetching trashed merchants.
var ErrFailedFindTrashedMerchants = response.NewErrorResponse("Failed to fetch trashed Merchants", http.StatusInternalServerError)

// ErrFailedFindMerchantById is returned when a merchant cannot be found by ID.
var ErrFailedFindMerchantById = response.NewErrorResponse("Failed to find Merchant by ID", http.StatusInternalServerError)

// ErrFailedFindByApiKey is returned when a merchant cannot be found by API key.
var ErrFailedFindByApiKey = response.NewErrorResponse("Failed to find Merchant by API key", http.StatusInternalServerError)

// ErrFailedFindByMerchantUserId is returned when a merchant cannot be found by user ID.
var ErrFailedFindByMerchantUserId = response.NewErrorResponse("Failed to find Merchant by User ID", http.StatusInternalServerError)

// ErrFailedSendEmail indicates a failure when sending an email related to merchant operations.
var ErrFailedSendEmail = response.NewErrorResponse("Failed to send email", http.StatusInternalServerError)

// ErrFailedFindAllTransactions is returned when fetching all transactions fails.
var ErrFailedFindAllTransactions = response.NewErrorResponse("Failed to fetch Merchant transactions", http.StatusInternalServerError)

// ErrFailedFindAllTransactionsByMerchant is returned when fetching transactions by merchant fails.
var ErrFailedFindAllTransactionsByMerchant = response.NewErrorResponse("Failed to fetch transactions by Merchant", http.StatusInternalServerError)

// ErrFailedFindAllTransactionsByApikey is returned when fetching transactions by API key fails.
var ErrFailedFindAllTransactionsByApikey = response.NewErrorResponse("Failed to fetch transactions by API key", http.StatusInternalServerError)

// ErrFailedFindMonthlyPaymentMethodsMerchant indicates failure fetching monthly payment methods.
var ErrFailedFindMonthlyPaymentMethodsMerchant = response.NewErrorResponse("Failed to get monthly payment methods", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodMerchant indicates failure fetching yearly payment methods.
var ErrFailedFindYearlyPaymentMethodMerchant = response.NewErrorResponse("Failed to get yearly payment method", http.StatusInternalServerError)

// ErrFailedFindMonthlyPaymentMethodByMerchants indicates failure fetching monthly payment methods by merchant.
var ErrFailedFindMonthlyPaymentMethodByMerchants = response.NewErrorResponse("Failed to get monthly payment methods by Merchant", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodByMerchants indicates failure fetching yearly payment methods by merchant.
var ErrFailedFindYearlyPaymentMethodByMerchants = response.NewErrorResponse("Failed to get yearly payment method by Merchant", http.StatusInternalServerError)

// ErrFailedFindMonthlyPaymentMethodByApikeys indicates failure fetching monthly payment methods by API key.
var ErrFailedFindMonthlyPaymentMethodByApikeys = response.NewErrorResponse("Failed to get monthly payment methods by API key", http.StatusInternalServerError)

// ErrFailedFindYearlyPaymentMethodByApikeys indicates failure fetching yearly payment methods by API key.
var ErrFailedFindYearlyPaymentMethodByApikeys = response.NewErrorResponse("Failed to get yearly payment method by API key", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmountMerchant indicates failure fetching monthly amounts.
var ErrFailedFindMonthlyAmountMerchant = response.NewErrorResponse("Failed to get monthly amount", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountMerchant indicates failure fetching yearly amounts.
var ErrFailedFindYearlyAmountMerchant = response.NewErrorResponse("Failed to get yearly amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmountByMerchants indicates failure fetching monthly amounts by merchant.
var ErrFailedFindMonthlyAmountByMerchants = response.NewErrorResponse("Failed to get monthly amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountByMerchants indicates failure fetching yearly amounts by merchant.
var ErrFailedFindYearlyAmountByMerchants = response.NewErrorResponse("Failed to get yearly amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindMonthlyAmountByApikeys indicates failure fetching monthly amounts by API key.
var ErrFailedFindMonthlyAmountByApikeys = response.NewErrorResponse("Failed to get monthly amount by API key", http.StatusInternalServerError)

// ErrFailedFindYearlyAmountByApikeys indicates failure fetching yearly amounts by API key.
var ErrFailedFindYearlyAmountByApikeys = response.NewErrorResponse("Failed to get yearly amount by API key", http.StatusInternalServerError)

// ErrFailedFindMonthlyTotalAmountMerchant indicates failure fetching monthly total amounts.
var ErrFailedFindMonthlyTotalAmountMerchant = response.NewErrorResponse("Failed to get monthly total amount", http.StatusInternalServerError)

// ErrFailedFindYearlyTotalAmountMerchant indicates failure fetching yearly total amounts.
var ErrFailedFindYearlyTotalAmountMerchant = response.NewErrorResponse("Failed to get yearly total amount", http.StatusInternalServerError)

// ErrFailedFindMonthlyTotalAmountByMerchants indicates failure fetching monthly total amounts by merchant.
var ErrFailedFindMonthlyTotalAmountByMerchants = response.NewErrorResponse("Failed to get monthly total amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindYearlyTotalAmountByMerchants indicates failure fetching yearly total amounts by merchant.
var ErrFailedFindYearlyTotalAmountByMerchants = response.NewErrorResponse("Failed to get yearly total amount by Merchant", http.StatusInternalServerError)

// ErrFailedFindMonthlyTotalAmountByApikeys indicates failure fetching monthly total amounts by API key.
var ErrFailedFindMonthlyTotalAmountByApikeys = response.NewErrorResponse("Failed to get monthly total amount by API key", http.StatusInternalServerError)

// ErrFailedFindYearlyTotalAmountByApikeys indicates failure fetching yearly total amounts by API key.
var ErrFailedFindYearlyTotalAmountByApikeys = response.NewErrorResponse("Failed to get yearly total amount by API key", http.StatusInternalServerError)

// ErrFailedCreateMerchant indicates failure when creating a new merchant.
var ErrFailedCreateMerchant = response.NewErrorResponse("Failed to create Merchant", http.StatusInternalServerError)

// ErrFailedUpdateMerchant indicates failure when updating a merchant.
var ErrFailedUpdateMerchant = response.NewErrorResponse("Failed to update Merchant", http.StatusInternalServerError)

// ErrFailedTrashMerchant indicates failure when soft-deleting (trashing) a merchant.
var ErrFailedTrashMerchant = response.NewErrorResponse("Failed to trash Merchant", http.StatusInternalServerError)

// ErrFailedRestoreMerchant indicates failure when restoring a trashed merchant.
var ErrFailedRestoreMerchant = response.NewErrorResponse("Failed to restore Merchant", http.StatusInternalServerError)

// ErrFailedDeleteMerchant indicates failure when permanently deleting a merchant.
var ErrFailedDeleteMerchant = response.NewErrorResponse("Failed to delete Merchant permanently", http.StatusInternalServerError)

// ErrFailedRestoreAllMerchants indicates failure when restoring all trashed merchants.
var ErrFailedRestoreAllMerchants = response.NewErrorResponse("Failed to restore all Merchants", http.StatusInternalServerError)

// ErrFailedDeleteAllMerchants indicates failure when permanently deleting all trashed merchants.
var ErrFailedDeleteAllMerchants = response.NewErrorResponse("Failed to delete all Merchants permanently", http.StatusInternalServerError)
