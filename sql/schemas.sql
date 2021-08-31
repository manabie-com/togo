
-- users definition
CREATE TABLE users (
	id VARCHAR(64) NOT NULL,
	password VARCHAR(64) NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);
CREATE INDEX idx_users_id ON users(id);

-- tasks definition
CREATE TABLE tasks (
	id VARCHAR(36) NOT NULL,
	content TEXT NOT NULL,
	user_id VARCHAR(64) NOT NULL,
    created_date VARCHAR(10) NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE INDEX idx_tasks_id ON tasks(id);

-- add some users
INSERT INTO users (id, password) VALUES ('firstUser', '$2a$10$K11x2DC2DJzSBDIpVlQDduS2Gxmc/KpapuMGLPMpaH08d1s140Fs6');
INSERT INTO users (id, password) VALUES ('testUser', '$2a$10$fYyWHHPGc3XhZjVYPiS7Y.f8LSfcJB3PgIiH9GuuSEOXubJ1y34Su');
INSERT INTO users (id, password) VALUES ('itestUser', '$2a$10$5flCeXVX/SzXiNs5ZenBFORHRDQ9YOg2fPJvhpZBso/.PHtpHQxzO');

-- add some tasks
INSERT INTO tasks (id, content, user_id, created_date) VALUES ('5edd0c84-5b22-4076-a243-10c8fc13d84c', 'some tasks', 'testUser', '2021-08-29');
INSERT INTO tasks (id, content, user_id, created_date) VALUES ('000c824f-5beb-4738-969b-e8b00c9a67d7', 'some integration testing tasks', 'itestUser', '2021-08-30');