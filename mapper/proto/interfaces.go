package protomapper

import (
	pb "github.com/MamangRust/monolith-payment-gateway-pb"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/response"
)

// AuthProtoMapper defines a set of methods to map authentication-related responses
// to their corresponding gRPC protobuf message formats.
//
//go:generate mockgen -source=interfaces.go -destination=mocks/mocks.go
type AuthProtoMapper interface {
	// ToProtoResponseVerifyCode converts a verification code response to a protobuf message.
	ToProtoResponseVerifyCode(status string, message string) *pb.ApiResponseVerifyCode

	// ToProtoResponseForgotPassword converts a forgot password response to a protobuf message.
	ToProtoResponseForgotPassword(status string, message string) *pb.ApiResponseForgotPassword

	// ToProtoResponseResetPassword converts a reset password response to a protobuf message.
	ToProtoResponseResetPassword(status string, message string) *pb.ApiResponseResetPassword

	// ToProtoResponseLogin converts a login response with token details to a protobuf message.
	ToProtoResponseLogin(status string, message string, response *response.TokenResponse) *pb.ApiResponseLogin

	// ToProtoResponseRegister converts a user registration response to a protobuf message.
	ToProtoResponseRegister(status string, message string, response *response.UserResponse) *pb.ApiResponseRegister

	// ToProtoResponseRefreshToken converts a token refresh response to a protobuf message.
	ToProtoResponseRefreshToken(status string, message string, response *response.TokenResponse) *pb.ApiResponseRefreshToken

	// ToProtoResponseGetMe converts a user profile fetch ("get me") response to a protobuf message.
	ToProtoResponseGetMe(status string, message string, response *response.UserResponse) *pb.ApiResponseGetMe
}

// UserProtoMapper defines methods for converting user-related data structures
// into gRPC protobuf response messages.
type UserProtoMapper interface {
	// ToProtoResponsesUser converts a list of user responses to a protobuf message.
	ToProtoResponsesUser(status string, message string, pbResponse []*response.UserResponse) *pb.ApiResponsesUser

	// ToProtoResponseUser converts a single user response to a protobuf message.
	ToProtoResponseUser(status string, message string, pbResponse *response.UserResponse) *pb.ApiResponseUser

	// ToProtoResponseUserDeleteAt converts a soft-deleted user response to a protobuf message.
	ToProtoResponseUserDeleteAt(status string, message string, pbResponse *response.UserResponseDeleteAt) *pb.ApiResponseUserDeleteAt

	// ToProtoResponseUserDelete returns a protobuf message indicating a user has been permanently deleted.
	ToProtoResponseUserDelete(status string, message string) *pb.ApiResponseUserDelete

	// ToProtoResponseUserAll returns a protobuf message for all users without pagination.
	ToProtoResponseUserAll(status string, message string) *pb.ApiResponseUserAll

	// ToProtoResponsePaginationUserDeleteAt returns a paginated protobuf message for soft-deleted users.
	ToProtoResponsePaginationUserDeleteAt(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponseDeleteAt) *pb.ApiResponsePaginationUserDeleteAt

	// ToProtoResponsePaginationUser returns a paginated protobuf message for active users.
	ToProtoResponsePaginationUser(pagination *pb.PaginationMeta, status string, message string, users []*response.UserResponse) *pb.ApiResponsePaginationUser
}

// RoleProtoMapper defines methods to map role-related data structures
// into gRPC protobuf response messages.
type RoleProtoMapper interface {
	// ToProtoResponseRoleAll returns a protobuf message for all roles without pagination.
	ToProtoResponseRoleAll(status string, message string) *pb.ApiResponseRoleAll

	// ToProtoResponseRoleDelete returns a protobuf message indicating a role has been deleted.
	ToProtoResponseRoleDelete(status string, message string) *pb.ApiResponseRoleDelete

	// ToProtoResponseRole converts a single role response to a protobuf message.
	ToProtoResponseRole(status string, message string, pbResponse *response.RoleResponse) *pb.ApiResponseRole

	// ToProtoResponsesRole converts a list of role responses to a protobuf message.
	ToProtoResponsesRole(status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsesRole

	// ToProtoResponsePaginationRole returns a paginated protobuf message for roles.
	ToProtoResponsePaginationRole(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.RoleResponse) *pb.ApiResponsePaginationRole

	// ToProtoResponsePaginationRoleDeleteAt returns a paginated protobuf message for soft-deleted roles.
	ToProtoResponsePaginationRoleDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.RoleResponseDeleteAt) *pb.ApiResponsePaginationRoleDeleteAt
}

