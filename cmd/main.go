package main

import (
	"context"
	"fmt"
	"graphql-go-template/client/mysql"
	"graphql-go-template/config"
	"graphql-go-template/datastores"
	serviceHttp "graphql-go-template/delivery/http"
	"graphql-go-template/delivery/http/graphql"
	"graphql-go-template/repository"
	"graphql-go-template/usecase"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

func main() {
	cfg := config.GetConfig()
	defer mysql.Disconnect()

	go func() {
		err := http.ListenAndServe("localhost:1997", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	// setup locale
	{
		loc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		time.Local = loc
	}

	client := mysql.GetClient
	repo := repository.New(client)

	datastores.Migrate(client(context.TODO()))

	useCase := usecase.New(repo)

	// graphql
	gql, err := graphql.NewHandler(repo)
	if err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(l)
	grpcL := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings("content-type", "application/grpc"))
	httpL := m.Match(cmux.HTTP1Fast())

	errs := make(chan error)

	// http
	{
		h := serviceHttp.NewHTTPHandler(useCase, gql)
		go func() {
			h.Listener = httpL
			errs <- h.Start("")
		}()
	}

	// gRPC
	{
		s := grpc.NewServer()

		go func() {
			// ll, _ := net.Listen("tcp", ":3001")
			errs <- s.Serve(grpcL)
		}()
	}

	go func() {
		errs <- m.Serve()
	}()

	// graceful
	// nolint
	go func() {
		c := make(chan os.Signal)
		signal.Notify(
			c,
			syscall.SIGINT,
			syscall.SIGTERM,
		)

		errs <- fmt.Errorf("%s", <-c)
	}()

	err = <-errs
	if err != nil {
		log.Fatal(err)
	}

	log.Println("exit")
}
