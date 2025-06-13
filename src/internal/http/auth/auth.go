package auth

import (
	"encoding/json"
	"github.com/AlexeyBMSTU/shop_backend/src/db"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/errorGen"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/tokenGen"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/validate"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, token string) {
	expiration := time.Now().Add(24 * time.Hour)
	cookie := http.Cookie{
		Name:     "X-Access-Token",
		Value:    token,
		Path:     "/",
		Expires:  expiration,
		HttpOnly: true,
		Secure:   false, // 'http': false, 'https': true
	}
	http.SetCookie(w, &cookie)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error decoding request body:", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("login attempt for user: %s\n", user.Username)

	if user.Username == "" || user.Email == "" || user.Password == "" {
		errorGen.ErrorGen(&w, "username, email, and password are required", http.StatusBadRequest)
		return
	}

	dbUser, err := db.GetUserByName(user.Username)
	if err != nil {
		errorGen.ErrorGen(&w, "user not found", http.StatusBadRequest)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		log.Println("invalid password for user:", user.Username)
		errorGen.ErrorGen(&w, "invalid password", http.StatusUnauthorized)
		return
	}

	if dbUser.Email != user.Email {
		errorGen.ErrorGen(&w, "invalid email", http.StatusUnauthorized)
		return
	}

	token, err := tokenGen.CreateToken(dbUser.Username)
	if err != nil {
		log.Println("could not create token:", err)
		http.Error(w, "could not create token", http.StatusInternalServerError)
		return
	}

	setCookie(w, token)
	w.WriteHeader(http.StatusOK)
	toDTO(&w, dbUser)

	log.Printf("user logged in: %s\n", dbUser.Username)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error decoding request body:", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("registration attempt for user: %s\n", user.Username)

	if user.Username == "" || user.Email == "" || user.Password == "" {
		errorGen.ErrorGen(&w, "username, email, and password are required", http.StatusBadRequest)
		return
	}

	err := validate.ValidatingUser(user)
	if err != nil {
		log.Println("error reg:", err)
		errorGen.ErrorGen(&w, "invalid credentials", http.StatusBadRequest)
		return
	}
	user.ID = uuid.New()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("error hashing password:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	user.Password = string(hashedPassword)

	err = db.AddUser(user)
	if err != nil {
		log.Println("error registration:", err)
		errorGen.ErrorGen(&w, "user already exists", http.StatusBadRequest)
		return
	}

	token, err := tokenGen.CreateToken(user.Username)
	if err != nil {
		log.Println("could not create token:", err)
		errorGen.ErrorGen(&w, "could not create token", http.StatusInternalServerError)
		return
	}

	setCookie(w, token)
	w.WriteHeader(http.StatusCreated)
	toDTO(&w, user)
	log.Printf("user  registered successfully: %s\n", user.Username)
}
