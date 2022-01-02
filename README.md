# Description
Program is simple which help user can create / login, after that they can create their's task. They have limit tasks at their configuration

# How to run ?
  1. make run
  2. Waiting some seconds for mongo replicaSet init and find their mem nodes
  3. Run test with Postman URL
    . 3.1 create your user: unit by username, MUST CONFIG "max_tasks" for him
    . 3.2 login with username, password
    . 3.3 create task with response JWT token
    . 3.4 check the task create with jwt token
    
# How to test ?
  1. cd app && make test

# Postman URL
https://www.getpostman.com/collections/899b3865b7abcd5537f6

# Tech points:
1. Code based on Golang and mongoDB, apply MVC and repo dependency parttern
  . Apply transaction in mongo, can apply event driven's architect consistency instead for avoid bottleneck create by lock mechanism
  . Apply JWT authentication, can apply token session based for easy tracing history login, and eviction leak token

2. Testing
  . Unit test with business function, repository function (mongo)
  . Integration test with go_mock 
  . E2e test with postman API

3. Deploy
  . Deploy with docker-compose
    . mongo replicaSet: support transaction, changeStream (If apply event driven)
    . alpine golang image: apply multistage build docker for minimize size program (~ 10MB)

4. Apply cli at main.go:
  . support easy run, manage Deployment by command, apply in K8s deployment yml

# How can go far ?
1. Apply event-driven architect support for easily expend micro service system 
2. Deploy with K8s for easy scale and zero downtime => Increase Scalability, Availability, Reliability
3. Apply log EFK for tracing log from: Service, Gateway, tracing if need
4. Apply load manage with Grafana / prometheus
5. Add more integration test at Handler layer, use mockGen
