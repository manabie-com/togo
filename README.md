# Pham Thanh Chi submission

## What I have done
- Decoupling storage and rate limit manager. Now you can use Postgres, SQLite, or anything you implement for them.
- Implemented Postgres as Database and Redis as rate limiter  
- Integration test for togo service

## What I will do when I have time
- Unit test for services, storages, and common
- Re-structure the services again for cleaner code and easier for unit-test, for now it quite a bit messing for me

## How to run

### Start postgres and redis
```bash
 docker-compose up -d
```

### Start togo server
```bash
 go run main.go
```