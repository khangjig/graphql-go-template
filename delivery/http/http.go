package http

import (
	"fmt"
	"github.com/graphql-go/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"graphql-go-template/delivery/http/demo"
	"graphql-go-template/usecase"
	"log"
	"net/http"
)

func NewHTTPHandler(useCase *usecase.UseCase, handleGql *handler.Handler) *echo.Echo {
	e := echo.New()

	logCfg := middleware.DefaultLoggerConfig

	logCfg.Skipper = func(c echo.Context) bool {
		return c.Request().URL.RequestURI() == "/health_check"
	}

	e.Use(middleware.LoggerWithConfig(logCfg))
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		e.DefaultHTTPErrorHandler(err, c)
	}
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch,
			http.MethodPost, http.MethodDelete, http.MethodOptions,
		},
	}))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		if c.Request().URL.RequestURI() != "/health_check" {
			request := fmt.Sprintf("%s", reqBody)

			if len(request) > 0 {
				log.Printf("%s", request)
			}
			log.Printf("%s", resBody)
		}
	}))

	e.GET("/health_check", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"message": "ok"})
	})

	e.POST("/graphql", echo.WrapHandler(handleGql))

	api := e.Group("/v1")
	demo.Init(api, useCase)

	return e
}
