-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user"
(
    "id"             SERIAL PRIMARY KEY,
    "username"       varchar(20) UNIQUE NOT NULL,
    "password"       varchar(72)        NOT NULL,
    "max_daily_task" int DEFAULT 5,
    "created_at"     timestamp          NOT NULL,
    "updated_at"     timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user";
-- +goose StatementEnd
