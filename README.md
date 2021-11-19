# Manabie TODO apis

## Introduction

- This is a simple backend for a todo service, right now this service can handle login/list/create simple tasks, to make it run
- Right now a user can add as many tasks as they want, we want the ability to limit N tasks per day.
- For example, users are limited to create only 5 tasks only per day, if the daily limit is reached, return 4xx code to the client and ignore the create request.

## Documentation

- Setting port, jwt token in env file
- Sqlite database path: manabie > databases > storages > data.db
- Testing user: user_id:firstUser && password:example
- Import ManabieTODO.postman_collection.json to postman

## Requirements

This library requires Go 1.11 or later.