from rest_framework import serializers
from .models import Todo, CustomUser
from rest_framework.validators import UniqueValidator
from django.core.validators import MaxLengthValidator
class TodoSerializer(serializers.ModelSerializer):
    class Meta:
        model = Todo
        fields = ["id", "task", "completed", "timestamp", "updated", "user"]

class UserSerializer(serializers.ModelSerializer):
    class Meta:
        model = CustomUser
        fields = ['id', 'username', 'todo_max']

class RegisterSerializer(serializers.ModelSerializer):
    username = serializers.CharField(
        validators=[UniqueValidator(queryset=CustomUser.objects.all()), MaxLengthValidator(150)]
    )
    password = serializers.CharField(min_length=4)

    class Meta:
        model = CustomUser
        fields = ['id', 'username', 'password', 'todo_max']
        extra_kwargs = {'password': {'write_only': True}}

    def create(self, validated_data):
        user = CustomUser.objects.create_user(validated_data['username'], validated_data['password'],
                                              todo_max=validated_data[
            'todo_max'])
        return user