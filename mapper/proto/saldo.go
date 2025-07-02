package protomapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type saldoProtoMapper struct {
}

// NewSaldoProtoMapper initializes and returns a new instance of saldoProtoMapper.
func NewSaldoProtoMapper() *saldoProtoMapper {
	return &saldoProtoMapper{}
}

// ToProtoResponseSaldo maps a SaldoResponse to an ApiResponseSaldo proto message.
//
// Args:
//
//	status: The status of the response.
//	message: The message accompanying the response.
//	pbResponse: The SaldoResponse to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseSaldo containing the mapped data.
func (s *saldoProtoMapper) ToProtoResponseSaldo(status string, message string, pbResponse *response.SaldoResponse) *pb.ApiResponseSaldo {
	return &pb.ApiResponseSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapResponseSaldo(pbResponse),
	}
}

// ToProtoResponsesSaldo maps a list of SaldoResponse to an ApiResponsesSaldo proto message.
//
// Args:
//
//	status: The status of the response.
//	message: The message accompanying the response.
//	pbResponse: A slice of SaldoResponse to be converted.
//
// Returns:
//
//	A pointer to an ApiResponsesSaldo containing the mapped data.
func (s *saldoProtoMapper) ToProtoResponsesSaldo(status string, message string, pbResponse []*response.SaldoResponse) *pb.ApiResponsesSaldo {
	return &pb.ApiResponsesSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapResponsesSaldo(pbResponse),
	}
}

