-- name: CreateCar :one
INSERT INTO cars (
  id ,
  owner_id,
  title ,
  model_id ,
  color ,
  model_name,
  description,
  tags,
  listing_status,
  images,
  owner_username
) VALUES (
  $1, $2, $3, $4,$5,$6,$7,$8,$9,$10,$11
) RETURNING *;

-- name: AddImagesToCar :one
UPDATE cars
SET  images= $2
where id=$1
RETURNING *;

-- name: DeleteCar :exec
DELETE FROM cars
where id = $1 ;

-- name: GetCarById :one
SELECT * FROM cars where id = $1;

-- name: SearchCarsFTS :many
-- go-type: title string
-- go-type: description string
-- go-type: model_name string
-- go-type: model_id string
-- go-type: color string
-- go-type: tags []string
SELECT *
FROM cars
WHERE 
ID != $1 AND 
    lower(title) LIKE '%' || $2 || '%'
    OR lower(description) LIKE '%' || $2 || '%'
    OR lower(model_name) LIKE '%' || $2 || '%'
    OR lower(model_id) LIKE '%' || $2 || '%'
    OR lower(color) LIKE '%' ||  $2 || '%'
    OR $2 = ANY(tags)
OFFSET $3
LIMIT $4;

-- name: GetUserOwnedCars :many
SELECT *
from cars 
where owner_id = $1 ;

-- name: UpdateCar :one
UPDATE cars
SET
  title = COALESCE(sqlc.narg(title), title),
  tags = COALESCE(sqlc.narg(tags), tags),
  listing_status = COALESCE(sqlc.narg(listing_status), listing_status),
  description = COALESCE(sqlc.narg(description), description),
  color = COALESCE(sqlc.narg(color), color),
  model_id = COALESCE(sqlc.narg(model_id), model_id),
  images = COALESCE(sqlc.narg(images), images),

  model_name = COALESCE(sqlc.narg(model_name), model_name)
  

WHERE
  id = sqlc.arg(id)
RETURNING *;

