## ToDo API Endpoint
This is an API endpoint that accepts todo tasks from users. A user can have a daily limit of tasks to input.

## Running the code locally
### With Python Virtual Environment
To run the code locally, refer to the following steps:
  1. Make sure that Python v3.6.3 is properly installed in the machine - [Python Installation Tutorial](https://www.tutorialspoint.com/how-to-install-python-in-windows)
  2. Download a copy of the source code and open the togo/togo/ directory in Command Prompt
```sh
cd {download_path}/togo/togo
```
  3. Set up a virtual environment (2 options):
```sh
pip install virtualenv                                      # installs the virtualenv dependency
python -m venv togo-test                                    # using python virtualenv
togo-test\Scripts\activate.bat                              # activates the virtual environment

conda create --name togo-test python=3.6.3                  # using anaconda
conda activate togo-test                                    # activates the virtual environment
```
  4. Install the necessary packages
```sh
pip install -r requirements.txt
```
  5. Run the following command to run the localhost server
```sh
python manage.py runserver
```

### With Docker
  - Docker repository: https://hub.docker.com/repository/docker/doremikaelatido/todo/general
  1. Pull the todo remote repository
```
docker pull doremikaelatido/todo
```
  2. Confirm successful pull using another command, "doremikaelatido/todo" must be listed down
```
docker images
```
  3. Make sure docker account is logged in, if not yet set up - provide the credentials
```
docker login
```
  4. Run the code with Docker
```
docker run -d -p 8000:8000 doremikaelatido/todo
```

At this point, http://localhost:8000/ should display the ApiRoot page

## Sample cURL command
  - curl -X POST -H "Authorization: Api-Key cvQYsddc.PZVUK5AY3vftSerjzbwqz2qgsoNdjB6h" -H "Content-Type: application/json" -H "Username: choerry" -d "{``\``"title``\``":``\``"cook``\``"}" http://localhost:8000/usertasks/
  - curl -X POST -H "Authorization: Api-Key xn58ox4d.vVIexLi9I5q2F5k6HythNqyNrIprB99X" -H "Content-Type: application/json" -H "Username: oliviahye" -d "{``\``"title``\``":``\``"animal crossing``\``"}" http://localhost:8000/usertasks/

##### The following information can be added to the data:
  - **title** - title of the task
  - **description** - additional information about the task (optional)
  - **start_time** (%Y-%m-%d %H:%M:%S) - when the task is expected to be started (optional)
  - **end_time** (%Y-%m-%d %H:%M:%S) - when the task is expected to be completed (optional)

## Unit Testing
In the same directory (from the steps to run the code locally (/togo/togo) OR through Docker container's CLI:
Run the command 
```sh
python manage.py test
```


## Additional Remarks
### What I love about my solution
- Architecture wise, what I love about my solution is I think I used the correct framework to write this solution in. I only started learning Django on my own a few months back but I strived to understand what the framework would be best used for. While I was working on this solution, I was able to explore Django even more and I found it very convenient to use. In addition to that, Python and Django are very widely used, so I think this solution will be easy to work with for other developers. On top of that, I had to do a lot of refreshing on best practices while I was writing this solution, so I think that makes me more content with the work I put out.
- Design wise, I thought a lot about the overall design of this solution. Since the requirements don't go into specifics, I had a lot of things to consider if I should include them in the scope of this solution or not. I decided to add some of these things because I thought about how this API could be expanded (possibly from requirements that clients may ask for in the future).
    - e.g. For a planner/calendar application, they might need the users' historical todo tasks - so I use an active flag instead of deleting them, they might also need the start and end times for a calendar visualization - so I added the additional fields
- I also like the implementation of "task expiration" after 24 hours (instead of midnight - the actual end of day) since I thought about people who have different body clocks. If a person wakes up around lunch time, they will only have the remaining 12 hours left to keep these todo tasks. I think this implementation would work better depending on situations.
### What I did not have enough time to complete?
- For the functionality of task deletion, I really had to think about how I should implement it. The requirement only specified the endpoint for **creating tasks** but I thought about adding the **deleting tasks** function as well to test the limitation for daily tasks per user. If I were to be honest, I do like my design for task expiration (after 24 hours) but things would be a lot more simple if the "refresh" happens every midnight (since it's possible to just run a one time job that deactivates the tasks that were expected to be accomplished by that day). I think both approaches have its pros and cons so I would appreciate it a lot if I could get a feedback on this.
- I think the logging and error handling could be improved by a lot. What's important is the basic functionality but I would definitely work on these two more with more time. Perhaps for logging - more information on the request being processed (using some of the json data in the logs) and for error handling - to be able to handle the more particular errors.
- I would also work on the validation for start and end time. For now it accepts requests with start_time and no end_time (just sets a default end_time to 1 hour past the start_time) but no validation for requests that go the other way around (with end_time and no start_time)