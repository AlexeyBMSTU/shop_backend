package auth

import (
	"encoding/json"
	"github.com/AlexeyBMSTU/shop_backend/src/db/init"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/tokenGen"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User.User
	init.initDB()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Login attempt for user: %s\n", user.Username)

	if user.Username == "validUser " && user.Password == "validPassword" {
		token, err := tokenGen.CreateToken(user.Username)
		if err != nil {
			log.Println("Could not create token:", err)
			http.Error(w, "Could not create token", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(token))
		log.Printf("Token created for user: %s\n", user.Username)
	} else {
		log.Println("Invalid credentials for user:", user.Username)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}
}
