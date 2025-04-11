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
    duration = 60 * 2
    data_sizes = [
        10,
        100,
        1000,
        10000,
        100000,
    ]
    data_range = max(data_sizes)

    numbers = range(1, data_range + 1)

    test_configs = []

    for data_size in data_sizes:
        data = random.sample(numbers, data_size)

        # Set the target to an unachievable level so we can test
        # the worse case scenario
        target = data_range + 1

        test_configs.append(
            {
                "target": target,
                "data": data,
            }
        )

    config_file.write_config(
        {
            "duration": duration,
            "test_configs": test_configs,
        }
    )

    # Run the test suite
    suite = create_suite()
    suite.run()
