DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS tasks CASCADE;

CREATE TABLE IF NOT EXISTS users
(
    id       BIGINT             NOT NULL,
    username VARCHAR(50)        NOT NULL,
    password TEXT               NOT NULL,
    max_todo BIGINT DEFAULT 5   NOT NULL,
    CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, username, password, max_todo) VALUES ('100', 'firstUser', 'example', 5);
INSERT INTO users (id, username, password, max_todo) VALUES ('200', 'secondUser', 'example', 5);

CREATE TABLE IF NOT EXISTS tasks
(
    id           BIGINT     NOT NULL,
    content      TEXT       NOT NULL,
    user_id      BIGINT     NOT NULL,
    created_date timestamp  NOT NULL,
    CONSTRAINT tasks_PK PRIMARY KEY (id),
    CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users (id)
);