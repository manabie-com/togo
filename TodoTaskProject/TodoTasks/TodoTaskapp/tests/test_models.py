from django.test import TestCase
from ..models import UserRecordTask, UserTaskAllow
from django.contrib.auth.models import User
from datetime import datetime
from django.utils.timezone import utc

now = datetime.utcnow().replace(tzinfo=utc)


class UserRecordTask_Test(TestCase):
    def setUp(self):
        iuser = User.objects.create(username='manh')
        UserRecordTask.objects.create(user=iuser, TaskTitle="Test Gateway", TaskDescription="Test 100 pcs Gateway product", TaskDay=now)

    def test_get_record(self):
        irecord = UserRecordTask.objects.get(user=1)
        self.assertEqual(irecord.get_record(), 'manh task: Test Gateway ' + str(now))


class UserTaskAllow_Test(TestCase):
    def setUp(self):
        iuser = User.objects.create(username='manh')
        UserTaskAllow.objects.create(user=iuser, task_allow=10, task_done=9, start_task_time=now, last_task_time=now)

    def test_get_record(self):
        iTaskAllow = UserTaskAllow.objects.get(user=1)
        self.assertEqual(iTaskAllow.get_taskallow(), 'manh allow 10 task. And done 9 task.')