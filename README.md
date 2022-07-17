## Setup

### Requirements

1. Install Docker

   - Install Docker Desktop `4.10.1` (current version) or Docker Engine `19.03.0` and Docker Compose `3.8`

2. Update your /etc/hosts

   Put the following configurations on your host machine's /etc/hosts.

   ```
   #Datastores
       127.0.0.1 postgresql.manabie.todo
   ```

### How to run your code locally?

1.  Run project:

    #### Run:

         make up

    #### Migrate: (Used when the new database is initialized)

         docker exec -it api.manabie.todo bash -c "make migrate-todo"

2.  Use these commands to manage your containers:

    #### Install

        make install

    #### Run

        make run

    #### Format

        make format

    #### Migrate DB

        make migrate-todo

    #### Drop DB

        make drop-todo

### A sample “curl” command to call your API

1.  Users:

    - Get list user

      ```
      curl --request GET 'http://localhost:8080/users'
      ```

2.  Setting limit for user

    - Show setting by userId

      ```
      curl --location --request GET 'http://localhost:8080/users/1/settings'
      ```

    - Create limit task by userId

      ```
      curl --location --request POST 'http://localhost:8080/users/1/settings' \
      --header 'Content-Type: application/json' \
      --data-raw '{
      "limit_task": 5
      }'
      ```

    - Update limit task by Id

      ```
      curl --location --request PUT 'http://localhost:8080/settings/1' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "limit_task": 10
      }'
      ```

3.  Show/Insert/Update/Delete todo task for user.

    - List task by user-id

      ```
      curl --location --request GET 'http://localhost:8080/users/1/tasks'
      ```

    - Create task by user-id

      ```
      curl --location --request POST 'http://localhost:8080/users/1/tasks' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "content": "whatever",
          "target_date": "2022-07-17"
      }'
      ```

    - Get task by id

      ```
      curl --location --request GET 'http://localhost:8080/tasks/1'
      ```

    - Update task by id

      ```
      curl --location --request PUT 'http://localhost:8080/tasks/1' \
      --header 'Content-Type: application/json' \
      --data-raw '{
          "id": 1,
          "member_id": 1,
          "content": "updated",
          "target_date": "2022-07-17T00:00:00Z",
          "created_at": "2022-07-17T14:44:46.981938Z"
      }'
      ```

    - Delete task by id

      ```
      curl --location --request DELETE 'http://localhost:8080/tasks/1'
      ```

### How to run your unit tests locally?

- Run a command in a running container (container: `api.manabie.todo`)

  #### Test (unit test)

        make test

  #### Test E2E

        make test-e2e

### What do you love about your solution?

- Build projects in different environments by `Docker`.
- Ensure data consistency when using `Transactions` and `Locking`.
- E2E testing and mock testing.
- Clean Architecture in Go by Repository pattern. ([reference](https://github.com/bxcodec/go-clean-arch#the-diagram))

### What else do you want us to know about however you do not have enough time to complete?

- Setup information (SECRET, PUBLIC) key by environment. (ansible, vault).
- Authentication and authorization for endpoints.
- Restore data when make test.

<hr />


### Notes

- We're using Golang at Manabie. **However**, we encourage you to use the programming language that you are most comfortable with because we want you to **shine** with all your skills and knowledge.

### How to submit your solution?

- Fork this repo and show us your development progress via a PR

### Interesting facts about Manabie

- Monthly there are about 2 million lines of code changes (inserted/updated/deleted) committed into our GitHub repositories. To avoid **regression bugs**, we write different kinds of **automated tests** (unit/integration (functionality)/end2end) as parts of the definition of done of our assigned tasks.
- We nurture the cultural values: **knowledge sharing** and **good communication**, therefore good written documents and readable, organizable, and maintainable code are in our blood when we build any features to grow our products.
- We have **collaborative** culture at Manabie. Feel free to ask trieu@manabie.com any questions. We are very happy to answer all of them.

Thank you for spending time to read and attempt our take-home assessment. We are looking forward to your submission.
