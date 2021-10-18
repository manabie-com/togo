CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "password" text NOT NULL,
  "max_todo" integer NOT NULL DEFAULT 5
);

CREATE TABLE "tasks" (
  "id" bigserial PRIMARY KEY,
  "content" text NOT NULL,
  "user_id" varchar NOT NULL,
  "created_date" text NOT NULL
);

ALTER TABLE "tasks" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("username");

INSERT INTO users (username, password) VALUES('testUser', 'password')