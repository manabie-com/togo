CREATE TABLE users
(
    id            SERIAL  PRIMARY KEY,
    name          TEXT    NOT NULL,
    password_hash TEXT    NOT NULL,
    plan          TEXT    NOT NULL DEFAULT 'free',
    max_todo      INTEGER NOT NULL DEFAULT 10
);

-- password: 'example'
-- INSERT INTO users (id, name, password_hash, plan, max_todo) VALUES(1, 'admin', '$2a$10$hA5N/hUvta0rhYi4/xBXP.Oi2laKCdOSaTfWm.6pBTmvq3D1CtvWO', 'free', 10); 