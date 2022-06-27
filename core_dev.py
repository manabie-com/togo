import os

from celery_ import init as service_init
from flask_ import init as api_init

RUN_MODE = os.getenv("APP_CONFIG_DEFAULT", "Config")


def run_api():
    app = api_init.factory(RUN_MODE, "base")
    kwargs = {
        "host": app.config.get("HOST"),
        "port": int(app.config.get("HOST_PORT")),
        "debug": app.config.get("DEBUG", False),
        "use_reloader": app.config.get("USE_RELOADER", False),
        **app.config.get("SERVER_OPTIONS", {}),
    }
    app.run(**kwargs)


if __name__ == "__main__":
    with service_init.factory(RUN_MODE):
        run_api()
