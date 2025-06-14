package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	user_db "github.com/AlexeyBMSTU/shop_backend/src/db/user"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"golang.org/x/crypto/bcrypt"
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
	testUser := User.User{
		Name:     "testuser",
		Password: "password123",
		Email:    new(string),
	}
	*testUser.Email = "test@example.com"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(testUser.Password), bcrypt.DefaultCost)
	testUser.Password = string(hashedPassword)

	user_db.AddUser = mockAddUser
	user_db.GetUserByName = mockGetUserByName

	if err := mockAddUser(testUser); err != nil {
		t.Fatalf("failed to add test user: %v", err)
	}

	tests := []struct {
		name       string
		body       User.User
		wantStatus int
	}{
		{"Valid Login", User.User{Name: "testuser", Password: "password123"}, http.StatusOK},
		{"Invalid Password", User.User{Name: "testuser", Password: "wrongpassword"}, http.StatusUnauthorized},
		{"User  Not Found", User.User{Name: "nonexistent", Password: "password123"}, http.StatusBadRequest},
		{"Missing Credentials", User.User{Name: "", Password: ""}, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			LoginHandler(w, req)

			res := w.Result()
			if res.StatusCode != tt.wantStatus {
				t.Errorf("got %v, want %v", res.StatusCode, tt.wantStatus)
			}
		})
	}
}

func TestRegisterHandler(t *testing.T) {
	user_db.AddUser = mockAddUser
	user_db.GetUserByName = mockGetUserByName

	tests := []struct {
		name       string
		body       User.User
		wantStatus int
	}{
		{"Valid Registration", User.User{Name: "newuser", Password: "newpassword123", Email: new(string)}, http.StatusCreated},
		{"Duplicate User", User.User{Name: "testuser", Password: "password123", Email: new(string)}, http.StatusBadRequest},
		{"Missing Credentials", User.User{Name: "", Password: ""}, http.StatusBadRequest},
	}

	*tests[0].body.Email = "new@example.com"
	*tests[1].body.Email = "test@example.com"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.body)
			req := httptest.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
			w := httptest.NewRecorder()

			RegisterHandler(w, req)

			res := w.Result()
			if res.StatusCode != tt.wantStatus {
				t.Errorf("got %v, want %v", res.StatusCode, tt.wantStatus)
			}
		})
	}
}
