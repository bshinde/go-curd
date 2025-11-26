package services

import (
	"testing"

	"github.com/example/golang_crud/models"
	"github.com/example/golang_crud/pkg/repository"
)

func TestUserService(t *testing.T) {
	r := repository.NewUserRepository()
	s := NewUserService(r)

	s.Create(models.NewUser(1, "A", "a@a.com"))
	if len(s.GetAll()) != 1 {
		t.Fatalf("expected 1")
	}

	if _, ok := s.GetByID(1); !ok {
		t.Fatalf("expected get by id to work")
	}

	u := models.NewUser(1, "B", "b@b.com")
	if !s.Update(u) {
		t.Fatalf("expected update to succeed")
	}
	if !s.Delete(1) {
		t.Fatalf("expected delete to succeed")
	}
}
