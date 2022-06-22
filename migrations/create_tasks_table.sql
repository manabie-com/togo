
CREATE TABLE tasks
(
    id           SERIAL PRIMARY KEY,
    name         TEXT                              NOT NULL,
    description  TEXT                              NOT NULL,
    created_at   TEXT                              NOT NULL,
    completed    BOOLEAN                           NOT NULL,
    username     TEXT REFERENCES users (username)  NOT NULL
);

-- DROP TABLE IF EXISTS tasks CASCADE;
