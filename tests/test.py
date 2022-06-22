from unittest import TestCase
import unittest
from rest_framework.test import APITestCase
import django

django.setup()


class UserDetailTask(TestCase):

    def setUp(self):
        return super(UserDetailTask, self).setUp()

    def tearDown(self):
        return super().tearDown()

    def test_limit_when_user_assign(self):
        self.assertEqual(2, 2)

    def test_user_does_not_exits(self):
        ...

    def test_task_does_not_exits(self):
        ...

    def test_current_date_gt_than_date_in_request(self):
        ...

    def test_success_assignment(self):
        ...

    def test_pick_limit_for_user_sucess(self):
        ...


if __name__ == '__main__':
    unittest.main()
