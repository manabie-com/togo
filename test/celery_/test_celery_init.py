from celery_.init import ServiceInit, factory
from mongoengine.connection import disconnect
from test.celery_.base_test_case import BaseTestCase


class CeleryInitTests(BaseTestCase):
    def test_default_configuration(self):
        app = factory("Mock")
        with app:
            self.assertIsInstance(app, ServiceInit)

        disconnect(alias="app-db")
