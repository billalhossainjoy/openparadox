package middleware

import "net/http"

// statusRecorder records the status code written by handlers.
// Important: ServeMux writes 404 by calling WriteHeader(404).
type statusRecorder struct {
	http.ResponseWriter
	status int
	wrote  bool
}

func (w *statusRecorder) WriteHeader(code int) {
	if w.wrote {
		return
	}
	w.status = code
	w.wrote = true
	w.ResponseWriter.WriteHeader(code)
}

func (w *statusRecorder) Write(b []byte) (int, error) {
	if !w.wrote {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}

func JSONHTTPErrorDefaults(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := &statusRecorder{ResponseWriter: w, status: http.StatusOK}

		next.ServeHTTP(sw, r)

		// If handler already wrote a non-404/405 response, do nothing.
		if sw.status != http.StatusNotFound && sw.status != http.StatusMethodNotAllowed {
			return
		}

		// If something already wrote a body, we should not try to overwrite it.
		// In practice, ServeMux 404/405 typically has no JSON body.
		if sw.wrote {
			// We can't safely change it now. So just return.
			// (If you want full control, use a custom router later.)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if sw.status == http.StatusNotFound {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(`{"error":"not found"}`))
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(`{"error":"method not allowed"}`))
	})
}
