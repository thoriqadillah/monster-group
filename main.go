package main

import (
	"context"
	"fmt"

	"github.com/thoriqadillah/monster-group/dao/product"
	"github.com/thoriqadillah/monster-group/database"
	"github.com/thoriqadillah/monster-group/model"
)

func main() {
	ctx := context.Background()
	db := database.NewConnection()
	defer db.Close()

	product := product.NewDAO(ctx, db)

	// Get all products
	var productModel model.Product
	products, _ := product.GetAll(&productModel)

	fmt.Println(products)

	// Update price based on pricelist
	err := product.UpdatePrice()
	if err != nil {
		panic(err)
	}

	updatedProducts, _ := product.GetAll(&productModel)

	fmt.Println(updatedProducts)
}
