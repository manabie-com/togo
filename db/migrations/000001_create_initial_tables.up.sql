CREATE TABLE users (
	id SERIAL PRIMARY KEY,
  username TEXT NOT NULL,
	password TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE tasks (
	id SERIAL PRIMARY KEY,
	content TEXT NOT NULL,
	user_id INTEGER,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	CONSTRAINT tasks_fk_users FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE INDEX tasks_created_at_idx ON tasks(created_at)