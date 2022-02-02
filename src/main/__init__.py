from util import *

from flask import Flask

from .api import api
from .api.error import default_error_handler, HTTPError
from src.main.util import EnhancedEncoder

def create_app(app_name):
    app = Flask(app_name)
    app.json_encoder = EnhancedEncoder
    app.register_blueprint(api)
    app.register_error_handler(HTTPError, default_error_handler)
    return app
