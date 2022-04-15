### Run the code
 * Create a virtual environment *(optional)*
```
python3 -m venv ./.venv
source ./.venv/bin/activate
```
 * Install requirements
```
pip install -r requirements.txt
```
 * Run the server

Go to Django project's root directory
```
cd akaru
```
Run migrations
```
python manage.py makemigrations
python manage.py migrate
```
Start the server
```
python manage.py runserver
```
Create a superuser
```
python manage.py createsuperuser
```
Check out the admin panel
```
127.0.0.1:8000/admin
```
You need to create a `UserProfile` by setting the limit for your users in the `User` page for these to be able to create tasks.

### Try with `curl`
```
curl --request POST --header "Content-Type: application/json" \
	--data '{"title": "Hello World!", "text": "Welcome!", "user": 1}' \
	127.0.0.1:8000/api/todo
```
### Run the tests
Run using Django's test-execution framework
```
python manage.py test todo
```
### "What do you love about your solution?"
- Scalable implementation of `UserProfile` that "extends" the built-in base `User` model by Django through a simple `OneToOneField`, getting access to its features; no need to reinvent the wheel. The existence of a  `UserProfile` of a `User` can also be used to indicate permission that it can create tasks.
- Timezone-aware, a feature of using the Django framework.
- Daily limit implementation dependent only on counting `TodoTask` entries in the database created during that day.
Scaling this system to include the usual CRUD operations will yield a case where if a `TodoTask` is deleted, it will not be counted against the daily limit.
One fix would be to implement a "soft-delete": add a field to the `TodoTask` model to indicate whether the task is "deleted" or not. This can also be extended to implementing a Trash feature, and so on.
> ### Requirements
> 
>   
> 
> - Implement one single API which accepts a todo task and records it
> 
> - There is a maximum **limit of N tasks per user** that can be added **per day**.
> 
> - Different users can have **different** maximum daily limit.
> 
> - Write integration (functional) tests
> 
> - Write unit tests
> 
> - Choose a suitable architecture to make your code simple, organizable, and maintainable
> 
> - Write a concise README
> 
> - How to run your code locally?
> 
> - A sample “curl” command to call your API
> 
> - How to run your unit tests locally?
> 
> - What do you love about your solution?
> 
> - What else do you want us to know about however you do not have enough time to complete?
> 
>   
> 
> ### Notes
> 
>   
> 
> - We're using Golang at Manabie. **However**, we encourage you to use the programming language that you are most comfortable with because we
> want you to **shine** with all your skills and knowledge.
> 
>   
> 
> ### How to submit your solution?
> 
>   
> 
> - Fork this repo and show us your development progress via a PR
> 
>   
> 
> ### Interesting facts about Manabie
> 
>   
> 
> - Monthly there are about 2 million lines of code changes (inserted/updated/deleted) committed into our GitHub repositories. To
> avoid **regression bugs**, we write different kinds of **automated
> tests** (unit/integration (functionality)/end2end) as parts of the
> definition of done of our assigned tasks.
> 
> - We nurture the cultural values: **knowledge sharing** and **good communication**, therefore good written documents and readable,
> organizable, and maintainable code are in our blood when we build any
> features to grow our products.
> 
> - We have **collaborative** culture at Manabie. Feel free to ask trieu@manabie.com any questions. We are very happy to answer all of
> them.
> 
>   
> 
> Thank you for spending time to read and attempt our take-home
> assessment. We are looking forward to your submission.
