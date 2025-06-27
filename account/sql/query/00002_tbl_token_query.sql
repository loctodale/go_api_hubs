-- name: CreateNewToken :exec
INSERT INTO tbl_token (user_id, public_key, refresh_token)
VALUES ($1, $2, $3)
    ON CONFLICT (user_id)
DO UPDATE SET
    public_key = EXCLUDED.public_key,
           refresh_token = EXCLUDED.refresh_token;

-- name: GetTokenByUser :one
SELECT user_id, public_key, refresh_token
FROM tbl_token
WHERE user_id = $1;