package errors

import (
	"encoding/json"

	"github.com/MamangRust/monolith-payment-gateway-shared/pb"
)

func GrpcErrorToJson(err *pb.ErrorResponse) string {
	jsonData, _ := json.Marshal(err)
	return string(jsonData)
}
