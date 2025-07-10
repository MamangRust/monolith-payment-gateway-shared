package transferrecordmapper

// transferRecordMapper provides methods to map Transfer database rows to TransferRecord domain models.
type transferRecordMapper struct {
	TransferQueryRecordMapper   TransferQueryRecordMapper
	TransferCommandRecordMapper TransferCommandRecordMapper
}

// NewTransferRecordMapper returns a new instance of transferRecordMapper which provides methods to map Transfer database rows to TransferRecord domain models.
func NewTransferRecordMapper() *transferRecordMapper {
	return &transferRecordMapper{
		TransferQueryRecordMapper:   NewTransferQueryRecordMapper(),
		TransferCommandRecordMapper: NewTransferCommandRecordMapper(),
	}
}
