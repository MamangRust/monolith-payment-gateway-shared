package refreshtoken_errors

import (
	"github.com/MamangRust/payment-gateway-monolith-grpc/shared/domain/response"
	"google.golang.org/grpc/codes"
)

var ErrGrpcRefreshToken = response.NewGrpcError("error", "refresh token failed", int(codes.Unauthenticated))
