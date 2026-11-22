package api

import (
	"net/http"

	"github.com/galiullindo/example-api-handler/internal/api/v2/bar"
	"github.com/galiullindo/example-api-handler/internal/api/v2/baz"
	"github.com/galiullindo/example-api-handler/internal/api/v2/foo"
)

type Services struct {
	Foo foo.Service
	Bar bar.Service
	Baz baz.Service
}

func NewRouter(s Services) http.Handler {
	mux := http.NewServeMux()

	foo.NewHandler(s.Foo).RegisterRoutes(mux)
	bar.NewHandler(s.Bar).RegisterRoutes(mux)
	baz.NewHandler(s.Baz).RegisterRoutes(mux)

	return mux
}
