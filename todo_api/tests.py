from django.test import TestCase

from django.urls import reverse
from rest_framework.test import APITestCase
from .models import CustomUser, Todo
from .serializers import CustomUserSerializer
from rest_framework import status
import base64


# Create User
class UsersTest(APITestCase):
    def setUp(self):
        self.test_number_todo_limit = 5
        self.test_user = CustomUser.objects.create_user('testuser', 'testpassword', number_todo_limit=self.test_number_todo_limit)
        self.credentials = base64.b64encode(b'testuser:testpassword')
        self.client.defaults['HTTP_AUTHORIZATION'] = 'Basic ' + self.credentials.decode("ascii")

        self.create_url = reverse('user-create')
        self.retrieve_url = reverse('user-get')

    # create user success
    def test_create_user(self):
        data = {
            'username': 'usertest',
            'password': 'passwordtest',
            'number_todo_limit': 5
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(CustomUser.objects.count(), 2)
        self.assertEqual(response.status_code, status.HTTP_200_OK)

    # create user without username
    def test_create_user_without_username(self):
        data = {
            'password': 'passwordtest',
            'number_todo_limit': 5
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with username is exist
    def test_create_user_with_username_is_exist(self):
        data = {
            'username': 'testuser',
            'password': 'testpassword',
            'number_todo_limit': 5
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with username exceed 150
    def test_create_user_with_username_exceed_150(self):
        data = {
            'username': 'foo' * 51,
            'password': 'passwordtest',
            'number_todo_limit': 5
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with short password
    def test_create_user_with_short_password(self):
        data = {
            'username': 'usertest',
            'password': 'foo',
            'number_todo_limit': 5
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with number_todo_limit is less than 0
    def test_create_user_with_number_todo_limit_lt_zero(self):
        data = {
            'username': 'usertest',
            'password': 'passwordtest',
            'number_todo_limit': -1
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with number_todo_limit is not integer
    def test_create_user_with_number_todo_limit_is_not_integer(self):
        data = {
            'username': 'usertest',
            'password': 'passwordtest',
            'number_todo_limit': 'foo'
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

    # create user with number_todo_limit is none
    def test_create_user_with_number_todo_limit_is_none(self):
        data = {
            'username': 'usertest',
            'password': 'passwordtest',
            'number_todo_limit': None
        }
        response = self.client.post(self.create_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)
        self.assertEqual(CustomUser.objects.count(), 1)

# Retrieve User
    # retrieve user success
    def test_retrieve_user(self):
        serializer = CustomUserSerializer(self.test_user)
        response = self.client.get(self.retrieve_url, format='json')
        self.assertEqual(response.status_code, status.HTTP_200_OK)

        self.assertEqual(serializer.data['username'], response.data['username'])

    # retrieve user id not exist


# Update User
    # update user success
    def test_update_user(self):
        data = {
            'number_todo_limit': 5
        }
        response = self.client.put(self.retrieve_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        self.assertEqual(data['number_todo_limit'], response.data['number_todo_limit'])

    # update user with number_todo_limit is less than 0
    def test_update_user_with_number_todo_limit_lt_zero(self):
        data = {
            'number_todo_limit': -1
        }
        response = self.client.put(self.retrieve_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

    # update user with number_todo_limit is not integer
    def test_update_user_with_number_todo_limit_is_not_integer(self):
        data = {
            'number_todo_limit': "foo"
        }
        response = self.client.put(self.retrieve_url, data, format='json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

    # update user with number_todo_limit is none
    def test_update_user_with_number_todo_limit_lt_zero(self):
        data = {
            'number_todo_limit': None
        }
        response = self.client.put(self.retrieve_url, data, format='json')

        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

