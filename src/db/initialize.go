package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
)

func InitializeDB() {
	var err error
	Database, err = pgx.Connect(context.Background(), "postgres://username:password@172.20.0.2/database_name")

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