// CardProtoMapper defines methods for mapping card-related domain responses
// to gRPC protobuf response messages.
type CardProtoMapper interface {
	// ToProtoResponseCard maps a single card response to a protobuf message.
	ToProtoResponseCard(status string, message string, card *response.CardResponse) *pb.ApiResponseCard

	// ToProtoResponsePaginationCard maps a paginated list of cards to a protobuf message.
	ToProtoResponsePaginationCard(pagination *pb.PaginationMeta, status string, message string, cards []*response.CardResponse) *pb.ApiResponsePaginationCard

	// ToProtoResponseCardDeleteAt returns a protobuf message indicating a card has been soft deleted.
	ToProtoResponseCardDeleteAt(status string, message string) *pb.ApiResponseCardDelete

	// ToProtoResponseCardAll returns all cards without pagination in a protobuf message.
	ToProtoResponseCardAll(status string, message string) *pb.ApiResponseCardAll

	// ToProtoResponsePaginationCardDeletedAt maps paginated soft-deleted cards to a protobuf message.
	ToProtoResponsePaginationCardDeletedAt(pagination *pb.PaginationMeta, status string, message string, cards []*response.CardResponseDeleteAt) *pb.ApiResponsePaginationCardDeleteAt

	// ToProtoResponseDashboardCard maps dashboard statistics for cards to a protobuf message.
	ToProtoResponseDashboardCard(status string, message string, dash *response.DashboardCard) *pb.ApiResponseDashboardCard

	// ToProtoResponseDashboardCardCardNumber maps card number statistics to a protobuf message.
	ToProtoResponseDashboardCardCardNumber(status string, message string, dash *response.DashboardCardCardNumber) *pb.ApiResponseDashboardCardNumber

	// ToProtoResponseMonthlyBalances maps monthly card balances to a protobuf message.
	ToProtoResponseMonthlyBalances(status string, message string, cards []*response.CardResponseMonthBalance) *pb.ApiResponseMonthlyBalance

	// ToProtoResponseYearlyBalances maps yearly card balances to a protobuf message.
	ToProtoResponseYearlyBalances(status string, message string, cards []*response.CardResponseYearlyBalance) *pb.ApiResponseYearlyBalance

	// ToProtoResponseMonthlyAmounts maps monthly card amounts to a protobuf message.
	ToProtoResponseMonthlyAmounts(status string, message string, cards []*response.CardResponseMonthAmount) *pb.ApiResponseMonthlyAmount

	// ToProtoResponseYearlyAmounts maps yearly card amounts to a protobuf message.
	ToProtoResponseYearlyAmounts(status string, message string, cards []*response.CardResponseYearAmount) *pb.ApiResponseYearlyAmount
}

