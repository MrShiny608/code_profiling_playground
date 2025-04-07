import inspect
import yaml
import os
from typing import Any, Dict


def read_config() -> Dict[str, Any]:
    caller_frame = inspect.stack()[1]
    file_path = os.path.abspath(caller_frame.filename)
    current_directory = os.path.dirname(file_path)
    config_file = os.path.join(current_directory, "config.yaml")

    with open(config_file, "r") as file:
        return yaml.safe_load(file)


def write_config(data: Dict[str, Any]) -> None:
    caller_frame = inspect.stack()[1]
    file_path = os.path.abspath(caller_frame.filename)
    current_directory = os.path.dirname(file_path)
    config_file = os.path.join(current_directory, "config.yaml")
    with open(config_file, "w+") as file:
        yaml.safe_dump(data, file)
