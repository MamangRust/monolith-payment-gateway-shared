package apimapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// AuthResponseMapper defines methods to map gRPC Auth responses to HTTP API responses.
//
//go:generate mockgen -source=interfaces.go -destination=mocks/interfaces.go
type AuthResponseMapper interface {
	// ToResponseVerifyCode maps gRPC response to API response for verify code operation.
	ToResponseVerifyCode(res *pb.ApiResponseVerifyCode) *response.ApiResponseVerifyCode

	// ToResponseForgotPassword maps gRPC response to API response for forgot password.
	ToResponseForgotPassword(res *pb.ApiResponseForgotPassword) *response.ApiResponseForgotPassword

	// ToResponseResetPassword maps gRPC response to API response for reset password.
	ToResponseResetPassword(res *pb.ApiResponseResetPassword) *response.ApiResponseResetPassword

	// ToResponseLogin maps gRPC response to API response for login.
	ToResponseLogin(res *pb.ApiResponseLogin) *response.ApiResponseLogin

	// ToResponseRegister maps gRPC response to API response for register.
	ToResponseRegister(res *pb.ApiResponseRegister) *response.ApiResponseRegister

	// ToResponseRefreshToken maps gRPC response to API response for refresh token.
	ToResponseRefreshToken(res *pb.ApiResponseRefreshToken) *response.ApiResponseRefreshToken

	// ToResponseGetMe maps gRPC response to API response for get user info.
	ToResponseGetMe(res *pb.ApiResponseGetMe) *response.ApiResponseGetMe
}

// RoleResponseMapper defines a set of methods to map gRPC Role API responses
type RoleResponseMapper interface {

	// ToApiResponseRoleAll maps a gRPC response containing all roles
	// into an HTTP API response format.
	ToApiResponseRoleAll(pbResponse *pb.ApiResponseRoleAll) *response.ApiResponseRoleAll

	// ToApiResponseRoleDelete maps a gRPC delete role response
	// into an HTTP API response format.
	ToApiResponseRoleDelete(pbResponse *pb.ApiResponseRoleDelete) *response.ApiResponseRoleDelete

	// ToApiResponseRole maps a single gRPC role response
	// into an HTTP API response format.
	ToApiResponseRole(pbResponse *pb.ApiResponseRole) *response.ApiResponseRole

	// ToApiResponsesRole maps a gRPC response containing multiple roles
	// into a list HTTP API response format.
	ToApiResponsesRole(pbResponse *pb.ApiResponsesRole) *response.ApiResponsesRole

	// ToApiResponsePaginationRole maps a paginated gRPC response of roles
	// into a paginated HTTP API response format.
	ToApiResponsePaginationRole(pbResponse *pb.ApiResponsePaginationRole) *response.ApiResponsePaginationRole

	// ToApiResponsePaginationRoleDeleteAt maps a paginated gRPC response
	// of soft-deleted roles into a paginated HTTP API response format.
	ToApiResponsePaginationRoleDeleteAt(pbResponse *pb.ApiResponsePaginationRoleDeleteAt) *response.ApiResponsePaginationRoleDeleteAt
}

