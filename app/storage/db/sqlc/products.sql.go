// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: products.sql

package db

import (
	"context"
	"database/sql"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
    name,
    article,
    brand,
    country_of_origin,
    description,
    price
) VALUES (
     $1, $2, $3, $4, $5, $6
 ) RETURNING id, name, article, brand, country_of_origin, description, price, rating, vote_count, vote_sum
`

type CreateProductParams struct {
	Name            string         `db:"name" json:"name"`
	Article         string         `db:"article" json:"article"`
	Brand           string         `db:"brand" json:"brand"`
	CountryOfOrigin string         `db:"country_of_origin" json:"country_of_origin"`
	Description     sql.NullString `db:"description" json:"description"`
	Price           string         `db:"price" json:"price"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, createProduct,
		arg.Name,
		arg.Article,
		arg.Brand,
		arg.CountryOfOrigin,
		arg.Description,
		arg.Price,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Article,
		&i.Brand,
		&i.CountryOfOrigin,
		&i.Description,
		&i.Price,
		&i.Rating,
		&i.VoteCount,
		&i.VoteSum,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :execrows
DELETE FROM products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int64) (int64, error) {
	result, err := q.db.ExecContext(ctx, deleteProduct, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getProduct = `-- name: GetProduct :one
SELECT id, name, article, brand, country_of_origin, description, price, rating, vote_count, vote_sum FROM products
WHERE id = $1
`

func (q *Queries) GetProduct(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProduct, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Article,
		&i.Brand,
		&i.CountryOfOrigin,
		&i.Description,
		&i.Price,
		&i.Rating,
		&i.VoteCount,
		&i.VoteSum,
	)
	return i, err
}

const getProductForUpdate = `-- name: GetProductForUpdate :one
SELECT id, name, article, brand, country_of_origin, description, price, rating, vote_count, vote_sum FROM products
WHERE id = $1 LIMIT 1
FOR UPDATE
`

func (q *Queries) GetProductForUpdate(ctx context.Context, id int64) (Product, error) {
	row := q.db.QueryRowContext(ctx, getProductForUpdate, id)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Article,
		&i.Brand,
		&i.CountryOfOrigin,
		&i.Description,
		&i.Price,
		&i.Rating,
		&i.VoteCount,
		&i.VoteSum,
	)
	return i, err
}

const listProduct = `-- name: ListProduct :many
SELECT id, name, article, brand, country_of_origin, description, price, rating, vote_count, vote_sum FROM products
ORDER BY id DESC
LIMIT $1
OFFSET $2
`

type ListProductParams struct {
	Limit  int32 `db:"limit" json:"limit"`
	Offset int32 `db:"offset" json:"offset"`
}

func (q *Queries) ListProduct(ctx context.Context, arg ListProductParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, listProduct, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Article,
			&i.Brand,
			&i.CountryOfOrigin,
			&i.Description,
			&i.Price,
			&i.Rating,
			&i.VoteCount,
			&i.VoteSum,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listProductCount = `-- name: ListProductCount :one
SELECT count(*) FROM products
`

func (q *Queries) ListProductCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, listProductCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products SET
    name = $2,
    article  = $3,
    brand = $4,
    country_of_origin = $5,
    description = $6,
    price = $7
WHERE id = $1 RETURNING id, name, article, brand, country_of_origin, description, price, rating, vote_count, vote_sum
`

type UpdateProductParams struct {
	ID              int64          `db:"id" json:"id"`
	Name            string         `db:"name" json:"name"`
	Article         string         `db:"article" json:"article"`
	Brand           string         `db:"brand" json:"brand"`
	CountryOfOrigin string         `db:"country_of_origin" json:"country_of_origin"`
	Description     sql.NullString `db:"description" json:"description"`
	Price           string         `db:"price" json:"price"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, updateProduct,
		arg.ID,
		arg.Name,
		arg.Article,
		arg.Brand,
		arg.CountryOfOrigin,
		arg.Description,
		arg.Price,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Article,
		&i.Brand,
		&i.CountryOfOrigin,
		&i.Description,
		&i.Price,
		&i.Rating,
		&i.VoteCount,
		&i.VoteSum,
	)
	return i, err
}
