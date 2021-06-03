CREATE TABLE IF NOT EXISTS manabie.tasks (
    id serial PRIMARY KEY,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    create_date TEXT NOT NULL,
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES manabie.users(id)
);
