-- Version: 1.1
-- Description: Create table users
CREATE TABLE users (
	id             UUID                     NOT NULL,
	name           TEXT                     NOT NULL,
	email          TEXT UNIQUE              NOT NULL,
	password_hash  TEXT                     NOT NULL,
	daily_max_todo INT                      NOT NULL,
	date_created   TIMESTAMP WITH TIME ZONE NOT NULL,
	date_updated   TIMESTAMP WITH TIME ZONE NOT NULL,

	PRIMARY KEY (id)
);

-- Version: 1.2
-- Description: Create table todos
CREATE TABLE todos (
	id           UUID                     NOT NULL,
	title        TEXT                     NOT NULL,
	content      TEXT                     NULL,
	user_id      UUID                     NOT NULL,
	date_created TIMESTAMP WITH TIME ZONE NOT NULL,
	date_updated TIMESTAMP WITH TIME ZONE NOT NULL,

	PRIMARY KEY (id),
	FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

