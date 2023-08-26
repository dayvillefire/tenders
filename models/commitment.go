package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

type Commitment struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	EventID   uuid.UUID `db:"event_id"`
	Event     *Event    `json:"event" belongs_to:"event"`
	UserID    uuid.UUID `db:"user_id"`
	User      *User     `json:"user" belongs_to:"user"`
	Attending bool      `json:"attending" db:"attending"`
	Note      string    `json:"note" db:"note"`
}

// String is not required by pop and may be deleted
func (c Commitment) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Commitments is not required by pop and may be deleted
type Commitments []Commitment

// String is not required by pop and may be deleted
func (c Commitments) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Commitment) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Note, Name: "Note"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Commitment) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Commitment) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
