package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(router http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			router.ServeHTTP(w, r)
			log.Println(r.Method, r.URL.Path, time.Since(start))
		})
}
