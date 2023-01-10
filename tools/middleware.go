package tools

import "net/http"

func (t *Tools) Middleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header != t.cfg.Authorization {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{ "errors": "401", "message": "invalid authorization" }`))
			return
		}
		next(w, r)
	}
}
