package service

import (
	"net/http"

	jwtauth "github.com/kuromii5/posts/internal/lib/jwtAuth"
)

func (s *Service) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			next.ServeHTTP(w, r)
			return
		}

		bearer := "Bearer "
		auth = auth[len(bearer):]

		validate, err := jwtauth.VerifyJWTToken(auth, s.secret)
		if err != nil || !validate.Valid {
			http.Error(w, "Invalid token", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
