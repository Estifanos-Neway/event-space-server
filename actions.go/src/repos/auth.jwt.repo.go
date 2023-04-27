package repos

import (
	"errors"
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
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(commons.RefreshTokenExpiresAfter)),
		},
	}
	return signJwt(claims, env.Env.JWT_REFRESH_SECRETE)
}

func signEmailVerificationToken(user types.User) (string, error) {
	claims := emailVerificationClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(commons.EmailVerificationExpiresAfter)),
		},
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

func verifyEmailVerificationToken(tokenString string) (*types.User, error) {
	token, err := jwt.ParseWithClaims(tokenString, &emailVerificationClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Env.JWT_SECRETE), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*emailVerificationClaims); ok && token.Valid {
		return &claims.User, nil
	} else {
		return nil, errors.New(invalidToken)
	}
}

func verifyRefreshToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &sessionRefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.Env.JWT_REFRESH_SECRETE), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*sessionRefreshClaims); ok && token.Valid {
		return claims.UserId, nil
	} else {
		return "", errors.New(invalidToken)
	}
}
