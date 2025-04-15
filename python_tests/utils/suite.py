import os
import subprocess
import sys


class Suite(object):
    def __init__(self, directory: str, exclude_file: str):
        self.directory = directory
        self.exclude_file = exclude_file

    def run(self):
        # Iterate the directory, finding files viable for execution
        for root, directories, _ in os.walk(self.directory):
            for directory in directories:
                main_file_path = os.path.join(root, directory, "main.py")
                if not os.path.isfile(main_file_path):
                    continue

                # Run the file in a new subprocess to reduce likelihood of code cache misses
                try:
                    env = os.environ.copy()
                    env["PYTHONPATH"] = os.getcwd()

                    subprocess.run(
                        [sys.executable, main_file_path],
                        cwd=root,
                        env=env,
                        text=True,
                        check=True,
                        stdout=sys.stdout,
                        stderr=sys.stderr,
                    )
                except subprocess.CalledProcessError as e:
                    print(f"Error running {main_file_path}:\n{e.stderr}", file=sys.stderr)
