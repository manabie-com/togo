# TOGO

## Description

&emsp;&emsp;**Simple API using Golang, Postgresql and jwt for authentication**	

## 1. How to run your code locally?
- Git clone repository

	```bash
	git clone https://github.com/qanghaa/togo.git
	```
- Go to ***togo*** directory
- Using Docker to run Postgresql
	```bash
	docker build -t togo .
	```
	
	```bash
	docker run -dp 4001:5432 togo
	```
- Dependencies Installation

	```bash
	go mod download
	```
- Configuration Server

    - create `.env` file
	
    - and add your secret enviroment variables below 
     
    ```
  PORT=3000
  CONNECT_STR=postgres://postgres:manabie@localhost:4001/togo?sslmode=disable
  SECRET_TOKEN=<your_secret_token>
	```
- Run Server
  ```bash
  go run main.go
	```
    OR
	```bash
  go run github.com/manabie-com/togo
  ```
- Go to Sample 'curl' in below document :point_down:
## 2. Sample “curl” Command:
 &emsp;&emsp;API Document: [here](https://documenter.getpostman.com/view/15522883/UzBvHPBC)

## 3. How to run your unit tests locally?
  - Go to ***togo*** directory
  - type in cmd: 
	```bash
	go test ./...
	```

## 4. What do you love about your solution?
  &emsp;&emsp;Overall, nothing outstanding. However I quite like the payment feature. Although the implementation is simple, it will help the user become a Premium user LOL. This feature helps users to overcome the limit of creating tasks in 1 day (20 tasks/day) instead òf 10 as usual. Not related to feature, probably Docker, I spent 2 days learning and trying to work on it and I was able to use it in this project :D.
  
## 5. What else do you want us to know about however you do not have enough time to complete?
  &emsp;&emsp;Probably Testing. I haven't writed unit test enough possible scenarios with my API yet. 
