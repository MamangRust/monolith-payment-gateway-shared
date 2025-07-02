package responseservice

import (
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// CardResponseMapper defines methods for converting internal Card records
// into HTTP/REST API response formats.
//
//go:generate mockgen -source=interfaces.go -destination=mocks/interfaces.go
type CardResponseMapper interface {
	// Converts a single card record into a CardResponse.
	ToCardResponse(card *record.CardRecord) *response.CardResponse

	// Converts a list of card records into a list of CardResponse.
	ToCardsResponse(cards []*record.CardRecord) []*response.CardResponse

	// Converts a soft-deleted card record into a CardResponseDeleteAt.
	ToCardResponseDeleteAt(card *record.CardRecord) *response.CardResponseDeleteAt

	// Converts a list of soft-deleted card records into a list of CardResponseDeleteAt.
	ToCardsResponseDeleteAt(cards []*record.CardRecord) []*response.CardResponseDeleteAt

	// Converts a monthly card balance record into a response object.
	ToGetMonthlyBalance(card *record.CardMonthBalance) *response.CardResponseMonthBalance

	// Converts multiple monthly card balance records into a list of response objects.
	ToGetMonthlyBalances(cards []*record.CardMonthBalance) []*response.CardResponseMonthBalance

	// Converts a yearly card balance record into a response object.
	ToGetYearlyBalance(card *record.CardYearlyBalance) *response.CardResponseYearlyBalance

	// Converts multiple yearly card balance records into a list of response objects.
	ToGetYearlyBalances(cards []*record.CardYearlyBalance) []*response.CardResponseYearlyBalance

	// Converts a monthly card amount record into a response object.
	ToGetMonthlyAmount(card *record.CardMonthAmount) *response.CardResponseMonthAmount

	// Converts multiple monthly card amount records into a list of response objects.
	ToGetMonthlyAmounts(cards []*record.CardMonthAmount) []*response.CardResponseMonthAmount

	// Converts a yearly card amount record into a response object.
	ToGetYearlyAmount(card *record.CardYearAmount) *response.CardResponseYearAmount

	// Converts multiple yearly card amount records into a list of response objects.
	ToGetYearlyAmounts(cards []*record.CardYearAmount) []*response.CardResponseYearAmount
}

// UserResponseMapper defines methods for converting internal User records
// into HTTP/REST API response formats.
type UserResponseMapper interface {
	// Converts a single user record into a UserResponse.
	ToUserResponse(user *record.UserRecord) *response.UserResponse

	// Converts a list of user records into a list of UserResponse.
	ToUsersResponse(users []*record.UserRecord) []*response.UserResponse

	// Converts a soft-deleted user record into a UserResponseDeleteAt.
	ToUserResponseDeleteAt(user *record.UserRecord) *response.UserResponseDeleteAt

	// Converts a list of soft-deleted user records into a list of UserResponseDeleteAt.
	ToUsersResponseDeleteAt(users []*record.UserRecord) []*response.UserResponseDeleteAt
}

// RoleResponseMapper defines methods to map RoleRecord domain models
// into structured API response representations.
type RoleResponseMapper interface {
	// Converts a single RoleRecord into RoleResponse.
	ToRoleResponse(role *record.RoleRecord) *response.RoleResponse

	// Converts a slice of RoleRecord into a slice of RoleResponse.
	ToRolesResponse(roles []*record.RoleRecord) []*response.RoleResponse

	// Converts a deleted RoleRecord into RoleResponseDeleteAt.
	ToRoleResponseDeleteAt(role *record.RoleRecord) *response.RoleResponseDeleteAt

	// Converts a slice of deleted RoleRecord into RoleResponseDeleteAt slice.
	ToRolesResponseDeleteAt(roles []*record.RoleRecord) []*response.RoleResponseDeleteAt
}

// RefreshTokenResponseMapper defines methods to map RefreshTokenRecord
// into their corresponding API response structures.
type RefreshTokenResponseMapper interface {
	// Converts a single RefreshTokenRecord into RefreshTokenResponse.
	ToRefreshTokenResponse(refresh *record.RefreshTokenRecord) *response.RefreshTokenResponse

	// Converts multiple RefreshTokenRecord into a slice of RefreshTokenResponse.
	ToRefreshTokenResponses(refreshs []*record.RefreshTokenRecord) []*response.RefreshTokenResponse
}

// SaldoResponseMapper defines methods to map Saldo-related records
// into structured API response formats.
type SaldoResponseMapper interface {
	// Converts a single SaldoRecord into SaldoResponse.
	ToSaldoResponse(saldo *record.SaldoRecord) *response.SaldoResponse

	// Converts a list of SaldoRecord into SaldoResponse list.
	ToSaldoResponses(saldos []*record.SaldoRecord) []*response.SaldoResponse

	// Converts a monthly total balance record into API response.
	ToSaldoMonthTotalBalanceResponse(ss *record.SaldoMonthTotalBalance) *response.SaldoMonthTotalBalanceResponse

	// Converts multiple monthly total balance records into API response list.
	ToSaldoMonthTotalBalanceResponses(ss []*record.SaldoMonthTotalBalance) []*response.SaldoMonthTotalBalanceResponse

	// Converts a yearly total balance record into API response.
	ToSaldoYearTotalBalanceResponse(ss *record.SaldoYearTotalBalance) *response.SaldoYearTotalBalanceResponse

	// Converts multiple yearly total balance records into API response list.
	ToSaldoYearTotalBalanceResponses(ss []*record.SaldoYearTotalBalance) []*response.SaldoYearTotalBalanceResponse

	// Converts a monthly balance record into API response.
	ToSaldoMonthBalanceResponse(ss *record.SaldoMonthSaldoBalance) *response.SaldoMonthBalanceResponse

	// Converts multiple monthly balance records into API response list.
	ToSaldoMonthBalanceResponses(ss []*record.SaldoMonthSaldoBalance) []*response.SaldoMonthBalanceResponse

	// Converts a yearly balance record into API response.
	ToSaldoYearBalanceResponse(ss *record.SaldoYearSaldoBalance) *response.SaldoYearBalanceResponse

	// Converts multiple yearly balance records into API response list.
	ToSaldoYearBalanceResponses(ss []*record.SaldoYearSaldoBalance) []*response.SaldoYearBalanceResponse

	// Converts a soft-deleted SaldoRecord into SaldoResponseDeleteAt.
	ToSaldoResponseDeleteAt(saldo *record.SaldoRecord) *response.SaldoResponseDeleteAt

	// Converts multiple soft-deleted SaldoRecords into SaldoResponseDeleteAt list.
	ToSaldoResponsesDeleteAt(saldos []*record.SaldoRecord) []*response.SaldoResponseDeleteAt
}

// TopupResponseMapper defines methods for mapping top-up records into
// structured API response objects used in the application.
type TopupResponseMapper interface {
	// Converts a single TopupRecord into a TopupResponse.
	ToTopupResponse(topup *record.TopupRecord) *response.TopupResponse

	// Converts a slice of TopupRecord into a slice of TopupResponse.
	ToTopupResponses(topups []*record.TopupRecord) []*response.TopupResponse

	// Converts a monthly top-up success record into a TopupResponseMonthStatusSuccess.
	ToTopupResponseMonthStatusSuccess(s *record.TopupRecordMonthStatusSuccess) *response.TopupResponseMonthStatusSuccess

	// Converts multiple monthly top-up success records into a slice of TopupResponseMonthStatusSuccess.
	ToTopupResponsesMonthStatusSuccess(topups []*record.TopupRecordMonthStatusSuccess) []*response.TopupResponseMonthStatusSuccess

	// Converts a yearly top-up success record into a TopupResponseYearStatusSuccess.
	ToTopupResponseYearStatusSuccess(s *record.TopupRecordYearStatusSuccess) *response.TopupResponseYearStatusSuccess

	// Converts multiple yearly top-up success records into a slice of TopupResponseYearStatusSuccess.
	ToTopupResponsesYearStatusSuccess(topups []*record.TopupRecordYearStatusSuccess) []*response.TopupResponseYearStatusSuccess

	// Converts a monthly top-up failed record into a TopupResponseMonthStatusFailed.
	ToTopupResponseMonthStatusFailed(s *record.TopupRecordMonthStatusFailed) *response.TopupResponseMonthStatusFailed

	// Converts multiple monthly top-up failed records into a slice of TopupResponseMonthStatusFailed.
	ToTopupResponsesMonthStatusFailed(topups []*record.TopupRecordMonthStatusFailed) []*response.TopupResponseMonthStatusFailed

	// Converts a yearly top-up failed record into a TopupResponseYearStatusFailed.
	ToTopupResponseYearStatusFailed(s *record.TopupRecordYearStatusFailed) *response.TopupResponseYearStatusFailed

	// Converts multiple yearly top-up failed records into a slice of TopupResponseYearStatusFailed.
	ToTopupResponsesYearStatusFailed(topups []*record.TopupRecordYearStatusFailed) []*response.TopupResponseYearStatusFailed

	// Converts a monthly method record into a TopupMonthMethodResponse.
	ToTopupMonthlyMethodResponse(s *record.TopupMonthMethod) *response.TopupMonthMethodResponse

	// Converts multiple monthly method records into a slice of TopupMonthMethodResponse.
	ToTopupMonthlyMethodResponses(s []*record.TopupMonthMethod) []*response.TopupMonthMethodResponse

	// Converts a yearly method record into a TopupYearlyMethodResponse.
	ToTopupYearlyMethodResponse(s *record.TopupYearlyMethod) *response.TopupYearlyMethodResponse

	// Converts multiple yearly method records into a slice of TopupYearlyMethodResponse.
	ToTopupYearlyMethodResponses(s []*record.TopupYearlyMethod) []*response.TopupYearlyMethodResponse

	// Converts a monthly amount record into a TopupMonthAmountResponse.
	ToTopupMonthlyAmountResponse(s *record.TopupMonthAmount) *response.TopupMonthAmountResponse

	// Converts multiple monthly amount records into a slice of TopupMonthAmountResponse.
	ToTopupMonthlyAmountResponses(s []*record.TopupMonthAmount) []*response.TopupMonthAmountResponse

	// Converts a yearly amount record into a TopupYearlyAmountResponse.
	ToTopupYearlyAmountResponse(s *record.TopupYearlyAmount) *response.TopupYearlyAmountResponse

	// Converts multiple yearly amount records into a slice of TopupYearlyAmountResponse.
	ToTopupYearlyAmountResponses(s []*record.TopupYearlyAmount) []*response.TopupYearlyAmountResponse

	// Converts a soft-deleted TopupRecord into a TopupResponseDeleteAt.
	ToTopupResponseDeleteAt(topup *record.TopupRecord) *response.TopupResponseDeleteAt

	// Converts multiple soft-deleted TopupRecord into a slice of TopupResponseDeleteAt.
	ToTopupResponsesDeleteAt(topups []*record.TopupRecord) []*response.TopupResponseDeleteAt
}

// TransactionResponseMapper defines a set of methods for converting transaction records
// from the data layer into structured API response objects for use in handlers.
type TransactionResponseMapper interface {
	// Converts a single transaction record into a TransactionResponse.
	ToTransactionResponse(transaction *record.TransactionRecord) *response.TransactionResponse

	// Converts multiple transaction records into a slice of TransactionResponse.
	ToTransactionsResponse(transactions []*record.TransactionRecord) []*response.TransactionResponse

	// Converts a monthly success status transaction record into a TransactionResponseMonthStatusSuccess.
	ToTransactionResponseMonthStatusSuccess(s *record.TransactionRecordMonthStatusSuccess) *response.TransactionResponseMonthStatusSuccess

	// Converts multiple monthly success status transaction records into a slice of TransactionResponseMonthStatusSuccess.
	ToTransactionResponsesMonthStatusSuccess(Transactions []*record.TransactionRecordMonthStatusSuccess) []*response.TransactionResponseMonthStatusSuccess

	// Converts a yearly success status transaction record into a TransactionResponseYearStatusSuccess.
	ToTransactionResponseYearStatusSuccess(s *record.TransactionRecordYearStatusSuccess) *response.TransactionResponseYearStatusSuccess

	// Converts multiple yearly success status transaction records into a slice of TransactionResponseYearStatusSuccess.
	ToTransactionResponsesYearStatusSuccess(Transactions []*record.TransactionRecordYearStatusSuccess) []*response.TransactionResponseYearStatusSuccess

	// Converts a monthly failed status transaction record into a TransactionResponseMonthStatusFailed.
	ToTransactionResponseMonthStatusFailed(s *record.TransactionRecordMonthStatusFailed) *response.TransactionResponseMonthStatusFailed

	// Converts multiple monthly failed status transaction records into a slice of TransactionResponseMonthStatusFailed.
	ToTransactionResponsesMonthStatusFailed(Transactions []*record.TransactionRecordMonthStatusFailed) []*response.TransactionResponseMonthStatusFailed

	// Converts a yearly failed status transaction record into a TransactionResponseYearStatusFailed.
	ToTransactionResponseYearStatusFailed(s *record.TransactionRecordYearStatusFailed) *response.TransactionResponseYearStatusFailed

	// Converts multiple yearly failed status transaction records into a slice of TransactionResponseYearStatusFailed.
	ToTransactionResponsesYearStatusFailed(Transactions []*record.TransactionRecordYearStatusFailed) []*response.TransactionResponseYearStatusFailed

	// Converts a monthly method record into a TransactionMonthMethodResponse.
	ToTransactionMonthlyMethodResponse(s *record.TransactionMonthMethod) *response.TransactionMonthMethodResponse

	// Converts multiple monthly method records into a slice of TransactionMonthMethodResponse.
	ToTransactionMonthlyMethodResponses(s []*record.TransactionMonthMethod) []*response.TransactionMonthMethodResponse

	// Converts a yearly method record into a TransactionYearMethodResponse.
	ToTransactionYearlyMethodResponse(s *record.TransactionYearMethod) *response.TransactionYearMethodResponse

	// Converts multiple yearly method records into a slice of TransactionYearMethodResponse.
	ToTransactionYearlyMethodResponses(s []*record.TransactionYearMethod) []*response.TransactionYearMethodResponse

	// Converts a monthly amount record into a TransactionMonthAmountResponse.
	ToTransactionMonthlyAmountResponse(s *record.TransactionMonthAmount) *response.TransactionMonthAmountResponse

	// Converts multiple monthly amount records into a slice of TransactionMonthAmountResponse.
	ToTransactionMonthlyAmountResponses(s []*record.TransactionMonthAmount) []*response.TransactionMonthAmountResponse

	// Converts a yearly amount record into a TransactionYearlyAmountResponse.
	ToTransactionYearlyAmountResponse(s *record.TransactionYearlyAmount) *response.TransactionYearlyAmountResponse

	// Converts multiple yearly amount records into a slice of TransactionYearlyAmountResponse.
	ToTransactionYearlyAmountResponses(s []*record.TransactionYearlyAmount) []*response.TransactionYearlyAmountResponse

	// Converts a soft-deleted transaction record into a TransactionResponseDeleteAt.
	ToTransactionResponseDeleteAt(transaction *record.TransactionRecord) *response.TransactionResponseDeleteAt

	// Converts multiple soft-deleted transaction records into a slice of TransactionResponseDeleteAt.
	ToTransactionsResponseDeleteAt(transactions []*record.TransactionRecord) []*response.TransactionResponseDeleteAt
}

// TransferResponseMapper defines a set of methods for converting transfer-related records
// from the database layer into structured API responses used by the application layer.
type TransferResponseMapper interface {
	// Converts a single transfer record into a TransferResponse.
	ToTransferResponse(transfer *record.TransferRecord) *response.TransferResponse

	// Converts multiple transfer records into a slice of TransferResponse.
	ToTransfersResponse(transfers []*record.TransferRecord) []*response.TransferResponse

	// Converts a monthly success status transfer record into a TransferResponseMonthStatusSuccess.
	ToTransferResponseMonthStatusSuccess(s *record.TransferRecordMonthStatusSuccess) *response.TransferResponseMonthStatusSuccess

	// Converts multiple monthly success status transfer records into a slice of TransferResponseMonthStatusSuccess.
	ToTransferResponsesMonthStatusSuccess(Transfers []*record.TransferRecordMonthStatusSuccess) []*response.TransferResponseMonthStatusSuccess

	// Converts a yearly success status transfer record into a TransferResponseYearStatusSuccess.
	ToTransferResponseYearStatusSuccess(s *record.TransferRecordYearStatusSuccess) *response.TransferResponseYearStatusSuccess

	// Converts multiple yearly success status transfer records into a slice of TransferResponseYearStatusSuccess.
	ToTransferResponsesYearStatusSuccess(Transfers []*record.TransferRecordYearStatusSuccess) []*response.TransferResponseYearStatusSuccess

	// Converts a monthly failed status transfer record into a TransferResponseMonthStatusFailed.
	ToTransferResponseMonthStatusFailed(s *record.TransferRecordMonthStatusFailed) *response.TransferResponseMonthStatusFailed

	// Converts multiple monthly failed status transfer records into a slice of TransferResponseMonthStatusFailed.
	ToTransferResponsesMonthStatusFailed(Transfers []*record.TransferRecordMonthStatusFailed) []*response.TransferResponseMonthStatusFailed

	// Converts a yearly failed status transfer record into a TransferResponseYearStatusFailed.
	ToTransferResponseYearStatusFailed(s *record.TransferRecordYearStatusFailed) *response.TransferResponseYearStatusFailed

	// Converts multiple yearly failed status transfer records into a slice of TransferResponseYearStatusFailed.
	ToTransferResponsesYearStatusFailed(Transfers []*record.TransferRecordYearStatusFailed) []*response.TransferResponseYearStatusFailed

	// Converts a monthly transfer amount record into a TransferMonthAmountResponse.
	ToTransferResponseMonthAmount(s *record.TransferMonthAmount) *response.TransferMonthAmountResponse

	// Converts multiple monthly transfer amount records into a slice of TransferMonthAmountResponse.
	ToTransferResponsesMonthAmount(s []*record.TransferMonthAmount) []*response.TransferMonthAmountResponse

	// Converts a yearly transfer amount record into a TransferYearAmountResponse.
	ToTransferResponseYearAmount(s *record.TransferYearAmount) *response.TransferYearAmountResponse

	// Converts multiple yearly transfer amount records into a slice of TransferYearAmountResponse.
	ToTransferResponsesYearAmount(s []*record.TransferYearAmount) []*response.TransferYearAmountResponse

	// Converts a soft-deleted transfer record into a TransferResponseDeleteAt.
	ToTransferResponseDeleteAt(transfer *record.TransferRecord) *response.TransferResponseDeleteAt

	// Converts multiple soft-deleted transfer records into a slice of TransferResponseDeleteAt.
	ToTransfersResponseDeleteAt(transfers []*record.TransferRecord) []*response.TransferResponseDeleteAt
}

// WithdrawResponseMapper defines methods to map withdrawal-related database records
// to structured API response objects used in handlers or services.
type WithdrawResponseMapper interface {
	// Converts a single WithdrawRecord into a WithdrawResponse.
	ToWithdrawResponse(withdraw *record.WithdrawRecord) *response.WithdrawResponse

	// Converts multiple WithdrawRecords into a slice of WithdrawResponse.
	ToWithdrawsResponse(withdraws []*record.WithdrawRecord) []*response.WithdrawResponse

	// Converts a monthly success-status WithdrawRecord into a WithdrawResponseMonthStatusSuccess.
	ToWithdrawResponseMonthStatusSuccess(s *record.WithdrawRecordMonthStatusSuccess) *response.WithdrawResponseMonthStatusSuccess

	// Converts multiple monthly success-status WithdrawRecords into a slice of WithdrawResponseMonthStatusSuccess.
	ToWithdrawResponsesMonthStatusSuccess(Withdraws []*record.WithdrawRecordMonthStatusSuccess) []*response.WithdrawResponseMonthStatusSuccess

	// Converts a yearly success-status WithdrawRecord into a WithdrawResponseYearStatusSuccess.
	ToWithdrawResponseYearStatusSuccess(s *record.WithdrawRecordYearStatusSuccess) *response.WithdrawResponseYearStatusSuccess

	// Converts multiple yearly success-status WithdrawRecords into a slice of WithdrawResponseYearStatusSuccess.
	ToWithdrawResponsesYearStatusSuccess(Withdraws []*record.WithdrawRecordYearStatusSuccess) []*response.WithdrawResponseYearStatusSuccess

	// Converts a monthly failed-status WithdrawRecord into a WithdrawResponseMonthStatusFailed.
	ToWithdrawResponseMonthStatusFailed(s *record.WithdrawRecordMonthStatusFailed) *response.WithdrawResponseMonthStatusFailed

	// Converts multiple monthly failed-status WithdrawRecords into a slice of WithdrawResponseMonthStatusFailed.
	ToWithdrawResponsesMonthStatusFailed(Withdraws []*record.WithdrawRecordMonthStatusFailed) []*response.WithdrawResponseMonthStatusFailed

	// Converts a yearly failed-status WithdrawRecord into a WithdrawResponseYearStatusFailed.
	ToWithdrawResponseYearStatusFailed(s *record.WithdrawRecordYearStatusFailed) *response.WithdrawResponseYearStatusFailed

	// Converts multiple yearly failed-status WithdrawRecords into a slice of WithdrawResponseYearStatusFailed.
	ToWithdrawResponsesYearStatusFailed(Withdraws []*record.WithdrawRecordYearStatusFailed) []*response.WithdrawResponseYearStatusFailed

	// Converts a monthly amount summary into a WithdrawMonthlyAmountResponse.
	ToWithdrawAmountMonthlyResponse(s *record.WithdrawMonthlyAmount) *response.WithdrawMonthlyAmountResponse

	// Converts multiple monthly amount summaries into a slice of WithdrawMonthlyAmountResponse.
	ToWithdrawsAmountMonthlyResponses(s []*record.WithdrawMonthlyAmount) []*response.WithdrawMonthlyAmountResponse

	// Converts a yearly amount summary into a WithdrawYearlyAmountResponse.
	ToWithdrawAmountYearlyResponse(s *record.WithdrawYearlyAmount) *response.WithdrawYearlyAmountResponse

	// Converts multiple yearly amount summaries into a slice of WithdrawYearlyAmountResponse.
	ToWithdrawsAmountYearlyResponses(s []*record.WithdrawYearlyAmount) []*response.WithdrawYearlyAmountResponse

	// Converts a soft-deleted WithdrawRecord into a WithdrawResponseDeleteAt.
	ToWithdrawResponseDeleteAt(withdraw *record.WithdrawRecord) *response.WithdrawResponseDeleteAt

	// Converts multiple soft-deleted WithdrawRecords into a slice of WithdrawResponseDeleteAt.
	ToWithdrawsResponseDeleteAt(withdraws []*record.WithdrawRecord) []*response.WithdrawResponseDeleteAt
}

// MerchantResponseMapper defines methods for converting merchant-related database
// records into API-compatible response types.
type MerchantResponseMapper interface {
	// Converts a single MerchantRecord to a MerchantResponse.
	ToMerchantResponse(merchant *record.MerchantRecord) *response.MerchantResponse

	// Converts multiple MerchantRecords to a slice of MerchantResponse.
	ToMerchantsResponse(merchants []*record.MerchantRecord) []*response.MerchantResponse

	// Converts a monthly total amount record into a MerchantResponseMonthlyTotalAmount.
	ToMerchantMonthlyTotalAmount(ms *record.MerchantMonthlyTotalAmount) *response.MerchantResponseMonthlyTotalAmount

	// Converts multiple monthly total amount records into a slice of MerchantResponseMonthlyTotalAmount.
	ToMerchantMonthlyTotalAmounts(ms []*record.MerchantMonthlyTotalAmount) []*response.MerchantResponseMonthlyTotalAmount

	// Converts a yearly total amount record into a MerchantResponseYearlyTotalAmount.
	ToMerchantYearlyTotalAmount(ms *record.MerchantYearlyTotalAmount) *response.MerchantResponseYearlyTotalAmount

	// Converts multiple yearly total amount records into a slice of MerchantResponseYearlyTotalAmount.
	ToMerchantYearlyTotalAmounts(ms []*record.MerchantYearlyTotalAmount) []*response.MerchantResponseYearlyTotalAmount

	// Converts a merchant transaction record into a MerchantTransactionResponse.
	ToMerchantTransactionResponse(merchant *record.MerchantTransactionsRecord) *response.MerchantTransactionResponse

	// Converts multiple merchant transaction records into a slice of MerchantTransactionResponse.
	ToMerchantsTransactionResponse(merchants []*record.MerchantTransactionsRecord) []*response.MerchantTransactionResponse

	// Converts a monthly payment method record into a MerchantResponseMonthlyPaymentMethod.
	ToMerchantMonthlyPaymentMethod(ms *record.MerchantMonthlyPaymentMethod) *response.MerchantResponseMonthlyPaymentMethod

	// Converts multiple monthly payment method records into a slice of MerchantResponseMonthlyPaymentMethod.
	ToMerchantMonthlyPaymentMethods(ms []*record.MerchantMonthlyPaymentMethod) []*response.MerchantResponseMonthlyPaymentMethod

	// Converts a yearly payment method record into a MerchantResponseYearlyPaymentMethod.
	ToMerchantYearlyPaymentMethod(ms *record.MerchantYearlyPaymentMethod) *response.MerchantResponseYearlyPaymentMethod

	// Converts multiple yearly payment method records into a slice of MerchantResponseYearlyPaymentMethod.
	ToMerchantYearlyPaymentMethods(ms []*record.MerchantYearlyPaymentMethod) []*response.MerchantResponseYearlyPaymentMethod

	// Converts a monthly amount record into a MerchantResponseMonthlyAmount.
	ToMerchantMonthlyAmount(ms *record.MerchantMonthlyAmount) *response.MerchantResponseMonthlyAmount

	// Converts multiple monthly amount records into a slice of MerchantResponseMonthlyAmount.
	ToMerchantMonthlyAmounts(ms []*record.MerchantMonthlyAmount) []*response.MerchantResponseMonthlyAmount

	// Converts a yearly amount record into a MerchantResponseYearlyAmount.
	ToMerchantYearlyAmount(ms *record.MerchantYearlyAmount) *response.MerchantResponseYearlyAmount

	// Converts multiple yearly amount records into a slice of MerchantResponseYearlyAmount.
	ToMerchantYearlyAmounts(ms []*record.MerchantYearlyAmount) []*response.MerchantResponseYearlyAmount

	// Converts a soft-deleted MerchantRecord into a MerchantResponseDeleteAt.
	ToMerchantResponseDeleteAt(merchant *record.MerchantRecord) *response.MerchantResponseDeleteAt

	// Converts multiple soft-deleted MerchantRecords into a slice of MerchantResponseDeleteAt.
	ToMerchantsResponseDeleteAt(merchants []*record.MerchantRecord) []*response.MerchantResponseDeleteAt
}

// MerchantDocumentResponseMapper defines methods for converting merchant document
// records into API-compatible response types.
type MerchantDocumentResponseMapper interface {
	// Converts a single MerchantDocumentRecord to a MerchantDocumentResponse.
	ToMerchantDocumentResponse(doc *record.MerchantDocumentRecord) *response.MerchantDocumentResponse

	// Converts multiple MerchantDocumentRecords to a slice of MerchantDocumentResponse.
	ToMerchantDocumentsResponse(docs []*record.MerchantDocumentRecord) []*response.MerchantDocumentResponse

	// Converts a soft-deleted MerchantDocumentRecord to its corresponding response.
	ToMerchantDocumentResponseDeleteAt(doc *record.MerchantDocumentRecord) *response.MerchantDocumentResponseDeleteAt

	// Converts multiple soft-deleted MerchantDocumentRecords to their corresponding responses.
	ToMerchantDocumentsResponseDeleteAt(docs []*record.MerchantDocumentRecord) []*response.MerchantDocumentResponseDeleteAt
}
