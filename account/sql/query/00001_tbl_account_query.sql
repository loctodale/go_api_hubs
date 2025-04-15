-- name: GetOneUserInfo :one
SELECT user_id, user_account, user_password, user_salt, user_role
FROM tbl_account
WHERE user_account = $1;

-- name: GetOneUserInfoAdmin :one
SELECT user_id, user_account, user_password, user_salt, user_login_time, user_logout_time, user_login_ip, user_created_at,user_updated_at, user_role
FROM tbl_account
WHERE user_account = $1;

-- name: CheckUserBaseExists :one
SELECT COUNT(*)
FROM tbl_account
WHERE user_account = $1;

-- name: AddUserBase :execresult
INSERT INTO tbl_account
(user_account, user_password, user_salt, user_role, user_created_at, user_updated_at)
VALUES ($1, $2, $3, $4, NOW(), NOW());

-- name: LoginUserBase :exec
UPDATE tbl_account
SET user_login_time = NOW(), user_logout_time = $1, user_login_ip = $2
WHERE user_account = $3 AND user_password = $4;

-- name: LogoutUserBase :exec
UPDATE tbl_account
SET user_logout_time = NOW()
WHERE user_account = $1;

