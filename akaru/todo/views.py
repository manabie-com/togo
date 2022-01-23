from asyncio import tasks
from django.contrib.auth.models import User
from django.utils import timezone

from rest_framework import generics
from rest_framework import status
from rest_framework.response import Response

from todo.models import TodoTask
from todo.models import UserProfile
from todo.serializers import TodoTaskSerializer

# Create your views here.
class TodoTaskCreateAPI(generics.CreateAPIView):
    """
    Implements a create-only endpoint that only provides 'post' method handler.
    """
    serializer_class = TodoTaskSerializer

    def post(self, request, *args, **kwargs):
        """
        Overwrite base method to check if task to be saved 
        will exceed the daily limit of the user
        """
        serializer = TodoTaskSerializer(data=request.data)
        if serializer.is_valid():
            # Get the UserProfile instance 
            # using the profile id from the request data
            profile_id = request.data.get('user')
            profile = UserProfile.objects.get(id=profile_id)

            # Get the timezone-aware date now and move time to exactly 00:00
            today = timezone.now().replace(
                hour=0, minute=0, second=0, microsecond=0)
            # Query the count of all tasks by the user created today
            num_tasks_today = TodoTask.objects.filter(
                user=profile_id, date_created__gt=today).count()

            # Check if the number of tasks created today 
            # does not exceed the user's limit  
            if profile.limit > num_tasks_today:
                serializer.save()
                return Response(
                    serializer.data, status=status.HTTP_201_CREATED)
            else:
                return Response(
                    {'error':'User reached task daily limit.'}, 
                    status=status.HTTP_429_TOO_MANY_REQUESTS)

        return Response(
            serializer.errors, status=status.HTTP_400_BAD_REQUEST)