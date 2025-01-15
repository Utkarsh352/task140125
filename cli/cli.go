package cli

import (
	"fmt"

	"github.com/utkarsh352/task140125/tokenManager"
)

type CLI struct {
	TokenManager *tokenManager.TokenManager
}

// To initialize a new CLI instance.
func NewCLI(tm *tokenManager.TokenManager) *CLI {
	return &CLI{
		TokenManager: tm,
	}
}

// To start the interactive CLI session.
func (cli *CLI) Run() {
	fmt.Println("=== Welcome to the Token Manager  (Created by Utkarsh352)===")
	for {
		cli.displayMenu()
		var choice int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			cli.simulateTokenUsage()
		case 2:
			cli.displayTokenUsage()
		case 3:
			cli.resetTokens()
		case 4:
			cli.exit()
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

// To show the main menu options.
func (cli *CLI) displayMenu() {
	fmt.Println("\n=== Token Manager Menu ===")
	fmt.Println("1. Simulate Token Usage")
	fmt.Println("2. View Token Usage Statistics")
	fmt.Println("3. Reset All Tokens")
	fmt.Println("4. Exit")
}

// To simulate token usage for a given number of operations.
func (cli *CLI) simulateTokenUsage() {
	var operations int
	fmt.Print("Enter the number of operations to simulate: ")
	_, err := fmt.Scan(&operations)
	if err != nil || operations <= 0 {
		fmt.Println("Invalid input. Please enter a positive integer.")
		return
	}

	cli.TokenManager.Simulate(operations)
	fmt.Printf("Simulation complete: %d operations performed.\n", operations)
}

// To display statistics of all tokens.
func (cli *CLI) displayTokenUsage() {
	leastUsedTokens, leastUsage := cli.TokenManager.GetStatistics()

	fmt.Println("\n=== Token Usage Counts ===")
	for _, token := range cli.TokenManager.GetAllTokens() {
		fmt.Printf("%s: %d uses\n", token.ID, token.Usage)
	}

	fmt.Printf("\n=== Least Used Token(s) (%d use(s)) ===\n", leastUsage)
	for _, token := range leastUsedTokens {
		fmt.Printf("%s\n", token.ID)
	}
}


// To reset the usage count of all tokens.
func (cli *CLI) resetTokens() {
	cli.TokenManager.ResetTokens()
	fmt.Println("All token usage counts have been reset.")
}

// To exit the program.
func (cli *CLI) exit() {
	fmt.Println("Exiting Token Manager (Credit: Utkarsh352)")
}
