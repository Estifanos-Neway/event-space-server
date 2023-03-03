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
