package baz

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/galiullindo/example-api-handler/internal/api/v2/respond"
	"github.com/galiullindo/example-api-handler/internal/model"
)

type Service interface {
	List() ([]model.Baz, error)
	Create(baz string) (model.Baz, error)
	Get(id int64) (model.Baz, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /baz", h.list)
	mux.HandleFunc("POST /baz", h.create)
	mux.HandleFunc("GET /baz/{id}", h.get)
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	list, err := h.service.List()
	if err != nil {
		respond.Write(w, []byte(`{"error":"internal server error"}`), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(list)
	if err != nil {
		respond.Write(w, []byte(`{"error":"internal server error"}`), http.StatusInternalServerError)
		return
	}

	respond.Write(w, body, http.StatusOK)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Baz string `json:"baz"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respond.Write(w, []byte(`{"error":"bad request"}`), http.StatusBadRequest)
		return
	}

	new, err := h.service.Create(input.Baz)
	if err != nil {
		respond.Write(w, []byte(`{"error":"internal server error"}`), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(new)
	if err != nil {
		respond.Write(w, []byte(`{"error":"internal server error"}`), http.StatusInternalServerError)
		return
	}

	respond.Write(w, body, http.StatusOK)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respond.Write(w, []byte(`{"error":"bad request"}`), http.StatusBadRequest)
		return
	}

	item, err := h.service.Get(id)
	if err != nil {
		if errors.Is(err, model.ErrBazNotFound) {
			respond.Write(w, []byte(`{"error":"not found"}`), http.StatusNotFound)
			return
		}
		respond.Write(w, []byte(`{"error":"internal server error"}`), http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(item)
	if err != nil {
		respond.Write(w, []byte(`{"error":"internal server error"}`), http.StatusInternalServerError)
		return
	}

	respond.Write(w, body, http.StatusOK)
}
