package pricelist

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thoriqadillah/monster-group/model"
)

type pricelistDAO struct {
	ctx       context.Context
	db        *sql.DB
	pricelist model.Pricelist
}

func NewModel(ctx context.Context, db *sql.DB) pricelistDAO {
	return pricelistDAO{
		ctx:       ctx,
		db:        db,
		pricelist: model.Pricelist{},
	}
}

func (p *pricelistDAO) GetAll() ([]model.Pricelist, error) {
	query := "SELECT id, price, category, brand FROM pricelists"
	rows, err := p.db.QueryContext(p.ctx, query)
	if err != nil {
		return nil, err
	}

	var pricelists []model.Pricelist

	for rows.Next() {
		err := rows.Scan(&p.pricelist.Id, &p.pricelist.Price, &p.pricelist.Category, &p.pricelist.Brand)
		if err != nil {
			return nil, err
		}

		pricelists = append(pricelists, p.pricelist)
	}

	defer rows.Close()

	return pricelists, nil
}
