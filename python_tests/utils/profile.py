import time
from typing import Any, Callable, Dict, Tuple


class Test(object):
    def __init__(self, n: int, work: Callable, args: Tuple[Any, ...], kwargs: Dict[str, Any]):
        self.n = n
        self.work = work
        self.args = args
        self.kwargs = kwargs


class Profile(object):
    def __init__(self, name: str, duration: int, test: Test):
        self.name = name
        self.test = test
        self.duration = duration

    def run(self) -> None:
        print(f"[Python] {self.name} - N={self.test.n}", end="", flush=True)

        # Pre-warm any online/adaptive optimisation (e.g. Specializing Adaptive Interpreter, pypy, etc)
        test = self.test
        for _ in range(10000):
            test.work(*test.args, **test.kwargs)

        # Run the real test
        iterations = 0
        start_time = time.time()

        while (time.time() - start_time) < self.duration:
            test.work(*test.args, **test.kwargs)
            iterations += 1

        end_time = time.time()
        duration = end_time - start_time

        print(f" {((duration/iterations)*1e9):.0f}ns", flush=True)
