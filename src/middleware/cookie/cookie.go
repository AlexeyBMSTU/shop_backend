package cookie

import (
	"net/http"
	"time"
)

var authCookieName = "X-Access-Token"

func SetCookie(w http.ResponseWriter, token string) {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     authCookieName,
		Value:    token,
		Path:     "/",
		Expires:  expiration,
		HttpOnly: true,
		Secure:   false, // true - HTTPS
		SameSite: http.SameSiteStrictMode,
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

func ClearCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     authCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   false, // true - HTTPS
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, &cookie)
}
