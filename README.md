## DaiTV

### Explain
In the first time, I thought the test only took a few hours. So i talk to the headhunter that i can do it in wednesday night.
But, this weekend, I'm have to go back to my hometown to join my brother' wedding. So, I don't have time to do this test in this weekend and i need to do it as fast as i can. Sorry about that.

Thank you!

### How to run this code locally
 - Install golang: https://go.dev/doc/install
 - Install mongodb: https://www.mongodb.com/docs/manual/installation/
 - Change `CONNECTION_STRING` in `be/env/Env.go` to your connection string.
 - Run commands bellow.
```sh
cd /root/project/directory
go run be/main.go
```

### Sample “curl” command to call API
Run this command to add todo task and get header value name `Token`
```sh
curl -H "Content-Type: application/json" -X POST -d '{"Text":"Todo taks text"}' -v http://localhost:8008/api/todo
```
Run this command to continue with the same user in command above.
```sh
curl -H "Content-Type: application/json" -H "Token:{Token header in command above}" -X POST -d '{"Text":"Todo taks text"}' -c - http://localhost:8008/api/todo
```

### How to run unit tests locally
 - run this code locally.
 - Run commands bellow.
```sh
cd /root/project/directory
go test ./...
```

### What do you love about your solution?
I think there is nothing special in my solution.

### What else do you want us to know about however you do not have enough time to complete?
 - Add more apis and code FrontEnd with VueJs to make it become a real todo list web app.
 - Create `Dockerfile`, `docker-composer.yml`.
 - Write test case more detail.