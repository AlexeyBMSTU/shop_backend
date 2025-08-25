package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
)

var mockUserDB = make(map[string]User.User)

func mockAddUser(user User.User) error {
	if _, exists := mockUserDB[user.Name]; exists {
		return fmt.Errorf("user already exists")
	}
	mockUserDB[user.Name] = user
	return nil
}

func mockGetUserByName(name string) (User.User, error) {
	user, exists := mockUserDB[name]
	if !exists {
		return User.User{}, fmt.Errorf("user not found")
	}
	return user, nil
}

func TestLoginHandler(t *testing.T) {
	//testUser := User.User{
	//	Name:     "testuser",
	//	Password: "password123",
	//	Email:    nil,
	//}
	//hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(testUser.Password), bcrypt.DefaultCost)
	//testUser.Password = string(hashedPassword)
	//
	////user_db.GetUserByName(testUser.Name)
	//if err := mockAddUser(testUser); err != nil {
	//	t.Fatalf("failed to add test user: %v", err)
	//}
	email := "test@test.com"
	tests := []struct {
		name       string
		body       User.User
		wantStatus int
	}{
		{"Valid Login", User.User{Name: "testuser", Password: "password123", Email: &email}, http.StatusOK},
		//{"Invalid Password", User.User{Name: "testuser", Password: "wrongpassword"}, http.StatusUnauthorized},
		//{"User  Not Found", User.User{Name: "nonexistent", Password: "password123"}, http.StatusBadRequest},
		//{"Missing Credentials", User.User{Name: "", Password: ""}, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println("эййй")

			body, _ := json.Marshal(tt.body)
			log.Println(tt.body)
			log.Println(string(body))
			req := httptest.NewRequest(http.MethodPost, "/api/v1/login/", bytes.NewBuffer(body))
			w := httptest.NewRecorder()
			log.Println(w.Body)
			LoginHandler(w, req)
			res := w.Result()
			if res.StatusCode != tt.wantStatus {
				t.Errorf("got %v, want %v", res.StatusCode, tt.wantStatus)
			}
		})
	}
}

//
//func TestRegisterHandler(t *testing.T) {
//
//	tests := []struct {
//		name       string
//		body       User.User
//		wantStatus int
//	}{
//		{"Valid Registration", User.User{Name: "newuser", Password: "newpassword123"}, http.StatusCreated},
//		{"Duplicate User", User.User{Name: "testuser", Password: "password123"}, http.StatusBadRequest},
//		{"Missing Credentials", User.User{Name: "", Password: ""}, http.StatusBadRequest},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			body, _ := json.Marshal(tt.body)
//			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
//			w := httptest.NewRecorder()
//
//			RegisterHandler(w, req)
//
//			res := w.Result()
//			if res.StatusCode != tt.wantStatus {
//				t.Errorf("got %v, want %v", res.StatusCode, tt.wantStatus)
//			}
//		})
//	}
//}
