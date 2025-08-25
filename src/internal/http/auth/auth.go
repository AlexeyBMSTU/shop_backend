package auth

import (
	"encoding/json"
	user_db "github.com/AlexeyBMSTU/shop_backend/src/db/user"
	"github.com/AlexeyBMSTU/shop_backend/src/middleware/cookie"
	"github.com/AlexeyBMSTU/shop_backend/src/middleware/verify_token"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/errorGen"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/tokenGen"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/validate"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user User.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error decoding request body:", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("login attempt for user: %s\n", user.Name)

	if user.Name == "" || user.Password == "" {
		errorGen.ErrorGen(&w, "name and password are required", http.StatusBadRequest)
		return
	}

	dbUser, err := user_db.GetUserByName(user.Name)
	if err != nil {
		errorGen.ErrorGen(&w, "user not found", http.StatusBadRequest)
		return
	}
	log.Printf("login 2 attempt for user: %s\n", user.Name)
	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password)); err != nil {
		log.Println("invalid password for user:", user.Name)
		errorGen.ErrorGen(&w, "user not found", http.StatusUnauthorized)
		return
	}

	if user.Email != nil {
		if dbUser.Email == nil || *user.Email != *dbUser.Email {
			errorGen.ErrorGen(&w, "user not found", http.StatusUnauthorized)
			return
		}
	}

	token, err := tokenGen.CreateToken(dbUser.Name)
	if err != nil {
		log.Println("could not create token:", err)
		http.Error(w, "could not create token", http.StatusInternalServerError)
		return
	}

	cookie.SetCookie(w, token)
	w.WriteHeader(http.StatusOK)
	toDTO(&w, dbUser)

	log.Printf("user logged in: %s\n", dbUser.Name)
	return
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user User.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error decoding request body:", err)
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	log.Printf("registration attempt for user2: %s\n", user.Name)

	if user.Name == "" || user.Password == "" {
		errorGen.ErrorGen(&w, "username and password are required", http.StatusBadRequest)
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

	err = user_db.AddUser(user)
	if err != nil {
		log.Println("error registration:", err)
		errorGen.ErrorGen(&w, "user already exists", http.StatusBadRequest)
		return
	}

	token, err := tokenGen.CreateToken(user.Name)
	if err != nil {
		log.Println("could not create token:", err)
		errorGen.ErrorGen(&w, "could not create token", http.StatusInternalServerError)
		return
	}

	cookie.SetCookie(w, token)
	w.WriteHeader(http.StatusCreated)
	toDTO(&w, user)

	log.Printf("user  registered successfully: %s\n", user.Name)
	return
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	_, err := verify_token.VerifyToken(w, r)
	if err != nil {
		log.Println("error verifying token:", err)
		errorGen.ErrorGen(&w, "invalid token", http.StatusUnauthorized)
		return
	}

	cookieValue, _ := cookie.GetCookie(r)

	username, _ := tokenGen.ExtractUsernameFromToken(cookieValue)

	cookie.ClearCookie(w, cookieValue)

	w.WriteHeader(http.StatusOK)

	dbUser, err := user_db.GetUserByName(username)
	if err != nil {
		log.Println("error getting user by name:", err)
		errorGen.ErrorGen(&w, "user not found", http.StatusBadRequest)
		return
	}

	toDTO(&w, dbUser)
	log.Printf("user logged out: %s\n", dbUser.Name)
	return
}

func MeHandler(w http.ResponseWriter, r *http.Request) {
	_, err := verify_token.VerifyToken(w, r)
	if err != nil {
		log.Println("error verifying token:", err)
		errorGen.ErrorGen(&w, "invalid token", http.StatusUnauthorized)
		return
	}
	cookieValue, _ := r.Cookie("X-Access-Token")

	token := cookieValue.Value

	username, err := tokenGen.ExtractUsernameFromToken(token)
	if err != nil {
		log.Println("error extracting username:", err)
		errorGen.ErrorGen(&w, "invalid token", http.StatusUnauthorized)
		return
	}
	dbUser, err := user_db.GetUserByName(username)
	if err != nil {
		log.Println("error getting user:", err)
		errorGen.ErrorGen(&w, "user not found", http.StatusUnauthorized)
		return
	}

	log.Printf("protected user: %s\n", username)

	w.WriteHeader(http.StatusOK)

	toDTO(&w, dbUser)

	log.Printf("user protected successfully: %s\n", username)

	return
}
