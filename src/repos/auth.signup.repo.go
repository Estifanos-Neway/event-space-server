package repos

import (
	"log"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/env"
	types "github.com/estifanos-neway/event-space-server/src/types"
)

func SignUpRepo(signUpInput types.SignUpInput) (int, string) {
	// validate correctness
	if err := signUpInput.IsValidSignUpInput(); err != nil {
		return 400, err.Error()
	}
	// validate uniqueness
	if existingUser, err := getUserByEmail(signUpInput.Email); err != nil {
		log.Println("usersByEmail", err)
		return 500, commons.InternalError
	} else if existingUser.Email == signUpInput.Email {
		return 409, emailAlreadyExist
	}
	// send email
	passwordHash := commons.Hash(signUpInput.Password)
	user := types.User{
		Email:        signUpInput.Email,
		Name:         signUpInput.Name,
		PasswordHash: string(passwordHash),
	}
	emailVerificationToken, err := signEmailVerificationToken(user)
	if err != nil {
		log.Println("signEmailVerificationToken", err)
		return 500, commons.InternalError
	}
	subject := "Email Verification"
	content := env.Env.EMAIL_VERIFICATION_URL + emailVerificationToken
	if err := commons.SendEmail(signUpInput.Email, &content, nil, nil, &subject, nil); err != nil {
		log.Println("SendEmail", err)
		return 500, commons.InternalError
	}
	// return token
	return 200, commons.Ok
}
