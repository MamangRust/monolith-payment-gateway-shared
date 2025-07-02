package merchant_errors

import "errors"

// ErrFindAllMerchantsFailed is returned when fetching all merchants fails.
var ErrFindAllMerchantsFailed = errors.New("failed to find all merchants")

// ErrFindActiveMerchantsFailed is returned when fetching active merchants fails.
var ErrFindActiveMerchantsFailed = errors.New("failed to find active merchants")

// ErrFindTrashedMerchantsFailed is returned when fetching trashed merchants fails.
var ErrFindTrashedMerchantsFailed = errors.New("failed to find trashed merchants")

// ErrFindMerchantByIdFailed is returned when a merchant cannot be found by ID.
var ErrFindMerchantByIdFailed = errors.New("failed to find merchant by ID")

// ErrFindMerchantByApiKeyFailed is returned when a merchant cannot be found by API key.
var ErrFindMerchantByApiKeyFailed = errors.New("failed to find merchant by API key")

// ErrFindMerchantByNameFailed is returned when a merchant cannot be found by name.
var ErrFindMerchantByNameFailed = errors.New("failed to find merchant by name")

// ErrFindMerchantByUserIdFailed is returned when a merchant cannot be found by user ID.
var ErrFindMerchantByUserIdFailed = errors.New("failed to find merchant by user ID")

// ErrGetMonthlyTotalAmountMerchantFailed indicates failure fetching monthly total amount for a merchant.
var ErrGetMonthlyTotalAmountMerchantFailed = errors.New("failed to get monthly total amount of merchant")

// ErrGetYearlyTotalAmountMerchantFailed indicates failure fetching yearly total amount for a merchant.
var ErrGetYearlyTotalAmountMerchantFailed = errors.New("failed to get yearly total amount of merchant")

// ErrGetMonthlyTotalAmountByMerchantsFailed indicates failure fetching monthly total amount for all merchants.
var ErrGetMonthlyTotalAmountByMerchantsFailed = errors.New("failed to get monthly total amount by merchants")

// ErrGetYearlyTotalAmountByMerchantsFailed indicates failure fetching yearly total amount for all merchants.
var ErrGetYearlyTotalAmountByMerchantsFailed = errors.New("failed to get yearly total amount by merchants")

// ErrGetMonthlyTotalAmountByApikeyFailed indicates failure fetching monthly total amount by API key.
var ErrGetMonthlyTotalAmountByApikeyFailed = errors.New("failed to get monthly total amount by API key")

// ErrGetYearlyTotalAmountByApikeyFailed indicates failure fetching yearly total amount by API key.
var ErrGetYearlyTotalAmountByApikeyFailed = errors.New("failed to get yearly total amount by API key")

// ErrGetMonthlyAmountMerchantFailed indicates failure fetching monthly amount for a merchant.
var ErrGetMonthlyAmountMerchantFailed = errors.New("failed to get monthly amount of merchant")

// ErrGetYearlyAmountMerchantFailed indicates failure fetching yearly amount for a merchant.
var ErrGetYearlyAmountMerchantFailed = errors.New("failed to get yearly amount of merchant")

// ErrGetMonthlyAmountByMerchantsFailed indicates failure fetching monthly amount for all merchants.
var ErrGetMonthlyAmountByMerchantsFailed = errors.New("failed to get monthly amount by merchants")

// ErrGetYearlyAmountByMerchantsFailed indicates failure fetching yearly amount for all merchants.
var ErrGetYearlyAmountByMerchantsFailed = errors.New("failed to get yearly amount by merchants")

// ErrGetMonthlyAmountByApikeyFailed indicates failure fetching monthly amount by API key.
var ErrGetMonthlyAmountByApikeyFailed = errors.New("failed to get monthly amount by API key")

// ErrGetYearlyAmountByApikeyFailed indicates failure fetching yearly amount by API key.
var ErrGetYearlyAmountByApikeyFailed = errors.New("failed to get yearly amount by API key")

// ErrGetMonthlyPaymentMethodsMerchantFailed indicates failure fetching monthly payment methods for a merchant.
var ErrGetMonthlyPaymentMethodsMerchantFailed = errors.New("failed to get monthly payment methods of merchant")

// ErrGetYearlyPaymentMethodMerchantFailed indicates failure fetching yearly payment methods for a merchant.
var ErrGetYearlyPaymentMethodMerchantFailed = errors.New("failed to get yearly payment method of merchant")

// ErrGetMonthlyPaymentMethodByMerchantsFailed indicates failure fetching monthly payment methods for all merchants.
var ErrGetMonthlyPaymentMethodByMerchantsFailed = errors.New("failed to get monthly payment method by merchants")

// ErrGetYearlyPaymentMethodByMerchantsFailed indicates failure fetching yearly payment methods for all merchants.
var ErrGetYearlyPaymentMethodByMerchantsFailed = errors.New("failed to get yearly payment method by merchants")

// ErrGetMonthlyPaymentMethodByApikeyFailed indicates failure fetching monthly payment methods by API key.
var ErrGetMonthlyPaymentMethodByApikeyFailed = errors.New("failed to get monthly payment method by API key")

// ErrGetYearlyPaymentMethodByApikeyFailed indicates failure fetching yearly payment methods by API key.
var ErrGetYearlyPaymentMethodByApikeyFailed = errors.New("failed to get yearly payment method by API key")

// ErrFindAllTransactionsFailed indicates failure fetching all transactions.
var ErrFindAllTransactionsFailed = errors.New("failed to find all merchant transactions")

// ErrFindAllTransactionsByMerchantFailed indicates failure fetching transactions by merchant ID.
var ErrFindAllTransactionsByMerchantFailed = errors.New("failed to find merchant transactions by merchant ID")

// ErrFindAllTransactionsByApiKeyFailed indicates failure fetching transactions by API key.
var ErrFindAllTransactionsByApiKeyFailed = errors.New("failed to find merchant transactions by API key")

// ErrCreateMerchantFailed indicates failure when creating a merchant.
var ErrCreateMerchantFailed = errors.New("failed to create merchant")

// ErrUpdateMerchantFailed indicates failure when updating a merchant.
var ErrUpdateMerchantFailed = errors.New("failed to update merchant")

// ErrUpdateMerchantStatusFailed indicates failure when updating merchant status.
var ErrUpdateMerchantStatusFailed = errors.New("failed to update merchant status")

// ErrTrashedMerchantFailed indicates failure when soft-deleting a merchant.
var ErrTrashedMerchantFailed = errors.New("failed to soft-delete (trash) merchant")

// ErrRestoreMerchantFailed indicates failure when restoring a trashed merchant.
var ErrRestoreMerchantFailed = errors.New("failed to restore merchant")

// ErrDeleteMerchantPermanentFailed indicates failure when permanently deleting a merchant.
var ErrDeleteMerchantPermanentFailed = errors.New("failed to permanently delete merchant")

// ErrRestoreAllMerchantFailed indicates failure when restoring all trashed merchants.
var ErrRestoreAllMerchantFailed = errors.New("failed to restore all merchants")

// ErrDeleteAllMerchantPermanentFailed indicates failure when permanently deleting all merchants.
var ErrDeleteAllMerchantPermanentFailed = errors.New("failed to permanently delete all merchants")
