CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE users (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES(uuid_generate_v4(), 'example', 5);

CREATE TABLE tasks (
	id uuid DEFAULT uuid_generate_v4() NOT NULL,
	content TEXT NOT NULL,
	user_id uuid NOT NULL,
    created_date TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);