package merchantdocumentapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/merchantdocument"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type MerchantDocumentBaseResponseMapper interface {
	// Maps a single merchant document to an HTTP API response.
	ToApiResponseMerchantDocument(doc *pb.ApiResponseMerchantDocument) *response.ApiResponseMerchantDocument
}

type MerchantDocumentQueryResponseMapper interface {
	MerchantDocumentBaseResponseMapper

	// Maps paginated merchant document responses to an HTTP API response.
	ToApiResponsePaginationMerchantDocument(docs *pb.ApiResponsePaginationMerchantDocument) *response.ApiResponsePaginationMerchantDocument

	// Maps paginated soft-deleted merchant document responses to an HTTP API response.
	ToApiResponsePaginationMerchantDocumentDeleteAt(docs *pb.ApiResponsePaginationMerchantDocumentAt) *response.ApiResponsePaginationMerchantDocumentDeleteAt

	// Maps all merchant documents to an API-compatible response.
	ToApiResponseMerchantDocumentAll(resp *pb.ApiResponseMerchantDocumentAll) *response.ApiResponseMerchantDocumentAll

	// Maps a soft-deleted single merchant document response to an API response.
	ToApiResponseMerchantDocumentDeleteAt(resp *pb.ApiResponseMerchantDocumentDelete) *response.ApiResponseMerchantDocumentDelete
}

type MerchantDocumentCommandResponseMapper interface {
	MerchantDocumentBaseResponseMapper
}
