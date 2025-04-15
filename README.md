# 🔬 Code Profiling Playground

[![Check me out on LinkedIn](https://img.shields.io/badge/LinkedIn-0077B5?logo=linkedin&logoColor=white)](https://www.linkedin.com/in/timothybrookes)
[![Blog](https://img.shields.io/badge/GitHub%20Pages-View%20Blog-green?logo=github&style=flat-square)](https://mrshiny608.github.io/MrShiny608/optimisation/2025/04/15/ProfilingCodeWithoutGettingTricked)

This repo is set up for **accurate performance profiling** of small blocks of code.

Functionally, it’s pretty straightforward — but as anyone who’s tried knows, **profiling correctly is deceptively hard**.

## 🔎 Key Files

- `*_tests/utils/profile.go` – minimal profiler loop
- `*_tests/utils/suite.go` – suite runner to manage testing in isolated executables
- `*_tests/suites/<test_name>/<implementation_name>` - where your code to be profiled should live, check the example of `two_sum`

You can plug your own logic into the test suites and generate comparative timings across multiple configurations. All the hard stuff (timing, iteration control, test isolation) is handled for you.

## 📖 Learn More

See the [blog post](https://mrshiny608.github.io/MrShiny608/optimisation/2025/04/15/ProfilingCodeWithoutGettingTricked) for an in-depth look at the _why_ behind this setup — including all the subtle ways benchmarking can lie to you, and how to get reliable, reproducible results instead.

## 🚀 Getting Started

### 1. Install Python & Dependencies

Ensure PIP is installed and in your `PATH`:

```bash
pip --version
```

Install [Pipenv](https://pipenv.pypa.io/):

```bash
pip install pipenv
```

Then install the dependencies:

```bash
pipenv install
```

### 2. Install Go

Ensure [Go](https://go.dev/doc/install) is installed and in your `PATH`:

```bash
go version
```

Then install any module dependencies (from the root of the Go test suite):

```bash
go mod tidy
```

### 3. Run a Test

From the root directory:

```bash
bash run.sh <test_name> -p   # Run Python test
bash run.sh <test_name> -g   # Run Go test
```

Replace `<test_name>` with the name of your test config.

### 4. Observe

Tests will always run first, there's no point profiling broken code, then profiling information will be output as it runs.

If you intend to leave the tests running for a long time on a remote machine, you can use `nohup` to prevent terminal hangup, this will by default direct all output to `nohup.out`.

```bash
nohup bash run.sh <test_name> -p -g &
```
