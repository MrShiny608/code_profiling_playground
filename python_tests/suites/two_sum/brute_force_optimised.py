from typing import Callable, List

from python_tests.utils.profile import Profile
from python_tests.utils import config_file


def create_test(data: List[int], target: int) -> Callable:
    def work() -> List[int]:
        for i, a in enumerate(data):
            compliment = target - a

            for j, b in enumerate(data, i):
                if b == compliment:
                    return [i, j]

        return []

    return work
    # def work() -> List[int]:
    #     length = len(data)
    #     for i in range(length):
    #         compliment = target - data[i]

    #         for j in range(i, length):
    #             if data[j] == compliment:
    #                 return [i, j]

    #     return []

    # return work


if __name__ == "__main__":
    config = config_file.read_config()
    data: List[int] = config["data"]
    target: int = config["target"]
    iterations: int = config["iterations"]

    p = Profile("Brute force (optimised)", iterations, create_test(data, target))
    p.run()
