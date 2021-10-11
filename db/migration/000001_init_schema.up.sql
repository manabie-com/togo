CREATE TABLE "users" (
  "id" uuid PRIMARY KEY NOT NULL,
  "password" text NOT NULL,
  "max_todo" integer NOT NULL DEFAULT 5
);

CREATE TABLE "tasks" (
  "id" uuid PRIMARY KEY NOT NULL,
  "content" text NOT NULL,
  "user_id" uuid NOT NULL,
  "created_date" text NOT NULL
);

ALTER TABLE "tasks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
