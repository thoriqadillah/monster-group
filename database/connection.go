package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

func NewConnection() *sql.DB {
	var (
		USERNAME = os.Getenv("DB_USERNAME")
		PASSWORD = os.Getenv("DB_PASSWORD")
		HOST     = os.Getenv("DB_HOST")
		SCHEMA   = os.Getenv("DB_SCHEMA")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", USERNAME, PASSWORD, HOST, SCHEMA)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(5 * time.Minute)

	return db
}
