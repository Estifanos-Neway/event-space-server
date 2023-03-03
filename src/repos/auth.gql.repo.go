package repos

import (
	"context"
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/env"
	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/hasura/go-graphql-client"
)

const GRAPHQL_SERVER_URL string = "http://localhost:8000/v1/graphql"

var httpClient *http.Client = &http.Client{
	Transport: roundTripper{rt: http.DefaultTransport},
}

var gqClient = graphql.NewClient(GRAPHQL_SERVER_URL, httpClient)

func (rt roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("x-hasura-admin-secret", env.Env.HASURA_GRAPHQL_ADMIN_SECRET)
	return rt.rt.RoundTrip(req)
}

func getUserByEmail(email string) (types.User, error) {
	query := usersByEmailQuery{}
	variables := map[string]interface{}{
		"useremail": email,
	}

	if err := gqClient.Query(context.Background(), &query, variables); err != nil {
		return query.UsersByEmail, err
	}

	return query.UsersByEmail, nil
}

func insertUser(user types.User) (*types.User, error) {
	mutation := insertUserMutation{}
	variables := map[string]interface{}{
		"email":        user.Email,
		"name":         user.Name,
		"passwordHash": user.PasswordHash,
	}

	if err := gqClient.Mutate(context.Background(), &mutation, variables); err != nil {
		return &mutation.InsertUsersOne.User, err
	}
	return &mutation.InsertUsersOne.User, nil
}

type insertUserMutation struct {
	InsertUsersOne struct {
		types.User
	} `graphql:"insertUsersOne(object:{email:$email,name:$name,passwordHash:$passwordHash })"`
}

func insertSessionRefreshToken(token string) error {
	mutation := insertSessionRefreshTokenMutation{}
	variables := map[string]interface{}{
		"token": token,
	}
	if err := gqClient.Mutate(context.Background(), &mutation, variables); err != nil {
		return err
	}
	return nil
}

type insertSessionRefreshTokenMutation struct {
	InsertSessionRefreshTokensOne struct {
		Id string
	} `graphql:"insertSessionRefreshTokensOne(object:{token:$token })"`
}

type sessionRefreshTokensQuery struct {
	SessionRefreshTokens []struct {
		Id string
	} `graphql:"sessionRefreshTokens(where:{token:{_eq:$token}},limit: 1)"`
}

func refreshTokenExists(token string) (bool, error) {
	query := sessionRefreshTokensQuery{}
	variables := map[string]interface{}{
		"token": token,
	}
	if err := gqClient.Query(context.Background(), &query, variables); err != nil {
		return false, err
	}
	return query.SessionRefreshTokens[0].Id != "", nil
}
