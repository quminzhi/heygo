# Sorting Algorithm Project

## Overview

This project is designed to implement and test two common sorting algorithms: Quick Sort and Merge Sort. It includes a set of unit tests and benchmarks to ensure the correctness and efficiency of the algorithms. Additionally, it provides functionality to generate code coverage reports.

## Project Structure
```plaintext
.
├── go.mod
├── main.go
└── sort
    ├── sort.go
    └── sort_test.go
```
- `go.mod`: Defines the Go module and its dependencies.
- `main.go`: The entry - point of the project.
- `sort/`: A directory containing the sorting algorithm implementations and test cases.
    - `sort.go`: Contains the implementation of Quick Sort and Merge Sort.
    - `sort_test.go`: Includes test and benchmark functions for the sorting algorithms.

## Prerequisites

- Go programming language (version 1.16 or higher recommended).

## Installation

Clone the project to your local machine:
```bash
git clone <repository_url>
cd <project_directory>
```

## Usage

### Running Tests

You can use the `Makefile` to execute various tasks related to testing and benchmarking.

#### Normal Tests

To run normal tests that verify the correctness of the sorting algorithms, use the following command:
```bash
make test
```
This command runs all the test functions in the `sort` directory.

#### Tests with Coverage

To run tests and display the code coverage percentage, use:
```bash
make test-cover
```
This will show how much of the code is covered by the tests.

#### Benchmarking

To run benchmarks and measure the performance of the sorting algorithms, use:
```bash
make benchmark
```
This will output metrics such as the average time per operation and the number of operations per second.

### Coverage Report

#### Generate Coverage Report
To generate a coverage report file, use:
```bash
make cover-report
```
This will create a file named `coverage.out` that contains detailed coverage information.

#### Visualize Coverage Report
To view the coverage report in a more visual way, use:
```bash
make view-cover
```
This will open an HTML page in your default browser, showing which parts of the code are covered by tests.

### Cleanup

To remove the generated coverage report file, use:
```bash
make clean
```

### Help
If you need a list of all available commands, use:
```bash
make help
```
This will display a summary of all the commands you can use with the `Makefile`.

## Contributing
If you want to contribute to this project, please fork the repository and create a pull request with your changes.
