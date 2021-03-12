# Todo Backend Service

## Overview
- Todo Backend Service is a simple backend for a good old todo service, right now this service can handle login/list/create simple tasks.  

## Install
### Backend
#### With Docker

_this is the recommended way to run todo backend service_

* copy provided `docker-compose.yml` and customize for your needs
* make sure you add `ADMIN_PASSWD=something...` for any SQL Connection
* pull prepared images from the DockerHub and start - `docker-compose pull && docker-compose up -d`
* alternatively compile from the sources - `docker-compose build && docker-compose up -d`

#### Without Docker

* download archive Go [stable release](https://golang.org/dl/) for your OS
* go run ./cmd/server.go 


## API Spec:
- We use Postman for write API Spec 
- Import Postman collection from `docs` to check example




## DB Schema
// TODO REPLACE SQL DB SCHEMA TO USE MIND MAP OR CLASS ENTITY DIAGRAM FOR READABILITY

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

CREATE TABLE tasks (
	id TEXT NOT NULL,
	content TEXT NOT NULL,
	user_id TEXT NOT NULL,
    created_date TEXT NOT NULL,
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
```

## What I have done
* Write integration tests for this project
* Make this code DRY
* Write unit test for services layer
* Split services layer to use case and transport layer follow clean arch
* Change from using SQLite to Postgres with docker-compose
## What I need to improve for not having enough time for complete 
* Write integration tests
* Write more document and API docs use swagger or something instead of postman collection

### Sequence diagram
![auth and create tasks request](https://github.com/manabie-com/togo/blob/master/docs/sequence.svg)
