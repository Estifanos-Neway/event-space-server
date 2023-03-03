package repos

import (
	"crypto/sha256"
	"log"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/env"
	types "github.com/estifanos-neway/event-space-server/src/types"
)

func SignupRepo(signUpInput types.SignUpInput) types.SimpleResponse {
	// validate correctness
	if err := signUpInput.IsValidSignInInput(); err != nil {
		return types.SimpleResponse{
			Code:    400,
			Message: err.Error(),
		}
	}
	// validate uniqueness
	if existingUser, err := getUserByEmail(signUpInput.Email); err != nil {
		log.Println("usersByEmail", err)
		return types.SimpleResponse{
			Code:    500,
			Message: commons.InternalError,
		}
	} else if existingUser.Email == signUpInput.Email {
		return types.SimpleResponse{
			Code:    400,
			Message: commons.InternalError,
		}
	}
	// send email
	passwordHash := sha256.Sum256([]byte(signUpInput.Password))
	user := types.User{
		Email:        signUpInput.Email,
		Name:         signUpInput.Name,
		PasswordHash: string(passwordHash[:]),
	}
	emailVerificationToken, err := signEmailVerificationToken(user)
	if err != nil {
		log.Println("signEmailVerificationToken", err)
		return types.SimpleResponse{
			Code:    500,
			Message: commons.InternalError,
		}
	}
	subject := "Email Verification"
	content := env.Env.EMAIL_VERIFICATION_URL + emailVerificationToken
	if err := commons.SendEmail(signUpInput.Email, &content, nil, nil, &subject, nil); err != nil {
		log.Println("SendEmail", err)
		return types.SimpleResponse{
			Code:    500,
			Message: commons.InternalError,
		}
	}
	// return token
	return types.SimpleResponse{
		Code:    200,
		Message: commons.Ok,
	}
}
