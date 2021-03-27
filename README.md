### Overview
To make it run:
 - Create `.env` file at root folder, then copy content of `.env.example` into it.
 - Run `docker-compose up`.

Or:
 - Run `make init-local`.
 - Run `make run`.

### System requirement
 - Installed `docker` and `docker-compose`
 - Installed Golang ~v1.14

### What i have done
 - Change identify field of `users` and `tasks` table to `integer` instead of text type.
 - Change `created_date` from `text` to `datetime`. It make availability for filter by datetime.
 - Dockerized this app.
 - Use environment variables to store application configs.
 - Switch database from SQLite to Postgres.
 - Use `pgcrypto` to hash and compare password.
 - Change login request method from `GET` to `POST` to prevent user's info expose in URL.
 - Split `service` to `domain` layer and `usecase` layer.
 - Move `storages` to `data` layer.
 - Implement clean architecture.
 - Write tests.

### Potential improvements
 - Write auth middleware
 - Write unit test for data layer
 - Write more details comments in source code
 - (optional) Implement refresh token function