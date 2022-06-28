from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.permissions import IsAuthenticated

from services.task_service import TaskService


class BaseTaskAPIView(APIView):
    permission_classes = (IsAuthenticated,)

    def __init__(self, **kwargs):
        super().__init__(**kwargs)
        self.task_service = TaskService()


class TasksView(BaseTaskAPIView):
    def get(self, request):
        data, http_code = self.task_service.get(request, is_list=True, task_id=None)
        return Response(data, status=http_code)

    def post(self, request):
        data, http_code = self.task_service.create(request=request)
        return Response(data, status=http_code)


class TaskView(BaseTaskAPIView):
    def get(self, request, id):
        data, http_code = self.task_service.get(request=request, task_id=id)
        return Response(data, status=http_code)

    def put(self, request, id):
        data, http_code = self.task_service.update(request=request, task_id=id)
        return Response(data, status=http_code)

    def delete(self, request, id):
        message, http_code = self.task_service.delete(request=request, task_id=id)
        return Response(message, status=http_code)
