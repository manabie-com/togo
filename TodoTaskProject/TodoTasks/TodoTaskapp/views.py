from django.shortcuts import render
from rest_framework import viewsets
from rest_framework import views
from rest_framework.response import Response
from rest_framework import status
from rest_framework.authentication import SessionAuthentication, BasicAuthentication
from rest_framework.permissions import IsAuthenticated

from .models import UserRecordTask, UserTaskAllow
from .serializers import UserRecordTaskSerializer, UserTaskAllowSerializer

#from datetime import datetime
import datetime

class UserRecordTasks(views.APIView):
    #authentication_classes = [SessionAuthentication, BasicAuthentication]
    #permission_classes = [IsAuthenticated]
    def get(self, request):
        _UserRecordTask = UserRecordTask.objects.all()
        serializer = UserRecordTaskSerializer(_UserRecordTask, many=True)
        return Response(serializer.data)

    def post(self, request):
        serializer = UserRecordTaskSerializer(data=request.data)
        if serializer.is_valid():
            _UserTaskAllow = UserTaskAllow.objects.filter(user=request.user).first()
            aday_in_seconds = 86400
            if (datetime.datetime.now().replace(tzinfo=datetime.timezone.utc) - _UserTaskAllow.start_task_time).total_seconds() > aday_in_seconds:
                #-------its been over a day!
                _UserTaskAllow.task_done = 0
                _UserTaskAllow.save(update_fields=["task_done"])
                #----else: its been not over a day!
            if _UserTaskAllow.task_allow > _UserTaskAllow.task_done:
                serializer.save(user=request.user)
                _UserTaskAllow.task_done += 1
                _UserTaskAllow.save(update_fields=["task_done"])
            else:
                return Response({"Warning": "You have expired today!"}, status=status.HTTP_400_BAD_REQUEST)
            return Response(serializer.data, status=status.HTTP_201_CREATED)

        return Response(serializer.data, status=status.HTTP_400_BAD_REQUEST)

class UserRecordTaskDetail(views.APIView):
    authentication_classes = [SessionAuthentication, BasicAuthentication]
    permission_classes = [IsAuthenticated]
    def get_object(self, pk):
        try:
            _UserRecordTask = UserRecordTask.objects.get(pk=pk)
            return _UserRecordTask
        except:
            return Response(status=status.HTTP_404_NOT_FOUND)

    def get(self, request, pk):
        _UserRecordTask = self.get_object(pk=pk)
        serialzer = UserRecordTaskSerializer(_UserRecordTask)
        return Response(serialzer.data, status=status.HTTP_200_OK)

    def put(self, request, pk):
        try:
            _UserRecordTask = self.get_object(pk=pk)
        except:
            return Response(status=status.HTTP_404_NOT_FOUND)
        serialzer = UserRecordTaskSerializer(instance=_UserRecordTask, data=request.data)
        if serialzer.is_valid():
            serialzer.save()
            return Response(serialzer.data, status=status.HTTP_202_ACCEPTED)
        return Response(serialzer.errors, status=status.HTTP_400_BAD_REQUEST)

    def delete(self, request, pk):
        try:
            _UserRecordTask = self.get_object(pk=pk)
        except:
            return Response(status=status.HTTP_404_NOT_FOUND)        
        _UserRecordTask.delete()
        return Response(status=status.HTTP_204_NO_CONTENT)

class UserTaskAllows(views.APIView):
    authentication_classes = [SessionAuthentication, BasicAuthentication]
    permission_classes = [IsAuthenticated]
    def get(self, request):
        _UserTaskAllow = UserTaskAllow.objects.all()
        serializer = UserTaskAllowSerializer(_UserTaskAllow, many=True)
        return Response(serializer.data, status=status.HTTP_200_OK)

    def post(self, request):
        serializer = UserTaskAllowSerializer(data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

class UserTaskAllow_Detail(views.APIView):
    authentication_classes = [SessionAuthentication, BasicAuthentication]
    permission_classes = [IsAuthenticated]
    def get_object(self, pk):
        try:
            #_UserTaskAllow=UserTaskAllow.objects.filter(user=user)
            return UserTaskAllow.objects.get(pk=pk)
        except:
            return Response(status=status.HTTP_404_NOT_FOUND)

    def get(self, request, pk):
        _UserTaskAllow = self.get_object(pk=pk)
        serializer = UserTaskAllowSerializer(_UserTaskAllow)
        return Response(serializer.data, status=status.HTTP_200_OK)

    def put(self, request, pk):
        _UserTaskAllow = self.get_object(pk=pk)
        serializer = UserTaskAllowSerializer(instance=_UserTaskAllow, data=request.data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_202_ACCEPTED)
        return Response(serializer.errors, status=status.HTTP_406_NOT_ACCEPTABLE)

    def delete(self, request, pk):
        _UserTaskAllow = self.get_object(pk=pk)
        _UserTaskAllow.delete()
        return Response(status=status.HTTP_204_NO_CONTENT)
