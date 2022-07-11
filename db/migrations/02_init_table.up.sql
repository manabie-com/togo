CREATE TABLE IF NOT EXISTS togo.users (
    id serial PRIMARY KEY,
    email text NULL,
    password text NOT NULL,
    name text NOT NULL,
    quota int NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS togo.tasks (
    id serial PRIMARY KEY,
    title text NOT NULL,
    description text,
    priority int NOT NULL DEFAULT 0,
    is_done boolean NOT NULL DEFAULT FALSE,
    user_id int NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP,

    CONSTRAINT task_user_fk FOREIGN KEY (user_id) REFERENCES togo.users(id)
);
