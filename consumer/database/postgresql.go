package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	url := os.Getenv("DATABASE_URL")
	var err error
	DB, err = sql.Open("postgres", url)
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Ping database error:", err)
	}
	fmt.Println("Connected to the database successfully!")
}

func CreateTable() {
	createTb := `
	CREATE TABLE IF NOT EXISTS skill (
		key TEXT PRIMARY KEY,
		name TEXT NOT NULL DEFAULT '',
		description TEXT NOT NULL DEFAULT '',
		logo TEXT NOT NULL DEFAULT '',
		tags TEXT [] NOT NULL DEFAULT '{}'
	);
    `
	_, err := DB.Exec(createTb)
	if err != nil {
		log.Fatal("Can't create table", err)
	}
	fmt.Println("Create table success")
}