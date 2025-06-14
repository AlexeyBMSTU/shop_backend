package user_db

import (
	"context"
	"errors"
	"github.com/AlexeyBMSTU/shop_backend/src/db"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"github.com/google/uuid"
	"log"
)

func AddUser(user User.User) error {
	insertUserQuery := `
	INSERT INTO users (user_uid, name, password, email) 
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`
	var userID int
	err := db.Database.QueryRow(context.Background(), insertUserQuery, user.ID, user.Name, user.Password, &user.Email).Scan(&userID)
	if err != nil {
		return err
	}
	log.Printf("user  added successfully with ID: %d\n", userID)
	return nil
}

func GetUserByName(name string) (User.User, error) {
	var user User.User

	query := `
	SELECT id, user_uid, name, password, email
	FROM users
	WHERE name = $1;
	`

	row := db.Database.QueryRow(context.Background(), query, name)

	var id uint64
	var email *string
	var username string
	var userUID uuid.UUID
	var password string

	err := row.Scan(&id, &userUID, &username, &password, &email)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return user, err
		}
		return user, err
	}

	user.ID = userUID
	user.Name = username
	user.Password = password
	user.Email = email

	return user, nil
}
