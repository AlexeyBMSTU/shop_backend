package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"os"
)

func InitializeDB() {
	var err error

	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	databaseName := os.Getenv("POSTGRES_DB")

	databaseURL := fmt.Sprintf("postgres://%s:%s@db:5432/%s", username, password, databaseName)
	Database, err = pgx.Connect(context.Background(), databaseURL)
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
	_, err := Database.Exec(context.Background(), createTableQuery)
	if err != nil {
		log.Println("Error executing create table query:", err)
		return err
	}

	var exists bool
	checkTableQuery := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = '%s');", tableName)
	err = Database.QueryRow(context.Background(), checkTableQuery).Scan(&exists)
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
