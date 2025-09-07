package user_db

import (
	"context"
	"errors"
	"github.com/AlexeyBMSTU/shop_backend/src/db"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
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

	log.Printf("login 1.3 attempt for user: %s\n", user.Name)
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
	log.Printf("login 1.4 attempt for user: %s\n", user.Name)
	user.ID = userUID
	user.Name = username
	user.Password = password
	if email != nil {
		user.Email = email
	}

	return user, nil
}

func UserExists(userUID uuid.UUID) (bool, error) {
	query := `SELECT 1 FROM users WHERE user_uid = $1 LIMIT 1;`
	var exists int
	err := db.Database.QueryRow(context.Background(), query, userUID).Scan(&exists)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return false, err
		}
		// Если нет строк — пользователь не найден
		if err.Error() == "no rows in result set" {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func UpdateUser(user User.User) error {
	exists, err := UserExists(user.ID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user not found")
	}

	setParts := []string{}
	args := []interface{}{}
	argPos := 1

	if user.Name != "" {
		setParts = append(setParts, "name = $"+strconv.Itoa(argPos))
		args = append(args, user.Name)
		argPos++
	}
	if user.Password != "" {
		setParts = append(setParts, "password = $"+strconv.Itoa(argPos))
		args = append(args, user.Password)
		argPos++
	}
	if user.Email != nil {
		setParts = append(setParts, "email = $"+strconv.Itoa(argPos))
		args = append(args, user.Email)
		argPos++
	}

	if len(setParts) == 0 {
		return nil
	}

	query := "UPDATE users SET " + strings.Join(setParts, ", ") + " WHERE user_uid = $" + strconv.Itoa(argPos)
	args = append(args, user.ID)

	_, err = db.Database.Exec(context.Background(), query, args...)
	if err != nil {
		return err
	}
	log.Printf("user updated successfully with UUID: %s\n", user.ID.String())
	return nil
}

func DeleteUser(userUID uuid.UUID) error {
	exists, err := UserExists(userUID)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New("user not found")
	}

	deleteQuery := `
	DELETE FROM users
	WHERE user_uid = $1;
	`
	_, err = db.Database.Exec(context.Background(), deleteQuery, userUID)
	if err != nil {
		return err
	}
	log.Printf("user deleted successfully with UUID: %s\n", userUID.String())
	return nil
}
