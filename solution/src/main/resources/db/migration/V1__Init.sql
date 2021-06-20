CREATE TABLE users (
                       id BIGSERIAL primary key,
                       username varchar(100) NOT NULL UNIQUE ,
                       password varchar(60) NOT NULL,
                       max_todo INTEGER DEFAULT 5 NOT NULL
);

-- tasks definition

CREATE TABLE tasks (
                       id BIGSERIAL primary key,
                       content varchar(1000) NOT NULL,
                       user_id bigint NOT NULL,
                       created_date timestamp NOT NULL,
                       CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);