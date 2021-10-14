CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "password" text NOT NULL,
  "max_todo" integer NOT NULL DEFAULT 5
);

CREATE TABLE "tasks" (
  "id" bigserial PRIMARY KEY NOT NULL,
  "content" text NOT NULL,
  "user_id" bigserial NOT NULL,
  "created_date" text NOT NULL
);

ALTER TABLE "tasks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
