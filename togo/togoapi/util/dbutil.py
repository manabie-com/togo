from turtle import update
from ..models import User, UserTask, Task
from django.http import HttpResponse
from datetime import datetime, timedelta

from ..util import detailvalidationutil
from django.db import DatabaseError, transaction

# returns the user object based on username
def user(username):
    return User.objects.filter(username=username)[0]

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
        UserTask.objects.filter(added_time__lte=datetime.now()-timedelta(days=1)).update(is_active=False)
    
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
    def createTaskRecord(user, request):
        if user.daily_limit > user.task_today: 
            CreateUtil.createNewTask(user.id, request)        
            return HttpResponse("Task successfully created...", status=201)
        
        return HttpResponse("Daily task limit has been reached!!", status=409)
    
    # schedule (start and end time) validation
    # proceed if the schedule works 
    # setup up a transaction for inserting new records to Task and UserTask tables
    @staticmethod
    def createNewTask(user_id, request):
        if detailvalidationutil.validSchedule(request.data.get("start_time"), request.data.get("end_time")):
            try:
                with transaction.atomic():
                    ut_id = CreateUtil.createdUserTaskId(user_id)
                    CreateUtil.createTask(ut_id, request)
            except DatabaseError as e:
                print(e)
                return HttpResponse("Could not create a new task...", status=404)

        else:
            return HttpResponse("Invalid schedule set...", status=409)
    
    # insert a user task from data provided and return the newly created task id
    @staticmethod
    def createdUserTaskId(user_id):
        ut = UserTask(
            user_id = user_id,
            added_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
            is_active=True
        )
        ut.save()
        return ut.id

    # insert a task from data provided
    @staticmethod
    def createTask(task_id, request):
        t = Task(
            task_id = task_id,
            title = request.data.get("title"),
            description = request.data.get("description"),
            start_time = request.data.get("start_time"),
            end_time=request.data.get("end_time")
        )
        t.save()