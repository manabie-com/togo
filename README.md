
## Run the application
First, you need to install maven and java with version >= 1.8.
Open terminal and  `cd` to folder `todotask`. After that, run this command  
`mvn spring-boot:run` to initialize the application.
## Verify running application
After initializing application successfully, you can test the api by sending this curl request.
```cookie
curl --location --request POST 'http://localhost:8080/task/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": 1,
    "task_name": "test",
    "task_description":"test",
    "target_date":"2022-07-08T00:00:00Z"
}'
```
## Run unit and integration tests
For unit testing and integration testing, run this command ```mvn clean test```

## Some 
This project uses H2, an in-memory database, for the purpose of making the application locally runnable without an external database like Postrges or Mysql. The schema and initial data defined in two files `schema.sql` and `data.sql`, are stored in folder ```src/main/java/resources```. The schema file is also used for testing. 