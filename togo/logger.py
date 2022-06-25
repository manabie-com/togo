import logging

from togo.settings import (
    TOGO_TASK_PICK_LIMIT_LOGGER,
    TOGO_TASK_PICK_LIMIT_MANUALLY_LOGGER
)


def get_logger(name=None):
    return logging.getLogger(name=name)


togo_task_pick_limit_logger = get_logger(TOGO_TASK_PICK_LIMIT_LOGGER)
togo_task_pick_limit_manually_logger = get_logger(TOGO_TASK_PICK_LIMIT_MANUALLY_LOGGER)
default_logger = get_logger("django")
