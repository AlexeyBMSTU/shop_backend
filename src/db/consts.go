package db

import "github.com/jackc/pgx/v4"

var Database *pgx.Conn

const createTableQueryUser = `
CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		user_uid VARCHAR(100) UNIQUE NOT NULL,
		name VARCHAR(100) UNIQUE NOT NULL,
    	password VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE
	);
`
