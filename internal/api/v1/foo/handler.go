package foo

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/galiullindo/example-api-handler/internal/api/v1/respond"
	"github.com/galiullindo/example-api-handler/internal/model"
)

type Service interface {
	List() ([]model.Foo, error)
	Create(foo string) (model.Foo, error)
	Get(id int64) (model.Foo, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
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

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Foo string `json:"foo"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respond.Write(w, []byte(`{"error":"bad request"}`), http.StatusBadRequest)
		return
	}

	new, err := h.service.Create(input.Foo)
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

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		respond.Write(w, []byte(`{"error":"bad request"}`), http.StatusBadRequest)
		return
	}

	item, err := h.service.Get(id)
	if err != nil {
		if errors.Is(err, model.ErrFooNotFound) {
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
