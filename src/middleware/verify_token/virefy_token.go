package verify_token

import (
	"context"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/tokenGen"
	"net/http"
)

var secretKey = []byte("secret_key")

func VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("X-Access-Token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, "bad request", http.StatusBadRequest)
			return
		}

		token := cookie.Value

		username, err := tokenGen.CreateToken(token)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
