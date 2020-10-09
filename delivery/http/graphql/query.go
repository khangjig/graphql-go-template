package graphql

import (
	"github.com/graphql-go/graphql"
	"graphql-go-template/delivery/http/graphql/demo"
	"graphql-go-template/repository"
)

func newQuery(repo *repository.Repository) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getDemo": demo.NewDemo(repo),
		},
	})
}
