# To Do API

##TECH
Node, Redis, AWS, Serverless
    
## PATHS

endpoints:
  - POST - https://{URL}/dev/create-user
  - POST - https://{URL}/dev/get-user
  - POST - https://{URL}/dev/create-task
  - POST - https://{URL}/dev/delete-task
  - GET - https://{URL}/dev/reset-limit


  

## AUTH
Header
```sh
Authorization: {TOKEN}
```

## Installation

To test the hosted version, just visit the URL path.
To deploy/invoke local by yourself, require:
- AWS
- Node
- Serverless framework
- Fill serverless.yml file with your AWS redis host, security groups and subnet.
```
serverless deploy --stage dev --aws-profile {your profile}
```

To run test:
- Node
- Redis
```
npm run test
```

Visit ```https://redis.io/download``` for installation guide.


## RUN

Visit the URL path with header, either by postman or curl:
CURL
```
curl --location --request POST 'https://{URL}/dev/create-user' \
--header 'Authorization: {TOKEN}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "limit": 1
}'
```
Run test/ run locally:
```
redis-server
npm run test
```
## API
### create-user
description:
provide ```limit```, return ```id```, which is user id, use through out the application.
payload:

```
{
    "limit": 1
}
```

sample response:
```
{
    "message": "SUCCESS",
    "id": "1642091273404",
    "limit": 1
}
```

### get-user
description:
provide ```id```, return users details.
payload:

```
{
    "id": "1642091273404"
}
```

sample response:
```
{
    "user": {
        "limit": "1",
        "remain": "0"
    },
    "task": {
        "1642091298256_task": "task description here"
    }
}
```

### create-task
description:
provide ```userId```, with ```description```. Return ```taskId```. Can be called n times, where n is the limit of user.
payload:

```
{
    "userId": "1642091273404",
    "description": "tmp 1"
}
```

sample response:
```
{
    "message": "SUCCESS",
    "taskId": "1642093823430_task",
    "description": "tmp 1"
}
```
```
{
    "message": "Create task limit reached"
}
```

### delete-task
description:
provide ```userId```, ```taskId```. delete that task
payload:

```
{
    "userId": "1642067046395",
    "taskId": "1642067166317_task"
}
```

sample response:
```
{
    "message": "Success delete"
}
```

### reset-limit
description:
reset all user limit back to original


sample response:
```
{
    "message": "Reset limit successful"
}
```

## Some notes
```
What do you love about your solution?
```
I was gonna fire some classic ExpressJS and some DB, and then I remember serverless architecture is perfect for this type of task.

Serverless reduces the complicated of any backend, plus reduce the complication of integration test. It also minimize coupling, which is nice when code base grow and any changes can break others.

There are a few downside of serverless, some which related to this particular tasks:
    - To test the function you need to mock events for the aws lambda.
    - Require aws to run true locally. Can always mock event as an input object through.
    - Require some network knowledge.
    
However, I think serverless is much better than traditional method in this case.

I also like redis as the DB of choice. Simple and fast for quick task. It should't be used as main db in production through.

```
What else do you want us to know about however you do not have enough time to complete?
```
I can only spend like 1 night of work on this task, but I think it covers pretty much all requirements. 
Maybe an API observer like sentry, or add CICD to run auto test everytime we deploy code, would be emphasize testing even more.
