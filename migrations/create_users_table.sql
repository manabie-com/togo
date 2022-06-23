CREATE TABLE users
(
    id            SERIAL        PRIMARY KEY,
    username      VARCHAR(200)  UNIQUE NOT NULL,
    password      VARCHAR(200)  NOT NULL,
    plan          VARCHAR(50)   NOT NULL DEFAULT 'free',
    max_todo      INTEGER       NOT NULL DEFAULT 10
);

-- password: 'admin'
-- INSERT INTO users (username, password, plan, max_todo) VALUES('admin', '$2a$10$b1cUVK/l7O0D1q4TU1IM7O/sUq7uXmZU.uLiSgQJoD2jFbPmzbK2a', 'free', 10); 

-- DROP TABLE IF EXISTS users CASCADE;