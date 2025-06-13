package auth

import (
	"encoding/json"
	"github.com/AlexeyBMSTU/shop_backend/src/db"
	"github.com/AlexeyBMSTU/shop_backend/src/models/Error"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/errorGen"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/tokenGen"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Login attempt for user: %s\n", user.Username)
	user, err := db.GetUserByName(user.Username)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
	}
	token, err := tokenGen.CreateToken(user.Username)
	if err != nil {
		log.Println("Could not create token:", err)
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token))
	w.WriteHeader(http.StatusCreated)
	log.Printf("User logged in: %s\n", user.Username)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User.User
	var error Error.Error
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("Registration attempt for user: %s\n", user.Username)

	if user.Username == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "Username, email, and password are required", http.StatusBadRequest)
		return
	}

	err := db.AddUser(user.Username, user.Email)
	if err != nil {
		log.Println("Error adding user to database:", err)
		errorGen.ErrorGen(w, "could registration", http.StatusBadRequest)
	}

	token, err := tokenGen.CreateToken(user.Username)
	if err != nil {
		log.Println("Could not create token:", err)
		http.Error(w, "Could not create token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(token))
	log.Printf("User  registered successfully: %s\n", user.Username)
}
