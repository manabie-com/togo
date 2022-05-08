from django.shortcuts import render
from django.db.models import Count
from django.http import HttpResponse

from rest_framework import viewsets

from .serializers import TaskSerializer, UserSerializer, UserTaskSerializer
from .models import User, UserTask, Task

from datetime import datetime, timedelta

from .util import authutil

class UserViewSet(viewsets.ModelViewSet):
    queryset = User.objects.all()
    serializer_class = UserSerializer

class UserTaskViewSet(viewsets.ModelViewSet):
    queryset = UserTask.objects.all()
    serializer_class = UserTaskSerializer

    def create(self, request):
        try: 
            username = request.META.get('HTTP_USERNAME')            
            if not authutil.usernameExists(username):
                return HttpResponse('Username does not exist', status=401)
        except:
            return HttpResponse('Missing a username header?', status=401)

        # insert a new usertask to UserTask
        # insert a new task to Task
        return HttpResponse(status=200)
            
    """
    What the DELETE method would do
    - set the active flag of a UserTask to False
        - no deletions - with consideration for archiving purposes
    - update the task_today of the users
    """

    def delete():
        timenow = datetime.datetime.now()
        
        expired = UserTask.objects.filter(added_time=(timenow-timedelta(hours=24), timenow))
        expired.update(is_active=False)

        # update the task_today values for all users

class TaskViewSet(viewsets.ModelViewSet):
    queryset = Task.objects.all()
    serializer_class = TaskSerializer

    def create(task):
        pass



