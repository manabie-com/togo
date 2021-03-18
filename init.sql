CREATE TABLE users (
    id TEXT PRIMARY KEY,
    password TEXT NOT NULL,
    max_todo INTEGER DEFAULT 5 NOT NULL
);

CREATE TABLE tasks (
    id TEXT NOT NULL PRIMARY KEY,
    content TEXT NOT NULL,
    user_id TEXT NOT NULL,
    created_date TEXT NOT NULL,
    CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);