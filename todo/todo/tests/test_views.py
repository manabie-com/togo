from django.http import response
from django.test import TestCase, Client
from django.urls import reverse, resolve
from todo.views import login, tasks
from todo.models import User, Task
from django.test.client import RequestFactory
import json

class TestViews(TestCase):

    def setUp(self):
        self.username = 'firstUser'
        self.password = 'example'

        self.user = User.objects.create(username = self.username, password = self.password)

        client = Client()

    def mock_login(self):
        credentials = {'user_id': self.username, 'password': self.password}
        return self.client.post('/login/', json.dumps(credentials), content_type="application/json")

    def test_login(self):
        response = self.mock_login()
        self.assertEqual(response.status_code, 200)

    def test_view_tasks_unauthorized(self):
        response = self.client.get('/tasks/?created_date=2000-01-01')
        self.assertEqual(response.status_code, 401)

    def test_view_tasks_authorized(self):
        self.mock_login()
        response = self.client.get('/tasks/?created_date=2000-01-01')
        self.assertEqual(response.status_code, 200)

    def test_create_tasks_unauthorized(self):
        content = {'content': 'sample content'}
        response = self.client.post('/tasks/', json.dumps(content), content_type="application/json")
        self.assertEqual(response.status_code, 401)

    def test_create_tasks_authorized(self):
        self.mock_login()
        content = {'content': 'sample content'}
        response = self.client.post('/tasks/', json.dumps(content), content_type="application/json")
        self.assertEqual(response.status_code, 200)
