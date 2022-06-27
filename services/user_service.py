from django.http import response
from rest_framework import status
from rest_framework_simplejwt.tokens import AccessToken

from utils import encrypting, response_formatting

from baseuser.serializers import UserSerializer
from baseuser.models import BaseUser
from constants import HTTPReponseMessage


class UserService:
    def get_user_from_access_token(self, request):
        content = {"auth": str(request.auth)}
        access_token = AccessToken(content["auth"])
        return access_token

    def get_user_object(self, user_id):
        return BaseUser.objects.get(pk=user_id)

    def get_single_user(self, request, user_id):
        current_user = self.get_user_from_access_token(request)
        is_super_user = self.__is_super_user(user=current_user)
        user_id = self.__decrypt_data(user_id)

        if is_super_user is True or current_user["user_id"] == user_id:
            user = BaseUser.objects.get(pk=user_id)
            serializer = UserSerializer(user)
            return serializer.data, status.HTTP_200_OK

        return response_formatting.get_format_message(
            HTTPReponseMessage.NOT_ALLOWED_VIEW, status.HTTP_403_FORBIDDEN
        )

    def get(self, request) -> tuple:
        user = self.get_user_from_access_token(request)
        is_super_user = self.__is_super_user(user=user)

        if not is_super_user:
            return response_formatting.get_format_message(
                HTTPReponseMessage.NOT_ALLOWED, status.HTTP_403_FORBIDDEN
            )

        users = BaseUser.objects.all()
        serializer = UserSerializer(users, many=True)

        return serializer.data, status.HTTP_200_OK

    def create(self, request) -> tuple:
        serializer = UserSerializer(data=request.data)

        if serializer.is_valid():
            serializer.save()
            return serializer.data, status.HTTP_201_CREATED

        return serializer.errors, status.HTTP_400_BAD_REQUEST

    def update(self, request, user_id) -> tuple:
        user = self.get_user_from_access_token(request)
        is_super_user = self.__is_super_user(user=user)

        if not is_super_user:
            return response_formatting.get_format_message(
                HTTPReponseMessage.NOT_ALLOWED, status.HTTP_403_FORBIDDEN
            )

        maximum_task_per_day_from_request = request.data.get(
            "maximum_task_per_day", None
        )

        if maximum_task_per_day_from_request is None:
            return response_formatting.get_format_message(
                HTTPReponseMessage.MISSING_FIELD, status.HTTP_400_BAD_REQUEST
            )

        updating_user = self.__get_user_by_id(user_id)
        if updating_user.maximum_task_per_day > maximum_task_per_day_from_request:
            return response_formatting.get_format_message(
                HTTPReponseMessage.INVALID_MAXIMUM_TASK_FIELD,
                status.HTTP_400_BAD_REQUEST,
            )

        serializer = UserSerializer(
            updating_user,
            data=request.data,
            partial=True,
        )

        if serializer.is_valid():
            serializer.save()
            return serializer.data, status.HTTP_200_OK

        return serializer.errors, status.HTTP_400_BAD_REQUEST

    def __is_super_user(self, user) -> bool:
        user_object = BaseUser.objects.get(pk=user["user_id"])
        return user_object.is_superuser

    def __get_user_by_id(self, user_id):
        try:
            return BaseUser.objects.get(pk=self.__decrypt_data(user_id))
        except BaseUser.DoesNotExist:
            raise response.Http404

    def __decrypt_data(self, id):
        return encrypting.decrypt(id)[0]
