package middleware

import (
	"net/http"
	"os"
)

func BasicAuthentication(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		uid, pass, ok := r.BasicAuth()
		if !ok || uid != os.Getenv("BASIC_AUTH_USER_ID") || pass != os.Getenv("BASIC_AUTH_PASSWORD") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
