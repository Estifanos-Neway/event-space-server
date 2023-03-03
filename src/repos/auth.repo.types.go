package repos

import (
	"errors"
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/env"
	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/golang-jwt/jwt/v4"
)

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
		return nil, errors.New(commons.Invalid_Token)
	}
}

type usersByEmailQuery struct {
	UsersByEmail types.User `graphql:"usersByEmail(args:{useremail:$useremail})"`
}

type roundTripper struct {
	rt http.RoundTripper
}

type httpsHasuraIoJwtClaims struct {
	XHasuraAllowedRoles []string `json:"x-hasura-allowed-roles"`
	XHasuraDefaultRole  string   `json:"x-hasura-default-role"`
	XHasuraUserId       string   `json:"X-Hasura-User-Id"`
}

type loginClaims struct {
	HttpsHasuraIoJwtClaims httpsHasuraIoJwtClaims `json:"https://hasura.io/jwt/claims"`
	jwt.RegisteredClaims
}

type sessionRefreshClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

type emailVerificationClaims struct {
	User types.User `json:"user"`
	jwt.RegisteredClaims
}
