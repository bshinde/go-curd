package router

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/example/golang_crud/models"
	"github.com/example/golang_crud/pkg/services"
)

func RegisterRoutes(mux *http.ServeMux, svc *services.UserService) {
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(svc.GetAll())
		case http.MethodPost:
			var u models.User
			if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			svc.Create(&u)
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(u)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		idStr := strings.TrimPrefix(r.URL.Path, "/users/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}
		switch r.Method {
		case http.MethodGet:
			u, ok := svc.GetByID(id)
			if !ok {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			_ = json.NewEncoder(w).Encode(u)
		case http.MethodPut:
			var u models.User
			if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
				http.Error(w, "bad request", http.StatusBadRequest)
				return
			}
			u.ID = id
			if !svc.Update(&u) {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			_ = json.NewEncoder(w).Encode(u)
		case http.MethodDelete:
			if !svc.Delete(id) {
				http.Error(w, "not found", http.StatusNotFound)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
