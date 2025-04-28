from typing import Callable, List

from python_tests.utils.profile import Profile, Test
from python_tests.utils import config_file


def create_test() -> Callable:
    hashmap = {}

    def work(data: List[int], target: int) -> List[int] | None:
        hashmap.clear()

        for i, a in enumerate(data):
            complement = target - a

            if complement in hashmap:
                return [hashmap[complement], i]

            hashmap[a] = i

        return None

    return work


if __name__ == "__main__":
    config = config_file.read_config()
    duration: int = config["duration"]
    target: int = config["target"]
    data_size: List[int] = config["data_size"]

    data: List[int] = [i + 1 for i in range(data_size)]

    test = Test(
        data_size,
        create_test(),
        args=(data, target),
        kwargs={},
    )

    p = Profile("Hashmap (cache clear)", duration, test)
    p.run()
