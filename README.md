## Overview
To make it run:
Import Postman collection from docs to check example
- `docker-compose up`
- Import Postman collection from `docs` to check example

## What I have (and have not) accomplished

- [x] Daily task limit functionality
- [x] Switch from SQLite to Postgres with `docker-compose`
- [x] Unit tests for `service` layer
- [x] Integration tests
- [x] DRY code
- [x] Change GET `/login` to POST `/login` to prevent user's info exposes in url
- [x] Use pgcrypto to hash password, only store password hash in postgres,
- [x] `usr` table and `task` table use identity column with integer type instead of text type
- [x] Build middlewares
- [x] Graceful shutdown
- [ ] (Optional) Unit tests for `storages` layer
- [ ] (Optional) Split `services` layer to `use_case` and `transport` layer

## Potential improvements
