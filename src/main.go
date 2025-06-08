package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"shop_backend/src/internal/http/auth"
	"shop_backend/src/middleware/verify_token"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/login", auth.LoginHandler).Methods("POST")
	router.Handle("/protected", verify_token.VerifyToken(http.HandlerFunc(ProtectedHandler))).Methods("GET")

	http.ListenAndServe(":10000", router)
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the protected area!"))
}