// CardResponseMapper defines methods for converting gRPC card-related responses
// (protobuf format) into HTTP-compatible API response structures used in the presentation layer.
type CardResponseMapper interface {

	// ToApiResponseCard maps a single gRPC card response to an HTTP API response.
	ToApiResponseCard(card *pb.ApiResponseCard) *response.ApiResponseCard

	// ToApiResponsesCard maps a paginated list of gRPC card responses
	// to a paginated HTTP API response.
	ToApiResponsesCard(cards *pb.ApiResponsePaginationCard) *response.ApiResponsePaginationCard

	// ToApiResponseCardAll maps a gRPC response containing all cards
	// to an HTTP API response format.
	ToApiResponseCardAll(card *pb.ApiResponseCardAll) *response.ApiResponseCardAll

	// ToApiResponseCardDeleteAt maps a soft-deleted gRPC card response
	// to an HTTP API response format.
	ToApiResponseCardDeleteAt(card *pb.ApiResponseCardDelete) *response.ApiResponseCardDelete

	// ToApiResponsesCardDeletedAt maps a paginated list of soft-deleted gRPC cards
	// to a paginated HTTP API response.
	ToApiResponsesCardDeletedAt(cards *pb.ApiResponsePaginationCardDeleteAt) *response.ApiResponsePaginationCardDeleteAt

	// ToApiResponseDashboardCard maps a gRPC response of general card dashboard statistics
	// to an HTTP API response.
	ToApiResponseDashboardCard(dash *pb.ApiResponseDashboardCard) *response.ApiResponseDashboardCard

	// ToApiResponseDashboardCardCardNumber maps a gRPC response of card-number-specific dashboard statistics
	// to an HTTP API response.
	ToApiResponseDashboardCardCardNumber(dash *pb.ApiResponseDashboardCardNumber) *response.ApiResponseDashboardCardNumber

	// ToApiResponseMonthlyBalances maps a gRPC response containing monthly balance statistics
	// to an HTTP API response format.
	ToApiResponseMonthlyBalances(cards *pb.ApiResponseMonthlyBalance) *response.ApiResponseMonthlyBalance

	// ToApiResponseYearlyBalances maps a gRPC response containing yearly balance statistics
	// to an HTTP API response format.
	ToApiResponseYearlyBalances(cards *pb.ApiResponseYearlyBalance) *response.ApiResponseYearlyBalance

	// ToApiResponseMonthlyAmounts maps a gRPC response containing monthly financial amounts
	// (e.g., top-up, withdrawal, transaction) to an HTTP API response format.
	ToApiResponseMonthlyAmounts(cards *pb.ApiResponseMonthlyAmount) *response.ApiResponseMonthlyAmount

	// ToApiResponseYearlyAmounts maps a gRPC response containing yearly financial amounts
	// (e.g., top-up, withdrawal, transaction) to an HTTP API response format.
	ToApiResponseYearlyAmounts(cards *pb.ApiResponseYearlyAmount) *response.ApiResponseYearlyAmount
}

// MerchantResponseMapper defines methods for converting gRPC merchant-related responses
// to HTTP-compatible API response formats used in the REST layer.
type MerchantResponseMapper interface {

	// Maps a single merchant gRPC response to an HTTP API response.
	ToApiResponseMerchant(merchants *pb.ApiResponseMerchant) *response.ApiResponseMerchant

	// Maps multiple merchant gRPC responses to a slice-based API response.
	ToApiResponseMerchants(merchants *pb.ApiResponsesMerchant) *response.ApiResponsesMerchant

	// Maps a paginated list of merchants to a paginated HTTP API response.
	ToApiResponsesMerchant(merchants *pb.ApiResponsePaginationMerchant) *response.ApiResponsePaginationMerchant

	// Maps a paginated list of soft-deleted merchants to an HTTP API response.
	ToApiResponsesMerchantDeleteAt(merchants *pb.ApiResponsePaginationMerchantDeleteAt) *response.ApiResponsePaginationMerchantDeleteAt

	// Maps a response containing all merchants to an HTTP API format.
	ToApiResponseMerchantAll(card *pb.ApiResponseMerchantAll) *response.ApiResponseMerchantAll

	// Maps a soft-deleted merchant response to an HTTP API format.
	ToApiResponseMerchantDeleteAt(card *pb.ApiResponseMerchantDelete) *response.ApiResponseMerchantDelete

	// Maps a paginated response of merchant transactions to an HTTP API response.
	ToApiResponseMerchantsTransactionResponse(merchants *pb.ApiResponsePaginationMerchantTransaction) *response.ApiResponsePaginationMerchantTransaction

	// Maps monthly payment method statistics of a merchant to an API response.
	ToApiResponseMonthlyPaymentMethods(ms *pb.ApiResponseMerchantMonthlyPaymentMethod) *response.ApiResponseMerchantMonthlyPaymentMethod

	// Maps yearly payment method statistics of a merchant to an API response.
	ToApiResponseYearlyPaymentMethods(ms *pb.ApiResponseMerchantYearlyPaymentMethod) *response.ApiResponseMerchantYearlyPaymentMethod

	// Maps monthly financial amounts (e.g., transactions, top-ups) to an API response.
	ToApiResponseMonthlyAmounts(ms *pb.ApiResponseMerchantMonthlyAmount) *response.ApiResponseMerchantMonthlyAmount

	// Maps yearly financial amounts (e.g., transactions, top-ups) to an API response.
	ToApiResponseYearlyAmounts(ms *pb.ApiResponseMerchantYearlyAmount) *response.ApiResponseMerchantYearlyAmount

	// Maps monthly total financial statistics across merchants to an API response.
	ToApiResponseMonthlyTotalAmounts(ms *pb.ApiResponseMerchantMonthlyTotalAmount) *response.ApiResponseMerchantMonthlyTotalAmount

	// Maps yearly total financial statistics across merchants to an API response.
	ToApiResponseYearlyTotalAmounts(ms *pb.ApiResponseMerchantYearlyTotalAmount) *response.ApiResponseMerchantYearlyTotalAmount
}

