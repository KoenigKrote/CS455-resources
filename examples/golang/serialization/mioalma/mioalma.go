package MioAlma

import (
	"fmt"
	"time"
)

type MioAlma struct {
	name      string
	uniqueID  int64
	birthDate Time
	interests string
}

func New(name string, uniqueID int64) *MioAlma {
	return &MioAlma{
		name:      name,
		uniqueID:  uniqueID,
		birthDate: time.Now(),
		interests: "daydreaming",
	}
}

func (m MioAlma) GetName() {
	return m.name
}

func (m *MioAlma) SetName(name string) {
	m.name = name
}

func (m MioAlma) GetUniqueID() int64 {
	return m.uniqueID
}

func (m MioAlma) GetBirthDate() Time {
	return m.birthDate
}

func (m MioAlma) GetBirthDateString() string {
	return m.birthDate.format("Jan 2 2006")
}

func (m MioAlma) String() string {
	return fmt.Sprintf("MioAlma [name=%s, uniqueID=%d, birthDate=%s, interests=%s]", m.name, m.uniqueID, m.birthDate.format("Jan 2 2006"), m.interests)
}
