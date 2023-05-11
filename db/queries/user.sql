-- name: GetUser :one
SELECT * FROM public."user"
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM public."user"
ORDER BY name;

-- name: CreateUser :one
INSERT INTO public."user" (
  name, bio
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM public."user"
WHERE id = $1;