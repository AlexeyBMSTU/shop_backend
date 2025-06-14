package auth

import (
	"encoding/json"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"net/http"
)

type AuthResponse struct {
	ID    string  `json:"user_uid"`
	Name  string  `json:"name"`
	Email *string `json:"email,omitempty"`
}

func toDTO(w *http.ResponseWriter, user User.User) {
	var a AuthResponse

	a.Email = user.Email
	a.Name = user.Name
	a.ID = user.ID.String()

	json.NewEncoder(*w).Encode(a)

	return
}
