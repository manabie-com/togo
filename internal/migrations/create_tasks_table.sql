-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tasks
(
    id           TEXT PRIMARY KEY,
    content      TEXT                       NOT NULL,
    created_date TEXT                       NOT NULL,
    user_id      TEXT REFERENCES users (id) NOT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS tasks CASCADE;