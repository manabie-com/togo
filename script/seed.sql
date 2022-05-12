CREATE TABLE IF NOT EXISTS "users"
(
    id          INTEGER PRIMARY KEY AUTOINCREMENT,
    username    TEXT    NOT NULL UNIQUE,
    daily_limit INTEGER NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "tasks"
(
    id         INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id    INTEGER,
    content    TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

INSERT INTO users (username, daily_limit)
VALUES ('user1', 1)
ON CONFLICT DO NOTHING;
INSERT INTO users (username, daily_limit)
VALUES ('user2', 2)
ON CONFLICT DO NOTHING;
INSERT INTO users (username, daily_limit)
VALUES ('admin', 3)
ON CONFLICT DO NOTHING;