// ToProtoResponseSaldoDelete maps a status and message to an ApiResponseSaldoDelete proto response.
//
// It is used to generate the response for the SaldoService.DeleteSaldo rpc method.
func (s *saldoProtoMapper) ToProtoResponseSaldoDelete(status string, message string) *pb.ApiResponseSaldoDelete {
	return &pb.ApiResponseSaldoDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseSaldoAll maps a status and message to an ApiResponseSaldoAll proto response.
//
// Args:
//
//	status: The status of the response.
//	message: The message accompanying the response.
//
// Returns:
//
//	A pointer to an ApiResponseSaldoAll containing the mapped data.
func (s *saldoProtoMapper) ToProtoResponseSaldoAll(status string, message string) *pb.ApiResponseSaldoAll {
	return &pb.ApiResponseSaldoAll{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseMonthTotalSaldo maps a status, message and a list of *response.SaldoMonthTotalBalanceResponse
// to a *pb.ApiResponseMonthTotalSaldo proto response.
//
// It is used to generate the response for the SaldoService.GetMonthlyTotalSaldoBalance rpc method.
func (s *saldoProtoMapper) ToProtoResponseMonthTotalSaldo(status string, message string, pbResponse []*response.SaldoMonthTotalBalanceResponse) *pb.ApiResponseMonthTotalSaldo {
	return &pb.ApiResponseMonthTotalSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoMonthTotalBalanceResponses(pbResponse),
	}
}

// ToProtoResponseYearTotalSaldo maps a status, message and a list of *response.SaldoYearTotalBalanceResponse
// to a *pb.ApiResponseYearTotalSaldo proto response.
//
// It is used to generate the response for the SaldoService.GetYearlyTotalSaldoBalance rpc method.
func (s *saldoProtoMapper) ToProtoResponseYearTotalSaldo(status string, message string, pbResponse []*response.SaldoYearTotalBalanceResponse) *pb.ApiResponseYearTotalSaldo {
	return &pb.ApiResponseYearTotalSaldo{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoYearTotalBalanceResponses(pbResponse),
	}
}

// ToProtoResponseMonthSaldoBalances maps a status, message and a list of *response.SaldoMonthBalanceResponse
// to a *pb.ApiResponseMonthSaldoBalances proto response.
//
// It is used to generate the response for the SaldoService.GetMonthlySaldoBalances rpc method.
func (s *saldoProtoMapper) ToProtoResponseMonthSaldoBalances(status string, message string, pbResponse []*response.SaldoMonthBalanceResponse) *pb.ApiResponseMonthSaldoBalances {
	return &pb.ApiResponseMonthSaldoBalances{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoMonthBalanceResponses(pbResponse),
	}
}

// ToProtoResponseYearSaldoBalances maps a status, message, and a list of
// *response.SaldoYearBalanceResponse to a *pb.ApiResponseYearSaldoBalances proto response.
//
// It is used to generate the response for the SaldoService.GetYearlySaldoBalances rpc method.
//
// Args:
//
//	status: The status of the response.
//	message: The message accompanying the response.
//	pbResponse: A slice of SaldoYearBalanceResponse to be converted.
//
// Returns:
//
//	A pointer to an ApiResponseYearSaldoBalances containing the mapped data.
func (s *saldoProtoMapper) ToProtoResponseYearSaldoBalances(status string, message string, pbResponse []*response.SaldoYearBalanceResponse) *pb.ApiResponseYearSaldoBalances {
	return &pb.ApiResponseYearSaldoBalances{
		Status:  status,
		Message: message,
		Data:    s.mapSaldoYearBalanceResponses(pbResponse),
	}
}

// ToProtoResponsePaginationSaldo maps a status, message, pagination meta, and a list of *response.SaldoResponse
// to a *pb.ApiResponsePaginationSaldo proto response.
//
// It is used to generate the response for the SaldoService.ListSaldo rpc method.
//
// Args:
//
//	pagination: The pagination data for the paginated response.
//	status: The status of the response.
//	message: The message accompanying the response.
//	pbResponse: A slice of SaldoResponse to be converted.
//
// Returns:
//
//	A pointer to an ApiResponsePaginationSaldo containing the status, message, data, and pagination data.
func (s *saldoProtoMapper) ToProtoResponsePaginationSaldo(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponse) *pb.ApiResponsePaginationSaldo {
	return &pb.ApiResponsePaginationSaldo{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesSaldo(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponsePaginationSaldoDeleteAt creates a new instance of ApiResponsePaginationSaldoDeleteAt.
//
// It maps the provided pagination metadata, status, message, and a list of SaldoResponseDeleteAt
// to its protobuf representation.
//
// Args:
//
//	pagination: The pagination data for the paginated response.
//	status: The status of the API response.
//	message: A descriptive message of the API response.
//	pbResponse: The list of SaldoResponseDeleteAt that needs to be converted.
//
// Returns:
//
//	A pointer to ApiResponsePaginationSaldoDeleteAt containing the status, message, saldo data, and pagination data.
func (s *saldoProtoMapper) ToProtoResponsePaginationSaldoDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponseDeleteAt) *pb.ApiResponsePaginationSaldoDeleteAt {
	return &pb.ApiResponsePaginationSaldoDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapResponsesSaldoDeleteAt(pbResponse),
		Pagination: mapPaginationMeta(pagination),
	}
}

// mapResponseSaldo maps a SaldoResponse to its protobuf representation.
//
// Args:
//
//	saldo: The SaldoResponse to be converted.
//
// Returns:
//
//	A pointer to SaldoResponse containing the mapped data.
func (s *saldoProtoMapper) mapResponseSaldo(saldo *response.SaldoResponse) *pb.SaldoResponse {
	return &pb.SaldoResponse{
		SaldoId:        int32(saldo.ID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int32(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
	}
}

// mapResponsesSaldo maps a list of SaldoResponse to a list of *pb.SaldoResponse proto responses.
//
// It iterates over each SaldoResponse in the input slice, converting
// them to their protobuf equivalent using the mapResponseSaldo function.
// This function is used to generate the response data for ListSaldo rpc methods.
func (s *saldoProtoMapper) mapResponsesSaldo(saldos []*response.SaldoResponse) []*pb.SaldoResponse {
	var responseSaldos []*pb.SaldoResponse

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.mapResponseSaldo(saldo))
	}

	return responseSaldos
}

// mapResponseSaldoDeleteAt maps a single response.SaldoResponseDeleteAt to a single *pb.SaldoResponseDeleteAt proto response.
//
// Args:
//
//	saldo: The SaldoResponseDeleteAt to be converted.
//
// Returns:
//
//	A pointer to a SaldoResponseDeleteAt containing the mapped data.
func (s *saldoProtoMapper) mapResponseSaldoDeleteAt(saldo *response.SaldoResponseDeleteAt) *pb.SaldoResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if saldo.DeletedAt != nil {
		deletedAt = wrapperspb.String(*saldo.DeletedAt)
	}

	return &pb.SaldoResponseDeleteAt{
		SaldoId:        int32(saldo.ID),
		CardNumber:     saldo.CardNumber,
		TotalBalance:   int32(saldo.TotalBalance),
		WithdrawTime:   saldo.WithdrawTime,
		WithdrawAmount: int32(saldo.WithdrawAmount),
		CreatedAt:      saldo.CreatedAt,
		UpdatedAt:      saldo.UpdatedAt,
		DeletedAt:      deletedAt,
	}
}

// mapResponsesSaldoDeleteAt maps a list of *response.SaldoResponseDeleteAt to a list of
// *pb.SaldoResponseDeleteAt proto responses.
//
// It iterates over each SaldoResponseDeleteAt in the input slice, converting
// them to their protobuf equivalent using the mapResponseSaldoDeleteAt function.
// This function is used to generate the response data for Saldo-related RPC methods
// involving deleted saldo records.
func (s *saldoProtoMapper) mapResponsesSaldoDeleteAt(saldos []*response.SaldoResponseDeleteAt) []*pb.SaldoResponseDeleteAt {
	var responseSaldos []*pb.SaldoResponseDeleteAt

	for _, saldo := range saldos {
		responseSaldos = append(responseSaldos, s.mapResponseSaldoDeleteAt(saldo))
	}

	return responseSaldos
}

// mapSaldoMonthTotalBalanceResponse maps a *response.SaldoMonthTotalBalanceResponse to a *pb.SaldoMonthTotalBalanceResponse.
//
// It maps the fields of the input response to the corresponding fields of the output response.
// The TotalBalance field is converted to int32.
//
// Args:
//
//	ss: The SaldoMonthTotalBalanceResponse to be converted.
//
// Returns:
//
//	A pointer to a SaldoMonthTotalBalanceResponse containing the mapped data.
func (s *saldoProtoMapper) mapSaldoMonthTotalBalanceResponse(ss *response.SaldoMonthTotalBalanceResponse) *pb.SaldoMonthTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = ss.TotalBalance
	}

	return &pb.SaldoMonthTotalBalanceResponse{
		Month:        ss.Month,
		Year:         ss.Year,
		TotalBalance: int32(totalBalance),
	}
}

