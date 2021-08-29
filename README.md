# TOGO Project

This is an upgraded back-end service based on the current Togo project

- Provide full features of Login/List/Create Tasks that were available before
- Build on NodeJS Express platform
- Migrate database from SQLite to Postgresql
- Improve and upgrade the DB structure to achieve high efficiency in performance and security
- There is backward compatibility with the interface that the client has integrated before

# Prerequisite

1. docker-compose: Your device have to install docker-compose before
2. Fork/Clone this repo and move to project folder
3. Please make sure is there not any running services in your device duplicate PORT with docker services, includes:
    1. Postgresql: 5432
    2. Postgres Admin: 8080
    3. Redis: 6379
    4. Application: 5050

P/S: It should be allow change docker services PORT dynamically, but it not available at time. It will handle later.

# Run project

## Docker and docker-compose

In the project folder, just run the below cmd, it will pull some images and build to container:

```
$ docker-compose up -d
```

P/S:

1. While image is being pull, it will run both unit test and integration test inside. If all passed, the service is make sure could be use. That's mean you don'
   t need to run any script for testing service.
2. When console run completely, wait more few seconds for app initialize
3. After that, you can access the APIs Service by: http://localhost:5050

# Optimize

I have found some issue from DB structure and have optimized them:

## User Entity

- Column id and password is defined data type as TEXT. It can't be limited and validated the length, this will cause the effect about the performance of space
  when storage and time when querying. So I edit those to VARCHAR with fixed-length for each one.
- However, the id column is being set as Primary Key, if it is TEXT it will, the indexing and the querying will max out more RAM.
- The max_todo column is typed as INTEGER, the value will use 4 bytes but the actual value won't need to be that big, so I changed it back to SMALLINT to save
  more
- The password column is storing a plain text value, which is not good in terms of security if SQL Injection is encountered. This is not a matter of the schema,
  but more about the backend logic. We can replace it by storing encrypted passwords to avoid risks. However, this issue is still handled in parallel to ensure
  backward compatibility with the client side

```
CREATE TABLE users (
	id VARCHAR(20) NOT NULL,
	password VARCHAR(255) NOT NULL,
	max_todo SMALLINT DEFAULT 5 NOT NULL,
	CONSTRAINT users_PK PRIMARY KEY (id)
);
```

## Task Entity

- The current column is storing the UUID code and the identifier is the primary key. Similar to above, instead of TEXT type, we will change it back to UUID
  type.
- Set default for id, assigning the id value to this field is not from the client side but is handled by the back-end. If this schema is replaced with a
  backend/framework, the developer may forget about this and throw a null error when inserting.
- user_id is a foreign key that references the user's id, so the data type should be the same as the id of table users.
- created_date is a field to store date values, so use DATE type to manage Date objects according to the system (If you need to manage time, use TIMESTEMPZ).
  With DATE type will avoid query problems if the date format is not the same, eg YYYY-MM-DD and YYYY/MM/DD or DD/MM/YYYY

```
CREATE TABLE tasks (
	id UUID NOT NULL DEFAULT gen_random_uuid(),
	content TEXT NOT NULL,
	user_id VARCHAR(20) NOT NULL,
	created_date DATE NOT NULL DEFAULT Now(),
	CONSTRAINT tasks_PK PRIMARY KEY (id),
	CONSTRAINT tasks_FK FOREIGN KEY (user_id) REFERENCES users(id)
);
```

# What I'm missing

- Organize ORM to map Model Entity and Repository with Database instead of simply processing query text as currently
- Build a DB Connector allow adapting other database, then I can switch to other db without having to change in the code
- Configure docker-compose to allow porting of services if there are other services running on the same port on OS
