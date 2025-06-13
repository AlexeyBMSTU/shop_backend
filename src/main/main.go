package main

import (
	"github.com/AlexeyBMSTU/shop_backend/src/db/init"
	"github.com/AlexeyBMSTU/shop_backend/src/internal/http/auth"
	"github.com/AlexeyBMSTU/shop_backend/src/middleware/verify_token"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	init.InitializeDB()
	router := mux.NewRouter()
	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	router.Handle("/protected", verify_token.VerifyToken(http.HandlerFunc(ProtectedHandler))).Methods("GET")

	http.ListenAndServe(":10000", router)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the protected area!"))
}
