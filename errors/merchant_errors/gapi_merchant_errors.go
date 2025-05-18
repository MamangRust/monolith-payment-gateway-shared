package merchant_errors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var (
	ErrGrpcMerchantNotFound      = response.NewGrpcError("merchant", "Merchant not found", int(codes.NotFound))
	ErrGrpcMerchantInvalidID     = response.NewGrpcError("merchant", "Invalid Merchant ID", int(codes.InvalidArgument))
	ErrGrpcMerchantInvalidUserID = response.NewGrpcError("merchant", "Invalid Merchant User ID", int(codes.InvalidArgument))
	ErrGrpcMerchantInvalidApiKey = response.NewGrpcError("merchant", "Invalid Merchant Api Key", int(codes.InvalidArgument))
	ErrGrpcMerchantInvalidMonth  = response.NewGrpcError("month", "Invalid Merchant Month", int(codes.InvalidArgument))
	ErrGrpcMerchantInvalidYear   = response.NewGrpcError("year", "Invalid Merchant Year", int(codes.InvalidArgument))

	ErrGrpcFailedCreateMerchant         = response.NewGrpcError("merchant", "Failed to create merchant", int(codes.Internal))
	ErrGrpcFailedUpdateMerchant         = response.NewGrpcError("merchant", "Failed to update merchant", int(codes.Internal))
	ErrGrpcValidateCreateMerchant       = response.NewGrpcError("merchant", "Invalid input for create merchant", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchant       = response.NewGrpcError("merchant", "Invalid input for update merchant", int(codes.InvalidArgument))
	ErrGrpcValidateUpdateMerchantStatus = response.NewGrpcError("merchant", "Invalid input for update merchant status", int(codes.InvalidArgument))
)
