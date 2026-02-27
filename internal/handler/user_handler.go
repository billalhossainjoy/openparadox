package handler

import (
	"encoding/json"
	"net/http"

	"github.com/billalhossainjoy/openparadox/internal/service"
)


type UserHandler struct{
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
		Email string `json:"email"`
	}

	if err:= json.NewDecoder(r.Body).Decode(&req); err!=nil {
		http.Error(w,err.Error(),http.StatusBadRequest)
		return
	}

	user, err := h.service.CreateUser(r.Context(), req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetUsers(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	id:= r.PathValue("id")

	user, err := h.service.GetUser(r.Context(), id);
	if err!= nil {
		writeError(w, http.StatusNotFound, "user not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
