package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"graphql-go-template/repository"
)

func NewHandler(repo *repository.Repository) (*handler.Handler, error) {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    newQuery(repo),
			Mutation: nil,
		},
	)
	if err != nil {
		return nil, err
	}

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	}), nil
}
