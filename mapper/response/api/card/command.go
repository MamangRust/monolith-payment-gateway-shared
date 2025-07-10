package cardapimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb/card"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

type cardCommandResponseMapper struct {
}

func NewCardCommandResponseMapper() CardCommandResponseMapper {
	return &cardCommandResponseMapper{}
}

// ToApiResponseCard maps the ApiResponseCard from the domain to the ApiResponseCard of the api.
//
// Args:
//   - card: A pointer to a pb.ApiResponseCard representing the ApiResponseCard from the domain.
//
// Returns:
//   - A pointer to a response.ApiResponseCard containing the mapped data, including status, message,
//     and data.
func (s *cardCommandResponseMapper) ToApiResponseCard(card *pb.ApiResponseCard) *response.ApiResponseCard {
	return &response.ApiResponseCard{
		Status:  card.Status,
		Message: card.Message,
		Data:    s.mapCardResponse(card.Data),
	}
}

// mapCardResponse maps a CardResponse from the domain representation to the API response representation.
//
// Args:
//   - card: A pointer to a pb.CardResponse representing the domain CardResponse object.
//
// Returns:
//   - A pointer to a response.CardResponse containing the mapped data, including ID, UserID,
//     CardNumber, CardType, ExpireDate, CVV, CardProvider, CreatedAt, and UpdatedAt.
func (s *cardCommandResponseMapper) mapCardResponse(card *pb.CardResponse) *response.CardResponse {
	return &response.CardResponse{
		ID:           int(card.Id),
		UserID:       int(card.UserId),
		CardNumber:   card.CardNumber,
		CardType:     card.CardType,
		ExpireDate:   card.ExpireDate,
		CVV:          card.Cvv,
		CardProvider: card.CardProvider,
		CreatedAt:    card.CreatedAt,
		UpdatedAt:    card.UpdatedAt,
	}
}
