# Togo

## Table of contents
* [Features](#features)
* [Easy Start](#easy-start)


## Features

These are the main features in this repo.

:heavy_check_mark: Login

:heavy_check_mark: List Tasks

:heavy_check_mark: Add Tasks

These are some function I want to improve but I don't have enough time.

:x: Create User

:x: Password Encryption

:x: Implement Transport layer interface

## Easy Start
This section aims to show you how to run Togo service.

1. ### Installation
```
git clone https://github.com/nohattee/togo
cd togo
cp .env.example .env
```
2. ### Run unit test & integration test
```
docker-compose up --build
docker-compose exec db psql -U togo_user togo_db
```
In the psql console type this:
```
create database togo_db_test;
exit;
```
Then, you will back to terminal console, type this:
```
go test ./...
```
3. ### Run without test
```
docker-compose up --build
```

