package recordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

type UserRecordMapping interface {
	// ToUserRecord maps a User database row to a UserRecord domain model.
	ToUserRecord(user *db.User) *record.UserRecord

	// ToUserRecordPagination maps a GetUsersWithPaginationRow to a UserRecord domain model.
	ToUserRecordPagination(user *db.GetUsersWithPaginationRow) *record.UserRecord

	// ToUsersRecordPagination maps a slice of GetUsersWithPaginationRow to a slice of UserRecord domain models.
	ToUsersRecordPagination(users []*db.GetUsersWithPaginationRow) []*record.UserRecord

	// ToUserRecordActivePagination maps a GetActiveUsersWithPaginationRow to a UserRecord domain model.
	ToUserRecordActivePagination(user *db.GetActiveUsersWithPaginationRow) *record.UserRecord

	// ToUsersRecordActivePagination maps a slice of GetActiveUsersWithPaginationRow to a slice of UserRecord domain models.
	ToUsersRecordActivePagination(users []*db.GetActiveUsersWithPaginationRow) []*record.UserRecord

	// ToUserRecordTrashedPagination maps a GetTrashedUsersWithPaginationRow to a UserRecord domain model.
	ToUserRecordTrashedPagination(user *db.GetTrashedUsersWithPaginationRow) *record.UserRecord

	// ToUsersRecordTrashedPagination maps a slice of GetTrashedUsersWithPaginationRow to a slice of UserRecord domain models.
	ToUsersRecordTrashedPagination(users []*db.GetTrashedUsersWithPaginationRow) []*record.UserRecord
}

// RoleRecordMapping defines a mapping function from a Role database row to a RoleRecord domain model.
type RoleRecordMapping interface {
	// ToRoleRecord maps a Role database row to a RoleRecord domain model.
	ToRoleRecord(role *db.Role) *record.RoleRecord

	// ToRolesRecord maps a slice of Role database rows to a slice of RoleRecord domain models.
	ToRolesRecord(roles []*db.Role) []*record.RoleRecord

	// ToRoleRecordAll maps a GetRolesRow to a RoleRecord domain model.
	ToRoleRecordAll(role *db.GetRolesRow) *record.RoleRecord

	// ToRolesRecordAll maps a slice of GetRolesRow to a slice of RoleRecord domain models.
	ToRolesRecordAll(roles []*db.GetRolesRow) []*record.RoleRecord

	// ToRoleRecordActive maps a GetActiveRolesRow to a RoleRecord domain model.
	ToRoleRecordActive(role *db.GetActiveRolesRow) *record.RoleRecord

	// ToRolesRecordActive maps a slice of GetActiveRolesRow to a slice of RoleRecord domain models.
	ToRolesRecordActive(roles []*db.GetActiveRolesRow) []*record.RoleRecord

	// ToRoleRecordTrashed maps a GetTrashedRolesRow to a RoleRecord domain model.
	ToRoleRecordTrashed(role *db.GetTrashedRolesRow) *record.RoleRecord

	// ToRolesRecordTrashed maps a slice of GetTrashedRolesRow to a slice of RoleRecord domain models.
	ToRolesRecordTrashed(roles []*db.GetTrashedRolesRow) []*record.RoleRecord
}

// UserRoleRecordMapping defines a mapping function from a UserRole database row to a UserRoleRecord domain model.
type UserRoleRecordMapping interface {
	// ToUserRoleRecord maps a UserRole database row to a UserRoleRecord domain model.
	ToUserRoleRecord(userRole *db.UserRole) *record.UserRoleRecord
}

// RefreshTokenRecordMapping defines a mapping function from a RefreshToken database row to a RefreshTokenRecord domain model.
type RefreshTokenRecordMapping interface {
	// ToRefreshTokenRecord maps a RefreshToken database row to a RefreshTokenRecord domain model.
	ToRefreshTokenRecord(refreshToken *db.RefreshToken) *record.RefreshTokenRecord

	// ToRefreshTokensRecord maps a slice of RefreshToken database rows to a slice of RefreshTokenRecord domain models.
	ToRefreshTokensRecord(refreshTokens []*db.RefreshToken) []*record.RefreshTokenRecord
}

// ResetTokenRecordMapping defines a mapping function from a ResetToken database row to a ResetTokenRecord domain model.
type ResetTokenRecordMapping interface {
	// ToResetTokenRecord maps a ResetToken database row to a ResetTokenRecord domain model.
	ToResetTokenRecord(resetToken *db.ResetToken) *record.ResetTokenRecord
}

// SaldoRecordMapping provides methods to map database rows to Saldo domain models.
type SaldoRecordMapping interface {
	// ToSaldoRecord maps a Saldo database row to a SaldoRecord domain model.
	ToSaldoRecord(saldo *db.Saldo) *record.SaldoRecord

	// ToSaldosRecord maps a slice of Saldo database rows to a slice of SaldoRecord domain models.
	ToSaldosRecord(saldos []*db.Saldo) []*record.SaldoRecord

	// ToSaldoMonthTotalBalance maps a MonthlyTotalSaldoBalanceRow to a SaldoMonthTotalBalance domain model.
	ToSaldoMonthTotalBalance(ss *db.GetMonthlyTotalSaldoBalanceRow) *record.SaldoMonthTotalBalance

	// ToSaldoMonthTotalBalances maps a slice of MonthlyTotalSaldoBalanceRow to a slice of SaldoMonthTotalBalance domain models.
	ToSaldoMonthTotalBalances(ss []*db.GetMonthlyTotalSaldoBalanceRow) []*record.SaldoMonthTotalBalance

	// ToSaldoYearTotalBalance maps a YearlyTotalSaldoBalancesRow to a SaldoYearTotalBalance domain model.
	ToSaldoYearTotalBalance(ss *db.GetYearlyTotalSaldoBalancesRow) *record.SaldoYearTotalBalance

	// ToSaldoYearTotalBalances maps a slice of YearlyTotalSaldoBalancesRow to a slice of SaldoYearTotalBalance domain models.
	ToSaldoYearTotalBalances(ss []*db.GetYearlyTotalSaldoBalancesRow) []*record.SaldoYearTotalBalance

	// ToSaldoMonthBalance maps a MonthlySaldoBalancesRow to a SaldoMonthSaldoBalance domain model.
	ToSaldoMonthBalance(ss *db.GetMonthlySaldoBalancesRow) *record.SaldoMonthSaldoBalance

	// ToSaldoMonthBalances maps a slice of MonthlySaldoBalancesRow to a slice of SaldoMonthSaldoBalance domain models.
	ToSaldoMonthBalances(ss []*db.GetMonthlySaldoBalancesRow) []*record.SaldoMonthSaldoBalance

	// ToSaldoYearSaldoBalance maps a YearlySaldoBalancesRow to a SaldoYearSaldoBalance domain model.
	ToSaldoYearSaldoBalance(ss *db.GetYearlySaldoBalancesRow) *record.SaldoYearSaldoBalance

	// ToSaldoYearSaldoBalances maps a slice of YearlySaldoBalancesRow to a slice of SaldoYearSaldoBalance domain models.
	ToSaldoYearSaldoBalances(ss []*db.GetYearlySaldoBalancesRow) []*record.SaldoYearSaldoBalance

	// ToSaldoRecordAll maps a SaldosRow to a SaldoRecord domain model.
	ToSaldoRecordAll(saldo *db.GetSaldosRow) *record.SaldoRecord

	// ToSaldosRecordAll maps a slice of SaldosRow to a slice of SaldoRecord domain models.
	ToSaldosRecordAll(saldos []*db.GetSaldosRow) []*record.SaldoRecord

	// ToSaldoRecordActive maps an ActiveSaldosRow to a SaldoRecord domain model.
	ToSaldoRecordActive(saldo *db.GetActiveSaldosRow) *record.SaldoRecord

	// ToSaldosRecordActive maps a slice of ActiveSaldosRow to a slice of SaldoRecord domain models.
	ToSaldosRecordActive(saldos []*db.GetActiveSaldosRow) []*record.SaldoRecord

	// ToSaldoRecordTrashed maps a TrashedSaldosRow to a SaldoRecord domain model.
	ToSaldoRecordTrashed(saldo *db.GetTrashedSaldosRow) *record.SaldoRecord

	// ToSaldosRecordTrashed maps a slice of TrashedSaldosRow to a slice of SaldoRecord domain models.
	ToSaldosRecordTrashed(saldos []*db.GetTrashedSaldosRow) []*record.SaldoRecord
}

