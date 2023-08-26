package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

type Event struct {
	ID                  uuid.UUID   `json:"id" db:"id"`
	CreatedAt           time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at" db:"updated_at"`
	Date                time.Time   `json:"date" db:"dateof"`
	BeginTime           int         `json:"begin_hour" db:"beginhour"`
	EndTime             int         `json:"end_hour" db:"endhour"`
	DepartmentID        uuid.UUID   `db:"department_id"`
	Department          *Department `json:"department" belongs_to:"department"`
	LocationID          uuid.UUID   `db:"location_id"`
	Location            *Location   `json:"location" belongs_to:"location"`
	StaffingRequirement string      `json:"staffing_required" db:"staffingreq"`
	StaffingOptional    string      `json:"staffing_optional" db:"staffingopt"`
	Active              bool        `json:"active" db:"active"`
}

// String is not required by pop and may be deleted
func (e Event) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Events is not required by pop and may be deleted
type Events []Event

// String is not required by pop and may be deleted
func (e Events) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Event) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