// MerchantDocumentResponseMapper defines methods to map gRPC merchant document responses
// into HTTP API response structures used in the presentation layer.
type MerchantDocumentResponseMapper interface {

	// Maps a single merchant document to an HTTP API response.
	ToApiResponseMerchantDocument(doc *pb.ApiResponseMerchantDocument) *response.ApiResponseMerchantDocument

	// Maps multiple merchant documents to a response slice.
	ToApiResponsesMerchantDocument(docs *pb.ApiResponsesMerchantDocument) *response.ApiResponsesMerchantDocument

	// Maps paginated merchant document responses to an HTTP API response.
	ToApiResponsePaginationMerchantDocument(docs *pb.ApiResponsePaginationMerchantDocument) *response.ApiResponsePaginationMerchantDocument

	// Maps paginated soft-deleted merchant document responses to an HTTP API response.
	ToApiResponsePaginationMerchantDocumentDeleteAt(docs *pb.ApiResponsePaginationMerchantDocumentAt) *response.ApiResponsePaginationMerchantDocumentDeleteAt

	// Maps all merchant documents to an API-compatible response.
	ToApiResponseMerchantDocumentAll(resp *pb.ApiResponseMerchantDocumentAll) *response.ApiResponseMerchantDocumentAll

	// Maps a soft-deleted single merchant document response to an API response.
	ToApiResponseMerchantDocumentDeleteAt(resp *pb.ApiResponseMerchantDocumentDelete) *response.ApiResponseMerchantDocumentDelete
}

// SaldoResponseMapper defines methods for transforming gRPC saldo-related responses
// into HTTP-friendly API response formats used in the RESTful layer.
type SaldoResponseMapper interface {

	// Converts a single saldo gRPC response to an API response.
	ToApiResponseSaldo(pbResponse *pb.ApiResponseSaldo) *response.ApiResponseSaldo

	// Converts multiple saldo responses into a slice-based API response.
	ToApiResponsesSaldo(pbResponse *pb.ApiResponsesSaldo) *response.ApiResponsesSaldo

	// Converts a deleted saldo response into an API response.
	ToApiResponseSaldoDelete(pbResponse *pb.ApiResponseSaldoDelete) *response.ApiResponseSaldoDelete

	// Converts all saldo records into a comprehensive API response.
	ToApiResponseSaldoAll(pbResponse *pb.ApiResponseSaldoAll) *response.ApiResponseSaldoAll

	// Converts monthly total saldo values into an API response.
	ToApiResponseMonthTotalSaldo(pbResponse *pb.ApiResponseMonthTotalSaldo) *response.ApiResponseMonthTotalSaldo

	// Converts yearly total saldo values into an API response.
	ToApiResponseYearTotalSaldo(pbResponse *pb.ApiResponseYearTotalSaldo) *response.ApiResponseYearTotalSaldo

	// Converts monthly saldo balances into an API response.
	ToApiResponseMonthSaldoBalances(pbResponse *pb.ApiResponseMonthSaldoBalances) *response.ApiResponseMonthSaldoBalances

	// Converts yearly saldo balances into an API response.
	ToApiResponseYearSaldoBalances(pbResponse *pb.ApiResponseYearSaldoBalances) *response.ApiResponseYearSaldoBalances

	// Converts paginated saldo responses into an API response.
	ToApiResponsePaginationSaldo(pbResponse *pb.ApiResponsePaginationSaldo) *response.ApiResponsePaginationSaldo

	// Converts paginated soft-deleted saldo responses into an API response.
	ToApiResponsePaginationSaldoDeleteAt(pbResponse *pb.ApiResponsePaginationSaldoDeleteAt) *response.ApiResponsePaginationSaldoDeleteAt
}

