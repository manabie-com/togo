### Requirements
- Implement one single API which accepts a todo task and records it 
  - There is a maximum **limit of N tasks per user** that can be added **per day**. #DONE
  - Different users can have **different** maximum daily limit. #DONE
- Write integration (functional) tests #PENDING
- Write unit tests #DONE
- Choose a suitable architecture to make your code simple, organizable, and maintainable #JAVE SPRING BOOT + POSTGRESQL
- Write a concise README #DONE

  
  
### README

1. Run Code locally
1.1 With IDE
- Install IntelliJ (https://www.jetbrains.com/idea/download/?fromIDE=#section=windows).
- Install POSTGRESQL (https://www.postgresql.org/download/windows/ I choose ver12).
- Install JRE, setup Evironment variables.
- Enable spring.jpa.hibernate.ddl-auto=create in \src\main\resources\application.properties for auto generate tables
1.2 With Deployed Service
- cd deploy (deployed service (manage-task-service2.jar) & deploy tool, config are there)
- Run CMD: WinSW.NET4 install
- Open Window Service Manager -> Select our service -> Start.

2.Sample “curl” command to call API
->register new user with default limit task = 3
curl -d "uid=hoavq&password=1" http://localhost:8080/api/user/register

->delete user
curl --request "DELETE"  http://localhost:8080/api/user/delete/hoavq

->create task for him 4 times to test
curl -d "uid=hoavq&createDate=20/04/2022&description=coding1" -H "Content-Type:application/json" http://localhost:8080/api/task/register
curl -d "uid=hoavq&createDate=20/04/2022&description=coding2" -H "Content-Type:application/json" http://localhost:8080/api/task/register
curl -d "uid=hoavq&createDate=20/04/2022&description=coding3" -H "Content-Type:application/json" http://localhost:8080/api/task/register
curl -d "uid=hoavq&createDate=20/04/2022&description=coding4" -H "Content-Type:application/json" http://localhost:8080/api/task/register

->list current task
curl -v http://localhost:8080/api/task/list

3.Run unit test locally
I run with IDE only.

4.What do you love about your solution?
My solution is simple, organizable, maintainable.

5.What else do you want us to know about however you do not have enough time to complete?
The mail came to me at 10h00 last Sunday but unfortunately it came to my spam mail. When i checked mail, it was Sunday evenning, so I have 4 evennings to complete the test.
Honestly, this's my first time that i create a microservices & intergration test (I used to use unit test but with C++ && Google test before).
The daily works don't require me to create microservices, i just call api from other system service & we use SOAP protocol not REST API.
The requirements are really not too difficult, i think so. But there's a lot of new things to me, i need more time to complete intergration tests.
Anyways, i learn alot from this test.
Thank you











