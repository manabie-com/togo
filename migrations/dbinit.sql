CREATE EXTENSION IF NOT EXISTS "pgcrypto";
DROP TABLE IF EXISTS tasks;
DROP TABLE IF EXISTS users;

-- users definition
CREATE TABLE users (
	id VARCHAR(20) NOT NULL,
	password VARCHAR(255) NOT NULL,
	max_todo SMALLINT DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

-- tasks definition
CREATE TABLE tasks (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	content TEXT NOT NULL,
	user_id VARCHAR(20) NOT NULL,
    created_date DATE NOT NULL DEFAULT Now(),
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
