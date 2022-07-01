CREATE TABLE users
(
    id            SERIAL        PRIMARY KEY,
    username      VARCHAR(200)  UNIQUE NOT NULL,
    password      VARCHAR(200)  NOT NULL,
    plan          VARCHAR(50)   NOT NULL DEFAULT 'free',
    max_todo      INTEGER       NOT NULL DEFAULT 10
);

INSERT INTO users (username, password, plan, max_todo) VALUES('admin', '$2a$10$b1cUVK/l7O0D1q4TU1IM7O/sUq7uXmZU.uLiSgQJoD2jFbPmzbK2a', 'free', 10); -- password: admin
INSERT INTO users (username, password, plan, max_todo) VALUES('admin1', '$2a$10$b1cUVK/l7O0D1q4TU1IM7O/sUq7uXmZU.uLiSgQJoD2jFbPmzbK2a', 'free', 10); -- password: admin
INSERT INTO users (username, password, plan, max_todo) VALUES('free', '$2a$10$sn4/wbXxUodhTDHviykz8OgD0X.xugS/BX2D7J6n5A9OLgfHCsWmC', 'free', 10); -- password: password
INSERT INTO users (username, password, plan, max_todo) VALUES('vip', '$2a$10$sn4/wbXxUodhTDHviykz8OgD0X.xugS/BX2D7J6n5A9OLgfHCsWmC', 'vip', 20); -- password: password


CREATE TABLE tasks
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(200)                      NOT NULL,
    description  VARCHAR(500)                      NOT NULL,
    created_at   DATE                              NOT NULL,
    completed    BOOLEAN                           NOT NULL,
    user_id      INTEGER REFERENCES users (id)     NOT NULL
);

INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('to do 1', 'test for to do list 1', CURRENT_DATE, 'false', 1); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('to do 2', 'test for to do list 2', CURRENT_DATE, 'false', 1); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('to do 3', 'test for to do list 3', CURRENT_DATE, 'false', 1); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('free test', 'test for free user', CURRENT_DATE, 'false', 3); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('free test', 'test for free user', CURRENT_DATE, 'false', 3); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('free test', 'test for free user', CURRENT_DATE, 'false', 3); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('free test', 'test for free user', CURRENT_DATE, 'false', 3); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('free test', 'test for free user', CURRENT_DATE, 'false', 3); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('free test', 'test for free user', CURRENT_DATE, 'false', 3); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('free test', 'test for free user', CURRENT_DATE, 'false', 3); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('free test', 'test for free user', CURRENT_DATE, 'false', 3); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('free test', 'test for free user', CURRENT_DATE, 'false', 3); 
INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('free test', 'test for free user', CURRENT_DATE, 'false', 3); 