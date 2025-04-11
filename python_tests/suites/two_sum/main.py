import os
import random

from python_tests.utils import config_file
from python_tests.utils.suite import Suite


def create_suite() -> Suite:
    file_path = os.path.abspath(__file__)
    current_directory = os.path.dirname(file_path)
    exclude_file = os.path.basename(file_path)

    return Suite(current_directory, exclude_file)


if __name__ == "__main__":
    # Prepare the config files
    iterations = 100000000
    data_range = 1000
    data_size = 5
    data = random.sample(range(1, data_range), data_size)
    target = sum(random.sample(data, 2))

    config_file.write_config(
        {
            "iterations": iterations,
            "target": target,
            "data": data,
        }
    )

    # Run the test suite
    suite = create_suite()
    suite.run()
