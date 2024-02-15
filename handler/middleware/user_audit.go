package middleware

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/mileusna/useragent"
)

type contextKey string

const (
	ContextKeyOS contextKey = "OS"
)

func UserAudit(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userAgent := useragent.Parse(r.UserAgent())
		ctx := context.WithValue(r.Context(), ContextKeyOS, userAgent.OS)
		startedAt := time.Now()
		h.ServeHTTP(w, r.WithContext(ctx))
		log.Println(time.Since(startedAt))
	}

	return http.HandlerFunc(fn)
}

func extractOS(ctx context.Context) string {
	os, _ := ctx.Value(ContextKeyOS).(string)
	return os
}
