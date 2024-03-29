package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	generatePrompt()
	for scanner.Scan() {
		// Handle the execution of the input
		if err := execInput(scanner.Text()); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		generatePrompt()
	}
}

func generatePrompt() {
	hostname, _ := os.Hostname()
	currentUser, _ := user.Current()
	currentDirectory, _ := os.Getwd()
	splitCurrentDirectory := strings.Split(currentDirectory, "/")
	fmt.Printf("%s|%s|%s> ", hostname, currentUser.Name, splitCurrentDirectory[len(splitCurrentDirectory) - 1])
}

func execInput(input string) error {
	// Remove the newline character
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate teh command and the arguments
	args := strings.Split(input, " ")

	// Check for built-in commands
	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path
		if len(args) < 2 {
			return os.Chdir("/")
		}
		// Change the directory and return the error
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// Pass the program and the arguments separately
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error
	return cmd.Run()
}