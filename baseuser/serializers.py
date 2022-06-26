from rest_framework import serializers
from rest_framework.exceptions import ValidationError
from .models import BaseUser

from utils import encrypting
from constants import HTTPReponseMessage


class UserSerializer(serializers.ModelSerializer):
    def to_representation(self, instance):
        data = super().to_representation(instance)
        data["id"] = encrypting.encrypt(data["id"])
        return data

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

    def update(self, instance, validated_data):
        for field in validated_data:
            if not self.is_allow_updated(field):
                raise ValidationError(
                    HTTPReponseMessage.NOT_ALLOWED_UPDATE_FIELD % field
                )
        return super().update(instance, validated_data)

    def is_allow_updated(self, field_name):
        allowed_fields = ("maximum_task_per_day",)
        return field_name in allowed_fields
