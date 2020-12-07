CREATE TABLE task (
    id BIGINT NOT NULL,
    content TEXT NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP DEFAULT NULL,
    CONSTRAINT tasks_PK PRIMARY KEY (id)
);

CREATE TABLE users (
    id BIGINT NOT NULL,
    name TEXT NOT NULL,
    password TEXT NOT NULL,
    CONSTRAINT users_PK PRIMARY KEY (id)
);

CREATE INDEX task_user_id ON task(user_id);
CREATE INDEX users_name ON users(name);


INSERT INTO users (id, name, password) VALUES(123456789, 'firstUser', '$2a$10$FTe4r.Zp62XQNMvfEZ5b/.xH2D.juLU9Z.9GHvlnBud65pS9c7L86');
