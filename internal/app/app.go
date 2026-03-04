package app

import (
	"net/http"
	"time"

	"github.com/billalhossainjoy/openparadox/internal/handler"
	"github.com/billalhossainjoy/openparadox/internal/middleware"
	"github.com/billalhossainjoy/openparadox/internal/repository"
	"github.com/billalhossainjoy/openparadox/internal/service"
)

type Deps struct {
	ReqTimeout time.Duration
}

func New(deps Deps) http.Handler {
	repo := repository.NewMemoryUserRepository()
	userService := service.NewUserService(repo)

	userH := handler.NewUserHandler(userService)
	healthH := handler.NewHealthHandler()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", healthH.Healthz)
	mux.HandleFunc("GET /readyz", healthH.Readyz)

	mux.HandleFunc("POST /users", userH.CreateUser)
	mux.HandleFunc("GET /users", userH.GetUsers)
	mux.HandleFunc("GET /user/{id}", userH.GetUser)

	return middleware.Chain(
		mux,
		middleware.Recover,
		middleware.RequestId,
		middleware.Logging,
		middleware.BodyLimit(1<<20),
		middleware.RequestTimeout(deps.ReqTimeout),
	)

}
