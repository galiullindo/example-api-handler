package bar

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/galiullindo/example-api-handler/internal/api/v2/respond"
	"github.com/galiullindo/example-api-handler/internal/model"
)

type Service interface {
	List() ([]model.Bar, error)
	Create(bar string) (model.Bar, error)
	Get(id int64) (model.Bar, error)
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /bar", h.list)
	mux.HandleFunc("POST /bar", h.create)
	mux.HandleFunc("GET /bar/{id}", h.get)
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
		Bar string `json:"bar"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		respond.Write(w, []byte(`{"error":"bad request"}`), http.StatusBadRequest)
		return
	}

	new, err := h.service.Create(input.Bar)
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
		if errors.Is(err, model.ErrBarNotFound) {
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
