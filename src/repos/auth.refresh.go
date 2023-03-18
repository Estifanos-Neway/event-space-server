package repos

import (
	"log"

	"github.com/estifanos-neway/event-space-server/src/commons"
)

func RefreshRepo(refreshToken string) (int, *UserLogin, string) {
	userId, err := verifyRefreshToken(refreshToken)
	if err != nil {
		return 400, nil, invalidToken
	}

	var tokenExist bool
	tokenExist, err = refreshTokenExists(refreshToken)
	if err != nil {
		log.Println("refreshTokenExists", err)
		return 500, nil, commons.InternalError
	}
	if !tokenExist {
		return 400, nil, invalidToken
	}

	var accessToken string
	accessToken, err = signAccessToken(userId)
	if err != nil {
		log.Println("signAccessToken", err)
		return 500, nil, commons.InternalError
	}
	userLogin := UserLogin{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return 200, &userLogin, commons.Ok
}
