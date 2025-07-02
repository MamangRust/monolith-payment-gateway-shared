package responseservice

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// merchantDocumentResponseMapper provides methods to map MerchantDocumentRecord domain models to MerchantDocumentResponse API-compatible response types.
type merchantDocumentResponseMapper struct{}

// NewMerchantDocumentResponseMapper creates a new instance of merchantDocumentResponseMapper which provides methods to map MerchantDocumentRecord domain models to MerchantDocumentResponse API-compatible response types.
func NewMerchantDocumentResponseMapper() *merchantDocumentResponseMapper {
	return &merchantDocumentResponseMapper{}
}

// ToMerchantDocumentResponse maps a single MerchantDocumentRecord to a MerchantDocumentResponse API-compatible response.
// Args:
//   - doc: A pointer to a MerchantDocumentRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantDocumentResponse containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (s *merchantDocumentResponseMapper) ToMerchantDocumentResponse(doc *record.MerchantDocumentRecord) *response.MerchantDocumentResponse {
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

// ToMerchantDocumentsResponse maps multiple MerchantDocumentRecords to a slice of MerchantDocumentResponse API-compatible responses.
// It constructs a slice of MerchantDocumentResponse by mapping each MerchantDocumentRecord to a MerchantDocumentResponse using
// the ToMerchantDocumentResponse method.
//
// Args:
//   - docs: A slice of pointers to MerchantDocumentRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantDocumentResponse containing the mapped data.
func (s *merchantDocumentResponseMapper) ToMerchantDocumentsResponse(docs []*record.MerchantDocumentRecord) []*response.MerchantDocumentResponse {
	var responses []*response.MerchantDocumentResponse
	for _, doc := range docs {
		responses = append(responses, s.ToMerchantDocumentResponse(doc))
	}
	return responses
}

// ToMerchantDocumentResponseDeleteAt maps a soft-deleted MerchantDocumentRecord to its corresponding response.
// Args:
//   - doc: A pointer to a MerchantDocumentRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a MerchantDocumentResponseDeleteAt containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (s *merchantDocumentResponseMapper) ToMerchantDocumentResponseDeleteAt(doc *record.MerchantDocumentRecord) *response.MerchantDocumentResponseDeleteAt {
	return &response.MerchantDocumentResponseDeleteAt{
		ID:           doc.ID,
		MerchantID:   doc.MerchantID,
		DocumentType: doc.DocumentType,
		DocumentURL:  doc.DocumentURL,
		Status:       doc.Status,
		Note:         doc.Note,
		CreatedAt:    doc.CreatedAt,
		UpdatedAt:    doc.UpdatedAt,
		DeletedAt:    doc.DeletedAt,
	}
}

// ToMerchantDocumentsResponseDeleteAt maps multiple soft-deleted MerchantDocumentRecords to a slice of MerchantDocumentResponseDeleteAt API-compatible responses.
// It constructs a slice of MerchantDocumentResponseDeleteAt by mapping each MerchantDocumentRecord to a MerchantDocumentResponseDeleteAt using
// the ToMerchantDocumentResponseDeleteAt method.
//
// Args:
//   - docs: A slice of pointers to MerchantDocumentRecord containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to MerchantDocumentResponseDeleteAt containing the mapped data.
func (s *merchantDocumentResponseMapper) ToMerchantDocumentsResponseDeleteAt(docs []*record.MerchantDocumentRecord) []*response.MerchantDocumentResponseDeleteAt {
	var responses []*response.MerchantDocumentResponseDeleteAt
	for _, doc := range docs {
		responses = append(responses, s.ToMerchantDocumentResponseDeleteAt(doc))
	}
	return responses
}
