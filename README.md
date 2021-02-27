## Overview
To make it run:
Import Postman collection from docs to check example
- `docker-compose up -d pg`
- `go run main.go`
- Import Postman collection (modified) from `docs` to check example

Or
- `docker-compose up -d pg`
- `docker-compose up -d togo`

## What I have (and have not) accomplished
- [x] Daily task limit functionality
- [x] Switch from SQLite to Postgres with `docker-compose`
- [x] Unit tests for `service` layer
- [x] Integration tests
- [x] DRY code
- [x] Change GET `/login` to POST `/login` to prevent user's info exposes in url
- [x] Use pgcrypto to hash password, only store password hash in postgres,
- [x] `usr` table and `task` table use identity column with integer type instead of text type
- [x] Change `create_date` column to `create_at` column with `timestamptz` type 
- [x] Build middlewares
- [x] Graceful shutdown
- [x] Use environment variables (`.env` file) to store postgres parameters and use them in app  
- [ ] (Optional) Unit tests for `storages` layer
- [ ] (Optional) Split `services` layer to `use_case` and `transport` layer

## Potential improvements
- Use more environment variables to store application parameters to prevent sensitive information leakages and conveniences.
- Give users JWT refresh token so that they do not have to log in again.
- Api response error more clearly.
- More detail document for code.  
- Improve project package structure.  
- More unit tests (especially for `storages` layer), integration tests.