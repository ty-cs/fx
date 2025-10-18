# GEMINI.md - Project Context for Gemini

This file provides context about the `fx` project for the Gemini AI assistant.

## Project Overview

`fx` is a command-line toolkit for essential financial calculations, written in Go. It's designed to be a simple and efficient tool for developers and anyone interested in finance.

The project uses the following main technologies:
- **Go**: The programming language used for the project.
- **Cobra**: A popular library for creating powerful and modern CLI applications in Go.
- **Lipgloss**: A library for styling terminal output, used here to format tables and summaries.

The current features include:
- **Compound Growth Calculation**: The `compound` command calculates the future value of an investment with compound growth and displays a year-by-year breakdown.
- **Kelly Criterion Calculation**: The `kelly` command calculates the optimal bet size based on the Kelly Criterion formula.
- **Money Shorthand**: The tool accepts money values in shorthand, such as `10k` for 10,000, `1m` for 1,000,000, and `1b` for 1,000,000,000.

## Building and Running

### Prerequisites

- Go 1.25 or later is required.

### Building the Project

To build the project, run the following command from the root directory:

```sh
go build -v ./...
```

### Running the Application

After building, you can run the application using:

```sh
./fx [command] [flags]
```

Alternatively, you can run it without building using:

```sh
go run main.go [command] [flags]
```

### Running Tests

To run the tests, use the following command:

```sh
go test -v ./...
```

## Development Conventions

### Project Structure

The project follows a standard Go project layout:
- `cmd/`: Contains the main application code, with each file representing a command.
- `pkg/`: Contains the core business logic, such as the financial calculators.
- `util/`: Contains utility functions, such as the custom money flag parser.
- `main.go`: The entry point of the application.
- `go.mod`: Manages the project's dependencies.

### Coding Style

The code is formatted using the standard Go formatting tools (`gofmt`). It is expected that all new code contributions are formatted accordingly.

### Dependencies

The project uses Go modules to manage its dependencies. To add a new dependency, use `go get`.
