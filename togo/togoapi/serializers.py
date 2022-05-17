from rest_framework import serializers

from .models import Task, User, UserTask

class TaskSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = Task
        fields = ('title', 'description', 'start_time', 'end_time')

class UserSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = User
        fields = ('id', 'username', 'daily_limit', 'task_today')

class UserTaskSerializer(serializers.HyperlinkedModelSerializer):
    class Meta:
        model = UserTask
        fields = ('id', 'user_id', 'added_time', 'is_active')