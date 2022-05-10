# Task Todo
## Guilde run in Intellij IDEA
Ensure you have JDK 8 (or newer) and Git installed

    java -version
    git --version

First clone the repository:

    git clone https://github.com/manabie-com/togo.git
    
Link Gradle Project

    Open this project on Intellij IDEA
    Right click on build.gradle file and select Link Gradle Project

Setup Run/Debug Configurations

    add a new configuration : Spring boot
    select main class is com/todo/TodoApplication.java
    
Run

## A sample “curl” command to call your API

Register

    curl --location --request POST 'http://localhost:8080/register' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "username": "aaaa@aaa.aa",
      "password": "a",
      "taskLimit": 2
    }'
    
Login to get token
    
    curl --location --request POST 'http://localhost:8080/authenticate' \
    --header 'Content-Type: application/json' \
    --data-raw '{
      "username": "aaaa@aaa.aa",
      "password": "a"
    }'
    
List all task of current user
    
    curl --location --request GET 'http://localhost:8080/tasks' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer <token>'
    
Accepts a todo task and records it
    
    curl --location --request POST 'http://localhost:8080/tasks' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer <token>' \
    --data-raw '{
      "content": "Test task !!!!!! "
    }'
