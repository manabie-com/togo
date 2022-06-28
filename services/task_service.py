from typing import OrderedDict
from django.http import response
from rest_framework import status
from todo.serializers import TaskSerializer
from todo.models import Task

from utils import encrypting, response_formatting
from django.utils import timezone
from datetime import timedelta

from services.user_service import UserService
from constants import HTTPReponseMessage


class TaskService:
    def __init__(self) -> None:
        self.user_service = UserService()

    def get(self, request, is_list=False, task_id=None) -> tuple:
        user = self.__get_user_from_token(request)
        serializer = TaskSerializer(
            self.__get_single_or_list(user["user_id"], is_list, task_id),
            many=is_list,
        )

        return serializer.data, status.HTTP_200_OK

    def create(self, request) -> tuple:
        user = self.__get_user_from_token(request)

        if not self.__can_be_added(user["user_id"]):
            return response_formatting.get_format_message(
                HTTPReponseMessage.EXCEED_MAXIMUM_ERROR, status.HTTP_400_BAD_REQUEST
            )

        data = OrderedDict()
        data.update(request.data)
        data["created_by"] = user["user_id"]
        serializer = TaskSerializer(data=data)

        if serializer.is_valid():
            serializer.save()
            return serializer.data, status.HTTP_201_CREATED

        return serializer.errors, status.HTTP_400_BAD_REQUEST

    def update(self, request, task_id) -> tuple:
        user = self.__get_user_from_token(request)
        task = self.__get_task_by_id(task_id, user["user_id"])
        serializer = TaskSerializer(task, data=request.data, partial=True)

        if serializer.is_valid():
            serializer.save()
            return serializer.data, status.HTTP_200_OK

        return serializer.errors, status.HTTP_400_BAD_REQUEST

    def delete(self, request, task_id) -> tuple:
        user = self.__get_user_from_token(request)
        task = self.__get_task_by_id(task_id, user["user_id"])

        task.delete()

        return response_formatting.get_format_message(
            HTTPReponseMessage.DELETE_SUCCESSFULL, status.HTTP_204_NO_CONTENT
        )

    def __get_user_from_token(self, request):
        return self.user_service.get_user_from_access_token(request)

    def __get_number_of_added_tasks(self, user_id):
        today = timezone.now()
        tomorrow = today + timedelta(days=1)
        yesterday = today - timedelta(days=1)

        return Task.objects.filter(
            created_by=user_id, created_at__gt=yesterday, created_at__lt=tomorrow
        ).count()

    def __can_be_added(self, user_id) -> bool:
        user = self.user_service.get_user_object(user_id)
        number_of_added_tasks = self.__get_number_of_added_tasks(user_id)
        maximum_task_of_current_user = user.maximum_task_per_day

        if number_of_added_tasks < maximum_task_of_current_user:
            return True
        return False

    def __decrypt_data(self, id):
        return encrypting.decrypt(id)[0]

    def __get_task_by_id(self, task_id, user_id):
        try:
            return Task.objects.get(pk=self.__decrypt_data(task_id), created_by=user_id)
        except Task.DoesNotExist:
            raise response.Http404

    def __get_single_or_list(self, user_id, is_list, task_id):
        if is_list is True:
            return Task.objects.filter(created_by=user_id)
        return self.__get_task_by_id(task_id, user_id)
