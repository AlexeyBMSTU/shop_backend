package http_routes

import (
	"github.com/AlexeyBMSTU/shop_backend/src/internal/http/auth"
	"github.com/AlexeyBMSTU/shop_backend/src/middleware/verify_token"
	"net/http"
)

type Route struct {
	Handler http.Handler
	Method  string
}

func GetRoutes() map[string]Route {
	return map[string]Route{
		"/login": {
			Handler: http.HandlerFunc(auth.LoginHandler),
			Method:  "POST",
		},
		"/registration": {
			Handler: http.HandlerFunc(auth.RegisterHandler),
			Method:  "POST",
		},
		"/protected": {
			Handler: verify_token.VerifyToken(http.HandlerFunc(ProtectedHandler)),
			Method:  "GET",
		},
	}
}

func ProtectedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome322 to the protected area"))
}
