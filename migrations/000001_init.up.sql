-- user table
CREATE TABLE users
(
    username        VARCHAR PRIMARY KEY,
    hashed_password VARCHAR   NOT NULL,
    max_todo        INTEGER            DEFAULT 5 NOT NULL,
    created_at      TIMESTAMP NOT NULL DEFAULT (NOW())
    updated_at      TIMESTAMP NOT NULL DEFAULT (NOW())
);

-- seed one user, hashed_password = "example"
INSERT INTO users (username, hashed_password, max_todo)
VALUES ('firstUser', '$2a$10$3jEtynoYdZJlw2fTUMjuCeGxHEjvc8a23gXMaidDW3yKPjMWFbb4W', 5);

-- tasks definition table

CREATE TABLE tasks
(
    id         SERIAL PRIMARY KEY,
    content    TEXT      NOT NULL,
    username   VARCHAR   NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT (NOW()),
    updated_at TIMESTAMP NOT NULL DEFAULT (NOW()),
    CONSTRAINT tasks_FK FOREIGN KEY (username) REFERENCES users (username)
);