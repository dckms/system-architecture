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

type RecognizerExporterSetter interface {
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

type Recognizer struct {
    id                        MemberId
    grade                     Grade
    availableEndorsementCount EndorsementCount
    pendingEndorsementCount   EndorsementCount
    version                   uint
    createdAt                 time.Time
}

func (r Recognizer) Export(ex RecognizerExporterSetter) {
    ex.SetId(r.id)
    ex.SetGrade(r.grade)
    ex.SetAvailableEndorsementCount(r.availableEndorsementCount)
    ex.SetPendingEndorsementCount(r.pendingEndorsementCount)
    ex.SetVersion(r.version)
    ex.SetCreatedAt(r.createdAt)
}

type RecognizerExporter struct {
    Id                        UintExporter
    Grade                     UintExporter
    AvailableEndorsementCount UintExporter
    PendingEndorsementCount   UintExporter
    Version                   uint
    CreatedAt                 time.Time
}

func (ex *RecognizerExporter) SetId(val MemberId) {
    val.Export(&ex.Id)
}

func (ex *RecognizerExporter) SetGrade(val Grade) {
    val.Export(&ex.Grade)
}

func (ex *RecognizerExporter) SetAvailableEndorsementCount(val EndorsementCount) {
    val.Export(&ex.AvailableEndorsementCount)
}

func (ex *RecognizerExporter) SetPendingEndorsementCount(val EndorsementCount) {
    val.Export(&ex.PendingEndorsementCount)
}

func (ex *RecognizerExporter) SetVersion(val uint) {
    ex.Version = val
}

func (ex *RecognizerExporter) SetCreatedAt(val time.Time) {
    ex.CreatedAt = val
}