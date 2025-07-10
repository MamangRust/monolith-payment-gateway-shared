package merchantdocumentservicemapper

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantDocumentCommandResponseMapper struct {
}

func NewMerchantDocumentCommandResponseMapper() MerchantDocumentCommandResponseMapper {
	return &merchantDocumentCommandResponseMapper{}
}

// ToMerchantDocumentResponse maps a single MerchantDocumentRecord to a MerchantDocumentResponse API-compatible response.
// Parameters:
//   - doc: A pointer to a MerchantDocumentRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantDocumentResponse containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (s *merchantDocumentCommandResponseMapper) ToMerchantDocumentResponse(doc *record.MerchantDocumentRecord) *response.MerchantDocumentResponse {
	return &response.MerchantDocumentResponse{
		ID:           doc.ID,
		MerchantID:   doc.MerchantID,
		DocumentType: doc.DocumentType,
		DocumentURL:  doc.DocumentURL,
		Status:       doc.Status,
		Note:         doc.Note,
		CreatedAt:    doc.CreatedAt,
		UpdatedAt:    doc.UpdatedAt,
	}
}
