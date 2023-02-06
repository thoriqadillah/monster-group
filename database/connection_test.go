package database

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestOpenConnection(t *testing.T) {
	db := NewConnection()
	defer db.Close()
}
