package main

import (
	"github.com/AlexeyBMSTU/shop_backend/src/db"
	http_routes "github.com/AlexeyBMSTU/shop_backend/src/internal/http"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	db.InitializeDB()
	router := mux.NewRouter()
	routes := http_routes.GetRoutes()

	for path, route := range routes {
		router.Handle(path, route.Handler).Methods(route.Method)
	}

	http.ListenAndServe(":10000", router)
	log.Println("Starting server on port 10000")
}
