package data

const TodosTableCreationQuery = `CREATE TABLE IF NOT EXISTS todos
(
	id serial,
	content TEXT NOT NULL,
	user_id INTEGER NOT NULL,
	created_date DATE NOT NULL DEFAULT CURRENT_DATE,
	CONSTRAINT todos_PK PRIMARY KEY (id),
	CONSTRAINT todos_FK FOREIGN KEY (user_id) REFERENCES users(id)
)`

const InitialUserQuery = `INSERT INTO users (name, max_todo) VALUES ('test_user', 3)`

const UsersTableCreationQuery = `CREATE TABLE IF NOT EXISTS users (
	id serial,
	name TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
)`
