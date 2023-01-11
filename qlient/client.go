//go:generate go run github.com/Khan/genqlient
package qlient

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

func NewClient() graphql.Client {
	graphqlClient := graphql.NewClient("http://localhost:8080/query", http.DefaultClient)
	return graphqlClient
}
