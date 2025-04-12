from typing import Callable, Dict, List

from python_tests.utils.profile import Profile, Test
from python_tests.utils import config_file


def create_test(data: List[int], target: int) -> Callable:
    def work() -> List[int] | None:
        hashmap = {}
        for i, a in enumerate(data):
            compliment = target - a

            if compliment in hashmap:
                return [hashmap[compliment], i]

            hashmap[a] = i

        return None

    return work


if __name__ == "__main__":
    config = config_file.read_config()
    duration: int = config["duration"]
    test_configs: List[Dict] = config["test_configs"]

    tests: List[Test] = []
    for test_config in test_configs:
        target: int = test_config["target"]
        data_size: List[int] = test_config["data_size"]
        data: List[int] = [0] * data_size

        test = Test(len(data), create_test(data, target))
        tests.append(test)

    p = Profile("Hashmap", duration, tests)
    p.run()
