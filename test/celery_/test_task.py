import os
import sys

from database.models.task import UserTask
from celery_.response import ServiceResponse
from flask_.error import InvalidUsage
from test.celery_.base_test_case import BaseTestCase

sys.path.append(os.getcwd())


class TaskTest(BaseTestCase):
    def setUp(self):
        super().add_mock_services_to_db()
        from celery_.services import task_service

        self.task_service = task_service
        UserTask.objects.insert(
            [
                UserTask(
                    user_id="6179257a0551dddc2be4478f",
                    request_number_per_day=0,
                    limit=5,
                ),
                UserTask(
                    user_id="6179257a0551dddc2be44799",
                    request_number_per_day=5,
                    limit=5,
                ),
            ]
        )

    def tearDown(self):
        super().drop_collection()
        UserTask.drop_collection()

    def test_post_task(self):
        user_id = "6179257a0551dddc2be4478f"

        result = self.task_service.post_task.s(user_id).apply().get()

        self.assertIs(type(result), ServiceResponse)
        self.assertEqual(
            UserTask.objects.get(user_id=user_id).request_number_per_day, 1
        )

    def test_post_task_reach_limit(self):
        user_id = "6179257a0551dddc2be44799"

        task = self.task_service.post_task.s(user_id).apply()

        self.assertIs(type(task.result), InvalidUsage)
        self.assertEqual(task.result.payload, "Bad request")
        self.assertEqual(
            task.result.message, "Reached the limit for number of requests already"
        )

    def test_initialize_user_task(self):
        post_signup_response = {"user_id": "6179257a0551dddc2be44711"}

        result = (
            self.task_service.initialize_user_task.s(post_signup_response).apply().get()
        )

        self.assertIs(type(result), ServiceResponse)
        self.assertIsNot(
            UserTask.objects.get(user_id=post_signup_response.get("user_id")), None
        )

    def test_reset_counter_request(self):
        result = self.task_service.reset_counter_request.s().apply().get()

        self.assertIs(type(result), ServiceResponse)
        self.assertEqual(
            UserTask.objects.get(
                user_id="6179257a0551dddc2be44799"
            ).request_number_per_day,
            0,
        )
