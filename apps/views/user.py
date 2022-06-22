from django.contrib.auth.models import User
from rest_framework.generics import CreateAPIView

from .utils.decorator import retry_get_limit_task
from ..exceptions.base import make_success_response
from ..exceptions.status_code import Code404, Code400
from ..serializers.user import DetailSerializer
from rest_framework.permissions import IsAuthenticated
from datetime import date, datetime
from apps.models.models.detail import Detail
from apps.models.models.task import Task
from apps.models.models.schedule import Schedule
from rest_framework import status
from rest_framework.response import Response


class CreateDetail(CreateAPIView):
    serializer_class = DetailSerializer
    queryset = Detail.objects.all()
    permission_classes = (IsAuthenticated,)

    @staticmethod
    @retry_get_limit_task
    def get_schedule_task(user, date_):
        return Schedule.objects.get(user=user, date=date_)

    def perform_create(self, serializer):
        try:
            data = serializer.initial_data
            date_ = datetime.strptime(data.get('date'), '%Y-%m-%d').date()
            task = Task.objects.filter(id=data.get('task'))
            user = User.objects.filter(id=data.get('user'))
            if date_ > date.today():
                raise Code400
            if not task.exists() or not user.exists():
                raise Code404
            return serializer.save(task=task.first(), user=user.first())
        except Exception as e:
            raise

    def create(self, request, *args, **kwargs):
        serializer = self.get_serializer(data=request.data)
        serializer.is_valid(raise_exception=True)
        self.perform_create(serializer)
        headers = self.get_success_headers(serializer.data)
        return make_success_response(data=serializer.data, status_code=status.HTTP_201_CREATED, headers=headers)
