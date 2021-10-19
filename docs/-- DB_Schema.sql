-- SQLite

DROP TABLE tasks;
DROP TABLE task_status;
DROP TABLE users;

-- users definition
CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);


-- task status definition
CREATE TABLE task_status (
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT NULL,
    CONSTRAINT task_status_PK PRIMARY KEY (code)
);

INSERT INTO task_status (code, name, description) VALUES('todo', 'To Do', 'Task need to be done');
INSERT INTO task_status (code, name, description) VALUES('doing', 'Doing', 'Task is doing');
INSERT INTO task_status (code, name, description) VALUES('done', 'Done', 'Task is already done, Yay!');
INSERT INTO task_status (code, name, description) VALUES('skip', 'Skipped', 'I dont want to do this task anymore');


-- tasks definition

CREATE TABLE tasks (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    status_code TEXT,
    due_date TEXT NULL,
    created_date TEXT NOT NULL,
    updated_at TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_users_FK FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT tasks_status_FK FOREIGN KEY (status_code) REFERENCES task_status(code)
);