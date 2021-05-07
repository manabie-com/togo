# Interview Restful API 
This is an application built with gin-gonic, jwt, gorm, logrus, postgresql, redis, docker, docker-compose. 
As Backend requirements: "This project include many issues from code to DB strucutre, feel free to optimize them". 
I have changed a lot of things. My structure base on awesome project lists using Gin web framework: https://github.com/photoprism/photoprism
##### New changed
 - I have chosen new structure source code and used Gin-gonic framework.
 - Changed SQLite to PostgreSQL and Gorm.
 - Improved testing by Postman with Swagger UI and UnitTest
 - Support to run source code with docker and docker-compose and wrote some script for starting, clean up docker images.
 - Improved JWT algorithm: HS256 -> RS256. Then, add 1 more step for hashing token(I chose AES encrypted and decrypted). Thus, the hacker can't decrypt and try to read the body of JWT. 
 - Provided Redis for caching. If the daily limit is reached, we can reject quickly and reduce query records from DB. 
 - Provided Logrus. Log anything into log file, separate by day.  
 - I am aware that we can use Interfaces and mock database calls, but I decided to be replicate the exact same process that happens in real life for testing.
   We can define a test database in our .env so, everyone can create a new database for the tests alone.
##### Improvement later
 - Try to support CICD and deploy source code with no downtime: kubernetes, docker swarm,CircleCI,...
 - Hide swagger UI (run as production mode)
##### Run project
 - With docker:
   + [Docker](https://www.docker.com/)  installed (**version 17.12.0+ is required**)
   + [Docker compose](https://docs.docker.com/compose/install/) installed (**version 1.21.0+ is required**)
   + Please see docker-compose.yml and change .env file. 
   + For your convenience, you can copy all environment variables from .env.example.docker and place it into your .env
   + Start app by command: ./dockerDeployLocal.sh
   + The swagger UI will be show at: http://127.0.0.1:9005/swagger/index.html . However, you need to migrate the database before you can test APIs 
   + Log in into pgadmin http://127.0.0.1:5050/ and create a new server connect to our db( see info at your .env) -> copy contents **migrations/1.0.0_migration_init.sql** and execute it. 
   + Try check healthy, server information, register a new user, log in, add a task.
 - Without docker:
   + Install Go: go version go1.15.5 linux/amd64(**version 1.15+ is required**)
   + Set up your GOPATH
   + In your $GOPATH, run command: mkdir -p go/{bin,src,pkg}
   + Run command: cd go/src
   + Run command: mkdir manabie-com
   + Clone satoshi-game source code into **manabie-com** (**Required: Don’t change this name. It’s should be manabie-com**)
   + Change environment variables of **.env** file
   + Build and run main.go
   + Then, migrate your database and testing.
##### Run tests
   + Create a new test db and change your **.env** file:
      <pre>
        TEST_DATABASE_HOST=
        TEST_DATABASE_PORT=
        TEST_DATABASE_USERNAME=
        TEST_DATABASE_PASSWORD=
        TEST_DATABASE_NAME=
        TEST_DATABASE_SSL_MODE=
        TEST_DATABASE_TIME_ZONE=
        TEST_DATABASE_QUERY_MAX_LIMIT=
      </pre>
   + Run test.