// TopupRecordMapping provides methods to map database rows related to top-ups to domain models.
type TopupRecordMapping interface {
	// ToTopupRecord maps a Topup database row to a TopupRecord domain model.
	ToTopupRecord(topup *db.Topup) *record.TopupRecord

	// ToTopupRecords maps a slice of Topup database rows to a slice of TopupRecord domain models.
	ToTopupRecords(topups []*db.Topup) []*record.TopupRecord

	// ToTopupByCardNumberRecord maps a GetTopupsByCardNumberRow database row to a TopupRecord domain model.
	ToTopupByCardNumberRecord(topup *db.GetTopupsByCardNumberRow) *record.TopupRecord

	// ToTopupByCardNumberRecords maps a slice of GetTopupsByCardNumberRow database rows to a slice of TopupRecord domain models.
	ToTopupByCardNumberRecords(topups []*db.GetTopupsByCardNumberRow) []*record.TopupRecord

	// ToTopupRecordMonthStatusSuccess maps a GetMonthTopupStatusSuccessRow database row to a TopupRecordMonthStatusSuccess domain model.
	ToTopupRecordMonthStatusSuccess(s *db.GetMonthTopupStatusSuccessRow) *record.TopupRecordMonthStatusSuccess

	// ToTopupRecordsMonthStatusSuccess maps a slice of GetMonthTopupStatusSuccessRow database rows to a slice of TopupRecordMonthStatusSuccess domain models.
	ToTopupRecordsMonthStatusSuccess(topups []*db.GetMonthTopupStatusSuccessRow) []*record.TopupRecordMonthStatusSuccess

	// ToTopupRecordYearStatusSuccess maps a GetYearlyTopupStatusSuccessRow database row to a TopupRecordYearStatusSuccess domain model.
	ToTopupRecordYearStatusSuccess(s *db.GetYearlyTopupStatusSuccessRow) *record.TopupRecordYearStatusSuccess

	// ToTopupRecordsYearStatusSuccess maps a slice of GetYearlyTopupStatusSuccessRow database rows to a slice of TopupRecordYearStatusSuccess domain models.
	ToTopupRecordsYearStatusSuccess(topups []*db.GetYearlyTopupStatusSuccessRow) []*record.TopupRecordYearStatusSuccess

	// ToTopupRecordMonthStatusFailed maps a GetMonthTopupStatusFailedRow database row to a TopupRecordMonthStatusFailed domain model.
	ToTopupRecordMonthStatusFailed(s *db.GetMonthTopupStatusFailedRow) *record.TopupRecordMonthStatusFailed

	// ToTopupRecordsMonthStatusFailed maps a slice of GetMonthTopupStatusFailedRow database rows to a slice of TopupRecordMonthStatusFailed domain models.
	ToTopupRecordsMonthStatusFailed(topups []*db.GetMonthTopupStatusFailedRow) []*record.TopupRecordMonthStatusFailed

	// ToTopupRecordYearStatusFailed maps a GetYearlyTopupStatusFailedRow database row to a TopupRecordYearStatusFailed domain model.
	ToTopupRecordYearStatusFailed(s *db.GetYearlyTopupStatusFailedRow) *record.TopupRecordYearStatusFailed

	// ToTopupRecordsYearStatusFailed maps a slice of GetYearlyTopupStatusFailedRow database rows to a slice of TopupRecordYearStatusFailed domain models.
	ToTopupRecordsYearStatusFailed(topups []*db.GetYearlyTopupStatusFailedRow) []*record.TopupRecordYearStatusFailed

	// ToTopupRecordMonthStatusSuccessByCardNumber maps a GetMonthTopupStatusSuccessCardNumberRow database row to a TopupRecordMonthStatusSuccess domain model.
	ToTopupRecordMonthStatusSuccessByCardNumber(s *db.GetMonthTopupStatusSuccessCardNumberRow) *record.TopupRecordMonthStatusSuccess

	// ToTopupRecordsMonthStatusSuccessByCardNumber maps a slice of GetMonthTopupStatusSuccessCardNumberRow database rows to a slice of TopupRecordMonthStatusSuccess domain models.
	ToTopupRecordsMonthStatusSuccessByCardNumber(topups []*db.GetMonthTopupStatusSuccessCardNumberRow) []*record.TopupRecordMonthStatusSuccess

	// ToTopupRecordYearStatusSuccessByCardNumber maps a GetYearlyTopupStatusSuccessCardNumberRow database row to a TopupRecordYearStatusSuccess domain model.
	ToTopupRecordYearStatusSuccessByCardNumber(s *db.GetYearlyTopupStatusSuccessCardNumberRow) *record.TopupRecordYearStatusSuccess

	// ToTopupRecordsYearStatusSuccessByCardNumber maps a slice of GetYearlyTopupStatusSuccessCardNumberRow database rows to a slice of TopupRecordYearStatusSuccess domain models.
	ToTopupRecordsYearStatusSuccessByCardNumber(topups []*db.GetYearlyTopupStatusSuccessCardNumberRow) []*record.TopupRecordYearStatusSuccess

	// ToTopupRecordMonthStatusFailedByCardNumber maps a GetMonthTopupStatusFailedCardNumberRow database row to a TopupRecordMonthStatusFailed domain model.
	ToTopupRecordMonthStatusFailedByCardNumber(s *db.GetMonthTopupStatusFailedCardNumberRow) *record.TopupRecordMonthStatusFailed

	// ToTopupRecordsMonthStatusFailedByCardNumber maps a slice of GetMonthTopupStatusFailedCardNumberRow database rows to a slice of TopupRecordMonthStatusFailed domain models.
	ToTopupRecordsMonthStatusFailedByCardNumber(topups []*db.GetMonthTopupStatusFailedCardNumberRow) []*record.TopupRecordMonthStatusFailed

	// ToTopupRecordYearStatusFailedByCardNumber maps a GetYearlyTopupStatusFailedCardNumberRow database row to a TopupRecordYearStatusFailed domain model.
	ToTopupRecordYearStatusFailedByCardNumber(s *db.GetYearlyTopupStatusFailedCardNumberRow) *record.TopupRecordYearStatusFailed

	// ToTopupRecordsYearStatusFailedByCardNumber maps a slice of GetYearlyTopupStatusFailedCardNumberRow database rows to a slice of TopupRecordYearStatusFailed domain models.
	ToTopupRecordsYearStatusFailedByCardNumber(topups []*db.GetYearlyTopupStatusFailedCardNumberRow) []*record.TopupRecordYearStatusFailed

	// ToTopupMonthlyMethod maps a GetMonthlyTopupMethodsRow database row to a TopupMonthMethod domain model.
	ToTopupMonthlyMethod(s *db.GetMonthlyTopupMethodsRow) *record.TopupMonthMethod

	// ToTopupMonthlyMethods maps a slice of GetMonthlyTopupMethodsRow database rows to a slice of TopupMonthMethod domain models.
	ToTopupMonthlyMethods(s []*db.GetMonthlyTopupMethodsRow) []*record.TopupMonthMethod

	// ToTopupYearlyMethod maps a GetYearlyTopupMethodsRow database row to a TopupYearlyMethod domain model.
	ToTopupYearlyMethod(s *db.GetYearlyTopupMethodsRow) *record.TopupYearlyMethod

	// ToTopupYearlyMethods maps a slice of GetYearlyTopupMethodsRow database rows to a slice of TopupYearlyMethod domain models.
	ToTopupYearlyMethods(s []*db.GetYearlyTopupMethodsRow) []*record.TopupYearlyMethod

	// ToTopupMonthlyAmount maps a GetMonthlyTopupAmountsRow database row to a TopupMonthAmount domain model.
	ToTopupMonthlyAmount(s *db.GetMonthlyTopupAmountsRow) *record.TopupMonthAmount

	// ToTopupMonthlyAmounts maps a slice of GetMonthlyTopupAmountsRow database rows to a slice of TopupMonthAmount domain models.
	ToTopupMonthlyAmounts(s []*db.GetMonthlyTopupAmountsRow) []*record.TopupMonthAmount

	// ToTopupYearlyAmount maps a GetYearlyTopupAmountsRow database row to a TopupYearlyAmount domain model.
	ToTopupYearlyAmount(s *db.GetYearlyTopupAmountsRow) *record.TopupYearlyAmount

	// ToTopupYearlyAmounts maps a slice of GetYearlyTopupAmountsRow database rows to a slice of TopupYearlyAmount domain models.
	ToTopupYearlyAmounts(s []*db.GetYearlyTopupAmountsRow) []*record.TopupYearlyAmount

	// ToTopupMonthlyMethodByCardNumber maps a GetMonthlyTopupMethodsByCardNumberRow database row to a TopupMonthMethod domain model.
	ToTopupMonthlyMethodByCardNumber(s *db.GetMonthlyTopupMethodsByCardNumberRow) *record.TopupMonthMethod

	// ToTopupMonthlyMethodsByCardNumber maps a slice of GetMonthlyTopupMethodsByCardNumberRow database rows to a slice of TopupMonthMethod domain models.
	ToTopupMonthlyMethodsByCardNumber(s []*db.GetMonthlyTopupMethodsByCardNumberRow) []*record.TopupMonthMethod

	// ToTopupYearlyMethodByCardNumber maps a GetYearlyTopupMethodsByCardNumberRow database row to a TopupYearlyMethod domain model.
	ToTopupYearlyMethodByCardNumber(s *db.GetYearlyTopupMethodsByCardNumberRow) *record.TopupYearlyMethod

	// ToTopupYearlyMethodsByCardNumber maps a slice of GetYearlyTopupMethodsByCardNumberRow database rows to a slice of TopupYearlyMethod domain models.
	ToTopupYearlyMethodsByCardNumber(s []*db.GetYearlyTopupMethodsByCardNumberRow) []*record.TopupYearlyMethod

	// ToTopupMonthlyAmountByCardNumber maps a GetMonthlyTopupAmountsByCardNumberRow database row to a TopupMonthAmount domain model.
	ToTopupMonthlyAmountByCardNumber(s *db.GetMonthlyTopupAmountsByCardNumberRow) *record.TopupMonthAmount

	// ToTopupMonthlyAmountsByCardNumber maps a slice of GetMonthlyTopupAmountsByCardNumberRow database rows to a slice of TopupMonthAmount domain models.
	ToTopupMonthlyAmountsByCardNumber(s []*db.GetMonthlyTopupAmountsByCardNumberRow) []*record.TopupMonthAmount

	// ToTopupYearlyAmountByCardNumber maps a GetYearlyTopupAmountsByCardNumberRow database row to a TopupYearlyAmount domain model.
	ToTopupYearlyAmountByCardNumber(s *db.GetYearlyTopupAmountsByCardNumberRow) *record.TopupYearlyAmount

	// ToTopupYearlyAmountsByCardNumber maps a slice of GetYearlyTopupAmountsByCardNumberRow database rows to a slice of TopupYearlyAmount domain models.
	ToTopupYearlyAmountsByCardNumber(s []*db.GetYearlyTopupAmountsByCardNumberRow) []*record.TopupYearlyAmount

	// ToTopupRecordAll maps a GetTopupsRow database row to a TopupRecord domain model.
	ToTopupRecordAll(topup *db.GetTopupsRow) *record.TopupRecord

	// ToTopupRecordsAll maps a slice of GetTopupsRow database rows to a slice of TopupRecord domain models.
	ToTopupRecordsAll(topups []*db.GetTopupsRow) []*record.TopupRecord

	// ToTopupRecordActive maps a GetActiveTopupsRow database row to a TopupRecord domain model.
	ToTopupRecordActive(topup *db.GetActiveTopupsRow) *record.TopupRecord

	// ToTopupRecordsActive maps a slice of GetActiveTopupsRow database rows to a slice of TopupRecord domain models.
	ToTopupRecordsActive(topups []*db.GetActiveTopupsRow) []*record.TopupRecord

	// ToTopupRecordTrashed maps a GetTrashedTopupsRow database row to a TopupRecord domain model.
	ToTopupRecordTrashed(topup *db.GetTrashedTopupsRow) *record.TopupRecord

	// ToTopupRecordsTrashed maps a slice of GetTrashedTopupsRow database rows to a slice of TopupRecord domain models.
	ToTopupRecordsTrashed(topups []*db.GetTrashedTopupsRow) []*record.TopupRecord
}

