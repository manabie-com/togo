from django.shortcuts import render

from .models import Todo, CustomUser, TodoList

from datetime import datetime, timezone, timedelta
import datetime
from rest_framework.views import APIView
from rest_framework.response import Response
from rest_framework import status
from rest_framework import permissions
from rest_framework import generics
# from .models import Todo, CustomUser
from todo_api.serializers import (
    RegisterSerializer,
    TodoSerializer,
    TodoListSerializer,
    CustomUserSerializer
)
# TodoSerializer, UserSerializer,

# Create your views here.
class UserDetailAPI(APIView):
    permission_classes = [permissions.IsAuthenticated]

    def get_object(self, user_id):
        '''
        Helper method to get the object with given todo_id, and user_id
        '''
        try:
            return CustomUser.objects.get(id=user_id)
        except CustomUser.DoesNotExist:
            return None

    # Retrieve
    def get(self, request, *args, **kwargs):
        user_instance = self.get_object(request.user.id)
        if not user_instance:
            return Response(
                {"res": "Object with user id does not exists"},
                status=status.HTTP_400_BAD_REQUEST
            )

        serializer = CustomUserSerializer(user_instance)
        return Response(serializer.data, status=status.HTTP_200_OK)

    # Update
    def put(self, request, *args, **kwargs):
        '''
        Updates the todo item with given todo_id if exists
        '''
        user_instance = self.get_object(request.user.id)
        if not user_instance:
            return Response(
                {"res": "Object with user id does not exists"},
                status=status.HTTP_400_BAD_REQUEST
            )
        data = {}
        if 'number_todo_limit' in request.data:
            data['number_todo_limit'] = request.data.get('number_todo_limit')

        serializer = CustomUserSerializer(instance = user_instance, data=data, partial = True)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_200_OK)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


class RegisterUserAPI(generics.GenericAPIView):
    serializer_class = RegisterSerializer

    def post(self, request, *args, **kwargs):
        serializer = self.get_serializer(data=request.data)
        serializer.is_valid(raise_exception=True)

        user = serializer.save()

        return Response({"user": CustomUserSerializer(user, context=self.get_serializer_context()).data})

class TodoListAPI(APIView):
    # add permission to check if user is authenticated
    permission_classes = [permissions.IsAuthenticated]
    serializer_class = TodoSerializer

    # 1. List all
    def get(self, request, *args, **kwargs):
        '''
        List all the todo items for given requested user
        '''
        todos = Todo.objects.filter(tag = request.user.id)
        serializer = TodoSerializer(todos, many=True)
        import pdb
        pdb.set_trace()
        return Response(serializer.data, status=status.HTTP_200_OK)

    # 2. Create
    def post(self, request, *args, **kwargs):
        '''
        Create the Todo with given todo data
        '''
        # import pdb
        today = datetime.date.today()
        todos = Todo.objects.filter(tag=request.user.id, date=today)
        # pdb.set_trace()
        if todos.count() >= request.user.number_todo_limit:

            return Response('Todo tasks per day for user: %s is limited by %d' % (request.user.username,
                                                                                  request.user.number_todo_limit),
                            status=status.HTTP_400_BAD_REQUEST)

        data = {}
        if 'title' in request.data:
            data['title'] = request.data.get('title')
        if 'description' in request.data:
            data['description'] = request.data.get('description')
        if 'date' in request.data:
            data['date'] = request.data.get('date')
        if 'status' in request.data:
            data['status'] = request.data.get('status')
        if 'priority' in request.data:
            data['priority'] = request.data.get('priority')
        if 'tag' in request.data:
            try:
                id_user = CustomUser.objects.get(username=request.data.get('tag')).id
                data['tag'] = id_user
            except:
                return Response(
                    ('%s is does not exists' % (request.data.get('tag'))),
                    status=status.HTTP_400_BAD_REQUEST
                )

        serializer = TodoSerializer(data=data)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_201_CREATED)

        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)


class TodoDetailApi(APIView):
    # add permission to check if user is authenticated
    permission_classes = [permissions.IsAuthenticated]

    def get_object(self, todo_id):
        '''
        Helper method to get the object with given todo_id, and user_id
        '''
        try:
            return Todo.objects.get(id=todo_id)
        except Todo.DoesNotExist:
            return None

    # Retrieve
    def get(self, request, todo_id, *args, **kwargs):
        '''
        Retrieves the Todo with given todo_id
        '''
        todo_instance = self.get_object(todo_id)
        if not todo_instance:
            return Response(
                {"res": "Object with todo id does not exists"},
                status=status.HTTP_400_BAD_REQUEST
            )

        serializer = TodoSerializer(todo_instance)
        return Response(serializer.data, status=status.HTTP_200_OK)

    # Update
    def put(self, request, todo_id, *args, **kwargs):
        '''
        Updates the todo item with given todo_id if exists
        '''
        todo_instance = self.get_object(todo_id)
        if not todo_instance:
            return Response(
                {"res": "Object with todo id does not exists"},
                status=status.HTTP_400_BAD_REQUEST
            )
        data = {}

        if 'title' in request.data:
            data['title'] = request.data.get('title')
        if 'description' in request.data:
            data['description'] = request.data.get('description')
        if 'date' in request.data:
            data['date'] = request.data.get('date')
        if 'status' in request.data:
            data['status'] = request.data.get('status')
        if 'priority' in request.data:
            data['priority'] = request.data.get('priority')
        if 'tag' in request.data:
            try:
                id_user = CustomUser.objects.get(username=request.data.get('tag')).id
                data['tag'] = id_user
            except:
                return Response(
                    ('%s is does not exists' % (request.data.get('tag'))),
                    status=status.HTTP_400_BAD_REQUEST
                )

        serializer = TodoSerializer(instance = todo_instance, data=data, partial = True)
        if serializer.is_valid():
            serializer.save()
            return Response(serializer.data, status=status.HTTP_200_OK)
        return Response(serializer.errors, status=status.HTTP_400_BAD_REQUEST)

    # Delete
    def delete(self, request, todo_id, *args, **kwargs):
        '''
        Deletes the todo item with given todo_id if exists
        '''
        todo_instance = self.get_object(todo_id)
        if not todo_instance:
            return Response(
                {"res": "Object with todo id does not exists"},
                status=status.HTTP_400_BAD_REQUEST
            )
        todo_instance.delete()
        return Response(
            {"res": "Object deleted!"},
            status=status.HTTP_200_OK
        )