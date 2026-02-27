package middleware

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
)

type ctxKey string

const RequestIdKey ctxKey = "request_id"

func GetRequestId(ctx context.Context)string {
	v, _ := ctx.Value(RequestIdKey).(string)
	return v
}

func RequestId(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		reqId:= r.Header.Get("X-Request-Id")
		if reqId == "" {
			reqId = newRequestId()
		}

		w.Header().Set("X-Request-Id", reqId)

		ctx := context.WithValue(r.Context(), RequestIdKey, reqId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func newRequestId() string{
	b:= make([]byte, 16)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
