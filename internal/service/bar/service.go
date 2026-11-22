package bar

import "github.com/galiullindo/example-api-handler/internal/model"

type Repository interface {
	List() ([]model.Bar, error)
	Create(bar string) (model.Bar, error)
	Get(id int64) (model.Bar, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) List() ([]model.Bar, error) {
	return s.repository.List()
}

func (s *Service) Create(bar string) (model.Bar, error) {
	return s.repository.Create(bar)
}

func (s *Service) Get(id int64) (model.Bar, error) {
	return s.repository.Get(id)
}
