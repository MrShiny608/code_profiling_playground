import os

from python_tests.utils import config_file
from python_tests.utils.suite import Suite


def create_suite() -> Suite:
    file_path = os.path.abspath(__file__)
    current_directory = os.path.dirname(file_path)
    exclude_file = os.path.basename(file_path)

    return Suite(current_directory, exclude_file)


if __name__ == "__main__":
    # Prepare the config files
    duration = 60
    data_sizes = [(i + 1) * 10 for i in range(100)]

    for data_size in data_sizes:
        # We don't actually need any real data, as we're intentionally hitting
        # worst case scenario so don't generate a consistent dataset for use across
        # tests, just let them generate an array of zeros of the correct size

        # Set the target to an unachievable level so we can test the worse case scenario
        target = -1

        config_file.write_config(
            {
                "duration": duration,
                "target": target,
                "data_size": data_size,
            }
        )

        # Run the test suite
        suite = create_suite()
        suite.run()
