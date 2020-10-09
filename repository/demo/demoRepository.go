package demo

import (
	"context"
	"graphql-go-template/model"
)

// Repository .
type Repository interface {
	GetAll(ctx context.Context) ([]model.Demo, error)
	GetByID(ctx context.Context, id int64) (*model.Demo, error)
	Create(ctx context.Context, obj *model.Demo) error
	Update(ctx context.Context, obj *model.Demo) error
}
