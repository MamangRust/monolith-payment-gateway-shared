package recordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

type resetTokenRecordMapper struct {
}

func NewResetTokenRecordMapper() *resetTokenRecordMapper {
	return &resetTokenRecordMapper{}
}

func (r *resetTokenRecordMapper) ToResetTokenRecord(resetToken *db.ResetToken) *record.ResetTokenRecord {
	return &record.ResetTokenRecord{
		ID:        int64(resetToken.ID),
		UserID:    resetToken.UserID,
		Token:     resetToken.Token,
		ExpiredAt: resetToken.ExpiryDate.String(),
	}
}
