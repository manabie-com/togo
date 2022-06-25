from datetime import date
from apps.models.schedule import Schedule
from django.contrib.auth.models import User
from django.test import TestCase, override_settings
from apps.tasks.task_limited_each_day import pick_limit_for_user, callback_get_limit_task
from apps.views.user_detail_task import CreateDetail
import pytest


class UserDetailTaskTestCase(TestCase):
    @classmethod
    def setUpTestData(cls):
        user_test_1 = User.objects.create(username='test_user_1')
        user_test_1.set_password('1')
        user_test_1.save()
        return super().setUpTestData()

    def setUp(self):
        self.date = date.today()
        self.user_test_1 = User.objects.get(username='test_user_1')
        self.schedule = Schedule.objects.create(id=1, user=self.user_test_1, limit=3, date=str(self.date))
        return super().setUp()

    def test_get_schedule_has_on_db(self):
        date_ = str(self.date)
        schedule = CreateDetail.get_schedule_task(self.user_test_1, date_)
        self.assertEqual(schedule, self.schedule)

    def test_get_schedule_does_not_exists(self):
        date_ = str(self.date)
        user_test_2 = User.objects.create(username='test_user_2')
        schedule = CreateDetail.get_schedule_task(user_test_2, date_)
        schedule_compare = Schedule.objects.get(id=schedule.id)
        self.assertEqual(schedule_compare, schedule)

    def test_pick_limit_for_user_failed(self):
        date_ = str(self.date)
        task = pick_limit_for_user.s(date_).apply()
        self.assertEqual(task.result, False)

    def test_pick_limit_for_user_success(self):
        date_ = str(self.date)
        Schedule.objects.all().delete()
        task = pick_limit_for_user.s(date_).apply()
        self.assertEqual(task.result, True)

    def test_callback_get_limit_task(self):
        date_ = str(self.date)
        Schedule.objects.all().delete()
        Schedule.objects.create(id=1, user=self.user_test_1, limit=3, date=date_)
        User.objects.create(username='test_user_2')
        callback_get_limit_task.s(False, date_).apply()
        len_all = len(Schedule.objects.all())
        len_user = len(User.objects.all())
        self.assertEqual(len_all, len_user)
