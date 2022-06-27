import os
import subprocess
import time

import six
from database.models.service import Service
from helpers import common
from mongoengine.connection import connect

string_types = six.string_types


class ServiceInit:
    def __init__(self, config):
        self.config = config
        self.path = "celery_.services"

    def __enter__(self):
        connect(host=self.config.APP_DB_URL, alias="app-db")
        self.services = Service.objects()
        self.running = []
        self.boot_all()

    def create_service(self):
        for service in self.services:
            if service.type == "beat":
                self.running.append(
                    subprocess.Popen(
                        [
                            "celery",
                            "-A",
                            f"{self.path}.{service.app}",
                            "beat",
                            "-s",
                            f"celerybeat-schedule-{service.app}",
                            "-l",
                            service.loglevel,
                            f"--pidfile={service.app}.pid",
                        ]
                    )
                )
            self.running.append(
                subprocess.Popen(
                    [
                        "celery",
                        "-A",
                        f"{self.path}.{service.app}",
                        "worker",
                        "-l",
                        service.loglevel,
                        "-Q",
                        service.queues,
                        f"--concurrency={service.concurrency}",
                        "-n",
                        f"wkr{service.code}@%h",
                        "--pool",
                        "solo",
                    ]
                )
            )
            time.sleep(1)

    def boot_all(self):
        self.create_service()

    def __exit__(self, *_):
        for service in self.running:
            service.kill()


def factory(config):
    config = common.config_str_to_obj("celery_.config", config)
    app = ServiceInit(config)
    return app
