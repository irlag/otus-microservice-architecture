-- name: GetProduct :one
SELECT * FROM products
WHERE id = $1;

-- name: GetProductForUpdate :one
SELECT * FROM products
WHERE id = $1 LIMIT 1
FOR UPDATE;

-- name: ListProduct :many
SELECT * FROM products
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: ListProductCount :one
SELECT count(*) FROM products;

-- name: CreateProduct :one
INSERT INTO products (
    name,
    article,
    brand,
    country_of_origin,
    description,
    price,
    version
) VALUES (
     $1, $2, $3, $4, $5, $6, $7
 ) RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET
    name = $2,
    article  = $3,
    brand = $4,
    country_of_origin = $5,
    description = $6,
    price = $7,
    version = $8
WHERE id = $1 RETURNING *;

-- name: DeleteProduct :execrows
DELETE FROM products
WHERE id = $1;