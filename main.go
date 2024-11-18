package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the terminal!")
	fmt.Println("Type 'help' to see the available commands.")

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("help - Show this help message")
			fmt.Println("exit - Exit the game")

		case "exit":
			fmt.Println("Goodbye")
			return

		default:
			fmt.Println("Unknown commands:", input)
		}
	}
}
