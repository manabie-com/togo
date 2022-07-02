-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    user_id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
    username character varying(200) NOT NULL UNIQUE,
    task_daily_limit int NOT NULL DEFAULT 5,
    created_by UUID NOT NULL,
    updated_by UUID,
    created_when timestamptz NOT NULL,
    updated_when timestamptz,
    is_active boolean NOT NULL DEFAULT true
);

CREATE TABLE tasks (
    task_id UUID PRIMARY KEY NOT NULL DEFAULT uuid_generate_v1(),
    user_id UUID REFERENCES users(user_id) NOT NULL,
    title character varying(200) NOT NULL,
    description text,
    created_by UUID NOT NULL,
    updated_by UUID,
    created_when timestamptz NOT NULL,
    updated_when timestamptz,
    is_active boolean NOT NULL DEFAULT true
);

CREATE INDEX users_username_is_active_idx ON users USING BTREE (username, is_active);

CREATE INDEX tasks_user_id_created_when_is_active_idx ON tasks USING BTREE (user_id, created_when, is_active);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS users_username_password_is_active_idx;

DROP INDEX IF EXISTS user_settings_user_id_is_active_idx;

DROP INDEX IF EXISTS tasks_user_id_is_active_created_when_idx;

DROP TABLE IF EXISTS tasks;

DROP TABLE IF EXISTS user_settings;

DROP TABLE IF EXISTS users;

DROP EXTENSION IF EXISTS "uuid-ossp";

-- +goose StatementEnd