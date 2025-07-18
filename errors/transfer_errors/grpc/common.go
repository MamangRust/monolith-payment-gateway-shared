package transfergrpcerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcTransferNotFound  = response.NewGrpcError("transfer", "Transfer not found", int(codes.NotFound))
	ErrGrpcTransferInvalidID = response.NewGrpcError("transfer", "Invalid Transfer ID", int(codes.InvalidArgument))
	ErrGrpcInvalidUserID     = response.NewGrpcError("card_id", "Invalid user ID", int(codes.InvalidArgument))
	ErrGrpcInvalidCardNumber = response.NewGrpcError("card_id", "Invalid card number", int(codes.InvalidArgument))
	ErrGrpcInvalidMonth      = response.NewGrpcError("month", "Invalid month", int(codes.InvalidArgument))
	ErrGrpcInvalidYear       = response.NewGrpcError("year", "Invalid year", int(codes.InvalidArgument))
)
