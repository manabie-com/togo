-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
	"id" TEXT NOT NULL,
	"password" TEXT NOT NULL,
	"max_todo" INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);
CREATE INDEX "users_id_idx" ON  "users" ("id");
INSERT INTO "users" ("id", "password", "max_todo") VALUES('firstUser', 'example', 5);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "users";
-- +goose StatementEnd
