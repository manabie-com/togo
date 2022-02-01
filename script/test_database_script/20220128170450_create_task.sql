-- +goose Up
-- +goose StatementBegin
CREATE TABLE tasks (
                       id SERIAL NOT NULL,
                       content TEXT NOT NULL,
                       user_id INT NOT NULL,
                       created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       deleted_at timestamp,
                       PRIMARY KEY(id),
                       CONSTRAINT fk_user
                           FOREIGN KEY(user_id)
                               REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks;
-- +goose StatementEnd
