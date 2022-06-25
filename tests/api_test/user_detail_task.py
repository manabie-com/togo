import unittest
from rest_framework.test import APITestCase
from rest_framework import status
from rest_framework_simplejwt.tokens import RefreshToken
from django.contrib.auth.models import User
from django.urls import reverse
from dateutil.relativedelta import relativedelta
from datetime import date

from apps.models.schedule import Schedule
from apps.models.task import Task


class UserDetailTaskAPITest(APITestCase):
    @classmethod
    def setUpTestData(cls):
        user_test_1 = User.objects.create(username='test_user_1')
        user_test_1.set_password('1')
        user_test_1.save()
        user_test_2 = User.objects.create(username='test_user_2')
        user_test_2.set_password('1')
        user_test_2.save()
        tasks_name = ['cooking', 'booking', 'ordering', 'playing', 'watching', 'seeking', 'trekking']
        for i, name in enumerate(tasks_name):
            Task.objects.create(**{'id': i, 'name': name})
        return super().setUpTestData()

    def setUp(self):
        self.user_test_1 = User.objects.get(username='test_user_1')
        self.user_test_2 = User.objects.get(username='test_user_2')
        self.task_1 = Task.objects.get(id=1)
        return super().setUp()

    def token(self):
        refresh = RefreshToken.for_user(self.user_test_1)
        self.client.login(username=self.user_test_1.username, password='1')
        self.client.credentials(HTTP_AUTHORIZATION=f'Bearer {refresh}')

    def test_assignment_success(self):
        self.token()
        example = {
            "task": self.task_1.id,
            "user": self.user_test_2.id,
            "date": "2022-06-22"
        }
        response = self.client.post(reverse('assignment'), example, format='json')
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)

    def test_assignment_failed_because_not_login(self):
        example = {
            "task": self.task_1.id,
            "user": self.user_test_2.id,
            "date": "2022-06-22"
        }
        response = self.client.post(reverse('assignment'), example, format='json')
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_user_does_not_exits(self):
        self.token()
        example = {
            "task": self.task_1.id,
            "user": 100000000000,
            "date": "2022-06-22"
        }
        response = self.client.post(reverse('assignment'), example)
        self.assertEqual(response.status_code, status.HTTP_404_NOT_FOUND)

    def test_task_does_not_exits(self):
        self.token()
        example = {
            "task": 100000000000,
            "user": self.user_test_1.id,
            "date": "2022-06-22"
        }
        response = self.client.post(reverse('assignment'), example)
        self.assertEqual(response.status_code, status.HTTP_404_NOT_FOUND)

    def test_current_date_gt_than_today_in_request(self):
        self.token()
        example = {
            "task": self.task_1.id,
            "user": self.user_test_1.id,
            "date": date.today() + relativedelta(days=+1)
        }
        response = self.client.post(reverse('assignment'), example)
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)

    def test_pick_limit_for_assignment_user(self):
        self.token()
        schedule = Schedule.objects.create(user=self.user_test_2, limit=3, date="2022-06-22")
        limit = schedule.limit
        response = None
        for i in range(limit + 2):
            example = {
                "task": i,
                "user": self.user_test_2.id,
                "date": "2022-06-22"
            }
            response = self.client.post(reverse('assignment'), example)
        self.assertEqual(response.status_code, status.HTTP_500_INTERNAL_SERVER_ERROR)

    def test_success_assigment_test_difference_day(self):
        self.test_pick_limit_for_assignment_user()
        self.token()
        example = {
            "task": self.task_1.id,
            "user": self.user_test_2.id,
            "date": str(date.today())
        }
        response = self.client.post(reverse('assignment'), example, format='json')
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)


if __name__ == '__main__':
    unittest.main()
