package demo

import (
	"context"
	"github.com/graphql-go/graphql"
	"graphql-go-template/model"
	"graphql-go-template/repository"
	rp "graphql-go-template/repository/demo"
)

var demo = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Demo",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
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

func NewGraphqlDemo(repo *repository.Repository) GraphqlDemoInterface {
	return GraphqlDemo{
		DemoRepo: repo.Demo,
	}
}

type GraphqlDemo struct {
	DemoRepo rp.Repository
}

func (g GraphqlDemo) GetListDemo() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(demo),
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			return g.DemoRepo.GetAll(context.Background())
		},
		Description: "demo",
	}
}

func (g GraphqlDemo) GetDemoByID() *graphql.Field {
	return &graphql.Field{
		Type: demo,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.Int,
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			id, ok := p.Args["id"].(int)
			if ok {
				return g.DemoRepo.GetByID(context.Background(), int64(id))
			}

			return nil, nil
		},
		Description: "demo",
	}
}

func (g GraphqlDemo) CreateDemo() *graphql.Field {
	return &graphql.Field{
		Type: demo,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			createDemo := model.Demo{
				Name: p.Args["name"].(string),
			}

			err := g.DemoRepo.Create(context.Background(), &createDemo)
			if err != nil {
				return nil, err
			}

			return createDemo, nil
		},
		Description: "demo",
	}
}

func (g GraphqlDemo) UpdateDemo() *graphql.Field {
	return &graphql.Field{
		Type: demo,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.Int),
			},
			"name": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			updateDemo := model.Demo{
				ID:   int64(p.Args["id"].(int)),
				Name: p.Args["name"].(string),
			}

			err := g.DemoRepo.Update(context.Background(), &updateDemo)
			if err != nil {
				return nil, err
			}

			return updateDemo, nil
		},
		Description: "demo",
	}
}
