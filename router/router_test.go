package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/example/golang_crud/models"
	"github.com/example/golang_crud/pkg/repository"
	"github.com/example/golang_crud/pkg/services"
)

func TestRouter_CRUD(t *testing.T) {
	mux := http.NewServeMux()
	svc := services.NewUserService(repository.NewUserRepository())
	RegisterRoutes(mux, svc)

	// POST create
	user := models.NewUser(1, "John", "john@example.com")
	b, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	// GET list
	req = httptest.NewRequest(http.MethodGet, "/users", nil)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}
	var list []models.User
	if err := json.NewDecoder(w.Body).Decode(&list); err != nil {
		t.Fatalf("decode list failed: %v", err)
	}
	if len(list) != 1 {
		t.Fatalf("expected 1 user, got %d", len(list))
	}

	// GET by id
	req = httptest.NewRequest(http.MethodGet, "/users/1", nil)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	// PUT update
	updated := models.User{Name: "Jane", Email: "jane@example.com"}
	b, _ = json.Marshal(updated)
	req = httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 on update, got %d", w.Code)
	}

	// DELETE
	req = httptest.NewRequest(http.MethodDelete, "/users/1", nil)
	w = httptest.NewRecorder()
	mux.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("expected 204 on delete, got %d", w.Code)
	}
}
