CREATE TABLE users
(
    id            SERIAL        PRIMARY KEY,
    username      TEXT          UNIQUE NOT NULL,
    password      TEXT          NOT NULL,
    plan          TEXT          NOT NULL DEFAULT 'free',
    max_todo      INTEGER       NOT NULL DEFAULT 10
);

-- password: 'example'
-- INSERT INTO users (username, password, plan, max_todo) VALUES('admin', '$2a$10$hA5N/hUvta0rhYi4/xBXP.Oi2laKCdOSaTfWm.6pBTmvq3D1CtvWO', 'free', 10); 

-- DROP TABLE IF EXISTS users CASCADE;