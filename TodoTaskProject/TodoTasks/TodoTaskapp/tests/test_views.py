from rest_framework import status
from django.test import TestCase, Client
from django.urls import reverse
from ..models import UserRecordTask, UserTaskAllow
from ..serializers import UserRecordTaskSerializer, UserTaskAllowSerializer
from django.contrib.auth.models import User
from datetime import datetime
from django.utils.timezone import utc

# initialize the APIClient app
client = Client()
