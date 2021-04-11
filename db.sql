CREATE TABLE users (
    id        TEXT PRIMARY KEY,
    password  TEXT NOT NULL,
    max_todo  INTEGER DEFAULT 5 NOT NULL
);

INSERT INTO users (id, password, max_todo) VALUES ('firstUser', 'example', 5);
INSERT INTO users (id, password, max_todo) VALUES ('spamUser', 'example', 0);

CREATE TABLE tasks (
    id            TEXT PRIMARY KEY,
    content       TEXT NOT NULL,
    user_id       TEXT NOT NULL REFERENCES users (id),
    created_date  TEXT NOT NULL
);
