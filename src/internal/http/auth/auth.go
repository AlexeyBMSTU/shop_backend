package auth

import (
	"encoding/json"
	"net/http"
	"shop_backend/src/models/User"
	"shop_backend/src/utils/tokenGen"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request2", http.StatusBadRequest)
		return
	}

	if user.Username == "validUser " && user.Password == "validPassword" {
		token, err := tokenGen.CreateToken(user.Username)
		if err != nil {
			http.Error(w, "Could not create token", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
	} else {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