type TransferRecordMapping interface {
	ToTransferRecord(transfer *db.Transfer) *record.TransferRecord
	ToTransfersRecord(transfers []*db.Transfer) []*record.TransferRecord

	// ToTransferRecordMonthStatusSuccess maps a GetMonthTransferStatusSuccessRow database row to a TransferRecordMonthStatusSuccess domain model.
	ToTransferRecordMonthStatusSuccess(s *db.GetMonthTransferStatusSuccessRow) *record.TransferRecordMonthStatusSuccess

	// ToTransferRecordsMonthStatusSuccess maps a slice of GetMonthTransferStatusSuccessRow database rows to a slice of TransferRecordMonthStatusSuccess domain models.
	ToTransferRecordsMonthStatusSuccess(Transfers []*db.GetMonthTransferStatusSuccessRow) []*record.TransferRecordMonthStatusSuccess

	// ToTransferRecordYearStatusSuccess maps a GetYearlyTransferStatusSuccessRow database row to a TransferRecordYearStatusSuccess domain model.
	ToTransferRecordYearStatusSuccess(s *db.GetYearlyTransferStatusSuccessRow) *record.TransferRecordYearStatusSuccess

	// ToTransferRecordsYearStatusSuccess maps a slice of GetYearlyTransferStatusSuccessRow database rows to a slice of TransferRecordYearStatusSuccess domain models.
	ToTransferRecordsYearStatusSuccess(Transfers []*db.GetYearlyTransferStatusSuccessRow) []*record.TransferRecordYearStatusSuccess

	// ToTransferRecordMonthStatusFailed maps a GetMonthTransferStatusFailedRow database row to a TransferRecordMonthStatusFailed domain model.
	ToTransferRecordMonthStatusFailed(s *db.GetMonthTransferStatusFailedRow) *record.TransferRecordMonthStatusFailed

	// ToTransferRecordsMonthStatusFailed maps a slice of GetMonthTransferStatusFailedRow database rows to a slice of TransferRecordMonthStatusFailed domain models.
	ToTransferRecordsMonthStatusFailed(Transfers []*db.GetMonthTransferStatusFailedRow) []*record.TransferRecordMonthStatusFailed

	// ToTransferRecordYearStatusFailed maps a GetYearlyTransferStatusFailedRow database row to a TransferRecordYearStatusFailed domain model.
	ToTransferRecordYearStatusFailed(s *db.GetYearlyTransferStatusFailedRow) *record.TransferRecordYearStatusFailed

	// ToTransferRecordsYearStatusFailed maps a slice of GetYearlyTransferStatusFailedRow database rows to a slice of TransferRecordYearStatusFailed domain models.
	ToTransferRecordsYearStatusFailed(Transfers []*db.GetYearlyTransferStatusFailedRow) []*record.TransferRecordYearStatusFailed

	// ToTransferRecordMonthStatusSuccessCardNumber maps a GetMonthTransferStatusSuccessCardNumberRow database row to a TransferRecordMonthStatusSuccess domain model.
	ToTransferRecordMonthStatusSuccessCardNumber(s *db.GetMonthTransferStatusSuccessCardNumberRow) *record.TransferRecordMonthStatusSuccess

	// ToTransferRecordsMonthStatusSuccessCardNumber maps a slice of GetMonthTransferStatusSuccessCardNumberRow database rows to a slice of TransferRecordMonthStatusSuccess domain models.
	ToTransferRecordsMonthStatusSuccessCardNumber(Transfers []*db.GetMonthTransferStatusSuccessCardNumberRow) []*record.TransferRecordMonthStatusSuccess

	// ToTransferRecordYearStatusSuccessCardNumber maps a GetYearlyTransferStatusSuccessCardNumberRow database row to a TransferRecordYearStatusSuccess domain model.
	ToTransferRecordYearStatusSuccessCardNumber(s *db.GetYearlyTransferStatusSuccessCardNumberRow) *record.TransferRecordYearStatusSuccess

	// ToTransferRecordsYearStatusSuccessCardNumber maps a slice of GetYearlyTransferStatusSuccessCardNumberRow database rows to a slice of TransferRecordYearStatusSuccess domain models.
	ToTransferRecordsYearStatusSuccessCardNumber(Transfers []*db.GetYearlyTransferStatusSuccessCardNumberRow) []*record.TransferRecordYearStatusSuccess

	// ToTransferRecordMonthStatusFailedCardNumber maps a GetMonthTransferStatusFailedCardNumberRow database row to a TransferRecordMonthStatusFailed domain model.
	ToTransferRecordMonthStatusFailedCardNumber(s *db.GetMonthTransferStatusFailedCardNumberRow) *record.TransferRecordMonthStatusFailed

	// ToTransferRecordsMonthStatusFailedCardNumber maps a slice of GetMonthTransferStatusFailedCardNumberRow database rows to a slice of TransferRecordMonthStatusFailed domain models.
	ToTransferRecordsMonthStatusFailedCardNumber(Transfers []*db.GetMonthTransferStatusFailedCardNumberRow) []*record.TransferRecordMonthStatusFailed

	// ToTransferRecordYearStatusFailedCardNumber maps a GetYearlyTransferStatusFailedCardNumberRow database row to a TransferRecordYearStatusFailed domain model.
	ToTransferRecordYearStatusFailedCardNumber(s *db.GetYearlyTransferStatusFailedCardNumberRow) *record.TransferRecordYearStatusFailed

	// ToTransferRecordsYearStatusFailedCardNumber maps a slice of GetYearlyTransferStatusFailedCardNumberRow database rows to a slice of TransferRecordYearStatusFailed domain models.
	ToTransferRecordsYearStatusFailedCardNumber(Transfers []*db.GetYearlyTransferStatusFailedCardNumberRow) []*record.TransferRecordYearStatusFailed

	// ToTransferMonthAmount maps a GetMonthlyTransferAmountsRow database row to a TransferMonthAmount domain model.
	ToTransferMonthAmount(ss *db.GetMonthlyTransferAmountsRow) *record.TransferMonthAmount

	// ToTransferMonthAmounts maps a slice of GetMonthlyTransferAmountsRow database rows to a slice of TransferMonthAmount domain models.
	ToTransferMonthAmounts(ss []*db.GetMonthlyTransferAmountsRow) []*record.TransferMonthAmount

	// ToTransferYearAmount maps a GetYearlyTransferAmountsRow database row to a TransferYearAmount domain model.
	ToTransferYearAmount(ss *db.GetYearlyTransferAmountsRow) *record.TransferYearAmount

	// ToTransferYearAmounts maps a slice of GetYearlyTransferAmountsRow database rows to a slice of TransferYearAmount domain models.
	ToTransferYearAmounts(ss []*db.GetYearlyTransferAmountsRow) []*record.TransferYearAmount

	// ToTransferMonthAmountSender maps a GetMonthlyTransferAmountsBySenderCardNumberRow database row to a TransferMonthAmount domain model.
	ToTransferMonthAmountSender(ss *db.GetMonthlyTransferAmountsBySenderCardNumberRow) *record.TransferMonthAmount

	// ToTransferMonthAmountsSender maps a slice of GetMonthlyTransferAmountsBySenderCardNumberRow database rows to a slice of TransferMonthAmount domain models.
	ToTransferMonthAmountsSender(ss []*db.GetMonthlyTransferAmountsBySenderCardNumberRow) []*record.TransferMonthAmount

	// ToTransferYearAmountSender maps a GetYearlyTransferAmountsBySenderCardNumberRow database row to a TransferYearAmount domain model.
	ToTransferYearAmountSender(ss *db.GetYearlyTransferAmountsBySenderCardNumberRow) *record.TransferYearAmount

	// ToTransferYearAmountsSender maps a slice of GetYearlyTransferAmountsBySenderCardNumberRow database rows to a slice of TransferYearAmount domain models.
	ToTransferYearAmountsSender(ss []*db.GetYearlyTransferAmountsBySenderCardNumberRow) []*record.TransferYearAmount

	// ToTransferMonthAmountReceiver maps a GetMonthlyTransferAmountsByReceiverCardNumberRow database row to a TransferMonthAmount domain model.
	ToTransferMonthAmountReceiver(ss *db.GetMonthlyTransferAmountsByReceiverCardNumberRow) *record.TransferMonthAmount

	// ToTransferMonthAmountsReceiver maps a slice of GetMonthlyTransferAmountsByReceiverCardNumberRow database rows to a slice of TransferMonthAmount domain models.
	ToTransferMonthAmountsReceiver(ss []*db.GetMonthlyTransferAmountsByReceiverCardNumberRow) []*record.TransferMonthAmount

	// ToTransferYearAmountReceiver maps a GetYearlyTransferAmountsByReceiverCardNumberRow database row to a TransferYearAmount domain model.
	ToTransferYearAmountReceiver(ss *db.GetYearlyTransferAmountsByReceiverCardNumberRow) *record.TransferYearAmount

	// ToTransferYearAmountsReceiver maps a slice of GetYearlyTransferAmountsByReceiverCardNumberRow database rows to a slice of TransferYearAmount domain models.
	ToTransferYearAmountsReceiver(ss []*db.GetYearlyTransferAmountsByReceiverCardNumberRow) []*record.TransferYearAmount

	// ToTransferRecordAll maps a GetTransfersRow database row to a TransferRecord domain model.
	ToTransferRecordAll(transfer *db.GetTransfersRow) *record.TransferRecord

	// ToTransferRecordsAll maps a slice of GetTransfersRow database rows to a slice of TransferRecord domain models.
	ToTransfersRecordAll(transfers []*db.GetTransfersRow) []*record.TransferRecord

	// ToTransferRecordActive maps a GetActiveTransfersRow database row to a TransferRecord domain model.
	ToTransferRecordActive(transfer *db.GetActiveTransfersRow) *record.TransferRecord

	// ToTransferRecordsActive maps a slice of GetActiveTransfersRow database rows to a slice of TransferRecord domain models.
	ToTransfersRecordActive(transfers []*db.GetActiveTransfersRow) []*record.TransferRecord

	// ToTransferRecordTrashed maps a GetTrashedTransfersRow database row to a TransferRecord domain model.
	ToTransferRecordTrashed(transfer *db.GetTrashedTransfersRow) *record.TransferRecord

	// ToTransferRecordsTrashed maps a slice of GetTrashedTransfersRow database rows to a slice of TransferRecord domain models.
	ToTransfersRecordTrashed(transfers []*db.GetTrashedTransfersRow) []*record.TransferRecord
}

