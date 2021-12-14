-- name: ResisterCurrency :one
INSERT INTO currencies (
  coin,
  name
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetCurrency :one
SELECT * FROM currencies
WHERE name = $1 LIMIT 1;

-- name: ListCurrencies :many
SELECT * FROM currencies
ORDER BY id
LIMIT $1
OFFSET $2;