// TopupResponseMapper defines methods to convert gRPC top-up responses into API responses
// for use in the HTTP layer of the application.
type TopupResponseMapper interface {

	// Converts a single top-up response into an API format.
	ToApiResponseTopup(s *pb.ApiResponseTopup) *response.ApiResponseTopup

	// Converts a soft-deleted top-up response into an API format.
	ToApiResponseTopupDeleteAt(s *pb.ApiResponseTopupDeleteAt) *response.ApiResponseTopupDeleteAt

	// Converts all top-up records into a general API response.
	ToApiResponseTopupAll(s *pb.ApiResponseTopupAll) *response.ApiResponseTopupAll

	// Converts a permanently deleted top-up response into an API format.
	ToApiResponseTopupDelete(s *pb.ApiResponseTopupDelete) *response.ApiResponseTopupDelete

	// Converts a paginated list of top-ups into an API response.
	ToApiResponsePaginationTopup(s *pb.ApiResponsePaginationTopup) *response.ApiResponsePaginationTopup

	// Converts a paginated list of soft-deleted top-ups into an API response.
	ToApiResponsePaginationTopupDeleteAt(s *pb.ApiResponsePaginationTopupDeleteAt) *response.ApiResponsePaginationTopupDeleteAt

	// Converts monthly successful top-up stats into an API response.
	ToApiResponseTopupMonthStatusSuccess(s *pb.ApiResponseTopupMonthStatusSuccess) *response.ApiResponseTopupMonthStatusSuccess

	// Converts yearly successful top-up stats into an API response.
	ToApiResponseTopupYearStatusSuccess(s *pb.ApiResponseTopupYearStatusSuccess) *response.ApiResponseTopupYearStatusSuccess

	// Converts monthly failed top-up stats into an API response.
	ToApiResponseTopupMonthStatusFailed(s *pb.ApiResponseTopupMonthStatusFailed) *response.ApiResponseTopupMonthStatusFailed

	// Converts yearly failed top-up stats into an API response.
	ToApiResponseTopupYearStatusFailed(s *pb.ApiResponseTopupYearStatusFailed) *response.ApiResponseTopupYearStatusFailed

	// Converts monthly top-up statistics by payment method into an API response.
	ToApiResponseTopupMonthMethod(s *pb.ApiResponseTopupMonthMethod) *response.ApiResponseTopupMonthMethod

	// Converts yearly top-up statistics by payment method into an API response.
	ToApiResponseTopupYearMethod(s *pb.ApiResponseTopupYearMethod) *response.ApiResponseTopupYearMethod

	// Converts monthly top-up amount statistics into an API response.
	ToApiResponseTopupMonthAmount(s *pb.ApiResponseTopupMonthAmount) *response.ApiResponseTopupMonthAmount

	// Converts yearly top-up amount statistics into an API response.
	ToApiResponseTopupYearAmount(s *pb.ApiResponseTopupYearAmount) *response.ApiResponseTopupYearAmount
}

