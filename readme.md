# Todo Fetcher

This Go program fetches TODO items from the [JSONPlaceholder](https://jsonplaceholder.typicode.com/) API and prints their titles and completion status. It uses goroutines to perform concurrent HTTP requests and handles errors gracefully.

## Features

* Fetches TODO items with even IDs from 2 to 40.
* Concurrent HTTP requests for better performance.
* Graceful error handling and reporting.

## Dependencies

This program uses the following Go standard libraries:

* `encoding/json` for JSON parsing.
* `fmt` for formatted I/O.
* `io/ioutil` for reading HTTP response bodies.
* `net/http` for making HTTP requests.
* `sync` for managing concurrency with wait groups.

## Installation

1. Ensure you have Go installed. If not, download and install it from the [official Go website](https://golang.org/dl/).
2. Clone the repository or copy the code into a new directory.

## Usage

1. Open a terminal and navigate to the directory containing the code.
2. Run the program with the following command:

   ```sh
   go run main.go
