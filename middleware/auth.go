package middleware

import (
	"context"
	"first-jwt/helpers"
	"net/http"
)
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		accessToken := r.Header.Get("Authorization")

		// if access token is empty
		if accessToken == "" {
			helpers.Response(w, 410, "Unauthorized", nil)
			return
		}

		user, err := helpers.ValidateToken(accessToken)
		if err != nil {
			helpers.Response(w, 410, err.Error(), nil)
			return
		}

		// Context here (variabel global / session)
		ctx := context.WithValue(r.Context(), "userinfo", user) // the most 
		next.ServeHTTP(w,r.WithContext(ctx)) // proses request baru dengan context
	})
}