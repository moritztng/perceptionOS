//go:generate go run github.com/Khan/genqlient
package qlient

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
)

func NewClient(url string) graphql.Client {
	graphqlClient := graphql.NewClient(url, http.DefaultClient)
	return graphqlClient
}
