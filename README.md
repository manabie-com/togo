### Overview

This is repository I am actively implementing to result the problems that I see from existing code forked from `master` branch

### Assignee
- Name: Hau Van Phuong
- Email: hvphuong98@gmail.com
## Note for code reviewer:
```
It would be great that you can leave your feedback or send me a notification after you have your decision.
```

### What I have done
1. I replaced the login method from get to post, because submiting sensitive information such as username, password, credit card on the url
param exposes a huge thread to the user and the system. The network package can be easily sniffed by hackers as piece of cake.
Using `POST` method resolves the problem as the network package with sensitive data will be encrypted at the `Application layer` by `HTTPS protocol`
which later send down to other layer, so even the package sniffed by hackers, cracking the package without private key that signed the SSL Certificate
will take them a while.
2. The code from master branch combines all logics for handling transport protocol (http), business logic into functions in `services` package, It make the bussiness logics tightly coupled with transport layer. So I decide to seperate it into `services` and `rest` package which aims to have the clear concern of each individual layer, cares about bussiness logic stuffs and handling rest api respectively. It embraces the `Seperate of concerns` theory, which is pretty suitable
for dependency injection.
3. Added middlewares for handling REST stuff likes method, authentication, logging...
4. I add mechanism to encrypt the password of user before saving into database (`bcrypt`).
5. I tried to do my best to organizing every components in its correct package. At the main function, I initialize all the dependencies (db connection, service, mux handler,...), then I injected the dependencies around.
6. I wrote unit-test with 70%+ test coverage for `service layer` (bussiness logic layer) (also improving)
7. I write a docker-compose file to run postgres database engine
8. I dumped the sqlite script and modified it to feed to postgres sql
### What I am missing
1. Integration test, I'm researching how to elegantly setup a integration test mechanism (mocking or database setup/teardown with docker and sql script)
2. Organizing a collections of `Postman` example to test the rest server (almost done)
3. Limiting task created by the `max_todo` field in database (5 is currently hard-coded, working on it)
4. ...
### Run the app
1. Start postgres engine by docker-compose (the default username/password is `phuonghau`/`phuonghau`)

```sh
docker-compose up
```
2. Seed the sample database (the password is `phuonghau`)
```
psql -U phuonghau -p 8899 -h localhost -f data_seed.sql
```
3. Start the server
```
go run main.go
```

4. Clean up
```
docker-compose down
```