// WithdrawRecordMapping defines a set of functions for mapping database query results
// into internal record representations used across different application layers
// (e.g., response rendering, logging, business logic processing).
type WithdrawRecordMapping interface {

	// Maps a single withdraw entity to its corresponding record.
	ToWithdrawRecord(withdraw *db.Withdraw) *record.WithdrawRecord

	// Maps a slice of withdraw entities to a slice of records.
	ToWithdrawsRecord(withdraws []*db.Withdraw) []*record.WithdrawRecord

	// Maps a withdraw row filtered by card number to its corresponding record.
	ToWithdrawByCardNumberRecord(withdraw *db.GetWithdrawsByCardNumberRow) *record.WithdrawRecord

	// Maps a slice of withdraw rows filtered by card number to a slice of records.
	ToWithdrawsByCardNumberRecord(withdraws []*db.GetWithdrawsByCardNumberRow) []*record.WithdrawRecord

	// Maps a monthly successful withdraw row to its corresponding record.
	ToWithdrawRecordMonthStatusSuccess(s *db.GetMonthWithdrawStatusSuccessRow) *record.WithdrawRecordMonthStatusSuccess

	// Maps a slice of monthly successful withdraw rows to a slice of records.
	ToWithdrawRecordsMonthStatusSuccess(withdraws []*db.GetMonthWithdrawStatusSuccessRow) []*record.WithdrawRecordMonthStatusSuccess

	// Maps a yearly successful withdraw row to its corresponding record.
	ToWithdrawRecordYearStatusSuccess(s *db.GetYearlyWithdrawStatusSuccessRow) *record.WithdrawRecordYearStatusSuccess

	// Maps a slice of yearly successful withdraw rows to a slice of records.
	ToWithdrawRecordsYearStatusSuccess(withdraws []*db.GetYearlyWithdrawStatusSuccessRow) []*record.WithdrawRecordYearStatusSuccess

	// Maps a monthly failed withdraw row to its corresponding record.
	ToWithdrawRecordMonthStatusFailed(s *db.GetMonthWithdrawStatusFailedRow) *record.WithdrawRecordMonthStatusFailed

	// Maps a slice of monthly failed withdraw rows to a slice of records.
	ToWithdrawRecordsMonthStatusFailed(withdraws []*db.GetMonthWithdrawStatusFailedRow) []*record.WithdrawRecordMonthStatusFailed

	// Maps a yearly failed withdraw row to its corresponding record.
	ToWithdrawRecordYearStatusFailed(s *db.GetYearlyWithdrawStatusFailedRow) *record.WithdrawRecordYearStatusFailed

	// Maps a slice of yearly failed withdraw rows to a slice of records.
	ToWithdrawRecordsYearStatusFailed(withdraws []*db.GetYearlyWithdrawStatusFailedRow) []*record.WithdrawRecordYearStatusFailed

	// Maps a monthly successful withdraw row filtered by card number to a record.
	ToWithdrawRecordMonthStatusSuccessCardNumber(s *db.GetMonthWithdrawStatusSuccessCardNumberRow) *record.WithdrawRecordMonthStatusSuccess

	// Maps a slice of monthly successful withdraw rows filtered by card number to records.
	ToWithdrawRecordsMonthStatusSuccessCardNumber(withdraws []*db.GetMonthWithdrawStatusSuccessCardNumberRow) []*record.WithdrawRecordMonthStatusSuccess

	// Maps a yearly successful withdraw row filtered by card number to a record.
	ToWithdrawRecordYearStatusSuccessCardNumber(s *db.GetYearlyWithdrawStatusSuccessCardNumberRow) *record.WithdrawRecordYearStatusSuccess

	// Maps a slice of yearly successful withdraw rows filtered by card number to records.
	ToWithdrawRecordsYearStatusSuccessCardNumber(withdraws []*db.GetYearlyWithdrawStatusSuccessCardNumberRow) []*record.WithdrawRecordYearStatusSuccess

	// Maps a monthly failed withdraw row filtered by card number to a record.
	ToWithdrawRecordMonthStatusFailedCardNumber(s *db.GetMonthWithdrawStatusFailedCardNumberRow) *record.WithdrawRecordMonthStatusFailed

	// Maps a slice of monthly failed withdraw rows filtered by card number to records.
	ToWithdrawRecordsMonthStatusFailedCardNumber(withdraws []*db.GetMonthWithdrawStatusFailedCardNumberRow) []*record.WithdrawRecordMonthStatusFailed

	// Maps a yearly failed withdraw row filtered by card number to a record.
	ToWithdrawRecordYearStatusFailedCardNumber(s *db.GetYearlyWithdrawStatusFailedCardNumberRow) *record.WithdrawRecordYearStatusFailed

	// Maps a slice of yearly failed withdraw rows filtered by card number to records.
	ToWithdrawRecordsYearStatusFailedCardNumber(withdraws []*db.GetYearlyWithdrawStatusFailedCardNumberRow) []*record.WithdrawRecordYearStatusFailed

	// Maps a monthly withdraw amount row to its corresponding record.
	ToWithdrawAmountMonthly(ss *db.GetMonthlyWithdrawsRow) *record.WithdrawMonthlyAmount

	// Maps a slice of monthly withdraw amount rows to records.
	ToWithdrawsAmountMonthly(ss []*db.GetMonthlyWithdrawsRow) []*record.WithdrawMonthlyAmount

	// Maps a yearly withdraw amount row to its corresponding record.
	ToWithdrawAmountYearly(ss *db.GetYearlyWithdrawsRow) *record.WithdrawYearlyAmount

	// Maps a slice of yearly withdraw amount rows to records.
	ToWithdrawsAmountYearly(ss []*db.GetYearlyWithdrawsRow) []*record.WithdrawYearlyAmount

	// Maps a monthly withdraw amount row filtered by card number to a record.
	ToWithdrawAmountMonthlyByCardNumber(ss *db.GetMonthlyWithdrawsByCardNumberRow) *record.WithdrawMonthlyAmount

	// Maps a slice of monthly withdraw amount rows filtered by card number to records.
	ToWithdrawsAmountMonthlyByCardNumber(ss []*db.GetMonthlyWithdrawsByCardNumberRow) []*record.WithdrawMonthlyAmount

	// Maps a yearly withdraw amount row filtered by card number to a record.
	ToWithdrawAmountYearlyByCardNumber(ss *db.GetYearlyWithdrawsByCardNumberRow) *record.WithdrawYearlyAmount

	// Maps a slice of yearly withdraw amount rows filtered by card number to records.
	ToWithdrawsAmountYearlyByCardNumber(ss []*db.GetYearlyWithdrawsByCardNumberRow) []*record.WithdrawYearlyAmount

	// Maps a full withdraw row to its corresponding record.
	ToWithdrawRecordAll(withdraw *db.GetWithdrawsRow) *record.WithdrawRecord

	// Maps a slice of full withdraw rows to a slice of records.
	ToWithdrawsRecordALl(withdraws []*db.GetWithdrawsRow) []*record.WithdrawRecord

	// Maps an active withdraw row to its corresponding record.
	ToWithdrawRecordActive(withdraw *db.GetActiveWithdrawsRow) *record.WithdrawRecord

	// Maps a slice of active withdraw rows to records.
	ToWithdrawsRecordActive(withdraws []*db.GetActiveWithdrawsRow) []*record.WithdrawRecord

	// Maps a trashed (soft-deleted) withdraw row to its corresponding record.
	ToWithdrawRecordTrashed(withdraw *db.GetTrashedWithdrawsRow) *record.WithdrawRecord

	// Maps a slice of trashed withdraw rows to records.
	ToWithdrawsRecordTrashed(withdraws []*db.GetTrashedWithdrawsRow) []*record.WithdrawRecord
}

