-- +goose Up
-- +goose StatementBegin
CREATE TABLE "task"
(
    "id"         SERIAL PRIMARY KEY,
    "title"      varchar(255) NOT NULL,
    "user_id"    int,
    "created_at" timestamp    NOT NULL,
    "updated_at" timestamp
);

ALTER TABLE "task"
    ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "task";
-- +goose StatementEnd
