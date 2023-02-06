package product

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thoriqadillah/monster-group/model"
)

type productDAO struct {
	ctx context.Context
	db  *sql.DB
}

func NewDAO(ctx context.Context, db *sql.DB) productDAO {
	return productDAO{
		ctx: ctx,
		db:  db,
	}
}

func (p *productDAO) GetAll(product *model.Product) ([]model.Product, error) {
	query := "SELECT id, name, price, category, brand FROM products"
	rows, err := p.db.QueryContext(p.ctx, query)
	if err != nil {
		return nil, err
	}

	var products []model.Product

	for rows.Next() {
		err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Category, &product.Brand)
		if err != nil {
			return nil, err
		}

		products = append(products, *product)
	}

	defer rows.Close()

	return products, nil
}

// Update price based on pricelist
func (p *productDAO) UpdatePrice() error {
	query := "UPDATE products INNER JOIN pricelists ON products.category = pricelists.category AND products.brand = pricelists.brand SET products.price = pricelists.price;"
	_, err := p.db.ExecContext(p.ctx, query)
	if err != nil {
		return err
	}

	return nil
}
