package apimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// saldoResponse provides methods to map gRPC saldo responses to HTTP API responses.
type saldoResponse struct {
}

// NewSaldoResponseMapper creates a new instance of saldoResponseMapper.
func NewSaldoResponseMapper() *saldoResponse {
	return &saldoResponse{}
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
func (s *saldoResponse) ToApiResponseSaldo(pbResponse *pb.ApiResponseSaldo) *response.ApiResponseSaldo {
	return &response.ApiResponseSaldo{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    *s.mapResponseSaldo(pbResponse.Data),
	}
}

// ToApiResponsesSaldo maps a gRPC response containing multiple saldo records to an HTTP API response. It constructs an ApiResponseSaldo by copying the status and message fields and mapping the saldo data slice to a slice of SaldoResponse.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponsesSaldo containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsesSaldo with mapped data.
func (s *saldoResponse) ToApiResponsesSaldo(pbResponse *pb.ApiResponsesSaldo) *response.ApiResponsesSaldo {
	return &response.ApiResponsesSaldo{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapResponsesSaldo(pbResponse.Data),
	}
}

// ToApiResponseSaldoDelete maps a gRPC response containing a deleted saldo record to an HTTP API response.
// It constructs an ApiResponseSaldoDelete by copying the status and message fields from the gRPC response.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseSaldoDelete containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseSaldoDelete with mapped data.
func (s *saldoResponse) ToApiResponseSaldoDelete(pbResponse *pb.ApiResponseSaldoDelete) *response.ApiResponseSaldoDelete {
	return &response.ApiResponseSaldoDelete{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseSaldoAll maps a gRPC response containing all saldo records to an HTTP API response. It constructs an ApiResponseSaldoAll by copying the status and message fields from the gRPC response.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseSaldoAll containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseSaldoAll with mapped data.
func (s *saldoResponse) ToApiResponseSaldoAll(pbResponse *pb.ApiResponseSaldoAll) *response.ApiResponseSaldoAll {
	return &response.ApiResponseSaldoAll{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
	}
}

// ToApiResponseMonthTotalSaldo maps a gRPC response containing monthly total saldo
// statistics to an HTTP API response. It constructs an ApiResponseMonthTotalSaldo by
// copying the status and message fields and mapping the saldo statistics to a
// SaldoMonthTotalBalanceResponse slice.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseMonthTotalSaldo containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMonthTotalSaldo with mapped data.
func (s *saldoResponse) ToApiResponseMonthTotalSaldo(pbResponse *pb.ApiResponseMonthTotalSaldo) *response.ApiResponseMonthTotalSaldo {
	return &response.ApiResponseMonthTotalSaldo{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapSaldoMonthTotalBalanceResponses(pbResponse.Data),
	}
}

// ToApiResponseYearTotalSaldo maps a gRPC response containing yearly total saldo
// statistics to an HTTP API response. It constructs an ApiResponseYearTotalSaldo by
// copying the status and message fields and mapping the saldo statistics to a
// SaldoYearTotalBalanceResponse slice.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseYearTotalSaldo containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseYearTotalSaldo with mapped data.
func (s *saldoResponse) ToApiResponseYearTotalSaldo(pbResponse *pb.ApiResponseYearTotalSaldo) *response.ApiResponseYearTotalSaldo {
	return &response.ApiResponseYearTotalSaldo{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapSaldoYearTotalBalanceResponses(pbResponse.Data),
	}
}

// ToApiResponseMonthSaldoBalances maps a gRPC response containing monthly saldo
// balance statistics to an HTTP API response. It constructs an
// ApiResponseMonthSaldoBalances by copying the status and message fields and mapping
// the saldo statistics to a SaldoMonthBalanceResponse slice.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseMonthSaldoBalances containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseMonthSaldoBalances with mapped data.
func (s *saldoResponse) ToApiResponseMonthSaldoBalances(pbResponse *pb.ApiResponseMonthSaldoBalances) *response.ApiResponseMonthSaldoBalances {
	return &response.ApiResponseMonthSaldoBalances{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapSaldoMonthBalanceResponses(pbResponse.Data),
	}
}

// ToApiResponseYearSaldoBalances maps a gRPC response containing yearly saldo
// balance statistics to an HTTP API response. It constructs an
// ApiResponseYearSaldoBalances by copying the status and message fields and
// mapping the saldo statistics to a SaldoYearBalanceResponse slice.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponseYearSaldoBalances containing the
//	gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponseYearSaldoBalances with mapped data.
func (s *saldoResponse) ToApiResponseYearSaldoBalances(pbResponse *pb.ApiResponseYearSaldoBalances) *response.ApiResponseYearSaldoBalances {
	return &response.ApiResponseYearSaldoBalances{
		Status:  pbResponse.Status,
		Message: pbResponse.Message,
		Data:    s.mapSaldoYearBalanceResponses(pbResponse.Data),
	}
}

// ToApiResponsePaginationSaldo maps a paginated gRPC response of saldo records to an HTTP API response.
// It constructs an ApiResponsePaginationSaldo by copying the status and message fields,
// mapping the saldo data slice to a slice of SaldoResponse, and including pagination metadata.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponsePaginationSaldo containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationSaldo with mapped data and pagination info.
func (s *saldoResponse) ToApiResponsePaginationSaldo(pbResponse *pb.ApiResponsePaginationSaldo) *response.ApiResponsePaginationSaldo {
	return &response.ApiResponsePaginationSaldo{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesSaldo(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
	}
}

// ToApiResponsePaginationSaldoDeleteAt maps a paginated gRPC response of soft-deleted saldo records
// to an HTTP API response. It constructs an ApiResponsePaginationSaldoDeleteAt by copying the
// status and message fields, mapping the saldo data slice to a slice of SaldoResponseDeleteAt,
// and including pagination metadata.
//
// Args:
//
//	pbResponse: A pointer to a pb.ApiResponsePaginationSaldoDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.ApiResponsePaginationSaldoDeleteAt with mapped data and pagination info.
func (s *saldoResponse) ToApiResponsePaginationSaldoDeleteAt(pbResponse *pb.ApiResponsePaginationSaldoDeleteAt) *response.ApiResponsePaginationSaldoDeleteAt {
	return &response.ApiResponsePaginationSaldoDeleteAt{
		Status:     pbResponse.Status,
		Message:    pbResponse.Message,
		Data:       s.mapResponsesSaldoDeleteAt(pbResponse.Data),
		Pagination: mapPaginationMeta(pbResponse.Pagination),
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
func (s *saldoResponse) mapResponseSaldo(saldo *pb.SaldoResponse) *response.SaldoResponse {
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

// mapResponsesSaldo maps a slice of gRPC SaldoResponses to a slice of HTTP API SaldoResponses.
// It constructs a slice of SaldoResponse by converting and copying relevant fields
// such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
// CreatedAt, and UpdatedAt from each gRPC SaldoResponse.
//
// Args:
//
//	saldos: A slice of pointers to pb.SaldoResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoResponse with mapped data.
func (s *saldoResponse) mapResponsesSaldo(saldos []*pb.SaldoResponse) []*response.SaldoResponse {
	var responseSaldos []*response.SaldoResponse

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.mapResponseSaldo(saldo))
	}

	return responseSaldos
}

// mapResponseSaldoDeleteAt maps a gRPC SaldoResponseDeleteAt to an HTTP API SaldoResponseDeleteAt.
// It constructs a SaldoResponseDeleteAt by converting and copying relevant fields
// such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
// CreatedAt, UpdatedAt, and DeletedAt.
//
// Args:
//
//	saldo: A pointer to a pb.SaldoResponseDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoResponseDeleteAt with mapped data.
func (s *saldoResponse) mapResponseSaldoDeleteAt(saldo *pb.SaldoResponseDeleteAt) *response.SaldoResponseDeleteAt {
	var deletedAt string
	if saldo.DeletedAt != nil {
		deletedAt = saldo.DeletedAt.Value
	}

	return &response.SaldoResponseDeleteAt{
		ID:             int(saldo.SaldoId),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
		DeletedAt:      &deletedAt,
	}
}

// mapResponsesSaldoDeleteAt maps a slice of gRPC SaldoResponseDeleteAt to a slice of HTTP API SaldoResponseDeleteAt.
// It constructs a slice of SaldoResponseDeleteAt by converting and copying relevant fields
// such as ID, CardNumber, TotalBalance, WithdrawTime, WithdrawAmount,
// CreatedAt, UpdatedAt, and DeletedAt from each gRPC SaldoResponseDeleteAt.
//
// Args:
//
//	saldos: A slice of pointers to pb.SaldoResponseDeleteAt containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoResponseDeleteAt with mapped data.
func (s *saldoResponse) mapResponsesSaldoDeleteAt(saldos []*pb.SaldoResponseDeleteAt) []*response.SaldoResponseDeleteAt {
	var responseSaldos []*response.SaldoResponseDeleteAt

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.mapResponseSaldoDeleteAt(saldo))
	}

	return responseSaldos
}

// mapSaldoMonthTotalBalanceResponse maps a gRPC SaldoMonthTotalBalanceResponse to an HTTP API SaldoMonthTotalBalanceResponse.
// It constructs a SaldoMonthTotalBalanceResponse by copying the Month and Year fields and mapping the TotalBalance field to an integer.
//
// Args:
//
//	ss: A pointer to a pb.SaldoMonthTotalBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoMonthTotalBalanceResponse with mapped data.
func (s *saldoResponse) mapSaldoMonthTotalBalanceResponse(ss *pb.SaldoMonthTotalBalanceResponse) *response.SaldoMonthTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = int(ss.TotalBalance)
	}

	return &response.SaldoMonthTotalBalanceResponse{
		Month:        ss.Month,
		Year:         ss.Year,
		TotalBalance: totalBalance,
	}
}

// mapSaldoMonthTotalBalanceResponses maps a slice of gRPC SaldoMonthTotalBalanceResponses to a slice of HTTP API SaldoMonthTotalBalanceResponses.
// It constructs a slice of SaldoMonthTotalBalanceResponse by copying the Month and Year fields and mapping the TotalBalance field to an integer from each gRPC SaldoMonthTotalBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to pb.SaldoMonthTotalBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoMonthTotalBalanceResponse with mapped data.
func (s *saldoResponse) mapSaldoMonthTotalBalanceResponses(ss []*pb.SaldoMonthTotalBalanceResponse) []*response.SaldoMonthTotalBalanceResponse {
	var saldoProtos []*response.SaldoMonthTotalBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoMonthTotalBalanceResponse(saldo))
	}
	return saldoProtos
}

