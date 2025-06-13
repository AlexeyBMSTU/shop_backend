package db

import (
	"context"
	"log"
)

func AddUser(name string, email string) error {
	// SQL-запрос для вставки нового пользователя
	insertUserQuery := `
	INSERT INTO users (name, email) 
	VALUES ($1, $2)
	RETURNING id;
	`
	var userID int
	err := db.QueryRow(context.Background(), insertUserQuery, name, email).Scan(&userID)
	if err != nil {
		return err
	}
	log.Printf("User  added successfully with ID: %d\n", userID)
	return nil
}
