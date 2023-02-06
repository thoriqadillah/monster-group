package database

import (
	"database/sql"
	"os"
	"time"
)

var (
	USERNAME = os.Getenv("DB_USERNAME")
	PASSWORD = os.Getenv("DB_PASSWORD")
	HOST     = os.Getenv("DB_HOST")
	SCHEMA   = os.Getenv("DB_SCHEMA")
)

func NewConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/monster_group")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
