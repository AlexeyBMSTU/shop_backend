package verify_token

import (
	"errors"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/errorGen"
	"net/http"
)

var secretKey = []byte("secret_key")

func VerifyToken(w http.ResponseWriter, r *http.Request) (*http.Cookie, error) {
	token, err := r.Cookie("X-Access-Token")
	if err != nil {
		if err == http.ErrNoCookie {
			errorGen.ErrorGen(&w, "unauthorized", http.StatusUnauthorized)
			return nil, errors.New("unauthorized")
		}
		http.Error(w, "bad request", http.StatusBadRequest)
		return nil, errors.New("bad request")
	}
	return token, nil
}
