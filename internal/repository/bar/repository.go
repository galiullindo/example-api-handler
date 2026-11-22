package bar

import (
	"sync"

	"github.com/galiullindo/example-api-handler/internal/model"
)

type Repository struct {
	mu      sync.Mutex
	storage map[int64]model.Bar
	nextID  int64
}

func NewRepository() *Repository {
	return &Repository{
		storage: make(map[int64]model.Bar),
		nextID:  1,
	}
}

func (r *Repository) List() ([]model.Bar, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	list := make([]model.Bar, 0, len(r.storage))
	for _, item := range r.storage {
		list = append(list, item)
	}

	return list, nil
}

func (r *Repository) Create(bar string) (model.Bar, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.nextID
	r.nextID++

	new := model.Bar{ID: id, Bar: bar}
	r.storage[id] = new

	return new, nil
}

func (r *Repository) Get(id int64) (model.Bar, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, exists := r.storage[id]
	if !exists {
		return model.Bar{}, model.ErrBarNotFound
	}

	return item, nil
}
