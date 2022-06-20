# Pham Thanh Chi submission

## What I have done
- Decoupling storage and rate limit manager. Now you can use Postgres, SQLite, or anything you implement for them.
- Implemented Postgres as Database and Redis as a rate limiter  
- Integration test for togo service

## What I will do when I have time
- Unit test for services, storage, and common
- Re-structure the services again for cleaner code and easier unit-test. For now, it is quite a bit messing for me

## How to run

### Start Postgres and Redis for testing
```bash
 docker-compose up -d postgres redis
```

### Start togo service
```bash
 docker-compose up -d togo
```