package repos

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/estifanos-neway/event-space-server/src/commons"
	"github.com/estifanos-neway/event-space-server/src/env"
	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/golang-jwt/jwt/v4"
	"github.com/hasura/go-graphql-client"
)

var httpClient *http.Client = &http.Client{
	Transport: roundTripper{rt: http.DefaultTransport},
}

var gqClient = graphql.NewClient(env.Env.GRAPHQL_SERVER_URL, httpClient)

func (rt roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("x-hasura-admin-secret", env.Env.HASURA_GRAPHQL_ADMIN_SECRET)
	return rt.rt.RoundTrip(req)
}

// usersByEmail(args: {useremail: "my"}) {  email  }

func getUserByEmail(email string) (types.User, error) {
	log.Println(env.Env.GRAPHQL_SERVER_URL)
	query := usersByEmailQuery{}
	variables := map[string]interface{}{
		"useremail": email,
	}

	if err := gqClient.Query(context.Background(), &query, variables); err != nil {
		return query.UsersByEmail, err
	}

	return query.UsersByEmail, nil
}

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
