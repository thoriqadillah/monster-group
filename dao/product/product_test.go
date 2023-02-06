package product

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thoriqadillah/monster-group/database"
)

func TestGetAll(t *testing.T) {
	ctx := context.Background()
	db := database.NewConnection()
	defer db.Close()

	product := NewDAO(ctx, db)

	products, err := product.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(products)
}

func TestUpdatePrice(t *testing.T) {
	ctx := context.Background()
	db := database.NewConnection()
	defer db.Close()

	product := NewDAO(ctx, db)

	err := product.UpdatePrice()
	if err != nil {
		panic(err)
	}
}
