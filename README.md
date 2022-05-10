### How to run your code locally?
`go run main.go`

### A sample “curl” command to call your API

- register
```
curl --location --request POST 'localhost:8888/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "test_user",
    "password": "passwordtest",
    "task_limit_per_day": 5
}'
```

- login
```
curl --location --request POST 'localhost:8888/login' \
--header 'Content-Type: application/json' \
--data-raw '{
     "username": "test_user",
    "password": "passwordtest"
}'
```

- add task
```
curl --location --request POST 'localhost:8888/task' \
--header 'Authorization: Bearer <access_token>' \
--form 'title="sample task"' \
--form 'description="sample description"'
```

### How to run your unit tests locally?
`go test ./...`

### What do you love about your solution?
Being this as my first go rest api ever created. I love everything about it. I also appreciate your tip to check other candidates PR and utlizing the power of open source. it helps me to build this app piece by piece, for the project structure and design pattern, I follow this post (https://dev.to/stevensunflash/using-domain-driven-design-ddd-in-golang-3ee5) since I found this easier to follow through than other tutorials out there. though I'm still exploring other projects on golang so hit me up if you have other suggestion/recommendations for learning materials and best practices that I can use for my learning.

### What else do you want us to know about however you do not have enough time to complete?
Apparently, due to my work experience having used proprietary technologies and a limited experience in java. theres a lot of things that I need to learn like automation testing and test-driven development. as I see on your company profile, I know you value high quality and highly tested codes to implement good quality products so this is one thing that I need to learn and focus on. another one is the use of docker. I dont have experience using it so I havent implemented it on the project.