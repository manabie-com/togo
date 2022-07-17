from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework.permissions import IsAuthenticated
from rest_framework import status
from django.contrib.auth.models import User
from togo.models import UserProfile, Task
from togo.serializers import UserProfileSerializer, TaskSerializer


class UserView(APIView):
    """
    API endpoint that allows users to be viewed or edited.
    """
    # GET /api/users/
    # Retrieve list of users and their user profiles
    def get(self, request):
        response_dict = { "message": "User profiles successfully retrieved.", "users": UserProfileSerializer(UserProfile.objects.all(), many=True).data }
        return Response(response_dict, status=status.HTTP_200_OK)

    # POST /api/users/
    # Create new user (registration)
    def post(self, request):
        user = User.objects.create_user(**request.data)
        user_profile = UserProfile.objects.get(user=user)
        response_dict = { "message": "User successfully created.", "user": UserProfileSerializer(user_profile).data }
        return Response(response_dict, status=status.HTTP_201_CREATED)

class TaskView(APIView):
    """
    API endpoint that allows CRUD actions for tasks.
    """
    permission_classes = (IsAuthenticated,)

    # GET /api/tasks/
    # Retrieve tasks list for currently logged-in user
    def get(self, request):
        tasks = Task.objects.filter(user=request.user)
        response_dict = { "message": "Task list successfully retrieved.", "tasks": TaskSerializer(tasks, many=True).data }
        return Response(response_dict, status=status.HTTP_200_OK)

    # POST /api/tasks/
    # Create new task for currently logged-in user
    def post(self, request):
        # Check if user exceeded task limit
        if Task.objects.filter(user=request.user).count() >= UserProfile.objects.get(user=request.user).task_limit:
            return Response({ "message": "Task limit exceeded." }, status=status.HTTP_403_FORBIDDEN)
        # Create task if not
        else:
            task = Task.objects.create(user=request.user, **request.data)
            response_dict = { "message": "Task successfully created.", "task": TaskSerializer(task).data }
            return Response(response_dict, status=status.HTTP_201_CREATED)

    # PUT /api/tasks/<int:task_id>/
    # Update task with given ID
    def put(self, request, task_id=None):
        task = Task.objects.get(id=task_id)
        task.name = request.data["name"]
        task.save()
        response_dict = { "message": "Task successfully updated.", "task": TaskSerializer(task).data }
        return Response(response_dict, status=status.HTTP_200_OK)

    # DELETE /api/tasks/<int:task_id>/
    # Delete task with given ID
    def delete(self, request, task_id=None):
        task = Task.objects.get(id=task_id)
        task.delete()
        response_dict = { "message": "Task successfully deleted." }
        return Response(response_dict, status=status.HTTP_200_OK)