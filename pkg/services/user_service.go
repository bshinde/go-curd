package services

import "github.com/example/golang_crud/models"
import "github.com/example/golang_crud/pkg/repository"

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(r *repository.UserRepository) *UserService {
	return &UserService{repo: r}
}

func (s *UserService) Create(u *models.User) {
	s.repo.Create(u)
}

func (s *UserService) GetAll() []*models.User {
	return s.repo.GetAll()
}

func (s *UserService) GetByID(id int) (*models.User, bool) {
	return s.repo.GetByID(id)
}

func (s *UserService) Update(u *models.User) bool {
	return s.repo.Update(u)
}

func (s *UserService) Delete(id int) bool {
	return s.repo.Delete(id)
}
