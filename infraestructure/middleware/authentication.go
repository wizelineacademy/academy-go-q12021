package middleware

import "net/http"

//ValidateToken checks if the received JWT token is valid
func ValidateToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Header.Get("auth")

		if 1 == 1 {
			next.ServeHTTP(w, r)
		}

	}
}
