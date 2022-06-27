import os


# base config class; extend it to your needs.
class Config(object):
    PROJECT_PATH = os.getenv("PROJECT_PATH", "")
    APP_DB_HOST = os.getenv("APP_DB_HOST", "localhost")
    APP_DB_PORT = os.getenv("APP_DB_PORT", "27017")
    APP_DB_NAME = os.getenv("APP_DB_NAME", "local")
    APP_DB_USER = os.getenv("APP_DB_USER", "")
    APP_DB_PASS = os.getenv("APP_DB_PASS", "")

    if APP_DB_USER and APP_DB_PASS:
        APP_DB_URL = f"mongodb://{APP_DB_USER}:{ APP_DB_PASS}@{APP_DB_HOST}:{APP_DB_PORT}/{APP_DB_NAME}"
    else:
        APP_DB_URL = f"mongodb://{APP_DB_HOST}:{APP_DB_PORT}/{APP_DB_NAME}"

    WORKER_POOL = "prefork"  # "prefork" celery defaults value
    USER_REQUEST_LIMIT = 5


class Mock(Config):
    APP_DB_HOST = "localhost"
    APP_DB_PORT = "27017"
    APP_DB_NAME = "local"
    APP_DB_USER = ""
    APP_DB_PASS = ""

    if APP_DB_USER and APP_DB_PASS:
        APP_DB_URL = f"mongomock://{APP_DB_USER}:{ APP_DB_PASS}@{APP_DB_HOST}:{APP_DB_PORT}/{APP_DB_NAME}"
    else:
        APP_DB_URL = f"mongomock://{APP_DB_HOST}:{APP_DB_PORT}/{APP_DB_NAME}"
