package protomapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type cardProtoMapper struct{}

// NewCardProtoMapper creates a new instance of cardProtoMapper.
func NewCardProtoMapper() *cardProtoMapper {
	return &cardProtoMapper{}
}

// ToProtoResponseCard maps a *response.CardResponse to a *pb.ApiResponseCard proto response.
//
// It is used to generate the response for the CardService.GetCard rpc method.
func (s *cardProtoMapper) ToProtoResponseCard(status string, message string, card *response.CardResponse) *pb.ApiResponseCard {
	return &pb.ApiResponseCard{
		Status:  status,
		Message: message,
		Data:    s.mapCardResponse(card),
	}
}

// ToProtoResponsePaginationCard maps a pagination meta, status, message and a list of *response.CardResponse
// to a *pb.ApiResponsePaginationCard proto response.
//
// It is used to generate the response for the CardService.ListCard rpc method.
func (s *cardProtoMapper) ToProtoResponsePaginationCard(pagination *pb.PaginationMeta, status string, message string, cards []*response.CardResponse) *pb.ApiResponsePaginationCard {
	return &pb.ApiResponsePaginationCard{
		Status:     status,
		Message:    message,
		Data:       s.mapCardResponses(cards),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponseCardDeleteAt maps a status and message to a *pb.ApiResponseCardDelete proto response.
//
// It is used to generate the response for the CardService.DeleteCard rpc method.
func (s *cardProtoMapper) ToProtoResponseCardDeleteAt(status string, message string) *pb.ApiResponseCardDelete {
	return &pb.ApiResponseCardDelete{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponseCardAll maps a status and message to a *pb.ApiResponseCardAll proto response.
//
// It is used to generate the response for the CardService.GetAllCard rpc method.
func (s *cardProtoMapper) ToProtoResponseCardAll(status string, message string) *pb.ApiResponseCardAll {
	return &pb.ApiResponseCardAll{
		Status:  status,
		Message: message,
	}
}

// ToProtoResponsePaginationCardDeletedAt maps a pagination meta, status, message and a list of *response.CardResponseDeleteAt
// to a *pb.ApiResponsePaginationCardDeleteAt proto response.
//
// It is used to generate the response for the CardService.ListDeletedCard rpc method.
func (s *cardProtoMapper) ToProtoResponsePaginationCardDeletedAt(pagination *pb.PaginationMeta, status string, message string, cards []*response.CardResponseDeleteAt) *pb.ApiResponsePaginationCardDeleteAt {
	return &pb.ApiResponsePaginationCardDeleteAt{
		Status:     status,
		Message:    message,
		Data:       s.mapCardResponsesDeleteAt(cards),
		Pagination: mapPaginationMeta(pagination),
	}
}

// ToProtoResponseDashboardCard maps a status, message and a *response.DashboardCard to a *pb.ApiResponseDashboardCard proto response.
//
// It is used to generate the response for the CardService.GetDashboardCard rpc method.
func (s *cardProtoMapper) ToProtoResponseDashboardCard(status string, message string, dash *response.DashboardCard) *pb.ApiResponseDashboardCard {
	return &pb.ApiResponseDashboardCard{
		Status:  status,
		Message: message,
		Data:    s.mapDashboardCard(dash),
	}
}

// ToProtoResponseDashboardCardCardNumber maps a status, message, and a *response.DashboardCardCardNumber
// to a *pb.ApiResponseDashboardCardNumber proto response.
//
// It is used to generate the response for the CardService.GetDashboardCardCardNumber rpc method.
func (s *cardProtoMapper) ToProtoResponseDashboardCardCardNumber(status string, message string, dash *response.DashboardCardCardNumber) *pb.ApiResponseDashboardCardNumber {
	return &pb.ApiResponseDashboardCardNumber{
		Status:  status,
		Message: message,
		Data:    s.mapDashboardCardCardNumber(dash),
	}
}

// ToProtoResponseMonthlyBalances maps a status, message and a list of *response.CardResponseMonthBalance
// to a *pb.ApiResponseMonthlyBalance proto response.
//
// It is used to generate the response for the CardService.GetMonthlyBalance rpc method.
func (s *cardProtoMapper) ToProtoResponseMonthlyBalances(status string, message string, cards []*response.CardResponseMonthBalance) *pb.ApiResponseMonthlyBalance {

	return &pb.ApiResponseMonthlyBalance{
		Status:  status,
		Message: message,
		Data:    s.mapMonthlyBalances(cards),
	}
}

// ToProtoResponseYearlyBalances maps a status, message and a list of *response.CardResponseYearlyBalance
// to a *pb.ApiResponseYearlyBalance proto response.
//
// It is used to generate the response for the CardService.GetYearlyBalance rpc method.
func (s *cardProtoMapper) ToProtoResponseYearlyBalances(status string, message string, cards []*response.CardResponseYearlyBalance) *pb.ApiResponseYearlyBalance {

	return &pb.ApiResponseYearlyBalance{
		Status:  status,
		Message: message,
		Data:    s.mapYearlyBalances(cards),
	}
}

// ToProtoResponseMonthlyAmounts maps a status, message and a list of *response.CardResponseMonthAmount
// to a *pb.ApiResponseMonthlyAmount proto response.
//
// It is used to generate the response for the CardService.GetMonthlyAmount rpc method.
func (s *cardProtoMapper) ToProtoResponseMonthlyAmounts(status string, message string, cards []*response.CardResponseMonthAmount) *pb.ApiResponseMonthlyAmount {

	return &pb.ApiResponseMonthlyAmount{
		Status:  status,
		Message: message,
		Data:    s.mapMonthlyAmounts(cards),
	}
}

// ToProtoResponseYearlyAmounts maps a status, message and a list of *response.CardResponseYearAmount
// to a *pb.ApiResponseYearlyAmount proto response.
//
// It is used to generate the response for the CardService.GetYearlyAmount rpc method.
func (s *cardProtoMapper) ToProtoResponseYearlyAmounts(status string, message string, cards []*response.CardResponseYearAmount) *pb.ApiResponseYearlyAmount {
	return &pb.ApiResponseYearlyAmount{
		Status:  status,
		Message: message,
		Data:    s.mapYearlyAmounts(cards),
	}
}

// mapCardResponse maps a *response.CardResponse to a *pb.CardResponse proto response.
//
// It is used to generate the response for the CardService.GetCard rpc method.
func (s *cardProtoMapper) mapCardResponse(card *response.CardResponse) *pb.CardResponse {
	return &pb.CardResponse{
		Id:           int32(card.ID),
		UserId:       int32(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		Cvv:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}

// mapCardResponses maps a list of *response.CardResponse to a list of *pb.CardResponse
// proto response.
//
// It is used to generate the response for the CardService.ListCard rpc method.
func (s *cardProtoMapper) mapCardResponses(roles []*response.CardResponse) []*pb.CardResponse {
	var responseRoles []*pb.CardResponse

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapCardResponse(role))
	}

	return responseRoles
}

// mapCardResponseDeleteAt maps a *response.CardResponseDeleteAt to a *pb.CardResponseDeleteAt proto response.
//
// This function converts a CardResponseDeleteAt, which includes deletion information, into its protobuf
// representation. It handles the conversion of all fields including a nullable DeletedAt field.
func (s *cardProtoMapper) mapCardResponseDeleteAt(card *response.CardResponseDeleteAt) *pb.CardResponseDeleteAt {
	var deletedAt *wrapperspb.StringValue
	if card.DeletedAt != nil {
		deletedAt = wrapperspb.String(*card.DeletedAt)
	}

	return &pb.CardResponseDeleteAt{
		Id:           int32(card.ID),
		UserId:       int32(card.UserID),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		Cvv:          card.CVV,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
		DeletedAt:    deletedAt,
	}
}

// mapCardResponsesDeleteAt maps a list of *response.CardResponseDeleteAt to a list of
// *pb.CardResponseDeleteAt proto response.
//
// It is used to generate the response for the CardService.ListDeletedCard rpc method.
func (s *cardProtoMapper) mapCardResponsesDeleteAt(roles []*response.CardResponseDeleteAt) []*pb.CardResponseDeleteAt {
	var responseRoles []*pb.CardResponseDeleteAt

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapCardResponseDeleteAt(role))
	}

	return responseRoles
}

// mapDashboardCard maps a *response.DashboardCard to a *pb.CardResponseDashboard proto response.
//
// This function converts a DashboardCard from the response domain model to its protobuf
// representation, handling the mapping of all financial metrics fields.
func (s *cardProtoMapper) mapDashboardCard(dash *response.DashboardCard) *pb.CardResponseDashboard {
	return &pb.CardResponseDashboard{
		TotalBalance:     *dash.TotalBalance,
		TotalWithdraw:    *dash.TotalWithdraw,
		TotalTopup:       *dash.TotalTopup,
		TotalTransfer:    *dash.TotalTransfer,
		TotalTransaction: *dash.TotalTransaction,
	}
}

// mapDashboardCardCardNumber maps a *response.DashboardCardCardNumber to a *pb.CardResponseDashboardCardNumber proto response.
//
// This function converts a DashboardCardCardNumber from the response domain model to its protobuf
// representation, handling the mapping of all financial metrics fields for a specific card.
func (s *cardProtoMapper) mapDashboardCardCardNumber(dash *response.DashboardCardCardNumber) *pb.CardResponseDashboardCardNumber {
	return &pb.CardResponseDashboardCardNumber{
		TotalBalance:          *dash.TotalBalance,
		TotalWithdraw:         *dash.TotalWithdraw,
		TotalTopup:            *dash.TotalTopup,
		TotalTransferSend:     *dash.TotalTransferSend,
		TotalTransferReceiver: *dash.TotalTransferReceiver,
		TotalTransaction:      *dash.TotalTransaction,
	}
}

// mapMonthlyBalance maps a *response.CardResponseMonthBalance to a *pb.CardResponseMonthlyBalance
// proto response.
//
// It is used to generate the response for the CardService.GetMonthlyBalance rpc method.
func (s *cardProtoMapper) mapMonthlyBalance(card *response.CardResponseMonthBalance) *pb.CardResponseMonthlyBalance {
	return &pb.CardResponseMonthlyBalance{
		Month:        card.Month,
		TotalBalance: card.TotalBalance,
	}
}

// mapMonthlyBalances maps a list of *response.CardResponseMonthBalance to a list of
// *pb.CardResponseMonthlyBalance proto responses.
//
// It iterates over each CardResponseMonthBalance in the input slice, converting
// them to their protobuf equivalent using the mapMonthlyBalance function. This
// function is used to generate the response data for monthly balance RPC methods.
func (s *cardProtoMapper) mapMonthlyBalances(roles []*response.CardResponseMonthBalance) []*pb.CardResponseMonthlyBalance {
	var responseRoles []*pb.CardResponseMonthlyBalance

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapMonthlyBalance(role))
	}

	return responseRoles
}

