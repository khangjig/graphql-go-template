package demo

import (
	"context"
	"graphql-go-template/model"
)

type IUseCase interface {
	GetDemo(ctx context.Context, ID int64) (*model.Demo, error)
}
