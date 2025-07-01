-- +goose Up
-- +goose StatementBegin
CREATE TYPE api_status AS ENUM('active', 'disable', 'pending');
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS tbl_apis (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    provider_id UUID NOT NULL,
    name varchar(255) NOT NULL,
    slug varchar(255) NOT NULL,
    description TEXT,
    category    TEXT,
    base_url    TEXT,
    doc_url     TEXT,
    status      api_status,
    created_at  TIMESTAMP
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_apis;
DROP TYPE api_status;
-- +goose StatementEnd
