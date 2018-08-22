package MioAlmaDos

import (
	"fmt"
	"time"
)

type MioAlmaDos struct {
	name      String
	uniqueID  int64
	birthDate Time
	gender    String
	interests String
}

func New(name, gender string, uniqueID int64) *MioAlmaDos {
	return &MioAlmaDos{
		name:      name,
		uniqueID:  uniqueID,
		birthDate: time.Now(),
		gender:    gender,
		interests: "daydreaming, space travel",
	}
}

// Go does not support function overloading
func NewNoGender(name string, uniqueID int64) *MioAlmaDos {
	return New(name, "", uniqueID)
}

func (m MioAlmaDos) GetGender() string {
	return m.gender
}

func (m *MioAlmaDos) SetGender(gender string) {
	m.gender = gender
}

func (m MioAlmaDos) GetName() string {
	return m.name
}

func (m *MioAlmaDos) SetName(name string) {
	m.name = name
}

func (m MioAlmaDos) GetUniqueID() int64 {
	return m.uniqueID
}

func (m MioAlmaDos) GetBirthDate() Time {
	return m.birthDate
}

func (m MioAlmaDos) GetBirthDateString() string {
	return m.birthDate.format("Jan 02 2006")
}

func (m MioAlmaDos) String() string {
	return fmt.Sprintf("MioAlmaDos [name=%s, uniqueID=%d, birthDate=%s, gender=%s, interests=%s]", m.name, m.uniqueID, m.birthDate.format("Jan 02 2006"), m.gender, m.interests)
}
