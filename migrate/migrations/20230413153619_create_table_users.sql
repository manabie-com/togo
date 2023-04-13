-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id         serial PRIMARY KEY,
    email VARCHAR(255) NOT NULL,
    password   TEXT      NOT NULL,
    limit_task   INTEGER            DEFAULT 5 NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX users_email_index ON users (email);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
