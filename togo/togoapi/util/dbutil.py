from turtle import update
from ..models import User, UserTask, Task
from django.http import HttpResponse
from datetime import datetime, timedelta

from ..util import detailvalidationutil
from django.db import DatabaseError, transaction

from django.utils import timezone

time_format = "%Y-%m-%d %H:%M:%S"
td = "description"
tst = "start_time"
tet = "end_time"
tt = "title"

# returns the user object based on username
def user(username):
    try:
        return User.objects.filter(username=username)[0]
    except:
        raise Exception("This username does not exist...")

# returns the usertask object based on the task id
def usertask(id):
    return UserTask.object.filter(id=id)[0]

# new class for separation 
# delete operations
class DeleteUtil():
    # deactivates tasks added over 24 hours ago
    # updates the task count for the last 24 hours    
    @staticmethod
    def deleteExpiredTasks():
        try:
            with transaction.atomic():
                DeleteUtil.deactivateTasks()
                DeleteUtil.updateTaskCountToday()
        except DatabaseError as e:
            print(e)
            return HttpResponse("Could not delete expired tasks....", status=401)

    # sets active = False for user tasks added over 24 hours ago
    @staticmethod
    def deactivateTasks():
        UserTask.objects.filter(added_time__lte=datetime.now(timezone.utc)-timedelta(days=1)).update(is_active=False)
    
    # sets a user's "task_today" field to number of their active tasks
    @staticmethod
    def updateTaskCountToday():
        allusers = User.objects.all()
        for user in allusers:
            user.task_today = UserTask.objects.filter(user_id = user.id, is_active=True).count()
            user.save()

# new class for separation
# create operations
class CreateUtil():
    # daily task limit validation, proceed if user has not exceeded the limit
    @staticmethod
    def createTaskRecord(user, data):
        if not detailvalidationutil.dailyLimitReached(user): 
            CreateUtil.createNewTask(user, data)       
            return HttpResponse("Task successfully created...", status=201)
        
        return HttpResponse("Daily task limit has been reached!!", status=409)
    
    # schedule (start and end time) validation
    # proceed if the schedule works 
    # setup up a transaction for inserting new records to Task and UserTask tables
    @staticmethod
    def createNewTask(user, data):
        if detailvalidationutil.validSchedule(data.get("start_time"), data.get("end_time")):
            try:
                with transaction.atomic():
                    ut_id = CreateUtil.createdUserTaskId(user.id)
                    CreateUtil.createTask(ut_id, data)
                    CreateUtil.incrementUserDailyTask(user)
            except DatabaseError as e:
                print(e)
                return HttpResponse("Could not create a new task...", status=404)

        else:
            return HttpResponse("Invalid schedule set...", status=409)
    
    # insert a new user task using data provided
    # uses the current time as "added_time"
    # return the newly created task id
    @staticmethod
    def createdUserTaskId(user_id):
        ut = UserTask(
            user_id = user_id,
            added_time = datetime.now(timezone.utc).strftime(time_format),
            is_active=True
        )
        ut.save()
        return ut.id

    # insert a new task using data provided
    @staticmethod
    def createTask(task_id, data):
        t = Task(
            task_id = task_id,
            title = data.get(tt)
        )

        t = CreateUtil.addOptionalDetails(t, data)
        t.save()

    # checks for optional details such as description, start_time, end_time
    # and adds them records accordingly
    # if start time exists - defaults end time to one hour later
    @staticmethod
    def addOptionalDetails(task, data):
        if td in data:
            task.description = data.get(td)

        if tst in data:
            time_str = data.get(tst)
            task.start_time = time_str
            task.end_time = CreateUtil.defaultEndTime(time_str)

        if tet in data:
            task.end_time = data.get(tet)
        
        return task

    @staticmethod
    def defaultEndTime(start_time):
        return (datetime.strptime(start_time, time_format) + timedelta(hours=1)).strftime(time_format)

    # increments the user's current "daily task" amount
    @staticmethod
    def incrementUserDailyTask(user):
        user.task_today += 1
        user.save()