// mapSaldoMonthTotalBalanceResponses maps a list of *response.SaldoMonthTotalBalanceResponse
// to a list of *pb.SaldoMonthTotalBalanceResponse proto responses.
//
// It iterates over each SaldoMonthTotalBalanceResponse in the input slice, converting
// them to their protobuf equivalent using the mapSaldoMonthTotalBalanceResponse function.
// This function is used to generate the response data for monthly total saldo balance RPC methods.
//
// Args:
//
//	ss: A slice of SaldoMonthTotalBalanceResponse to be converted.
//
// Returns:
//
//	A slice of SaldoMonthTotalBalanceResponse containing the mapped data.
func (s *saldoProtoMapper) mapSaldoMonthTotalBalanceResponses(ss []*response.SaldoMonthTotalBalanceResponse) []*pb.SaldoMonthTotalBalanceResponse {
	var saldoProtos []*pb.SaldoMonthTotalBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoMonthTotalBalanceResponse(saldo))
	}
	return saldoProtos
}

// mapSaldoYearTotalBalanceResponse maps a *response.SaldoYearTotalBalanceResponse to a *pb.SaldoYearTotalBalanceResponse proto response.
//
// It takes a SaldoYearTotalBalanceResponse from the response domain model and converts it
// into its protobuf representation, ensuring that the year and total balance fields are accurately mapped.
//
// Args:
//
//	ss: The SaldoYearTotalBalanceResponse to be converted.
//
// Returns:
//
//	A pointer to a SaldoYearTotalBalanceResponse containing the mapped data.
func (s *saldoProtoMapper) mapSaldoYearTotalBalanceResponse(ss *response.SaldoYearTotalBalanceResponse) *pb.SaldoYearTotalBalanceResponse {
	totalBalance := 0

	if ss.TotalBalance != 0 {
		totalBalance = ss.TotalBalance
	}

	return &pb.SaldoYearTotalBalanceResponse{
		Year:         ss.Year,
		TotalBalance: int32(totalBalance),
	}
}

