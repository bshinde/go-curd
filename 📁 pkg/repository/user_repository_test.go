package repository

import (
	"testing"

	"github.com/example/golang_crud/models"
)

func TestUserRepository_CRUD(t *testing.T) {
	r := NewUserRepository()

	u := models.NewUser(1, "A", "a@a.com")
	r.Create(u)

	all := r.GetAll()
	if len(all) != 1 {
		t.Fatalf("expected 1 user, got %d", len(all))
	}

	got, ok := r.GetByID(1)
	if !ok || got.ID != 1 {
		t.Fatalf("GetByID failed")
	}

	u2 := models.NewUser(1, "B", "b@b.com")
	updated := r.Update(u2)
	if !updated {
		t.Fatalf("expected update to succeed")
	}
	got, _ = r.GetByID(1)
	if got.Name != "B" {
		t.Fatalf("update didn't apply")
	}

	deleted := r.Delete(1)
	if !deleted {
		t.Fatalf("expected delete to succeed")
	}
	if len(r.GetAll()) != 0 {
		t.Fatalf("expected 0 users after delete")
	}
}
