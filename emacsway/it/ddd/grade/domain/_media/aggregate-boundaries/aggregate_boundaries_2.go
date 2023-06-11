package grade_2

import (
    "errors"
    "time"
)

type MemberId uint64
type Grade uint
type EndorsementCount uint
type ArtifactId uint64
type ArtifactDescription string
type ArtifactStatus uint8
type CompetenceId uint64
type CompetenceName string

type Weight uint8

const (
    LowerWeight  = 0
    PeerWeight   = 1
    HigherWeight = 2

    Expert       = Grade(5)
    Candidate    = Grade(4)
    Grade1       = Grade(3)
    Grade2       = Grade(2)
    Grade3       = Grade(1)
    WithoutGrade = Grade(0)

    Proposed     = ArtifactStatus(0)
    Accepted     = ArtifactStatus(1)
)

type Specialist struct {
    id                       MemberId
    grade                    Grade
    receivedEndorsements     []Endorsement
    assignments              []Assignment
    version                  uint
    createdAt                time.Time
}

func (s *Specialist) ReceiveEndorsement(e Endorser, aId ArtifactId, t time.Time) error {
    if e.GetGrade() < s.grade {
        return errors.New(
            "it is allowed to receive endorsements only from members with equal or higher grade",
        )
    }
    if !e.CanCompleteEndorsement() {
        return errors.New(
            "endorser is not able to complete endorsement",
        )
    }
    if uint64(e.GetId()) == uint64(s.id) {
        return errors.New(
            "endorser can't endorse himself",
        )
    }
    for _, v := range s.receivedEndorsements {
        if v.IsEndorsedBy(e.GetId(), aId) {
            return errors.New("this artifact has already been endorsed by the recogniser")
        }
    }
    s.receivedEndorsements = append(s.receivedEndorsements, Endorsement{
        e.GetId(), e.GetGrade(), e.GetVersion(),
        s.id, s.grade, s.version,
        aId, t,
    })
    s.actualizeGrade(t)
    return nil
}

func (s *Specialist) actualizeGrade(t time.Time) {
    if s.grade == WithoutGrade && s.getReceivedEndorsementCount() >= 6 {
        s.setGrade(Grade3, t)
    } else if s.grade == Grade3 && s.getReceivedEndorsementCount() >= 10 {
        s.setGrade(Grade2, t)
    } else if s.grade == Grade2 && s.getReceivedEndorsementCount() >= 14 {
        s.setGrade(Grade1, t)
    } else if s.grade == Grade1 && s.getReceivedEndorsementCount() >= 20 {
        s.setGrade(Candidate, t)
    } else if s.grade == Candidate && s.getReceivedEndorsementCount() >= 40 {
        s.setGrade(Expert, t)
    }
}
func (s Specialist) getReceivedEndorsementCount() uint {
    var counter uint
    for _, v := range s.receivedEndorsements {
        if v.GetSpecialistGrade() == s.grade {
            counter += uint(v.GetWeight())
        }
    }
    return counter
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

type Endorsement struct {
    endorserId        MemberId
    endorserGrade     Grade
    endorserVersion   uint
    specialistId        MemberId
    specialistGrade     Grade
    specialistVersion   uint
    artifactId          ArtifactId
    createdAt           time.Time
}

func (e Endorsement) IsEndorsedBy(rId MemberId, aId ArtifactId) bool {
    return e.endorserId == rId && e.artifactId == aId
}

func (e Endorsement) GetSpecialistGrade() Grade {
    return e.specialistGrade
}

func (e Endorsement) GetWeight() Weight {
    if e.endorserGrade == e.specialistGrade {
        return PeerWeight
    } else if e.endorserGrade > e.specialistGrade {
        return HigherWeight
    }
    return LowerWeight
}

type Assignment struct {
    specialistId       MemberId
    specialistVersion  uint
    assignedGrade      Grade
    createdAt          time.Time
}

type Endorser struct {
    id                        MemberId
    grade                     Grade
    availableEndorsementCount EndorsementCount
    pendingEndorsementCount   EndorsementCount
    version                   uint
    createdAt                 time.Time
}

func (e Endorser) GetId() MemberId {
    return e.id
}

func (e Endorser) GetGrade() Grade {
    return e.grade
}

func (e Endorser) GetVersion() uint {
    return e.version
}

func (e Endorser) canReserveEndorsement() bool {
    return e.availableEndorsementCount > e.pendingEndorsementCount
}

func (e Endorser) CanCompleteEndorsement() bool {
    return e.pendingEndorsementCount > 0 && e.availableEndorsementCount >= e.pendingEndorsementCount
}

func (e *Endorser) ReserveEndorsement() error {
    if !e.canReserveEndorsement() {
        return errors.New("no endorsement can be reserved")
    }
    e.pendingEndorsementCount += 1
    return nil
}

func (e *Endorser) ReleaseEndorsementReservation() {
    e.pendingEndorsementCount -= 1
}

func (e *Endorser) CompleteEndorsement() error {
    if e.availableEndorsementCount == 0 {
        return errors.New("no endorsement is available")
    }
    if e.pendingEndorsementCount == 0 {
        return errors.New("there is no endorsement reservation")
    }
    e.availableEndorsementCount -= 1
    e.pendingEndorsementCount -= 1
    return nil
}

func (e *Endorser) IncreaseVersion() {
    e.version += 1
}

type Artifact struct {
    id            ArtifactId
    status        ArtifactStatus
    description   ArtifactDescription
    competenceIds []CompetenceId
    createdAt     time.Time
}

type Competence struct {
    id        CompetenceId
    name      CompetenceName
    createdAt time.Time
}
