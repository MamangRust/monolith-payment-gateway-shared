package saldoapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/saldo"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type saldoCommandResponseMapper struct {
}

func NewSaldoCommandResponseMapper() SaldoCommandResponseMapper {
	return &saldoCommandResponseMapper{}
}

// ToApiResponseSaldo maps a gRPC saldo response to an HTTP API response. It constructs an ApiResponseSaldo by copying the status and message fields and mapping the saldo data to a SaldoResponse.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseSaldo containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseSaldo with mapped data.
func (s *saldoCommandResponseMapper) ToApiResponseSaldo(pbResponse *pb.ApiResponseSaldo) *response.ApiResponseSaldo {
	return &response.ApiResponseSaldo{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponseSaldo(pbResponse.Data),
	}
}

// mapResponseSaldo maps a gRPC SaldoResponse to an HTTP API SaldoResponse.
// It constructs a SaldoResponse by converting and copying relevant fields
// such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
// CreatedAt, and UpdatedAt.
//
// Args:
//
//	saldo: A pointer to a pb.SaldoResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoResponse with mapped data.
func (s *saldoCommandResponseMapper) mapResponseSaldo(saldo *pb.SaldoResponse) *response.SaldoResponse {
	return &response.SaldoResponse{
		ID:             int(saldo.SaldoId),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}
}
