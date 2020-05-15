package user

import "github.com/Abacode7/delivery-service/pkg/domain"

type Service interface {
	GetUser(id int) (*domain.User, error)
	GetAllUser()map[string]interface{}
	CreateUser(user domain.User) map[string]interface{}
}

type service struct {
	repository domain.UserRepository
}

// NewService is the constructor for service
func NewService(repo domain.UserRepository) service {
	return service{repo}
}

func (s *service) GetUser(id int) (*domain.User, error) {
	return nil, nil
}

func (s *service) GetAllUser() map[string]interface{}{
	return nil
}

func (s *service) CreateUser(user domain.User) map[string]interface{}{
	return nil
}