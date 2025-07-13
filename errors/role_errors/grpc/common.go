package rolegrpcerrors

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"google.golang.org/grpc/codes"
)

var ErrGrpcRoleInvalidId = response.NewGrpcError("error", "Invalid Role ID", int(codes.NotFound))
