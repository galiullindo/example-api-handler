package foo

import "github.com/galiullindo/example-api-handler/internal/model"

type Repository interface {
	List() ([]model.Foo, error)
	Create(foo string) (model.Foo, error)
	Get(id int64) (model.Foo, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) List() ([]model.Foo, error) {
	return s.repository.List()
}

func (s *Service) Create(foo string) (model.Foo, error) {
	return s.repository.Create(foo)
}

func (s *Service) Get(id int64) (model.Foo, error) {
	return s.repository.Get(id)
}
