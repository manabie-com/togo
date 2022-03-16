## Description

This repository for testing interview with Manabie.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

- NodeJS (version 12.6 or higher, I'm using v17.0.1)

- Postgres


## Installation environment

```bash
# Install npx
$ npm install npx

# Install yarn
$ npm install --global yarn

# Invoke the Prisma CLI by prefixing it with npx
$ npx prisma
```

## Setting up
Follow all step bellow to setup your dev environment
1. Start your Postgres (I'm using docker-compose for environment setup)
```bash
$ docker-compose up -d
```

2. Setup environment variables.
```bash
$ cp .env.example .env
```

3. Install NPM packages

```bash
yarn install
```

## Running the app

```bash
# To map your data model to the database schema, you need to use the prisma migrate CLI commands
$ yarn migrate

# development
$ yarn start

# watch mode
$ yarn start:dev
```

## Test

```bash
# Unit tests: Let's remove all rows in table "Task", "User" on your local Database Postgres, then run test
$ yarn test
```

## Sample API
- After start app, let's connect this link on your browser: http://localhost:3000/graphql

```bash
# Signup
mutation signup($createUserInput: CreateUserInput!) {
  signup(createUserInput: $createUserInput) {
    id
    email
    name
    maxJob
    token
  }
}
----------------------------------------------
# Query Variables:
{
  "createUserInput": {
    "email": "vu.nguyen@gmail.com",
    "name": "vu nguyen",
    "password": "123456",
    "maxJob": 5
  }
}
```

```bash
# Login
mutation login(
$loginUserInput: LoginUserInput!
) {
  login(loginUserInput: $loginUserInput) {
    id
    email
    name
    maxJob
    token
  }
}
----------------------------------------------
# Query Variables:
{
  "loginUserInput": {
    "email": "vu.nguyen@gmail.com",
    "password": "123456"
  }
}
```

```bash
# createTask
mutation createTask(
$createTaskInput: CreateTaskInput!
) {
  createTask(createTaskInput: $createTaskInput) {
    id
    createdAt
    updatedAt
    title
    content
    userId
    user {
      id
      name
      email
      maxJob
    }
  }
}
----------------------------------------------
# Query Variables:
{
  "createTaskInput": {
    "title": "task1",
    "content": "this is content of task1"
  }
}
# HTTP HEADERS:
{
  "Authorization": "replace with token get from login API"
}
```

```bash
# taskById
query taskById(
$id: Float!
) {
  taskById(id: $id) {
    id
    createdAt
    updatedAt
    title
    content
    userId
    user {
      id
      name
      email
      maxJob
    }
  }
}
----------------------------------------------
# Query Variables:
{
  "id": 1
}
# HTTP HEADERS:
{
  "Authorization": "replace with token get from login API"
}
```

```bash
# tasks
query tasks(
$searchString: String
$skip: Float
$take: Float
$orderBy: TaskOrderByUpdatedAtInput
) {
  tasks(searchString: $searchString, skip: $skip, take: $take, orderBy: $orderBy) {
    id
    createdAt
    updatedAt
    title
    content
    user {
      id
      email
      name
      maxJob
    }
    userId
  }
}
----------------------------------------------
# Query Variables:
{
  "searchString": "task",
  "skip": 0,
  "take": 3,
  "orderBy": {
    "updatedAt": "desc"
  }
}
# HTTP HEADERS:
{
  "Authorization": "replace with token get from login API"
}
```

## Passion
- I love new architect, and I see the prisma is simple to buid faster app and it's good connect with PostgreSQL and other.
- I'm learning NestJS, so I've just apply in this project
- I think NestJS is new, reliable, scalable framework NodeJS.
- NestJS support for strongtyped like TypeScript.
- Especially, I love graphql, a query language for APIs.