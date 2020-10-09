package demo

import (
	"github.com/graphql-go/graphql"
)

type GraphqlDemoInterface interface {
	GetListDemo() *graphql.Field
	GetDemoByID() *graphql.Field
	CreateDemo() *graphql.Field
	UpdateDemo() *graphql.Field
}
