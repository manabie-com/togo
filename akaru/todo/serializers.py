from django.contrib.auth.models import User
from rest_framework import serializers

from todo.models import UserProfile
from todo.models import TodoTask


class TodoTaskSerializer(serializers.ModelSerializer):
    class Meta:
        model = TodoTask
        fields = '__all__'