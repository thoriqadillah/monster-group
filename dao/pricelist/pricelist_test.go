package pricelist

import (
	"context"
	"fmt"
	"testing"

	"github.com/thoriqadillah/monster-group/database"
)

func TestGetAll(t *testing.T) {
	ctx := context.Background()
	db := database.NewConnection()
	defer db.Close()

	pricelist := NewModel(ctx, db)

	pricelists, err := pricelist.GetAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(pricelists)
}
