#!/bin/bash
set -e
export PGPASSWORD=$POSTGRES_PASSWORD;
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    BEGIN;
        CREATE EXTENSION pgcrypto;
        CREATE TABLE IF NOT EXISTS users (
            id TEXT NOT NULL,
            password TEXT NOT NULL,
            CONSTRAINT users_PK PRIMARY KEY (id)
        );
        CREATE TABLE tasks (
            id TEXT NOT NULL,
            content TEXT NOT NULL,
            user_id TEXT NOT NULL,
            status TEXT NOT NULL,
            created_date TEXT NOT NULL,
            target_date TEXT NOT NULL,
            CONSTRAINT tasks_PK PRIMARY KEY (id),
            CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
        );
    COMMIT;
EOSQL