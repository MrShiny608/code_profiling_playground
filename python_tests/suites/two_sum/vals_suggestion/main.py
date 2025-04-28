from typing import Callable, List

from python_tests.utils.profile import Profile, Test
from python_tests.utils import config_file


def create_test() -> Callable:
    def work(data: List[int], target: int) -> List[int] | None:
        # Add indices to the data
        data = [(data[i], i) for i in range(len(data))]

        # Sort it by value, preserving the indices
        data.sort(key=lambda elem: elem[0])

        i = 0
        j = len(data) - 1
        while i <= j:
            total = data[i][0] + data[j][0]
            if total == target:
                return [data[i][1], data[j][1]]
            elif total < target:
                i += 1
            elif total > target:
                j -= 1

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

    p = Profile("Val's Suggestion", duration, test)
    p.run()
