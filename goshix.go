package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// errors
var ErrNoPath = errors.New("path required")

func parser(input string) []string {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Split(input, " ")
	return args 
}

func executor(args []string) error {
	// built-in commands
	switch args[0] {
	// cd command
	case "cd":
	 	if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])
	// exit command
	case "exit":
		os.Exit(0)
	}

	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func main() {
	fmt.Println("Welcome to Goshix!")

	reader := bufio.NewReader(os.Stdin)

	prompt := "Goshix: > "

	for {
		fmt.Print(prompt)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		args := parser(input)
		if err = executor(args); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
