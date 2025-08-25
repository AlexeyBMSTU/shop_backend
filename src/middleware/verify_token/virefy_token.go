package verify_token

import (
	"errors"
	"net/http"
)

var secretKey = []byte("secret_key")

func VerifyToken(w http.ResponseWriter, r *http.Request) (*http.Cookie, error) {
	token, err := r.Cookie("X-Access-Token")
	if err != nil {
		if err == http.ErrNoCookie {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return nil, errors.New("unauthorized")
		}
		http.Error(w, "bad request", http.StatusBadRequest)
		return nil, errors.New("bad request")
	}
	return token, nil
	//token := cookie.Value
	//
	//username, err := tokenGen.ExtractUsernameFromToken(token)
	//if err != nil {
	//	http.Error(w, "unauthorized", http.StatusUnauthorized)
	//	return "", errors.New("unauthorized")
	//}
	//
	//return username, nil
}

//
//func VerifyToken(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		cookie, err := r.Cookie("X-Access-Token")
//		if err != nil {
//			if err == http.ErrNoCookie {
//				http.Error(w, "unauthorized", http.StatusUnauthorized)
//				return
//			}
//			http.Error(w, "bad request", http.StatusBadRequest)
//			return
//		}
//
//		token := cookie.Value
//
//		username, err := tokenGen.ExtractUsernameFromToken(token)
//		if err != nil {
//			http.Error(w, "unauthorized", http.StatusUnauthorized)
//			return
//		}
//
//		return username
//
//		//toDTO(&w, "authorization", username)
//	})
//}