// mapYearlyBalance maps a *response.CardResponseYearlyBalance to a *pb.CardResponseYearlyBalance proto response.
//
// This function takes a CardResponseYearlyBalance from the response domain model and converts it
// into its protobuf representation, ensuring that the year and total balance fields are accurately mapped.
func (s *cardProtoMapper) mapYearlyBalance(card *response.CardResponseYearlyBalance) *pb.CardResponseYearlyBalance {
	return &pb.CardResponseYearlyBalance{
		Year:         card.Year,
		TotalBalance: card.TotalBalance,
	}
}

// mapYearlyBalances maps a list of *response.CardResponseYearlyBalance to a list of
// *pb.CardResponseYearlyBalance proto responses.
//
// It iterates over each CardResponseYearlyBalance in the input slice, converting
// them to their protobuf equivalent using the mapYearlyBalance function. This
// function is used to generate the response data for yearly balance RPC methods.
func (s *cardProtoMapper) mapYearlyBalances(roles []*response.CardResponseYearlyBalance) []*pb.CardResponseYearlyBalance {
	var responseRoles []*pb.CardResponseYearlyBalance

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapYearlyBalance(role))
	}

	return responseRoles
}

// mapMonthlyAmount maps a *response.CardResponseMonthAmount to a *pb.CardResponseMonthlyAmount proto response.
//
// It is used to generate the response for the CardService.GetMonthlyAmount rpc method.
func (s *cardProtoMapper) mapMonthlyAmount(card *response.CardResponseMonthAmount) *pb.CardResponseMonthlyAmount {
	return &pb.CardResponseMonthlyAmount{
		Month:       card.Month,
		TotalAmount: card.TotalAmount,
	}
}

