import time
from typing import Callable, List


class Test(object):
    def __init__(self, n: int, work: Callable):
        self.n = n
        self.work = work


class Profile(object):
    def __init__(self, name: str, duration: int, tests: List[Test]):
        self.name = name
        self.tests = tests
        self.duration = duration

    def run(self) -> None:
        print(f"Running test in Python\n - name: {self.name}")

        for test in self.tests:
            print(f" - N={test.n}", end="")

            # Pre-warm any online/adaptive optimisation (e.g. Specializing Adaptive Interpreter, pypy, etc)
            for _ in range(10000):
                test.work()

            # Run the real test
            iterations = 0
            start_time = time.time()

            while (time.time() - start_time) < self.duration:
                test.work()
                iterations += 1

            end_time = time.time()
            duration = end_time - start_time

            print(f" {((duration/iterations)*1e9):.0f}ns")
