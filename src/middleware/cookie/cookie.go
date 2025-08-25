package cookie

import (
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, token string) {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "X-Access-Token",
		Value:    token,
		Path:     "/",
		Expires:  expiration,
		HttpOnly: true,
		Secure:   false, // true - HTTPS
	}
	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request) (string, error) {
	cookie, err := r.Cookie("X-Access-Token")
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

func ClearCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:     token,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   false, // true - HTTPS
	}
	http.SetCookie(w, &cookie)
}
