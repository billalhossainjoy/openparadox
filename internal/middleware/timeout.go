package middleware

import (
	"context"
	"net/http"
	"time"
)

func RequestTimeout(timeout time.Duration) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), timeout)
			defer cancel()

			r = r.WithContext(ctx)

			done:= make(chan struct{})

			go func ()  {
				next.ServeHTTP(w,r)
				close(done)
			}()

			select{
			case <-done:
				return
			case <-ctx.Done():
				http.Error(w, "request timeout", http.StatusGatewayTimeout)
			}

		})
	}
}