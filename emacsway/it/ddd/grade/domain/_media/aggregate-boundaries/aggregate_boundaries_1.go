package grade_1

import (
    "errors"
    "time"
)

type MemberId uint64
type Grade uint
type AvailableEndorsementCount uint
type ReceivedEndorsementCount uint
type EndorsementId uint64
type ArtifactDescription string

type Weight uint8

const (
    PeerWeight   = 1
    HigherWeight = 2
)

const (
    Expert       = Grade(5)
    Candidate    = Grade(4)
    Grade1       = Grade(3)
    Grade2       = Grade(2)
    Grade3       = Grade(1)
    WithoutGrade = Grade(0)
)

type Specialist struct {
    id                       MemberId
    grade                    Grade
    receivedEndorsementCount ReceivedEndorsementCount
    assignments              []Assignment
    version                  uint
    createdAt                time.Time
}

func (s Specialist) GetId() MemberId {
    return s.id
}

func (s Specialist) GetGrade() Grade {
    return s.grade
}

func (s Specialist) GetVersion() uint {
    return s.version
}

func (s *Specialist) IncreaseReceivedEndorsementCount(w Weight, t time.Time) {
    s.receivedEndorsementCount += ReceivedEndorsementCount(w)
    if s.grade == WithoutGrade && s.receivedEndorsementCount >= 6 {
        s.setGrade(Grade3, t)
        s.receivedEndorsementCount = 0
    } else if s.grade == Grade3 && s.receivedEndorsementCount >= 10 {
        s.setGrade(Grade2, t)
        s.receivedEndorsementCount = 0
    } else if s.grade == Grade2 && s.receivedEndorsementCount >= 14 {
        s.setGrade(Grade1, t)
        s.receivedEndorsementCount = 0
    } else if s.grade == Grade1 && s.receivedEndorsementCount >= 20 {
        s.setGrade(Candidate, t)
        s.receivedEndorsementCount = 0
    } else if s.grade == Candidate && s.receivedEndorsementCount >= 40 {
        s.setGrade(Expert, t)
        s.receivedEndorsementCount = 0
    }
}

func (s *Specialist) setGrade(g Grade, t time.Time) {
    s.assignments = append(s.assignments, Assignment{
        s.id, s.version, g, t,
    })
    s.grade = g
}

func (s *Specialist) IncreaseVersion() {
    s.version += 1
}

type Assignment struct {
    specialistId        MemberId
    specialistVersion   uint
    assignedGrade       Grade
    createdAt           time.Time
}

type Endorser struct {
    id                        MemberId
    grade                     Grade
    availableEndorsementCount AvailableEndorsementCount
    version                   uint
    createdAt                 time.Time
}

func (e Endorser) Endorse(s Specialist, aDesc ArtifactDescription, t time.Time) (Endorsement, error) {
    if e.grade < s.grade {
        return Endorsement{}, errors.New(
            "it is allowed to endorse only members with equal or lower grade",
        )
    }
    if e.availableEndorsementCount == 0 {
        return Endorsement{}, errors.New(
            "you have reached the limit of available recommendations this year",
        )
    }
    if uint64(e.id) == uint64(s.GetId()) {
        return Endorsement{}, errors.New(
            "endorser can't endorse himself",
        )
    }
    return Endorsement{
        e.id, e.grade, e.version,
        s.id, s.grade, s.GetVersion(),
        aDesc, t,
    }, nil
}

func (e *Endorser) DecreaseAvailableEndorsementCount() error {
    if e.availableEndorsementCount == 0 {
        return errors.New("no endorsement is available")
    }
    e.availableEndorsementCount -= 1
    return nil
}

func (e *Endorser) IncreaseVersion() {
    e.version += 1
}

type Endorsement struct {
    endorserId        MemberId
    endorserGrade     Grade
    endorserVersion   uint
    specialistId        MemberId
    specialistGrade     Grade
    specialistVersion   uint
    artifactDescription ArtifactDescription
    createdAt           time.Time
}
