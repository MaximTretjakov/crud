-- +goose Up
CREATE TABLE IF NOT EXISTS blog (
    id int NOT NULL,
    title text,
    text text,
    tags text,
    created_at timestamp,
    updated_at timestamp
);

-- +goose Down
DROP TABLE blog;