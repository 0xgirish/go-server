-- +goose Up
CREATE TABLE todos (
    id SERIAL PRIMARY KEY,
    task TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT false
);

-- +goose Down
DROP TABLE todos;