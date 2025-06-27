-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tbl_token (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL UNIQUE,
    public_key TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    CONSTRAINT fk_user
    FOREIGN KEY (user_id)
    REFERENCES tbl_account(user_id)
    ON DELETE CASCADE
    ON UPDATE CASCADE
    );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS tbl_token;
-- +goose StatementEnd
