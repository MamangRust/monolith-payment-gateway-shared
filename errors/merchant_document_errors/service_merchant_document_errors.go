package merchantdocument_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// ErrMerchantDocumentNotFoundRes is returned when the requested merchant document cannot be found.
var ErrMerchantDocumentNotFoundRes = response.NewErrorResponse("Merchant Document not found", http.StatusNotFound)

// ErrFailedFindAllMerchantDocuments is returned when failing to fetch all merchant documents.
var ErrFailedFindAllMerchantDocuments = response.NewErrorResponse("Failed to fetch Merchant Documents", http.StatusInternalServerError)

// ErrFailedFindActiveMerchantDocuments is returned when failing to fetch active merchant documents.
var ErrFailedFindActiveMerchantDocuments = response.NewErrorResponse("Failed to fetch active Merchant Documents", http.StatusInternalServerError)

// ErrFailedFindTrashedMerchantDocuments is returned when failing to fetch trashed merchant documents.
var ErrFailedFindTrashedMerchantDocuments = response.NewErrorResponse("Failed to fetch trashed Merchant Documents", http.StatusInternalServerError)

// ErrFailedFindMerchantDocumentById is returned when failing to find a merchant document by ID.
var ErrFailedFindMerchantDocumentById = response.NewErrorResponse("Failed to find Merchant Document by ID", http.StatusInternalServerError)

// ErrFailedCreateMerchantDocument is returned when failing to create a new merchant document.
var ErrFailedCreateMerchantDocument = response.NewErrorResponse("Failed to create Merchant Document", http.StatusInternalServerError)

// ErrFailedUpdateMerchantDocument is returned when failing to update an existing merchant document.
var ErrFailedUpdateMerchantDocument = response.NewErrorResponse("Failed to update Merchant Document", http.StatusInternalServerError)

// ErrFailedTrashMerchantDocument is returned when failing to move a merchant document to trash.
var ErrFailedTrashMerchantDocument = response.NewErrorResponse("Failed to trash Merchant Document", http.StatusInternalServerError)

// ErrFailedRestoreMerchantDocument is returned when failing to restore a trashed merchant document.
var ErrFailedRestoreMerchantDocument = response.NewErrorResponse("Failed to restore Merchant Document", http.StatusInternalServerError)

// ErrFailedDeleteMerchantDocument is returned when failing to permanently delete a merchant document.
var ErrFailedDeleteMerchantDocument = response.NewErrorResponse("Failed to delete Merchant Document permanently", http.StatusInternalServerError)

// ErrFailedRestoreAllMerchantDocuments is returned when failing to restore all trashed merchant documents.
var ErrFailedRestoreAllMerchantDocuments = response.NewErrorResponse("Failed to restore all Merchant Documents", http.StatusInternalServerError)

// ErrFailedDeleteAllMerchantDocuments is returned when failing to permanently delete all merchant documents.
var ErrFailedDeleteAllMerchantDocuments = response.NewErrorResponse("Failed to delete all Merchant Documents permanently", http.StatusInternalServerError)
