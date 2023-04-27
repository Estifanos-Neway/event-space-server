package types

import (
	"errors"
	"fmt"

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
	if len(this.Password) < commons.MinPasswordLength {
		return errors.New(fmt.Sprintf("Password_Must_Be_At_Least_%v_Characters", commons.MinPasswordLength))
	}
	return nil
}

type SignUpInput struct {
	Name string
	SignInInput
}

func (this SignUpInput) IsValidSignUpInput() error {
	if err := this.IsValidSignInInput(); err != nil {
		return err
	}
	if len(this.Name) == 0 {
		return errors.New("Invalid_Name")
	}
	return nil
}
