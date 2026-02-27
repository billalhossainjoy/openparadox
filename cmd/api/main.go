package main

import (
	"net/http"

	"github.com/billalhossainjoy/openparadox/internal/handler"
	"github.com/billalhossainjoy/openparadox/internal/middleware"
	"github.com/billalhossainjoy/openparadox/internal/repository"
	"github.com/billalhossainjoy/openparadox/internal/service"
)


func main() {
	repo:= repository.NewMemoryUserRepository()
	service:= service.NewUserService(repo)
	handler:= handler.NewUserHandler(service)

	mux:= http.NewServeMux()
	loggedMux:= middleware.Logging(mux)
	mux.HandleFunc("POST /users", handler.CreateUser)
	mux.HandleFunc("GET /users", handler.GetUsers)
	mux.HandleFunc("GET /user/{id}", handler.GetUser)

	http.ListenAndServe(":8080", loggedMux)
}