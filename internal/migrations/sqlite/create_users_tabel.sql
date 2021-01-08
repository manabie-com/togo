-- users definition

CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', '$2a$10$Px50y37hZA.W4h8t2hvDMeIyenU3kDNWx0NCZpBtUyHHJUbW3e1Uu', 5);