// CardRecordMapping provides methods to map database rows to Card domain models.
type CardRecordMapping interface {
	// ToCardEmailRecord maps a GetUserEmailByCardNumberRow to a CardEmailRecord.
	ToCardEmailRecord(card *db.GetUserEmailByCardNumberRow) *record.CardEmailRecord

	// ToCardRecord maps a Card to a CardRecord.
	ToCardRecord(card *db.Card) *record.CardRecord
	// ToCardsRecord maps a slice of GetCardsRow to a slice of CardRecord.
	ToCardsRecord(cards []*db.GetCardsRow) []*record.CardRecord

	// ToCardRecordActive maps a GetActiveCardsWithCountRow to a CardRecord.
	ToCardRecordActive(card *db.GetActiveCardsWithCountRow) *record.CardRecord
	// ToCardRecordsActive maps a slice of GetActiveCardsWithCountRow to a slice of CardRecord.
	ToCardRecordsActive(cards []*db.GetActiveCardsWithCountRow) []*record.CardRecord

	// ToCardRecordTrashed maps a GetTrashedCardsWithCountRow to a CardRecord.
	ToCardRecordTrashed(card *db.GetTrashedCardsWithCountRow) *record.CardRecord
	// ToCardRecordsTrashed maps a slice of GetTrashedCardsWithCountRow to a slice of CardRecord.
	ToCardRecordsTrashed(cards []*db.GetTrashedCardsWithCountRow) []*record.CardRecord

	// ToMonthlyBalance maps a GetMonthlyBalancesRow to a CardMonthBalance.
	ToMonthlyBalance(card *db.GetMonthlyBalancesRow) *record.CardMonthBalance
	// ToMonthlyBalances maps a slice of GetMonthlyBalancesRow to a slice of CardMonthBalance.
	ToMonthlyBalances(cards []*db.GetMonthlyBalancesRow) []*record.CardMonthBalance

	// ToYearlyBalance maps a GetYearlyBalancesRow to a CardYearlyBalance.
	ToYearlyBalance(card *db.GetYearlyBalancesRow) *record.CardYearlyBalance
	// ToYearlyBalances maps a slice of GetYearlyBalancesRow to a slice of CardYearlyBalance.
	ToYearlyBalances(cards []*db.GetYearlyBalancesRow) []*record.CardYearlyBalance

	// ToMonthlyTopupAmount maps a GetMonthlyTopupAmountRow to a CardMonthAmount.
	ToMonthlyTopupAmount(card *db.GetMonthlyTopupAmountRow) *record.CardMonthAmount
	// ToMonthlyTopupAmounts maps a slice of GetMonthlyTopupAmountRow to a slice of CardMonthAmount.
	ToMonthlyTopupAmounts(cards []*db.GetMonthlyTopupAmountRow) []*record.CardMonthAmount

	// ToYearlyTopupAmount maps a GetYearlyTopupAmountRow to a CardYearAmount.
	ToYearlyTopupAmount(card *db.GetYearlyTopupAmountRow) *record.CardYearAmount
	// ToYearlyTopupAmounts maps a slice of GetYearlyTopupAmountRow to a slice of CardYearAmount.
	ToYearlyTopupAmounts(cards []*db.GetYearlyTopupAmountRow) []*record.CardYearAmount

	// ToMonthlyWithdrawAmount maps a GetMonthlyWithdrawAmountRow to a CardMonthAmount.
	ToMonthlyWithdrawAmount(card *db.GetMonthlyWithdrawAmountRow) *record.CardMonthAmount
	// ToMonthlyWithdrawAmounts maps a slice of GetMonthlyWithdrawAmountRow to a slice of CardMonthAmount.
	ToMonthlyWithdrawAmounts(cards []*db.GetMonthlyWithdrawAmountRow) []*record.CardMonthAmount

	// ToYearlyWithdrawAmount maps a GetYearlyWithdrawAmountRow to a CardYearAmount.
	ToYearlyWithdrawAmount(card *db.GetYearlyWithdrawAmountRow) *record.CardYearAmount
	// ToYearlyWithdrawAmounts maps a slice of GetYearlyWithdrawAmountRow to a slice of CardYearAmount.
	ToYearlyWithdrawAmounts(cards []*db.GetYearlyWithdrawAmountRow) []*record.CardYearAmount

	// ToMonthlyTransactionAmount maps a GetMonthlyTransactionAmountRow to a CardMonthAmount.
	ToMonthlyTransactionAmount(card *db.GetMonthlyTransactionAmountRow) *record.CardMonthAmount
	// ToMonthlyTransactionAmounts maps a slice of GetMonthlyTransactionAmountRow to a slice of CardMonthAmount.
	ToMonthlyTransactionAmounts(cards []*db.GetMonthlyTransactionAmountRow) []*record.CardMonthAmount

	// ToYearlyTransactionAmount maps a GetYearlyTransactionAmountRow to a CardYearAmount.
	ToYearlyTransactionAmount(card *db.GetYearlyTransactionAmountRow) *record.CardYearAmount
	// ToYearlyTransactionAmounts maps a slice of GetYearlyTransactionAmountRow to a slice of CardYearAmount.
	ToYearlyTransactionAmounts(cards []*db.GetYearlyTransactionAmountRow) []*record.CardYearAmount

	// ToMonthlyTransferSenderAmount maps a GetMonthlyTransferAmountSenderRow to a CardMonthAmount.
	ToMonthlyTransferSenderAmount(card *db.GetMonthlyTransferAmountSenderRow) *record.CardMonthAmount
	// ToMonthlyTransferSenderAmounts maps a slice of GetMonthlyTransferAmountSenderRow to a slice of CardMonthAmount.
	ToMonthlyTransferSenderAmounts(cards []*db.GetMonthlyTransferAmountSenderRow) []*record.CardMonthAmount

	// ToYearlyTransferSenderAmount maps a GetYearlyTransferAmountSenderRow to a CardYearAmount.
	ToYearlyTransferSenderAmount(card *db.GetYearlyTransferAmountSenderRow) *record.CardYearAmount
	// ToYearlyTransferSenderAmounts maps a slice of GetYearlyTransferAmountSenderRow to a slice of CardYearAmount.
	ToYearlyTransferSenderAmounts(cards []*db.GetYearlyTransferAmountSenderRow) []*record.CardYearAmount

	// ToMonthlyTransferReceiverAmount maps a GetMonthlyTransferAmountReceiverRow to a CardMonthAmount.
	ToMonthlyTransferReceiverAmount(card *db.GetMonthlyTransferAmountReceiverRow) *record.CardMonthAmount
	// ToMonthlyTransferReceiverAmounts maps a slice of GetMonthlyTransferAmountReceiverRow to a slice of CardMonthAmount.
	ToMonthlyTransferReceiverAmounts(cards []*db.GetMonthlyTransferAmountReceiverRow) []*record.CardMonthAmount

	// ToYearlyTransferReceiverAmount maps a GetYearlyTransferAmountReceiverRow to a CardYearAmount.
	ToYearlyTransferReceiverAmount(card *db.GetYearlyTransferAmountReceiverRow) *record.CardYearAmount
	// ToYearlyTransferReceiverAmounts maps a slice of GetYearlyTransferAmountReceiverRow to a slice of CardYearAmount.
	ToYearlyTransferReceiverAmounts(cards []*db.GetYearlyTransferAmountReceiverRow) []*record.CardYearAmount

	// ToMonthlyBalanceCardNumber maps a GetMonthlyBalancesByCardNumberRow to a CardMonthBalance.
	ToMonthlyBalanceCardNumber(card *db.GetMonthlyBalancesByCardNumberRow) *record.CardMonthBalance
	// ToMonthlyBalancesCardNumber maps a slice of GetMonthlyBalancesByCardNumberRow to a slice of CardMonthBalance.
	ToMonthlyBalancesCardNumber(cards []*db.GetMonthlyBalancesByCardNumberRow) []*record.CardMonthBalance

	// ToYearlyBalanceCardNumber maps a GetYearlyBalancesByCardNumberRow to a CardYearlyBalance.
	ToYearlyBalanceCardNumber(card *db.GetYearlyBalancesByCardNumberRow) *record.CardYearlyBalance
	// ToYearlyBalancesCardNumber maps a slice of GetYearlyBalancesByCardNumberRow to a slice of CardYearlyBalance.
	ToYearlyBalancesCardNumber(cards []*db.GetYearlyBalancesByCardNumberRow) []*record.CardYearlyBalance

	// ToMonthlyTopupAmountByCardNumber maps a GetMonthlyTopupAmountByCardNumberRow to a CardMonthAmount.
	ToMonthlyTopupAmountByCardNumber(card *db.GetMonthlyTopupAmountByCardNumberRow) *record.CardMonthAmount
	// ToMonthlyTopupAmountsByCardNumber maps a slice of GetMonthlyTopupAmountByCardNumberRow to a slice of CardMonthAmount.
	ToMonthlyTopupAmountsByCardNumber(cards []*db.GetMonthlyTopupAmountByCardNumberRow) []*record.CardMonthAmount

	// ToYearlyTopupAmountByCardNumber maps a GetYearlyTopupAmountByCardNumberRow to a CardYearAmount.
	ToYearlyTopupAmountByCardNumber(card *db.GetYearlyTopupAmountByCardNumberRow) *record.CardYearAmount
	// ToYearlyTopupAmountsByCardNumber maps a slice of GetYearlyTopupAmountByCardNumberRow to a slice of CardYearAmount.
	ToYearlyTopupAmountsByCardNumber(cards []*db.GetYearlyTopupAmountByCardNumberRow) []*record.CardYearAmount

	// ToMonthlyWithdrawAmountByCardNumber maps a GetMonthlyWithdrawAmountByCardNumberRow to a CardMonthAmount.
	ToMonthlyWithdrawAmountByCardNumber(card *db.GetMonthlyWithdrawAmountByCardNumberRow) *record.CardMonthAmount
	// ToMonthlyWithdrawAmountsByCardNumber maps a slice of GetMonthlyWithdrawAmountByCardNumberRow to a slice of CardMonthAmount.
	ToMonthlyWithdrawAmountsByCardNumber(cards []*db.GetMonthlyWithdrawAmountByCardNumberRow) []*record.CardMonthAmount

	// ToYearlyWithdrawAmountByCardNumber maps a GetYearlyWithdrawAmountByCardNumberRow to a CardYearAmount.
	ToYearlyWithdrawAmountByCardNumber(card *db.GetYearlyWithdrawAmountByCardNumberRow) *record.CardYearAmount
	// ToYearlyWithdrawAmountsByCardNumber maps a slice of GetYearlyWithdrawAmountByCardNumberRow to a slice of CardYearAmount.
	ToYearlyWithdrawAmountsByCardNumber(cards []*db.GetYearlyWithdrawAmountByCardNumberRow) []*record.CardYearAmount

	// ToMonthlyTransactionAmountByCardNumber maps a GetMonthlyTransactionAmountByCardNumberRow to a CardMonthAmount.
	ToMonthlyTransactionAmountByCardNumber(card *db.GetMonthlyTransactionAmountByCardNumberRow) *record.CardMonthAmount
	// ToMonthlyTransactionAmountsByCardNumber maps a slice of GetMonthlyTransactionAmountByCardNumberRow to a slice of CardMonthAmount.
	ToMonthlyTransactionAmountsByCardNumber(cards []*db.GetMonthlyTransactionAmountByCardNumberRow) []*record.CardMonthAmount

	// ToYearlyTransactionAmountByCardNumber maps a GetYearlyTransactionAmountByCardNumberRow to a CardYearAmount.
	ToYearlyTransactionAmountByCardNumber(card *db.GetYearlyTransactionAmountByCardNumberRow) *record.CardYearAmount
	// ToYearlyTransactionAmountsByCardNumber maps a slice of GetYearlyTransactionAmountByCardNumberRow to a slice of CardYearAmount.
	ToYearlyTransactionAmountsByCardNumber(cards []*db.GetYearlyTransactionAmountByCardNumberRow) []*record.CardYearAmount

	// ToMonthlyTransferSenderAmountByCardNumber maps a GetMonthlyTransferAmountBySenderRow to a CardMonthAmount.
	ToMonthlyTransferSenderAmountByCardNumber(card *db.GetMonthlyTransferAmountBySenderRow) *record.CardMonthAmount
	// ToMonthlyTransferSenderAmountsByCardNumber maps a slice of GetMonthlyTransferAmountBySenderRow to a slice of CardMonthAmount.
	ToMonthlyTransferSenderAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountBySenderRow) []*record.CardMonthAmount

	// ToYearlyTransferSenderAmountByCardNumber maps a GetYearlyTransferAmountBySenderRow to a CardYearAmount.
	ToYearlyTransferSenderAmountByCardNumber(card *db.GetYearlyTransferAmountBySenderRow) *record.CardYearAmount
	// ToYearlyTransferSenderAmountsByCardNumber maps a slice of GetYearlyTransferAmountBySenderRow to a slice of CardYearAmount.
	ToYearlyTransferSenderAmountsByCardNumber(cards []*db.GetYearlyTransferAmountBySenderRow) []*record.CardYearAmount

	// ToMonthlyTransferReceiverAmountByCardNumber maps a GetMonthlyTransferAmountByReceiverRow to a CardMonthAmount.
	ToMonthlyTransferReceiverAmountByCardNumber(card *db.GetMonthlyTransferAmountByReceiverRow) *record.CardMonthAmount
	// ToMonthlyTransferReceiverAmountsByCardNumber maps a slice of GetMonthlyTransferAmountByReceiverRow to a slice of CardMonthAmount.
	ToMonthlyTransferReceiverAmountsByCardNumber(cards []*db.GetMonthlyTransferAmountByReceiverRow) []*record.CardMonthAmount

	// ToYearlyTransferReceiverAmountByCardNumber maps a GetYearlyTransferAmountByReceiverRow to a CardYearAmount.
	ToYearlyTransferReceiverAmountByCardNumber(card *db.GetYearlyTransferAmountByReceiverRow) *record.CardYearAmount
	// ToYearlyTransferReceiverAmountsByCardNumber maps a slice of GetYearlyTransferAmountByReceiverRow to a slice of CardYearAmount.
	ToYearlyTransferReceiverAmountsByCardNumber(cards []*db.GetYearlyTransferAmountByReceiverRow) []*record.CardYearAmount
}

