# Togo

## Table of contents
* [Usage](#usage)
* [Features](#features)
* [Diagrams](#diagrams)
* [Easy Start](#easy-start)

## Usage
This is a Postman collection which contains REST API of Togo service.

[Togo collection](docs/togo_v2.postman_collection.json)
## Features

These are the main features in this repo.

:heavy_check_mark: Login

:heavy_check_mark: List Tasks

:heavy_check_mark: Add Tasks

These are some function I want to improve but I don't have enough time.

:x: Create User

:x: Password Encryption

:x: Implement Transport layer interface

## Diagrams
1. ### Sequence Diagram

![Sequence](https://raw.githubusercontent.com/nohattee/togo/master/docs/togo-sequence.svg)

2. ### ERD Diagram

![ERD](https://raw.githubusercontent.com/nohattee/togo/master/docs/togo-erd.svg)

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
docker-compose up --build -d
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

