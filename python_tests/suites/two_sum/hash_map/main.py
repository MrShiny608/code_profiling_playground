from typing import Callable, List

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
    target: int = config["target"]
    data_size: List[int] = config["data_size"]

    data: List[int] = [i + 1 for i in range(data_size)]

    test = Test(data_size, create_test(data, target))

    p = Profile("Hashmap", duration, test)
    p.run()
