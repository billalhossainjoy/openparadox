package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/billalhossainjoy/openparadox/internal/domain"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func writeJson(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJson(w, status, ErrorResponse{Error: message})
}

func writeDomainError(w http.ResponseWriter, err error) {
	switch {
		case errors.Is(err, domain.ErrInvalidInput):
			writeError(w, http.StatusBadRequest, domain.ErrInvalidInput.Error())
		case errors.Is(err, domain.ErrNotFound):
			writeError(w, http.StatusNotFound, domain.ErrNotFound.Error())
		case errors.Is(err, domain.ErrAlreadyExists):
			writeError(w, http.StatusConflict, domain.ErrAlreadyExists.Error())
		default:
			writeError(w, http.StatusInternalServerError, "internal server error")
	}
}