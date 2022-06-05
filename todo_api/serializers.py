from rest_framework import serializers
from rest_framework.validators import UniqueValidator
from django.core.validators import MaxLengthValidator
from .models import Todo, CustomUser, TodoList

class TodoSerializer(serializers.ModelSerializer):
    class Meta:
        model = Todo
        fields = [
            'title',
            'description',
            'status',
            'date',
            'priority',
            'tag'
        ]


class TodoListSerializer(serializers.ModelSerializer):
    class Meta:
        model = TodoList
        fields = [
            'id_user',
            'id_todo',
            'date'
        ]


class CustomUserSerializer(serializers.ModelSerializer):
    class Meta:
        model = CustomUser
        fields = [
            'username',
            'number_todo_limit'
        ]

class RegisterSerializer(serializers.ModelSerializer):
    username = serializers.CharField(
        validators=[UniqueValidator(queryset=CustomUser.objects.all()), MaxLengthValidator(50)]
    )
    password = serializers.CharField(min_length=4)

    class Meta:
        model = CustomUser
        fields = ['id', 'username', 'password', 'number_todo_limit']
        extra_kwargs = {'password': {'write_only': True}}

    def create(self, validated_data):
        if 'number_todo_limit' not in validated_data:
            data = {}
            data['number_todo_limit'] = CustomUser._meta.get_field(
                'number_todo_limit'
            ).get_default()
            validated_data.update(data)
        user = CustomUser.objects.create_user(validated_data['username'], validated_data['password'],
                                              number_todo_limit=validated_data[
            'number_todo_limit'])
        return user
