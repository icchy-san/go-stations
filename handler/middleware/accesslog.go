package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type UserAccessLog struct {
	Timestamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}

func AccessLog(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		os := extractOS(r.Context())

		currentTime := time.Now()

		h.ServeHTTP(w, r)

		time.Sleep(1 * time.Second) // To test calculating the latency
		latency := time.Since(currentTime).Milliseconds()

		ual := UserAccessLog{
			Timestamp: currentTime,
			Latency:   latency,
			OS:        os,
			Path:      r.URL.Path,
		}

		res, err := json.Marshal(&ual)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println(string(res))
	}

	return http.HandlerFunc(fn)
}
