SELECT 'CREATE DATABASE togo_test'
WHERE NOT EXISTS(SELECT FROM pg_database WHERE datname = 'togo_test')
\gexec

CREATE TABLE IF NOT EXISTS users
(
    id       TEXT              NOT NULL,
    password TEXT              NOT NULL,
    max_todo INTEGER DEFAULT 5 NOT NULL,
    CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo)
VALUES ('firstUser', 'example', 5),
       ('secondUser', 'example', 5);

CREATE TABLE IF NOT EXISTS tasks
(
    id           TEXT NOT NULL,
    content      TEXT NOT NULL,
    user_id      TEXT NOT NULL,
    created_date TEXT NOT NULL,
    CONSTRAINT tasks_PK PRIMARY KEY (id),
    CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users (id)
);

INSERT INTO tasks
VALUES ('e1da0b9b-7ecc-44f9-82ff-4623cc50446a', 'first content', 'firstUser', '2020-06-29'),
       ('055261ab-8ba8-49e1-a9e8-e9f725ba9104', 'second content', 'firstUser', '2020-06-29'),
       ('2bf3d510-c0fb-41e9-ad12-4b9a60b37e7a', 'another content', 'firstUser', '2020-06-29'),
       ('e35e13f8-35f3-409f-8e2f-f3e0173fcca3', 'sadsa', 'firstUser', '2020-08-10'),
       ('2a73a4d5-dd05-4c77-bcbd-f5e51a6d6809', 'sadsad', 'firstUser', '2020-08-11')