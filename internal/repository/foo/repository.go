package foo

import (
	"sync"

	"github.com/galiullindo/example-api-handler/internal/model"
)

type Repository struct {
	mu      sync.Mutex
	storage map[int64]model.Foo
	nextID  int64
}

func NewRepository() *Repository {
	return &Repository{
		storage: make(map[int64]model.Foo),
		nextID:  1,
	}
}

func (r *Repository) List() ([]model.Foo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	list := make([]model.Foo, 0, len(r.storage))
	for _, item := range r.storage {
		list = append(list, item)
	}

	return list, nil
}

func (r *Repository) Create(foo string) (model.Foo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id := r.nextID
	r.nextID++

	new := model.Foo{ID: id, Foo: foo}
	r.storage[id] = new

	return new, nil
}

func (r *Repository) Get(id int64) (model.Foo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	item, exists := r.storage[id]
	if !exists {
		return model.Foo{}, model.ErrFooNotFound
	}

	return item, nil
}
