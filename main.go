package main

import (
	"github.com/utkarsh352/task140125/cli"
	"github.com/utkarsh352/task140125/tokenManager"
)

func main() {
	// Initialize the TokenManager
	tokenManager := tokenManager.NewTokenManager()
	defer tokenManager.Stop()

	// Initialize and run the CLI
	cli := cli.NewCLI(tokenManager)
	cli.Run()
}
