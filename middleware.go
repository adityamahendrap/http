package main

import (
	"net/http"
)

func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "empty auth", http.StatusUnauthorized)
			return
		}

		isMatch := (username == USERNAME) && (password == PASSWORD)
		if !isMatch {
			http.Error(w, "invalid auth", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