// TransactionResponseMapper defines methods for converting gRPC transaction-related responses
// into API responses suitable for the HTTP/REST layer.
type TransactionResponseMapper interface {

	// Converts monthly transaction stats with success status into an API response.
	ToApiResponseTransactionMonthStatusSuccess(pbResponse *pb.ApiResponseTransactionMonthStatusSuccess) *response.ApiResponseTransactionMonthStatusSuccess

	// Converts yearly transaction stats with success status into an API response.
	ToApiResponseTransactionYearStatusSuccess(pbResponse *pb.ApiResponseTransactionYearStatusSuccess) *response.ApiResponseTransactionYearStatusSuccess

	// Converts monthly transaction stats with failed status into an API response.
	ToApiResponseTransactionMonthStatusFailed(pbResponse *pb.ApiResponseTransactionMonthStatusFailed) *response.ApiResponseTransactionMonthStatusFailed

	// Converts yearly transaction stats with failed status into an API response.
	ToApiResponseTransactionYearStatusFailed(pbResponse *pb.ApiResponseTransactionYearStatusFailed) *response.ApiResponseTransactionYearStatusFailed

	// Converts monthly transaction statistics grouped by payment method into an API response.
	ToApiResponseTransactionMonthMethod(pbResponse *pb.ApiResponseTransactionMonthMethod) *response.ApiResponseTransactionMonthMethod

	// Converts yearly transaction statistics grouped by payment method into an API response.
	ToApiResponseTransactionYearMethod(pbResponse *pb.ApiResponseTransactionYearMethod) *response.ApiResponseTransactionYearMethod

	// Converts monthly transaction amount statistics into an API response.
	ToApiResponseTransactionMonthAmount(pbResponse *pb.ApiResponseTransactionMonthAmount) *response.ApiResponseTransactionMonthAmount

	// Converts yearly transaction amount statistics into an API response.
	ToApiResponseTransactionYearAmount(pbResponse *pb.ApiResponseTransactionYearAmount) *response.ApiResponseTransactionYearAmount

	// Converts a single transaction response into an API response.
	ToApiResponseTransaction(pbResponse *pb.ApiResponseTransaction) *response.ApiResponseTransaction

	// Converts multiple transaction responses into a grouped API response.
	ToApiResponseTransactions(pbResponse *pb.ApiResponseTransactions) *response.ApiResponseTransactions

	// Converts a deleted transaction response into an API response.
	ToApiResponseTransactionDelete(pbResponse *pb.ApiResponseTransactionDelete) *response.ApiResponseTransactionDelete

	// Converts all transaction records into a general API response.
	ToApiResponseTransactionAll(pbResponse *pb.ApiResponseTransactionAll) *response.ApiResponseTransactionAll

	// Converts paginated transaction results into an API response.
	ToApiResponsePaginationTransaction(pbResponse *pb.ApiResponsePaginationTransaction) *response.ApiResponsePaginationTransaction

	// Converts paginated soft-deleted transaction results into an API response.
	ToApiResponsePaginationTransactionDeleteAt(pbResponse *pb.ApiResponsePaginationTransactionDeleteAt) *response.ApiResponsePaginationTransactionDeleteAt
}

// TransferResponseMapper defines methods for mapping gRPC transfer-related responses
// into HTTP/REST API responses.
type TransferResponseMapper interface {

	// Converts monthly successful transfer statistics into an API response.
	ToApiResponseTransferMonthStatusSuccess(pbResponse *pb.ApiResponseTransferMonthStatusSuccess) *response.ApiResponseTransferMonthStatusSuccess

	// Converts yearly successful transfer statistics into an API response.
	ToApiResponseTransferYearStatusSuccess(pbResponse *pb.ApiResponseTransferYearStatusSuccess) *response.ApiResponseTransferYearStatusSuccess

	// Converts monthly failed transfer statistics into an API response.
	ToApiResponseTransferMonthStatusFailed(pbResponse *pb.ApiResponseTransferMonthStatusFailed) *response.ApiResponseTransferMonthStatusFailed

	// Converts yearly failed transfer statistics into an API response.
	ToApiResponseTransferYearStatusFailed(pbResponse *pb.ApiResponseTransferYearStatusFailed) *response.ApiResponseTransferYearStatusFailed

	// Converts monthly total transfer amount statistics into an API response.
	ToApiResponseTransferMonthAmount(pbResponse *pb.ApiResponseTransferMonthAmount) *response.ApiResponseTransferMonthAmount

	// Converts yearly total transfer amount statistics into an API response.
	ToApiResponseTransferYearAmount(pbResponse *pb.ApiResponseTransferYearAmount) *response.ApiResponseTransferYearAmount

	// Converts a single transfer response into an API response.
	ToApiResponseTransfer(pbResponse *pb.ApiResponseTransfer) *response.ApiResponseTransfer

	// Converts a list of transfers into a grouped API response.
	ToApiResponseTransfers(pbResponse *pb.ApiResponseTransfers) *response.ApiResponseTransfers

	// Converts a deleted transfer response into an API response.
	ToApiResponseTransferDelete(pbResponse *pb.ApiResponseTransferDelete) *response.ApiResponseTransferDelete

	// Converts all transfer records into an API response.
	ToApiResponseTransferAll(pbResponse *pb.ApiResponseTransferAll) *response.ApiResponseTransferAll

	// Converts paginated transfer records into an API response.
	ToApiResponsePaginationTransfer(pbResponse *pb.ApiResponsePaginationTransfer) *response.ApiResponsePaginationTransfer

	// Converts paginated soft-deleted transfer records into an API response.
	ToApiResponsePaginationTransferDeleteAt(pbResponse *pb.ApiResponsePaginationTransferDeleteAt) *response.ApiResponsePaginationTransferDeleteAt
}

