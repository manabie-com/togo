CREATE TABLE users (
    id serial PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (256) NOT NULL,
    max_task_per_day INTEGER DEFAULT 5 NOT NULL
);