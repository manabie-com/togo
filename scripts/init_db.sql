
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
('firstUser', 'example', 5),
('secondUser', 'example', 10),
('thirdUser', 'example', 7);