// TransactionRecordMapping defines a set of functions for converting transaction-related
// database rows into application-level record structures used for processing,
// logging, and API responses.
type TransactionRecordMapping interface {

	// Converts a single transaction entity to a transaction record.
	ToTransactionRecord(transaction *db.Transaction) *record.TransactionRecord

	// Converts a slice of transaction entities to a slice of transaction records.
	ToTransactionsRecord(transactions []*db.Transaction) []*record.TransactionRecord

	// Converts a transaction row filtered by card number to a transaction record.
	ToTransactionByCardNumberRecord(transaction *db.GetTransactionsByCardNumberRow) *record.TransactionRecord

	// Converts a slice of transaction rows filtered by card number to transaction records.
	ToTransactionsByCardNumberRecord(transactions []*db.GetTransactionsByCardNumberRow) []*record.TransactionRecord

	// Converts a row of successful monthly transaction status to a record.
	ToTransactionRecordMonthStatusSuccess(s *db.GetMonthTransactionStatusSuccessRow) *record.TransactionRecordMonthStatusSuccess

	// Converts multiple rows of successful monthly transaction status to records.
	ToTransactionRecordsMonthStatusSuccess(transactions []*db.GetMonthTransactionStatusSuccessRow) []*record.TransactionRecordMonthStatusSuccess

	// Converts a row of successful yearly transaction status to a record.
	ToTransactionRecordYearStatusSuccess(s *db.GetYearlyTransactionStatusSuccessRow) *record.TransactionRecordYearStatusSuccess

	// Converts multiple rows of successful yearly transaction status to records.
	ToTransactionRecordsYearStatusSuccess(transactions []*db.GetYearlyTransactionStatusSuccessRow) []*record.TransactionRecordYearStatusSuccess

	// Converts a row of failed monthly transaction status to a record.
	ToTransactionRecordMonthStatusFailed(s *db.GetMonthTransactionStatusFailedRow) *record.TransactionRecordMonthStatusFailed

	// Converts multiple rows of failed monthly transaction status to records.
	ToTransactionRecordsMonthStatusFailed(transactions []*db.GetMonthTransactionStatusFailedRow) []*record.TransactionRecordMonthStatusFailed

	// Converts a row of failed yearly transaction status to a record.
	ToTransactionRecordYearStatusFailed(s *db.GetYearlyTransactionStatusFailedRow) *record.TransactionRecordYearStatusFailed

	// Converts multiple rows of failed yearly transaction status to records.
	ToTransactionRecordsYearStatusFailed(transactions []*db.GetYearlyTransactionStatusFailedRow) []*record.TransactionRecordYearStatusFailed

	// Converts a row of successful monthly transaction status filtered by card number to a record.
	ToTransactionRecordMonthStatusSuccessCardNumber(s *db.GetMonthTransactionStatusSuccessCardNumberRow) *record.TransactionRecordMonthStatusSuccess

	// Converts multiple rows of successful monthly transaction status filtered by card number to records.
	ToTransactionRecordsMonthStatusSuccessCardNumber(transactions []*db.GetMonthTransactionStatusSuccessCardNumberRow) []*record.TransactionRecordMonthStatusSuccess

	// Converts a row of successful yearly transaction status filtered by card number to a record.
	ToTransactionRecordYearStatusSuccessCardNumber(s *db.GetYearlyTransactionStatusSuccessCardNumberRow) *record.TransactionRecordYearStatusSuccess

	// Converts multiple rows of successful yearly transaction status filtered by card number to records.
	ToTransactionRecordsYearStatusSuccessCardNumber(transactions []*db.GetYearlyTransactionStatusSuccessCardNumberRow) []*record.TransactionRecordYearStatusSuccess

	// Converts a row of failed monthly transaction status filtered by card number to a record.
	ToTransactionRecordMonthStatusFailedCardNumber(s *db.GetMonthTransactionStatusFailedCardNumberRow) *record.TransactionRecordMonthStatusFailed

	// Converts multiple rows of failed monthly transaction status filtered by card number to records.
	ToTransactionRecordsMonthStatusFailedCardNumber(transactions []*db.GetMonthTransactionStatusFailedCardNumberRow) []*record.TransactionRecordMonthStatusFailed

	// Converts a row of failed yearly transaction status filtered by card number to a record.
	ToTransactionRecordYearStatusFailedCardNumber(s *db.GetYearlyTransactionStatusFailedCardNumberRow) *record.TransactionRecordYearStatusFailed

	// Converts multiple rows of failed yearly transaction status filtered by card number to records.
	ToTransactionRecordsYearStatusFailedCardNumber(transactions []*db.GetYearlyTransactionStatusFailedCardNumberRow) []*record.TransactionRecordYearStatusFailed

	// Converts a row of monthly payment method statistics to a record.
	ToTransactionMonthlyMethod(ss *db.GetMonthlyPaymentMethodsRow) *record.TransactionMonthMethod

	// Converts multiple rows of monthly payment method statistics to records.
	ToTransactionMonthlyMethods(ss []*db.GetMonthlyPaymentMethodsRow) []*record.TransactionMonthMethod

	// Converts a row of yearly payment method statistics to a record.
	ToTransactionYearlyMethod(ss *db.GetYearlyPaymentMethodsRow) *record.TransactionYearMethod

	// Converts multiple rows of yearly payment method statistics to records.
	ToTransactionYearlyMethods(ss []*db.GetYearlyPaymentMethodsRow) []*record.TransactionYearMethod

	// Converts a row of monthly transaction amounts to a record.
	ToTransactionMonthlyAmount(ss *db.GetMonthlyAmountsRow) *record.TransactionMonthAmount

	// Converts multiple rows of monthly transaction amounts to records.
	ToTransactionMonthlyAmounts(ss []*db.GetMonthlyAmountsRow) []*record.TransactionMonthAmount

	// Converts a row of yearly transaction amounts to a record.
	ToTransactionYearlyAmount(ss *db.GetYearlyAmountsRow) *record.TransactionYearlyAmount

	// Converts multiple rows of yearly transaction amounts to records.
	ToTransactionYearlyAmounts(ss []*db.GetYearlyAmountsRow) []*record.TransactionYearlyAmount

	// Converts a row of monthly payment methods filtered by card number to a record.
	ToTransactionMonthlyMethodByCardNumber(ss *db.GetMonthlyPaymentMethodsByCardNumberRow) *record.TransactionMonthMethod

	// Converts multiple rows of monthly payment methods filtered by card number to records.
	ToTransactionMonthlyMethodsByCardNumber(ss []*db.GetMonthlyPaymentMethodsByCardNumberRow) []*record.TransactionMonthMethod

	// Converts a row of yearly payment methods filtered by card number to a record.
	ToTransactionYearlyMethodByCardNumber(ss *db.GetYearlyPaymentMethodsByCardNumberRow) *record.TransactionYearMethod

	// Converts multiple rows of yearly payment methods filtered by card number to records.
	ToTransactionYearlyMethodsByCardNumber(ss []*db.GetYearlyPaymentMethodsByCardNumberRow) []*record.TransactionYearMethod

	// Converts a row of monthly transaction amounts filtered by card number to a record.
	ToTransactionMonthlyAmountByCardNumber(ss *db.GetMonthlyAmountsByCardNumberRow) *record.TransactionMonthAmount

	// Converts multiple rows of monthly transaction amounts filtered by card number to records.
	ToTransactionMonthlyAmountsByCardNumber(ss []*db.GetMonthlyAmountsByCardNumberRow) []*record.TransactionMonthAmount

	// Converts a row of yearly transaction amounts filtered by card number to a record.
	ToTransactionYearlyAmountByCardNumber(ss *db.GetYearlyAmountsByCardNumberRow) *record.TransactionYearlyAmount

	// Converts multiple rows of yearly transaction amounts filtered by card number to records.
	ToTransactionYearlyAmountsByCardNumber(ss []*db.GetYearlyAmountsByCardNumberRow) []*record.TransactionYearlyAmount

	// Converts a full transaction row to a transaction record.
	ToTransactionRecordAll(transaction *db.GetTransactionsRow) *record.TransactionRecord

	// Converts multiple full transaction rows to transaction records.
	ToTransactionsRecordAll(transactions []*db.GetTransactionsRow) []*record.TransactionRecord

	// Converts an active transaction row to a transaction record.
	ToTransactionRecordActive(transaction *db.GetActiveTransactionsRow) *record.TransactionRecord

	// Converts multiple active transaction rows to transaction records.
	ToTransactionsRecordActive(transactions []*db.GetActiveTransactionsRow) []*record.TransactionRecord

	// Converts a trashed (soft-deleted) transaction row to a transaction record.
	ToTransactionRecordTrashed(transaction *db.GetTrashedTransactionsRow) *record.TransactionRecord

	// Converts multiple trashed (soft-deleted) transaction rows to transaction records.
	ToTransactionsRecordTrashed(transactions []*db.GetTrashedTransactionsRow) []*record.TransactionRecord
}

