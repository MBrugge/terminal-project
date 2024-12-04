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
	helpCommand  = "help"
	exitCommand  = "exit"
	resetCommand = "reset"

	// Colours
	colourReset  = "\033[0m"
	colourRed    = "\033[31m"
	colourYellow = "\033[33m"
	colourCyan   = "\033[36m"
)

func main() {
	clearScreen()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(colourText("Welcome to the terminal!", colourCyan))
	fmt.Println(colourText("Type 'help' to see the available commands.", colourYellow))

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
		fmt.Println(colourText("Available commands:", colourCyan))
		fmt.Println(colourText("help - Show this help message", colourYellow))
		fmt.Println(colourText("exit - Exit the program", colourYellow))
		fmt.Println(colourText("reset - Reset the program", colourYellow))

	case exitCommand:
		fmt.Println(colourText("Goodbye", colourRed))
		os.Exit(0)

	case resetCommand:
		fmt.Println(colourText("Restarting the program", colourRed))
		restartProgram()

	default:
		fmt.Println("Unknown command:", input)
	}
}

func restartProgram() {
	cmd := exec.Command(os.Args[0])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error restarting the program:", err)
	}
	os.Exit(0)
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

func colourText(text, colour string) string {
	return colour + text + colourReset
}
