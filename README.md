### Notes
This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make it run:
- `docker-compose up -d --build`
- Import Postman collection from docs to check example

#### DB Schema
```sql
CREATE DATABASE todos;
-- users definition

CREATE TABLE users (
    id serial NOT NULL primary key,
    email varchar(255) NOT NULL,
    username  varchar(255) NOT NULL,
    password text NOT NULL,
    max_todo int DEFAULT 0 NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    CONSTRAINT users_Uniquel UNIQUE (email, username)
);

INSERT INTO users (email, username, password, max_todo, created_at, updated_at) VALUES('me@here.com', 'Todo User', '$2a$12$orZppdmhH.KRrxcZcjx0NeLPtIDpaf2GNUben4Rz7w53e5dSQJgdq', 0, '2021-05-17 00:00:00', '2021-05-17 00:00:00')


-- tasks definition

CREATE TABLE tasks (
    id serial NOT NULL primary key,
    content varchar(255),
    user_id int NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);

INSERT INTO tasks (content, user_id, created_at, updated_at) VALUES('Sample Content Todo', 1,  '2021-05-17 00:00:00', '2021-05-17 00:00:00')
```
### NEEDS TO IMPROVE
- Using other libraries for GO that will help a project structure like an MVC pattern
- Error handling and JSON Response should have a common/centralize function
- Security for all the routes that needs json token should be common/centralize

### ADDITIONAL THAT I'VE MADE
- Update And Delete task/todo


