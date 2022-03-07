-- name: ListFoodPosts :many
SELECT *
FROM foodposts fp join postdetails pd on fp.id = pd."postId";

-- name: CreateFoodPost :one
insert into foodposts(
    id,
    "restaurantName",
    address1,
    address2,
    city,
    state,
    country,
    zipcode,
    "userId",
    "createdBy",
    "createdDate",
    "updatedBy",
    "updatedDate",
    notes
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14
) returning *;

-- name: ListPostDetails :many
SELECT *
FROM postdetails;

-- name: CreatePostDetails :one
insert into postdetails(
    id,
    "postId",
    item,
    experience
) values ($1, $2, $3, $4) returning *;

