DROP DATABASE IF EXISTS togo;
CREATE DATABASE togo;
\c togo;

DROP TABLE IF EXISTS users;
CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

-- tasks definition
DROP TABLE IF EXISTS tasks;
CREATE TABLE tasks (
	id SERIAL NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    created_date TIMESTAMP NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);