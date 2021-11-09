from django.test import TestCase, Client
from django.urls import reverse, resolve
from todo.views import login, tasks
from todo.models import User, Task

class TestModels(TestCase):

    def setUp(self):
        self.username = "username"
        self.password = "password"
        self.content = "sample content"

        self.user = User.objects.create(username = self.username, password = self.password)
        self.task = Task.objects.create(content = self.content, user_id = self.user)

    def test_user_created(self):
        self.assertEquals(self.user.username, self.username)
        self.assertEquals(self.user.password, self.password)

    def test_tasks_created(self):
        self.assertEquals(self.task.user_id, self.user)
        self.assertEquals(self.task.content, self.content)
