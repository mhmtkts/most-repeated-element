# Most Repeated Element Algorithm

This project implements an algorithm in Go that finds the most repeated element in a string array.

## Project Structure

```
recursive-number-algo
├── src
│   ├── main.go          # Entry point of the application
│   ├── analyzer 
│   │   └── analyzer .go # # Implementation of the finder function
│   └── test
│       └── analyzer_test.go # Unit tests for the finder function
├── go.mod               # Module definition for the Go project
├── .gitignore           # Files and directories to be ignored by Git
└── README.md            # Documentation for the project
```

## How to Run the Project

1. Ensure you have Go installed on your machine.
2. Clone the repository to your local machine.
3. Navigate to the project directory.
4. Run the application using the command:

```bash
go run src/main.go
```

## Algorithm Details

The `FindMostRepeated` function takes a string array as input and returns the most frequently occurring element. The implementation can be found in `src/analyzer/analyzer.go`.

Example:
```go
input := []string{"apple", "pie", "apple", "red", "red", "red"}
result := analyzer.FindMostRepeated(input)
// Output: "red"
```

## Testing

Unit tests for the function are located in  `analyzer_test.go.`. You can run the tests using the command:

```bash
go test src/tests -v
```

## Contribution

Feel free to contribute to this project by submitting issues or pull requests.