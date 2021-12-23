*** Notes ***
  - Architech: please have a look file architect.jpg

  - Set limit task per day: in file app.env
      LIMIT_TASK_PER_DAY 
  - Run server:
    docker-compose build
    docker-compose up
    *** Note ***
      at the first time, it maybe not run correctly, because database haven't already yet, please  ctr + c, then docker-compuse up again
  
  - Test:
    make test

*** Complete Featured ***
  - all features is completed(post, patch, get, delete) and match with front-end, please check front-end

    curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"user_id":"1","content":"content 1"}' \
    http://localhost:8080/tasks
  

  curl -v http://localhost:8080/tasks


*** Missing ***
I haven't enough time to finish all of test case, so i only write some test case for Controller, Service, Repository.
