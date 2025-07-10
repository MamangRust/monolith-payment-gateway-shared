package merchantdocumentapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchantdocument"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type merchantDocumentCommandResponseMapepr struct{}

// NewMerchantDocumentCommandResponseMapper returns a new instance of
// merchantDocumentCommandResponseMapepr which provides methods to map
// gRPC merchant document responses to HTTP API responses for command operations.
func NewMerchantDocumentCommandResponseMapper() MerchantDocumentCommandResponseMapper {
	return &merchantDocumentCommandResponseMapepr{}
}

// ToApiResponseMerchantDocument maps a gRPC merchant document response to an HTTP API response.
// It constructs an ApiResponseMerchantDocument by copying the status and message fields
// and mapping the document data to a MerchantDocumentResponse.
//
// Args:
//
//	doc: A pointer to a pb.ApiResponseMerchantDocument containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMerchantDocument with mapped data.
func (m *merchantDocumentCommandResponseMapepr) ToApiResponseMerchantDocument(doc *pb.ApiResponseMerchantDocument) *response.ApiResponseMerchantDocument {
	return &response.ApiResponseMerchantDocument{
		Status:  doc.Status,
		Message: doc.Message,
		Data:    m.mapMerchantDocument(doc.Data),
	}
}

// mapMerchantDocument maps a gRPC merchant document response to an HTTP API response.
// It constructs a MerchantDocumentResponse by copying the fields from the gRPC response.
//
// Args:
//
//	doc: A pointer to a pb.MerchantDocument containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.MerchantDocumentResponse with the mapped data.
func (m *merchantDocumentCommandResponseMapepr) mapMerchantDocument(doc *pb.MerchantDocument) *response.MerchantDocumentResponse {
	return &response.MerchantDocumentResponse{
		ID:           int(doc.DocumentId),
		MerchantID:   int(doc.MerchantId),
		DocumentType: doc.DocumentType,
		DocumentURL:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         doc.Note,
		CreatedAt:    doc.UploadedAt,
		UpdatedAt:    doc.UpdatedAt,
	}
}
