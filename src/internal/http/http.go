package http_routes

import (
	"github.com/AlexeyBMSTU/shop_backend/src/internal/http/auth"
	"net/http"
)

type Route struct {
	Handler http.Handler
	Method  string
}

func GetRoutes() map[string]Route {
	return map[string]Route{
		"/api/v1/auth/login/": {
			Handler: http.HandlerFunc(auth.LoginHandler),
			Method:  "POST",
		},
		"/api/v1/auth/registration/": {
			Handler: http.HandlerFunc(auth.RegisterHandler),
			Method:  "POST",
		},
		"/api/v1/auth/logout/": {
			Handler: http.HandlerFunc(auth.LogoutHandler),
			Method:  "POST",
		},
		"/api/v1/auth/me/": {
			Handler: http.HandlerFunc(auth.MeHandler),
			Method:  "GET",
		},
	}
}
