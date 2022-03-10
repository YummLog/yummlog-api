-- name: ListFoodPosts :many
SELECT fp.*, pd.item, pd.experience
FROM foodposts fp join postdetails pd on fp.id = pd.post_id
ORDER BY fp.created_date, fp.id asc;

-- name: CreateFoodPost :one
insert into foodposts(
    id,
    restaurant_name,
    address1,
    address2,
    city,
    state,
    country,
    zipcode,
    user_id,
    created_by,
    created_date,
    updated_by,
    updated_date,
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
    post_id,
    item,
    experience
) values ($1, $2, $3, $4) returning *;

