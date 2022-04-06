CREATE TABLE tasks (
    id VARCHAR(50) PRIMARY KEY,
    detail TEXT NOT NULL,
    user_id bigint NOT NULL,
    created_date DATE NOT NULL,
    CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
