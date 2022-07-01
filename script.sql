CREATE TABLE users
(
    id serial NOT NULL,
    email VARCHAR(40) NOT NULL,
    name VARCHAR(30) NOT NULL,
    password VARCHAR(20) NOT NULL,
    is_payment boolean DEFAULT false,
    limit_day_tasks smallint DEFAULT 10,
    is_active boolean DEFAULT true,
    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT email UNIQUE (email)
);


CREATE TABLE tasks
(
    id serial NOT NULL,
    name VARCHAR(100) NOT NULL,
    content text,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    user_id integer NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id)
        REFERENCES users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