// MerchantProtoMapper defines methods for mapping merchant-related domain responses
// to gRPC protobuf response messages.
type MerchantProtoMapper interface {
	// ToProtoResponsePaginationMerchant maps paginated merchant responses to a protobuf message.
	ToProtoResponsePaginationMerchant(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantResponse) *pb.ApiResponsePaginationMerchant

	// ToProtoResponseMerchants maps a list of merchant responses to a protobuf message.
	ToProtoResponseMerchants(status string, message string, res []*response.MerchantResponse) *pb.ApiResponsesMerchant

	// ToProtoResponseMerchant maps a single merchant response to a protobuf message.
	ToProtoResponseMerchant(status string, message string, res *response.MerchantResponse) *pb.ApiResponseMerchant

	// ToProtoResponseMerchantAll maps all merchant responses (non-paginated) to a protobuf message.
	ToProtoResponseMerchantAll(status string, message string) *pb.ApiResponseMerchantAll

	// ToProtoResponseMerchantDelete returns a protobuf message indicating a merchant has been deleted.
	ToProtoResponseMerchantDelete(status string, message string) *pb.ApiResponseMerchantDelete

	// ToProtoResponsePaginationMerchantDeleteAt maps paginated soft-deleted merchants to a protobuf message.
	ToProtoResponsePaginationMerchantDeleteAt(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantResponseDeleteAt) *pb.ApiResponsePaginationMerchantDeleteAt

	// ToProtoResponsePaginationMerchantTransaction maps paginated merchant transaction summaries to a protobuf message.
	ToProtoResponsePaginationMerchantTransaction(pagination *pb.PaginationMeta, status string, message string, merchants []*response.MerchantTransactionResponse) *pb.ApiResponsePaginationMerchantTransaction

	// ToProtoResponseMonthlyPaymentMethods maps monthly payment method statistics to a protobuf message.
	ToProtoResponseMonthlyPaymentMethods(status string, message string, ms []*response.MerchantResponseMonthlyPaymentMethod) *pb.ApiResponseMerchantMonthlyPaymentMethod

	// ToProtoResponseYearlyPaymentMethods maps yearly payment method statistics to a protobuf message.
	ToProtoResponseYearlyPaymentMethods(status string, message string, ms []*response.MerchantResponseYearlyPaymentMethod) *pb.ApiResponseMerchantYearlyPaymentMethod

	// ToProtoResponseMonthlyAmounts maps monthly merchant transaction amounts to a protobuf message.
	ToProtoResponseMonthlyAmounts(status string, message string, ms []*response.MerchantResponseMonthlyAmount) *pb.ApiResponseMerchantMonthlyAmount

	// ToProtoResponseYearlyAmounts maps yearly merchant transaction amounts to a protobuf message.
	ToProtoResponseYearlyAmounts(status string, message string, ms []*response.MerchantResponseYearlyAmount) *pb.ApiResponseMerchantYearlyAmount

	// ToProtoResponseMonthlyTotalAmounts maps monthly total transaction amounts across all merchants to a protobuf message.
	ToProtoResponseMonthlyTotalAmounts(status string, message string, ms []*response.MerchantResponseMonthlyTotalAmount) *pb.ApiResponseMerchantMonthlyTotalAmount

	// ToProtoResponseYearlyTotalAmounts maps yearly total transaction amounts across all merchants to a protobuf message.
	ToProtoResponseYearlyTotalAmounts(status string, message string, ms []*response.MerchantResponseYearlyTotalAmount) *pb.ApiResponseMerchantYearlyTotalAmount
}

// MerchantDocumentProtoMapper defines methods for mapping merchant document-related responses
// into gRPC protobuf response messages.
type MerchantDocumentProtoMapper interface {
	// ToProtoResponseMerchantDocument maps a single merchant document to a protobuf response.
	ToProtoResponseMerchantDocument(status string, message string, doc *response.MerchantDocumentResponse) *pb.ApiResponseMerchantDocument

	// ToProtoResponsesMerchantDocument maps a list of merchant documents to a protobuf response.
	ToProtoResponsesMerchantDocument(status string, message string, docs []*response.MerchantDocumentResponse) *pb.ApiResponsesMerchantDocument

	// ToProtoResponsePaginationMerchantDocument maps paginated merchant documents to a protobuf response.
	ToProtoResponsePaginationMerchantDocument(pagination *pb.PaginationMeta, status string, message string, docs []*response.MerchantDocumentResponse) *pb.ApiResponsePaginationMerchantDocument

	// ToProtoResponsePaginationMerchantDocumentDeleteAt maps paginated soft-deleted merchant documents to a protobuf response.
	ToProtoResponsePaginationMerchantDocumentDeleteAt(pagination *pb.PaginationMeta, status string, message string, docs []*response.MerchantDocumentResponseDeleteAt) *pb.ApiResponsePaginationMerchantDocumentAt

	// ToProtoResponseMerchantDocumentDelete returns a response indicating a merchant document has been deleted.
	ToProtoResponseMerchantDocumentDelete(status string, message string) *pb.ApiResponseMerchantDocumentDelete

	// ToProtoResponseMerchantDocumentAll maps all merchant documents to a protobuf response.
	ToProtoResponseMerchantDocumentAll(status string, message string) *pb.ApiResponseMerchantDocumentAll
}

