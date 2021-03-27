GRANT ALL PRIVILEGES
ON DATABASE manabie_db TO demo;
CREATE EXTENSION pgcrypto;

CREATE SEQUENCE togo_users_id_seq START 1;
CREATE SEQUENCE togo_task_id_seq START 1;

CREATE TABLE users
(
    id               int8    NOT NULL DEFAULT nextval('togo_users_id_seq'::regclass),
    username         varchar(36) NOT NULL,
    password         varchar(255) NOT NULL,
    max_todo int4             DEFAULT 5 NOT NULL,
    CONSTRAINT users_PK PRIMARY KEY (id)
);
ALTER SEQUENCE togo_users_id_seq OWNED BY users.id;
CREATE UNIQUE INDEX uniq_username ON users (username);

INSERT INTO users("username", "password") VALUES('firstUser', crypt('example', gen_salt('bf', 8)));

-- tasks definition
CREATE TABLE tasks
(
    id           int8      NOT NULL DEFAULT nextval('togo_task_id_seq'::regclass),
    content      TEXT      NOT NULL,
    user_id      int8      NOT NULL,
    created_date timestamp NOT NULL,
    CONSTRAINT tasks_PK PRIMARY KEY (id),
    CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users (id)
);
ALTER SEQUENCE togo_task_id_seq OWNED BY tasks.id;