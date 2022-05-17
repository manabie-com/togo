from django.urls import reverse
from rest_framework.test import APITestCase
from .models import CustomUser, Todo
from .serializers import UserSerializer
from rest_framework import status
import base64


# Create User
class UsersTest(APITestCase):
    def setUp(self):
        self.test_todo_max = 3
        self.test_user = CustomUser.objects.create_user('testuser', 'testpassword', todo_max=self.test_todo_max)
        self.credentials = base64.b64encode(b'testuser:testpassword')
        self.client.defaults['HTTP_AUTHORIZATION'] = 'Basic ' + self.credentials.decode("ascii")

        self.create_url = reverse('user-create')
        self.retrieve_url = reverse('user-get')

    # create user success
    def test_create_user(self):
        data = {
            'username': 'foobar',
            'password': 'foopassword',
            'todo_max': 1
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(CustomUser.objects.count(), 2)
        self.assertEqual(response.status_code, status.HTTP_200_OK)

    # create user without username
    def test_create_user_without_username(self):
        data = {
            'password': 'foopassword',
            'todo_max': 1
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with username is exist
    def test_create_user_with_username_is_exist(self):
        data = {
            'username': 'testuser',
            'password': 'testpassword',
            'todo_max': 1
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with username exceed 150
    def test_create_user_with_username_exceed_150(self):
        data = {
            'username': 'foo' * 51,
            'password': 'foopassword',
            'todo_max': 1
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with short password
    def test_create_user_with_short_password(self):
        data = {
            'username': 'foobar',
            'password': 'foo',
            'todo_max': 1
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with todo_max is less than 0
    def test_create_user_with_todomax_lt_zero(self):
        data = {
            'username': 'foobar',
            'password': 'foopassword',
            'todo_max': -1
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with todo_max is not integer
    def test_create_user_with_todomax_is_not_integer(self):
        data = {
            'username': 'foobar',
            'password': 'foopassword',
            'todo_max': 'foo'
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with todo_max is none
    def test_create_user_with_todomax_is_none(self):
        data = {
            'username': 'foobar',
            'password': 'foopassword',
            'todo_max': None
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

# Retrieve User
    # retrieve user success
    def test_retrieve_user(self):
        serializer = UserSerializer(self.test_user)
        response = self.client.get(self.retrieve_url, format='json')
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(serializer.data['id'], response.data['id'])

    # retrieve user id not exist


# Update User
    # update user success
    def test_update_user(self):
        data = {
            'todo_max': 1
        }
        response = self.client.put(self.retrieve_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(data['todo_max'], response.data['todo_max'])

    # update user with todo_max is less than 0
    def test_update_user_with_todomax_lt_zero(self):
        data = {
            'todo_max': -1
        }
        response = self.client.put(self.retrieve_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

    # update user with todo_max is not integer
    def test_update_user_with_todomax_is_not_integer(self):
        data = {
            'todo_max': "foo"
        }
        response = self.client.put(self.retrieve_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

    # update user with todo_max is none
    def test_update_user_with_todomax_lt_zero(self):
        data = {
            'todo_max': None
        }
        response = self.client.put(self.retrieve_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)


# Create task

class TasksTest(APITestCase):
    def setUp(self):
        self.test_todo_max = 2
        self.test_user = CustomUser.objects.create_user('testuser', 'testpassword', todo_max=self.test_todo_max)
        self.credentials = base64.b64encode(b'testuser:testpassword')
        self.client.defaults['HTTP_AUTHORIZATION'] = 'Basic ' + self.credentials.decode("ascii")
        self.test_user_serializer = UserSerializer(self.test_user)

        self.test_task = Todo.objects.create(task='test task', completed=False, user=self.test_user)

        self.create_url = reverse('task-list')
        self.retrieve_url_name = 'task-detail'

    # create task success
    def test_create_task(self):
        data = {
            'task': 'footask',
            'completed': False,
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(Todo.objects.filter(user=self.test_user_serializer.data['id']).count(), 2)
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)

    # create task without task name
    def test_create_task_without_task_name(self):
        data = {
            'completed': False,
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(Todo.objects.filter(user=self.test_user_serializer.data['id']).count(), 1)
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

    # create task with completed is not boolean
    def test_create_task_with_completed_is_not_boolean(self):
        data = {
            'task': 'footask',
            'completed': 'foo',
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(Todo.objects.filter(user=self.test_user_serializer.data['id']).count(), 1)
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

    # create task exceed daily limit (todo_max)
    def test_create_task_exceed_daily_limit(self):
        Todo.objects.create(task='test task 01', completed=False, user=self.test_user)
        data = {
            'task': 'footask',
            'completed': 'foo',
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(Todo.objects.filter(user=self.test_user_serializer.data['id']).count(), self.test_todo_max)
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

# Retrieve task
    # retrieve list task by user id
    def test_retrieve_task_by_user_id(self):
        response = self.client.get(self.create_url, format='json')
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(len(response.data), 1)
    # retrieve task by task id

# Update task
    # update task success
    # update task with empty task name
    # update task with completed is not boolean


