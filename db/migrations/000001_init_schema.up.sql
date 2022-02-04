CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "full_name" varchar NOT NULL UNIQUE,
  "maximum" int NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tasks" (
  "id" SERIAL PRIMARY KEY,
  "title" varchar NOT NULL,
  "content" varchar NOT NULL,
  "is_complete" boolean NOT NULL,
  "user_id" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "tasks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

CREATE INDEX ON "tasks" ("user_id");