from typing import Callable, Dict, List

from python_tests.utils.profile import Profile, Test
from python_tests.utils import config_file


def create_test(data: List[int], target: int) -> Callable:
    def work() -> List[int] | None:
        length = len(data)
        for i, a in enumerate(data):
            compliment = target - a

            for j in range(i + 1, length):
                if data[j] == compliment:
                    return [i, j]

        return None

    return work


if __name__ == "__main__":
    config = config_file.read_config()
    duration: int = config["duration"]
    test_configs: List[Dict] = config["test_configs"]

    tests: List[Test] = []

    for test_config in test_configs:
        data: List[int] = test_config["data"]
        target: int = test_config["target"]

        test = Test(len(data), create_test(data, target))
        tests.append(test)

    p = Profile("Brute force", duration, tests)
    p.run()
