-- users definition

CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

--- pwd: 'example'
INSERT INTO users (id, password, max_todo) VALUES('firstUser', '$2a$10$hA5N/hUvta0rhYi4/xBXP.Oi2laKCdOSaTfWm.6pBTmvq3D1CtvWO', 5);