// mapSaldoYearTotalBalanceResponse maps a gRPC SaldoYearTotalBalanceResponse to an HTTP API SaldoYearTotalBalanceResponse.
// It constructs a SaldoYearTotalBalanceResponse by converting and copying the Year
// field and mapping the TotalBalance field to an integer.
//
// Args:
//
//	ss: A pointer to a pb.SaldoYearTotalBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoYearTotalBalanceResponse with mapped data.
func (s *saldoResponse) mapSaldoYearTotalBalanceResponse(ss *pb.SaldoYearTotalBalanceResponse) *response.SaldoYearTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = int(ss.TotalBalance)
	}

	return &response.SaldoYearTotalBalanceResponse{
		Year:         ss.Year,
		TotalBalance: totalBalance,
	}
}

// mapSaldoYearTotalBalanceResponses maps a slice of gRPC SaldoYearTotalBalanceResponses to a slice of HTTP API SaldoYearTotalBalanceResponses.
// It constructs a slice of SaldoYearTotalBalanceResponse by copying the Year field and mapping the TotalBalance field to an integer from each gRPC SaldoYearTotalBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to pb.SaldoYearTotalBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoYearTotalBalanceResponse with mapped data.
func (s *saldoResponse) mapSaldoYearTotalBalanceResponses(ss []*pb.SaldoYearTotalBalanceResponse) []*response.SaldoYearTotalBalanceResponse {
	var saldoProtos []*response.SaldoYearTotalBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoYearTotalBalanceResponse(saldo))
	}
	return saldoProtos
}

