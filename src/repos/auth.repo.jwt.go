package repos

import (
	"time"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/env"
	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/golang-jwt/jwt/v4"
)

func signAccessToken(userId string) (string, error) {
	claims := loginClaims{
		httpsHasuraIoJwtClaims{
			XHasuraAllowedRoles: []string{"user"},
			XHasuraDefaultRole:  "user",
			XHasuraUserId:       userId,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(commons.AccessTokenExpiresAfter)),
		},
	}
	return signJwt(claims, env.Env.JWT_SECRETE)
}

func signRefreshToken(userId string) (string, error) {
	claims := sessionRefreshClaims{
		userId,
		jwt.RegisteredClaims{},
	}
	return signJwt(claims, env.Env.JWT_REFRESH_SECRETE)
}

func signEmailVerificationToken(user types.User) (string, error) {
	claims := emailVerificationClaims{
		user,
		jwt.RegisteredClaims{},
	}
	return signJwt(claims, env.Env.JWT_SECRETE)
}

func signJwt(claims jwt.Claims, key string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(key))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
