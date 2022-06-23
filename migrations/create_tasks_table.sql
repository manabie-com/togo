
CREATE TABLE tasks
(
    id           SERIAL PRIMARY KEY,
    name         VARCHAR(200)                      NOT NULL,
    description  VARCHAR(500)                      NOT NULL,
    created_at   DATE                              NOT NULL,
    completed    BOOLEAN                           NOT NULL,
    user_id      INTEGER REFERENCES users (id)     NOT NULL
);

-- DROP TABLE IF EXISTS tasks CASCADE;
-- INSERT INTO tasks (name, description, created_at, completed, user_id) VALUES('to do 1', 'test for to do list 1', '16/06/2011', 'false', 1); 