# How to run

1. Clone code
2. Go to folder `docker-php/code/todo-app`
3. Run:
    ```
    composer install
    ```
    Note: Install composer if you have not
4. Run:
    ```
    php artisan migrate
    ```
5. Run:
    ```
    php artisan jwt:secret 
    ```
6. Go to folder `docker-php` and follow the instructions in the file huong-dan.txt
   1. List Api
       - User:

         |Name | Url | Method | Params
         | ------ | ------ | ------ | -----|
         | Register | api/users/register | POST | username, password
         | Login | api/auth/login | POST | username, password
         | Logout | api/auth/logout | POST |
         | Me | api/auth/me | GET|
       - Task:
      
         |Name | Url | Method | Params
         | ------ | ------ | ------ | -----|
         | Create | api/tasks | POST | content
         | Detail | api/tasks/:id | GET |
         | List | api/tasks | GET | page, limit
