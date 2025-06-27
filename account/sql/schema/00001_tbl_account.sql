-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS tbl_account (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_account VARCHAR(255) NOT NULL UNIQUE,
    user_password VARCHAR(255) NOT NULL,
    user_salt VARCHAR(255) NOT NULL,

    user_login_time TIMESTAMP NULL,
    user_logout_time TIMESTAMP NULL,
    user_login_ip VARCHAR(45),
    user_role INT NOT NULL,

    user_created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );

-- Trigger để cập nhật user_updated_at khi update dòng
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.user_updated_at = CURRENT_TIMESTAMP;
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER set_updated_at
    BEFORE UPDATE ON tbl_account
    FOR EACH ROW
    EXECUTE PROCEDURE update_updated_at_column();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_account;
-- +goose StatementEnd
