package baz

import "github.com/galiullindo/example-api-handler/internal/model"

type Repository interface {
	List() ([]model.Baz, error)
	Create(baz string) (model.Baz, error)
	Get(id int64) (model.Baz, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) List() ([]model.Baz, error) {
	return s.repository.List()
}

func (s *Service) Create(baz string) (model.Baz, error) {
	return s.repository.Create(baz)
}

func (s *Service) Get(id int64) (model.Baz, error) {
	return s.repository.Get(id)
}
