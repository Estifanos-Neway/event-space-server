package repos

import "log"

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
			return 500, nil, InternalError
		}
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
