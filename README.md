### Submission
- Which functions have been done:
  - APIs to create, read, update, and delete todo tasks.
  - Simple authorization mechanism when creating task.
  - Rate limit of tasks per user can be added per day.

- How I do:
  - APIs: 
    - Follow RESTful architecture.
    - Use embedded database(H2) to save tasks(tasks table)- which contains task information, user-config(user_config table) - which contains rate limit config per user.
    - Use embedded Redis to save rate limit config, and added task times per user.
  - Authorization mechanism: base on HTTP Authorization request header and value is userId - a numeric string, ex: "123123", "34241", "647544"
  - Rate limit mechanism: 
    - Use Redis Hashes to store and check the rate limit config.
    - Use HINCRBY to update the rate limit per user, to guarantee the amniotic operations.

- Pros and Cons:
  - Pros:
    - Guarantees each Redis hashes save rate limit counter contains a maximum of 1000 entries, which can best suit with hash-zip map-max-entries Redis setting and can reduce the used memory efficiently. Because I save one hash key per user, which can populate to millions when active users shoot up.
    - Quite fast with Redis to storage and guarantees atomic get and check rate limit operation.
    - Implement basic handler application exception.
  - Cons:
    - Simple function.
    - Embedded database and embedded Redis need to remove in the production application(I use them for easy integration purposes).
    - HINCRBY Redis function will not work as expected when 1 user creates more than 2000 tasks at a time.
- What I will complete if have enough time:
  - Remove steal redis cached.
  - Handle application exception.
