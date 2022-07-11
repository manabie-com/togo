### Run the application

First, you need to install `maven` and `java` with `version >= 1.8`.
Open terminal and  `cd` to folder `todotask`. After that, run this command `mvn spring-boot:run` to run the
application.

### Verify running application

After running application successfully, you can test the api by sending this `CURL` request.

```cookie
curl --location --request PUT 'http://localhost:8080/task/add' \
--header 'Content-Type: application/json' \
--data-raw '{
    "user_id": 1,
    "task_name": "test",
    "task_description":"test",
    "target_date":"2022-07-08T00:00:00Z"
}'
```

### Run unit and integration tests

For unit testing and integration testing, run this command ```mvn clean test```

### Some highlights about the project

1. This application contains three layer: controller, service and repository. Each layer communicates with each other
   via interface not their concrete methods. Therefore, each layer is easily replaceable and extendable.
2. `ZonedTimeDate` is used instead of `Date`. This new class helps us maintain daily limit logic for different users in
   different timezones.
3. `Spring Validation` and `Global Exception Handling` are used to passively verify request format and throws predefined
   response if request format is invalid. For example, if field `user_id` or `target_date` in `AddTaskRequest` is null,
   application will return a bad request response immediately.
4. This project uses H2, an in-memory database, for the purpose of making the application locally runnable without an
   external database like Postrges or Mysql. The schema and initial data defined in two files `schema.sql`
   and `data.sql`, are stored in folder ```src/main/java/resources```. The schema file is also used for testing.

### Other notes

1. Database and server configuartion are configure in file ```src/main/resources/application.properties```. Server will
   start at the default port `8080`
