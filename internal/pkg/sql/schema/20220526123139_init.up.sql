CREATE TABLE "users"
(
    "username"        varchar(50) PRIMARY KEY,
    "full_name"       varchar(50)        NOT NULL,
    "hashed_password" varchar(100)        NOT NULL,
    "email"           varchar(50) UNIQUE NOT NULL,
    "created_at"      timestamptz        NOT NULL DEFAULT (now()),
    "tier_id"         int                NOT NULL
);

CREATE TABLE "tasks"
(
    "id"          serial PRIMARY KEY,
    "name"        varchar     NOT NULL,
    "assignee"    varchar,
    "assign_date" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "description" varchar,
    "status"      varchar(50) NOT NULL,
    "creator"     varchar(50) NOT NULL,
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    "start_date"  timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
    "end_date"    timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "tiers"
(
    "id"          SMALLSERIAL PRIMARY KEY,
    "name"        varchar(50) NOT NULL,
    "description" varchar     NOT NULL
);

CREATE TABLE "limits"
(
    "tier_id" SMALLSERIAL,
    "action"  varchar NOT NULL,
    "value"   int     NOT NULL,
    PRIMARY KEY ("tier_id", "action")
);

CREATE INDEX ON "tasks" ("assignee");

ALTER TABLE "users"
    ADD FOREIGN KEY ("tier_id") REFERENCES "tiers" ("id");

ALTER TABLE "tasks"
    ADD FOREIGN KEY ("assignee") REFERENCES "users" ("username");

ALTER TABLE "tasks"
    ADD FOREIGN KEY ("creator") REFERENCES "users" ("username");

ALTER TABLE "limits"
    ADD FOREIGN KEY ("tier_id") REFERENCES "tiers" ("id");




-- TODO: this is a mock db for demo, remove later

INSERT INTO tiers ("name", description)
VALUES ('free', 'asd');
INSERT INTO tiers ("name", description)
VALUES ('standard', 'asd');


INSERT INTO limits (tier_id, "action", value)
VALUES (1, 'receive:task', 5);
INSERT INTO limits (tier_id, "action", value)
VALUES (2, 'receive:task', 10);


INSERT INTO users (username, full_name, hashed_password, email, created_at, tier_id)
VALUES ('free1', 'free1', '$2a$10$e9OfC7ft.heOXdOsZtQKC.nJY1CAwdJ1dZdpR0How0f88WjLu2nom', 'free1@gmail.com', now(), 1);
INSERT INTO users (username, full_name, hashed_password, email, created_at, tier_id)
VALUES ('standard1', 'standard1', '$2a$10$e9OfC7ft.heOXdOsZtQKC.nJY1CAwdJ1dZdpR0How0f88WjLu2nom', 'standard1@gmail.com', now(), 2);


INSERT INTO tasks ("name", assignee, assign_date, description, status, creator, created_at, start_date, end_date)
VALUES ('task 1', 'free1', now(), 'api create user', 'todo', 'free1', now(),
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone,
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone);
INSERT INTO tasks ("name", assignee, assign_date, description, status, creator, created_at, start_date, end_date)
VALUES ('task 2', 'free1', now(), 'api update user', 'todo', 'free1', now(),
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone,
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone);
INSERT INTO tasks ("name", assignee, assign_date, description, status, creator, created_at, start_date, end_date)
VALUES ('task 3', 'free1', now(), 'api delete user', 'todo', 'free1', now(),
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone,
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone);
INSERT INTO tasks ("name", assignee, assign_date, description, status, creator, created_at, start_date, end_date)
VALUES ('task 4', 'free1', now(), 'api list user', 'todo', 'free1', now(),
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone,
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone);
INSERT INTO tasks ("name", assignee, assign_date, description, status, creator, created_at, start_date, end_date)
VALUES ('task 5', 'free1', now(), 'api list user', 'todo', 'free1', now(),
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone,
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone);
INSERT INTO tasks ("name", assignee, assign_date, description, status, creator, created_at, start_date, end_date)
VALUES ('task 6', NULL, now(), 'api list user', 'todo', 'free1', now(),
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone,
        '0001-01-01 00:00:00+00:00:00'::timestamp with time zone);