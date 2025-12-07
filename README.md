# Expense Tracker

A lightweight command-line application for tracking personal expenses, built with Go. Simple, fast, and no database required—just a JSON file and your terminal.

## Features

- **Add expenses** with descriptions and amounts
- **List all expenses** in a clean, formatted table
- **Delete expenses** by ID
- **View spending summary** with total calculations
- **Persistent storage** using a JSON file in your home directory
- **Comprehensive test coverage** for core functionality

## Prerequisites

- Go 1.16 or higher
- Basic familiarity with the command line

## Installation

### From Source

1. Clone the repository:
```bash
git clone https://github.com/ayoblt/expense-tracker.git
cd expense-tracker
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o expense-tracker
```

4. (Optional) Move the binary to your PATH:
```bash
# On Linux/macOS
sudo mv expense-tracker /usr/local/bin/

# On Windows, add the directory to your PATH environment variable
```

## Usage

### Add an Expense

Record a new expense with a description and amount:

```bash
expense-tracker add --description "Lunch at Chicken Republic" --amount 3500
# Output: Expense added successfully (ID: 1)
```

Short flags are also supported:
```bash
expense-tracker add -d "Transport fare" -a 500
```

### List All Expenses

View all recorded expenses in a formatted table:

```bash
expense-tracker list
```

Output:
```
ID  Description                    Amount
1   Lunch at Chicken Republic      ₦3500
2   Transport fare                 ₦500
```

### Delete an Expense

Remove an expense by its ID:

```bash
expense-tracker delete --id 2
# Output: Expense deleted successfully
```

### View Summary

Calculate the total of all expenses:

```bash
expense-tracker summary
# Output: Total expenses: ₦4000
```

### Default Behavior

Running the command without any subcommand will display the expense list:

```bash
expense-tracker
```

## Data Storage

Expenses are stored in a JSON file located at `~/expenses.json`. The file is created automatically when you add your first expense.

**Example structure:**
```json
{
  "expenses": [
    {
      "id": 1,
      "description": "Lunch",
      "amount": 3500
    }
  ],
  "next_id": 2
}
```

## Project Structure

```
expense-tracker/
├── cmd/
│   ├── root.go       # Root command and CLI setup
│   ├── add.go        # Add expense command
│   ├── list.go       # List expenses command
│   ├── delete.go     # Delete expense command
│   └── summary.go    # Summary command
├── internal/
│   ├── expense/
│   │   ├── expense.go      # Core expense logic
│   │   └── expense_test.go # Expense tests
│   └── storage/
│       ├── storage.go      # File I/O operations
│       └── storage_test.go # Storage tests
├── main.go           # Application entry point
├── go.mod            # Go module definition
└── README.md         # Project documentation
```

## Running Tests

Run all tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run tests for a specific package:
```bash
go test ./internal/expense
go test ./internal/storage
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework for Go

## Development

### Adding New Features

1. Create a new command file in the `cmd/` directory
2. Implement the command logic using Cobra
3. Add the command to `rootCmd` in the `init()` function
4. Write tests for any new functionality

### Code Organization

- **cmd/**: Contains all CLI command definitions
- **internal/expense/**: Core business logic for expense management
- **internal/storage/**: Handles reading/writing to the JSON file

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

Please ensure your code:
- Passes all existing tests
- Includes tests for new features
- Follows Go best practices and formatting (`go fmt`)

## License

This project is open source and available under the MIT License.

## Author

Built by [ayoblt](https://github.com/ayoblt)

## Roadmap

Potential future enhancements:
- [ ] Filter expenses by date range
- [ ] Category-based expense tracking
- [ ] Export to CSV/PDF
- [ ] Monthly/weekly summaries
- [ ] Budget alerts and limits
- [ ] Multi-currency support

## Acknowledgments

This project was inspired by the need for a simple, terminal-based expense tracker that doesn't require complex setup or external databases.
