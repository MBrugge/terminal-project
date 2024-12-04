package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
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
	printText(colourText("Welcome to the terminal!", colourCyan))
	printText(colourText("Type 'help' to see the available commands.", colourYellow))

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			printText("Error reading input: " + err.Error())
			continue
		}
		input = strings.TrimSpace(input)
		handleCommand(input)
	}
}

func handleCommand(input string) {
	switch input {
	case helpCommand:
		printText(colourText("Available commands:", colourCyan))
		printText(colourText("help - Show this help message", colourYellow))
		printText(colourText("exit - Exit the program", colourYellow))
		printText(colourText("reset - Reset the program", colourYellow))

	case exitCommand:
		printText(colourText("Goodbye", colourRed))
		os.Exit(0)

	case resetCommand:
		printText(colourText("Restarting the program", colourRed))
		restartProgram()

	default:
		printText(colourText("Unknown command: ", colourRed) + input)
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

func printText(text string) {
	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(25 * time.Millisecond)
	}
	fmt.Println()
}
