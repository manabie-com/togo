""" Setup logger for program """
import os
import logging
import sys

from logging import Logger, Filter, LogRecord
from logging.config import dictConfig
from yaml import safe_load, YAMLError


# class AccessFilter(Filter):
# 	def filter(self, record: LogRecord):
# 		msg = record.msg
# 		level = record.levelname
# 		if level == "INFO" and "access" in msg.lower():
# 			return 1
# 		return 0

class ErrorFilter(Filter):

    def filter(self, record: LogRecord):
        msg = record.msg
        level = record.levelname
        if level == "ERROR":
            return 1
        return 0


class InfoFilter(Filter):

    def filter(self, record: LogRecord):
        msg = record.msg
        level = record.levelname
        if level == "INFO":
            return 1
        return 0


def setup_logger() -> Logger:
    """
    Create a logger with basic setup : info, error, access log files from `config/logger.yml` file
    """
    with open(os.path.join(os.getcwd(), "config", "logger.yml")) as reader:
        try:
            d = safe_load(reader)
            d["filters"]["inf"] = {
                "()": InfoFilter
            }
            d["filters"]["err"] = {
                "()": ErrorFilter
            }
            dictConfig(d)
            return logging.getLogger("program")
        except YAMLError as e:
            print("Logger configuration is invalid, check logger configuration!")
            print(e)
            sys.exit(1)
        except Exception as e:
            print(e)
            sys.exit(1)


logger = setup_logger()