type MerchantDocumentMapping interface {
	// ToGetMerchantDocument maps a MerchantDocument database row to a MerchantDocumentRecord domain model.
	ToGetMerchantDocument(doc *db.MerchantDocument) *record.MerchantDocumentRecord

	// ToMerchantDocumentRecord maps a GetMerchantDocumentsRow to a MerchantDocumentRecord domain model.
	ToMerchantDocumentRecord(doc *db.GetMerchantDocumentsRow) *record.MerchantDocumentRecord
	// ToMerchantDocumentsRecord maps a slice of GetMerchantDocumentsRow to a slice of MerchantDocumentRecord domain models.
	ToMerchantDocumentsRecord(docs []*db.GetMerchantDocumentsRow) []*record.MerchantDocumentRecord

	// ToMerchantDocumentActiveRecord maps a GetActiveMerchantDocumentsRow to a MerchantDocumentRecord domain model.
	ToMerchantDocumentActiveRecord(doc *db.GetActiveMerchantDocumentsRow) *record.MerchantDocumentRecord
	// ToMerchantDocumentsActiveRecord maps a slice of GetActiveMerchantDocumentsRow to a slice of MerchantDocumentRecord domain models.
	ToMerchantDocumentsActiveRecord(docs []*db.GetActiveMerchantDocumentsRow) []*record.MerchantDocumentRecord

	// ToMerchantDocumentTrashedRecord maps a GetTrashedMerchantDocumentsRow to a MerchantDocumentRecord domain model.
	ToMerchantDocumentTrashedRecord(doc *db.GetTrashedMerchantDocumentsRow) *record.MerchantDocumentRecord
	// ToMerchantDocumentsTrashedRecord maps a slice of GetTrashedMerchantDocumentsRow to a slice of MerchantDocumentRecord domain models.
	ToMerchantDocumentsTrashedRecord(docs []*db.GetTrashedMerchantDocumentsRow) []*record.MerchantDocumentRecord
}

