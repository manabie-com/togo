CREATE DATABASE togo;
CREATE TABLE IF NOT EXISTS users
(
    id       VARCHAR(100) NOT NULL,
    password VARCHAR(200) NOT NULL,
    max_todo INT          NOT NULL Default 5,
    PRIMARY KEY (id)
    );

CREATE TABLE IF NOT EXISTS tasks
(
    id           VARCHAR(100) NOT NULL,
    content      TEXT         NOT NULL,
    user_id      VARCHAR(100) NOT NULL,
    created_date DATE         NOT NULL DEFAULT CURRENT_DATE,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES users (id)
    );

INSERT INTO users (id, password, max_todo)
VALUES ('firstUser', '$2a$14$BdgOuNVBU7sdGW9rIDIIv.MWXDdvTVKyTppb3qW03bmvz/6hhA1FO', 5);

INSERT INTO tasks (id, content, user_id, created_date)
VALUES ('e1da0b9b-7ecc-44f9-82ff-4623cc50446a', 'first content', 'firstUser', '2020-06-29'),
       ('055261ab-8ba8-49e1-a9e8-e9f725ba9104', 'second content', 'firstUser', '2020-06-29');

