package auth

import (
	"encoding/json"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"net/http"
)

type AuthResponse struct {
	ID       string `json:"user_uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func toDTO(w *http.ResponseWriter, user User.User) {
	var a AuthResponse

	a.Email = user.Email
	a.Username = user.Username
	a.ID = user.ID.String()

	json.NewEncoder(*w).Encode(a)

	return
}