// mapMonthlyAmounts maps a list of *response.CardResponseMonthAmount to a list of
// *pb.CardResponseMonthlyAmount proto responses.
//
// It iterates over each CardResponseMonthAmount in the input slice, converting
// them to their protobuf equivalent using the mapMonthlyAmount function. This
// function is used to generate the response data for monthly amount RPC methods.
func (s *cardProtoMapper) mapMonthlyAmounts(roles []*response.CardResponseMonthAmount) []*pb.CardResponseMonthlyAmount {
	var responseRoles []*pb.CardResponseMonthlyAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapMonthlyAmount(role))
	}

	return responseRoles
}

// mapYearlyAmount maps a *response.CardResponseYearAmount to a *pb.CardResponseYearlyAmount proto response.
//
// This function takes a CardResponseYearAmount from the response domain model and converts it
// into its protobuf representation, ensuring that the year and total amount fields are accurately mapped.
func (s *cardProtoMapper) mapYearlyAmount(card *response.CardResponseYearAmount) *pb.CardResponseYearlyAmount {
	return &pb.CardResponseYearlyAmount{
		Year:        card.Year,
		TotalAmount: card.TotalAmount,
	}
}

// mapYearlyAmounts maps a list of *response.CardResponseYearAmount to a list of
// *pb.CardResponseYearlyAmount proto responses.
//
// It iterates over each CardResponseYearAmount in the input slice, converting
// them to their protobuf equivalent using the mapYearlyAmount function. This
// function is used to generate the response data for yearly amount RPC methods.
func (s *cardProtoMapper) mapYearlyAmounts(roles []*response.CardResponseYearAmount) []*pb.CardResponseYearlyAmount {
	var responseRoles []*pb.CardResponseYearlyAmount

	for _, role := range roles {
		responseRoles = append(responseRoles, s.mapYearlyAmount(role))
	}

	return responseRoles
}
