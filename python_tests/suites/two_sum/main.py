import os

from python_tests.utils import config_file
from python_tests.utils.suite import Suite
import random


def create_suite() -> Suite:
    file_path = os.path.abspath(__file__)
    current_directory = os.path.dirname(file_path)
    exclude_file = os.path.basename(file_path)

    return Suite(current_directory, exclude_file)


if __name__ == "__main__":
    suite = create_suite()

    data_size = 5
    data = random.sample(range(1, 1000), data_size)
    target = sum(random.sample(data, 2))

    config_file.write_config(
        {
            "iterations": 100000000,
            "target": target,
            "data": data,
        }
    )

    suite.run()
