package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/example/golang_crud/pkg/repository"
	"githubgithub.com/example/golang_crud/pkg/services"
	"github.com/example/golang_crud/router"
)

func main() {
	repo := repository.NewUserRepository()
	svc := services.NewUserService(repo)

	mux := http.NewServeMux()
	router.RegisterRoutes(mux, svc)

	fmt.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
