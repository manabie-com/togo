
-- users definition
SELECT '
    CREATE TABLE users (
        id TEXT NOT NULL,
        password TEXT NOT NULL,
        max_todo INTEGER NOT NULL,
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

CREATE INDEX ON "users" ("id", "password");
CREATE INDEX ON "tasks" ("created_date", "user_id");
CREATE INDEX ON "tasks" ("user_id");

INSERT INTO users (id, password, max_todo) VALUES
('firstUser', '50d858e0985ecc7f60418aaf0cc5ab587f42c2570a884095a9e8ccacd0f6545c', 5),
('secondUser', '50d858e0985ecc7f60418aaf0cc5ab587f42c2570a884095a9e8ccacd0f6545c', 10),
('thirdUser', '50d858e0985ecc7f60418aaf0cc5ab587f42c2570a884095a9e8ccacd0f6545c', 7),
('fourthUser', '50d858e0985ecc7f60418aaf0cc5ab587f42c2570a884095a9e8ccacd0f6545c', 1);

INSERT INTO tasks (id, content, user_id, created_date) VALUES
('fa258101-173d-412b-b39f-074cfd74710b', 'content', 'fourthUser', (SELECT CURRENT_DATE));