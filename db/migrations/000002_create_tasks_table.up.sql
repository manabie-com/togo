CREATE TABLE tasks (
	id VARCHAR(50) NOT NULL,
	content TEXT NOT NULL,
	user_id bigint NOT NULL,
    created_date DATE NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);