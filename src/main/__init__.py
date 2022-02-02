from flask import Flask, redirect

from .api import api, API_PREFIX
from .api.error import default_error_handler, HTTPError
from src.main.util import EnhancedEncoder


def create_app(app_name):
    app = Flask(app_name)
    app.json_encoder = EnhancedEncoder

    @app.route("/")
    def homepage():
        return redirect(API_PREFIX)

    app.register_blueprint(api)
    app.register_error_handler(HTTPError, default_error_handler)
    return app
