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

	WithoutGrade = Grade(0)
	Grade3       = Grade(1)
	Grade2       = Grade(2)
	Grade1       = Grade(3)
	Candidate    = Grade(4)
	Expert       = Grade(5)

	Proposed     = ArtifactStatus(0)
	Accepted     = ArtifactStatus(1)
)

type Endorsed struct {
	id                       EndorsedId
	memberId                 MemberId
	grade                    Grade
	receivedEndorsements     []Endorsement
	gradeLogEntries          []GradeLogEntry
	version                  int
	createdAt                time.Time
}

func (e *Endorsed) ReceiveEndorsement(r Recognizer, a ArtifactId, t time.Time) error {
	if r.GetGrade() < e.grade {
		return errors.New(
			"it is allowed to receive endorsements only from members with equal or higher grade",
		)
	}
	if !r.CanEndorse() {
		return errors.New(
			"recognizer is not able to endorse",
		)
	}
	e.receivedEndorsements = append(e.receivedEndorsements, Endorsement{
		r.GetId(), r.GetGrade(), r.GetVersion(),
		e.id, e.grade, e.version,
		a, t,
	})
	return nil
}

func (e *Endorsed) IncreaseReceivedEndorsementCount(dt time.Time) {
	if e.grade == WithoutGrade && e.getReceivedEndorsementCount() >= 6 {
		e.setGrade(Grade3, dt)
	} else if e.grade == Grade3 && e.getReceivedEndorsementCount() >= 10 {
		e.setGrade(Grade2, dt)
	} else if e.grade == Grade2 && e.getReceivedEndorsementCount() >= 14 {
		e.setGrade(Grade1, dt)
	} else if e.grade == Grade1 && e.getReceivedEndorsementCount() >= 20 {
		e.setGrade(Candidate, dt)
	} else if e.grade == Candidate && e.getReceivedEndorsementCount() >= 10 {
		e.setGrade(Expert, dt)
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

func (e *Endorsed) setGrade(g Grade, dt time.Time) {
	e.gradeLogEntries = append(e.gradeLogEntries, GradeLogEntry{
		e.id, e.version, g, dt,
	})
	e.grade = g
}

func (e *Endorsed) IncreaseVersion() {
	e.version += 1
}

type Endorsement struct {
	recognizerId        RecognizerId
	recognizerGrade     Grade
	recognizerVersion   int
	endorsedId          EndorsedId
	endorsedGrade       Grade
	endorsedVersion     int
	artifactId          ArtifactId
	createdAt           time.Time
}

func (e Endorsement) GetEndorsedId() EndorsedId {
	return e.endorsedId
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

type GradeLogEntry struct {
	endorsedId          EndorsedId
	endorsedVersion     int
	assignedGrade       Grade
	createdAt           time.Time
}

type Recognizer struct {
	id                        RecognizerId
	memberId                  MemberId
	grade                     Grade
	availableEndorsementCount EndorsementCount
	pendingEndorsementCount   EndorsementCount
	version                   int
	createdAt                 time.Time
}

func (r Recognizer) GetId() RecognizerId {
	return r.id
}

func (r Recognizer) GetGrade() Grade {
	return r.grade
}

func (r Recognizer) GetVersion() int {
	return r.version
}

func (r Recognizer) CanEndorse() bool {
	return r.availableEndorsementCount - r.pendingEndorsementCount >= 0
}

func (r *Recognizer) DecreaseAvailableEndorsementCount() error {
	if r.availableEndorsementCount == 0 {
		return errors.New("no endorsement is available")
	}
	r.availableEndorsementCount -= 1
	return nil
}

func (r *Recognizer) IncreasePendingEndorsementCount() {
	r.pendingEndorsementCount += 1
}

func (r *Recognizer) DecreasePendingEndorsementCount() error {
	if r.pendingEndorsementCount >= r.availableEndorsementCount {
		return errors.New("can't reserve an endorsement")
	}
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