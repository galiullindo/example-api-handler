package main

import (
	"net/http"

	"github.com/galiullindo/example-api-handler/internal/api/v1"
	"github.com/galiullindo/example-api-handler/internal/api/v1/bar"
	"github.com/galiullindo/example-api-handler/internal/api/v1/baz"
	"github.com/galiullindo/example-api-handler/internal/api/v1/foo"
	repoBar "github.com/galiullindo/example-api-handler/internal/repository/bar"
	repoBaz "github.com/galiullindo/example-api-handler/internal/repository/baz"
	repoFoo "github.com/galiullindo/example-api-handler/internal/repository/foo"
	svcBar "github.com/galiullindo/example-api-handler/internal/service/bar"
	svcBaz "github.com/galiullindo/example-api-handler/internal/service/baz"
	svcFoo "github.com/galiullindo/example-api-handler/internal/service/foo"
)

func main() {
	fooRepository := repoFoo.NewRepository()
	barRepository := repoBar.NewRepository()
	bazRepository := repoBaz.NewRepository()

	fooService := svcFoo.NewService(fooRepository)
	barService := svcBar.NewService(barRepository)
	bazService := svcBaz.NewService(bazRepository)

	handlers := api.Handlers{
		Foo: foo.NewHandler(fooService),
		Bar: bar.NewHandler(barService),
		Baz: baz.NewHandler(bazService),
	}

	apiRouter := http.NewServeMux()
	handlers.RegisterRoutes(apiRouter)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))

	server := http.Server{
		Addr:    "localhost:8081",
		Handler: mainRouter,
	}

	server.ListenAndServe()
}
