# fx

A command-line toolkit for essential financial calculations.

`fx` is a simple and efficient tool for developers and anyone interested in finance, providing quick calculations for compound growth and Kelly Criterion.

## Installation

### Prerequisites

- Go 1.25 or later is required.

### Using `go install`

You can install `fx` directly using `go install`:

```sh
go install github.com/ty-cs/fx@latest
```

### Building from Source

To build the project, clone the repository and run the following command from the root directory:

```sh
go build -v ./...
```

## Usage

After building, you can run the application using:

```sh
fx [command] [flags]
```

Alternatively, you can run it without building using:

```sh
go run main.go [command] [flags]
```

### Commands

#### `compound`

Calculates the future value of an investment with compound growth and displays a year-by-year breakdown.

**Flags:**

- `-p`, `--principal`: The initial investment principal (required).
- `-c`, `--contribution`: The annual contribution.
- `-g`, `--growth`: The annual percentage increase in the contribution (e.g., 3 for 3%).
- `-y`, `--years`: The total investment duration in years (required).
- `-r`, `--rate`: The expected annual rate of return (e.g., 8 for 8%).

**Example:**

```sh
fx compound --principal 10k --contribution 5k --growth 3 --years 20 --rate 8
```

#### `kelly`

Calculates the optimal bet size based on the Kelly Criterion formula.

**Flags:**

- `-p`, `--win-probability`: Win probability (e.g., 0.6 for 60%) (required).
- `-o`, `--net-odds`: Net odds (e.g., 0.2 for 20% return) (required).

**Example:**

```sh
fx kelly --win-probability 0.6 --net-odds 0.2
```

## Features

- **Compound Growth Calculation**: The `compound` command calculates the future value of an investment with compound growth and displays a year-by-year breakdown.
- **Kelly Criterion Calculation**: The `kelly` command calculates the optimal bet size based on the Kelly Criterion formula.
- **Money Shorthand**: The tool accepts money values in shorthand, such as `10k` for 10,000, `1m` for 1,000,000, and `1b` for 1,000,000,000.

## Dependencies

- [Cobra](https://github.com/spf13/cobra): A popular library for creating powerful and modern CLI applications in Go.
- [Lipgloss](https://github.com/charmbracelet/lipgloss): A library for styling terminal output.

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.