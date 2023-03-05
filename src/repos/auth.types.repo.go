package repos

import (
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/golang-jwt/jwt/v4"
)

type usersByEmailQuery struct {
	UsersByEmail types.User `graphql:"usersByEmail(args:{useremail:$useremail})"`
}
type usersById struct {
	UsersById types.User `graphql:"usersByPk(id:$id)"`
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

type UserLogin struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type uuid string
