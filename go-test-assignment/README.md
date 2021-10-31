### Setup database

1. Pull & run docker iamge

```bash
docker run -it -d -p 5432:5432 --name postgres-local -e POSTGRES_PASSWORD=password postgres
```

2. Go into docker container

```bash
docker exec -it postgres-local bash
```

3. Login pg

```bash
psql -h localhost -U postgres
```

4. Create user

```bash
CREATE USER "user_login" WITH PASSWORD 'password';
ALTER ROLE "user_login" WITH SUPERUSER;
```

5. Create Database

```bash
CREATE DATABASE dbtest;
```

6. Export ENV

```bash
export PG_HOST=127.0.0.1
export PG_PORT=5432
export PG_USER=user_login
export PG_PASS=password
export PG_DB=dbtest
export PORT=9090
```

6. Insert test data, please take a look for db.sql

### Install go packages & run app

1. Install go packages

```bash
go get -u ./...
```

2. Init go vendor for go modules management

```bash
go mod vendor
```

3. Run app

```bash
go run main.go
```

### Explaining project

1. No one uses http GET for login.

- Instead of GET to POST.

2. Change NoSQL Live to Postgres

3.Update columns.

- No one uses ID of column with TEXT type, expect ony NoSQL using uuid -> Change to bigserialize
- Change created_date from TEXT to timestamp and set default and create INDEX for it
- Change user_id of task to INT8 and update INDEX for it
- Add a username into Users table and check UNIQUE
- No one exposes raw password, which is very risk. I applied bcrypt hash to encrypt password.

4. DDD architecture

- Currently, I applied DDD (3 layers) and used self-container components, which helps me narrow down business logic in a single directory
- Domain contains endpoints and logic core
- Infrastructure contains drivers, services, providers, middleware...etc
- Pkgs contains utils, common, re-useable codes

5. TODO if I have time
- Apply redis
- Black list token after user logout.
- Apply golint for checking clean code

6. Apply unit test
