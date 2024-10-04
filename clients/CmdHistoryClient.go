package clients

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

func LastCommand() []string {
	var lastCommands []string
	// Get the user's current shell from the environment variables
	shell := os.Getenv("SHELL")

	// Get the current user information to locate the history file
	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error getting user:", err)
		return lastCommands
	}

	var historyFilePath string

	// Detect the shell and set the appropriate history file path
	if strings.Contains(shell, "bash") {
		historyFilePath = filepath.Join(usr.HomeDir, ".bash_history")
	} else if strings.Contains(shell, "zsh") {
		historyFilePath = filepath.Join(usr.HomeDir, ".zsh_history")
	} else {
		fmt.Println("Unsupported shell or unable to detect shell.")
		return lastCommands
	}

	// Read the history file directly
	historyFile, err := os.ReadFile(historyFilePath)
	if err != nil {
		fmt.Println("Error reading history file:", err)
		return lastCommands
	}

	// Split the history into lines
	historyLines := strings.Split(string(historyFile), "\n")

	// Find the last 5 non-empty commands

	for i := len(historyLines) - 1; i >= 0 && len(lastCommands) < 5; i-- {
		trimmedLine := strings.TrimSpace(historyLines[i])
		if trimmedLine != "" {
			lastCommands = append(lastCommands, trimmedLine)
		}
	}

	return lastCommands
}