// SaldoProtoMapper defines methods for mapping saldo (balance)-related responses
// into gRPC protobuf response messages.
type SaldoProtoMapper interface {
	// ToProtoResponseSaldo maps a single saldo record to a protobuf response.
	ToProtoResponseSaldo(status string, message string, pbResponse *response.SaldoResponse) *pb.ApiResponseSaldo

	// ToProtoResponsesSaldo maps a list of saldo records to a protobuf response.
	ToProtoResponsesSaldo(status string, message string, pbResponse []*response.SaldoResponse) *pb.ApiResponsesSaldo

	// ToProtoResponseSaldoDelete returns a response indicating a saldo record has been deleted.
	ToProtoResponseSaldoDelete(status string, message string) *pb.ApiResponseSaldoDelete

	// ToProtoResponseSaldoAll returns all saldo records without pagination.
	ToProtoResponseSaldoAll(status string, message string) *pb.ApiResponseSaldoAll

	// ToProtoResponseMonthTotalSaldo maps monthly total saldo to a protobuf response.
	ToProtoResponseMonthTotalSaldo(status string, message string, pbResponse []*response.SaldoMonthTotalBalanceResponse) *pb.ApiResponseMonthTotalSaldo

	// ToProtoResponseYearTotalSaldo maps yearly total saldo to a protobuf response.
	ToProtoResponseYearTotalSaldo(status string, message string, pbResponse []*response.SaldoYearTotalBalanceResponse) *pb.ApiResponseYearTotalSaldo

	// ToProtoResponseMonthSaldoBalances maps monthly saldo balances to a protobuf response.
	ToProtoResponseMonthSaldoBalances(status string, message string, pbResponse []*response.SaldoMonthBalanceResponse) *pb.ApiResponseMonthSaldoBalances

	// ToProtoResponseYearSaldoBalances maps yearly saldo balances to a protobuf response.
	ToProtoResponseYearSaldoBalances(status string, message string, pbResponse []*response.SaldoYearBalanceResponse) *pb.ApiResponseYearSaldoBalances

	// ToProtoResponsePaginationSaldo maps paginated saldo records to a protobuf response.
	ToProtoResponsePaginationSaldo(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponse) *pb.ApiResponsePaginationSaldo

	// ToProtoResponsePaginationSaldoDeleteAt maps paginated soft-deleted saldo records to a protobuf response.
	ToProtoResponsePaginationSaldoDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.SaldoResponseDeleteAt) *pb.ApiResponsePaginationSaldoDeleteAt
}

