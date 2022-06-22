import logging


from ..settings import (
    TOGO_TASK_PICK_LIMIT_LOGGER
)


def get_logger(name=None):
    return logging.getLogger(name=name)


togo_task_pick_limit_logger = get_logger(TOGO_TASK_PICK_LIMIT_LOGGER)
default_logger = get_logger("django")