-- name: CreateApis :execresult
INSERT INTO tbl_apis
(provider_id, name, slug, description, category, base_url, doc_url, status, created_at)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW());

-- name: GetApis :many
SELECT id, provider_id, name, slug, description, category, base_url, doc_url, status, created_at
FROM tbl_apis;

-- name: GetApiById :one
SELECT id, provider_id, name, slug, description, category, base_url, doc_url, status, created_at
FROM tbl_apis
WHERE id = $1;
