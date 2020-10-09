package demo

import (
	"graphql-go-template/repository"
	"graphql-go-template/repository/demo"
)

type UseCase struct {
	DemoRepo demo.Repository
}

func New(repo *repository.Repository) IUseCase {
	return &UseCase{
		DemoRepo: repo.Demo,
	}
}
