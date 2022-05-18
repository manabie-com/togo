import uuid

from factory import django, Faker
from faker import Faker as FakerClass
from rest_framework import status

from rest_framework.test import APITestCase, APIClient

from authentication import models

faker_obj = FakerClass()


COMMON_PASSWORD = faker_obj.password(
    length=12,
    special_chars=True,
    digits=True,
    upper_case=True,
    lower_case=True
)


def check_token_success(status_code, data):
    return "access" in data \
           and "refresh" in data \
           and "user" in data \
           and "refresh_token_expire_at" in data \
           and "token_expire_at" in data \
           and status_code == status.HTTP_200_OK


def fake_gmail():
    str_random = str(uuid.uuid1())
    return str_random+"@gmail.com"


class UserFactory(django.DjangoModelFactory):
    first_name = Faker('first_name')
    last_name = Faker('last_name')
    email = fake_gmail()
    password = COMMON_PASSWORD

    class Meta:
        model = models.User


class TestCaseBase(APITestCase):

    @classmethod
    def setUpClass(cls):
        super(TestCaseBase, cls).setUpClass()
        cls.client = APIClient()
        cls.faker_obj = faker_obj
