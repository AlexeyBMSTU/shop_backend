package profile

import (
	"encoding/json"
	user_db "github.com/AlexeyBMSTU/shop_backend/src/db/user"
	"github.com/AlexeyBMSTU/shop_backend/src/middleware/cookie"
	"github.com/AlexeyBMSTU/shop_backend/src/middleware/verify_token"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/errorGen"
	"github.com/AlexeyBMSTU/shop_backend/src/utils/tokenGen"
	"log"
	"net/http"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user User.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error decoding request body:", err)
		errorGen.ErrorGen(&w, "bad request", http.StatusBadRequest)
		return
	}

	_, err := verify_token.VerifyToken(w, r)
	if err != nil {
		log.Println("error verifying token:", err)
		return
	}
	cookieValue, _ := r.Cookie("X-Access-Token")

	token := cookieValue.Value
	id, err := tokenGen.ExtractIDFromToken(token)
	if err != nil {
		log.Println("error extracting id:", err)
		errorGen.ErrorGen(&w, "invalid token", http.StatusInternalServerError)
		return
	}
	user.ID = id

	err = user_db.UpdateUser(user)
	if err != nil {
		log.Println("error updating user:", err)
		errorGen.ErrorGen(&w, "user already exists", http.StatusBadRequest)
		return
	}

	toDTO(&w, user)

	log.Println("user updated successfully")
	return
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
	_, err := verify_token.VerifyToken(w, r)
	if err != nil {
		log.Println("error verifying token:", err)
		return
	}
	cookieValue, _ := r.Cookie("X-Access-Token")

	token := cookieValue.Value
	id, err := tokenGen.ExtractIDFromToken(token)
	if err != nil {
		log.Println("error extracting id:", err)
		errorGen.ErrorGen(&w, "invalid token", http.StatusInternalServerError)
		return
	}
	err = user_db.DeleteUser(id)
	if err != nil {
		log.Println("error deleting user:", err)
		errorGen.ErrorGen(&w, "user not found", http.StatusBadRequest)
		return
	}
	cookie.ClearCookie(w, token)

	w.WriteHeader(http.StatusOK)
	log.Println("user deleted successfully")
	return
}
