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

func (s *Specialist) ReceiveEndorsement(r Recognizer, aId ArtifactId, t time.Time) error {
    if r.GetGrade() < s.grade {
        return errors.New(
            "it is allowed to receive endorsements only from members with equal or higher grade",
        )
    }
    if !r.CanCompleteEndorsement() {
        return errors.New(
            "recognizer is not able to complete endorsement",
        )
    }
    if uint64(r.GetId()) == uint64(s.id) {
        return errors.New(
            "recognizer can't endorse himself",
        )
    }
    for _, v := range s.receivedEndorsements {
        if v.IsEndorsedBy(r.GetId(), aId) {
            return errors.New("this artifact has already been endorsed by the recogniser")
        }
    }
    s.receivedEndorsements = append(s.receivedEndorsements, Endorsement{
        r.GetId(), r.GetGrade(), r.GetVersion(),
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
    recognizerId        MemberId
    recognizerGrade     Grade
    recognizerVersion   uint
    specialistId        MemberId
    specialistGrade     Grade
    specialistVersion   uint
    artifactId          ArtifactId
    createdAt           time.Time
}

func (e Endorsement) IsEndorsedBy(rId MemberId, aId ArtifactId) bool {
    return e.recognizerId == rId && e.artifactId == aId
}

func (e Endorsement) GetSpecialistGrade() Grade {
    return e.specialistGrade
}

func (e Endorsement) GetWeight() Weight {
    if e.recognizerGrade == e.specialistGrade {
        return PeerWeight
    } else if e.recognizerGrade > e.specialistGrade {
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

type Recognizer struct {
    id                        MemberId
    grade                     Grade
    availableEndorsementCount EndorsementCount
    pendingEndorsementCount   EndorsementCount
    version                   uint
    createdAt                 time.Time
}

func (r Recognizer) GetId() MemberId {
    return r.id
}

func (r Recognizer) GetGrade() Grade {
    return r.grade
}

func (r Recognizer) GetVersion() uint {
    return r.version
}

func (r Recognizer) canReserveEndorsement() bool {
    return r.availableEndorsementCount > r.pendingEndorsementCount
}

func (r Recognizer) CanCompleteEndorsement() bool {
    return r.pendingEndorsementCount > 0 && r.availableEndorsementCount >= r.pendingEndorsementCount
}

func (r *Recognizer) ReserveEndorsement() error {
    if !r.canReserveEndorsement() {
        return errors.New("no endorsement can be reserved")
    }
    r.pendingEndorsementCount += 1
    return nil
}

func (r *Recognizer) ReleaseEndorsementReservation() {
    r.pendingEndorsementCount -= 1
}

func (r *Recognizer) CompleteEndorsement() error {
    if r.availableEndorsementCount == 0 {
        return errors.New("no endorsement is available")
    }
    if r.pendingEndorsementCount == 0 {
        return errors.New("there is no endorsement reservation")
    }
    r.availableEndorsementCount -= 1
    r.pendingEndorsementCount -= 1
    return nil
}

func (r *Recognizer) IncreaseVersion() {
    r.version += 1
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
