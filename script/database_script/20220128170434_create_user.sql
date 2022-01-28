-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
                       id SERIAL NOT NULL,
                       password TEXT NOT NULL,
                       limit_task INT8 DEFAULT 3 NOT NULL,
                       created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       deleted_at timestamp,
                       PRIMARY KEY(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
