BEGIN;

CREATE TABLE IF NOT EXISTS tasks(
    id varchar(50) NOT NULL,
    content varchar(50) NOT NULL,
    create_date varchar(50) NOT NULL,
    user_id varchar(50) NOT NULL,
    CONSTRAINT task_PK PRIMARY KEY (id),
    CONSTRAINT task_FK FOREIGN KEY (user_id) REFERENCES users(id)
);

COMMIT;
