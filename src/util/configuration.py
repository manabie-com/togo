""" Handle configuration for the program """
import sys
from yaml import safe_load, YAMLError
import os


def load_config() -> dict:
    """
    Load configuration from config/config.toml file
    """
    config_fp = os.getenv('config_file', default='config/config.yml')
    config_path = os.path.join(os.getcwd(), config_fp)
    try:
        with open(config_path, "r") as reader:
            CONFIG = safe_load(reader.read())
            return CONFIG
    except YAMLError as e:
        print(e)
        print("Error in loading configuration! Please check configuration file.")
        sys.exit(0)
