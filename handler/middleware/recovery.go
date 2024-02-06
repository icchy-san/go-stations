package middleware

import (
	"encoding/json"
	"net/http"
)

func Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// TODO: ここに実装をする
		defer func() {
			// when err is not nil, it means panic occurred
			err := recover()
			if err != nil {
				jsonBody, _ := json.Marshal(map[string]string{
					"error": "there exist an internal server error",
				})

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(jsonBody)
			}
		}()

		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
