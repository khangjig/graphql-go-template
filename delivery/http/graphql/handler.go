package graphql

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"github.com/pkg/errors"
	"graphql-go-template/delivery/http/graphql/demo"
	"graphql-go-template/repository"
)

func Handler(repo *repository.Repository) (*handler.Handler, error) {
	demoHandle := demo.NewGraphqlDemo(repo)

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query:    newQuery(demoHandle),
			Mutation: newMutation(demoHandle),
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "schema graphql")
	}

	return handler.New(
		&handler.Config{
			Schema:   &schema,
			Pretty:   true,
			GraphiQL: true,
		},
	), nil
}

func newQuery(demoHandle demo.GraphqlDemoInterface) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"getListDemo": demoHandle.GetListDemo(),
			"getDemoByID": demoHandle.GetDemoByID(),
		},
	})
}

func newMutation(demoHandle demo.GraphqlDemoInterface) *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createDemo": demoHandle.CreateDemo(),
			"updateDemo": demoHandle.UpdateDemo(),
		},
	})
}
