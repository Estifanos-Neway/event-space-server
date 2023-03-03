package repos

import (
	"log"

	"github.com/estifanos-neway/event-space-server/src/commons"
	types "github.com/estifanos-neway/event-space-server/src/types"
)

func SignInRepo(signInInput types.SignInInput) (int, *UserLogin, string) {
	if err := signInInput.IsValidSignInInput(); err != nil {
		return 400, nil, err.Error()
	}
	passwordHash := commons.Hash(signInInput.Password)
	user, err := getUserByEmail(signInInput.Email)
	if err != nil {
		log.Println("usersByEmail", err)
		return 500, nil, InternalError
	} else if user.PasswordHash != passwordHash {
		return 404, nil, commons.Not_Found
	}

	var accessToken, refreshToken string
	accessToken, err = signAccessToken(user.Id)
	if err != nil {
		log.Println("signAccessToken", err)
		return 500, nil, InternalError
	}
	refreshToken, err = signRefreshToken(user.Id)
	if err != nil {
		log.Println("signRefreshToken", err)
		return 500, nil, InternalError
	}
	userLogin := UserLogin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	if err = insertSessionRefreshToken(refreshToken); err != nil {
		log.Println("insertSessionRefreshToken", err)
		return 500, nil, InternalError
	}
	return 200, &userLogin, ""
}