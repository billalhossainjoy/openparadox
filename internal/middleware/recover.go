package middleware

import (
	"log"
	"net/http"
)

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		defer func () {
			if rec:= recover(); rec != nil {
				reqId := GetRequestId(r.Context())
				log.Printf("rid=%s panic=%v", reqId, rec)
			}
		}()

		next.ServeHTTP(w, r)
	})
}