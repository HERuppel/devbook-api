package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("NAME_REQUIRED")
	}
	if user.Nick == "" {
		return errors.New("NICK_REQUIRED")
	}
	if user.Email == "" {
		return errors.New("EMAIL_REQUIRED")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("INAVALID_FORM_EMAIL")
	}

	if step == "CREATE" && user.Password == "" {
		return errors.New("PASSWORD_REQUIRED")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
