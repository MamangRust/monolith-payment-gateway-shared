package saldorecordmapper

// SaldoRecordMapper provides methods to map Saldo database rows to SaldoRecord domain models.
type saldoRecordMapepr struct {
	SaldoQueryRecordMapping     SaldoQueryRecordMapping
	SaldoCommandRecordMapping   SaldoCommandRecordMapping
	SaldoStatisticRecordMapping SaldoStatisticRecordMapping
}

// NewSaldoRecordMapper returns a new instance of SaldoRecordMapper which provides methods to map Saldo database rows to SaldoRecord domain models.
func NewSaldoRecordMapper() *saldoRecordMapepr {
	return &saldoRecordMapepr{
		SaldoQueryRecordMapping:     NewSaldoQueryRecordMapper(),
		SaldoCommandRecordMapping:   NewSaldoCommandRecordMapper(),
		SaldoStatisticRecordMapping: NewSaldoStatisticRecordMapper(),
	}
}
