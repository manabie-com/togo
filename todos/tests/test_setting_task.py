from django.urls import reverse
from rest_framework_simplejwt.tokens import RefreshToken
from rest_framework.test import APITestCase, APIClient

from authentication.models import User
from todos.tests import TestCaseBase, SettingTaskFactory, check_setting_success, check_setting_fail


class SettingTaskTestCase(TestCaseBase):

    @classmethod
    def setUpClass(cls):
        super(SettingTaskTestCase, cls).setUpClass()
        cls.api_url = reverse('task-settings')

    def test_create_setting_success(self):
        print("TEST CREATE SETTING SUCCESS")

        data = {
            'limit': 5
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
        self.assertTrue(check_setting_success(response.status_code, response.data))

    def test_create_setting_fail_with_exists_setting(self):
        print("TEST CREATE SETTING FAIL WITH EXISTS SETTING")
        user_exists = User.objects.create(
            email='js@js.com',
            password='js.sj',
            first_name="Vi",
            last_name="Luong")

        self.settings_exists = SettingTaskFactory.create(user=user_exists)
        data = {
            'limit': 5
        }
        client = APIClient()
        refresh_obj = RefreshToken.for_user(user_exists)
        client.credentials(HTTP_AUTHORIZATION='Bearer ' + str(refresh_obj.access_token))
        response = client.post(self.api_url, data)
        self.assertTrue(check_setting_fail(response.status_code, response.data))
