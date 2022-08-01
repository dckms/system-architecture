package qualifying_grade_2

import (
    "errors"
    "time"
)

type EndorsedId uint64
type MemberId uint64
type Grade uint
type EndorsementCount uint
type RecognizerId uint64
type ArtifactId uint64
type ArtifactDescription string
type ArtifactStatus uint8
type ExpertiseAreaId uint64
type ExpertiseAreaName string

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

type Endorsed struct {
    id                       EndorsedId
    memberId                 MemberId
    grade                    Grade
    receivedEndorsements     []Endorsement
    assignments              []Assignment
    version                  uint
    createdAt                time.Time
}

func (e *Endorsed) ReceiveEndorsement(r Recognizer, aId ArtifactId, t time.Time) error {
    if r.GetGrade() < e.grade {
        return errors.New(
            "it is allowed to receive endorsements only from members with equal or higher grade",
        )
    }
    if !r.CanCompleteEndorsement() {
        return errors.New(
            "recognizer is not able to complete endorsement",
        )
    }
    if uint64(r.GetId()) == uint64(e.id) {
        return errors.New(
            "recognizer can't endorse himself",
        )
    }
    for _, v := range e.receivedEndorsements {
        if v.IsEndorsedBy(r.GetId(), aId) {
            return errors.New("this artifact has already been endorsed by the recogniser")
        }
    }
    e.receivedEndorsements = append(e.receivedEndorsements, Endorsement{
        r.GetId(), r.GetGrade(), r.GetVersion(),
        e.id, e.grade, e.version,
        aId, t,
    })
    e.actualizeGrade(t)
    return nil
}

func (e *Endorsed) actualizeGrade(t time.Time) {
    if e.grade == WithoutGrade && e.getReceivedEndorsementCount() >= 6 {
        e.setGrade(Grade3, t)
    } else if e.grade == Grade3 && e.getReceivedEndorsementCount() >= 10 {
        e.setGrade(Grade2, t)
    } else if e.grade == Grade2 && e.getReceivedEndorsementCount() >= 14 {
        e.setGrade(Grade1, t)
    } else if e.grade == Grade1 && e.getReceivedEndorsementCount() >= 20 {
        e.setGrade(Candidate, t)
    } else if e.grade == Candidate && e.getReceivedEndorsementCount() >= 40 {
        e.setGrade(Expert, t)
    }
}
func (e Endorsed) getReceivedEndorsementCount() uint {
    var counter uint
    for _, v := range e.receivedEndorsements {
        if v.GetEndorsedGrade() == e.grade {
            counter += uint(v.GetWeight())
        }
    }
    return counter
}

func (e *Endorsed) setGrade(g Grade, t time.Time) {
    e.assignments = append(e.assignments, Assignment{
        e.id, e.version, g, t,
    })
    e.grade = g
}

func (e *Endorsed) IncreaseVersion() {
    e.version += 1
}

type Endorsement struct {
    recognizerId        RecognizerId
    recognizerGrade     Grade
    recognizerVersion   uint
    endorsedId          EndorsedId
    endorsedGrade       Grade
    endorsedVersion     uint
    artifactId          ArtifactId
    createdAt           time.Time
}

func (e Endorsement) IsEndorsedBy(rId RecognizerId, aId ArtifactId) bool {
    return e.recognizerId == rId && e.artifactId == aId
}

func (e Endorsement) GetEndorsedGrade() Grade {
    return e.endorsedGrade
}

func (e Endorsement) GetWeight() Weight {
    if e.recognizerGrade == e.endorsedGrade {
        return PeerWeight
    } else if e.recognizerGrade > e.endorsedGrade {
        return HigherWeight
    }
    return LowerWeight
}

type Assignment struct {
    endorsedId          EndorsedId
    endorsedVersion     uint
    assignedGrade       Grade
    createdAt           time.Time
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

func (r Recognizer) GetId() RecognizerId {
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
    id                  ArtifactId
    artifactDescription ArtifactDescription
    status              ArtifactStatus
    createdAt           time.Time
}

type ExpertiseArea struct {
    id        ExpertiseAreaId
    name      ExpertiseAreaName
    createdAt time.Time
}
