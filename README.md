### Setup environment
1. Install [Redis](https://redis.io/docs/getting-started/installation/)
2. Install [Python](https://www.python.org/downloads/) (required python 3.x)
3. Install [Mongodb](https://www.mongodb.com/docs/manual/installation/)
4. Install requirement package
  ```sh
    pip install pip requirements.txt
  ```
5. Set up database
  ```sh
    python database/set_up.py
  ```

### How to run
- Start server on the local:
  ```
    python core_dev.py
  ```
- You can test APIs using `curl`:
 1. First, you should create a user 
```cosole
curl -d "{\"username\":\"user\",\"password\":\"test\"}" -H "Content-Type: application/json" -X POST "http://localhost:5000/auth/signup"
```
 2. Login user
```cosole
curl -d "{\"username\":\"user\",\"password\":\"test\"}" -H "Content-Type: application/json" -X POST "http://localhost:5000/auth/signin"
```
 3. After logined, you will get a token as below
```json
{
    "data": {
        "bearer_token": "Bearer: eyJ0eXAiOiJKV1..."
    },
    "message": ""
}
```
 4. You have to add token for authentication when call API to do task
```cosole
curl -H "Content-Type: application/json" -H "accept: */*" -H "Authorization: Bearer: eyJ0eXAiOiJKV1..." -X POST "http://localhost:5000/task"
```
- Run the testcases:
```cosole
python -m pytest
```

### What do you love about your solution?
Base of development of REST API server with micro service architecture. Each service handles a small business.

I chose Flask as the library to handle our API requirement. Flask core is fairly small (comparing to Django) and extensible framework. I use (a bit) modified version on MVC in this project. Each blueprint (high level routes) come with a view (the URL) and a controller to call one or a series of asynchronous task(s) from services.

I chose Celery to handle multi processing requirement. For inter process communication we use Redis (PUB/SUB), Redis can also be used as great caching mechanism along the way.

This system is easily scalable and distributable, we just need an accessible Redis server to handle our communication, after that we can run unlimited number of service only instances of the app.

As mentioned before I am going with MVC, but I did not mentioned anything about the models. Models come with services (which we call their tasks from controllers).

The idea is that each service handles at most one model (several services can handle the same model). Well at least that's the "wish", in practice that is not always achievable as sometimes you need interim data between models and services which are not easily doable with celery functionalities. So the moral of the story is that, you can go with more than one model in a service if and only if you have no other way around it.

Celery provides us with many useful functionalities such as group, chain and chords (for more information please check the celery docs), using this capabilities we can solve most of data coupled problems.

Celery also provides us with 2 different type of processes, worker and beat. Workers can be used for sync/async calls to normal "working" tasks and beat time can be used to schedule recurring tasks.

### What else do you want us to know about however you do not have enough time to complete?
- Implement integration test. I tried for many hours to implement but failed. Having problems building test environment for celery and flask.
- Develop more features:
  - Auto update the limit of user by number request per day.
  - Reset counter request base on the time zone.
