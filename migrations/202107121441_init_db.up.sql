CREATE TABLE users (
	id varchar(36) NOT NULL,
	password varchar(20) NOT NULL,
	max_todo int DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

-- tasks definition

CREATE TABLE tasks (
	id varchar(36) NOT NULL,
	content TEXT NOT NULL,
	user_id varchar(36) NOT NULL,
    created_date varchar(20) NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);

-- create pair index
CREATE INDEX tasks_ucd
ON tasks(user_id,created_date);

