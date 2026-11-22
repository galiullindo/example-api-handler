package baz

import (
	"sync"

	"github.com/galiullindo/example-api-handler/internal/model"
)

type Repository struct {
	mu      sync.Mutex
	storage map[int64]model.Baz
	nextID  int64
}

func NewRepository() *Repository {
	return &Repository{
		storage: make(map[int64]model.Baz),
		nextID:  1,
	}
}

func (r *Repository) List() ([]model.Baz, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	list := make([]model.Baz, 0, len(r.storage))
	for _, item := range r.storage {
		list = append(list, item)
	}

	return list, nil
}

func (r *Repository) Create(baz string) (model.Baz, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.nextID
	r.nextID++

	new := model.Baz{ID: id, Baz: baz}
	r.storage[id] = new

	return new, nil
}

func (r *Repository) Get(id int64) (model.Baz, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, exists := r.storage[id]
	if !exists {
		return model.Baz{}, model.ErrBazNotFound
	}

	return item, nil
}
