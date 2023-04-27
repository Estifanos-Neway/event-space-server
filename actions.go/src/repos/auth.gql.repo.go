package repos

import (
	"context"
	"net/http"

	"github.com/estifanos-neway/event-space-server/src/env"
	"github.com/estifanos-neway/event-space-server/src/types"
	"github.com/hasura/go-graphql-client"
)

var httpClient *http.Client = &http.Client{
	Transport: roundTripper{rt: http.DefaultTransport},
}

// var gqClient = graphql.NewClient(env.Env.GRAPHQL_SERVER_URL, httpClient)

func getGqClient() *graphql.Client {
	return graphql.NewClient(env.Env.GRAPHQL_SERVER_URL, httpClient)
}

func (rt roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("x-hasura-admin-secret", env.Env.HASURA_GRAPHQL_ADMIN_SECRET)
	return rt.rt.RoundTrip(req)
}

func getUserByEmail(email string) (types.User, error) {
	var gqClient = getGqClient()
	query := usersByEmailQuery{}
	variables := map[string]interface{}{
		"useremail": email,
	}

	if err := gqClient.Query(context.Background(), &query, variables); err != nil {
		return query.UsersByEmail, err
	}

	return query.UsersByEmail, nil
}

func getUserById(id uuid) (types.User, error) {
	var gqClient = getGqClient()
	query := usersById{}
	variables := map[string]interface{}{
		"id": id,
	}

	if err := gqClient.Query(context.Background(), &query, variables); err != nil {
		return query.UsersById, err
	}

	return query.UsersById, nil
}

func insertUser(user types.User) (*types.User, error) {
	var gqClient = getGqClient()
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

func insertSessionRefreshToken(token string, userId uuid) error {
	var gqClient = getGqClient()
	mutation := insertSessionRefreshTokenMutation{}
	variables := map[string]interface{}{
		"token":  token,
		"userId": userId,
	}
	if err := gqClient.Mutate(context.Background(), &mutation, variables); err != nil {
		return err
	}
	return nil
}

type insertSessionRefreshTokenMutation struct {
	InsertSessionRefreshTokensOne struct {
		Id string
	} `graphql:"insertSessionRefreshTokensOne(object:{token:$token, userId:$userId })"`
}

type sessionRefreshTokensQuery struct {
	SessionRefreshTokens []struct {
		Id string
	} `graphql:"sessionRefreshTokens(where:{token:{_eq:$token}},limit: 1)"`
}

func refreshTokenExists(token string) (bool, error) {
	var gqClient = getGqClient()
	query := sessionRefreshTokensQuery{}
	variables := map[string]interface{}{
		"token": token,
	}
	if err := gqClient.Query(context.Background(), &query, variables); err != nil {
		return false, err
	}
	return len(query.SessionRefreshTokens) > 0 && query.SessionRefreshTokens[0].Id != "", nil
}

func DeleteRefreshToken(token string) error {
	var gqClient = getGqClient()
	mutation := deleteSessionRefreshTokensMutation{}
	variables := map[string]interface{}{
		"token": token,
	}
	if err := gqClient.Mutate(context.Background(), &mutation, variables); err != nil {
		return err
	}
	return nil
}

type deleteSessionRefreshTokensMutation struct {
	DeleteSessionRefreshTokens struct {
		Returning []struct {
			Id string
		}
	} `graphql:"deleteSessionRefreshTokens(where:{token:{_eq:$token}})"`
}
