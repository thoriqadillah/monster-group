package pricelist

import (
	"context"
	"fmt"
	"testing"

	"github.com/thoriqadillah/monster-group/database"
	"github.com/thoriqadillah/monster-group/model"
)

func TestGetAll(t *testing.T) {
	ctx := context.Background()
	db := database.NewConnection()
	defer db.Close()

	pricelist := NewModel(ctx, db)

	var model model.Pricelist
	pricelists, err := pricelist.GetAll(&model)
	if err != nil {
		panic(err)
	}

	fmt.Println(pricelists)
}
