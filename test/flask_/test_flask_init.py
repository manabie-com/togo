import unittest

import werkzeug
from celery_.config import Config
from flask_.init import ApiInit


class AppInitTests(unittest.TestCase):
    def setUp(self):
        self.test_app = ApiInit("test")
        self.test_client = self.test_app.test_client()

        @self.test_app.get("/test")
        def test_route():
            return "test"

        @self.test_app.post("/test")
        def test_route_post():
            return "test"

    def test_empty_configuration(self):
        self.test_app.configure(None)
        self.assertEqual(self.test_app.blueprints, {})

    def test_normal_configuration(self):
        cnf = Config()
        cnf.BLUEPRINTS = ["auth", "task"]
        self.test_app.configure(cnf)
        self.assertEqual(len(self.test_app.blueprints), 2)

    def test_wrong_blueprints(self):
        cnf = Config()
        cnf.BLUEPRINTS = ["nonexisting"]
        self.assertRaises(
            werkzeug.utils.ImportStringError, self.test_app.configure, cnf
        )
