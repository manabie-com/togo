from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.permissions import (
    AllowAny,
    IsAuthenticated,
)

from services.user_service import UserService


class BaseUserAPIView(APIView):
    def __init__(self, **kwargs) -> None:
        super().__init__(**kwargs)
        self.user_service = UserService()


class RegistrationAPIView(BaseUserAPIView):
    permission_classes = (AllowAny,)

    def post(self, request):
        data, http_code = self.user_service.create(request=request)
        return Response(data, status=http_code)


class UserListAPIView(BaseUserAPIView):
    permission_classes = (IsAuthenticated,)

    def get(self, request):
        data, http_code = self.user_service.get(request=request)
        return Response(data, status=http_code)


class UserDetailAPIView(BaseUserAPIView):
    permission_classes = (IsAuthenticated,)

    def get(self, request, id):
        data, http_code = self.user_service.get_single_user(user_id=id)
        return Response(data, status=http_code)

    def put(self, request, id):
        data, http_code = self.user_service.update(request=request, user_id=id)
        return Response(data, status=http_code)
