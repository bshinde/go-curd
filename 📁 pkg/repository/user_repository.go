package repository

import "github.com/example/golang_crud/models"

type UserRepository struct {
	users map[int]*models.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{users: make(map[int]*models.User)}
}

func (r *UserRepository) Create(u *models.User) {
	r.users[u.ID] = &models.User{ID: u.ID, Name: u.Name, Email: u.Email}
}

func (r *UserRepository) GetAll() []*models.User {
	out := make([]*models.User, 0, len(r.users))
	for _, u := range r.users {
		out = append(out, &models.User{ID: u.ID, Name: u.Name, Email: u.Email})
	}
	return out
}

func (r *UserRepository) GetByID(id int) (*models.User, bool) {
	u, ok := r.users[id]
	if !ok {
		return nil, false
	}
	return &models.User{ID: u.ID, Name: u.Name, Email: u.Email}, true
}

func (r *UserRepository) Update(u *models.User) bool {
	_, ok := r.users[u.ID]
	if !ok {
		return false
	}
	r.users[u.ID] = &models.User{ID: u.ID, Name: u.Name, Email: u.Email}
	return true
}

func (r *UserRepository) Delete(id int) bool {
	_, ok := r.users[id]
	if !ok {
		return false
	}
	delete(r.users, id)
	return true
}
