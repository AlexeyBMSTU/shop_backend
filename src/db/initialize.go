package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"log"
)

var db *pgx.Conn

func InitializeDB() {
	var err error
	db, err = pgx.Connect(context.Background(), "postgres://username:password@172.20.0.2/database_name")

	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	err = createTables("users", createTableQueryUser)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}
}

func createTables(tableName string, query string) error {
	createTableQuery := query

	log.Printf("Creating table %s...", tableName)
	_, err := db.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Println("Error executing create table query:", err)
		return err
	}

	var exists bool
	checkTableQuery := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = '%s');", tableName)
	err = db.QueryRow(context.Background(), checkTableQuery).Scan(&exists)
	if err != nil {
		log.Println("Error checking if table exists:", err)
		return err
	}

	if exists {
		log.Printf("Table %s already exists.", tableName)
	} else {
		log.Printf("Table %s created successfully.", tableName)
	}

	return nil
}

func AddUser(user User.User) error {
	insertUserQuery := `
	INSERT INTO users (user_uid, name, password, email) 
	VALUES ($1, $2, $3, $4)
	RETURNING id;
	`
	var userID int
	err := db.QueryRow(context.Background(), insertUserQuery, user.ID, user.Username, user.Password, user.Email).Scan(&userID)
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
	row := db.QueryRow(context.Background(), query, name)
	var id uint64
	var email string
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
	user.Username = username
	user.Email = email
	user.Password = password
	return user, nil
}
