CREATE TABLE users
(
    id text NOT NULL,
    password text NOT NULL,
    max_todo integer NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

CREATE TABLE tasks
(
    id text NOT NULL,
    content text NOT NULL,
    user_id text NOT NULL,
    created_date text NOT NULL,
    CONSTRAINT "tasks_PK" PRIMARY KEY (id),
    CONSTRAINT "tasks_FK" FOREIGN KEY (user_id)
        REFERENCES public.users (id) MATCH SIMPLE
);

INSERT INTO tasks (id, content, user_id, created_date) VALUES ('e1da0b9b-7ecc-44f9-82ff-4623cc50446a', 'first content', 'firstUser', '2020-06-29');
INSERT INTO tasks (id, content, user_id, created_date) VALUES ('055261ab-8ba8-49e1-a9e8-e9f725ba9104', 'second content', 'firstUser', '2020-06-29');
INSERT INTO tasks (id, content, user_id, created_date) VALUES ('2bf3d510-c0fb-41e9-ad12-4b9a60b37e7a', 'another content', 'firstUser', '2020-06-29');