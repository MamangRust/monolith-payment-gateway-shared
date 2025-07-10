package merchantrecordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

// merchantDocumentCommandMapper provides methods to map MerchantDocument database rows to MerchantDocumentRecord domain models.
type merchantDocumentCommandMapper struct {
}

// NewMerchantDocumentCommandMapper returns a new instance of
// merchantDocumentCommandMapper which provides methods to map
// MerchantDocument database rows to MerchantDocumentRecord domain models.
func NewMerchantDocumentCommandMapper() MerchantCommandRecordMapper {
	return &merchantDocumentCommandMapper{}
}

// ToMerchantRecord maps a Merchant to a MerchantRecord domain model.
//
// Parameters:
//   - merchant: A pointer to a Merchant representing the database row.
//
// Returns:
//   - A pointer to a MerchantRecord containing the mapped data, including
//     ID, Name, ApiKey, UserID, Status, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantDocumentCommandMapper) ToMerchantRecord(merchant *db.Merchant) *record.MerchantRecord {
	var deletedAt *string

	if merchant.DeletedAt.Valid {
		formatedDeletedAt := merchant.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formatedDeletedAt
	}

	return &record.MerchantRecord{
		ID:        int(merchant.MerchantID),
		Name:      merchant.Name,
		ApiKey:    merchant.ApiKey,
		UserID:    int(merchant.UserID),
		Status:    merchant.Status,
		CreatedAt: merchant.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt: merchant.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt: deletedAt,
	}
}