- How to run my code locally:
  - Enviroment: jdk8, docker.
  - Existed data when start app:
    - Tasks table:

    |ID| NAME|CONTENT|USER_ID|  
    |------------------------------------|-------- |-------------------|-------|
    |4df7f33a-a32c-40f9-aa30-0a598348c302|	test-1 |	content test-1   |1231231|
    |71b2649e-5fb3-440a-b810-c25772a4c15f|	test-2 |	content test-2   |1231231|
    |be6bad9d-3492-4d2b-a90c-f13083ef9efe|	test-3 |	content test-3   |1231232|
    |95fd9350-e9d8-48b2-8b39-ee94667df6c1|	test-4 |	content test-4   |1231232|
    |359240af-927f-477a-bbe7-ddf90829a080|	test-5 |	content test-5   |1231233|
    |8810c75d-6678-445a-b906-8b8ca8ef7a52|	test-6 |	content test-6   |1231233|
    |9b8918c1-8a9f-4858-a1cd-4e6060eca82c|	test-7 |	content test-7   |1231234|
    |69c1206f-bcbf-4f25-b599-e6080c98b60d|	test-8 |	content test-8   |1231234|
    |cfdefc6b-7ff1-4cef-884e-d27b92bb991b|	test-9 |	content test-9   |1231235|
    |248c89d4-17ea-4c36-8522-f5886bbe51c1|  test-10|    content test-10|1231235|
    - User Configs table:

    |ID|USER_ID|CONFIG_TYPE|VALUE|  
    |------------------------------------|-------- |-------------------|-------|
    |7a2b1b74-7b5c-4ada-9cd4-64c6b36c73c7|	1231231 |	RATE_LIMIT_ADD_TASK_PER_DAY|	10 |
    |d93ba141-81a8-4c19-8bbb-5878497cbfa1|	1231232 |	RATE_LIMIT_ADD_TASK_PER_DAY|	11 |
    |010270f8-89f1-4342-bee2-10bf59f29da8|	1231233 |	RATE_LIMIT_ADD_TASK_PER_DAY|	5  |
    |938be0c2-bc39-45bd-8360-d2cded6dfc42|	1231234 |	RATE_LIMIT_ADD_TASK_PER_DAY|	15 |
    |ae4e4918-2dde-4d22-9bb1-d0f8e2115379|	1231235 |	RATE_LIMIT_ADD_TASK_PER_DAY|	20 |
    |660d81f7-9359-4384-9ab8-e9f82e4fa4e7|	23223332|	RATE_LIMIT_ADD_TASK_PER_DAY|	10 |
    |007d8a96-8069-4d8e-ac4b-464c08c43bda|	23223333|	RATE_LIMIT_ADD_TASK_PER_DAY|	5  |
    |0af65b14-94c3-466c-a85c-c7f98ac52433|	23223334|	RATE_LIMIT_ADD_TASK_PER_DAY|	5  |
    |0d94c99a-684b-4b67-8d93-67bb2503d064|	23223335|	RATE_LIMIT_ADD_TASK_PER_DAY|	7  |
    |37edf7a8-ee1c-4720-914b-5349ccc7dc9c|	23223336|	RATE_LIMIT_ADD_TASK_PER_DAY|	8  |
    |347a462b-5f4d-4277-91e7-959f52c7add6|	23223337|	RATE_LIMIT_ADD_TASK_PER_DAY|	9  |

  - Start App:
    - need to install docker
    - move to todo/todo-task-application
    - run script to build docker images: "docker build -t todo-task-app ."
    - run script to run docker container: "docker run -p 8888:8888 todo-app-v2"
    - curl:
      - create tasks: 
        - curl --location --request POST 'http://localhost:8888/todo-task/v1/tasks' \
          --header 'Authorization: 1231231' \
          --header 'Content-Type: application/json' \
          --data-raw '{
              "name": "123344",
              "content": "test1 - content"

          }'
      - get tasks:
        - curl --location --request GET 'http://localhost:8888/todo-task/v1/tasks'
      - get tasks/{id}:
        - curl --location --request GET 'http://localhost:8888/todo-task/v1/tasks/a66d5aa4-40d1-4f07-9cd6-46e3355e6c20'
      - update tasks/{id};
        - curl --location --request PUT 'http://localhost:8888/todo-task/v1/tasks/1a229915-41cd-4f69-8846-3545c526b0fd' \
          --header 'Content-Type: application/json' \
          --data-raw '{
              "name": "test-121312aaaaaaa3",
              "content": "cont232333333333ent test-1"
          }'
      - delete tasks/{id}:
        - curl --location --request DELETE 'http://localhost:8888/todo-task/v1/tasks/1a229915-41cd-4f69-8846-3545c526b0fd'
  - Run Unit Test:
    - move to todo/todo-task-application
    - run script: "mvn -U clean test"
  - Run Integration Test:
    - move to todo/todo-task-application
    - run script: "mvn failsafe:integration-test"
  

### Requirements

- Implement one single API which accepts a todo task and records it
  - There is a maximum **limit of N tasks per user** that can be added **per day**.
  - Different users can have **different** maximum daily limits.
- Write integration (functional) tests
- Write unit tests
- Choose a suitable architecture to make your code simple, organizable, and maintainable
- Write a concise README
  - How to run your code locally?
  - A sample “curl” command to call your API
  - How to run your unit tests locally?
  - What do you love about your solution?
  - What else do you want us to know about however you do not have enough time to complete?

### Notes

- We're using Golang at Manabie. **However**, we encourage you to use the programming language that you are most comfortable with because we want you to **shine** with all your skills and knowledge.

### How to submit your solution?

- Fork this repo and show us your development progress via a PR

### Interesting facts about Manabie

- Monthly there are about 2 million lines of code changes (inserted/updated/deleted) committed into our GitHub repositories. To avoid **regression bugs**, we write different kinds of **automated tests** (unit/integration (functionality)/end2end) as parts of the definition of done of our assigned tasks.
- We nurture the cultural values: **knowledge sharing** and **good communication**, therefore good written documents and readable, organizable, and maintainable code are in our blood when we build any features to grow our products.
- We have **collaborative** culture at Manabie. Feel free to ask trieu@manabie.com any questions. We are very happy to answer all of them.

Thank you for spending time to read and attempt our take-home assessment. We are looking forward to your submission.

