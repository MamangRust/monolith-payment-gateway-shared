package withdrawrepositoryerrors

import "errors"

// ErrGetMonthWithdrawStatusSuccessFailed is used when the system fails to get the monthly withdraw status success
var ErrGetMonthWithdrawStatusSuccessFailed = errors.New("failed to get monthly withdraw status success")

// ErrGetYearlyWithdrawStatusSuccessFailed is used when the system fails to get the yearly withdraw status success
var ErrGetYearlyWithdrawStatusSuccessFailed = errors.New("failed to get yearly withdraw status success")

// ErrGetMonthWithdrawStatusSuccessByCardFailed is used when the system fails to get the monthly withdraw status success by card number
var ErrGetMonthWithdrawStatusSuccessByCardFailed = errors.New("failed to get monthly withdraw status success by card number")

// ErrGetYearlyWithdrawStatusSuccessByCardFailed is used when the system fails to get the yearly withdraw status success by card number
var ErrGetYearlyWithdrawStatusSuccessByCardFailed = errors.New("failed to get yearly withdraw status success by card number")

// ErrGetMonthWithdrawStatusFailedFailed is used when the system fails to get the monthly withdraw status failed
var ErrGetMonthWithdrawStatusFailedFailed = errors.New("failed to get monthly withdraw status failed")

// ErrGetYearlyWithdrawStatusFailedFailed is used when the system fails to get the yearly withdraw status failed
var ErrGetYearlyWithdrawStatusFailedFailed = errors.New("failed to get yearly withdraw status failed")

// ErrGetMonthWithdrawStatusFailedByCardFailed is used when the system fails to get the monthly withdraw status failed by card number
var ErrGetMonthWithdrawStatusFailedByCardFailed = errors.New("failed to get monthly withdraw status failed by card number")

// ErrGetYearlyWithdrawStatusFailedByCardFailed is used when the system fails to get the yearly withdraw status failed by card number
var ErrGetYearlyWithdrawStatusFailedByCardFailed = errors.New("failed to get yearly withdraw status failed by card number")
