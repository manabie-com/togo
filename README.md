# Technical Assignment for Backend Engineer
!["golang"](https://miro.medium.com/max/3152/1*Ifpd_HtDiK9u6h68SZgNuA.png)

## How can I run the app ?
```
sudo docker-compose up
```
- API port `:8000`
- Database (PostgreSQL) port `:35432`

For test
```
go test -v ./test/...
```
### What is missing ?
- Integration tests: the time is really limited to me for adding this kind of test. I am currently having the seminar exams so that I cannot research and do my best in this testing.

## Changes I believe that would useful
### Software Architecture: 
___
I closely flow the `Clean Architecture` of Uncle Bob 
1. The structure (in my view):
```
.
├── api
│   └── http
│       ├── v1
├── controller
├── domain
├── entity
├── infra
│   ├── repository
├── schema
```
- `api` present for web server as a third party framework, ... My point to seperate this layer framework to the domain business code
- `controller` as the routing. It's responsiblity would be like the transport layer that comunicate between web framework and our application
- `domain` is our core business code, it just like a stand-alone module which isolate with every outer layer includes the database
- `repository` is our data layer that comunicate between our application and the database
- `enitty` is data structure for our model in database. It would only used in repository and called from domain
- `schema` is reuqest-response structure presentation. It's useful for checking validate the request before go through the application depper

Please feel free to check the source and improve the source code with feedbacks

**Diagram about the architecture would be upload soon**

## Things I want to improve
1. Global Error handle for whole app. It will catch the errors and check them in only one place. It will avoid us from repeating code and we don't need to check the errors in ervery layer

2. Research deeper about testing for both unit test and integration test in golang
3. More clean-code and restructure some layers
4. Dealing with scaling problems and benmarcking \
...\
and more other stuffs