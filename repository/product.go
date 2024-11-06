package repository

import (
	"database/sql"
	"math"
	"project/model"
)

type Product struct {
	Db *sql.DB
}

func CreateProduct(db *sql.DB) *Product {
	return &Product{Db: db}
}

func (repo *Product) Create(product *model.Product) error {
	query := `INSERT INTO products (name, category_id, location_id) SELECT $1, $2, $3 FROM sessions WHERE id=$4 AND expired_at IS NULL RETURNING id`
	err := repo.Db.QueryRow(query, product.Name, product.CategoryId, product.LocationId, product.SessionId).Scan(&product.Id)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Product) Update(product *model.Product) error {
	query := `UPDATE products SET balance=$2, updated_at=NOW()
				FROM sessions
				WHERE products.id=$1 AND sessions.id=$3 AND expired_at IS NULL
				RETURNING products.name`
	err := repo.Db.QueryRow(query, product.Id, product.Balance, product.SessionId).Scan(&product.Name)
	if err != nil {
		return err
	}
	return nil
}

func (repo *Product) Search(criteria model.Search) (int, int, int, float64, []model.Product, error) {
	pattern := "%" + criteria.Name + "%"

	var count int
	queryCount := `SELECT COUNT(*)
				FROM products
				WHERE products.name ILIKE $1`
	err := repo.Db.QueryRow(queryCount, pattern).Scan(&count)

	query := `SELECT products.id, products.name
				FROM products
				WHERE products.name ILIKE $1
				LIMIT $2
				OFFSET $3`

	offset := (criteria.Page - 1) * criteria.Limit
	rows, err := repo.Db.Query(query, pattern, criteria.Limit, offset)
	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.Id, &product.Name); err != nil {
			return 0, 0, 0, 0, []model.Product{}, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return 0, 0, 0, 0, []model.Product{}, err
	}
	return criteria.Page, criteria.Limit, count, math.Ceil(float64(count) / float64(criteria.Limit)), products, nil
}
