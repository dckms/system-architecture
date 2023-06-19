package grade_1

import (
    "time"
)

type Exporter[T any] interface {
    SetState(T)
}

type Exportable[T any] interface {
    ExportTo(Exporter[T])
}

type ExportableUint uint

func (e ExportableUint) Export(ex Exporter[uint]) {
    ex.SetState(uint(e))
}

type MemberId ExportableUint
type Grade ExportableUint
type EndorsementCount ExportableUint

type EndorserExporterSetter interface {
    SetId(MemberId)
    SetGrade(Grade)
    SetAvailableEndorsementCount(EndorsementCount)
    SetPendingEndorsementCount(EndorsementCount)
    SetVersion(uint)
    SetCreatedAt(time.Time)
}

type UintExporter uint

func (e *UintExporter) SetState(value uint) {
    *e = UintExporter(value)
}

type Endorser struct {
    id                        MemberId
    grade                     Grade
    availableEndorsementCount EndorsementCount
    pendingEndorsementCount   EndorsementCount
    version                   uint
    createdAt                 time.Time
}

func (e Endorser) Export(ex EndorserExporterSetter) {
    ex.SetId(e.id)
    ex.SetGrade(e.grade)
    ex.SetAvailableEndorsementCount(e.availableEndorsementCount)
    ex.SetPendingEndorsementCount(e.pendingEndorsementCount)
    ex.SetVersion(e.version)
    ex.SetCreatedAt(e.createdAt)
}

type EndorserExporter struct {
    Id                        UintExporter
    Grade                     UintExporter
    AvailableEndorsementCount UintExporter
    PendingEndorsementCount   UintExporter
    Version                   uint
    CreatedAt                 time.Time
}

func (ex *EndorserExporter) SetId(val MemberId) {
    val.Export(&ex.Id)
}

func (ex *EndorserExporter) SetGrade(val Grade) {
    val.Export(&ex.Grade)
}

func (ex *EndorserExporter) SetAvailableEndorsementCount(val EndorsementCount) {
    val.Export(&ex.AvailableEndorsementCount)
}

func (ex *EndorserExporter) SetPendingEndorsementCount(val EndorsementCount) {
    val.Export(&ex.PendingEndorsementCount)
}

func (ex *EndorserExporter) SetVersion(val uint) {
    ex.Version = val
}

func (ex *EndorserExporter) SetCreatedAt(val time.Time) {
    ex.CreatedAt = val
}
