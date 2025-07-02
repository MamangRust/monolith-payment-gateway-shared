package responseservice

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// saldoResponseMapper provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types.
type saldoResponseMapper struct {
}

// NewSaldoResponseMapper returns a new instance of saldoResponseMapper which provides methods to map SaldoRecord domain models to SaldoResponse API-compatible response types.
func NewSaldoResponseMapper() *saldoResponseMapper {
	return &saldoResponseMapper{}
}

// ToSaldoResponse maps a single SaldoRecord to a SaldoResponse.
//
// Args:
//   - saldo: A pointer to a SaldoRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoResponse containing the mapped data.
func (s *saldoResponseMapper) ToSaldoResponse(saldo *record.SaldoRecord) *response.SaldoResponse {
	return &response.SaldoResponse{
		ID:             saldo.ID,
		CardNumber:     saldo.CardNumber,
		TotalBalance:   saldo.TotalBalance,
		WithdrawAmount: saldo.WithdrawAmount,
		WithdrawTime:   saldo.WithdrawTime,
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}
}

// ToSaldoResponses maps a list of SaldoRecord into a list of SaldoResponse.
//
// Args:
//
//	saldos: A slice of pointers to SaldoRecord containing the data to be mapped.
//
// Returns:
//
//	A slice of pointers to SaldoResponse containing the mapped data.
func (s *saldoResponseMapper) ToSaldoResponses(saldos []*record.SaldoRecord) []*response.SaldoResponse {
	var responses []*response.SaldoResponse

	for _, response := range saldos {
		responses = append(responses, s.ToSaldoResponse(response))
	}

	return responses
}

// ToSaldoResponseDeleteAt maps a single SaldoRecord to a SaldoResponseDeleteAt.
//
// Args:
//   - saldo: A pointer to a SaldoRecord containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoResponseDeleteAt containing the mapped data.
func (s *saldoResponseMapper) ToSaldoResponseDeleteAt(saldo *record.SaldoRecord) *response.SaldoResponseDeleteAt {
	return &response.SaldoResponseDeleteAt{
		ID:             saldo.ID,
		CardNumber:     saldo.CardNumber,
		TotalBalance:   saldo.TotalBalance,
		WithdrawAmount: saldo.WithdrawAmount,
		WithdrawTime:   saldo.WithdrawTime,
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
		DeletedAt:      saldo.DeletedAt,
	}
}

// ToSaldoResponsesDeleteAt maps a list of SaldoRecord into a list of SaldoResponseDeleteAt.
//
// Args:
//
//	saldos: A slice of pointers to SaldoRecord containing the data to be mapped.
//
// Returns:
//
//	A slice of pointers to SaldoResponseDeleteAt containing the mapped data.
func (s *saldoResponseMapper) ToSaldoResponsesDeleteAt(saldos []*record.SaldoRecord) []*response.SaldoResponseDeleteAt {
	var responses []*response.SaldoResponseDeleteAt

	for _, response := range saldos {
		responses = append(responses, s.ToSaldoResponseDeleteAt(response))
	}

	return responses
}

// ToSaldoMonthTotalBalanceResponse maps a SaldoMonthTotalBalance record to a SaldoMonthTotalBalanceResponse.
//
// Args:
//   - ss: A pointer to a SaldoMonthTotalBalance containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoMonthTotalBalanceResponse containing the mapped data, including Month, Year, and TotalBalance.
func (s *saldoResponseMapper) ToSaldoMonthTotalBalanceResponse(ss *record.SaldoMonthTotalBalance) *response.SaldoMonthTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = ss.TotalBalance
	}

	return &response.SaldoMonthTotalBalanceResponse{
		Month:        ss.Month,
		Year:         ss.Year,
		TotalBalance: totalBalance,
	}
}

// ToSaldoMonthTotalBalanceResponses maps a list of SaldoMonthTotalBalance records to a list of SaldoMonthTotalBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to SaldoMonthTotalBalance containing the data to be mapped.
//
// Returns:
//
//	A slice of pointers to SaldoMonthTotalBalanceResponse containing the mapped data, including Month, Year, and TotalBalance.
func (s *saldoResponseMapper) ToSaldoMonthTotalBalanceResponses(ss []*record.SaldoMonthTotalBalance) []*response.SaldoMonthTotalBalanceResponse {
	var saldoResponses []*response.SaldoMonthTotalBalanceResponse
	for _, saldo := range ss {
		saldoResponses = append(saldoResponses, s.ToSaldoMonthTotalBalanceResponse(saldo))
	}
	return saldoResponses
}

