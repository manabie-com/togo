### Overview
This is a simple backend for a good old todo service, right now this service can handle login/list/create simple tasks.  

## How to run

```bash
> cp .env.example .env
> make run
```

## How to test

```bash
> make test
```
 
#### Backend working progress
|                                                            |     |
| ---------------------------------------------------------- | --- |
| Limit N daily tasks for each user                          | ✔   |
| Dockerized this service                                    | ✔   |
| Write integration tests                                    | ✔   |
| Make this code DRY                                         | ✔   |
| Write unit test for `services` layer                       | ✔   |
| Change from using `SQLite` to `Postgres` with `docker-compose`  | ✔   |
| Split `services` layer to `use case` and `transport` layer | ✔   |
| Write unit test for `storages` layer                       |    |

#### Potential improvements
- Change request method for login from GET to POST
- Hash user password
- Using middleware for handling authentication

### DB Schema
```sql
-- users definition

CREATE TABLE users (
	id TEXT NOT NULL,
	password TEXT NOT NULL,
	max_todo INTEGER DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);

INSERT INTO users (id, password, max_todo) VALUES('firstUser', 'example', 5);

-- tasks definition

CREATE TABLE IF NOT EXISTS tasks (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    created_date DATE NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE INDEX IF NOT EXISTS idx_created_date ON tasks (created_date);
```

### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