// mapSaldoMonthBalanceResponse maps a gRPC SaldoMonthBalanceResponse to an HTTP API SaldoMonthBalanceResponse.
// It constructs a SaldoMonthBalanceResponse by copying the Month field and mapping the TotalBalance field to an integer.
//
// Args:
//
//	ss: A pointer to a pb.SaldoMonthBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoMonthBalanceResponse with mapped data.
func (s *saldoResponse) mapSaldoMonthBalanceResponse(ss *pb.SaldoMonthBalanceResponse) *response.SaldoMonthBalanceResponse {
	return &response.SaldoMonthBalanceResponse{
		Month:        ss.Month,
		TotalBalance: int(ss.TotalBalance),
	}
}

// mapSaldoMonthBalanceResponses maps a slice of gRPC SaldoMonthBalanceResponses to a slice of HTTP API SaldoMonthBalanceResponses.
// It constructs a slice of SaldoMonthBalanceResponse by copying the Month field and mapping the TotalBalance field to an integer from each gRPC SaldoMonthBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to pb.SaldoMonthBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoMonthBalanceResponse with mapped data.
func (s *saldoResponse) mapSaldoMonthBalanceResponses(ss []*pb.SaldoMonthBalanceResponse) []*response.SaldoMonthBalanceResponse {
	var saldoProtos []*response.SaldoMonthBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoMonthBalanceResponse(saldo))
	}
	return saldoProtos
}

// mapSaldoYearBalanceResponse maps a gRPC SaldoYearBalanceResponse to an HTTP API SaldoYearBalanceResponse.
// It constructs a SaldoYearBalanceResponse by copying the Year field and mapping the TotalBalance field to an integer.
//
// Args:
//
//	ss: A pointer to a pb.SaldoYearBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A pointer to a response.SaldoYearBalanceResponse with mapped data.
func (s *saldoResponse) mapSaldoYearBalanceResponse(ss *pb.SaldoYearBalanceResponse) *response.SaldoYearBalanceResponse {
	return &response.SaldoYearBalanceResponse{
		Year:         ss.Year,
		TotalBalance: int(ss.TotalBalance),
	}
}

// mapSaldoYearBalanceResponses maps a slice of gRPC SaldoYearBalanceResponses to a slice of HTTP API SaldoYearBalanceResponses.
// It constructs a slice of SaldoYearBalanceResponse by copying the Year field and mapping the TotalBalance field to an integer from each gRPC SaldoYearBalanceResponse.
//
// Args:
//
//	ss: A slice of pointers to pb.SaldoYearBalanceResponse containing the gRPC response data.
//
// Returns:
//
//	A slice of pointers to response.SaldoYearBalanceResponse with mapped data.
func (s *saldoResponse) mapSaldoYearBalanceResponses(ss []*pb.SaldoYearBalanceResponse) []*response.SaldoYearBalanceResponse {
	var saldoProtos []*response.SaldoYearBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoYearBalanceResponse(saldo))
	}
	return saldoProtos
}
