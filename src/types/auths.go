package types

import (
	"errors"

	"github.com/estifanos-neway/event-space-server/src/commons"
)

type SignInInput struct {
	Email    string
	Password string
}

func (this SignInInput) IsValidSignInInput() error {
	email, err := commons.ParseEmail(this.Email)
	if err != nil {
		return errors.New("Invalid_Email")
	}
	this.Email = email
	if len(this.Password) == commons.MinPasswordLength {
		return errors.New("Invalid_Password")
	}
	return nil
}

type SignUpInput struct {
	Name string
	SignInInput
}

func (this SignUpInput) IsValid() error {
	if err := this.IsValidSignInInput(); err != nil {
		return err
	}
	if len(this.Name) == 0 {
		return errors.New("Invalid_Name")
	}
	return nil
}
