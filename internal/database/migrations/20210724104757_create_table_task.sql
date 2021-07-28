-- +goose Up
-- +goose StatementBegin
CREATE TABLE "tasks" (
	"id" TEXT NOT NULL,
	"content" TEXT NOT NULL,
	"user_id" TEXT NOT NULL,
    "created_date" TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE INDEX "tasks_id_idx" ON  "tasks" ("id");
CREATE INDEX "tasks_user_id_idx" ON  "tasks" ("user_id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "tasks";
-- +goose StatementEnd
