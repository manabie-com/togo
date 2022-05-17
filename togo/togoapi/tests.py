from datetime import datetime
from django.test import TestCase
from rest_framework.test import APITestCase, APIClient

from rest_framework import status

from .util import detailvalidationutil
from .util import dbutil

from datetime import datetime, timedelta
from django.utils import timezone
from django.test.client import RequestFactory

from .models import User, UserTask
from rest_framework_api_key.models import APIKey

endpoint = "/usertasks/"
todo = {
            "title": "lessons", 
            "description": "study english"
        }
ht_auth = "HTTP_AUTHORIZATION"
ht_user = "HTTP_USERNAME"

time_format = "%Y-%m-%d %H:%M:%S"

class AddTodoAPITestCase(APITestCase):

    temp_key = None
    temp_user = "temp_user"

    correct_header_wrong = {
        ht_auth: "Api-Key qkYpjT1D.Yg3aa1kv4ghPmh5lg2NCMi5PWmIp8Cy4", 
        ht_user: "test_user"
    }

    def setUp(self):
        key, self.temp_key = APIKey.objects.create_key(name="my-temp-key")
        user = User(username=self.temp_user,daily_limit=2)
        user.save()

        self.correct_header = {
            ht_auth: "Api-Key " + str(self.temp_key), 
            ht_user: self.temp_user
        }
    
    def test_create_todo_no_auth_header(self):
        response = self.client.post(endpoint, todo)
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_create_todo_wrong_auth_header(self):
        self.correct_header_wrong[ht_auth] = "Basic " + self.temp_key 

        response = self.client.post(endpoint, todo, **self.correct_header_wrong)
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_create_todo_wrong_api_key(self):
        self.correct_header_wrong[ht_auth] = "Api-Key qkYpjT1D.Yg3aa1kv4ghPmh5lg2NCMi5PWmIp8Cy5" 
        
        response = self.client.post(endpoint, todo, **self.correct_header_wrong)
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_create_todo_no_username_header(self):
        self.correct_header_wrong.pop(ht_user) 

        response = self.client.post(endpoint, todo, **self.correct_header_wrong)
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_create_todo_username_not_exists(self):
        self.correct_header_wrong[ht_user] = "idontexist" 

        response = self.client.post(endpoint, todo, **self.correct_header_wrong)
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_create_todo_valid(self):
        response = self.client.post(endpoint, todo, **self.correct_header)
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)

class ValidationUtilTestCase(APITestCase):
    temp_user = "limit_two_test"

    def setUp(self):
        key, self.temp_key = APIKey.objects.create_key(name="my-temp-key")
        user = User(username=self.temp_user,daily_limit=2,task_today=0)
        user.save()

        self.correct_header = {
            ht_auth: "Api-Key " + str(self.temp_key), 
            ht_user: self.temp_user
        }

    def test_valid_schedule_valid_pair(self):
        start = (datetime.now(timezone.utc) + timedelta(minutes=1)).strftime(time_format)
        end = (datetime.now(timezone.utc) + timedelta(hours=1)).strftime(time_format)
        isValid = detailvalidationutil.validSchedule(start, end)
        self.assertEqual(isValid, True)

    def test_valid_schedule_invalid_pair(self):
        start = (datetime.now(timezone.utc) + timedelta(minutes=1)).strftime(time_format)
        end = (datetime.now(timezone.utc) - timedelta(hours=1)).strftime(time_format)
        isValid = detailvalidationutil.validSchedule(start, end)
        self.assertEqual(isValid, False)

    def test_valid_schedule_invalid_past_start(self):
        isValid = detailvalidationutil.validSchedule("2022-05-10 10:10:10", "2022-05-10 12:30:10")
        self.assertEqual(isValid, False)

    def test_daily_limit_reached_true(self):
        r1 = self.client.post(endpoint, todo, **self.correct_header)
        r2 = self.client.post(endpoint, todo, **self.correct_header)

        limitHasBeenReached = detailvalidationutil.dailyLimitReached(dbutil.user(self.temp_user))
        self.assertEqual(limitHasBeenReached, True)

    def test_daily_limit_reached_false(self):
        r1 = self.client.post(endpoint, todo, **self.correct_header)

        limitHasBeenReached = detailvalidationutil.dailyLimitReached(dbutil.user(self.temp_user))
        self.assertEqual(limitHasBeenReached, False)

class DBUtilTestCase(APITestCase):
    temp_user = "dbutil_test"

    def setUp(self):
        key, self.temp_key = APIKey.objects.create_key(name="my-temp-key")
        user = User(username=self.temp_user,daily_limit=2)
        user.save()

        self.correct_header = {
            ht_auth: "Api-Key " + str(self.temp_key), 
            ht_user: self.temp_user
        }

    def test_increment_user_daily_task(self):
        u = dbutil.user(self.temp_user)
        dbutil.CreateUtil.incrementUserDailyTask(u)
        self.assertEqual(u.task_today, 1)

    def test_default_end_time(self):
        st = "2022-05-10 10:10:10"
        et = dbutil.CreateUtil.defaultEndTime(st)
        self.assertEqual(et, "2022-05-10 11:10:10")

