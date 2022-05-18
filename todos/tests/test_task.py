from django.urls import reverse
from rest_framework_simplejwt.tokens import RefreshToken
from rest_framework.test import APITestCase, APIClient

from authentication.models import User
from todos.models import Task
from todos.tests import TestCaseBase, check_task_fail_with_limit, check_task_success, SettingTaskFactory


class TaskTestCase(TestCaseBase):

    @classmethod
    def setUpClass(cls):
        super(TaskTestCase, cls).setUpClass()
        cls.api_url = reverse('tasks')

    def test_create_task_success(self):
        print("TEST CREATE TASK SUCCESS")

        data = {
            'name': "Task 1"
        }
        user_exists = User.objects.create(
            email='js@js.com',
            password='js.sj',
            first_name="Vi",
            last_name="Luong")
        client = APIClient()
        refresh_obj = RefreshToken.for_user(user_exists)
        client.credentials(HTTP_AUTHORIZATION='Bearer ' + str(refresh_obj.access_token))
        response = client.post(self.api_url, data)
        self.assertTrue(check_task_success(response.status_code, response.data))

    def test_create_task_fail_with_greater_limit(self):
        print("TEST CREATE TASK FAIL WITH LIMIT")

        data = {
            'name': "Task 1"
        }
        client = APIClient()

        user_exists = User.objects.create(
            email='js@js.com',
            password='js.sj',
            first_name="Vi",
            last_name="Luong")
        settings_exists = SettingTaskFactory.create(user=user_exists)
        objs = (Task(name='Task %s' % i, setting=settings_exists) for i in range(5))
        Task.objects.bulk_create(objs)
        refresh_obj = RefreshToken.for_user(user_exists)
        client.credentials(HTTP_AUTHORIZATION='Bearer ' + str(refresh_obj.access_token))
        response = client.post(self.api_url, data)
        self.assertTrue(check_task_fail_with_limit(response.status_code, response.data))
