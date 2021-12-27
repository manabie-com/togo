CREATE TABLE
    "users" (
        "username" VARCHAR PRIMARY KEY,
        "hashed_password" VARCHAR NOT NULL,
        "full_name" VARCHAR NOT NULL,
        "email" VARCHAR UNIQUE NOT NULL,
        "cap" BIGINT NOT NULL,
        "password_change_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

CREATE TABLE
    "tasks" (
        "id" bigserial PRIMARY KEY,
        "owner" VARCHAR NOT NULL,
        "content" VARCHAR NOT NULL,
        "quantity" BIGINT NOT NULL,
        "created_at" timestamptz NOT NULL DEFAULT (now())
    );

ALTER TABLE
    "tasks"
ADD
    FOREIGN KEY ("owner") REFERENCES "users" ("username");

CREATE INDEX
ON "users" ("hashed_password");

CREATE INDEX
ON "users" ("full_name");

CREATE INDEX
ON "users" ("cap");

CREATE INDEX
ON "users" ("password_change_at");

CREATE INDEX
ON "tasks" ("owner");

CREATE INDEX
ON "tasks" ("quantity");

COMMENT
ON COLUMN "users"."cap" IS 'non negative';

COMMENT
ON COLUMN "tasks"."quantity" IS 'must be positive';

-- Generate admin
INSERT INTO
    users (
        username,
        hashed_password,
        full_name,
        email,
        cap,
        password_change_at,
        created_at
    )
VALUES (
        'admin',
        '$2a$10$VxkKRxRSov1e2LzNXc1aden5kkDAJEM5RF5n60HauC/zLpFhx/jfe',
        'Admin',
        'admin@email.com',
        '10',
        '0001-01-01 07:00:00.000',
        '2021-12-26 22:22:49.644'
    );