// TopupProtoMapper defines methods for mapping top-up-related domain responses
// into gRPC protobuf response messages.
type TopupProtoMapper interface {
	// ToProtoResponseTopup maps a single top-up record to a protobuf response.
	ToProtoResponseTopup(status string, message string, s *response.TopupResponse) *pb.ApiResponseTopup

	// ToProtoResponseTopupDeletAt maps a soft-deleted top-up record to a protobuf response.
	ToProtoResponseTopupDeletAt(status string, message string, s *response.TopupResponseDeleteAt) *pb.ApiResponseTopupDeleteAt

	// ToProtoResponseTopupDelete returns a response indicating a top-up record has been deleted.
	ToProtoResponseTopupDelete(status string, message string) *pb.ApiResponseTopupDelete

	// ToProtoResponseTopupAll returns all top-up records in a protobuf response.
	ToProtoResponseTopupAll(status string, message string) *pb.ApiResponseTopupAll

	// ToProtoResponsePaginationTopup maps paginated top-up records to a protobuf response.
	ToProtoResponsePaginationTopup(pagination *pb.PaginationMeta, status string, message string, s []*response.TopupResponse) *pb.ApiResponsePaginationTopup

	// ToProtoResponsePaginationTopupDeleteAt maps paginated soft-deleted top-up records to a protobuf response.
	ToProtoResponsePaginationTopupDeleteAt(pagination *pb.PaginationMeta, status string, message string, s []*response.TopupResponseDeleteAt) *pb.ApiResponsePaginationTopupDeleteAt

	// ToProtoResponseTopupMonthStatusSuccess maps monthly successful top-ups to a protobuf response.
	ToProtoResponseTopupMonthStatusSuccess(status string, message string, s []*response.TopupResponseMonthStatusSuccess) *pb.ApiResponseTopupMonthStatusSuccess

	// ToProtoResponseTopupYearStatusSuccess maps yearly successful top-ups to a protobuf response.
	ToProtoResponseTopupYearStatusSuccess(status string, message string, s []*response.TopupResponseYearStatusSuccess) *pb.ApiResponseTopupYearStatusSuccess

	// ToProtoResponseTopupMonthStatusFailed maps monthly failed top-ups to a protobuf response.
	ToProtoResponseTopupMonthStatusFailed(status string, message string, s []*response.TopupResponseMonthStatusFailed) *pb.ApiResponseTopupMonthStatusFailed

	// ToProtoResponseTopupYearStatusFailed maps yearly failed top-ups to a protobuf response.
	ToProtoResponseTopupYearStatusFailed(status string, message string, s []*response.TopupResponseYearStatusFailed) *pb.ApiResponseTopupYearStatusFailed

	// ToProtoResponseTopupMonthMethod maps monthly top-up methods to a protobuf response.
	ToProtoResponseTopupMonthMethod(status string, message string, s []*response.TopupMonthMethodResponse) *pb.ApiResponseTopupMonthMethod

	// ToProtoResponseTopupYearMethod maps yearly top-up methods to a protobuf response.
	ToProtoResponseTopupYearMethod(status string, message string, s []*response.TopupYearlyMethodResponse) *pb.ApiResponseTopupYearMethod

	// ToProtoResponseTopupMonthAmount maps monthly top-up amounts to a protobuf response.
	ToProtoResponseTopupMonthAmount(status string, message string, s []*response.TopupMonthAmountResponse) *pb.ApiResponseTopupMonthAmount

	// ToProtoResponseTopupYearAmount maps yearly top-up amounts to a protobuf response.
	ToProtoResponseTopupYearAmount(status string, message string, s []*response.TopupYearlyAmountResponse) *pb.ApiResponseTopupYearAmount
}

