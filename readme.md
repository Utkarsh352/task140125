### Token Manager Project

The **Token Manager** is a Go-based application for managing token usage efficiently. It provides features to simulate operations on tokens, view token usage statistics, reset tokens, and more via a command-line interface (CLI).

---

## Features

1. **Simulate Token Usage**  
   Simulate a specified number of operations where tokens are selected and their usage counts are incremented.
   
2. **View Token Statistics**  
   Display the usage count of all tokens and identify the least-used tokens.

3. **Reset All Tokens**  
   Reset the usage count of all tokens to zero, either automatically every 24 hours or manually through the CLI.

4. **Interactive CLI**  
   An easy-to-use menu-driven interface to access all functionality.

5. **Concurrency-Safe Management**  
   Token operations are thread-safe, ensuring correctness in concurrent scenarios.

---

## Project Structure

```
gogo/
├── cli/                   # Command-line interface logic
│   └── cli.go             # Interactive CLI implementation
├── model/                 # Models and definitions
│   └── token.go           # Token struct definition
├── tokenManager/          # Core token management logic
│   └── tokenmanager.go    # TokenManager implementation
├── main.go                # Entry point for the application
├── go.mod                 # Go module file
└── readme.md              # Project documentation
```

---

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/utkarsh352/task140125.git
   cd task140125
   ```

2. Ensure you have Go installed (Go 1.18+ recommended).  

3. Build the application:
   ```sh
   go build -o tokenManager
   ```

4. Run the application:
   ```sh
   ./tokenManager
   ```

---

## Usage

Once the application is started, you will see the interactive menu:

```
=== Welcome to the Token Manager CLI ===

=== Token Manager Menu ===
1. Simulate Token Usage
2. View Token Usage Statistics
3. Reset All Tokens
4. Exit

Enter your choice:
```

### Menu Options

1. **Simulate Token Usage**  
   Select this option to simulate a number of operations:
   - Enter the number of operations to simulate.
   - The application will update the token usage counts based on random selections.

2. **View Token Usage Statistics**  
   Displays:
   - Usage counts for all tokens.
   - The least-used tokens with their usage counts.

3. **Reset All Tokens**  
   Resets all token usage counts to zero manually.

4. **Exit**  
   Cleanly exits the program.

---

## Example

### Simulate Token Usage:
```sh
Enter your choice: 1
Enter the number of operations to simulate: 500
Simulation complete: 500 operations performed.
```

### View Token Usage:
```sh
Enter your choice: 2

=== Token Usage Counts ===
Token 1: 5 uses
Token 2: 3 uses
...

=== Least Used Token(s) (0 use(s)) ===
Token 8
Token 12
...
```

### Reset All Tokens:
```sh
Enter your choice: 3
All token usage counts have been reset.
```

### Exit:
```sh
Enter your choice: 4
Exiting Token Manager. Goodbye!
```

---

## Testing

To test the program:

1. Run the application.
2. Navigate through the CLI options.
3. Confirm that statistics, reset operations, and token usage behave as expected.

---

## Future Enhancements

- **Token Expiry**: Add a feature to expire tokens after a certain duration of inactivity.
- **Database Integration**: Store tokens and usage data in a database for persistence.
- **Web Interface**: Provide a web-based UI for interacting with the token manager.

---

## License

This project is open-source and available under the [MIT License](LICENSE).  
