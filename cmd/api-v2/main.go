package main

import (
	"net/http"

	"github.com/galiullindo/example-api-handler/internal/api/v2"
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

	services := api.Services{
		Foo: svcFoo.NewService(fooRepository),
		Bar: svcBar.NewService(barRepository),
		Baz: svcBaz.NewService(bazRepository),
	}

	apiRouter := api.NewRouter(services)

	mainRouter := http.NewServeMux()
	mainRouter.Handle("/api/v2/", http.StripPrefix("/api/v2", apiRouter))

	server := http.Server{
		Addr:    "localhost:8082",
		Handler: mainRouter,
	}

	server.ListenAndServe()
}
