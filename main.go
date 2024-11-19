package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	helpCommand = "help"
	exitCommand = "exit"
)

func main() {
	clearScreen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the terminal!")
	fmt.Println("Type 'help' to see the available commands.")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		input = strings.TrimSpace(input)
		handleCommand(input)
	}
}

func handleCommand(input string) {
	switch input {
	case helpCommand:
		fmt.Println("Available commands:")
		fmt.Println("help - Show this help message")
		fmt.Println("exit - Exit the game")

	case exitCommand:
		fmt.Println("Goodbye")
		os.Exit(0)

	default:
		fmt.Println("Unknown command:", input)
	}
}

func clearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
