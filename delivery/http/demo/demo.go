package demo

import (
	"github.com/labstack/echo/v4"
	"graphql-go-template/util"
	"graphql-go-template/util/myerror"
)

func (r *Route) getDemo(c echo.Context) error {
	ctx := &util.CustomEchoContext{Context: c}

	response, err := r.useCase.Demo.GetDemo(ctx, 1)
	if err != nil {
		return util.Response.Error(ctx, err.(myerror.MyError))
	}

	return util.Response.Success(c, response)
}
