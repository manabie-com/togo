CREATE TABLE IF NOT EXISTS tasks (
    id TEXT NOT NULL,
    content TEXT NOT NULL,
    user_id TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT tasks_pk PRIMARY KEY (id),
	CONSTRAINT tasks_fk FOREIGN KEY (user_id) REFERENCES users (id)
);