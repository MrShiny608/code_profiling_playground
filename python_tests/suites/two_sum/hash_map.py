from typing import Callable, List

from python_tests.utils.profile import Profile
from python_tests.utils import config_file


def create_test(data: List[int], target: int) -> Callable:
    def work() -> List[int]:
        hashmap = {}
        for i, a in enumerate(data):
            hashmap[a] = i

        for i, b in enumerate(data):
            compliment = target - b

            if compliment in hashmap:
                return [i, hashmap[compliment]]

        return []

    return work


if __name__ == "__main__":
    config = config_file.read_config()
    data: List[int] = config["data"]
    target: int = config["target"]
    iterations: int = config["iterations"]

    p = Profile("Hashmap", iterations, create_test(data, target))
    p.run()
