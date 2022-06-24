from rest_framework import serializers
from .models import BaseUser

from utils import encrypting
from todo.serializers import TaskSerializer


class UserSerializer(serializers.ModelSerializer):
    tasks = TaskSerializer(many=True, read_only=True)

    def to_representation(self, instance):
        data = super().to_representation(instance)
        data["id"] = encrypting.encrypt(data["id"])
        return data

    class Meta:
        model = BaseUser
        fields = (
            "id",
            "username",
            "maximum_task_per_day",
            "tasks",
        )
        extra_kwargs = {"password": {"write_only": True}}

    def create(self, validated_data):
        password = validated_data.get("password", None)
        instance = self.Meta.model(**validated_data)
        if password is not None:
            instance.set_password(password)
        instance.save()
        return instance
