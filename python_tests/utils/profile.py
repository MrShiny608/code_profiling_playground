import time
from typing import Callable


class Profile(object):
    def __init__(self, name: str, count: int, test: Callable):
        self.name = name
        self.test = test
        self.count = count
        self.duration = 0

    def run(self) -> None:
        print(f"Running test\n - name: {self.name}")

        # Pre-warm any online/adaptive optimisation (e.g. Specializing Adaptive Interpreter, pypy, etc)
        for _ in range(100000):
            self.test()

        # Run the real test
        start_time = time.time()

        for _ in range(self.count):
            self.test()

        end_time = time.time()
        self.duration = end_time - start_time

        print(f" - duration: {self.duration}s\n - iterations: {self.count}\n - average: {(self.duration/self.count)*1e9}ns")
