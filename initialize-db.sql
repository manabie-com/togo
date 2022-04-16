CREATE DATABASE tododb WITH ENCODING 'UTF8' LC_COLLATE='English_United States' LC_CTYPE='English_United States';

create user todo_user with encrypted password 'secret';

grant all privileges on database tododb to todo_user;

GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public to todo_user;

CREATE TABLE IF NOT EXISTS users (
 username VARCHAR(15) PRIMARY KEY,
 name VARCHAR(100) NOT NULL,
 email VARCHAR(50) NOT NULL,
 password VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS tasks (
  id SERIAL PRIMARY KEY,        
  title VARCHAR(100) NOT NULL,
  description VARCHAR(200) NOT NULL,
  username VARCHAR(15) NOT NULL,
  create_date DATE default CURRENT_DATE,
  CONSTRAINT fk_username FOREIGN KEY(username) REFERENCES users(username)
);

CREATE TABLE IF NOT EXISTS task_limit_config (
   username VARCHAR(15) NOT NULL,
   task_limit INTEGER NOT NULL,
   CONSTRAINT fk_username_config FOREIGN KEY(username) REFERENCES users(username)
);

insert into users (username, name, email, password) values ('qgdomingo', 'Gio Domingo', 'qgdomingo@sample.com', 'secret');
insert into task_limit_config values ('qgdomingo', '10');
insert into users (username, name, email, password) values ('todo_test_user', 'Test User', 'testuser@sample.com', 'secret');
insert into task_limit_config values ('todo_test_user', '5');
insert into tasks (title, description, username) values ('Insert test task data', 'lorem ipsum 1', 'qgdomingo');
insert into tasks (title, description, username) values ('Get test data', 'lorem ipsum 2', 'qgdomingo');
insert into tasks (title, description, username) values ('Update test data', 'lorem ipsum 3', 'qgdomingo');
insert into tasks (title, description, username) values ('Delete test data', 'lorem ipsum 4', 'qgdomingo');