
## How to run

Follow steps below to run

```bash
# Start database postgres
$ make startDB

# Run seed database
$ make seedDB

# Test
$ make test

# Stop
$ make stop
```

## What is missing
- Hash password user
- Structure project
- Validate

## What is done
 - Structure project follow [The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
 - Integration tests
 - Unit tests

## Want to improve
- Support user delete a task or many tasks
- Limit N task per day base on timezone of user
- Cover full unit test
- Manage dependency injection by [uber fx](https://github.com/uber-go/fx)

### DB Schema
```sql
-- users definition

CREATE TABLE users (
	id INTEGER primary_key NOT NULL,
	email varchar(256) unique NOT NULL,
	password varchar(512) NOT NULL,
	max_todo INTEGER NOT NULL,
	created_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
);


-- tasks definition

CREATE TABLE tasks (
	id INTEGER primary_key NOT NULL,
	email varchar(2048) NOT NULL,
	status varchar(64) NOT NULL,
	user_id INTEGER NOT NULL,
    created_date timestamp NOT NULL,
);
```
