# Togo

Golang application which accepts a todo task and records it if the user has not yet reached the limited number of tasks per day.

## Usage

### I. Run server

Make sure you have Docker installed, otherwise you may download it from this [link](https://www.docker.com/products/docker-desktop/).

If you already have Docker, simply follow the steps to deploy the code.

1. Clone the repository

```Shell
$ git clone git@github.com:jrpespinas/togo.git
```

2. Change directory

```Shell
$ cd togo
```

3. Run docker compose to deploy the application

```Shell
$ docker compose up
```

### II. Sample Request

Once the server is running you may make a simple post request to create a task

```Shell
$ curl -X POST http://localhost:8080/tasks -H 'Content-Type: application/json' -d '{"title":"sample title","description":"sample description"}'
```

You should have received a response such as this:

```json
{
  "status": "Success",
  "code": 200,
  "message": {
    "id": "cabhte81hrh6mgum9d7g",
    "title": "sample title",
    "description": "sample description",
    "created_at": "2022-06-01T08:09:29.0564148Z"
  }
}
```
