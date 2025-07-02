package merchantdocument_errors

import "errors"

// ErrFindAllMerchantDocumentsFailed is returned when failing to retrieve all merchant documents.
var ErrFindAllMerchantDocumentsFailed = errors.New("failed to find all merchant documents")

// ErrFindActiveMerchantDocumentsFailed is returned when failing to retrieve active merchant documents.
var ErrFindActiveMerchantDocumentsFailed = errors.New("failed to find active merchant documents")

// ErrFindTrashedMerchantDocumentsFailed is returned when failing to retrieve trashed merchant documents.
var ErrFindTrashedMerchantDocumentsFailed = errors.New("failed to find trashed merchant documents")

// ErrFindMerchantDocumentByIdFailed is returned when failing to retrieve a merchant document by ID.
var ErrFindMerchantDocumentByIdFailed = errors.New("failed to find merchant document by ID")

// ErrCreateMerchantDocumentFailed is returned when failing to create a new merchant document.
var ErrCreateMerchantDocumentFailed = errors.New("failed to create merchant document")

// ErrUpdateMerchantDocumentFailed is returned when failing to update an existing merchant document.
var ErrUpdateMerchantDocumentFailed = errors.New("failed to update merchant document")

// ErrUpdateMerchantDocumentStatusFailed is returned when failing to update the status of a merchant document.
var ErrUpdateMerchantDocumentStatusFailed = errors.New("failed to update merchant document status")

// ErrTrashedMerchantDocumentFailed is returned when failing to move a merchant document to trash.
var ErrTrashedMerchantDocumentFailed = errors.New("failed to soft-delete (trash) merchant document")

// ErrRestoreMerchantDocumentFailed is returned when failing to restore a trashed merchant document.
var ErrRestoreMerchantDocumentFailed = errors.New("failed to restore merchant document")

// ErrDeleteMerchantDocumentPermanentFailed is returned when failing to permanently delete a merchant document.
var ErrDeleteMerchantDocumentPermanentFailed = errors.New("failed to permanently delete merchant document")

// ErrRestoreAllMerchantDocumentsFailed is returned when failing to restore all trashed merchant documents.
var ErrRestoreAllMerchantDocumentsFailed = errors.New("failed to restore all merchant documents")

// ErrDeleteAllMerchantDocumentsPermanentFailed is returned when failing to permanently delete all merchant documents.
var ErrDeleteAllMerchantDocumentsPermanentFailed = errors.New("failed to permanently delete all merchant documents")
