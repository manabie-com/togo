# Togo

## Table of contents

* [Postman Collection] (#postmant-collection)
* [Features] (#features)
* [Diagrams] (#diagrams)
* [Start] (#start)


## Postman Collection

This is a Postman collection contains REST API of Togo service.
[Postman Collection] (docs/togo.postman_collection.json)

## Features

There are a few main features in this repo.
:heavy_check_mark: Login
:heavy_check_mark: Retrieve tasks
:heavy_check_mark: Create task

There are also a few functionalities/ultilities I want to improve but I don't have enough time.

:x: Apply Swagger(Open API 3.0) tools to create docs and help with generate boilerplate codes for request/response definitions.
:x: Implement user permission to have an admin account to create many other users and also implement API for create/update/delete users.

## Diagrams

1.  ### Sequence Diagram

![Sequence] (https://raw.githubusercontent.com/mirageruler/togo/master/docs/togo-sequence.svg)

2. ### ERD Diagram

![ERD] (https://raw.githubusercontent.com/mirageruler/togo/master/docs/togo-erd.svg)

## Start

Follow these steps to run Togo service.

1. ### Installation

```
git clone https://github.com/mirageruler/togo
cd togo
cp .env.example .env
```

2. ### Run uint test & integration test

```
docker-compose up --build -d
docker-compose exec db psql -U togo_user togo_db
```

In the psql console type this:

``` 
create database togo_db_test;
exit;
```

Finally, comeback to terminal console, type this:

```
go test ./...
```

3. ### Run without test

```
docker-compose up --build
```