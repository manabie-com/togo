-- users definition

CREATE TABLE users (
                       id TEXT NOT NULL,
                       password TEXT NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_DATE,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_DATE,
                       deleted_at TIMESTAMP NULL,
                       CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password) VALUES('firstUser', 'example');

-- tasks definition

CREATE TABLE tasks (
                       id TEXT NOT NULL,
                       content TEXT NOT NULL,
                       user_id TEXT NOT NULL,
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_DATE,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_DATE,
                       deleted_at TIMESTAMP NULL,
                       CONSTRAINT tasks_PK PRIMARY KEY (id),
                       CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);