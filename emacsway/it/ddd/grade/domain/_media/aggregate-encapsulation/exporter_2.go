package grade_2

import (
    "time"
)

type Exportable[T any] interface {
    Export() T
}

type ExportableUint uint

func (e ExportableUint) Export() uint {
    return uint(e)
}

type MemberId ExportableUint
type Grade ExportableUint
type EndorsementCount ExportableUint

type EndorserState struct {
    Id                        uint
    Grade                     uint
    AvailableEndorsementCount uint
    PendingEndorsementCount   uint
    Version                   uint
    CreatedAt                 time.Time
}

type Endorser struct {
    id                        MemberId
    grade                     Grade
    availableEndorsementCount EndorsementCount
    pendingEndorsementCount   EndorsementCount
    version                   uint
    createdAt                 time.Time
}

func (e Endorser) Export() EndorserState {
    return EndorserState{
        e.id.Export(), e.grade.Export(),
        e.availableEndorsementCount.Export(),
        e.pendingEndorsementCount.Export(),
        e.version, e.createdAt,
    }
}
