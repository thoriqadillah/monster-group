package database

import (
	"database/sql"
	"time"
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