// TransactionProtoMapper defines methods for mapping transaction-related domain responses
// into gRPC protobuf response messages.
type TransactionProtoMapper interface {
	// ToProtoResponseTransactionMonthStatusSuccess maps monthly successful transactions to a protobuf response.
	ToProtoResponseTransactionMonthStatusSuccess(status string, message string, pbResponse []*response.TransactionResponseMonthStatusSuccess) *pb.ApiResponseTransactionMonthStatusSuccess

	// ToProtoResponseTransactionYearStatusSuccess maps yearly successful transactions to a protobuf response.
	ToProtoResponseTransactionYearStatusSuccess(status string, message string, pbResponse []*response.TransactionResponseYearStatusSuccess) *pb.ApiResponseTransactionYearStatusSuccess

	// ToProtoResponseTransactionMonthStatusFailed maps monthly failed transactions to a protobuf response.
	ToProtoResponseTransactionMonthStatusFailed(status string, message string, pbResponse []*response.TransactionResponseMonthStatusFailed) *pb.ApiResponseTransactionMonthStatusFailed

	// ToProtoResponseTransactionYearStatusFailed maps yearly failed transactions to a protobuf response.
	ToProtoResponseTransactionYearStatusFailed(status string, message string, pbResponse []*response.TransactionResponseYearStatusFailed) *pb.ApiResponseTransactionYearStatusFailed

	// ToProtoResponseTransactionMonthMethod maps monthly transaction methods to a protobuf response.
	ToProtoResponseTransactionMonthMethod(status string, message string, pbResponse []*response.TransactionMonthMethodResponse) *pb.ApiResponseTransactionMonthMethod

	// ToProtoResponseTransactionYearMethod maps yearly transaction methods to a protobuf response.
	ToProtoResponseTransactionYearMethod(status string, message string, pbResponse []*response.TransactionYearMethodResponse) *pb.ApiResponseTransactionYearMethod

	// ToProtoResponseTransactionMonthAmount maps monthly transaction amounts to a protobuf response.
	ToProtoResponseTransactionMonthAmount(status string, message string, pbResponse []*response.TransactionMonthAmountResponse) *pb.ApiResponseTransactionMonthAmount

	// ToProtoResponseTransactionYearAmount maps yearly transaction amounts to a protobuf response.
	ToProtoResponseTransactionYearAmount(status string, message string, pbResponse []*response.TransactionYearlyAmountResponse) *pb.ApiResponseTransactionYearAmount

	// ToProtoResponseTransaction maps a single transaction to a protobuf response.
	ToProtoResponseTransaction(status string, message string, pbResponse *response.TransactionResponse) *pb.ApiResponseTransaction

	// ToProtoResponseTransactions maps a list of transactions to a protobuf response.
	ToProtoResponseTransactions(status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponseTransactions

	// ToProtoResponseTransactionDelete returns a response indicating a transaction has been deleted.
	ToProtoResponseTransactionDelete(status string, message string) *pb.ApiResponseTransactionDelete

	// ToProtoResponseTransactionAll returns all transactions in a protobuf response.
	ToProtoResponseTransactionAll(status string, message string) *pb.ApiResponseTransactionAll

	// ToProtoResponsePaginationTransaction maps paginated transaction records to a protobuf response.
	ToProtoResponsePaginationTransaction(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponse) *pb.ApiResponsePaginationTransaction

	// ToProtoResponsePaginationTransactionDeleteAt maps paginated soft-deleted transactions to a protobuf response.
	ToProtoResponsePaginationTransactionDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransactionResponseDeleteAt) *pb.ApiResponsePaginationTransactionDeleteAt
}

type TransferProtoMapper interface {
	// ToProtoResponseTransferMonthStatusSuccess maps monthly successful transfers to a protobuf response.
	ToProtoResponseTransferMonthStatusSuccess(status string, message string, pbResponse []*response.TransferResponseMonthStatusSuccess) *pb.ApiResponseTransferMonthStatusSuccess
	// ToProtoResponseTransferYearStatusSuccess maps yearly successful transfers to a protobuf response.
	ToProtoResponseTransferYearStatusSuccess(status string, message string, pbResponse []*response.TransferResponseYearStatusSuccess) *pb.ApiResponseTransferYearStatusSuccess
	// ToProtoResponseTransferMonthStatusFailed maps monthly failed transfers to a protobuf response.
	ToProtoResponseTransferMonthStatusFailed(status string, message string, pbResponse []*response.TransferResponseMonthStatusFailed) *pb.ApiResponseTransferMonthStatusFailed
	// ToProtoResponseTransferYearStatusFailed maps yearly failed transfers to a protobuf response.
	ToProtoResponseTransferYearStatusFailed(status string, message string, pbResponse []*response.TransferResponseYearStatusFailed) *pb.ApiResponseTransferYearStatusFailed
	// ToProtoResponseTransferMonthAmount maps monthly transfer amounts to a protobuf response.
	ToProtoResponseTransferMonthAmount(status string, message string, pbResponse []*response.TransferMonthAmountResponse) *pb.ApiResponseTransferMonthAmount
	// ToProtoResponseTransferYearAmount maps yearly transfer amounts to a protobuf response.
	ToProtoResponseTransferYearAmount(status string, message string, pbResponse []*response.TransferYearAmountResponse) *pb.ApiResponseTransferYearAmount
	// ToProtoResponseTransfer maps a single transfer to a protobuf response.
	ToProtoResponseTransfer(status string, message string, pbResponse *response.TransferResponse) *pb.ApiResponseTransfer
	// ToProtoResponseTransfers maps a list of transfers to a protobuf response.
	ToProtoResponseTransfers(status string, message string, pbResponse []*response.TransferResponse) *pb.ApiResponseTransfers
	// ToProtoResponseTransferDelete returns a response indicating a transfer has been deleted.
	ToProtoResponseTransferDelete(status string, message string) *pb.ApiResponseTransferDelete
	// ToProtoResponseTransferAll returns all transfers in a protobuf response.
	ToProtoResponseTransferAll(status string, message string) *pb.ApiResponseTransferAll
	// ToProtoResponsePaginationTransfer maps paginated transfer records to a protobuf response.
	ToProtoResponsePaginationTransfer(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransferResponse) *pb.ApiResponsePaginationTransfer
	// ToProtoResponsePaginationTransferDeleteAt maps paginated soft-deleted transfers to a protobuf response.
	ToProtoResponsePaginationTransferDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.TransferResponseDeleteAt) *pb.ApiResponsePaginationTransferDeleteAt
}

