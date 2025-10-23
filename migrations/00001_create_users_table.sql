-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    telegram_id BIGINT UNIQUE NOT NULL,
    phone TEXT,
    username TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_users_telegram_id ON users(telegram_id);

-- +goose Down
DROP INDEX IF EXISTS idx_users_telegram_id;
DROP TABLE IF EXISTS users;