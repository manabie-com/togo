## System requirement
- Java openjdk version 16 ([Download](https://jdk.java.net/16/))
- Maven >= 3.5 ([Install](https://maven.apache.org/install.html))

## Summary
This solution uses spring boot version `2.4.7` with following starters
- Web
- Spring JPA for data layer
- Spring security for managing authentication
Database migration use flyway for versioning the change

## How to run
```bash
mvn clean package
docker-compose -f ./app-meta/docker-compose.yml up -d
```

## Test
Using [`test.http`](./test.http) to make request to localhost