package merchantdocument_errors

import (
	"net/http"

	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"github.com/labstack/echo/v4"
)

var ErrApiInvalidMerchantDocumentID = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Invalid merchant document ID", http.StatusBadRequest)
}

var ErrApiFailedFindAllMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all merchant documents", http.StatusInternalServerError)
}

var ErrApiFailedFindByIdMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch merchant document by ID", http.StatusInternalServerError)
}

var ErrApiFailedFindAllActiveMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all active merchant documents", http.StatusInternalServerError)
}

var ErrApiFailedFindAllTrashedMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to fetch all trashed merchant documents", http.StatusInternalServerError)
}

var ErrApiFailedCreateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to create merchant document", http.StatusInternalServerError)
}

var ErrApiFailedUpdateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant document", http.StatusInternalServerError)
}

var ErrApiFailedUpdateMerchantDocumentStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to update merchant document status", http.StatusInternalServerError)
}

var ErrApiFailedTrashMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to trash merchant document", http.StatusInternalServerError)
}

var ErrApiFailedRestoreMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore merchant document", http.StatusInternalServerError)
}

var ErrApiFailedDeleteMerchantDocumentPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete merchant document", http.StatusInternalServerError)
}

var ErrApiFailedRestoreAllMerchantDocuments = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to restore all merchant documents", http.StatusInternalServerError)
}

var ErrApiFailedDeleteAllMerchantDocumentsPermanent = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "Failed to permanently delete all merchant documents", http.StatusInternalServerError)
}

var ErrApiValidateCreateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid create merchant document request", http.StatusBadRequest)
}

var ErrApiValidateUpdateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant document request", http.StatusBadRequest)
}

var ErrApiValidateUpdateMerchantDocumentStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "validation failed: invalid update merchant document status request", http.StatusBadRequest)
}

// ErrApiBindCreateMerchantDocument is returned when binding the create request payload fails.
var ErrApiBindCreateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid create merchant document request", http.StatusBadRequest)
}

// ErrApiBindUpdateMerchantDocument is returned when binding the update request payload fails.
var ErrApiBindUpdateMerchantDocument = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant document request", http.StatusBadRequest)
}

// ErrApiBindUpdateMerchantDocumentStatus is returned when binding the update status request payload fails.
var ErrApiBindUpdateMerchantDocumentStatus = func(c echo.Context) error {
	return response.NewApiErrorResponse(c, "error", "bind failed: invalid update merchant document status request", http.StatusBadRequest)
}
