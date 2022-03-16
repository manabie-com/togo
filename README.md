
### How to run the code locally:
- PHP 7.3 ~ 8.0
- Composer 1.10.13+
- MySQL or MariaDB
- After you cloned the source code from github, go to .env file and edit the following codes:
	DB_CONNECTION=mysql
	DB_HOST=127.0.0.1
	DB_PORT=3306
	DB_DATABASE=your-database-name
	DB_USERNAME=your-database-username
	DB_PASSWORD=your-database-password
- Run the following commands from the terminal:
	- php artisan migrate
	- php artisan passport:install
- To run the application, run this command from the terminal : php artisan serve

### Sample curl command to call api:
- To register: curl -X POST http://127.0.0.1:8000/api/v1/register -d 'name=yourname&email=youremail@gmail.com&password=yourpassword&password_confirmation=yourpassword'

- To login: curl -X POST http://127.0.0.1:8000/api/v1/login -d 'email=youremail@gmail.com&password=yourpassword'

- To add todo task: curl -H "Authorization: Bearer your_access_token" -X POST http://127.0.0.1:8000/api/v1/todo -d 'content=your_content'

### How to run the unit tests locally:
- Run this command from the terminal and watch the result: php artisan test

### How to run the unit tests locally:
- I seperated the limit of task per user per day for easier developing
- The repository modules all share the same base class, so that all the basic query functions only need to write once

### Further development:
- Have more api to edit the limit of task
- Seperate the validate codes to its own module

