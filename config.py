import os
basedir = os.path.abspath(os.path.dirname(__file__))

class Config:
    DEBUG = False
    SSL_REDIRECT = False
    SQLALCHEMY_TRACK_MODIFICATIONS = False
    DB_HOST = os.environ.get("DB_HOST")
    DB_USER = os.environ.get("DB_USER") or "dummy"
    DB_NAME = os.environ.get("DB_NAME") or "todo"
    DB_PORT = os.environ.get("DB_PORT") or 5432
    DB_PASSWORD = os.environ.get("DB_PASSWORD") or "dummypassword"

    SQLALCHEMY_DATABASE_URI = os.environ.get('DEV_DATABASE_URL') or \
        f"postgresql://{DB_USER}:{DB_PASSWORD}@postgres/{DB_NAME}"

    @staticmethod
    def init_app(app):
        pass

class DevelopmentConfig(Config):
    DEBUG = True

class TestingConfig(Config):
    TESTING = True


config={
    "default": DevelopmentConfig,
    "development": DevelopmentConfig,
    "testing": TestingConfig,
}