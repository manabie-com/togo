### Requirements/Specification

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limit.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Write a concise README (see readme proper)
  - How to run your code locally? 
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?

### System Requirements
 * [Jdk 11+](https://www.oracle.com/ph/java/technologies/javase/jdk11-archive-downloads.html)
 * [Maven 3.8.5](https://maven.apache.org/download.cgi)
 * [Postgres Server 11+](https://www.postgresql.org/download/)
 
 
 ### env Setup
* Ensure that the specified versions of requirements are installed.
* Ensure postgres server connectivity
LinuxOS/MacOS:
> `psql -h localhost -U $psql_user`
* create database
> `CREATE DATABASE todo-test`

### Running/Building locally via cli
* navigate into project's root 
> `cd $projectDir\`
* clean project
> `mvn clean`
* install dependencies 
> `mvn clean install` 
* run migration scripts
> `mvn liquibase:update`
* Run the project
> `mvn spring-boot:run`

### Sample cURL request for testing and health check
> `curl --location --request GET 'http://localhost:8008/api/v1/version'`


### Run Unit Tests
* navigate into project's root
> `cd $projectDir\ `
* run test command
> `mvn clean test`


### What do I love about my solution?
* Microservice ready.
