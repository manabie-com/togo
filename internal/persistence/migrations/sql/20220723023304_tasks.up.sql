CREATE TABLE "tasks"
(
    "id"         UUID      NOT NULL,
    PRIMARY KEY ("id"),
    "user_id"    UUID      NOT NULL,
    "title"      TEXT,
    "note"       TEXT,
    "status"     TEXT,
    "due_date"   TIMESTAMP,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL
)
