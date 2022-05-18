import datetime

from factory import django, SubFactory, Sequence
from faker import Faker as FakerClass
from rest_framework import status
from rest_framework.test import APITestCase, APIClient
# from rest_framework_simplejwt.tokens import RefreshToken
from rest_framework.authtoken.models import Token
from rest_framework_simplejwt.tokens import RefreshToken

from authentication.models import User
from authentication.tests import UserFactory
from todos.models import SettingTask, Task

faker_obj = FakerClass()


def check_setting_success(status_code, data):
    return "id" in data \
           and "date" in data \
           and "user" in data \
           and "user" in data \
           and "limit" in data \
           and status_code == status.HTTP_200_OK


def check_setting_fail(status_code, data):
    return "id" not in data \
           and status_code == status.HTTP_400_BAD_REQUEST


def check_task_success(status_code, data):
    return 'id' in data \
           and status_code == status.HTTP_200_OK


def check_task_fail_with_limit(status_code, data):
    return 'error' in data and status_code == status.HTTP_400_BAD_REQUEST

class SettingTaskFactory(django.DjangoModelFactory):
    date = datetime.date.today()
    limit = faker_obj.random_int(0, 5)
    user = SubFactory(UserFactory)

    class Meta:
        model = SettingTask


class TaskFactory(django.DjangoModelFactory):
    setting = SubFactory(UserFactory)
    is_accepted = True
    name = Sequence(lambda n: 'Task %d' % n)

    class Meta:
        model = Task


class TestCaseBase(APITestCase):

    @classmethod
    def setUpClass(cls):
        super(TestCaseBase, cls).setUpClass()
        cls.client = APIClient()
        cls.faker_obj = faker_obj
