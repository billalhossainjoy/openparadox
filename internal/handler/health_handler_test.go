package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler_Healthz(t *testing.T) {
	h := NewHealthHandler()

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	h.Healthz(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Extected 200, got %d", rec.Code)
	}
	if rec.Body.String() != "ok" {
		t.Fatalf("expected body ok , got %q", rec.Body.String())
	}
}

func TestHealthHandler_readyz(t *testing.T) {
	h := NewHealthHandler()

	req := httptest.NewRequest(http.MethodGet, "/readyz", nil)
	rec := httptest.NewRecorder()

	h.Readyz(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	if rec.Body.String() != "ready" {
		t.Fatalf("expected body ok, got %q", rec.Body.String())
	}
}
