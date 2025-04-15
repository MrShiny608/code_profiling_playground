import os
import subprocess
import sys


class Suite(object):
    def __init__(self, directory: str, exclude_file: str):
        self.directory = directory
        self.exclude_file = exclude_file

    def run(self):
        # Iterate the directory, finding files viable for execution
        for root, _, files in os.walk(self.directory):
            for file in files:
                if file == self.exclude_file:
                    continue

                if not file.endswith(".py"):
                    continue

                if file.endswith("_test.py"):
                    continue

                # Run the file in a new subprocess to reduce liklihood of code cache misses
                file_path = os.path.join(root, file)
                try:
                    env = os.environ.copy()
                    env["PYTHONPATH"] = os.getcwd()

                    subprocess.run(
                        [sys.executable, file_path],
                        cwd=root,
                        env=env,
                        text=True,
                        check=True,
                        stdout=sys.stdout,
                        stderr=sys.stderr,
                    )
                except subprocess.CalledProcessError as e:
                    print(f"Error running {file_path}:\n{e.stderr}", file=sys.stderr)
