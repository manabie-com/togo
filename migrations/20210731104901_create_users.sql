-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         serial PRIMARY KEY,
    username   TEXT      NOT NULL,
    password   TEXT      NOT NULL,
    max_todo   INTEGER            DEFAULT 5 NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX users_username_index ON users (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
