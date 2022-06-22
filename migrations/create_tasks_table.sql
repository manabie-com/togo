
CREATE TABLE tasks
(
    id           SERIAL PRIMARY KEY,
    name         TEXT                          NOT NULL,
    description  TEXT                          NOT NULL,
    created_at   TEXT                          NOT NULL,
    completed    BOOLEAN                       NOT NULL,
    user_id      INTEGER REFERENCES users (id) NOT NULL
);

-- DROP TABLE IF EXISTS tasks CASCADE;
