package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logging(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r * http.Request) {
		start := time.Now()

		next.ServeHTTP(w,r)

		reqId:= GetRequestId(r.Context())

		log.Printf("rid=%s method=%s path=%s dur=%s",
		reqId,
		r.Method,
		r.URL.Path,
		time.Since(start),
	)
	})
}