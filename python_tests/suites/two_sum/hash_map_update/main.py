from typing import Callable, List, Tuple, Generator

from python_tests.utils.profile import Profile, Test
from python_tests.utils import config_file


def enumerate_to_dict(data: List[int]) -> Generator[Tuple[int, int], None, None]:
    for i, a in enumerate(data):
        yield a, i


def create_test() -> Callable:
    def work(data: List[int], target: int) -> List[int] | None:
        hashmap = {}
        hashmap.update(enumerate_to_dict(data))

        for i, a in enumerate(data):
            complement = target - a

            if complement in hashmap:
                return [hashmap[complement], i]

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

    p = Profile("Hashmap (update)", duration, test)
    p.run()
