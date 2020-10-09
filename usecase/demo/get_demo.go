package demo

import (
	"context"
	"graphql-go-template/model"
	"graphql-go-template/util/myerror"
)

func (u *UseCase) GetDemo(ctx context.Context, ID int64) (*model.Demo, error) {
	data, err := u.DemoRepo.GetByID(ctx, ID)
	if err != nil {
		return nil, myerror.ErrDemoGetDemo(err)
	}

	return data, nil
}
