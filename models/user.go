package models

import (
	"encoding/json"
	"time"

	"github.com/dayvillefire/tenders/common"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
)

const (
	UserBitActive        = 1
	UserBitSceneSupport  = 2
	UserBitFirefighter   = 4
	UserBitInstructor    = 8
	UserBitAdministrator = 16
	UserBitSiteAdmin     = 1024
)

type User struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
	ShortCode    string     `json:"short_code" db:"shortcode"`
	PIN          string     `json:"pin" db:"pin"`
	FirstName    string     `json:"first_name" db:"firstname"`
	LastName     string     `json:"last_name" db:"lastname"`
	Department   Department `json:"department" db:"department" belongs_to:"department"`
	EmailAddress string     `json:"email" db:"email"`
	LoginEmail   string     `json:"loginemail" db:"loginemail"`
	PhoneNumber  string     `json:"phone" db:"phone"`
	BitField     int64      `json:"bitfield" db:"bitfield"`
	PictureURL   string     `json:"string" db:"picture"`
	Active       bool       `json:"active" db:"active"`
}

// UserExistsByEmail checks for an existant user
func UserExistsByEmail(email string) bool {
	var user User
	query := common.DB.Where("loginemail = ?", email).Limit(1)
	err := query.All(&user)
	if err != nil || user.ID.String() == "" {
		return false
	}
	return true
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
