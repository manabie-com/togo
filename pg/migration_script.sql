DROP TABLE IF EXISTS tasks;

DROP TABLE IF EXISTS users;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
	id uuid DEFAULT uuid_generate_v4 (),
	user_name VARCHAR(32) NOT NULL,
	password VARCHAR(255) NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (user_name, password, max_todo) VALUES('firstUser', 'password', 5);
INSERT INTO users (user_name, password, max_todo) VALUES('secondUser', 'password', 1);

CREATE TABLE tasks (
	id uuid DEFAULT uuid_generate_v4 (),
	content TEXT NOT NULL,
	user_id uuid NOT NULL,
    created_date VARCHAR(10) NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX tasks_date ON tasks USING btree
(
    created_date
);