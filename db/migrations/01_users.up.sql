CREATE TABLE IF NOT EXISTS manabie.users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,

	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO manabie.users (id, password, max_todo) VALUES('firstUser', 'example', 5);
