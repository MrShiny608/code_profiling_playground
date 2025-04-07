from typing import Callable, List

from python_tests.utils.profile import Profile
from python_tests.utils import config_file


def create_test(data: List[int], target: int) -> Callable:
    def work() -> List[int]:
        hashmap = {}
        for i, a in enumerate(data):
            compliment = target - a

            if compliment in hashmap:
                return [i, hashmap[compliment]]

            hashmap[a] = i

        return []

    return work


if __name__ == "__main__":
    config = config_file.read_config()
    data: List[int] = config["data"]
    target: int = config["target"]
    iterations: int = config["iterations"]

    p = Profile("Hashmap (optimised)", iterations, create_test(data, target))
    p.run()
