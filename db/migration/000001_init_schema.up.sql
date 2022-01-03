CREATE TABLE
    "users" (
        "username" VARCHAR PRIMARY KEY,
        "hashed_password" VARCHAR NOT NULL,
        "full_name" VARCHAR NOT NULL,
        "email" VARCHAR UNIQUE NOT NULL,
        "daily_cap" BIGINT NOT NULL,
        "daily_quantity" BIGINT NOT NULL,
        "password_change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

CREATE TABLE
    "tasks" (
        "id" bigserial PRIMARY KEY,
        "name" VARCHAR NOT NULL,
        "owner" VARCHAR NOT NULL,
        "content" VARCHAR NOT NULL,
        "content_change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

ALTER TABLE
    "tasks"
ADD
    FOREIGN KEY ("owner") REFERENCES "users" ("username");

CREATE INDEX
ON "users" ("full_name");

CREATE INDEX
ON "users" ("daily_cap");

CREATE INDEX
ON "users" ("daily_quantity");

CREATE INDEX
ON "users" ("password_change_at");

CREATE INDEX
ON "tasks" ("name");

CREATE INDEX
ON "tasks" ("owner");

CREATE INDEX
ON "tasks" ("content");

CREATE INDEX
ON "tasks" ("content_change_at");

COMMENT
ON COLUMN "users"."daily_cap" IS 'non negative';

COMMENT
ON COLUMN "users"."daily_quantity" IS 'non negative';

COMMENT
ON COLUMN "tasks"."name" IS 'unique per owner';

-- Generate admin
INSERT INTO
    users (
        username,
        hashed_password,
        full_name,
        email,
        daily_cap,
        daily_quantity,
        password_change_at,
        created_at
    )
VALUES (
        'admin',
        '$2a$10$VxkKRxRSov1e2LzNXc1aden5kkDAJEM5RF5n60HauC/zLpFhx/jfe',
        'Admin',
        'admin@email.com',
        '10',
        '0',
        '0001-01-01 07:00:00.000',
        '2021-12-26 22:22:49.644'
    );
