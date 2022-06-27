import os
from typing import Dict, List


class Config(object):
    DEBUG = os.getenv("FLASK_DEBUG", "1") == "1"
    HOST = os.getenv("FLASK_HOST", "0.0.0.0")
    HOST_PORT = os.getenv("FLASK_PORT", "5000")
    SECRET_KEY = os.getenv("FLASK_SECRET_KEY", "CHANGE_ME")

    EXTENSIONS: List[str] = ["_jwt"]
    BLUEPRINTS: List[str] = ["task", "auth"]


class Mock(Config):
    EXTENSIONS: List[str] = ["_jwt"]
    BLUEPRINTS: List[str] = ["task", "auth"]
