### HOW TO RUN:
- To make it run: `docker-compose up` and waiting docker containers setup finish
- Login with method POST : `username=firstUser` and `password=example`
- User Test `togo.postman_collection.json` from `docs` folder to test
- Run docker-compose before test end-to-end TOGO application
### WHAT'S DONE:
- Successful to add new logic users are limited to create only 5 task only per day, 
if the daily limit is reached, return 4xx code to client and ignore the create request.
- Make unit test for `services` layer
- Write unit test for `storages` layer
- Split `services` layer to `use case` and `transport` layer
- Change from using `SQLite` to `Postgres` with `docker-compose`
- Make integration tests 
- Make this code DRY
- Refactor, optimize code
- Change login from method `GET` to `POST`,login from `user_id` and `password` to `username` and `password`
- Fix database structure: 
	+ Create new column `username` in `user` table, 
	+ Change `id, user_id` to `bigint`,  
	+ Change `created_date` to `timestamp`
- Add graceful shutdown
### WHAT ELSE WANT TO IMPROVE IF HAVE ENOUGH TIME
- Add more test case for unit test.
- Make integration test better.
- Add more validate.
- Make code clean, structure better.
- Add logger, monitoring, concurrency


DB SCHEMA OPTIMIZE; 

DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS tasks CASCADE;

CREATE TABLE IF NOT EXISTS users
(
    id       BIGINT             NOT NULL,
    username VARCHAR(50)        NOT NULL,
    password TEXT               NOT NULL,
    max_todo BIGINT DEFAULT 5   NOT NULL,
    CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, username, password, max_todo) VALUES ('100', 'firstUser', 'example', 5);
INSERT INTO users (id, username, password, max_todo) VALUES ('200', 'secondUser', 'example', 5);

CREATE TABLE IF NOT EXISTS tasks
(
    id           BIGINT     NOT NULL,
    content      TEXT       NOT NULL,
    user_id      BIGINT     NOT NULL,
    created_date timestamp  NOT NULL,
    CONSTRAINT tasks_PK PRIMARY KEY (id),
    CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users (id)
);