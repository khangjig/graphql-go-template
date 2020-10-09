package repository

import (
	"context"
	"github.com/jinzhu/gorm"
	"graphql-go-template/repository/demo"
)

type Repository struct {
	Demo demo.Repository
}

func New(getClient func(ctx context.Context) *gorm.DB) *Repository {
	return &Repository{
		Demo: demo.NewPG(getClient),
	}
}
