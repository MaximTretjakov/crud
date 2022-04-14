-- +goose Up
CREATE TABLE IF NOT EXISTS blog (
    id SERIAL,
    title varchar(100),
    text varchar(100),
    tags varchar(20),
    created_at timestamptz default now(),
    updated_at timestamptz 
);

-- +goose Down
DROP TABLE blog;