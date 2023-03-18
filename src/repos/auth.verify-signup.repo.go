package repos

import (
	"log"

	"github.com/estifanos-neway/event-space-server/src/commons"
)

func VerifySignupRepo(verificationToken string) (int, *UserLogin, string) {
	// verify
	user, err := verifyEmailVerificationToken(verificationToken)
	if err != nil {
		return 400, nil, invalidToken
	}

	// insert user
	user, err = insertUser(*user)
	if err != nil {
		if err.Error() != emailDuplicationMessage {
			return 409, nil, emailAlreadyExist
		} else {
			log.Println("insertUser", err)
			return 500, nil, commons.InternalError
		}
	}

	var accessToken, refreshToken string
	accessToken, err = signAccessToken(user.Id)
	if err != nil {
		log.Println("signAccessToken", err)
		return 500, nil, commons.InternalError
	}
	refreshToken, err = signRefreshToken(user.Id)
	if err != nil {
		log.Println("signRefreshToken", err)
		return 500, nil, commons.InternalError
	}
	userLogin := UserLogin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err = insertSessionRefreshToken(refreshToken, uuid(user.Id)); err != nil {
		log.Println("insertSessionRefreshToken", err)
		return 500, nil, commons.InternalError
	}
	return 200, &userLogin, commons.Ok
}