// ToSaldoYearTotalBalanceResponse maps a SaldoYearTotalBalance record to a SaldoYearTotalBalanceResponse.
//
// Args:
//   - ss: A pointer to a SaldoYearTotalBalance containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoYearTotalBalanceResponse containing the mapped data, including Year, and TotalBalance.
func (s *saldoResponseMapper) ToSaldoYearTotalBalanceResponse(ss *record.SaldoYearTotalBalance) *response.SaldoYearTotalBalanceResponse {
	return &response.SaldoYearTotalBalanceResponse{
		Year:         ss.Year,
		TotalBalance: ss.TotalBalance,
	}
}

// ToSaldoYearTotalBalanceResponses maps a list of SaldoYearTotalBalance records to a list of SaldoYearTotalBalanceResponse.
//
// Args:
//   - ss: A slice of pointers to SaldoYearTotalBalance containing the data to be mapped.
//
// Returns:
//   - A slice of pointers to SaldoYearTotalBalanceResponse containing the mapped data, including Year, and TotalBalance.
func (s *saldoResponseMapper) ToSaldoYearTotalBalanceResponses(ss []*record.SaldoYearTotalBalance) []*response.SaldoYearTotalBalanceResponse {
	var saldoResponses []*response.SaldoYearTotalBalanceResponse
	for _, saldo := range ss {
		saldoResponses = append(saldoResponses, s.ToSaldoYearTotalBalanceResponse(saldo))
	}
	return saldoResponses
}

// ToSaldoMonthBalanceResponse maps a SaldoMonthSaldoBalance record to a SaldoMonthBalanceResponse.
//
// Args:
//   - ss: A pointer to a SaldoMonthSaldoBalance containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoMonthBalanceResponse containing the mapped data, including Month, and TotalBalance.
func (s *saldoResponseMapper) ToSaldoMonthBalanceResponse(ss *record.SaldoMonthSaldoBalance) *response.SaldoMonthBalanceResponse {
	return &response.SaldoMonthBalanceResponse{
		Month:        ss.Month,
		TotalBalance: ss.TotalBalance,
	}
}

// ToSaldoMonthBalanceResponses maps a list of SaldoMonthSaldoBalance records to a list of SaldoMonthBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to SaldoMonthSaldoBalance containing the data to be mapped.
//
// Returns:
//
//	A slice of pointers to SaldoMonthBalanceResponse containing the mapped data, including Month, and TotalBalance.
func (s *saldoResponseMapper) ToSaldoMonthBalanceResponses(ss []*record.SaldoMonthSaldoBalance) []*response.SaldoMonthBalanceResponse {
	var saldoResponses []*response.SaldoMonthBalanceResponse
	for _, saldo := range ss {
		saldoResponses = append(saldoResponses, s.ToSaldoMonthBalanceResponse(saldo))
	}
	return saldoResponses
}

// ToSaldoYearBalanceResponse maps a SaldoYearSaldoBalance record to a SaldoYearBalanceResponse.
//
// Args:
//   - ss: A pointer to a SaldoYearSaldoBalance containing the data to be mapped.
//
// Returns:
//   - A pointer to a SaldoYearBalanceResponse containing the mapped data, including Year, and TotalBalance.
func (s *saldoResponseMapper) ToSaldoYearBalanceResponse(ss *record.SaldoYearSaldoBalance) *response.SaldoYearBalanceResponse {
	return &response.SaldoYearBalanceResponse{
		Year:         ss.Year,
		TotalBalance: ss.TotalBalance,
	}
}

// ToSaldoYearBalanceResponses maps a list of SaldoYearSaldoBalance records to a list of SaldoYearBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to SaldoYearSaldoBalance containing the data to be mapped.
//
// Returns:
//
//	A slice of pointers to SaldoYearBalanceResponse containing the mapped data, including Year, and TotalBalance.
func (s *saldoResponseMapper) ToSaldoYearBalanceResponses(ss []*record.SaldoYearSaldoBalance) []*response.SaldoYearBalanceResponse {
	var saldoResponses []*response.SaldoYearBalanceResponse
	for _, saldo := range ss {
		saldoResponses = append(saldoResponses, s.ToSaldoYearBalanceResponse(saldo))
	}
	return saldoResponses
}
