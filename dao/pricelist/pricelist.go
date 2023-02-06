package pricelist

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thoriqadillah/monster-group/model"
)

type pricelistDAO struct {
	ctx context.Context
	db  *sql.DB
}

func NewModel(ctx context.Context, db *sql.DB) pricelistDAO {
	return pricelistDAO{
		ctx: ctx,
		db:  db,
	}
}

func (p *pricelistDAO) GetAll(pricelist *model.Pricelist) ([]model.Pricelist, error) {
	query := "SELECT id, price, category, brand FROM pricelists"
	rows, err := p.db.QueryContext(p.ctx, query)
	if err != nil {
		return nil, err
	}

	var pricelists []model.Pricelist

	for rows.Next() {
		err := rows.Scan(&pricelist.Id, &pricelist.Price, &pricelist.Category, &pricelist.Brand)
		if err != nil {
			return nil, err
		}

		pricelists = append(pricelists, *pricelist)
	}

	defer rows.Close()

	return pricelists, nil
}