type WithdrawalProtoMapper interface {
	// ToProtoResponseWithdraw maps a single withdraw to a protobuf response.
	ToProtoResponseWithdraw(status string, message string, withdraw *response.WithdrawResponse) *pb.ApiResponseWithdraw

	// ToProtoResponsesWithdraw maps a list of withdraws to a protobuf response.
	ToProtoResponsesWithdraw(status string, message string, pbResponse []*response.WithdrawResponse) *pb.ApiResponsesWithdraw

	// ToProtoResponseWithdrawDelete returns a response indicating a withdraw has been deleted.
	ToProtoResponseWithdrawDelete(status string, message string) *pb.ApiResponseWithdrawDelete

	// ToProtoResponseWithdrawAll returns all withdraws in a protobuf response.
	ToProtoResponseWithdrawAll(status string, message string) *pb.ApiResponseWithdrawAll

	// ToProtoResponsePaginationWithdraw maps paginated withdraw records to a protobuf response.
	ToProtoResponsePaginationWithdraw(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponse) *pb.ApiResponsePaginationWithdraw

	// ToProtoResponsePaginationWithdrawDeleteAt maps paginated soft-deleted withdraw records to a protobuf response.
	ToProtoResponsePaginationWithdrawDeleteAt(pagination *pb.PaginationMeta, status string, message string, pbResponse []*response.WithdrawResponseDeleteAt) *pb.ApiResponsePaginationWithdrawDeleteAt

	// ToProtoResponseWithdrawMonthStatusSuccess maps monthly successful withdraws to a protobuf response.
	ToProtoResponseWithdrawMonthStatusSuccess(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusSuccess) *pb.ApiResponseWithdrawMonthStatusSuccess

	// ToProtoResponseWithdrawYearStatusSuccess maps yearly successful withdraws to a protobuf response.
	ToProtoResponseWithdrawYearStatusSuccess(status string, message string, pbResponse []*response.WithdrawResponseYearStatusSuccess) *pb.ApiResponseWithdrawYearStatusSuccess

	// ToProtoResponseWithdrawMonthStatusFailed maps monthly failed withdraws to a protobuf response.
	ToProtoResponseWithdrawMonthStatusFailed(status string, message string, pbResponse []*response.WithdrawResponseMonthStatusFailed) *pb.ApiResponseWithdrawMonthStatusFailed

	// ToProtoResponseWithdrawYearStatusFailed maps yearly failed withdraws to a protobuf response.
	ToProtoResponseWithdrawYearStatusFailed(status string, message string, pbResponse []*response.WithdrawResponseYearStatusFailed) *pb.ApiResponseWithdrawYearStatusFailed

	// ToProtoResponseWithdrawMonthAmount maps monthly withdraw amounts to a protobuf response.
	ToProtoResponseWithdrawMonthAmount(status string, message string, pbResponse []*response.WithdrawMonthlyAmountResponse) *pb.ApiResponseWithdrawMonthAmount

	// ToProtoResponseWithdrawYearAmount maps yearly withdraw amounts to a protobuf response.
	ToProtoResponseWithdrawYearAmount(status string, message string, pbResponse []*response.WithdrawYearlyAmountResponse) *pb.ApiResponseWithdrawYearAmount
}
