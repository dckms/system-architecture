package qualifying_grade_1

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

func (e ExportableUint) ExportTo(ex Exporter[uint]) {
    ex.SetState(uint(e))
}

type RecognizerId ExportableUint
type MemberId ExportableUint
type Grade ExportableUint
type EndorsementCount ExportableUint

type RecognizerExporter interface {
    SetState(
        id Exporter[uint],
        memberId Exporter[uint],
        grade Exporter[uint],
        availableEndorsementCount Exporter[uint],
        pendingEndorsementCount Exporter[uint],
        version uint,
        createdAt time.Time,
    )
}

type UintExporter uint

func (e *UintExporter) SetState(value uint) {
    *e = UintExporter(value)
}

type Recognizer struct {
    id                        RecognizerId
    memberId                  MemberId
    grade                     Grade
    availableEndorsementCount EndorsementCount
    pendingEndorsementCount   EndorsementCount
    version                   uint
    createdAt                 time.Time
}

func (r Recognizer) ExportTo(ex RecognizerExporter) {
    var id, memberId UintExporter
    var grade, availableEndorsementCount, pendingEndorsementCount UintExporter

    r.id.ExportTo(&id)
    r.memberId.ExportTo(&memberId)
    r.grade.ExportTo(&grade)
    r.availableEndorsementCount.ExportTo(&availableEndorsementCount)
    r.pendingEndorsementCount.ExportTo(&pendingEndorsementCount)
    ex.SetState(
        &id, &memberId, &grade, &availableEndorsementCount,
        &pendingEndorsementCount, r.version, r.createdAt,
    )
}
