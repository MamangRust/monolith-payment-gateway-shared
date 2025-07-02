package recordmapper

import (
	db "github.com/MamangRust/monolith-payment-gateway-pkg/database/schema"
	"github.com/MamangRust/monolith-payment-gateway-shared/domain/record"
)

// merchantDocumentRecordMapper provides methods to map MerchantDocument database rows to MerchantDocumentRecord domain models.
type merchantDocumentRecordMapper struct {
}

// NewMerchantDocumentRecordMapper returns a new instance of merchantDocumentRecordMapper which provides methods to map MerchantDocument database rows to MerchantDocumentRecord domain models.
func NewMerchantDocumentRecordMapper() *merchantDocumentRecordMapper {
	return &merchantDocumentRecordMapper{}
}

// MerchantDocumentRecord maps a MerchantDocument database row to a MerchantDocumentRecord domain model.
//
// Args:
//   - merchantdocument: A pointer to a MerchantDocument representing the database row.
//
// Returns:
//   - A pointer to a MerchantDocumentRecord containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantDocumentRecordMapper) ToGetMerchantDocument(doc *db.MerchantDocument) *record.MerchantDocumentRecord {
	return &record.MerchantDocumentRecord{
		ID:           int(doc.DocumentID),
		MerchantID:   int(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentURL:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         doc.Note.String,
		CreatedAt:    doc.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    doc.UpdatedAt.Time.Format("2006-01-02"),
	}
}

// ToMerchantDocumentRecord maps a GetMerchantDocumentsRow database row to a MerchantDocumentRecord domain model.
//
// Args:
//   - doc: A pointer to a GetMerchantDocumentsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantDocumentRecord containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantDocumentRecordMapper) ToMerchantDocumentRecord(doc *db.GetMerchantDocumentsRow) *record.MerchantDocumentRecord {
	var deletedAt *string

	if doc.DeletedAt.Valid {
		formattedDeletedAt := doc.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formattedDeletedAt
	}

	return &record.MerchantDocumentRecord{
		ID:           int(doc.DocumentID),
		MerchantID:   int(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentURL:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         doc.Note.String,
		CreatedAt:    doc.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    doc.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

// ToMerchantDocumentsRecord maps a slice of GetMerchantDocumentsRow database rows to a slice of MerchantDocumentRecord domain models.
//
// Args:
//   - docs: A slice of pointers to GetMerchantDocumentsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantDocumentRecord containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantDocumentRecordMapper) ToMerchantDocumentsRecord(docs []*db.GetMerchantDocumentsRow) []*record.MerchantDocumentRecord {
	var records []*record.MerchantDocumentRecord
	for _, doc := range docs {
		records = append(records, m.ToMerchantDocumentRecord(doc))
	}
	return records
}

// ToMerchantDocumentActiveRecord maps a GetActiveMerchantDocumentsRow database row to a MerchantDocumentRecord domain model.
//
// Args:
//   - doc: A pointer to a GetActiveMerchantDocumentsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantDocumentRecord containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantDocumentRecordMapper) ToMerchantDocumentActiveRecord(doc *db.GetActiveMerchantDocumentsRow) *record.MerchantDocumentRecord {
	var deletedAt *string
	if doc.DeletedAt.Valid {
		formattedDeletedAt := doc.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formattedDeletedAt
	}
	return &record.MerchantDocumentRecord{
		ID:           int(doc.DocumentID),
		MerchantID:   int(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentURL:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         doc.Note.String,
		CreatedAt:    doc.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    doc.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

// ToMerchantDocumentsActiveRecord maps a slice of GetActiveMerchantDocumentsRow database rows to a slice of MerchantDocumentRecord domain models.
//
// Args:
//   - docs: A slice of pointers to GetActiveMerchantDocumentsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantDocumentRecord containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantDocumentRecordMapper) ToMerchantDocumentsActiveRecord(docs []*db.GetActiveMerchantDocumentsRow) []*record.MerchantDocumentRecord {
	var records []*record.MerchantDocumentRecord
	for _, doc := range docs {
		records = append(records, m.ToMerchantDocumentActiveRecord(doc))
	}
	return records
}

// ToMerchantDocumentTrashedRecord maps a GetTrashedMerchantDocumentsRow database row to a MerchantDocumentRecord domain model.
//
// Args:
//   - doc: A pointer to a GetTrashedMerchantDocumentsRow representing the database row.
//
// Returns:
//   - A pointer to a MerchantDocumentRecord containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantDocumentRecordMapper) ToMerchantDocumentTrashedRecord(doc *db.GetTrashedMerchantDocumentsRow) *record.MerchantDocumentRecord {
	var deletedAt *string
	if doc.DeletedAt.Valid {
		formattedDeletedAt := doc.DeletedAt.Time.Format("2006-01-02")
		deletedAt = &formattedDeletedAt
	}

	return &record.MerchantDocumentRecord{
		ID:           int(doc.DocumentID),
		MerchantID:   int(doc.MerchantID),
		DocumentType: doc.DocumentType,
		DocumentURL:  doc.DocumentUrl,
		Status:       doc.Status,
		Note:         doc.Note.String,
		CreatedAt:    doc.CreatedAt.Time.Format("2006-01-02"),
		UpdatedAt:    doc.UpdatedAt.Time.Format("2006-01-02"),
		DeletedAt:    deletedAt,
	}
}

// ToMerchantDocumentsTrashedRecord maps a slice of GetTrashedMerchantDocumentsRow database rows to a slice of
// MerchantDocumentRecord domain models.
//
// Args:
//   - docs: A slice of pointers to GetTrashedMerchantDocumentsRow representing the database rows.
//
// Returns:
//   - A slice of pointers to MerchantDocumentRecord containing the mapped data, including
//     ID, MerchantID, DocumentType, DocumentURL, Status, Note, CreatedAt, UpdatedAt, and DeletedAt.
func (m *merchantDocumentRecordMapper) ToMerchantDocumentsTrashedRecord(docs []*db.GetTrashedMerchantDocumentsRow) []*record.MerchantDocumentRecord {
	var records []*record.MerchantDocumentRecord
	for _, doc := range docs {
		records = append(records, m.ToMerchantDocumentTrashedRecord(doc))
	}
	return records
}
