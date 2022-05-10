from django.shortcuts import render
from django.db.models import Count
from django.http import HttpResponse

from rest_framework import viewsets

from .serializers import TaskSerializer, UserSerializer, UserTaskSerializer
from .models import User, UserTask, Task
from .util import dbutil

class UserViewSet(viewsets.ModelViewSet):
    queryset = User.objects.all()
    serializer_class = UserSerializer

class UserTaskViewSet(viewsets.ModelViewSet):
    queryset = UserTask.objects.all()
    serializer_class = UserTaskSerializer

    # create - function executed for POST http request
    # validate the username provided to which the task will be associated with
    # proceed to creating the task if valid username
    def create(self, request):
        self.destroy()

        try: 
            username = request.META.get('HTTP_USERNAME')
            user = dbutil.user(username)
            if not user:
                return HttpResponse('Username does not exist', status=401)
            return dbutil.CreateUtil.createTaskRecord(user, request)

        except Exception as e:
            print(e)
            return HttpResponse('Missing a username header?', status=401)

    # destroy - function executed for DELETE http request
    def destroy(self):
        dbutil.DeleteUtil.deleteExpiredTasks()

class TaskViewSet(viewsets.ModelViewSet):
    queryset = Task.objects.all()
    serializer_class = TaskSerializer





