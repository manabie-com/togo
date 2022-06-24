from django.contrib.auth import authenticate
from rest_framework import serializers
from .models import BaseUser


class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = BaseUser
        fields = (
            "id",
            "username",
            "password",
            "maximum_task_per_day",
        )
        extra_kwargs = {"password": {"write_only": True}}

    def create(self, validated_data):
        password = validated_data.get("password", None)
        instance = self.Meta.model(**validated_data)
        if password is not None:
            instance.set_password(password)
        instance.save()
        return instance
