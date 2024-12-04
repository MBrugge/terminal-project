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
	helpCommand     = "help"
	whereAmICommand = "whereami"
	whoAmICommand   = "whoami"
	videoCommand    = "video"
	githubCommand   = "github"
	resetCommand    = "reset"
	exitCommand     = "exit"

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
		printText(colourText("help - You know what this command does.", colourYellow))
		printText(colourText("whereami - Where am I right now?", colourYellow))
		printText(colourText("whoami - Who am I?", colourYellow))
		printText(colourText("video - What could this be?", colourYellow))
		printText(colourText("github - Go to my github", colourYellow))
		printText(colourText("reset - Reset the program", colourYellow))
		printText(colourText("exit - Exit the program", colourYellow))

	case whereAmICommand:
		printText(colourText("You are presumably sitting, or standing, behind a device which has a terminal. You probably got here after finding my project and being intrigued. Whether this was curiosity or a want to find out more about me I cannot tell. I will welcome you all the same.", colourYellow))

	case whoAmICommand:
		printText(colourText("That is something you find out every single day.", colourYellow))

	case videoCommand:
		printText(colourText("Opening video...", colourCyan))
		openLink("https://www.youtube.com/watch?v=dQw4w9WgXcQ")

	case githubCommand:
		printText(colourText("Opening github...", colourCyan))
		openLink("https://github.com/MBrugge")

	case resetCommand:
		printText(colourText("Restarting the program", colourRed))
		restartProgram()

	case exitCommand:
		printText(colourText("Goodbye", colourRed))
		os.Exit(0)

	default:
		printText(colourText("Unknown command: ", colourRed) + input)
	}
}

func restartProgram() {
	cmd := exec.Command(os.Args[0])
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	err := cmd.Run()
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
	lineLength := 80
	currentLength := 0

	for _, char := range text {
		fmt.Print(string(char))
		currentLength++
		if currentLength >= lineLength && char == ' ' {
			fmt.Println()
			currentLength = 0
		}
		time.Sleep(25 * time.Millisecond)
	}
	fmt.Println()
}

func openLink(url string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		cmd = exec.Command("xdg-open", url)
	}
	err := cmd.Start()
	if err != nil {
		printText(colourText("Error opening link: "+err.Error(), colourRed))
	}
}
