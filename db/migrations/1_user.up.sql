BEGIN;

CREATE TABLE IF NOT EXISTS users (
    id varchar(50) NOT NULL,
    username varchar(50) NOT NULL,
    password varchar(50) NOT NULL,
    max_task_per_day INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT user_PK PRIMARY KEY (id)
);

COMMIT;
