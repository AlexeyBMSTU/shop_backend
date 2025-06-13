package init

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

var db *pgx.Conn

func init() {
	log.Println("Pre-Creating users table...")
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
