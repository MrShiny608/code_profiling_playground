import time
from typing import Callable


class Profile(object):
    def __init__(self, name: str, iterations: int, test: Callable):
        self.name = name
        self.test = test
        self.iterations = iterations
        self.duration = 0

    def run(self) -> None:
        print(f"Running test in Python\n - name: {self.name}")

        # Pre-warm any online/adaptive optimisation (e.g. Specializing Adaptive Interpreter, pypy, etc)
        for _ in range(100000):
            self.test()

        # Run the real test
        start_time = time.time()

        for _ in range(self.iterations):
            self.test()

        end_time = time.time()
        self.duration = end_time - start_time

        print(f" - duration: {self.duration}s\n - iterations: {self.iterations}\n - average: {(self.duration/self.iterations)*1e9}ns")
