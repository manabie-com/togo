from django.shortcuts import render

# Create your views here.
from rest_framework import viewsets

from .serializers import TaskSerializer, UserSerializer, UserTaskSerializer
from .models import User, UserTask, Task

from datetime import datetime, timedelta

from django.db.models import Count

class UserViewSet(viewsets.ModelViewSet):
    queryset = User.objects.all()
    serializer_class = UserSerializer

class UserTaskViewSet(viewsets.ModelViewSet):
    queryset = UserTask.objects.all()
    serializer_class = UserTaskSerializer

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



