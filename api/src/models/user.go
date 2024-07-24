package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/fernandomocrosky/DevBookGo/src/security"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

func (user *User) Prepare(step string) error {
	if err := user.validate(step); err != nil {
		return err
	}

	if err := user.Format(step); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("name must not be empty")
	}

	if user.Email == "" {
		return errors.New("email must not be empty")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("invalid email format")
	}

	if user.Nick == "" {
		return errors.New("nick must not be empty")
	}

	if step == "register" && user.Password == "" {
		return errors.New("password must not be empty")
	}

	return nil
}

func (user *User) Format(step string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Nick = strings.TrimSpace(user.Nick)

	if step == "register" {
		hashedKey, err := security.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashedKey)
	}

	return nil
}
