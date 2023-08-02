-- +goose Up
INSERT INTO todos (task, completed) VALUES ('Buy milk', false);

-- +goose Down
DELETE FROM todos WHERE task = 'Buy milk';