package transferapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/transfer"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type transferCommandResponseMapper struct{}

func NewTransferCommandResponseMapper() TransferCommandResponseMapper {
	return &transferCommandResponseMapper{}
}

// ToApiResponseTransfer maps a single gRPC transfer response to an HTTP API response.
//
// Args:
//   - pbResponse: A pointer to a pb.ApiResponseTransfer containing the gRPC response data.
//
// Returns:
//   - A pointer to a response.ApiResponseTransfer containing the mapped data, including status, message, and a single mapped transfer response.
func (m *transferCommandResponseMapper) ToApiResponseTransfer(pbResponse *pb.ApiResponseTransfer) *response.ApiResponseTransfer {
	return &response.ApiResponseTransfer{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    m.mapResponseTransfer(pbResponse.Data),
	}
}

// mapResponseTransfer maps a TransferResponse from a protobuf message to a response.TransferResponse.
//
// Args:
//   - transfer: A pointer to a pb.TransferResponse containing the data to be mapped.
//
// Returns:
//   - A pointer to a response.TransferResponse containing the mapped data.
func (t *transferCommandResponseMapper) mapResponseTransfer(transfer *pb.TransferResponse) *response.TransferResponse {
	return &response.TransferResponse{
		ID:             int(transfer.Id),
		TransferNo:     transfer.TransferNo,
		TransferFrom:   transfer.TransferFrom,
		TransferTo:     transfer.TransferTo,
		TransferAmount: int(transfer.TransferAmount),
		TransferTime:   transfer.TransferTime,
		CreatedAt:      transfer.CreatedAt,
		UpdatedAt:      transfer.UpdatedAt,
	}
}
