package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	fmt.Println("Hello Goshix")

	reader := bufio.NewReader(os.Stdin)
	for {
		// input indicator
		fmt.Print("> ")

		// read stdin (keyboard)
		input, err := reader.ReadString('\n')
		if err != nil {	
			fmt.Fprintln(os.Stderr, err)
		}

		// execution of input
		if err = execInput(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func execInput(input string) error {
	// remove newline char
	input = strings.TrimSuffix(input, "\n")

	// separate command and arguments
	args := strings.Split(input, " ")

	// prepare command to execute, pass program and arguments separate
	cmd := exec.Command(args[0], args[1:]...)

	// correct output
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// execute command
	return cmd.Run()
}
