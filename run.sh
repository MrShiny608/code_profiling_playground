#!/bin/sh

# Read the test suite and flags from the command line
if [ -z "${1}" ]; then
  echo "Usage: ${0} <test_suite> [-p] [-g]"
  exit 1
fi
TEST_SUITE="${1}"
shift

RUN_PYTHON=false
RUN_GO=false

# Parse flags
while [ "$#" -gt 0 ]; do
  case "$1" in
    -p)
      RUN_PYTHON=true
      ;;
    -g)
      RUN_GO=true
      ;;
    *)
      echo "Unknown flag: $1"
      echo "Usage: $0 <test_suite> [-p] [-g]"
      exit 1
      ;;
  esac
  shift
done

# Run the python tests if -p flag is set
if [ "${RUN_PYTHON}" = "true" ]; then
  python python_tests/suites/${TEST_SUITE}/main.py
fi

# Run the go tests if -g flag is set
if [ "${RUN_GO}" = "true" ]; then
  go build -o test_suite.out go_tests/suites/${TEST_SUITE}/main.go
  chmod +x test_suite.out
  ./test_suite.out
  rm test_suite.out
fi
