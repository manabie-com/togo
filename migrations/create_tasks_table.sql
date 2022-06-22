
CREATE TABLE tasks
(
    id           TEXT PRIMARY KEY,
    name         TEXT                       NOT NULL,
    description  TEXT                       NOT NULL,
    created_date TEXT                       NOT NULL,
    checked      BOOLEAN                    NOT NULL,
    user_id      TEXT REFERENCES users (id) NOT NULL
);

-- DROP TABLE IF EXISTS tasks CASCADE;
