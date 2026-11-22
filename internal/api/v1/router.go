package api

import (
	"net/http"
)

type FooHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type BarHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type BazHandler interface {
	List(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type Handlers struct {
	Foo FooHandler
	Bar BarHandler
	Baz BazHandler
}

func (h *Handlers) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /foo", h.Foo.List)
	mux.HandleFunc("POST /foo", h.Foo.Create)
	mux.HandleFunc("GET /foo/{id}", h.Foo.Get)

	mux.HandleFunc("GET /bar", h.Bar.List)
	mux.HandleFunc("POST /bar", h.Bar.Create)
	mux.HandleFunc("GET /bar/{id}", h.Bar.Get)

	mux.HandleFunc("GET /baz", h.Baz.List)
	mux.HandleFunc("POST /baz", h.Baz.Create)
	mux.HandleFunc("GET /baz/{id}", h.Baz.Get)
}
