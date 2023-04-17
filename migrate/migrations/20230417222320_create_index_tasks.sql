-- +goose Up
-- +goose StatementBegin
CREATE INDEX ON "tasks" ("user_id" );
CREATE INDEX ON "tasks" ("is_done" );
CREATE INDEX ON "tasks" ("created_date" );
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX "tasks_user_id_idx";
DROP INDEX "tasks_is_done_idx";
DROP INDEX "tasks_created_date_idx";
-- +goose StatementEnd
