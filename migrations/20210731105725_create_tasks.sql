-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks
(
    id           serial PRIMARY KEY,
    content      TEXT      NOT NULL,
    user_id      int4      NOT NULL,
    created_date DATE      NOT NULL DEFAULT CURRENT_DATE,
    completed    BOOLEAN   NOT NULL DEFAULT FALSE,
    created_at   TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at   TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT tasks_fk FOREIGN KEY (user_id) REFERENCES users (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks
-- +goose StatementEnd
