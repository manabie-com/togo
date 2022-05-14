from rest_framework import status
from rest_framework.test import APIRequestFactory
from django.test import TestCase, Client
from django.urls import reverse
from ..models import UserRecordTask, UserTaskAllow
from ..serializers import UserRecordTaskSerializer, UserTaskAllowSerializer
from django.contrib.auth.models import User
from datetime import datetime
from django.utils.timezone import utc

# initialize the APIClient app
client = Client()
now = datetime.utcnow().replace(tzinfo=utc)
print(now)

class GetAllUserRecordTasksTest(TestCase):
    """ Test module for GET all puppies API """

    def setUp(self):
        iuser = User.objects.create(username='manh')
        UserRecordTask.objects.create(user=iuser, TaskTitle="Test Gateway",
                                      TaskDescription="Test 100 pcs Gateway product", TaskDay=now)

        
    def test_get_all_userrecordtasks(self):
        # get API response
        response = client.get('/userrecords/') 
        print(response.data)
        # get data from db
        userrecords = UserRecordTask.objects.all()
        serializer = UserRecordTaskSerializer(userrecords, many=True)
        self.assertEqual(response.data, serializer.data)
        self.assertEqual(response.status_code, status.HTTP_200_OK)

