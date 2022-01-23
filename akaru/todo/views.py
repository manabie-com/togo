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
    serializer_class = TodoTaskSerializer

    def post(self, request, *args, **kwargs):
        serializer = TodoTaskSerializer(data=request.data)
        if serializer.is_valid():
            profile_id = request.data.get('user')
            profile = UserProfile.objects.get(id=profile_id)
            today = timezone.now().replace(
                hour=0, minute=0, second=0, microsecond=0)
            num_tasks_today = TodoTask.objects.filter(
                user=profile_id, date_created__gt=today).count()

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