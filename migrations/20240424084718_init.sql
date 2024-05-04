-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    username VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(1000) NOT NULL
);

-- +goose Down
DROP TABLE users;
