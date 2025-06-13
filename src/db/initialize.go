package db

import (
	"context"
	"errors"
	"github.com/AlexeyBMSTU/shop_backend/src/models/User"
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
	// Создаем таблицы
	err = createTables()
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}
}

// createTables создает необходимые таблицы в базе данных

func createTables() error {
	// SQL-запрос для создания таблицы users
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL
	);
	`

	log.Println("Creating users table...")
	// Выполняем запрос
	_, err := db.Exec(context.Background(), createUsersTable)
	if err != nil {
		return err
	}

	log.Println("Users table created successfully.")
	return nil
}

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

func GetUserByName(name string) (User.User, error) {
	var user User.User
	query := `
	SELECT id, name, email
	FROM users
	WHERE name = $1;
	`
	row := db.QueryRow(context.Background(), query, name)
	var id uint64
	var email string
	var username string
	err := row.Scan(&id, &username, &email)
	if err != nil {
		if errors.Is(err, context.Canceled) {
			return user, err
		}
		return user, err
	}
	user.ID = id
	user.Username = username
	user.Email = email
	return user, nil
}
