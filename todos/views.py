import datetime
from rest_framework import generics, status
from rest_framework.parsers import FormParser, MultiPartParser, JSONParser
from rest_framework.permissions import AllowAny
from rest_framework.response import Response
from rest_framework.views import APIView

from todos.models import Task, SettingTask
from todos.serializers import TaskSerializer, SettingTaskSerializer


class TaskSetting(APIView):
    parser_classes = [FormParser, MultiPartParser, JSONParser]

    def post(self, request):
        serializer = SettingTaskSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save(user=self.request.user)

            return Response(serializer.data, status=status.HTTP_200_OK)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


class TaskView(APIView):

    parser_classes = [FormParser, MultiPartParser, JSONParser]

    def get(self, request, format=None):
        today = datetime.date.today()
        setting_task, created = SettingTask.objects.get_or_create(date=today, user=self.request.user)
        tasks = Task.objects.filter(setting=setting_task)
        serializer = TaskSerializer(tasks, many=True)
        return Response(serializer.data, status=status.HTTP_200_OK)

    def post(self, request):
        today = datetime.date.today()
        setting_task, created = SettingTask.objects.get_or_create(date=today, user=self.request.user)
        serializer = TaskSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save(setting=setting_task)

            return Response(serializer.data, status=status.HTTP_200_OK)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


class TaskDetailView(generics.RetrieveUpdateAPIView):
    queryset = Task.objects.all()
    serializer_class = TaskSerializer
