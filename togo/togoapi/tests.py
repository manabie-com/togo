from datetime import datetime
from django.test import TestCase
from rest_framework.test import APITestCase
from rest_framework import status

class TestAddTodo(APITestCase):

    correct_header = {
        "HTTP_AUTHORIZATION": "Api-Key qkYpjT1D.Yg3aa1kv4ghPmh5lg2NCMi5PWmIp8Cy4", 
        "HTTP_USERNAME": "choerry"
    }
    
    correct_header_wrong = {
        "HTTP_AUTHORIZATION": "Api-Key qkYpjT1D.Yg3aa1kv4ghPmh5lg2NCMi5PWmIp8Cy4", 
        "HTTP_USERNAME": "choerry"
    }

    todo = {
        "title": "grocery", 
        "description": "buy eggs", 
        "start_time":"2022-05-16 15:15:00", 
        "end_time":"2022-05-16 16:00:00",
        "added_time": datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    }

    endpoint = "/usertasks/"
  
    def test_create_todo_no_auth_header(self):
        response = self.client.post(self.endpoint, self.todo)
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_create_todo_wrong_auth_header(self):
        self.correct_header_wrong["Authorization"] = "Basic qkYpjT1D.Yg3aa1kv4ghPmh5lg2NCMi5PWmIp8Cy4" 

        response = self.client.post(self.endpoint, self.todo, **self.correct_header_wrong)
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_create_todo_wrong_api_key(self):
        self.correct_header_wrong["Authorization"] = "Api-Key qkYpjT1D.Yg3aa1kv4ghPmh5lg2NCMi5PWmIp8Cy5" 
        
        response = self.client.post(self.endpoint, self.todo, **self.correct_header_wrong)
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_create_todo_no_username_header(self):
        self.correct_header_wrong.pop("HTTP_USERNAME") 

        response = self.client.post(self.endpoint, self.todo, **self.correct_header_wrong)
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    def test_create_todo_username_not_exists(self):
        self.correct_header_wrong["Username"] = "idontexist" 

        response = self.client.post(self.endpoint, self.todo, **self.correct_header_wrong)
        self.assertEqual(response.status_code, status.HTTP_401_UNAUTHORIZED)

    #TODO: response is returning 401 - double check if headers are being set properly
    def test_create_todo_valid(self):
        response = self.client.post(self.endpoint, self.todo, **self.correct_header)
        print(response.data, response.content)
        self.assertEqual(response.status_code, status.HTTP_201_CREATED)
