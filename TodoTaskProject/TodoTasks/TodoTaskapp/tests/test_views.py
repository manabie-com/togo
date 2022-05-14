from rest_framework import status
from rest_framework.test import APIRequestFactory
from django.test import TestCase, Client
from django.urls import reverse
from ..models import UserRecordTask, UserTaskAllow
from ..serializers import UserRecordTaskSerializer, UserTaskAllowSerializer
from django.contrib.auth.models import User
from datetime import datetime
from django.utils.timezone import utc
import base64
from rest_framework import HTTP_HEADER_ENCODING, status
import json

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
        
class GetSingleUserRecordTaskTest(TestCase):
    """ Test module for GET single puppy API """

    def setUp(self):
        username = "manh"
        password = "Foxconn168!!"

        # Create database user. It needs to be created with django set_password function.
        iuser=User.objects.create_user(username=username, password=password)

        # Generate base64 credentials string
        credentials = f"{username}:{password}"
        base64_credentials = base64.b64encode(
            credentials.encode(HTTP_HEADER_ENCODING)
        ).decode(HTTP_HEADER_ENCODING)
        #-------------------------------------------------
        self.record1 = UserRecordTask.objects.create(user=iuser, TaskTitle="Test Gateway",
                                      TaskDescription="Test 100 pcs Gateway product", TaskDay=now)
        self.response = client.get('/userrecord/1/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}")

    def test_get_valid_single_userrecordtask(self):
        response = self.response
        irecord = UserRecordTask.objects.get(pk=self.record1.pk)
        serializer = UserRecordTaskSerializer(irecord)
        print(serializer.data)
        self.assertEqual(response.data, serializer.data)
        self.assertEqual(response.status_code, status.HTTP_200_OK)

    #def test_get_invalid_single_userrecordtask(self):
    #    response = client.get('/userrecord/1/')
    #    self.assertEqual(response.status_code, status.HTTP_404_NOT_FOUND)


class UpdateSingleUserRecordTaskTest(TestCase):
    """ Test module for updating an existing puppy record """
    def defaultconverter(self, a):
        if isinstance(self, datetime):
            return self.__str__()

    def setUp(self):
        username = "manh1"
        password = "Foxconn168!!"

        # Create database user. It needs to be created with django set_password function.
        iuser=User.objects.create_user(username=username, password=password)

        # Generate base64 credentials string
        credentials = f"{username}:{password}"
        self.base64_credentials = base64.b64encode(
            credentials.encode(HTTP_HEADER_ENCODING)
        ).decode(HTTP_HEADER_ENCODING)
        #-------------------------------------------------
        self.task1 = UserRecordTask.objects.create(user=iuser, TaskTitle="Test Gateway",
                                      TaskDescription="Test 100 pcs Gateway product", TaskDay=now)
        self.task2 = UserRecordTask.objects.create(user=iuser, TaskTitle="Test Cable Modem",
                                      TaskDescription="Test 100 pcs Cable Modem product", TaskDay=now)
        self.valid_payload = {
            'user': 'manh1',
            'TaskTitle': 'Test Gateway',
            'TaskDescription': 'Test 50 pcs Gateway product',
            'TaskDay': str(now)
        }
        self.invalid_payload = {
            'user': '',
            'TaskTitle': 'Test Cable Modem',
            'TaskDescription': 'Test 100 pcs Cable Modem product',
            'TaskDay': str(now)
        }

    def test_valid_update_userrecordtask(self):
        base64_credentials = self.base64_credentials
        response = client.put('/userrecord/' + str(self.task1.pk) + '/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}",
            data=json.dumps(self.valid_payload, indent=4, sort_keys=True, default=str),
            content_type='application/json')
        self.assertEqual(response.status_code, status.HTTP_202_ACCEPTED)#HTTP_204_NO_CONTENT)

    '''def test_invalid_update_userrecordtask(self):
        base64_credentials = self.base64_credentials
        response = client.put('/userrecord/' + str(self.task2.pk) + '/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}",
            data=json.dumps(self.invalid_payload, indent=4, sort_keys=True, default=str),
            content_type='application/json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)'''


class DeleteSingleUserRecordTaskTest(TestCase):
    """ Test module for deleting an existing puppy record """

    def setUp(self):
        username = "manh1"
        password = "Foxconn168!!"

        # Create database user. It needs to be created with django set_password function.
        iuser=User.objects.create_user(username=username, password=password)

        # Generate base64 credentials string
        credentials = f"{username}:{password}"
        self.base64_credentials = base64.b64encode(
            credentials.encode(HTTP_HEADER_ENCODING)
        ).decode(HTTP_HEADER_ENCODING)
        #-------------------------------------------------
        self.task1 = UserRecordTask.objects.create(user=iuser, TaskTitle="Test Gateway",
                                      TaskDescription="Test 100 pcs Gateway product", TaskDay=now)
        self.task2 = UserRecordTask.objects.create(user=iuser, TaskTitle="Test Cable Modem",
                                      TaskDescription="Test 100 pcs Cable Modem product", TaskDay=now)

    def test_valid_delete_userrecordtask(self):
        base64_credentials = self.base64_credentials
        response = client.delete('/userrecord/' + str(self.task1.pk) + '/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}")
        self.assertEqual(response.status_code, status.HTTP_204_NO_CONTENT)

    def test_invalid_delete_userrecordtask(self):
        base64_credentials = self.base64_credentials
        response = client.delete('/userrecord/100/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}")
        self.assertEqual(response.status_code, status.HTTP_404_NOT_FOUND)


class GetAllUserTaskAllowsTest(TestCase):
    """ Test module for GET all puppies API """

    def setUp(self):
        iuser = User.objects.create(username='manh')
        UserTaskAllow.objects.create(user=iuser, task_allow=3,
                                      task_done=1, start_task_time=now, last_task_time=now)

        
    def test_get_all_usertaskallows(self):
        # get API response
        username = "manh2"
        password = "Foxconn168!!"

        # Create database user. It needs to be created with django set_password function.
        iuser=User.objects.create_user(username=username, password=password)

        # Generate base64 credentials string
        credentials = f"{username}:{password}"
        base64_credentials = base64.b64encode(
            credentials.encode(HTTP_HEADER_ENCODING)
        ).decode(HTTP_HEADER_ENCODING)
        #-------------------------------------------------        
        response = client.get('/usertaskallows/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}") 
        print(response.data)
        # get data from db
        taskallows = UserTaskAllow.objects.all()
        serializer = UserTaskAllowSerializer(taskallows, many=True)
        self.assertEqual(response.data, serializer.data)
        self.assertEqual(response.status_code, status.HTTP_200_OK)
        
class GetSingleUserTaskAllowTest(TestCase):
    """ Test module for GET single puppy API """

    def setUp(self):
        username = "manh3"
        password = "Foxconn168!!"

        # Create database user. It needs to be created with django set_password function.
        iuser=User.objects.create_user(username=username, password=password)

        # Generate base64 credentials string
        credentials = f"{username}:{password}"
        base64_credentials = base64.b64encode(
            credentials.encode(HTTP_HEADER_ENCODING)
        ).decode(HTTP_HEADER_ENCODING)
        #-------------------------------------------------
        self.taskallow1 = UserTaskAllow.objects.create(user=iuser, task_allow=3,
                                      task_done=1, start_task_time=now, last_task_time=now)
        self.response = client.get('/usertaskallow/1/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}")

    def test_get_valid_single_usertaskallow(self):
        response = self.response
        itaskallow = UserTaskAllow.objects.get(pk=self.taskallow1.pk)
        serializer = UserTaskAllowSerializer(itaskallow)
        print(serializer.data)
        self.assertEqual(response.data, serializer.data)
        self.assertEqual(response.status_code, status.HTTP_200_OK)

    #def test_get_invalid_single_usertaskallow(self):
    #    response = client.get('/usertaskallow/1/')
    #    self.assertEqual(response.status_code, status.HTTP_404_NOT_FOUND)

class UpdateSingleUserTaskAllowTest(TestCase):
    """ Test module for updating an existing puppy record """
    def defaultconverter(self, a):
        if isinstance(self, datetime):
            return self.__str__()

    def setUp(self):
        username = "manh4"
        password = "Foxconn168!!"

        # Create database user. It needs to be created with django set_password function.
        iuser=User.objects.create_user(username=username, password=password)

        # Generate base64 credentials string
        credentials = f"{username}:{password}"
        self.base64_credentials = base64.b64encode(
            credentials.encode(HTTP_HEADER_ENCODING)
        ).decode(HTTP_HEADER_ENCODING)
        #-------------------------------------------------
        self.allowtask1 = UserTaskAllow.objects.create(user=iuser, task_allow=3,
                                      task_done=1, start_task_time=now, last_task_time=now)
        self.allowtask1 = UserTaskAllow.objects.create(user=iuser, task_allow=3,
                                      task_done=2, start_task_time=now, last_task_time=now)
        self.valid_payload = {
            'user': 'manh4',
            'task_allow': 3,
            'task_done': 1,
            'start_task_time': str(now),
            'last_task_time': str(now)
        }
        self.invalid_payload = {
            'user': '',
            'task_allow': 3,
            'task_done': 2,
            'start_task_time': str(now),
            'last_task_time': str(now)
        }

    def test_valid_update_usertaskallow(self):
        base64_credentials = self.base64_credentials
        response = client.put('/usertaskallow/' + str(self.allowtask1.pk) + '/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}",
            data=json.dumps(self.valid_payload, indent=4, sort_keys=True, default=str),
            content_type='application/json')
        self.assertEqual(response.status_code, status.HTTP_202_ACCEPTED)#HTTP_204_NO_CONTENT)

    '''def test_invalid_update_usertaskallow(self):
        base64_credentials = self.base64_credentials
        response = client.put('/usertaskallow/' + str(self.allowtask2.pk) + '/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}",
            data=json.dumps(self.invalid_payload, indent=4, sort_keys=True, default=str),
            content_type='application/json')
        self.assertEqual(response.status_code, status.HTTP_400_BAD_REQUEST)'''


class DeleteSingleUserTaskAllowTest(TestCase):
    """ Test module for deleting an existing puppy record """

    def setUp(self):
        username = "manh5"
        password = "Foxconn168!!"

        # Create database user. It needs to be created with django set_password function.
        iuser=User.objects.create_user(username=username, password=password)

        # Generate base64 credentials string
        credentials = f"{username}:{password}"
        self.base64_credentials = base64.b64encode(
            credentials.encode(HTTP_HEADER_ENCODING)
        ).decode(HTTP_HEADER_ENCODING)
        #-------------------------------------------------
        self.task1 = UserTaskAllow.objects.create(user=iuser, TaskTitle="Test Gateway",
                                      TaskDescription="Test 100 pcs Gateway product", TaskDay=now)
        self.task2 = UserTaskAllow.objects.create(user=iuser, TaskTitle="Test Cable Modem",
                                      TaskDescription="Test 100 pcs Cable Modem product", TaskDay=now)

    def test_valid_delete_usertaskallow(self):
        base64_credentials = self.base64_credentials
        response = client.delete('/usertaskallow/' + str(self.task1.pk) + '/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}")
        self.assertEqual(response.status_code, status.HTTP_204_NO_CONTENT)

    def test_invalid_delete_usertaskallow(self):
        base64_credentials = self.base64_credentials
        response = client.delete('/usertaskallow/100/', HTTP_AUTHORIZATION=f"Basic {base64_credentials}")
        self.assertEqual(response.status_code, status.HTTP_404_NOT_FOUND)
