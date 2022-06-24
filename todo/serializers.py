from rest_framework import serializers
from .models import Task

from utils import encrypting


class TaskSerializer(serializers.ModelSerializer):
    def to_representation(self, instance):
        data = super().to_representation(instance)
        data["id"] = encrypting.encrypt(data["id"])
        return data

    class Meta:
        model = Task
        exclude = ("created_by",)