// UserResponseMapper defines methods for mapping gRPC user-related responses
// into HTTP/REST API responses.
type UserResponseMapper interface {

	// Converts a single user response into an API response.
	ToApiResponseUser(pbResponse *pb.ApiResponseUser) *response.ApiResponseUser

	// Converts a soft-deleted user response into an API response.
	ToApiResponseUserDeleteAt(pbResponse *pb.ApiResponseUserDeleteAt) *response.ApiResponseUserDeleteAt

	// Converts multiple users into a grouped API response.
	ToApiResponsesUser(pbResponse *pb.ApiResponsesUser) *response.ApiResponsesUser

	// Converts a permanently deleted user response into an API response.
	ToApiResponseUserDelete(pbResponse *pb.ApiResponseUserDelete) *response.ApiResponseUserDelete

	// Converts all user records into an API response.
	ToApiResponseUserAll(pbResponse *pb.ApiResponseUserAll) *response.ApiResponseUserAll

	// Converts paginated soft-deleted users into an API response.
	ToApiResponsePaginationUserDeleteAt(pbResponse *pb.ApiResponsePaginationUserDeleteAt) *response.ApiResponsePaginationUserDeleteAt

	// Converts paginated user records into an API response.
	ToApiResponsePaginationUser(pbResponse *pb.ApiResponsePaginationUser) *response.ApiResponsePaginationUser
}

// WithdrawResponseMapper defines methods for converting gRPC withdraw-related responses
// into HTTP/REST API responses.
type WithdrawResponseMapper interface {

	// Converts a single withdraw response into an API response.
	ToApiResponseWithdraw(pbResponse *pb.ApiResponseWithdraw) *response.ApiResponseWithdraw

	// Converts a list of withdraw responses into a grouped API response.
	ToApiResponsesWithdraw(pbResponse *pb.ApiResponsesWithdraw) *response.ApiResponsesWithdraw

	// Converts a permanently deleted withdraw response into an API response.
	ToApiResponseWithdrawDelete(pbResponse *pb.ApiResponseWithdrawDelete) *response.ApiResponseWithdrawDelete

	// Converts all withdraw records into an API response.
	ToApiResponseWithdrawAll(pbResponse *pb.ApiResponseWithdrawAll) *response.ApiResponseWithdrawAll

	// Converts paginated withdraw records into an API response.
	ToApiResponsePaginationWithdraw(pbResponse *pb.ApiResponsePaginationWithdraw) *response.ApiResponsePaginationWithdraw

	// Converts paginated soft-deleted withdraw records into an API response.
	ToApiResponsePaginationWithdrawDeleteAt(pbResponse *pb.ApiResponsePaginationWithdrawDeleteAt) *response.ApiResponsePaginationWithdrawDeleteAt

	// Converts monthly successful withdraw statistics into an API response.
	ToApiResponseWithdrawMonthStatusSuccess(pbResponse *pb.ApiResponseWithdrawMonthStatusSuccess) *response.ApiResponseWithdrawMonthStatusSuccess

	// Converts yearly successful withdraw statistics into an API response.
	ToApiResponseWithdrawYearStatusSuccess(pbResponse *pb.ApiResponseWithdrawYearStatusSuccess) *response.ApiResponseWithdrawYearStatusSuccess

	// Converts monthly failed withdraw statistics into an API response.
	ToApiResponseWithdrawMonthStatusFailed(pbResponse *pb.ApiResponseWithdrawMonthStatusFailed) *response.ApiResponseWithdrawMonthStatusFailed

	// Converts yearly failed withdraw statistics into an API response.
	ToApiResponseWithdrawYearStatusFailed(pbResponse *pb.ApiResponseWithdrawYearStatusFailed) *response.ApiResponseWithdrawYearStatusFailed

	// Converts monthly total withdraw amount statistics into an API response.
	ToApiResponseWithdrawMonthAmount(pbResponse *pb.ApiResponseWithdrawMonthAmount) *response.ApiResponseWithdrawMonthAmount

	// Converts yearly total withdraw amount statistics into an API response.
	ToApiResponseWithdrawYearAmount(pbResponse *pb.ApiResponseWithdrawYearAmount) *response.ApiResponseWithdrawYearAmount
}
