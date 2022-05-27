### 1. How to run your code locally?
* build docker images: "docker-compose build"
* run docker compose: "docker-compose up -d"
* [optional] horizontal scaling API deployment with 3 replicas
  * "docker-compose up --scale expressjs=3 -d"

### 2. Sample “curl” command to call your API:
* get dummy user:
  * curl --location --request GET 'http://localhost/users
* add a todo:
  * curl --location --request POST 'http://localhost/users/<userId>/todos' --header 'Content-Type: application/json' --data-raw '{ "name": "test" }'
* get list todo of user:
  * curl --location --request GET 'http://localhost/users/<userId>/todos'

### 3. How to run your unit tests locally?
 running test in docker container
  * get docker container name of expressjs: "docker ps"
  * docker name eg: "togo_expressjs_1"
  * login to the container: "docker exec -it togo_expressjs_1 bash"
  * run unit test: "npm run test"


### 4. What do you love about your solution?
* simple
* readable
* maintainable
* horizontal scaling
  * load balancer with nginx to distribute requests to horizontal scaled
  * deal with high traffic


### 5. What else do you want us to know about however you do not have enough time to complete?
* clustering mongodb for adapting horizontal scaling API service
* API authentication
* 90% coverage