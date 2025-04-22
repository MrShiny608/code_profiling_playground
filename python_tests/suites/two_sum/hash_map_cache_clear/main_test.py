from typing import List
from pytest_subtests import SubTests

from python_tests.suites.two_sum.hash_map_cache_clear.main import create_test


class Args(object):
    def __init__(self, data: List[int], target: int):
        self.data = data
        self.target = target


class Result(object):
    def __init__(self, indices: List[int] | None):
        self.indices = indices


class Config(object):
    def __init__(self, name: str, args: Args, result: Result):
        self.name = name
        self.args = args
        self.result = result


def test_create_test(subtests: SubTests):
    configs: List[Config] = [
        Config(
            "returns the correct indices",
            Args(
                data=[1, 2, 3, 4],
                target=4,
            ),
            Result(
                indices=[0, 2],
            ),
        ),
        Config(
            "returns none when target isn't achievable",
            Args(
                data=[1, 2, 3, 4],
                target=-1,
            ),
            Result(
                indices=None,
            ),
        ),
    ]

    for config in configs:
        with subtests.test(config.name):
            # Arrange
            args = config.args
            result = config.result
            work = create_test()

            # Act
            indices = work(args.data, args.target)

            # Assert
            assert indices == result.indices
