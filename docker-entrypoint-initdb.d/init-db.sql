--DROP TABLE IF EXISTS togo_app;
--
--CREATE TABLE togo_app;

-- users definition
SELECT '
    CREATE TABLE users (
        id TEXT NOT NULL,
        password TEXT NOT NULL,
        max_todo INTEGER DEFAULT 5 NOT NULL,
        CONSTRAINT users_PK PRIMARY KEY (id));
'
WHERE NOT EXISTS (SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_name='users')\gexec

-- tasks definition
SELECT '
    CREATE TABLE tasks (
        id TEXT NOT NULL,
        content TEXT NOT NULL,
        user_id TEXT NOT NULL,
        created_date TEXT NOT NULL,
        CONSTRAINT tasks_PK PRIMARY KEY (id),
        CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id));
'
WHERE NOT EXISTS (SELECT table_name FROM information_schema.tables WHERE table_schema='public' AND table_name='tasks')\gexec

INSERT INTO users (id, password, max_todo) VALUES
('firstUser', 'example', 5),
('secondUser', 'example', 5),
('thirdUser', 'example', 5);

INSERT INTO tasks (id, content, user_id, created_date) VALUES
('e1da0b9b-7ecc-44f9-82ff-4623cc50446a', 'first content', 'firstUser', '2020-06-29'),
('055261ab-8ba8-49e1-a9e8-e9f725ba9104', 'second content', 'firstUser', '2020-06-29'),
('2bf3d510-c0fb-41e9-ad12-4b9a60b37e7a', 'another content', 'firstUser', '2020-06-29'),
('8b5acb6a-7511-45c2-841f-1d955ea3f27e', 'this is a content', 'secondUser', (SELECT CURRENT_DATE)),
('a3e67d11-6b52-42fc-bb25-071e8f44fc78', 'this is a content', 'secondUser', (SELECT CURRENT_DATE)),
('af81bf90-d22c-4df4-a323-46835661e57f', 'this is a content', 'secondUser', (SELECT CURRENT_DATE)),
('dad2bc21-2241-44b5-9ecc-08ae783de59e', 'this is a content', 'secondUser', (SELECT CURRENT_DATE)),
('db5f18e3-c9a3-4bce-af05-6bcb6ce9efb4', 'this is a content', 'secondUser', (SELECT CURRENT_DATE));

-- grant permission
GRANT ALL PRIVILEGES ON DATABASE togo_app TO postgres;