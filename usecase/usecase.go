package usecase

import (
	"graphql-go-template/repository"
	"graphql-go-template/usecase/demo"
)

type UseCase struct {
	Demo demo.IUseCase
}

func New(repo *repository.Repository) *UseCase {
	return &UseCase{
		Demo: demo.New(repo),
	}
}
