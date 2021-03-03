package models

import (
	"encoding/json"
	"fmt"
	"log"
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
	ID           uuid.UUID `json:"id" db:"id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
	ShortCode    string    `json:"short_code" db:"shortcode"`
	PIN          string    `json:"pin" db:"pin"`
	FirstName    string    `json:"first_name" db:"firstname"`
	LastName     string    `json:"last_name" db:"lastname"`
	DepartmentID uuid.UUID `json:"department" db:"department"`
	EmailAddress string    `json:"email" db:"email"`
	LoginEmail   string    `json:"loginemail" db:"loginemail"`
	PhoneNumber  string    `json:"phone" db:"phone"`
	BitField     int64     `json:"bitfield" db:"bitfield"`
	PictureURL   string    `json:"string" db:"picture"`
	Active       bool      `json:"active" db:"active"`
}

// UserExistsByEmail checks for an existant user
func UserExistsByEmail(email string) bool {
	var user User
	query := common.DB.Where("loginemail = ?", email).Limit(1)
	err := query.First(&user)
	if err != nil || user.ID.String() == "" {
		return false
	}
	return true
}

// UserByShortCode returns a user looked up by their shortcode
func UserByShortCode(shortcode string) (User, error) {
	var user User
	query := common.DB.Where("shortcode = ?", shortcode).Limit(1)
	err := query.First(&user)
	return user, err
}

// UserAuth returns a user looked up by their shortcode and pin
func UserAuth(shortcode string, pin string) (User, error) {
	var user User
	query := common.DB.Where("shortcode = ? AND pin = ?", shortcode, pin).Limit(1)
	err := query.First(&user)
	if err != nil {
		log.Printf("UserAuth: ERR: %s", err.Error())
	}
	if err == nil && user.PIN == "" {
		err = fmt.Errorf("user or PIN incorrect")
	}
	return user, err
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

// NewShortCode generates a new short code that is not in use
func NewShortCode() string {
	shortCode := common.ShortCode(4)
	// TODO: FIXME: XXX: Need to do collision checking
	return shortCode
}
