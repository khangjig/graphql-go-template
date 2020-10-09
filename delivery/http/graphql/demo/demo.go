package demo

import (
	"context"
	"graphql-go-template/repository"

	"github.com/graphql-go/graphql"
)

var demo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Demo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.ID,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"created_at": &graphql.Field{
				Type: graphql.String,
			},
			"updated_at": &graphql.Field{
				Type: graphql.String,
			},
		},
		Description: "get demo",
	},
)

func NewDemo(repo *repository.Repository) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(demo),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return repo.Demo.GetAll(context.Background())
		},
		Description: "demo",
	}
}
