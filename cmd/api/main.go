package main

import (
	"net/http"

	"github.com/billalhossainjoy/openparadox/internal/config"
	"github.com/billalhossainjoy/openparadox/internal/handler"
	"github.com/billalhossainjoy/openparadox/internal/middleware"
	"github.com/billalhossainjoy/openparadox/internal/repository"
	"github.com/billalhossainjoy/openparadox/internal/service"
)


func main() {
	cfg :=config.Load()

	repo:= repository.NewMemoryUserRepository()
	service:= service.NewUserService(repo)
	userHandler:= handler.NewUserHandler(service)
	healthHandler:= handler.NewHealthHandler()

	mux:= http.NewServeMux()




	app:= middleware.Chain(
		mux,
		middleware.Recover,
		middleware.RequestId,
		middleware.Logging,
		middleware.RequestTimeout(cfg.ReqTimeout),
	)


	// Health Routes
	mux.HandleFunc("GET /healthz", healthHandler.Healthz)
	mux.HandleFunc("GET /readyz", healthHandler.Readyz)

	// User Routes
	mux.HandleFunc("POST /users", userHandler.CreateUser)
	mux.HandleFunc("GET /users", userHandler.GetUsers)
	mux.HandleFunc("GET /user/{id}", userHandler.GetUser)

	http.ListenAndServe(":8080", app)
}