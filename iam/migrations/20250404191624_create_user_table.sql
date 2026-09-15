-- +goose Up

CREATE TABLE users (
    user_uuid UUID PRIMARY KEY REFERENCES users(uuid) ON DELETE CASCADE,
    login VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    notification_methods JSONB NOT NULL DEFAULT '[]'::jsonb,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS user_infos;
DROP TABLE IF EXISTS users;
