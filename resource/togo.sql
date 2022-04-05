CREATE TABLE "users" (
  "username" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "daily_cap" bigint NOT NULL DEFAULT (0),
  "daily_quantity" bigint NOT NULL DEFAULT (0),
  "password_change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tasks" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "owner" varchar NOT NULL,
  "content" varchar NOT NULL,
  "deleted" bool NOT NULL DEFAULT (false),
  "content_change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "deleted_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "tasks" ADD FOREIGN KEY ("owner") REFERENCES "users" ("username");

CREATE INDEX ON "users" ("full_name");

CREATE INDEX ON "users" ("daily_cap");

CREATE INDEX ON "users" ("daily_quantity");

CREATE INDEX ON "users" ("password_change_at");

CREATE INDEX ON "tasks" ("name");

CREATE INDEX ON "tasks" ("owner");

CREATE INDEX ON "tasks" ("deleted");

CREATE INDEX ON "tasks" ("content");

CREATE INDEX ON "tasks" ("deleted_at");

CREATE INDEX ON "tasks" ("content_change_at");

COMMENT ON COLUMN "users"."daily_cap" IS 'non negative';

COMMENT ON COLUMN "users"."daily_quantity" IS 'non negative';

COMMENT ON COLUMN "tasks"."name" IS 'unique per owner';