// mapSaldoYearTotalBalanceResponses maps a list of *response.SaldoYearTotalBalanceResponse
// to a list of *pb.SaldoYearTotalBalanceResponse proto responses.
//
// It iterates over each SaldoYearTotalBalanceResponse in the input slice, converting
// them to their protobuf equivalent using the mapSaldoYearTotalBalanceResponse function.
// This function is used to generate the response data for yearly total saldo balance RPC methods.
func (s *saldoProtoMapper) mapSaldoYearTotalBalanceResponses(ss []*response.SaldoYearTotalBalanceResponse) []*pb.SaldoYearTotalBalanceResponse {
	var saldoProtos []*pb.SaldoYearTotalBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoYearTotalBalanceResponse(saldo))
	}
	return saldoProtos
}

// mapSaldoMonthBalanceResponse maps a *response.SaldoMonthBalanceResponse to a *pb.SaldoMonthBalanceResponse proto response.
//
// It takes a SaldoMonthBalanceResponse from the response domain model and converts it
// into its protobuf representation, ensuring that the month and total balance fields are accurately mapped.
//
// Args:
//
//	ss: The SaldoMonthBalanceResponse to be converted.
//
// Returns:
//
//	A pointer to a SaldoMonthBalanceResponse containing the mapped data.
func (s *saldoProtoMapper) mapSaldoMonthBalanceResponse(ss *response.SaldoMonthBalanceResponse) *pb.SaldoMonthBalanceResponse {
	return &pb.SaldoMonthBalanceResponse{
		Month:        ss.Month,
		TotalBalance: int32(ss.TotalBalance),
	}
}

// mapSaldoMonthBalanceResponses maps a list of *response.SaldoMonthBalanceResponse
// to a list of *pb.SaldoMonthBalanceResponse proto responses.
//
// It iterates over each SaldoMonthBalanceResponse in the input slice, converting
// them to their protobuf equivalent using the mapSaldoMonthBalanceResponse function.
// This function is used to generate the response data for monthly balance RPC methods.
func (s *saldoProtoMapper) mapSaldoMonthBalanceResponses(ss []*response.SaldoMonthBalanceResponse) []*pb.SaldoMonthBalanceResponse {
	var saldoProtos []*pb.SaldoMonthBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoMonthBalanceResponse(saldo))
	}
	return saldoProtos
}

// mapSaldoYearBalanceResponse maps a *response.SaldoYearBalanceResponse to a *pb.SaldoYearBalanceResponse proto response.
//
// It takes a SaldoYearBalanceResponse from the response domain model and converts it
// into its protobuf representation, ensuring that the year and total balance fields are accurately mapped.
//
// Args:
//
//	ss: The SaldoYearBalanceResponse to be converted.
//
// Returns:
//
//	A pointer to SaldoYearBalanceResponse containing the mapped data.
func (s *saldoProtoMapper) mapSaldoYearBalanceResponse(ss *response.SaldoYearBalanceResponse) *pb.SaldoYearBalanceResponse {
	return &pb.SaldoYearBalanceResponse{
		Year:         ss.Year,
		TotalBalance: int32(ss.TotalBalance),
	}
}

// mapSaldoYearBalanceResponses maps a list of *response.SaldoYearBalanceResponse
// to a list of *pb.SaldoYearBalanceResponse proto responses.
//
// It iterates over each SaldoYearBalanceResponse in the input slice, converting
// them to their protobuf equivalent using the mapSaldoYearBalanceResponse function.
// This function is used to generate the response data for yearly balance RPC methods.
func (s *saldoProtoMapper) mapSaldoYearBalanceResponses(ss []*response.SaldoYearBalanceResponse) []*pb.SaldoYearBalanceResponse {
	var saldoProtos []*pb.SaldoYearBalanceResponse
	for _, saldo := range ss {
		saldoProtos = append(saldoProtos, s.mapSaldoYearBalanceResponse(saldo))
	}
	return saldoProtos
}
