package main

import (
	"context"
	"fmt"

	"github.com/thoriqadillah/monster-group/dao/product"
	"github.com/thoriqadillah/monster-group/database"
)

func main() {
	ctx := context.Background()
	db := database.NewConnection()
	defer db.Close()

	product := product.NewDAO(ctx, db)

	products, _ := product.GetAll() // Get all products
	fmt.Println(products)

	err := product.UpdatePrice() // Update price based on pricelist
	if err != nil {
		panic(err)
	}

	updatedProducts, _ := product.GetAll()
	fmt.Println(updatedProducts)
}
