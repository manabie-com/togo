# TOGO

## Description

**Simple API using Golang, Postgresql and jwt for authentication**	

## How to run your code locally?
- Git clone repository

	```bash
	git clone https://github.com/qanghaa/togo.git
	```
- Dependencies Installation

	```bash
	go install
	```
- Configuration Server

	create `.env` file
	
	add your secret enviroment variables below
    ```
  PORT=3000
  CONNECT_STR=postgres://<your_username>:<your_password>@localhost/<your_db_name>?sslmode=disable
  SECRET_TOKEN=<your_secret_token>
	```
- go run main.go

	```bash
  OR 
  ```
- go run github.com/manabie-com/togo

## Sample “curl” Command:** [here](https://documenter.getpostman.com/view/15522883/UzBvHPBC)

## How to run your unit tests locally?
  - Go to ***togo*** directory
  - type in cmd: ```go run test ./...```

## What do you love about your solution?
  Overall, nothing outstanding. However I quite like the payment feature. Although the implementation is simple, it will help the user become a Premium user LOL.
  This feature helps users to overcome the limit of creating tasks in 1 day (20 tasks/day) instead òf 10 as usual.
  
## What else do you want us to know about however you do not have enough time to complete?
  Probably Docker. I spent 2 days learning and trying to work on it. However I failed and can't make it easier for you to use my source code.