// MerchantRecordMapping defines a set of methods for mapping merchant-related
// database rows into internal application records, used for reporting, logging,
// analytics, and API responses.
type MerchantRecordMapping interface {

	// Converts a single merchant entity to a merchant record.
	ToMerchantRecord(merchant *db.Merchant) *record.MerchantRecord

	// Converts a slice of merchant entities to merchant records.
	ToMerchantsRecord(merchants []*db.Merchant) []*record.MerchantRecord

	// Converts a monthly total amount row into a merchant monthly total amount record.
	ToMerchantMonthlyTotalAmount(ms *db.GetMonthlyTotalAmountMerchantRow) *record.MerchantMonthlyTotalAmount

	// Converts multiple monthly total amount rows into merchant monthly total amount records.
	ToMerchantMonthlyTotalAmounts(ms []*db.GetMonthlyTotalAmountMerchantRow) []*record.MerchantMonthlyTotalAmount

	// Converts a yearly total amount row into a merchant yearly total amount record.
	ToMerchantYearlyTotalAmount(ms *db.GetYearlyTotalAmountMerchantRow) *record.MerchantYearlyTotalAmount

	// Converts multiple yearly total amount rows into merchant yearly total amount records.
	ToMerchantYearlyTotalAmounts(ms []*db.GetYearlyTotalAmountMerchantRow) []*record.MerchantYearlyTotalAmount

	// Converts a transaction row to a merchant transaction record.
	ToMerchantTransactionRecord(merchant *db.FindAllTransactionsRow) *record.MerchantTransactionsRecord

	// Converts multiple transaction rows to merchant transaction records.
	ToMerchantsTransactionRecord(merchants []*db.FindAllTransactionsRow) []*record.MerchantTransactionsRecord

	// Converts a full merchant row to a merchant record.
	ToMerchantGetAllRecord(merchant *db.GetMerchantsRow) *record.MerchantRecord

	// Converts multiple full merchant rows to merchant records.
	ToMerchantsGetAllRecord(merchants []*db.GetMerchantsRow) []*record.MerchantRecord

	// Converts a monthly payment method row to a merchant monthly payment method record.
	ToMerchantMonthlyPaymentMethod(ms *db.GetMonthlyPaymentMethodsMerchantRow) *record.MerchantMonthlyPaymentMethod

	// Converts multiple monthly payment method rows to merchant monthly payment method records.
	ToMerchantMonthlyPaymentMethods(ms []*db.GetMonthlyPaymentMethodsMerchantRow) []*record.MerchantMonthlyPaymentMethod

	// Converts a yearly payment method row to a merchant yearly payment method record.
	ToMerchantYearlyPaymentMethod(ms *db.GetYearlyPaymentMethodMerchantRow) *record.MerchantYearlyPaymentMethod

	// Converts multiple yearly payment method rows to merchant yearly payment method records.
	ToMerchantYearlyPaymentMethods(ms []*db.GetYearlyPaymentMethodMerchantRow) []*record.MerchantYearlyPaymentMethod

	// Converts a monthly amount row to a merchant monthly amount record.
	ToMerchantMonthlyAmount(ms *db.GetMonthlyAmountMerchantRow) *record.MerchantMonthlyAmount

	// Converts multiple monthly amount rows to merchant monthly amount records.
	ToMerchantMonthlyAmounts(ms []*db.GetMonthlyAmountMerchantRow) []*record.MerchantMonthlyAmount

	// Converts a yearly amount row to a merchant yearly amount record.
	ToMerchantYearlyAmount(ms *db.GetYearlyAmountMerchantRow) *record.MerchantYearlyAmount

	// Converts multiple yearly amount rows to merchant yearly amount records.
	ToMerchantYearlyAmounts(ms []*db.GetYearlyAmountMerchantRow) []*record.MerchantYearlyAmount

	// Converts a transaction row filtered by merchant to a merchant transaction record.
	ToMerchantTransactionByMerchantRecord(merchant *db.FindAllTransactionsByMerchantRow) *record.MerchantTransactionsRecord

	// Converts multiple transaction rows filtered by merchant to merchant transaction records.
	ToMerchantsTransactionByMerchantRecord(merchants []*db.FindAllTransactionsByMerchantRow) []*record.MerchantTransactionsRecord

	// Converts a monthly payment method row filtered by merchant to a merchant monthly payment method record.
	ToMerchantMonthlyPaymentMethodByMerchant(ms *db.GetMonthlyPaymentMethodByMerchantsRow) *record.MerchantMonthlyPaymentMethod

	// Converts multiple monthly payment method rows filtered by merchant to merchant monthly payment method records.
	ToMerchantMonthlyPaymentMethodsByMerchant(ms []*db.GetMonthlyPaymentMethodByMerchantsRow) []*record.MerchantMonthlyPaymentMethod

	// Converts a yearly payment method row filtered by merchant to a merchant yearly payment method record.
	ToMerchantYearlyPaymentMethodByMerchant(ms *db.GetYearlyPaymentMethodByMerchantsRow) *record.MerchantYearlyPaymentMethod

	// Converts multiple yearly payment method rows filtered by merchant to merchant yearly payment method records.
	ToMerchantYearlyPaymentMethodsByMerchant(ms []*db.GetYearlyPaymentMethodByMerchantsRow) []*record.MerchantYearlyPaymentMethod

	// Converts a monthly amount row filtered by merchant to a merchant monthly amount record.
	ToMerchantMonthlyAmountByMerchant(ms *db.GetMonthlyAmountByMerchantsRow) *record.MerchantMonthlyAmount

	// Converts multiple monthly amount rows filtered by merchant to merchant monthly amount records.
	ToMerchantMonthlyAmountsByMerchant(ms []*db.GetMonthlyAmountByMerchantsRow) []*record.MerchantMonthlyAmount

	// Converts a yearly amount row filtered by merchant to a merchant yearly amount record.
	ToMerchantYearlyAmountByMerchant(ms *db.GetYearlyAmountByMerchantsRow) *record.MerchantYearlyAmount

	// Converts multiple yearly amount rows filtered by merchant to merchant yearly amount records.
	ToMerchantYearlyAmountsByMerchant(ms []*db.GetYearlyAmountByMerchantsRow) []*record.MerchantYearlyAmount

	// Converts a monthly total amount row filtered by merchant to a merchant monthly total amount record.
	ToMerchantMonthlyTotalAmountByMerchant(ms *db.GetMonthlyTotalAmountByMerchantRow) *record.MerchantMonthlyTotalAmount

	// Converts multiple monthly total amount rows filtered by merchant to merchant monthly total amount records.
	ToMerchantMonthlyTotalAmountsByMerchant(ms []*db.GetMonthlyTotalAmountByMerchantRow) []*record.MerchantMonthlyTotalAmount

	// Converts a yearly total amount row filtered by merchant to a merchant yearly total amount record.
	ToMerchantYearlyTotalAmountByMerchant(ms *db.GetYearlyTotalAmountByMerchantRow) *record.MerchantYearlyTotalAmount

	// Converts multiple yearly total amount rows filtered by merchant to merchant yearly total amount records.
	ToMerchantYearlyTotalAmountsByMerchant(ms []*db.GetYearlyTotalAmountByMerchantRow) []*record.MerchantYearlyTotalAmount

	// Converts a transaction row filtered by API key to a merchant transaction record.
	ToMerchantTransactionByApikeyRecord(merchant *db.FindAllTransactionsByApikeyRow) *record.MerchantTransactionsRecord

	// Converts multiple transaction rows filtered by API key to merchant transaction records.
	ToMerchantsTransactionByApikeyRecord(merchants []*db.FindAllTransactionsByApikeyRow) []*record.MerchantTransactionsRecord

	// Converts a monthly payment method row filtered by API key to a merchant monthly payment method record.
	ToMerchantMonthlyPaymentMethodByApikey(ms *db.GetMonthlyPaymentMethodByApikeyRow) *record.MerchantMonthlyPaymentMethod

	// Converts multiple monthly payment method rows filtered by API key to merchant monthly payment method records.
	ToMerchantMonthlyPaymentMethodsByApikey(ms []*db.GetMonthlyPaymentMethodByApikeyRow) []*record.MerchantMonthlyPaymentMethod

	// Converts a yearly payment method row filtered by API key to a merchant yearly payment method record.
	ToMerchantYearlyPaymentMethodByApikey(ms *db.GetYearlyPaymentMethodByApikeyRow) *record.MerchantYearlyPaymentMethod

	// Converts multiple yearly payment method rows filtered by API key to merchant yearly payment method records.
	ToMerchantYearlyPaymentMethodsByApikey(ms []*db.GetYearlyPaymentMethodByApikeyRow) []*record.MerchantYearlyPaymentMethod

	// Converts a monthly amount row filtered by API key to a merchant monthly amount record.
	ToMerchantMonthlyAmountByApikey(ms *db.GetMonthlyAmountByApikeyRow) *record.MerchantMonthlyAmount

	// Converts multiple monthly amount rows filtered by API key to merchant monthly amount records.
	ToMerchantMonthlyAmountsByApikey(ms []*db.GetMonthlyAmountByApikeyRow) []*record.MerchantMonthlyAmount

	// Converts a yearly amount row filtered by API key to a merchant yearly amount record.
	ToMerchantYearlyAmountByApikey(ms *db.GetYearlyAmountByApikeyRow) *record.MerchantYearlyAmount

	// Converts multiple yearly amount rows filtered by API key to merchant yearly amount records.
	ToMerchantYearlyAmountsByApikey(ms []*db.GetYearlyAmountByApikeyRow) []*record.MerchantYearlyAmount

	// Converts a monthly total amount row filtered by API key to a merchant monthly total amount record.
	ToMerchantMonthlyTotalAmountByApikey(ms *db.GetMonthlyTotalAmountByApikeyRow) *record.MerchantMonthlyTotalAmount

	// Converts multiple monthly total amount rows filtered by API key to merchant monthly total amount records.
	ToMerchantMonthlyTotalAmountsByApikey(ms []*db.GetMonthlyTotalAmountByApikeyRow) []*record.MerchantMonthlyTotalAmount

	// Converts a yearly total amount row filtered by API key to a merchant yearly total amount record.
	ToMerchantYearlyTotalAmountByApikey(ms *db.GetYearlyTotalAmountByApikeyRow) *record.MerchantYearlyTotalAmount

	// Converts multiple yearly total amount rows filtered by API key to merchant yearly total amount records.
	ToMerchantYearlyTotalAmountsByApikey(ms []*db.GetYearlyTotalAmountByApikeyRow) []*record.MerchantYearlyTotalAmount

	// Converts an active merchant row to a merchant record.
	ToMerchantActiveRecord(merchant *db.GetActiveMerchantsRow) *record.MerchantRecord

	// Converts multiple active merchant rows to merchant records.
	ToMerchantsActiveRecord(merchants []*db.GetActiveMerchantsRow) []*record.MerchantRecord

	// Converts a trashed (soft-deleted) merchant row to a merchant record.
	ToMerchantTrashedRecord(merchant *db.GetTrashedMerchantsRow) *record.MerchantRecord

	// Converts multiple trashed merchant rows to merchant records.
	ToMerchantsTrashedRecord(merchants []*db.GetTrashedMerchantsRow) []*record.MerchantRecord
}
