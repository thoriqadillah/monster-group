package product

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thoriqadillah/monster-group/database"
	"github.com/thoriqadillah/monster-group/model"
)

func TestGetAll(t *testing.T) {
	ctx := context.Background()
	db := database.NewConnection()
	defer db.Close()

	product := NewDAO(ctx, db)

	var model model.Product
	products, err := product.GetAll(&model)
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
