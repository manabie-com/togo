# TOGO

## Description

&emsp;&emsp;**Simple API using Golang, Postgresql and jwt for authentication**	

## 1. How to run your code locally?
- Requirements
	- Install [Docker Engine](https://docs.docker.com/engine/install/)
	- Install [docker-compose](https://docs.docker.com/compose/install/)
- Git clone repository

	```bash
	git clone https://github.com/qanghaa/togo.git
	```
- Go to ***togo***:open_file_folder: directory
- Using `docker-compose` commands
	```bash
	# this command make sure next command working as expected
	sudo docker-compose down --volumes .
	```
	
	```bash
	sudo docker-compose up
	```
	
## 2. Sample “curl” Command:
 &emsp;&emsp;API Document: [here](https://documenter.getpostman.com/view/15522883/UzBvHPBC)
 <br> &emsp;&emsp;Note: ***Using Bearer Authorization Header for endpoints required*** 

## 3. How to run your unit tests locally?
  - Go to ***togo*** directory
  - type in cmd: 
	```bash
	go test ./...
	```

## 4. What do you love about your solution?
  &emsp;&emsp;Overall, nothing outstanding. However I quite like the payment feature. Although the implementation is simple, it will help the user become a Premium user LOL. This feature helps users to overcome the limit of creating tasks in 1 day (20 tasks/day) instead òf 10 as usual. Not related to feature, probably Docker, I spent 2 days learning and trying to work on it and I was able to use it in this project :D.
  
## 5. What else do you want us to know about however you do not have enough time to complete?
Probably Testing. I haven't writed unit test enough possible scenarios with my API yet. 
