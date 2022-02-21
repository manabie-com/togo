-- Create users table ----------
CREATE TABLE users
(
    id          uuid PRIMARY KEY,
    email       text        NOT NULL,
    name        text        NOT NULL,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    deleted_at timestamptz
);
-- ADD some records ----------
INSERT INTO users(id, email, name, created_at, updated_at)
VALUES ('0a1a88c2-18f9-42dd-b4b0-99ee4dc77751', 'chipv.bka@gmail.com', 'Chi Pham', '2022-02-02 01:42:22.103187+00',
        '2022-02-02 01:42:29.832697+00');
INSERT INTO users(id, email, name, created_at, updated_at)
VALUES ('dc334e08-3842-4dac-9338-8f30ac5e2369', 'chipv.@gmail.com', 'Chi Handsome', '2022-02-02 01:42:22.103187+00',
        '2022-02-02 01:42:29.832697+00');

-- Create limitations table ----------
CREATE TABLE limitations
(
    id          serial PRIMARY KEY,
    user_id     uuid        NOT NULL,
    limit_tasks bigint      NOT NULL DEFAULT 1,
    created_at timestamptz NOT NULL,
    updated_at timestamptz NOT NULL,
    deleted_at timestamptz
);
ALTER TABLE limitations
    ADD CONSTRAINT limitations_fk_user FOREIGN KEY (user_id) REFERENCES users (id);
CREATE INDEX limitations_user_id_idx ON limitations (user_id);

-- ADD some records on limitations_user_id_idx----------
INSERT INTO limitations(user_id, limit_tasks, created_at, updated_at)
VALUES ('0a1a88c2-18f9-42dd-b4b0-99ee4dc77751', 5, '2022-02-02 01:42:22.103187+00',
        '2022-02-02 01:42:29.832697+00');
INSERT INTO limitations(user_id, limit_tasks, created_at, updated_at)
VALUES ('dc334e08-3842-4dac-9338-8f30ac5e2369', 10, '2022-02-02 01:42:22.103187+00',
        '2022-02-02 01:42:29.832697+00');

-- Create Enum for the task ----------
CREATE TYPE task_priorities AS ENUM ('URGENT', 'HIGH', 'MEDIUM', 'NORMAL');
CREATE TYPE statuses AS ENUM ('TODO', 'INPROGRESS', 'COMPLETED');

-- Create user_tasks table ----------
CREATE TABLE tasks
(
    id          uuid PRIMARY KEY,
    title       text            NOT NULL,
    status      statuses        NOT NULL,
    priority    task_priorities NOT NULL,
    user_id     uuid            NOT NULL,
    created_at timestamptz     NOT NULL,
    updated_at timestamptz     NOT NULL,
    deleted_at timestamptz
);
INSERT INTO tasks(id, title, status, priority, user_id, created_at, updated_at)
VALUES ('dc334e08-3842-4dac-9338-8f30ac5e2369', 'task 1', 'TODO', 'URGENT', 'dc334e08-3842-4dac-9338-8f30ac5e2369', '2022-02-02 01:42:22.103187+00',
        '2022-02-02 01:42:29.832697+00');

CREATE INDEX user_tasks_user_id_idx ON tasks (user_id);